package blob

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

// VersionedHash computes the EIP-4844 versioned hash from raw commitment bytes.
// versionedHash = 0x01 || SHA256(commitment)[1:]
func VersionedHash(commitmentBytes []byte) common.Hash {
	h := sha256.Sum256(commitmentBytes)
	h[0] = 0x01
	return common.Hash(h)
}

// KzgToVersionedHash computes the EIP-4844 versioned hash from a hex-encoded KZG commitment.
func KzgToVersionedHash(kzgCommitment string) common.Hash {
	commitBytes, err := hex.DecodeString(StripHexPrefix(kzgCommitment))
	if err != nil {
		return common.Hash{}
	}
	return VersionedHash(commitBytes)
}

// StripHexPrefix removes the "0x" or "0X" prefix from a hex string.
func StripHexPrefix(s string) string {
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}
	return s
}
