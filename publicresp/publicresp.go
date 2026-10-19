// Package publicresp defines the public response bytes R that a worker
// commits for a public-output job, and the hashes every party derives from
// them. The worker encodes R and commits keccak256(R) as the job's
// responseCiphertextHash; the disputer decodes R, replays the job and
// compares; a contract (AgentContract.deliver) checks keccak256(R) against
// the commitment and hands R to its handler. One encoder here is what makes
// those three agree byte for byte.
//
// R is JSON canonicalised with RFC 8785 (JCS), so the same object always
// gives the same bytes whichever side produced it. See
// lightchain-agents/docs/plans/drafts/04-worker-disputer.md §3.3 and
// 03-manifest-and-actions.md §5.1.
package publicresp

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gowebpki/jcs"
)

// Version is the R format this package writes. Decode accepts only this one:
// R is a commitment, and a party that does not understand a version must not
// guess at it.
const Version = 1

// ZeroHash is the formatHash of a job without a format, and the value of any
// absent hash field.
const ZeroHash = "0x0000000000000000000000000000000000000000000000000000000000000000"

// Params are the inference parameters the worker actually sent to the engine.
// They are echoed so the disputer can replay with the same values and refuse
// a worker that ran the job under different ones (draft 04 §4.2).
type Params struct {
	Model       string  `json:"model"`
	NumCtx      int     `json:"numCtx"`
	NumPredict  int     `json:"numPredict"`
	Seed        int     `json:"seed"`
	Temperature float64 `json:"temperature"`
}

// Response is R before canonicalisation.
//
// Exactly one of Output and Answer is set: Output when the job had a format
// and the engine's text parsed as JSON (Valid says whether it also validated
// against the schema), Answer with the raw text otherwise. Action is set only
// for action tools (draft 03 §5.1) and is derived from Output by a transcoder
// the disputer re-runs; this package carries it and checks its shape, nothing
// more.
type Response struct {
	V          int             `json:"v"`
	PromptHash string          `json:"promptHash"`
	FormatHash string          `json:"formatHash"`
	Params     Params          `json:"params"`
	Valid      bool            `json:"valid"`
	Output     json.RawMessage `json:"output,omitempty"`
	Answer     *string         `json:"answer,omitempty"`
	Action     string          `json:"action,omitempty"`
}

var (
	// ErrNotCanonical marks bytes that are not the JCS form of their own
	// content. A commitment over non-canonical bytes would let two parties
	// disagree about one answer.
	ErrNotCanonical = errors.New("public response is not canonical JSON")
	// ErrInvalid marks a response whose fields are inconsistent.
	ErrInvalid = errors.New("invalid public response")
)

// PromptHash commits to the plaintext envelope exactly as decrypted. The
// envelope carries a salt in public mode, so the hash does not leak short
// prompts.
func PromptHash(plainEnvelope []byte) string {
	return hexHash(crypto.Keccak256(plainEnvelope))
}

// FormatHash is keccak256 of the JCS form of the job's format (a JSON Schema),
// or ZeroHash when the job had none. The profile, the envelope and R all
// carry the same value for the same schema whatever whitespace or key order
// each used.
func FormatHash(format json.RawMessage) (string, error) {
	if len(bytes.TrimSpace(format)) == 0 {
		return ZeroHash, nil
	}
	canon, err := jcs.Transform(format)
	if err != nil {
		return "", fmt.Errorf("format: %w", err)
	}
	return hexHash(crypto.Keccak256(canon)), nil
}

// Encode validates r and returns its canonical bytes R. keccak256 of the
// result is what the worker commits on chain.
func Encode(r Response) ([]byte, error) {
	if err := r.check(); err != nil {
		return nil, err
	}
	raw, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return jcs.Transform(raw)
}

// Decode parses R, refusing bytes that are not canonical or not a valid
// version-1 response. The returned Response re-encodes to exactly b.
func Decode(b []byte) (Response, error) {
	canon, err := jcs.Transform(b)
	if err != nil {
		return Response{}, fmt.Errorf("%w: %w", ErrNotCanonical, err)
	}
	if !bytes.Equal(canon, b) {
		return Response{}, ErrNotCanonical
	}
	var r Response
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&r); err != nil {
		return Response{}, fmt.Errorf("%w: %w", ErrInvalid, err)
	}
	if err := r.check(); err != nil {
		return Response{}, err
	}
	return r, nil
}

