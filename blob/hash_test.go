package blob

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionedHash(t *testing.T) {
	t.Parallel()

	commitment := make([]byte, 48)
	for i := range commitment {
		commitment[i] = byte(i)
	}

	hash := VersionedHash(commitment)

	assert.Equal(t, byte(0x01), hash[0])

	expected := sha256.Sum256(commitment)
	assert.Equal(t, expected[1:], hash[1:])
}

func TestKzgToVersionedHash(t *testing.T) {
	t.Parallel()

	commitment := make([]byte, 48)
	for i := range commitment {
		commitment[i] = byte(i)
	}
	commitHex := "0x" + hex.EncodeToString(commitment)

	hash := KzgToVersionedHash(commitHex)

	assert.Equal(t, byte(0x01), hash[0])

	expected := sha256.Sum256(commitment)
	assert.Equal(t, expected[1:], hash[1:])
}

func TestKzgToVersionedHash_InvalidHex(t *testing.T) {
	t.Parallel()

	hash := KzgToVersionedHash("0xZZZZZZ")
	assert.Equal(t, [32]byte{}, [32]byte(hash))
}

func TestStripHexPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"0xabcd", "abcd"},
		{"0Xabcd", "abcd"},
		{"abcd", "abcd"},
		{"", ""},
		{"0x", ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, StripHexPrefix(tt.input))
	}
}
