// Package crypto provides cryptographic primitives for the LightChain protocol.
// Includes ECDH key exchange (P-256) and AES-256-GCM symmetric encryption.
package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"errors"
	"fmt"
)

// ErrInvalidPublicKey is returned when a public key is nil or invalid.
var ErrInvalidPublicKey = errors.New("invalid public key")

// GenerateKeyPair generates a new ECDH P-256 key pair for use in session
// key exchange between consumers, workers, and disputers.
func GenerateKeyPair() (*ecdh.PrivateKey, error) {
	key, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate ECDH key pair: %w", err)
	}
	return key, nil
}

// DeriveSharedSecret performs an ECDH key exchange using the local private key
// and a remote public key, producing a shared secret suitable for deriving
// symmetric encryption keys.
func DeriveSharedSecret(priv *ecdh.PrivateKey, remotePub *ecdh.PublicKey) ([]byte, error) {
	if remotePub == nil {
		return nil, ErrInvalidPublicKey
	}
	secret, err := priv.ECDH(remotePub)
	if err != nil {
		return nil, fmt.Errorf("derive shared secret: %w", err)
	}
	return secret, nil
}