// FromEngineText turns what the engine returned into R's output fields, the
// same way on the worker and on the disputer. With a format, text that parses
// as a JSON object or array is the output and valid; anything else (a
// generation cut mid-object by numPredict, prose around the JSON) is kept
// verbatim as the answer with valid false, never retried, so a replay can
// reproduce it (draft 04 §4.3). Without a format the text is always the
// answer. Validation against the schema itself is not done here yet: valid
// means "parsed", and handlers must still check the shape they expect.
func FromEngineText(text string, hasFormat bool) (output json.RawMessage, answer *string, valid bool) {
	if hasFormat {
		trimmed := bytes.TrimSpace([]byte(text))
		if len(trimmed) > 0 && (trimmed[0] == '{' || trimmed[0] == '[') && json.Valid(trimmed) {
			return json.RawMessage(trimmed), nil, true
		}
	}
	return nil, &text, false
}

// Hash is keccak256(R): the value committed as responseCiphertextHash for a
// public-output job.
func Hash(encoded []byte) string {
	return hexHash(crypto.Keccak256(encoded))
}

// OutputHash is what the disputer compares between the worker's R and its
// own replay for an exact-match job (draft 04 §5.1 step 3): keccak256 over
// the canonical output (or the answer text) and the valid flag. Key order
// and whitespace inside the output cannot cause a dispute because the output
// is already canonical inside R.
func (r Response) OutputHash() string {
	var body []byte
	if r.Output != nil {
		// Canonicalised here too, so the hash of a Response built from the
		// engine's text equals the hash of the same Response decoded from R.
		if canon, err := jcs.Transform(r.Output); err == nil {
			body = canon
		} else {
			body = r.Output
		}
	} else if r.Answer != nil {
		body = []byte(*r.Answer)
	}
	valid := byte(0)
	if r.Valid {
		valid = 1
	}
	return hexHash(crypto.Keccak256(body, []byte{valid}))
}

func (r Response) check() error {
	switch {
	case r.V != Version:
		return fmt.Errorf("%w: version %d, want %d", ErrInvalid, r.V, Version)
	case !isHash(r.PromptHash):
		return fmt.Errorf("%w: promptHash is not a 32-byte hex hash", ErrInvalid)
	case !isHash(r.FormatHash):
		return fmt.Errorf("%w: formatHash is not a 32-byte hex hash", ErrInvalid)
	case r.Params.Model == "":
		return fmt.Errorf("%w: params.model is empty", ErrInvalid)
	case r.Params.NumCtx <= 0 || r.Params.NumPredict <= 0:
		return fmt.Errorf("%w: params.numCtx and params.numPredict must be positive", ErrInvalid)
	case (r.Output == nil) == (r.Answer == nil):
		return fmt.Errorf("%w: exactly one of output and answer must be set", ErrInvalid)
	case r.Output != nil && !json.Valid(r.Output):
		return fmt.Errorf("%w: output is not JSON", ErrInvalid)
	case r.Answer != nil && r.Valid:
		return fmt.Errorf("%w: a raw answer cannot be valid against a schema", ErrInvalid)
	case r.Action != "" && r.Output == nil:
		return fmt.Errorf("%w: action without output", ErrInvalid)
	case r.Action != "" && !isHexBytes(r.Action):
		return fmt.Errorf("%w: action is not 0x-prefixed hex", ErrInvalid)
	}
	return nil
}

func hexHash(h []byte) string {
	return "0x" + hex.EncodeToString(h)
}

func isHash(s string) bool {
	return len(s) == 66 && isHexBytes(s)
}

func isHexBytes(s string) bool {
	if len(s) < 2 || s[:2] != "0x" || len(s)%2 != 0 {
		return false
	}
	_, err := hex.DecodeString(s[2:])
	return err == nil
}
