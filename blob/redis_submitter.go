package blob

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisBlobSubmitter stores blob data in Redis keyed by sha256 hash.
// This is a dev-mode replacement for the real EIP-4844 blob submitter,
// used when no Beacon API is available (e.g. Anvil).
type RedisBlobSubmitter struct {
	rdb *redis.Client
}

// NewRedisBlobSubmitter creates a Redis-backed blob submitter for local dev.
func NewRedisBlobSubmitter(rdb *redis.Client) *RedisBlobSubmitter {
	return &RedisBlobSubmitter{rdb: rdb}
}

// SubmitBlobTx stores data in Redis and returns its sha256 hash as the blob identifier.
// The hash serves the same role as the KZG versioned hash in real EIP-4844 blobs.
func (s *RedisBlobSubmitter) SubmitBlobTx(ctx context.Context, data []byte) ([][32]byte, error) {
	hash := sha256.Sum256(data)
	key := fmt.Sprintf("blob:%x", hash)

	if err := s.rdb.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("store blob in redis: %w", err)
	}

	return [][32]byte{hash}, nil
}
