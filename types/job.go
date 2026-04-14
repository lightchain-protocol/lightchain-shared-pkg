package types

import (
	"fmt"
	"math/big"
)

// JobID uniquely identifies a job within the JobRegistry contract.
type JobID uint64

// JobState represents the lifecycle state of a job.
// Values match the Solidity enum in JobRegistry.
type JobState uint8

const (
	// JobStateSubmitted is the initial state when a job is posted on-chain.
	JobStateSubmitted JobState = iota
	// JobStateAcknowledged indicates the worker has accepted the job.
	JobStateAcknowledged
	// JobStateCompleted indicates the worker has submitted a result.
	JobStateCompleted
	// JobStateTimedOut indicates the job was not acknowledged or completed within its deadline.
	JobStateTimedOut
	// JobStateDisputed indicates the consumer or disputer has challenged the result.
	JobStateDisputed
	// JobStateResolved indicates the dispute has been adjudicated.
	JobStateResolved
	// JobStateReleased indicates the escrowed fee has been released to the worker.
	JobStateReleased
)

// String returns a human-readable representation of the job state.
func (s JobState) String() string {
	switch s {
	case JobStateSubmitted:
		return "Submitted"
	case JobStateAcknowledged:
		return "Acknowledged"
	case JobStateCompleted:
		return "Completed"
	case JobStateTimedOut:
		return "TimedOut"
	case JobStateDisputed:
		return "Disputed"
	case JobStateResolved:
		return "Resolved"
	case JobStateReleased:
		return "Released"
	default:
		return fmt.Sprintf("JobState(%d)", s)
	}
}

// Job represents a single AI inference job within a session.
// Maps to the Solidity Job struct in JobRegistry.
type Job struct {
	// ID is the unique on-chain job identifier.
	ID JobID
	// SessionID is the session this job belongs to.
	SessionID SessionID
	// Worker is the address of the worker assigned to execute this job.
	Worker WorkerAddr
	// State tracks the job's lifecycle stage.
	State JobState
	// EscrowedFee is the amount (in wei) locked on-chain until the job is settled.
	EscrowedFee *big.Int
	// PromptBlobHash is the EIP-4844 blob hash containing the input prompt data.
	PromptBlobHash [32]byte
	// ResponseBlobHash is the EIP-4844 blob hash containing the worker's response.
	ResponseBlobHash [32]byte
	// SubmittedAt is the Unix timestamp when the job was submitted on-chain.
	SubmittedAt uint64
	// AckTimestamp is the Unix timestamp when the worker acknowledged the job.
	AckTimestamp uint64
	// CompletedAt is the Unix timestamp when the worker submitted the result.
	CompletedAt uint64
	// Deadline is the Unix timestamp by which the worker must complete the job.
	Deadline uint64
	// DisputeFiler is the address of the party that filed a dispute, if any.
	DisputeFiler WorkerAddr
	// DisputeBond is the bond amount (in wei) posted by the dispute filer.
	DisputeBond *big.Int
	// ReExecutionBlobHash is the EIP-4844 blob hash of the disputer's re-execution result.
	ReExecutionBlobHash [32]byte
	// SimilarityScore is the similarity score computed during dispute resolution.
	SimilarityScore uint64
	// DisputeCreatedAt is the Unix timestamp when the dispute was created.
	DisputeCreatedAt uint64
	// ResponseCiphertextHash is keccak256 of the full encrypted response payload.
	ResponseCiphertextHash [32]byte
	// SubmitBlockNumber is the EL block number when the job was submitted.
	SubmitBlockNumber uint64
	// CompletionBlockNumber is the EL block number when the job was completed.
	CompletionBlockNumber uint64
}
