package types_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lightchain/pkg/types"
)

const testWorkerAddr = "0xab5801a7d398351b8be11c439e05c5b3259aec9b"

func newDrainTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return rdb, mr
}

func TestDrainRedisKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "checksummed address",
			addr: "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
			want: "worker:0xab5801a7d398351b8be11c439e05c5b3259aec9b:drain",
		},
		{
			name: "lowercase address",
			addr: "0xab5801a7d398351b8be11c439e05c5b3259aec9b",
			want: "worker:0xab5801a7d398351b8be11c439e05c5b3259aec9b:drain",
		},
		{
			name: "no 0x prefix",
			addr: "ab5801a7d398351b8be11c439e05c5b3259aec9b",
			want: "worker:0xab5801a7d398351b8be11c439e05c5b3259aec9b:drain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, types.DrainRedisKey(tt.addr))
		})
	}
}

func TestSetDraining_writesMarkerWithTTL(t *testing.T) {
	t.Parallel()
	rdb, mr := newDrainTestRedis(t)
	ctx := context.Background()

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))

	key := types.DrainRedisKey(testWorkerAddr)
	val, err := mr.Get(key)
	require.NoError(t, err)
	assert.Equal(t, types.DrainMarkerValue, val)

	ttl := mr.TTL(key)
	assert.Greater(t, ttl, time.Duration(0))
	assert.LessOrEqual(t, ttl, time.Hour)
}

func TestSetDraining_isIdempotentAndRefreshesTTL(t *testing.T) {
	t.Parallel()
	rdb, mr := newDrainTestRedis(t)
	ctx := context.Background()

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))
	mr.FastForward(30 * time.Minute)

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))

	ttl := mr.TTL(types.DrainRedisKey(testWorkerAddr))
	assert.Greater(t, ttl, 30*time.Minute,
		"TTL should be refreshed to a full hour after the second SetDraining")
}

func TestUndrain_removesMarker(t *testing.T) {
	t.Parallel()
	rdb, mr := newDrainTestRedis(t)
	ctx := context.Background()

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))
	require.True(t, mr.Exists(types.DrainRedisKey(testWorkerAddr)))

	require.NoError(t, types.Undrain(ctx, rdb, testWorkerAddr))
	assert.False(t, mr.Exists(types.DrainRedisKey(testWorkerAddr)))
}

func TestUndrain_isIdempotentOnAbsentMarker(t *testing.T) {
	t.Parallel()
	rdb, _ := newDrainTestRedis(t)
	ctx := context.Background()

	assert.NoError(t, types.Undrain(ctx, rdb, testWorkerAddr))
}

func TestIsDraining_reflectsMarkerState(t *testing.T) {
	t.Parallel()
	rdb, _ := newDrainTestRedis(t)
	ctx := context.Background()

	got, err := types.IsDraining(ctx, rdb, testWorkerAddr)
	require.NoError(t, err)
	assert.False(t, got, "fresh worker should not be draining")

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))
	got, err = types.IsDraining(ctx, rdb, testWorkerAddr)
	require.NoError(t, err)
	assert.True(t, got, "worker should report draining after SetDraining")

	require.NoError(t, types.Undrain(ctx, rdb, testWorkerAddr))
	got, err = types.IsDraining(ctx, rdb, testWorkerAddr)
	require.NoError(t, err)
	assert.False(t, got, "worker should not report draining after Undrain")
}

func TestIsDraining_falseAfterTTLExpiry(t *testing.T) {
	t.Parallel()
	rdb, mr := newDrainTestRedis(t)
	ctx := context.Background()

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))

	got, err := types.IsDraining(ctx, rdb, testWorkerAddr)
	require.NoError(t, err)
	assert.True(t, got)

	mr.FastForward(2 * time.Hour)

	got, err = types.IsDraining(ctx, rdb, testWorkerAddr)
	require.NoError(t, err)
	assert.False(t, got, "drain marker should be gone after TTL expiry")
}

func TestDrain_independentFromHeartbeatHash(t *testing.T) {
	t.Parallel()
	rdb, mr := newDrainTestRedis(t)
	ctx := context.Background()

	hbKey := types.HeartbeatRedisKey(testWorkerAddr)
	require.NoError(t, rdb.HSet(ctx, hbKey, types.HBFieldStatus, types.HeartbeatStatusActive).Err())
	require.NoError(t, rdb.PExpire(ctx, hbKey, 30*time.Second).Err())

	require.NoError(t, types.SetDraining(ctx, rdb, testWorkerAddr, time.Hour))

	hbStatus, err := rdb.HGet(ctx, hbKey, types.HBFieldStatus).Result()
	require.NoError(t, err)
	assert.Equal(t, types.HeartbeatStatusActive, hbStatus,
		"heartbeat status field must be unaffected by SetDraining (decoupled keys)")

	hbTTL := mr.TTL(hbKey)
	assert.LessOrEqual(t, hbTTL, 30*time.Second,
		"heartbeat key TTL must not be extended by SetDraining")

	drainTTL := mr.TTL(types.DrainRedisKey(testWorkerAddr))
	assert.Greater(t, drainTTL, 30*time.Second,
		"drain key TTL must be independent and longer than heartbeat TTL")
}
