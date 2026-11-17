package inferenceapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Params is the set of request parameters every endpoint a route can reach
// accepts. A pinned route sends require_parameters, so a request carrying a
// parameter its provider does not list is refused outright: qwen3-vl-8b's
// only OpenRouter provider takes no "reasoning", and gemma-4's take no
// "reasoning_effort". The worker and the disputer drop such a parameter the
// same way, so the job and its re-run stay the same request.
type Params map[string]bool

// Accepts reports whether name may be sent. Nil Params (not known) accepts
// everything, which is what was sent before parameters were looked up.
func (p Params) Accepts(name string) bool {
	return p == nil || p[name]
}

// paramsTimeout bounds one endpoints lookup.
const paramsTimeout = 5 * time.Second

// FetchParams reads the route's endpoints from an OpenRouter-style
// GET {baseURL}/models/{model}/endpoints and returns the parameters all the
// endpoints the route can reach accept: the pinned provider's (its tag is the
// provider slug, optionally followed by "/..."), at the pinned quantization.
// No endpoint matching the route is an error, since every request on it would
// be refused.
func FetchParams(ctx context.Context, client *http.Client, baseURL, apiKey string, r Route) (Params, error) {
	ctx, cancel := context.WithTimeout(ctx, paramsTimeout)
	defer cancel()
	u := strings.TrimRight(baseURL, "/") + "/models/" + escapeModel(r.Model) + "/endpoints"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("endpoints request for %s: %w", r, err)
	}
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("endpoints request for %s: %w", r, err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("endpoints for %s: status %d: %s", r, resp.StatusCode, strings.TrimSpace(string(body)))
	}
	var body struct {
		Data struct {
			Endpoints []struct {
				Tag                 string   `json:"tag"`
				Quantization        string   `json:"quantization"`
				SupportedParameters []string `json:"supported_parameters"`
			} `json:"endpoints"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decode endpoints for %s: %w", r, err)
	}
	var out Params
	for _, e := range body.Data.Endpoints {
		if !servesRoute(e.Tag, e.Quantization, r) {
			continue
		}
		have := make(Params, len(e.SupportedParameters))
		for _, p := range e.SupportedParameters {
			have[p] = true
		}
		if out == nil {
			out = have
			continue
		}
		for p := range out {
			if !have[p] {
				delete(out, p)
			}
		}
	}
	if out == nil {
		return nil, fmt.Errorf("%w: %s", ErrNoEndpoint, r)
	}
	return out, nil
}

// servesRoute reports whether an endpoint (tag "provider[/...]") is one the
// route's provider object can send a request to.
func servesRoute(tag, quantization string, r Route) bool {
	tag = strings.ToLower(tag)
	if p := strings.ToLower(r.Provider); p != "" && tag != p && !strings.HasPrefix(tag, p+"/") {
		return false
	}
	return r.Quantization == "" || strings.EqualFold(quantization, r.Quantization)
}

// escapeModel escapes each path segment of a model id (openai/gpt-oss-20b,
// qwen/qwen3-8b:free), keeping its slashes.
func escapeModel(model string) string {
	parts := strings.Split(model, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return strings.Join(parts, "/")
}

// paramsRetryAfter is how long a failed lookup is answered from memory
// before it is tried again, so a hung endpoints API costs one request a
// paramsTimeout wait, not every request.
const paramsRetryAfter = time.Minute

// ParamsCache looks each route's Params up once per process. A failed lookup
// is retried after paramsRetryAfter; until one succeeds, Get returns the
// error and the caller sends every parameter (nil Params).
type ParamsCache struct {
	client  *http.Client
	baseURL string
	apiKey  string

	mu     sync.Mutex
	seen   map[Route]Params
	failed map[Route]paramsFailure
	now    func() time.Time
}

type paramsFailure struct {
	at  time.Time
	err error
}

// NewParamsCache creates a cache reading from baseURL with client.
func NewParamsCache(client *http.Client, baseURL, apiKey string) *ParamsCache {
	return &ParamsCache{
		client: client, baseURL: baseURL, apiKey: apiKey,
		seen: make(map[Route]Params), failed: make(map[Route]paramsFailure), now: time.Now,
	}
}

// Get returns the route's Params, or nil and the lookup error.
func (c *ParamsCache) Get(ctx context.Context, r Route) (Params, error) {
	c.mu.Lock()
	p, ok := c.seen[r]
	f, failed := c.failed[r]
	c.mu.Unlock()
	if ok {
		return p, nil
	}
	if failed && c.now().Sub(f.at) < paramsRetryAfter {
		return nil, f.err
	}
	p, err := FetchParams(ctx, c.client, c.baseURL, c.apiKey, r)
	c.mu.Lock()
	defer c.mu.Unlock()
	if err != nil {
		c.failed[r] = paramsFailure{at: c.now(), err: err}
		return nil, err
	}
	delete(c.failed, r)
	c.seen[r] = p
	return p, nil
}

// ErrNoEndpoint reports a route no endpoint serves; FetchParams wraps it.
var ErrNoEndpoint = errors.New("no endpoint serves the route")
