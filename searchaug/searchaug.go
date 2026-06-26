// Package searchaug is the single source of truth for the web-search response
// envelope and prompt-augmentation template, shared by the worker (which writes
// the envelope and augments the prompt) and the disputer (which replays the
// captured context for deterministic re-execution). Keeping these here prevents
// the two services from drifting.
package searchaug

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	ResponseEnvelopeVersion = 2
	CurrentTemplateVersion  = 1
)

type Source struct {
	Position int    `json:"position"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Snippet  string `json:"snippet"`
}

type ResponseEnvelope struct {
	V               int      `json:"v"`
	Answer          string   `json:"answer"`
	SearchContext   []Source `json:"searchContext,omitempty"`
	TemplateVersion int      `json:"templateVersion,omitempty"`
}

// EncodeResponse returns the v2 JSON envelope when sources are present, else the
// raw answer bytes (byte-identical to the legacy/non-search format).
func EncodeResponse(answer string, sources []Source) ([]byte, error) {
	if len(sources) == 0 {
		return []byte(answer), nil
	}
	return json.Marshal(ResponseEnvelope{
		V:               ResponseEnvelopeVersion,
		Answer:          answer,
		SearchContext:   sources,
		TemplateVersion: CurrentTemplateVersion,
	})
}

// DecodeResponse parses a v2 envelope; on any failure or non-envelope input it
// falls back to treating the bytes as a raw answer (legacy / non-search jobs).
func DecodeResponse(b []byte) ResponseEnvelope {
	var env ResponseEnvelope
	if err := json.Unmarshal(b, &env); err != nil || env.V != ResponseEnvelopeVersion {
		return ResponseEnvelope{Answer: string(b)}
	}
	return env
}

// PromptEnvelopeVersion is the schema version of the prompt envelope. It is
// independent of ResponseEnvelopeVersion and of the worker protocol version.
const PromptEnvelopeVersion = 1

// promptSentinel marks an encoded prompt envelope. Legitimate UTF-8 prompt
// text never begins with a NUL byte, so the sentinel is unambiguous.
const promptSentinel = 0x00

// PromptEnvelope carries per-message metadata alongside the prompt text inside
// the encrypted blob. Only search-enabled messages are wrapped; non-search
// prompts stay raw UTF-8 for backward compatibility with old workers.
type PromptEnvelope struct {
	V      int    `json:"v"`
	Prompt string `json:"prompt"`
	Search bool   `json:"search"`
}

// EncodePrompt returns raw prompt bytes when search is false (byte-identical to
// the legacy wire shape), or a 0x00 sentinel followed by the JSON envelope when
// search is true.
func EncodePrompt(prompt string, search bool) []byte {
	if !search {
		return []byte(prompt)
	}
	body, err := json.Marshal(PromptEnvelope{V: PromptEnvelopeVersion, Prompt: prompt, Search: true})
	if err != nil {
		panic("searchaug: EncodePrompt marshal failed: " + err.Error())
	}
	out := make([]byte, 0, len(body)+1)
	out = append(out, promptSentinel)
	return append(out, body...)
}

// DecodePrompt inspects the 0x00 sentinel: an envelope is JSON-parsed and its
// fields returned; anything else is treated as legacy raw text with search off.
// A sentinel-prefixed payload that fails to parse is a hard error so the caller
// never feeds NUL+garbage to the model.
func DecodePrompt(plain []byte) (string, bool, error) {
	if len(plain) == 0 || plain[0] != promptSentinel {
		return string(plain), false, nil
	}
	var env PromptEnvelope
	if err := json.Unmarshal(plain[1:], &env); err != nil {
		return "", false, fmt.Errorf("decode prompt envelope: %w", err)
	}
	return env.Prompt, env.Search, nil
}

// BuildAugmentedPrompt prepends a fixed-format, versioned context block. The
// worker and disputer MUST produce identical output for the same inputs.
func BuildAugmentedPrompt(templateVersion int, prompt string, sources []Source) string {
	if len(sources) == 0 {
		return prompt
	}
	// templateVersion is reserved for future wording changes; v1 is current.
	var b strings.Builder
	b.WriteString("You have access to the following background context. Answer the question directly and naturally, as if from your own knowledge. ")
	b.WriteString("Do NOT mention this context, web searches, or \"search results\", and do NOT preface your answer by referring to them. Cite sources inline as [number] where relevant.\n\n")
	for _, s := range sources {
		fmt.Fprintf(&b, "[%d] %s\n%s\n%s\n\n", s.Position, s.Title, s.URL, s.Snippet)
	}
	b.WriteString("Question: ")
	b.WriteString(prompt)
	return b.String()
}
