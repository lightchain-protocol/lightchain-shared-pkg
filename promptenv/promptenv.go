// Package promptenv decodes the JSON prompt envelope carried inside a job's
// encrypted prompt blob. The worker serves a job from it and the disputer
// re-runs the job from it; sharing one decoder is what keeps the two from
// giving the model different input.
package promptenv

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Envelope is the decrypted prompt payload.
//
// Prompts used to be raw UTF-8 text, which meant a deployed vision model had
// no way to receive an image: the blob carried a bare string and the worker
// passed it straight through. The envelope adds a version and an image list
// while staying backward compatible - a payload that is not a versioned
// envelope is treated as text, so older clients keep working unchanged.
type Envelope struct {
	// Version is 1 for the first envelope format, 2 once any voice field is
	// in use, 3 for a self-contained job (see Messages). Its presence is what distinguishes an envelope from a prompt
	// that merely happens to be valid JSON. Decode accepts any version >= 1:
	// unknown fields are ignored by encoding/json, so an old worker decoding
	// a v2 payload simply drops the voice fields (forward compatibility)
	// instead of rejecting a prompt it could still answer.
	Version int    `json:"v"`
	Text    string `json:"text"`
	// Images are base64-encoded, without a data: prefix, in the form Ollama
	// expects. They ride on the final user turn.
	Images []string `json:"images,omitempty"`

	// --- Voice fields (envelope v2) ---
	//
	// Audio is a base64-encoded audio clip (no data: prefix) carrying a
	// spoken prompt. The worker transcribes it via the STT sidecar and
	// merges the transcript into Text before inference. AudioFormat names
	// the container/codec ("wav", "mp3", ...) so the transcriber can pick a
	// decoder; empty defaults to "wav".
	//
	// The audio bytes stay inside the encrypted prompt blob on DA, so the
	// input side of settlement (what the consumer submitted) still covers
	// the voice prompt bit-for-bit. The *transcript* is derived context,
	// not settled content: it is re-derivable from the blob but not
	// bit-deterministic across whisper versions, so the response hash
	// continues to cover only the model's output text until the non-text
	// settlement commitment lands (contract-side scope, tracked
	// separately).
	Audio       string `json:"audio,omitempty"`
	AudioFormat string `json:"audioFormat,omitempty"`
	// AudioResponse opts in to spoken output: the worker synthesizes the
	// response text via the TTS sidecar and delivers the audio by
	// reference (encrypted DA blob + `audio` frame descriptor). The audio
	// is deliberately NOT part of the settlement ciphertext - the
	// text-only settlement invariant is unchanged.
	AudioResponse bool `json:"audioResponse,omitempty"`
	// Voice selects the TTS voice (sidecar-specific name, e.g. a Kokoro
	// voice id). Empty falls back to the worker's configured default.
	Voice string `json:"voice,omitempty"`

	// Messages is the whole conversation of a self-contained job (envelope
	// v3), the last turn the user's. Such a job goes to the model's chat
	// call with exactly these messages and no history rebuilt from earlier
	// jobs in its session. Text stays empty; it is text only.
	Messages []Message `json:"messages,omitempty"`
}

// Message is one turn of a self-contained job's conversation.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// SelfContainedVersion is the envelope version that carries Messages.
const SelfContainedVersion = 3

// ErrInvalidSelfContained marks a self-contained envelope that must be
// refused rather than served: serving it with parts dropped or guessed at
// would answer a different question than the one paid for.
var ErrInvalidSelfContained = errors.New("invalid self-contained prompt")

// SelfContained reports whether the envelope carries its whole conversation.
func (e *Envelope) SelfContained() bool {
	return e.Version >= SelfContainedVersion
}

// maxPromptImages bounds how many images one prompt may carry. The prompt
// blob is capped at 126,972 bytes by EIP-4844 encoding, so this is a guard
// against a pathological payload rather than the real limit - the blob
// encoder rejects anything oversized long before this matters.
const maxPromptImages = 8

// maxPromptAudioB64 bounds the base64 audio payload one prompt may carry
// (~96 KiB decoded). Same rationale as maxPromptImages: the blob cap is the
// real limit, this rejects pathological envelopes early and keeps an audio
// prompt from crowding out the response blob budget unnoticed.
const maxPromptAudioB64 = 131072

// errTooManyImages is returned rather than silently truncating, because a
// dropped image changes the answer the consumer paid for.
var errTooManyImages = errors.New("prompt carries too many images")

// errAudioTooLarge rejects an oversized audio payload rather than silently
// truncating it, for the same reason: a clipped clip transcribes to
// different words than the consumer spoke.
var errAudioTooLarge = errors.New("prompt audio exceeds size limit")

// Decode interprets a decrypted prompt payload.
//
// Anything that is not a well-formed versioned envelope is returned as plain
// text. That fallback is deliberate and load-bearing: a user prompt can
// legitimately be a JSON document, and misreading one as an envelope would
// silently drop the actual question.
func Decode(raw []byte) (Envelope, error) {
	var env Envelope
	if err := json.Unmarshal(raw, &env); err != nil || env.Version < 1 {
		return Envelope{Text: string(raw)}, nil
	}
	if len(env.Images) > maxPromptImages {
		return Envelope{}, fmt.Errorf("%w: %d > %d", errTooManyImages, len(env.Images), maxPromptImages)
	}
	if len(env.Audio) > maxPromptAudioB64 {
		return Envelope{}, fmt.Errorf("%w: %d > %d base64 chars", errAudioTooLarge, len(env.Audio), maxPromptAudioB64)
	}
	if env.SelfContained() {
		if err := checkSelfContained(&env); err != nil {
			return Envelope{}, err
		}
	}
	return env, nil
}

func checkSelfContained(env *Envelope) error {
	switch {
	case len(env.Images) > 0 || env.Audio != "":
		return fmt.Errorf("%w: carries images or audio", ErrInvalidSelfContained)
	case len(env.Messages) == 0:
		return fmt.Errorf("%w: no messages", ErrInvalidSelfContained)
	case env.Messages[len(env.Messages)-1].Role != "user":
		return fmt.Errorf("%w: last message is not the user's", ErrInvalidSelfContained)
	}
	for i, m := range env.Messages {
		switch m.Role {
		case "system", "user", "assistant":
		default:
			return fmt.Errorf("%w: message %d has unknown role %q", ErrInvalidSelfContained, i, m.Role)
		}
	}
	return nil
}

// Bytes reports the plaintext size for logging, counting image and audio
// data because that is what dominates a multimodal prompt, and the messages
// of a self-contained one.
func (e *Envelope) Bytes() int {
	n := len(e.Text) + len(e.Audio)
	for _, img := range e.Images {
		n += len(img)
	}
	for _, m := range e.Messages {
		n += len(m.Content)
	}
	return n
}
