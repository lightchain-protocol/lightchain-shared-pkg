// Package types defines shared domain types for the LightChain protocol.
// All types map to their Solidity counterparts in the on-chain contracts.
package types

import "fmt"

// ModelID represents a whitelisted AI model identifier (maps to Solidity bytes32).
type ModelID [32]byte

// WorkerAddr represents a worker's Ethereum address (maps to Solidity address).
// Defined as [20]byte to avoid pulling in go-ethereum at this stage.
// Convert to common.Address at integration boundaries.
type WorkerAddr [20]byte

// ToHex returns the hex-encoded address with 0x prefix.
func (w WorkerAddr) ToHex() string {
	return fmt.Sprintf("0x%x", w[:])
}

// BlobRef represents a reference to an EIP-4844 blob (maps to Solidity bytes32).
type BlobRef [32]byte

// WorkerStatus represents the registration state of a worker.
type WorkerStatus uint8

const (
	// WorkerStatusActive indicates the worker is registered and accepting jobs.
	WorkerStatusActive WorkerStatus = iota
	// WorkerStatusInactive indicates the worker has deregistered or paused.
	WorkerStatusInactive
	// WorkerStatusSuspended indicates the worker has been penalised and barred from new jobs.
	WorkerStatusSuspended
)

// String returns a human-readable representation of the worker status.
func (s WorkerStatus) String() string {
	switch s {
	case WorkerStatusActive:
		return "Active"
	case WorkerStatusInactive:
		return "Inactive"
	case WorkerStatusSuspended:
		return "Suspended"
	default:
		return fmt.Sprintf("WorkerStatus(%d)", s)
	}
}
