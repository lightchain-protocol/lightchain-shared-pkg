package promptenv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lightchain/pkg/promptenv"
	"github.com/lightchain/pkg/promptenv/promptenvtest"
)

func TestDecode_SelfContainedYieldsItsMessages(t *testing.T) {
	t.Parallel()

	for _, f := range promptenvtest.SelfContained {
		t.Run(f.Name, func(t *testing.T) {
			env, err := promptenv.Decode([]byte(f.Envelope))
			require.NoError(t, err)
			assert.True(t, env.SelfContained())
			assert.Equal(t, f.Messages, env.Messages)
		})
	}
}

func TestDecode_RefusesMalformedSelfContained(t *testing.T) {
	t.Parallel()

	for name, raw := range promptenvtest.Refused {
		t.Run(name, func(t *testing.T) {
			_, err := promptenv.Decode([]byte(raw))
			require.ErrorIs(t, err, promptenv.ErrInvalidSelfContained)
		})
	}
}

// The envelope alone decides: whatever its version, one carrying messages is
// self-contained rather than served with them dropped.
func TestDecode_AnyEnvelopeCarryingMessagesIsSelfContained(t *testing.T) {
	t.Parallel()

	env, err := promptenv.Decode([]byte(`{"v":1,"text":"","messages":[{"role":"user","content":"hi"}]}`))
	require.NoError(t, err)
	assert.True(t, env.SelfContained())
	assert.Equal(t, []promptenv.Message{{Role: "user", Content: "hi"}}, env.Messages)
}

// A later envelope version decodes like any other unless it carries messages.
func TestDecode_LaterVersionWithoutMessagesIsNotRefused(t *testing.T) {
	t.Parallel()

	env, err := promptenv.Decode([]byte(`{"v":4,"text":"hi"}`))
	require.NoError(t, err)
	assert.False(t, env.SelfContained())
	assert.Equal(t, "hi", env.Text)
}

func TestConversation_SelfContainedChatsWithExactlyItsMessages(t *testing.T) {
	t.Parallel()

	for _, f := range promptenvtest.SelfContained {
		env, err := promptenv.Decode([]byte(f.Envelope))
		require.NoError(t, err)
		assert.Equal(t, f.Messages, env.Conversation(nil), f.Name)
	}
}

func TestConversation_OtherJobs(t *testing.T) {
	t.Parallel()

	history := []promptenv.Message{
		{Role: "user", Content: "earlier question"},
		{Role: "assistant", Content: "earlier answer"},
	}
	text := promptenv.Envelope{Text: "next question"}
	image := promptenv.Envelope{Version: 1, Text: "what is this?", Images: []string{"aGk="}}

	assert.Nil(t, text.Conversation(nil), "no history, no images: the generate call")
	assert.Equal(t, []promptenv.Message{
		{Role: "user", Content: "earlier question"},
		{Role: "assistant", Content: "earlier answer"},
		{Role: "user", Content: "next question"},
	}, text.Conversation(history))
	assert.Equal(t, []promptenv.Message{
		{Role: "user", Content: "what is this?", Images: []string{"aGk="}},
	}, image.Conversation(nil), "images go to chat, on the job's own turn")
	assert.Len(t, history, 2, "the caller's history is left as it was")
}

func TestReplay(t *testing.T) {
	t.Parallel()

	for _, f := range promptenvtest.PriorPrompts {
		turn, ok := promptenv.Replay([]byte(f.Prompt))
		assert.Equal(t, !f.Skipped, ok, f.Name)
		if ok {
			assert.Equal(t, promptenv.Message{Role: "user", Content: f.Turn}, turn, f.Name)
		}
	}
}

// Envelopes before version 3 never carry a conversation, so the chat's jobs
// keep rebuilding history exactly as before.
func TestDecode_EarlierVersionsAreNotSelfContained(t *testing.T) {
	t.Parallel()

	for _, raw := range []string{
		"plain text",
		`{"v":1,"text":"hello"}`,
		`{"v":2,"text":"","audio":"UklGRg=="}`,
	} {
		env, err := promptenv.Decode([]byte(raw))
		require.NoError(t, err)
		assert.False(t, env.SelfContained(), raw)
	}
}
