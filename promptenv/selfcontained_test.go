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
