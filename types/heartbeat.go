// Package types defines shared domain types for the LightChain protocol.

package types

import (
	"fmt"
	"strings"
)

// HeartbeatRedisKey returns the canonical Redis key for a worker's heartbeat.
// Both worker and dispatcher must use this format to ensure heartbeat visibility.
// addr can be with or without "0x" prefix; the result is always lowercase with "0x".
// Format: "worker:0x{lowercase_hex}:health"
func HeartbeatRedisKey(addr string) string {
	a := strings.ToLower(addr)
	if !strings.HasPrefix(a, "0x") {
		a = "0x" + a
	}
	return fmt.Sprintf("worker:%s:health", a)
}

// Heartbeat HSET field names — shared contract between worker and dispatcher.
// Worker writes these fields via HSET; dispatcher reads them via HGETALL/HMGET.
const (
	HBFieldLastHeartbeat = "lastHeartbeat" // Unix seconds (int64)
	HBFieldActiveJobs    = "activeJobs"    // int
	HBFieldMaxJobs       = "maxJobs"       // int
	HBFieldLatencyMs     = "latencyMs"     // int (0 until instrumented)
	HBFieldGPUUtil       = "gpuUtil"       // float64 as string (0 until instrumented)
	HBFieldStatus        = "status"        // "active" | "stale" | "draining"
	HBFieldModels        = "models"        // JSON-encoded []string
	HBFieldOllamaStatus  = "ollamaStatus"  // "ready" | "unreachable"
	HBFieldUptime        = "uptimeSeconds" // int64
)

// HeartbeatStatusActive is the value for a healthy, job-accepting worker.
const HeartbeatStatusActive = "active"

// HeartbeatStatusStale indicates the worker has not sent a heartbeat recently.
const HeartbeatStatusStale = "stale"

// HeartbeatStatusDraining indicates the worker is finishing current jobs but not accepting new ones.
const HeartbeatStatusDraining = "draining"
