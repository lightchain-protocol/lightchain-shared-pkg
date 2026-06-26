package searchaug

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecode_RoundTrip(t *testing.T) {
	srcs := []Source{{1, "T1", "https://a", "s1"}, {2, "T2", "https://b", "s2"}}
	b, err := EncodeResponse("the answer", srcs)
	require.NoError(t, err)
	env := DecodeResponse(b)
	assert.Equal(t, ResponseEnvelopeVersion, env.V)
	assert.Equal(t, "the answer", env.Answer)
	assert.Equal(t, srcs, env.SearchContext)
}

func TestEncode_NoSources_IsRawAnswer(t *testing.T) {
	b, err := EncodeResponse("plain answer", nil)
	require.NoError(t, err)
	assert.Equal(t, "plain answer", string(b)) // byte-identical to legacy
	env := DecodeResponse(b)
	assert.Equal(t, 0, env.V)
	assert.Equal(t, "plain answer", env.Answer)
	assert.Nil(t, env.SearchContext)
}

func TestDecode_LegacyRawBytes(t *testing.T) {
	env := DecodeResponse([]byte("just a worker answer"))
	assert.Equal(t, "just a worker answer", env.Answer)
	assert.Nil(t, env.SearchContext)
}

func TestDecode_V2EnvelopeEmptyAnswerNotMisdecoded(t *testing.T) {
	b, err := json.Marshal(ResponseEnvelope{V: ResponseEnvelopeVersion, Answer: "", TemplateVersion: 1})
	require.NoError(t, err)
	env := DecodeResponse(b)
	assert.Equal(t, ResponseEnvelopeVersion, env.V) // NOT treated as raw
	assert.Equal(t, "", env.Answer)
}

func TestDecode_NonEnvelopeJSONFallsBack(t *testing.T) {
	env := DecodeResponse([]byte(`{"foo":1}`))
	assert.Equal(t, 0, env.V)
	assert.Equal(t, `{"foo":1}`, env.Answer)
}

func TestBuildAugmentedPrompt_Deterministic_NoPreface(t *testing.T) {
	srcs := []Source{{1, "T1", "https://a", "s1"}}
	out := BuildAugmentedPrompt(CurrentTemplateVersion, "original question", srcs)
	assert.Contains(t, out, "original question")
	assert.Contains(t, out, "https://a")
	assert.Contains(t, out, "[1]")
	assert.Contains(t, out, "Do NOT mention this context")
	assert.Equal(t, out, BuildAugmentedPrompt(CurrentTemplateVersion, "original question", srcs))
}

func TestDecodePrompt_LegacyRawText(t *testing.T) {
	prompt, search, err := DecodePrompt([]byte("hello world"))
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if prompt != "hello world" || search {
		t.Fatalf("got (%q, %v), want (hello world, false)", prompt, search)
	}
}

func TestEncodeDecodePrompt_RoundTripSearch(t *testing.T) {
	enc := EncodePrompt("what is the weather?", true)
	if len(enc) == 0 || enc[0] != 0x00 {
		t.Fatalf("search envelope must start with 0x00 sentinel, got %v", enc[:1])
	}
	prompt, search, err := DecodePrompt(enc)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if prompt != "what is the weather?" || !search {
		t.Fatalf("got (%q, %v), want (what is the weather?, true)", prompt, search)
	}
}

func TestEncodePrompt_NoSearchIsRawBytes(t *testing.T) {
	enc := EncodePrompt("plain prompt", false)
	if string(enc) != "plain prompt" {
		t.Fatalf("non-search EncodePrompt must be raw bytes, got %q", enc)
	}
}

func TestDecodePrompt_SentinelButBadJSON(t *testing.T) {
	_, _, err := DecodePrompt([]byte{0x00, '{', 'n', 'o'})
	if err == nil {
		t.Fatal("expected hard error for 0x00-prefixed unparseable payload")
	}
}

func TestDecodePrompt_Empty(t *testing.T) {
	prompt, search, err := DecodePrompt(nil)
	if err != nil || prompt != "" || search {
		t.Fatalf("got (%q, %v, %v), want empty/false/nil", prompt, search, err)
	}
}
