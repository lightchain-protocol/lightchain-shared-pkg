package types

import (
	"fmt"
	"math/big"
)

// SessionID uniquely identifies a session within the JobRegistry contract.
type SessionID uint64

// SessionStatus represents the lifecycle state of a session.
// Values match the Solidity enum in JobRegistry.
type SessionStatus uint8

const (
	// SessionStatusActive indicates the session is open and accepting jobs.
	SessionStatusActive SessionStatus = iota
	// SessionStatusReassigning indicates the worker is being swapped and the session is temporarily paused.
	SessionStatusReassigning
	// SessionStatusClosed indicates the session has ended and no further jobs will be accepted.
	SessionStatusClosed
)

// String returns a human-readable representation of the session status.
func (s SessionStatus) String() string {
	switch s {
	case SessionStatusActive:
		return "Active"
	case SessionStatusReassigning:
		return "Reassigning"
	case SessionStatusClosed:
		return "Closed"
	default:
		return fmt.Sprintf("SessionStatus(%d)", s)
	}
}

// Session represents an AI inference session between a consumer and a worker.
// Maps to the Solidity Session struct in JobRegistry.
type Session struct {
	// ID is the unique on-chain session identifier.
	ID SessionID
	// User is the consumer's Ethereum address (reuses WorkerAddr [20]byte type).
	User WorkerAddr
	// ModelID identifies the AI model requested for this session.
	ModelID ModelID
	// Worker is the assigned worker's Ethereum address.
	Worker WorkerAddr
	// Status tracks the session's lifecycle state.
	Status SessionStatus
	// EncWorkerKey is the symmetric key encrypted with the worker's public key.
	EncWorkerKey []byte
	// EncDisputerKey is the symmetric key encrypted with the disputer's public key.
	EncDisputerKey []byte
	// JobCount is the total number of jobs submitted in this session.
	JobCount uint64
	// LastActivityAt is the Unix timestamp of the most recent job submission.
	LastActivityAt uint64
	// ReassignCount is the number of times the session's worker has been replaced.
	ReassignCount uint64
	// Deposit is the escrowed amount (in wei) held for this session.
	Deposit *big.Int
	// ExcludedWorkers is the list of worker addresses that may not be assigned
	// to this session (e.g. previously deregistered or reassigned workers).
	ExcludedWorkers []WorkerAddr
}
