package inferenceapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRoute(t *testing.T) {
	cases := []struct {
		in   string
		want Route
	}{
		{"openai/gpt-oss-20b", Route{Model: "openai/gpt-oss-20b"}},
		{"qwen/qwen3-8b:free", Route{Model: "qwen/qwen3-8b:free"}},
		{"openai/gpt-oss-20b@deepinfra", Route{Model: "openai/gpt-oss-20b", Provider: "deepinfra"}},
		{"openai/gpt-oss-20b@deepinfra:FP4", Route{Model: "openai/gpt-oss-20b", Provider: "deepinfra", Quantization: "fp4"}},
		{"qwen/qwen3-8b:free@groq:fp8", Route{Model: "qwen/qwen3-8b:free", Provider: "groq", Quantization: "fp8"}},
		{"google/gemma-4-26b-a4b-it@:bf16", Route{Model: "google/gemma-4-26b-a4b-it", Quantization: "bf16"}},
		{" gpt-oss:20b ", Route{Model: "gpt-oss:20b"}},
	}
	for _, c := range cases {
		got, err := ParseRoute(c.in)
		require.NoError(t, err, c.in)
		assert.Equal(t, c.want, got, c.in)
	}
}

func TestParseRoute_Errors(t *testing.T) {
	for _, in := range []string{"", "@deepinfra", "openai/gpt-oss-20b@", "openai/gpt-oss-20b@:", "m@p:fp7"} {
		_, err := ParseRoute(in)
		assert.Error(t, err, in)
	}
}

func TestRoute_StringRoundTrips(t *testing.T) {
	for _, in := range []string{"openai/gpt-oss-20b", "openai/gpt-oss-20b@deepinfra", "a/b:free@groq:fp8", "m@:bf16"} {
		r, err := ParseRoute(in)
		require.NoError(t, err)
		assert.Equal(t, in, r.String())
	}
}

func TestPreferences(t *testing.T) {
	assert.Nil(t, Route{Model: "m"}.Preferences(), "an unpinned route sends no provider object")

	b, err := json.Marshal(Route{Model: "m", Provider: "deepinfra", Quantization: "fp4"}.Preferences())
	require.NoError(t, err)
	assert.JSONEq(t, `{"order":["deepinfra"],"allow_fallbacks":false,"quantizations":["fp4"],"require_parameters":true,"data_collection":"deny"}`, string(b))

	b, err = json.Marshal(Route{Model: "m", Quantization: "fp8"}.Preferences())
	require.NoError(t, err)
	assert.JSONEq(t, `{"quantizations":["fp8"],"require_parameters":true,"data_collection":"deny"}`, string(b))
}

func TestParseRoutes(t *testing.T) {
	got, err := ParseRoutes("gpt-oss:20b=openai/gpt-oss-20b@deepinfra:fp4; qwen3-vl:8b=qwen/qwen3-vl-8b-instruct;")
	require.NoError(t, err)
	assert.Equal(t, map[string]Route{
		"gpt-oss:20b": {Model: "openai/gpt-oss-20b", Provider: "deepinfra", Quantization: "fp4"},
		"qwen3-vl:8b": {Model: "qwen/qwen3-vl-8b-instruct"},
	}, got)

	got, err = ParseRoutes("  ")
	require.NoError(t, err)
	assert.Nil(t, got)

	_, err = ParseRoutes("a=x;a=y;b;c=m@p:bad")
	require.Error(t, err)
	assert.Contains(t, err.Error(), `duplicate entry for model "a"`)
	assert.Contains(t, err.Error(), `malformed entry "b"`)
	assert.Contains(t, err.Error(), `unknown quantization "bad"`)
}
