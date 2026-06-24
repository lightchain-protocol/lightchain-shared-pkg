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
