package chain

import (
	"context"
	"math/big"

	"github.com/lightchain/pkg/types"
)

// AIConfig mirrors the Solidity IAIConfig read-only contract interface.
// It provides access to protocol configuration parameters including model fees,
// slashing rates, dispute settings, and infrastructure timeouts.
//
// The full interface is defined here as a reference. Consuming services should
// define narrow subsets:
//   - Dispatcher: model fees, min stake
//   - Worker: timeouts, max output tokens
//   - Disputer: similarity thresholds, sampling rate, dispute parameters
type AIConfig interface {
	// --- Model Parameters ---

	// GetModelFee returns the flat fee for a model.
	// Solidity: getModelFee(bytes32) view returns (uint256).
	GetModelFee(ctx context.Context, modelID types.ModelID) (*big.Int, error)

	// GetMinWorkerStake returns the minimum global stake required for a worker.
	// Amendment A-2: replaces per-model GetModelMinStake — single global stake.
	// Solidity: getMinWorkerStake() view returns (uint256).
	GetMinWorkerStake(ctx context.Context) (*big.Int, error)

	// GetModelMaxOutputTokens returns the maximum output tokens for a model.
	GetModelMaxOutputTokens(ctx context.Context, modelID types.ModelID) (*big.Int, error)

	// GetSuspensionThreshold returns the offense count at which a worker is suspended.
	// Solidity: getSuspensionThreshold() view returns (uint256).
	GetSuspensionThreshold(ctx context.Context) (*big.Int, error)

	// GetSuspensionCooldown returns the time a worker must wait after suspension.
	// Solidity: getSuspensionCooldown() view returns (uint256).
	GetSuspensionCooldown(ctx context.Context) (*big.Int, error)

	// IsModelEnabled returns whether a model is currently enabled.
	// Solidity: isModelEnabled(bytes32) view returns (bool).
	IsModelEnabled(ctx context.Context, modelID types.ModelID) (bool, error)

	// --- Infrastructure Addresses ---

	// GetDispatcherAddress returns the address of the dispatcher contract.
	// Solidity: getDispatcherAddress() view returns (address).
	GetDispatcherAddress(ctx context.Context) (types.WorkerAddr, error)

	// GetDisputerAddress returns the address of the disputer contract.
	// Solidity: getDisputerAddress() view returns (address).
	GetDisputerAddress(ctx context.Context) (types.WorkerAddr, error)

	// --- Fee Distribution ---

	// GetWorkerFeeBps returns the worker's share of job fees in basis points.
	GetWorkerFeeBps(ctx context.Context) (*big.Int, error)

	// GetProtocolFeeBps returns the protocol's share of job fees in basis points.
	GetProtocolFeeBps(ctx context.Context) (*big.Int, error)

	// GetBurnFeeBps returns the burn share of job fees in basis points.
	GetBurnFeeBps(ctx context.Context) (*big.Int, error)

	// --- Slashing ---

	// GetTimeoutSlashBps returns the slash rate for ack timeout in basis points.
	GetTimeoutSlashBps(ctx context.Context) (*big.Int, error)

	// GetCompletionTimeoutSlashBps returns the slash rate for completion timeout in basis points.
	GetCompletionTimeoutSlashBps(ctx context.Context) (*big.Int, error)

	// GetDisputeSlashBps returns the slash rate for lost disputes in basis points.
	GetDisputeSlashBps(ctx context.Context) (*big.Int, error)

	// --- Disputes ---

	// GetDisputeBondMultiplier returns the multiplier for dispute bond calculation.
	GetDisputeBondMultiplier(ctx context.Context) (*big.Int, error)

	// GetDisputeWindow returns the dispute window duration in seconds.
	GetDisputeWindow(ctx context.Context) (*big.Int, error)

	// GetResolutionTimeout returns the maximum time for dispute resolution in seconds.
	GetResolutionTimeout(ctx context.Context) (*big.Int, error)

	// GetSamplingRateBps returns the automatic dispute sampling rate in basis points.
	GetSamplingRateBps(ctx context.Context) (*big.Int, error)

	// GetSimilarityThreshold returns the similarity threshold for dispute resolution.
	GetSimilarityThreshold(ctx context.Context) (*big.Int, error)

	// GetCanarySimilarityThreshold returns the similarity threshold for canary jobs.
	GetCanarySimilarityThreshold(ctx context.Context) (*big.Int, error)

	// --- Infrastructure ---

	// GetBlobRetentionPeriod returns the blob retention period in seconds.
	GetBlobRetentionPeriod(ctx context.Context) (*big.Int, error)

	// GetSessionInactivityTimeout returns the session inactivity timeout in seconds.
	GetSessionInactivityTimeout(ctx context.Context) (*big.Int, error)

	// GetMaxReassignments returns the maximum number of session reassignments.
	GetMaxReassignments(ctx context.Context) (*big.Int, error)

	// GetAckTimeout returns the job acknowledgement timeout in seconds.
	GetAckTimeout(ctx context.Context) (*big.Int, error)

	// GetCompletionTimeout returns the job completion timeout in seconds.
	GetCompletionTimeout(ctx context.Context) (*big.Int, error)

	// --- Fee Calculation ---

	// CalculateJobFee calculates the total fee for a job given the model.
	// Solidity: calculateJobFee(bytes32) view returns (uint256).
	CalculateJobFee(ctx context.Context, modelID types.ModelID) (*big.Int, error)
}
