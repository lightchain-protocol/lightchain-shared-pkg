// Package types — worker drain marker primitives.

package types

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// DrainRedisKey returns the canonical Redis key for a worker's drain marker.
// The drain marker is stored independently from the heartbeat hash: its
// presence signals that the dispatcher should not select the worker for new
// sessions, even if the worker's heartbeat is otherwise healthy.
//
// addr can be with or without "0x" prefix; the result is always lowercase
// with "0x". Format: "worker:0x{lowercase_hex}:drain"
func DrainRedisKey(addr string) string {
	a := strings.ToLower(addr)
	if !strings.HasPrefix(a, "0x") {
		a = "0x" + a
	}
	return fmt.Sprintf("worker:%s:drain", a)
}

// DrainMarkerValue is the canonical value stored at the drain key.
// Callers should rely on key existence, not value equality — see IsDraining.
const DrainMarkerValue = "1"

// SetDraining writes the drain marker for addr with the given TTL.
// Idempotent: an existing marker is overwritten and its TTL refreshed.
//
// Drain TTL must outlive the on-chain dispute window plus operator slack
// (see docs/worker-drain-plan.md). Compute it at the call site;
// pkg/types intentionally does not depend on contract bindings.
func SetDraining(ctx context.Context, rdb *redis.Client, addr string, ttl time.Duration) error {
	key := DrainRedisKey(addr)
	if err := rdb.Set(ctx, key, DrainMarkerValue, ttl).Err(); err != nil {
		return fmt.Errorf("set drain %s: %w", key, err)
	}
	return nil
}

// Undrain removes the drain marker for addr.
// Idempotent: returns nil if the marker is already absent.
func Undrain(ctx context.Context, rdb *redis.Client, addr string) error {
	key := DrainRedisKey(addr)
	if err := rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("delete drain %s: %w", key, err)
	}
	return nil
}

// IsDraining reports whether a drain marker exists for addr.
// A non-nil error indicates a Redis failure.
func IsDraining(ctx context.Context, rdb *redis.Client, addr string) (bool, error) {
	key := DrainRedisKey(addr)
	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("check drain %s: %w", key, err)
	}
	return n > 0, nil
}
