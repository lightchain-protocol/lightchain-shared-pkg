package inferenceapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// endpointsJSON is the shape OpenRouter returns, trimmed to what is read.
const endpointsJSON = `{"data":{"id":"m","endpoints":[
 {"tag":"parasail/bf16","quantization":"bf16","supported_parameters":["max_tokens","temperature","seed"]},
 {"tag":"baseten/fp4","quantization":"fp4","supported_parameters":["max_tokens","seed","reasoning","reasoning_effort"]},
 {"tag":"baseten/fp4/us","quantization":"fp4","supported_parameters":["max_tokens","seed","reasoning"]},
 {"tag":"deepinfra","quantization":"unknown","supported_parameters":["max_tokens","reasoning"]}
]}}`

func endpointsServer(t *testing.T, calls *atomic.Int32, status int) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		assert.Equal(t, "/v1/models/qwen/qwen3-8b:free/endpoints", r.URL.Path, "the model id keeps its slash")
		assert.Equal(t, "Bearer sk-x", r.Header.Get("Authorization"))
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, endpointsJSON)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestFetchParams(t *testing.T) {
	t.Parallel()
	var calls atomic.Int32
	srv := endpointsServer(t, &calls, http.StatusOK)
	get := func(provider, quant string) (Params, error) {
		return FetchParams(context.Background(), srv.Client(), srv.URL+"/v1/", "sk-x",
			Route{Model: "qwen/qwen3-8b:free", Provider: provider, Quantization: quant})
	}

	p, err := get("parasail", "bf16")
	require.NoError(t, err)
	assert.Equal(t, Params{"max_tokens": true, "temperature": true, "seed": true}, p)
	assert.False(t, p.Accepts("reasoning"))

	// a provider slug matches its sub-tags; only what all of them take is accepted
	p, err = get("baseten", "fp4")
	require.NoError(t, err)
	assert.True(t, p.Accepts("reasoning"))
	assert.False(t, p.Accepts("reasoning_effort"), "baseten/fp4/us does not take it")

	p, err = get("BaseTen", "")
	require.NoError(t, err)
	assert.True(t, p.Accepts("seed"))

	// unpinned: every endpoint may serve it
	p, err = get("", "")
	require.NoError(t, err)
	assert.Equal(t, Params{"max_tokens": true}, p)

	_, err = get("parasail", "fp8")
	require.ErrorIs(t, err, ErrNoEndpoint)
	_, err = get("base", "")
	require.ErrorIs(t, err, ErrNoEndpoint, "a slug prefix is not a provider")
}

func TestFetchParams_HTTPError(t *testing.T) {
	t.Parallel()
	var calls atomic.Int32
	srv := endpointsServer(t, &calls, http.StatusNotFound)
	_, err := FetchParams(context.Background(), srv.Client(), srv.URL+"/v1", "sk-x", Route{Model: "qwen/qwen3-8b:free"})
	require.ErrorContains(t, err, "status 404")
}

func TestParams_NilAcceptsEverything(t *testing.T) {
	t.Parallel()
	var p Params
	assert.True(t, p.Accepts("reasoning"))
}

func TestParamsCache(t *testing.T) {
	t.Parallel()
	var calls atomic.Int32
	status := http.StatusInternalServerError
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, endpointsJSON)
	}))
	t.Cleanup(srv.Close)
	c := NewParamsCache(srv.Client(), srv.URL, "")
	r := Route{Model: "m", Provider: "parasail"}

	_, err := c.Get(context.Background(), r)
	require.Error(t, err)
	status = http.StatusOK
	p, err := c.Get(context.Background(), r)
	require.NoError(t, err, "a failure is not cached")
	_, err = c.Get(context.Background(), r)
	require.NoError(t, err)
	assert.False(t, p.Accepts("reasoning"))
	assert.Equal(t, int32(2), calls.Load(), "a success is")
}
