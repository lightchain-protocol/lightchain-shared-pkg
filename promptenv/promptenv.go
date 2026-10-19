// Package promptenv decodes the JSON prompt envelope carried inside a job's
// encrypted prompt blob. The worker serves a job from it and the disputer
// re-runs the job from it; sharing one decoder is what keeps the two from
// giving the model different input.
package promptenv

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/gowebpki/jcs"
	"github.com/lightchain/pkg/searchaug"
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
	// in use, 3 for a self-contained job (see Messages), 4 for a public-output
	// job (see Output). Its presence is
	// what distinguishes an envelope from a prompt that merely happens to be
	// valid JSON. Decode accepts any version >= 1: unknown fields are
	// ignored by encoding/json, so an old worker decoding a v2 payload
	// simply drops the voice fields (forward compatibility) instead of
	// rejecting a prompt it could still answer. Messages are the exception:
	// they are checked whatever the version, and version 3 must carry them.
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
	// jobs in its session. It is text only and its Text stays empty.
	Messages []Message `json:"messages,omitempty"`

	// --- Public output (envelope v4) ---
	//
	// Output is "public" when the answer is committed and published in
	// plaintext as the public response bytes R (package publicresp) instead
	// of encrypted under the session key; empty or "private" is the
	// encrypted path. The prompt stays encrypted either way. An old worker
	// drops this field and encrypts, so the version is not the gate: the
	// PUBLIC_OUTPUT session capability is (draft 04 §3.2).
	Output string `json:"output,omitempty"`
	// Format is a JSON Schema the engine is asked to constrain its output
	// to, passed through as Ollama's `format`. At most MaxFormatBytes in
	// canonical form. Its JCS hash is R.formatHash.
	Format json.RawMessage `json:"format,omitempty"`
	// Match names the dispute rule: "similarity" (the default) or "exact",
	// which requires Output "public" and a Format.
	Match string `json:"match,omitempty"`
	// Params are the inference parameters the job requires. Required when
	// Output is "public": the worker uses exactly these instead of operator
	// overrides and echoes them in R, so a replay can match.
	Params *Params `json:"params,omitempty"`
	// Salt is 32 random bytes as 0x-hex, required when Output is "public",
	// so R.promptHash (keccak of this envelope's bytes) is not guessable
	// for a short prompt.
	Salt string `json:"salt,omitempty"`
}

// Params are the inference parameters a public-output job fixes.
type Params struct {
	NumCtx     int `json:"numCtx"`
	NumPredict int `json:"numPredict"`
}

// Output values.
const (
	OutputPrivate = "private"
	OutputPublic  = "public"
)

// Match values.
const (
	MatchSimilarity = "similarity"
	MatchExact      = "exact"
)

// MaxFormatBytes bounds the canonical size of Format.
const MaxFormatBytes = 16 * 1024

// ErrInvalidPublic marks a public-output envelope that must be refused rather
// than served: a job served under guessed parameters could never be replayed.
var ErrInvalidPublic = errors.New("invalid public-output prompt")

// Public reports whether the answer is to be committed in plaintext.
func (e *Envelope) Public() bool {
	return e.Output == OutputPublic
}

// Exact reports whether the job is disputed by exact match on its output.
func (e *Envelope) Exact() bool {
	return e.Match == MatchExact
}

// CanonicalFormat is the JCS form of Format, or nil without one. It is what
// the engine is sent and what R.formatHash commits to.
func (e *Envelope) CanonicalFormat() (json.RawMessage, error) {
	if len(e.Format) == 0 {
		return nil, nil
	}
	return jcs.Transform(e.Format)
}

// Message is one turn of a conversation: of a self-contained job's
// envelope, or of the chat input Conversation builds for any job.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	// Images ride on the user turn of a multimodal job, never on the wire:
	// a self-contained job is text only.
	Images []string `json:"-"`
}

// selfContainedVersion is the envelope version that carries Messages.
const selfContainedVersion = 3

// ErrInvalidSelfContained marks a self-contained envelope that must be
// refused rather than served: serving it with parts dropped or guessed at
// would answer a different question than the one paid for.
var ErrInvalidSelfContained = errors.New("invalid self-contained prompt")

// SelfContained reports whether the envelope carries its whole conversation.
// The messages alone decide it, whatever the version, so they are never
// dropped.
func (e *Envelope) SelfContained() bool {
	return len(e.Messages) > 0
}

// Conversation is the chat input a job gives the model, or nil when the job
// goes to the single-prompt generate call with its Text. A self-contained job
// always chats with exactly its own messages, even a single user turn, and
// takes no history. Any other job chats when it has history (its session's
// earlier turns, rebuilt with Replay) or images, its Text and images being
// the last user turn.
func (e *Envelope) Conversation(history []Message) []Message {
	if e.SelfContained() {
		return e.Messages
	}
	if len(history) == 0 && len(e.Images) == 0 {
		return nil
	}
	return append(slices.Clip(history), Message{Role: "user", Content: e.Text, Images: e.Images})
}

