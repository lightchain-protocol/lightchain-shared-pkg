// Package promptenvtest holds the prompt envelopes and sessions that the
// worker and disputer handler tests both run. Each side proves it hands the
// model a fixture's messages, so together they prove both sides give the
// model the same input.
package promptenvtest

import (
	"fmt"

	"github.com/lightchain/pkg/promptenv"
	"github.com/lightchain/pkg/searchaug"
	"github.com/lightchain/pkg/sessionhistory"
)

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

// PriorPrompt is an earlier job's prompt in a session and the user turn it
// adds to the history rebuilt for a later chat job. Skipped marks a
// self-contained job, which adds nothing, its answer included.
type PriorPrompt struct {
	Name    string
	Prompt  string
	Turn    string
	Skipped bool
}

// PriorPrompts covers every prompt shape a session's earlier job can have.
var PriorPrompts = []PriorPrompt{
	{Name: "raw text", Prompt: "plain question", Turn: "plain question"},
	{Name: "search wrapped", Prompt: string(searchaug.EncodePrompt("searched question", true)), Turn: "searched question"},
	{Name: "multimodal envelope", Prompt: `{"v":1,"text":"envelope question"}`, Turn: "envelope question"},
	{Name: "self-contained", Prompt: SelfContained[1].Envelope, Skipped: true},
}

// SessionJob is a job as the chain holds it: its JobSubmitted event, and the
// prompt and answer in its blobs.
type SessionJob struct {
	ID, SessionID, SubmitBlock uint64
	Prompt, Answer             string
}

// Session is a chat job, Current, that a worker serves and the disputer later
// re-runs, and the other jobs on chain around it. History is the
// conversation both must rebuild for Current.
type Session struct {
	Current SessionJob
	Chain   []SessionJob
	History []promptenv.Message
}

// JobSubmitted returns the ids of sessionID's jobs, Current included, whose
// JobSubmitted event lies in blocks [fromBlock, toBlock]: what a node answers
// a history lookup with.
func (s *Session) JobSubmitted(sessionID, fromBlock, toBlock uint64) []uint64 {
	var ids []uint64
	for _, j := range append([]SessionJob{s.Current}, s.Chain...) {
		if j.SessionID == sessionID && j.SubmitBlock >= fromBlock && j.SubmitBlock <= toBlock {
			ids = append(ids, j.ID)
		}
	}
	return ids
}

// BusySession is a session on a busy chain. Current's earlier jobs, one per
// PriorPrompts shape, all sit more than 50 job ids behind it, other sessions'
// jobs between. The chain also holds session jobs neither side may replay:
// one just before the look-back window, and two after Current, one of them in
// Current's own block.
var BusySession = func() Session {
	const session = 7
	s := Session{Current: SessionJob{ID: 200, SessionID: session, SubmitBlock: 120000, Prompt: "next question"}}
	windowStart := s.Current.SubmitBlock - sessionhistory.LookbackBlocks
	for i, p := range PriorPrompts {
		j := SessionJob{
			ID: 20 + 40*uint64(i), SessionID: session, SubmitBlock: windowStart + 10000*uint64(i),
			Prompt: p.Prompt, Answer: fmt.Sprintf("answer %d", i+1),
		}
		s.Chain = append(s.Chain, j)
		if !p.Skipped {
			s.History = append(s.History,
				promptenv.Message{Role: "user", Content: p.Turn},
				promptenv.Message{Role: "assistant", Content: j.Answer})
		}
	}
	s.Chain = append(
		s.Chain,
		SessionJob{ID: 5, SessionID: session, SubmitBlock: windowStart - 1, Prompt: "too old", Answer: "too old to replay"},
		SessionJob{ID: 150, SessionID: 8, SubmitBlock: 110000, Prompt: "another session", Answer: "not this one"},
		SessionJob{ID: 199, SessionID: 9, SubmitBlock: s.Current.SubmitBlock, Prompt: "a third session", Answer: "not this one"},
		SessionJob{ID: 201, SessionID: session, SubmitBlock: s.Current.SubmitBlock, Prompt: "same block, later", Answer: "not yet"},
		SessionJob{ID: 202, SessionID: session, SubmitBlock: s.Current.SubmitBlock + 1, Prompt: "later", Answer: "not yet"},
	)
	return s
}()
