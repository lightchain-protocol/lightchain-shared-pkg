package crypto

import (
	"crypto/ecdh"
	"errors"
	"fmt"
)

const (
	// p256UncompressedKeySize is the byte length of an uncompressed P-256 public key.
	p256UncompressedKeySize = 65
	// sessionKeySize is the required byte length for a session key (AES-256).
	sessionKeySize = 32
)

// ErrSessionKeyLength is returned when the session key is not exactly 32 bytes.
var ErrSessionKeyLength = errors.New("session key must be exactly 32 bytes")

// ErrEncWorkerKeyTooShort is returned when the encrypted worker key blob is shorter
// than the 65-byte uncompressed ephemeral public key prefix.
var ErrEncWorkerKeyTooShort = errors.New("encrypted worker key too short: must be at least 65 bytes")

// ErrInvalidPrivateKey is returned when a private key is nil or invalid.
var ErrInvalidPrivateKey = errors.New("invalid private key")

// EncryptSessionKey encrypts a 32-byte session key for delivery to a remote party.
// It generates an ephemeral ECDH P-256 key pair, derives a shared secret with
// remotePub, and encrypts sessionKey using AES-256-GCM with that shared secret.
//
// Output format: ephemeralPub.Bytes() (65 bytes) || AES-GCM(sharedSecret, sessionKey)
func EncryptSessionKey(sessionKey []byte, remotePub *ecdh.PublicKey) ([]byte, error) {
	if len(sessionKey) != sessionKeySize {
		return nil, ErrSessionKeyLength
	}
	if remotePub == nil {
		return nil, ErrInvalidPublicKey
	}

	ephemeralPriv, err := GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("encrypt session key: generate ephemeral key pair: %w", err)
	}

	sharedSecret, err := DeriveSharedSecret(ephemeralPriv, remotePub)
	if err != nil {
		return nil, fmt.Errorf("encrypt session key: derive shared secret: %w", err)
	}

	ciphertext, err := Encrypt(sharedSecret, sessionKey)
	if err != nil {
		return nil, fmt.Errorf("encrypt session key: encrypt: %w", err)
	}

	ephemeralPubBytes := ephemeralPriv.PublicKey().Bytes()
	out := make([]byte, 0, len(ephemeralPubBytes)+len(ciphertext))
	out = append(out, ephemeralPubBytes...)
	out = append(out, ciphertext...)
	return out, nil
}

// DecryptSessionKey decrypts an encrypted session key using the worker's ECDH private key.
// The input must be at least 65 bytes: the first 65 bytes are the uncompressed ephemeral
// P-256 public key, followed by the AES-GCM ciphertext of the session key.
//
// Returns the 32-byte session key on success.
func DecryptSessionKey(encWorkerKey []byte, workerPriv *ecdh.PrivateKey) ([]byte, error) {
	if workerPriv == nil {
		return nil, ErrInvalidPrivateKey
	}
	if len(encWorkerKey) < p256UncompressedKeySize {
		return nil, ErrEncWorkerKeyTooShort
	}

	ephemeralPubBytes := encWorkerKey[:p256UncompressedKeySize]
	ciphertext := encWorkerKey[p256UncompressedKeySize:]

	ephemeralPub, err := ecdh.P256().NewPublicKey(ephemeralPubBytes)
	if err != nil {
		return nil, fmt.Errorf("decrypt session key: parse ephemeral public key: %w", err)
	}

	sharedSecret, err := DeriveSharedSecret(workerPriv, ephemeralPub)
	if err != nil {
		return nil, fmt.Errorf("decrypt session key: derive shared secret: %w", err)
	}

	sessionKey, err := Decrypt(sharedSecret, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("decrypt session key: decrypt: %w", err)
	}

	if len(sessionKey) != sessionKeySize {
		return nil, fmt.Errorf("decrypt session key: %w (got %d bytes)", ErrSessionKeyLength, len(sessionKey))
	}

	return sessionKey, nil
}
