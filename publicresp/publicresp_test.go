package publicresp

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gowebpki/jcs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	promptHash = "0x1111111111111111111111111111111111111111111111111111111111111111"
	formatHash = "0x2222222222222222222222222222222222222222222222222222222222222222"
)

func params() Params {
	return Params{Model: "llama3-8b", NumCtx: 8192, NumPredict: 512, Seed: 42, Temperature: 0}
}

// The JCS library is pinned, so this proves the pin still canonicalises the
// way RFC 8785 says (its §3.2.3 example): key order, number formatting and
// string escapes are all exercised.
func TestJCS_RFC8785Vector(t *testing.T) {
	t.Parallel()

	in := `{
  "numbers": [333333333.33333329, 1E30, 4.50, 2e-3, 0.000000000000000000000000001],
  "string": "\u20ac$\u000F\u000aA'\u0042\u0022\u005c\\\"\u002f",
  "literals": [null, true, false]
}`
	want := `{"literals":[null,true,false],"numbers":[333333333.3333333,1e+30,4.5,0.002,1e-27],"string":"€$\u000f\nA'B\"\\\\\"/"}`
	out, err := jcs.Transform([]byte(in))
	require.NoError(t, err)
	assert.Equal(t, want, string(out))
}

func TestEncode_IsCanonicalAndRoundTrips(t *testing.T) {
	t.Parallel()

	r := Response{
		V:          Version,
		PromptHash: promptHash,
		FormatHash: formatHash,
		Params:     params(),
		Valid:      true,
		Output:     json.RawMessage(`{"zeta": 1, "alpha": [1.0, 2.50], "name": "x"}`),
	}
	enc, err := Encode(r)
	require.NoError(t, err)
	assert.Equal(t,
		`{"formatHash":"`+formatHash+`","output":{"alpha":[1,2.5],"name":"x","zeta":1},"params":{"model":"llama3-8b","numCtx":8192,"numPredict":512,"seed":42,"temperature":0},"promptHash":"`+promptHash+`","v":1,"valid":true}`,
		string(enc))

	back, err := Decode(enc)
	require.NoError(t, err)
	again, err := Encode(back)
	require.NoError(t, err)
	assert.Equal(t, enc, again, "decode then encode is the identity")
	assert.Equal(t, Hash(enc), Hash(again))
}

// The same object with different whitespace or key order gives the same R.
func TestEncode_IgnoresKeyOrderAndWhitespace(t *testing.T) {
	t.Parallel()

	a := Response{V: Version, PromptHash: promptHash, FormatHash: ZeroHash, Params: params(), Valid: true,
		Output: json.RawMessage(`{"b":2,"a":1}`)}
	b := a
	b.Output = json.RawMessage("{ \"a\" : 1 ,\n \"b\" : 2 }")
	ea, err := Encode(a)
	require.NoError(t, err)
	eb, err := Encode(b)
	require.NoError(t, err)
	assert.Equal(t, ea, eb)
	assert.Equal(t, a.OutputHash(), b.OutputHash())
}

func TestEncode_RawAnswerWhenNoFormat(t *testing.T) {
	t.Parallel()

	text := "The answer is 42."
	r := Response{V: Version, PromptHash: promptHash, FormatHash: ZeroHash, Params: params(), Answer: &text}
	enc, err := Encode(r)
	require.NoError(t, err)
	assert.Contains(t, string(enc), `"answer":"The answer is 42."`)
	assert.NotContains(t, string(enc), `"output"`)

	back, err := Decode(enc)
	require.NoError(t, err)
	require.NotNil(t, back.Answer)
	assert.Equal(t, text, *back.Answer)
	assert.False(t, back.Valid)
}

// An invalid output (truncated JSON from numPredict) is committed as a raw
// answer with valid false, never retried (draft 04 §4.3).
func TestEncode_InvalidOutputIsRawAnswerNotValid(t *testing.T) {
	t.Parallel()

	cut := `{"summary":"the mod`
	r := Response{V: Version, PromptHash: promptHash, FormatHash: formatHash, Params: params(), Answer: &cut, Valid: false}
	_, err := Encode(r)
	require.NoError(t, err)

	r.Valid = true
	_, err = Encode(r)
	require.ErrorIs(t, err, ErrInvalid)
}

