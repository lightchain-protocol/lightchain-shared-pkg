package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptSessionKey_roundtrip(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	sessionKey := make([]byte, 32)
	_, err = rand.Read(sessionKey)
	require.NoError(t, err)

	enc, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)
	require.NotEmpty(t, enc)

	got, err := DecryptSessionKey(enc, workerPriv)
	require.NoError(t, err)
	assert.Equal(t, sessionKey, got)
}

func TestEncryptSessionKey_uniqueCiphertexts(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	sessionKey := make([]byte, 32)
	_, err = rand.Read(sessionKey)
	require.NoError(t, err)

	enc1, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)

	enc2, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)

	// Each call generates a fresh ephemeral key pair, so ciphertexts must differ.
	assert.NotEqual(t, enc1, enc2)
}

func TestDecryptSessionKey_wrongPrivateKey(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	otherPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	sessionKey := make([]byte, 32)
	_, err = rand.Read(sessionKey)
	require.NoError(t, err)

	enc, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)

	_, err = DecryptSessionKey(enc, otherPriv)
	require.Error(t, err)
}

func TestDecryptSessionKey_truncatedInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []byte
	}{
		{name: "empty", input: []byte{}},
		{name: "one byte", input: []byte{0x04}},
		{name: "63 bytes", input: make([]byte, 63)},
		{name: "64 bytes", input: make([]byte, 64)},
	}

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := DecryptSessionKey(tt.input, workerPriv)
			require.ErrorIs(t, err, ErrEncWorkerKeyTooShort)
		})
	}
}

func TestDecryptSessionKey_nilWorkerPriv(t *testing.T) {
	t.Parallel()
	_, err := DecryptSessionKey(make([]byte, 100), nil)
	require.ErrorIs(t, err, ErrInvalidPrivateKey)
}

func TestEncryptSessionKey_nilRemotePub(t *testing.T) {
	t.Parallel()

	sessionKey := make([]byte, 32)
	_, err := rand.Read(sessionKey)
	require.NoError(t, err)

	_, err = EncryptSessionKey(sessionKey, nil)
	require.ErrorIs(t, err, ErrInvalidPublicKey)
}

func TestEncryptSessionKey_wrongSessionKeyLength(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	tests := []struct {
		name   string
		keyLen int
	}{
		{name: "empty", keyLen: 0},
		{name: "16 bytes (AES-128)", keyLen: 16},
		{name: "31 bytes", keyLen: 31},
		{name: "33 bytes", keyLen: 33},
		{name: "64 bytes", keyLen: 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := make([]byte, tt.keyLen)
			_, err := EncryptSessionKey(key, workerPriv.PublicKey())
			require.ErrorIs(t, err, ErrSessionKeyLength)
		})
	}
}

func TestDecryptSessionKey_resultIs32Bytes(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	sessionKey := make([]byte, 32)
	_, err = rand.Read(sessionKey)
	require.NoError(t, err)

	enc, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)

	got, err := DecryptSessionKey(enc, workerPriv)
	require.NoError(t, err)
	assert.Len(t, got, 32)
}

func TestDecryptSessionKey_invalidEphemeralPubKey(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	// Construct a blob that is long enough (65 + some bytes) but whose first 65
	// bytes are not a valid uncompressed P-256 point.
	garbage := make([]byte, p256UncompressedKeySize+32)
	_, err = rand.Read(garbage)
	require.NoError(t, err)
	// Ensure the first byte is NOT 0x04 so it is definitely invalid.
	garbage[0] = 0x00

	_, err = DecryptSessionKey(garbage, workerPriv)
	require.Error(t, err)
}

func TestEncryptSessionKey_outputPrefixedWithEphemeralPub(t *testing.T) {
	t.Parallel()

	workerPriv, err := GenerateKeyPair()
	require.NoError(t, err)

	sessionKey := make([]byte, 32)
	_, err = rand.Read(sessionKey)
	require.NoError(t, err)

	enc, err := EncryptSessionKey(sessionKey, workerPriv.PublicKey())
	require.NoError(t, err)

	// First 65 bytes must be a valid uncompressed P-256 public key.
	require.GreaterOrEqual(t, len(enc), p256UncompressedKeySize)
	_, err = ecdh.P256().NewPublicKey(enc[:p256UncompressedKeySize])
	require.NoError(t, err, "first 65 bytes of output must be a valid P-256 uncompressed public key")
}
