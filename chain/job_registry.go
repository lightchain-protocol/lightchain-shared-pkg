package chain

import (
	"context"
	"math/big"

	"github.com/lightchain/pkg/types"
)

// JobRegistry mirrors the Solidity IJobRegistry contract interface.
// It manages sessions, jobs, and the dispute lifecycle.
type JobRegistry interface {
	// --- Sessions ---

	// Nonce returns the current session-creation nonce for the given account.
	// Solidity: nonce(address) view returns (uint256).
	Nonce(ctx context.Context, account types.WorkerAddr) (uint64, error)

	// CreateSession creates a new AI inference session between a consumer and a
	// worker. Returns the new session ID. Solidity: createSession(...) payable returns (uint256).
	CreateSession(
		ctx context.Context,
		modelID types.ModelID,
		worker types.WorkerAddr,
		encWorkerKey []byte,
		encDisputerKey []byte,
		dispatcherSignature []byte,
		expiry uint64,
		deposit *big.Int,
	) (types.SessionID, error)

	// ReassignSession triggers worker reassignment for a session.
	// Solidity: reassignSession(uint256).
	ReassignSession(ctx context.Context, sessionID types.SessionID) error

	// UpdateSessionKey updates the encrypted session keys after reassignment.
	// Solidity: updateSessionKey(uint256, bytes, bytes).
	UpdateSessionKey(ctx context.Context, sessionID types.SessionID, encWorkerKey []byte, encDisputerKey []byte) error

	// CloseSession closes a session and refunds remaining deposit.
	// Solidity: closeSession(uint256).
	CloseSession(ctx context.Context, sessionID types.SessionID) error

	// --- Jobs ---

	// SubmitJob submits a new inference job within a session. Returns the new
	// job ID. Solidity: submitJob(uint256, bytes32[], uint256) payable returns (uint256).
	SubmitJob(
		ctx context.Context,
		sessionID types.SessionID,
		blobHashes [][32]byte,
		dataLength uint64,
		fee *big.Int,
	) (types.JobID, error)

	// AcknowledgeJob marks a job as acknowledged by the worker.
	// Solidity: acknowledgeJob(uint256).
	AcknowledgeJob(ctx context.Context, jobID types.JobID) error

	// CompleteJob marks a job as completed with response blob hashes and the
	// keccak256 hash of the full encrypted response ciphertext.
	// Solidity: completeJob(uint256, bytes32[], bytes32).
	CompleteJob(
		ctx context.Context,
		jobID types.JobID,
		responseBlobHashes [][32]byte,
		responseCiphertextHash [32]byte,
	) error

	// --- Disputes ---

	// DisputeJob files a dispute against a completed job.
	// Solidity: disputeJob(uint256) payable.
	DisputeJob(ctx context.Context, jobID types.JobID, bond *big.Int) error

	// ResolveDispute resolves a dispute with the re-execution result.
	// Solidity: resolveDispute(uint256, bool, bytes32[], uint256).
	ResolveDispute(
		ctx context.Context,
		jobID types.JobID,
		workerGuilty bool,
		reExecutionBlobHashes [][32]byte,
		similarityScore uint64,
	) error

	// ClaimTimeout claims a timeout on a job (ack or completion timeout).
	// Solidity: claimTimeout(uint256).
	ClaimTimeout(ctx context.Context, jobID types.JobID) error

	// ReleaseJob releases escrowed fees after the dispute window expires.
	// Solidity: releaseJob(uint256).
	ReleaseJob(ctx context.Context, jobID types.JobID) error

	// ReleaseJobs batch-releases escrowed fees for multiple jobs.
	// Solidity: releaseJobs(uint256[]).
	ReleaseJobs(ctx context.Context, jobIDs []types.JobID) error

	// DisputeResponseMismatch raises a dispute when worker's off-chain response
	// doesn't match on-chain commitment. Solidity: disputeResponseMismatch(uint256, bytes, bytes).
	DisputeResponseMismatch(
		ctx context.Context,
		jobID types.JobID,
		ciphertext []byte,
		signature []byte,
	) error
}