func TestDecode_RefusesNonCanonicalBytes(t *testing.T) {
	t.Parallel()

	canon := `{"formatHash":"` + ZeroHash + `","output":{"a":1},"params":{"model":"m","numCtx":1,"numPredict":1,"seed":42,"temperature":0},"promptHash":"` + promptHash + `","v":1,"valid":true}`
	_, err := Decode([]byte(canon))
	require.NoError(t, err)

	for name, b := range map[string]string{
		"whitespace":    strings.Replace(canon, `"v":1`, `"v": 1`, 1),
		"key order":     strings.Replace(strings.Replace(canon, `"v":1,`, ``, 1), `"valid":true`, `"valid":true,"v":1`, 1),
		"unknown field": strings.Replace(canon, `"v":1`, `"v":1,"x":1`, 1),
		"not json":      "nope",
	} {
		_, err := Decode([]byte(b))
		assert.Error(t, err, name)
	}
}

func TestDecode_RefusesWrongVersionAndBadFields(t *testing.T) {
	t.Parallel()

	base := Response{V: Version, PromptHash: promptHash, FormatHash: ZeroHash, Params: params(), Valid: true,
		Output: json.RawMessage(`{}`)}
	cases := map[string]func(*Response){
		"version 2":        func(r *Response) { r.V = 2 },
		"short promptHash": func(r *Response) { r.PromptHash = "0x1234" },
		"bad formatHash":   func(r *Response) { r.FormatHash = "zz" },
		"no model":         func(r *Response) { r.Params.Model = "" },
		"zero numCtx":      func(r *Response) { r.Params.NumCtx = 0 },
		"both set":         func(r *Response) { s := "x"; r.Answer = &s },
		"neither set":      func(r *Response) { r.Output = nil },
		"action no hex":    func(r *Response) { r.Action = "abc" },
	}
	for name, mutate := range cases {
		r := base
		mutate(&r)
		_, err := Encode(r)
		assert.ErrorIs(t, err, ErrInvalid, name)
	}
}

func TestAction_CarriedVerbatim(t *testing.T) {
	t.Parallel()

	r := Response{V: Version, PromptHash: promptHash, FormatHash: formatHash, Params: params(), Valid: true,
		Output: json.RawMessage(`{"action":{"kind":"none"}}`), Action: "0x00"}
	enc, err := Encode(r)
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(string(enc), `{"action":"0x00",`), "action sorts first under JCS (draft 03 §5.1)")
	back, err := Decode(enc)
	require.NoError(t, err)
	assert.Equal(t, "0x00", back.Action)
}

func TestFormatHash(t *testing.T) {
	t.Parallel()

	empty, err := FormatHash(nil)
	require.NoError(t, err)
	assert.Equal(t, ZeroHash, empty)

	a, err := FormatHash(json.RawMessage(`{"type":"object","required":["x"]}`))
	require.NoError(t, err)
	b, err := FormatHash(json.RawMessage("{ \"required\" : [\"x\"], \"type\": \"object\" }"))
	require.NoError(t, err)
	assert.Equal(t, a, b, "same schema, same hash")
	assert.Len(t, a, 66)

	_, err = FormatHash(json.RawMessage(`{bad`))
	assert.Error(t, err)
}

func TestPromptHash_IsKeccakOfExactBytes(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", PromptHash(nil), "keccak256 of empty input")
	assert.NotEqual(t, PromptHash([]byte(`{"v":4}`)), PromptHash([]byte(`{"v": 4}`)), "the envelope is hashed as sent, not canonicalised")
}

func TestOutputHash_DistinguishesValidFlagAndContent(t *testing.T) {
	t.Parallel()

	a := Response{Output: json.RawMessage(`{"a":1}`), Valid: true}
	b := Response{Output: json.RawMessage(`{"a":1}`), Valid: false}
	c := Response{Output: json.RawMessage(`{"a":2}`), Valid: true}
	assert.NotEqual(t, a.OutputHash(), b.OutputHash())
	assert.NotEqual(t, a.OutputHash(), c.OutputHash())
	assert.Equal(t, a.OutputHash(), Response{Output: json.RawMessage(`{"a":1}`), Valid: true}.OutputHash())
}
