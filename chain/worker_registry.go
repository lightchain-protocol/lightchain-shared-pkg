// Package chain defines Go interface definitions that mirror the Solidity
// smart contracts in the LightChain protocol. These are reference interfaces
// — consuming services define narrow subsets per the "accept interfaces,
// return structs" convention.
package chain

import (
	"context"
	"math/big"

	"github.com/lightchain/pkg/types"
)

// WorkerRegistry mirrors the Solidity IWorkerRegistry contract interface.
// It manages worker registration, model support, staking, and eligibility checks.
//
// Amendment A-2: single global stake replaces per-model staking. Workers register
// with msg.value as their stake, top up or withdraw via dedicated functions, and
// slashing is expressed in basis points against the global stake.
type WorkerRegistry interface {
	// RegisterWorker registers msg.sender as a worker with the given ECDH
	// encryption public key, staking the sent value. Solidity: registerWorker(bytes) payable.
	RegisterWorker(ctx context.Context, encryptionPubKey []byte, stake *big.Int) error

	// DeregisterWorker removes msg.sender from the active worker set.
	// Solidity: deregisterWorker().
	DeregisterWorker(ctx context.Context) error

	// AddSupportedModel adds a model to the worker's supported set.
	// Solidity: addSupportedModel(bytes32).
	AddSupportedModel(ctx context.Context, modelID types.ModelID) error

	// RemoveSupportedModel removes a model from the worker's supported set.
	// Solidity: removeSupportedModel(bytes32).
	RemoveSupportedModel(ctx context.Context, modelID types.ModelID) error

	// --- Staking (A-2: global stake) ---

	// TopUpStake adds msg.value to the worker's global stake.
	// Solidity: topUpStake() payable.
	TopUpStake(ctx context.Context, amount *big.Int) error

	// WithdrawStake withdraws the specified amount from the worker's global stake.
	// Solidity: withdrawStake(uint256).
	WithdrawStake(ctx context.Context, amount *big.Int) error

	// GetWorkerStake returns the current global stake for a worker.
	// Solidity: getWorkerStake(address) view returns (uint256).
	GetWorkerStake(ctx context.Context, worker types.WorkerAddr) (*big.Int, error)

	// --- Admin / Internal ---

	// Slash slashes a worker's stake by the given basis points.
	// Solidity: slash(address, uint256).
	Slash(ctx context.Context, worker types.WorkerAddr, bps *big.Int) error

	// IncrementActiveJobs increments the per-model active job count for a worker.
	// Solidity: incrementActiveJobs(address, bytes32).
	IncrementActiveJobs(ctx context.Context, worker types.WorkerAddr, modelID types.ModelID) error

	// DecrementActiveJobs decrements the per-model active job count for a worker.
	// Solidity: decrementActiveJobs(address, bytes32).
	DecrementActiveJobs(ctx context.Context, worker types.WorkerAddr, modelID types.ModelID) error

	// IsWorkerRegistered returns whether a worker is registered.
	// Solidity: isWorkerRegistered(address) view returns (bool).
	IsWorkerRegistered(ctx context.Context, worker types.WorkerAddr) (bool, error)

	// --- View functions ---

	// IsEligible checks whether a worker is eligible to serve a given model.
	// Solidity: isEligible(address, bytes32) view returns (bool).
	IsEligible(ctx context.Context, worker types.WorkerAddr, modelID types.ModelID) (bool, error)

	// GetEligibleWorkers returns all workers eligible to serve a given model.
	// Solidity: getEligibleWorkers(bytes32) view returns (address[]).
	GetEligibleWorkers(ctx context.Context, modelID types.ModelID) ([]types.WorkerAddr, error)

	// GetWorkerEncryptionKey returns the ECDH public key for a worker.
	// Solidity: getWorkerEncryptionKey(address) view returns (bytes).
	GetWorkerEncryptionKey(ctx context.Context, worker types.WorkerAddr) ([]byte, error)

	// SelectEligibleWorker selects a pseudo-random eligible worker for a model,
	// excluding the given set of addresses. Used by the dispatcher for reassignment.
	// Solidity: selectEligibleWorker(bytes32, address[], uint256) view returns (address).
	SelectEligibleWorker(ctx context.Context, modelID types.ModelID, excluded []types.WorkerAddr, seed *big.Int) (types.WorkerAddr, error)
}