// Replay is the user turn an earlier job's decrypted prompt adds to the
// history rebuilt for a later job: a search prompt's question, a multimodal
// prompt's text and images, raw text as it is. A prompt that fails to decode
// is replayed raw rather than dropped. ok is false for a self-contained job,
// even one in the search wrapper that a worker refuses: it carries its own
// conversation, so neither it nor its answer belongs in another job's history.
func Replay(plain []byte) (turn Message, ok bool) {
	raw := Message{Role: "user", Content: string(plain)}
	text, _, err := searchaug.DecodePrompt(plain)
	if err != nil {
		return raw, true
	}
	env, err := Decode([]byte(text))
	if err != nil {
		return raw, true
	}
	if env.SelfContained() {
		return Message{}, false
	}
	return Message{Role: "user", Content: env.Text, Images: env.Images}, true
}

// UnwrapSearch takes a decrypted prompt out of its web-search wrapper
// (searchaug.DecodePrompt): the text to search on and decode, and whether the
// job asked for a search. A self-contained envelope inside the wrapper is
// refused with ErrInvalidSelfContained, its text still returned: the job is
// text only, and a search would run on and augment the envelope's JSON
// rather than its conversation.
func UnwrapSearch(plain []byte) (text string, search bool, err error) {
	text, search, err = searchaug.DecodePrompt(plain)
	if err != nil || !search {
		return text, search, err
	}
	if env, decErr := Decode([]byte(text)); env.SelfContained() || errors.Is(decErr, ErrInvalidSelfContained) {
		return text, search, fmt.Errorf("%w: wrapped for a web search", ErrInvalidSelfContained)
	}
	return text, search, nil
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
	// Checked first: images or audio of any size make a self-contained
	// prompt one to refuse outright.
	if env.SelfContained() || env.Version == selfContainedVersion {
		if err := checkSelfContained(&env); err != nil {
			return Envelope{}, err
		}
	}
	if len(env.Images) > maxPromptImages {
		return Envelope{}, fmt.Errorf("%w: %d > %d", errTooManyImages, len(env.Images), maxPromptImages)
	}
	if len(env.Audio) > maxPromptAudioB64 {
		return Envelope{}, fmt.Errorf("%w: %d > %d base64 chars", errAudioTooLarge, len(env.Audio), maxPromptAudioB64)
	}
	if err := checkPublic(&env); err != nil {
		return Envelope{}, err
	}
	return env, nil
}

// checkPublic validates the v4 fields on any envelope that uses them. A
// private envelope with none of them set passes untouched.
func checkPublic(env *Envelope) error {
	switch env.Output {
	case "", OutputPrivate, OutputPublic:
	default:
		return fmt.Errorf("%w: unknown output %q", ErrInvalidPublic, env.Output)
	}
	switch env.Match {
	case "", MatchSimilarity:
	case MatchExact:
		if !env.Public() {
			return fmt.Errorf("%w: exact match needs public output", ErrInvalidPublic)
		}
		if len(env.Format) == 0 {
			return fmt.Errorf("%w: exact match needs a format", ErrInvalidPublic)
		}
	default:
		return fmt.Errorf("%w: unknown match %q", ErrInvalidPublic, env.Match)
	}
	if len(env.Format) > 0 {
		canon, err := jcs.Transform(env.Format)
		if err != nil {
			return fmt.Errorf("%w: format is not JSON: %w", ErrInvalidPublic, err)
		}
		if len(canon) == 0 || canon[0] != '{' {
			return fmt.Errorf("%w: format is not a JSON object", ErrInvalidPublic)
		}
		if len(canon) > MaxFormatBytes {
			return fmt.Errorf("%w: format is %d canonical bytes, max %d", ErrInvalidPublic, len(canon), MaxFormatBytes)
		}
	}
	if !env.Public() {
		return nil
	}
	switch {
	case env.Params == nil:
		return fmt.Errorf("%w: params are required", ErrInvalidPublic)
	case env.Params.NumCtx <= 0 || env.Params.NumPredict <= 0:
		return fmt.Errorf("%w: params.numCtx and params.numPredict must be positive", ErrInvalidPublic)
	case !isSalt(env.Salt):
		return fmt.Errorf("%w: salt must be 32 bytes as 0x-hex", ErrInvalidPublic)
	case env.Audio != "" || env.AudioResponse:
		return fmt.Errorf("%w: voice is not supported with public output", ErrInvalidPublic)
	}
	return nil
}

func isSalt(s string) bool {
	if len(s) != 66 || s[:2] != "0x" {
		return false
	}
	_, err := hex.DecodeString(s[2:])
	return err == nil
}

func checkSelfContained(env *Envelope) error {
	switch {
	case len(env.Images) > 0 || env.Audio != "":
		return fmt.Errorf("%w: carries images or audio", ErrInvalidSelfContained)
	case env.Text != "":
		return fmt.Errorf("%w: carries text beside its messages", ErrInvalidSelfContained)
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
