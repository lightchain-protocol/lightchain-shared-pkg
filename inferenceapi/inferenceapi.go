// Package inferenceapi is the shared contract between a worker that serves jobs
// from a hosted OpenAI-compatible API and the disputer that audits it: the
// on-chain capability such a worker declares, the model route syntax both read
// from their configuration, and the provider-routing object both send. A job
// is only reproducible when the disputer re-runs it on the route the worker
// served it on, so the two must build the same request from the same route.
package inferenceapi

import (
	"errors"
	"fmt"
	"strings"
)

// Capability is the WorkerRegistry capability an API-backed worker declares.
// The disputer audits such a worker against the hosted reference instead of
// its local model, and a consumer whose prompts must not leave the worker
// excludes it from a request (SessionManager RequestTerms.excludedCaps).
const Capability = "api-inference"

// Route is where one on-chain model name is served from.
type Route struct {
	// Model is the provider's model id, e.g. openai/gpt-oss-20b.
	Model string
	// Provider pins the upstream provider (an OpenRouter provider slug, e.g.
	// deepinfra). Empty leaves the choice to the endpoint, which may move a
	// request between providers with different weights or quantizations.
	Provider string
	// Quantization pins the weights' precision (OpenRouter's values: int4,
	// int8, fp4, fp6, fp8, fp16, bf16, fp32). Empty accepts any.
	Quantization string
}

// Pinned reports whether the route fixes the provider or the quantization.
func (r Route) Pinned() bool {
	return r.Provider != "" || r.Quantization != ""
}

// String is the route in the syntax ParseRoute reads.
func (r Route) String() string {
	s := r.Model
	if r.Pinned() {
		s += "@" + r.Provider
		if r.Quantization != "" {
			s += ":" + r.Quantization
		}
	}
	return s
}

// ProviderPreferences is OpenRouter's "provider" request object.
type ProviderPreferences struct {
	Order             []string `json:"order,omitempty"`
	AllowFallbacks    *bool    `json:"allow_fallbacks,omitempty"`
	Quantizations     []string `json:"quantizations,omitempty"`
	RequireParameters bool     `json:"require_parameters,omitempty"`
	DataCollection    string   `json:"data_collection,omitempty"`
}

// Preferences returns the provider object for the route, or nil for an
// unpinned route, which is sent with none so any OpenAI-compatible endpoint
// accepts it. A pinned route never falls back to another provider, only goes
// to providers that honour every request parameter (the seed among them), and
// only to providers that do not keep prompts.
func (r Route) Preferences() *ProviderPreferences {
	if !r.Pinned() {
		return nil
	}
	noFallbacks := false
	p := &ProviderPreferences{RequireParameters: true, DataCollection: "deny"}
	if r.Provider != "" {
		p.Order = []string{r.Provider}
		p.AllowFallbacks = &noFallbacks
	}
	if r.Quantization != "" {
		p.Quantizations = []string{r.Quantization}
	}
	return p
}

var quantizations = map[string]bool{
	"int4": true, "int8": true, "fp4": true, "fp6": true, "fp8": true,
	"fp16": true, "bf16": true, "fp32": true,
}

// ParseRoute reads "model[@provider[:quantization]]", e.g.
// openai/gpt-oss-20b@deepinfra:fp4. Model ids carry slashes and colons
// (qwen/qwen3-8b:free), so the provider part starts at the last "@"; provider
// slugs carry no colon, so its quantization starts at the first ":". An empty
// provider pins only the quantization (model@:fp8).
func ParseRoute(s string) (Route, error) {
	s = strings.TrimSpace(s)
	model, pin, pinned := s, "", false
	if i := strings.LastIndex(s, "@"); i >= 0 {
		model, pin, pinned = s[:i], s[i+1:], true
	}
	r := Route{Model: strings.TrimSpace(model)}
	if r.Model == "" {
		return Route{}, fmt.Errorf("route %q: empty model id", s)
	}
	if !pinned {
		return r, nil
	}
	provider, quant, _ := strings.Cut(pin, ":")
	r.Provider, r.Quantization = strings.TrimSpace(provider), strings.ToLower(strings.TrimSpace(quant))
	if !r.Pinned() {
		return Route{}, fmt.Errorf("route %q: empty pin after @", s)
	}
	if r.Quantization != "" && !quantizations[r.Quantization] {
		return Route{}, fmt.Errorf("route %q: unknown quantization %q", s, r.Quantization)
	}
	return r, nil
}

// ParseRoutes reads "name=route;name2=route2", the INFERENCE_API_MODELS
// syntax. Entries split on ";" and then at the first "=", because model names
// carry colons (gpt-oss:20b) and routes carry slashes, colons and "@" but
// neither carries "=" or ";". Every malformed entry is reported.
func ParseRoutes(raw string) (map[string]Route, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}
	out := make(map[string]Route)
	var errs []error
	for _, entry := range strings.Split(raw, ";") {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}
		name, spec, found := strings.Cut(entry, "=")
		name = strings.TrimSpace(name)
		if !found || name == "" || strings.TrimSpace(spec) == "" {
			errs = append(errs, fmt.Errorf("malformed entry %q (want name=provider-model[@provider[:quantization]])", entry))
			continue
		}
		if _, dup := out[name]; dup {
			errs = append(errs, fmt.Errorf("duplicate entry for model %q", name))
			continue
		}
		r, err := ParseRoute(spec)
		if err != nil {
			errs = append(errs, fmt.Errorf("model %q: %w", name, err))
			continue
		}
		out[name] = r
	}
	return out, errors.Join(errs...)
}
