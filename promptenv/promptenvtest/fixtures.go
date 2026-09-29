// Package promptenvtest holds the self-contained prompt envelopes that the
// worker and disputer handler tests both run. Each side proves it hands the
// model a fixture's messages, so together they prove both sides give the
// model the same input.
package promptenvtest

import "github.com/lightchain/pkg/promptenv"

// Fixture is a self-contained envelope as it travels inside the encrypted
// prompt blob, and the messages the model must be given for it.
type Fixture struct {
	Name     string
	Envelope string
	Messages []promptenv.Message
}

// SelfContained are well-formed version-3 envelopes.
var SelfContained = []Fixture{
	{
		Name:     "single user turn",
		Envelope: `{"v":3,"text":"","messages":[{"role":"user","content":"What is 2+2?"}]}`,
		Messages: []promptenv.Message{{Role: "user", Content: "What is 2+2?"}},
	},
	{
		Name: "system message and a multi-turn conversation",
		Envelope: `{"v":3,"text":"","messages":[` +
			`{"role":"system","content":"Answer in one word."},` +
			`{"role":"user","content":"Capital of France?"},` +
			`{"role":"assistant","content":"Paris"},` +
			`{"role":"user","content":"And of \"Italy\"?\nOne word."}]}`,
		Messages: []promptenv.Message{
			{Role: "system", Content: "Answer in one word."},
			{Role: "user", Content: "Capital of France?"},
			{Role: "assistant", Content: "Paris"},
			{Role: "user", Content: "And of \"Italy\"?\nOne word."},
		},
	},
}

// Refused are version-3 envelopes a worker must refuse rather than serve,
// keyed by what is wrong with them.
var Refused = map[string]string{
	"carries images": `{"v":3,"text":"","messages":[{"role":"user","content":"what is this?"}],"images":["aGk="]}`,
	"carries audio":  `{"v":3,"text":"","messages":[{"role":"user","content":"transcribe"}],"audio":"UklGRg=="}`,
	"carries too many images": `{"v":3,"text":"","messages":[{"role":"user","content":"what are these?"}],` +
		`"images":["aGk=","aGk=","aGk=","aGk=","aGk=","aGk=","aGk=","aGk=","aGk="]}`,
	"carries text":          `{"v":3,"text":"hello","messages":[{"role":"user","content":"hello"}]}`,
	"no messages":           `{"v":3,"text":""}`,
	"last message not user": `{"v":3,"text":"","messages":[{"role":"user","content":"hi"},{"role":"assistant","content":"hello"}]}`,
	"unknown role":          `{"v":3,"text":"","messages":[{"role":"developer","content":"be brief"},{"role":"user","content":"hi"}]}`,
}
