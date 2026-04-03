package blob

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
)

// RedisBlobFetcher reads blob data from Redis keyed by versioned hash.
// This is a dev-mode replacement for the Beacon API blob fetcher,
// used when no Beacon API is available (e.g. Anvil).
type RedisBlobFetcher struct {
	rdb *redis.Client
}

// NewRedisBlobFetcher creates a Redis-backed blob fetcher for local dev.
func NewRedisBlobFetcher(rdb *redis.Client) *RedisBlobFetcher {
	return &RedisBlobFetcher{rdb: rdb}
}

// FetchBlob retrieves blob data from Redis using the versioned hash as key.
// The blockNumber parameter is ignored in Redis mode.
func (f *RedisBlobFetcher) FetchBlob(ctx context.Context, versionedHash common.Hash, _ uint64) ([]byte, error) {
	key := fmt.Sprintf("blob:%x", versionedHash)
	data, err := f.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("fetch blob %s from redis: %w", versionedHash.Hex(), err)
	}
	return data, nil
}
