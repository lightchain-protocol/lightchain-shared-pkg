package crypto

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKeyPair(t *testing.T) {
	t.Parallel()

	key, err := GenerateKeyPair()
	require.NoError(t, err)
	require.NotNil(t, key)
	require.NotNil(t, key.PublicKey())
}

func TestDeriveSharedSecret(t *testing.T) {
	t.Parallel()

	alice, err := GenerateKeyPair()
	require.NoError(t, err)

	bob, err := GenerateKeyPair()
	require.NoError(t, err)

	// Alice derives shared secret with Bob's public key
	secretAB, err := DeriveSharedSecret(alice, bob.PublicKey())
	require.NoError(t, err)

	// Bob derives shared secret with Alice's public key
	secretBA, err := DeriveSharedSecret(bob, alice.PublicKey())
	require.NoError(t, err)

	// Both sides produce the same shared secret
	assert.Equal(t, secretAB, secretBA)
	assert.NotEmpty(t, secretAB)
}

func TestDeriveSharedSecret_nilPublicKey(t *testing.T) {
	t.Parallel()

	key, err := GenerateKeyPair()
	require.NoError(t, err)

	_, err = DeriveSharedSecret(key, nil)
	require.ErrorIs(t, err, ErrInvalidPublicKey)
}

func TestEncryptDecryptRoundtrip(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		plaintext []byte
	}{
		{name: "hello world", plaintext: []byte("hello world")},
		{name: "empty plaintext", plaintext: nil},
		{name: "binary data", plaintext: []byte{0x00, 0xff, 0x01, 0xfe}},
		{name: "large payload", plaintext: make([]byte, 1024*1024)},
	}

	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ciphertext, err := Encrypt(key, tt.plaintext)
			require.NoError(t, err)
			require.NotEmpty(t, ciphertext)

			decrypted, err := Decrypt(key, ciphertext)
			require.NoError(t, err)
			assert.Equal(t, tt.plaintext, decrypted)
		})
	}
}

func TestDecryptWrongKey(t *testing.T) {
	t.Parallel()

	key1 := make([]byte, 32)
	_, err := rand.Read(key1)
	require.NoError(t, err)

	key2 := make([]byte, 32)
	_, err = rand.Read(key2)
	require.NoError(t, err)

	plaintext := []byte("secret message")
	ciphertext, err := Encrypt(key1, plaintext)
	require.NoError(t, err)

	_, err = Decrypt(key2, ciphertext)
	require.Error(t, err)
}

func TestEncrypt_invalidKeyLength(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		keyLen int
	}{
		{name: "too short", keyLen: 16},
		{name: "too long", keyLen: 64},
		{name: "empty", keyLen: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := make([]byte, tt.keyLen)
			_, err := Encrypt(key, []byte("test"))
			require.ErrorIs(t, err, ErrInvalidKeyLength)
		})
	}
}

func TestDecrypt_invalidKeyLength(t *testing.T) {
	t.Parallel()

	key := make([]byte, 16)
	_, err := Decrypt(key, []byte("some ciphertext that is long enough"))
	require.ErrorIs(t, err, ErrInvalidKeyLength)
}

func TestDecrypt_ciphertextTooShort(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.NoError(t, err)

	_, err = Decrypt(key, []byte("short"))
	require.ErrorIs(t, err, ErrCiphertextTooShort)
}

func TestEncrypt_uniqueCiphertexts(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.NoError(t, err)

	plaintext := []byte("same input")

	ct1, err := Encrypt(key, plaintext)
	require.NoError(t, err)

	ct2, err := Encrypt(key, plaintext)
	require.NoError(t, err)

	// Random nonces mean identical plaintext produces different ciphertext
	assert.NotEqual(t, ct1, ct2)
}
