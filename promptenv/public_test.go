package promptenv

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const salt = "0x0102030405060708091011121314151617181920212223242526272829303132"

func publicEnvelope(extra string) string {
	return `{"v":4,"text":"","messages":[{"role":"user","content":"Summarise this."}],` +
		`"output":"public","params":{"numCtx":8192,"numPredict":512},"salt":"` + salt + `"` + extra + `}`
}

func TestDecode_V4PublicEnvelope(t *testing.T) {
	t.Parallel()

	env, err := Decode([]byte(publicEnvelope(`,"format":{"type":"object","properties":{"summary":{"type":"string"}}},"match":"exact"`)))
	require.NoError(t, err)
	assert.Equal(t, 4, env.Version)
	assert.True(t, env.Public())
	assert.True(t, env.Exact())
	assert.True(t, env.SelfContained(), "v4 is v3 plus the public fields")
	require.NotNil(t, env.Params)
	assert.Equal(t, 8192, env.Params.NumCtx)
	assert.Equal(t, 512, env.Params.NumPredict)
	assert.Equal(t, salt, env.Salt)

	canon, err := env.CanonicalFormat()
	require.NoError(t, err)
	assert.Equal(t, `{"properties":{"summary":{"type":"string"}},"type":"object"}`, string(canon))
}

// A public job without a format is a plain text answer committed in the
// clear: allowed, similarity-disputed.
func TestDecode_V4PublicWithoutFormat(t *testing.T) {
	t.Parallel()

	env, err := Decode([]byte(publicEnvelope("")))
	require.NoError(t, err)
	assert.True(t, env.Public())
	assert.False(t, env.Exact())
	canon, err := env.CanonicalFormat()
	require.NoError(t, err)
	assert.Nil(t, canon)
}

// An old worker reads v4 as v3: it answers and encrypts. The capability bit
// on the session, not the version, keeps such a worker away from public jobs
// (draft 04 §3.2). This pins that Decode on an envelope with the public
// fields removed still succeeds, which is what an old binary effectively does.
func TestDecode_V4MinusPublicFieldsIsAValidV3(t *testing.T) {
	t.Parallel()

	env, err := Decode([]byte(`{"v":4,"text":"","messages":[{"role":"user","content":"hi"}]}`))
	require.NoError(t, err)
	assert.False(t, env.Public())
	assert.True(t, env.SelfContained())
}

// A private envelope is untouched by the v4 checks.
func TestDecode_PrivateEnvelopeIgnoresPublicRules(t *testing.T) {
	t.Parallel()

	for _, raw := range []string{
		`{"v":1,"text":"hello"}`,
		`{"v":3,"text":"","messages":[{"role":"user","content":"hi"}],"output":"private"}`,
		`{"v":3,"text":"","messages":[{"role":"user","content":"hi"}],"format":{"type":"object"}}`,
	} {
		env, err := Decode([]byte(raw))
		require.NoError(t, err, raw)
		assert.False(t, env.Public())
	}
}

// A format that is not JSON cannot occur: the envelope itself would not parse
// and Decode would hand the bytes back as plain text, which is the existing
// fallback for anything that is not an envelope.
func TestDecode_V4Refused(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"unknown output": `{"v":4,"text":"hi","output":"plaintext"}`,
		"unknown match":  `{"v":4,"text":"hi","match":"fuzzy"}`,
		"exact without public": `{"v":4,"text":"","messages":[{"role":"user","content":"hi"}],` +
			`"format":{"type":"object"},"match":"exact"}`,
		"exact without format": publicEnvelope(`,"match":"exact"`),
		"format not an object": publicEnvelope(`,"format":["type"]`),
		"format too large":     publicEnvelope(`,"format":{"description":"` + strings.Repeat("x", MaxFormatBytes) + `"}`),
		"no params": `{"v":4,"text":"","messages":[{"role":"user","content":"hi"}],` +
			`"output":"public","salt":"` + salt + `"}`,
		"zero numPredict": `{"v":4,"text":"","messages":[{"role":"user","content":"hi"}],` +
			`"output":"public","params":{"numCtx":8192,"numPredict":0},"salt":"` + salt + `"}`,
		"no salt": `{"v":4,"text":"","messages":[{"role":"user","content":"hi"}],` +
			`"output":"public","params":{"numCtx":8192,"numPredict":512}}`,
		"short salt": `{"v":4,"text":"","messages":[{"role":"user","content":"hi"}],` +
			`"output":"public","params":{"numCtx":8192,"numPredict":512},"salt":"0x0102"}`,
		"voice with public": `{"v":4,"text":"","audio":"UklGRg==","output":"public",` +
			`"params":{"numCtx":8192,"numPredict":512},"salt":"` + salt + `"}`,
	}
	for name, raw := range cases {
		_, err := Decode([]byte(raw))
		require.Error(t, err, name)
		assert.ErrorIs(t, err, ErrInvalidPublic, name)
	}
}

// Public output is not the self-contained flag: a v4 single-turn text prompt
// without messages is also fine, so the gateway can use either shape.
func TestDecode_V4PublicTextPrompt(t *testing.T) {
	t.Parallel()

	env, err := Decode([]byte(`{"v":4,"text":"What is 2+2?","output":"public",` +
		`"params":{"numCtx":2048,"numPredict":64},"salt":"` + salt + `"}`))
	require.NoError(t, err)
	assert.True(t, env.Public())
	assert.False(t, env.SelfContained())
	assert.Equal(t, "What is 2+2?", env.Text)
}
