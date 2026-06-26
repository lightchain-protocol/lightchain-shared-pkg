// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IJobRegistryJob is an auto generated low-level Go binding around an user-defined struct.
type IJobRegistryJob struct {
	SessionId              *big.Int
	Worker                 common.Address
	State                  uint8
	EscrowedFee            *big.Int
	PromptBlobHash         [32]byte
	ResponseBlobHash       [32]byte
	SubmittedAt            *big.Int
	AckTimestamp           *big.Int
	CompletedAt            *big.Int
	Deadline               *big.Int
	DisputeFiler           common.Address
	DisputeBond            *big.Int
	ReExecutionBlobHash    [32]byte
	SimilarityScore        *big.Int
	DisputeCreatedAt       *big.Int
	ResponseCiphertextHash [32]byte
	SubmitBlockNumber      *big.Int
	CompletionBlockNumber  *big.Int
}

// IJobRegistrySession is an auto generated low-level Go binding around an user-defined struct.
type IJobRegistrySession struct {
	User            common.Address
	ModelId         [32]byte
	Worker          common.Address
	Status          uint8
	EncWorkerKey    []byte
	EncDisputerKey  []byte
	JobCount        *big.Int
	LastActivityAt  *big.Int
	ReassignCount   *big.Int
	Deposit         *big.Int
	ExcludedWorkers []common.Address
}

// JobRegistryMetaData contains all meta data concerning the JobRegistry contract.
var JobRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acknowledgeJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRefund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimTimeout\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeSession\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"responseBlobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"responseCiphertextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSession\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dispatcherSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegateAllowance\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndAuthorize\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disputeJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disputeResponseMismatch\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ciphertext\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIJobRegistry.Job\",\"components\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumIJobRegistry.JobState\"},{\"name\":\"escrowedFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"promptBlobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"responseBlobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"submittedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ackTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"disputeFiler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"disputeBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reExecutionBlobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"similarityScore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"disputeCreatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"responseCiphertextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"submitBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completionBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSession\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIJobRegistry.Session\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIJobRegistry.SessionStatus\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jobCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastActivityAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reassignCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"excludedWorkers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_workerRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aiConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feePool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegateAuthorized\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingRefund\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prepaidBalanceOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reassignSession\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseJobs\",\"inputs\":[{\"name\":\"jobIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"workerGuilty\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reExecutionBlobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_similarityScore\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegateAllowance\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegateAuthorization\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitJob\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"submitJobOnBehalf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sunsetGuardian\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSessionKey\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dispatcherSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBalance\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"workerBalance\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegateAllowanceSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateAuthorizationSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeCreated\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"disputer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"workerGuilty\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"similarityScore\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeDistributed\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"workerShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feePoolShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianUpdated\",\"inputs\":[{\"name\":\"oldGuardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGuardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobAcknowledged\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobCompleted\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"responseBlobHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"responseCiphertextHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobReleased\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"workerShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feePoolShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobSubmitted\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobTimedOut\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundClaimed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundCreditedToBalance\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundEscrowed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionClosed\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionCreated\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionKeyUpdated\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionReassigned\",\"inputs\":[{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newWorker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrew\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerWithdrawal\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// JobRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use JobRegistryMetaData.ABI instead.
var JobRegistryABI = JobRegistryMetaData.ABI

// JobRegistry is an auto generated Go binding around an Ethereum contract.
type JobRegistry struct {
	JobRegistryCaller     // Read-only binding to the contract
	JobRegistryTransactor // Write-only binding to the contract
	JobRegistryFilterer   // Log filterer for contract events
}

// JobRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobRegistrySession struct {
	Contract     *JobRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobRegistryCallerSession struct {
	Contract *JobRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// JobRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobRegistryTransactorSession struct {
	Contract     *JobRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// JobRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobRegistryRaw struct {
	Contract *JobRegistry // Generic contract binding to access the raw methods on
}

// JobRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobRegistryCallerRaw struct {
	Contract *JobRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// JobRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobRegistryTransactorRaw struct {
	Contract *JobRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobRegistry creates a new instance of JobRegistry, bound to a specific deployed contract.
func NewJobRegistry(address common.Address, backend bind.ContractBackend) (*JobRegistry, error) {
	contract, err := bindJobRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobRegistry{JobRegistryCaller: JobRegistryCaller{contract: contract}, JobRegistryTransactor: JobRegistryTransactor{contract: contract}, JobRegistryFilterer: JobRegistryFilterer{contract: contract}}, nil
}

// NewJobRegistryCaller creates a new read-only instance of JobRegistry, bound to a specific deployed contract.
func NewJobRegistryCaller(address common.Address, caller bind.ContractCaller) (*JobRegistryCaller, error) {
	contract, err := bindJobRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobRegistryCaller{contract: contract}, nil
}

// NewJobRegistryTransactor creates a new write-only instance of JobRegistry, bound to a specific deployed contract.
func NewJobRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*JobRegistryTransactor, error) {
	contract, err := bindJobRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobRegistryTransactor{contract: contract}, nil
}

// NewJobRegistryFilterer creates a new log filterer instance of JobRegistry, bound to a specific deployed contract.
func NewJobRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*JobRegistryFilterer, error) {
	contract, err := bindJobRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobRegistryFilterer{contract: contract}, nil
}

// bindJobRegistry binds a generic wrapper to an already deployed contract.
func bindJobRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobRegistry *JobRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobRegistry.Contract.JobRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobRegistry *JobRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.Contract.JobRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobRegistry *JobRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobRegistry.Contract.JobRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobRegistry *JobRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobRegistry *JobRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobRegistry *JobRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_JobRegistry *JobRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_JobRegistry *JobRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _JobRegistry.Contract.UPGRADEINTERFACEVERSION(&_JobRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_JobRegistry *JobRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _JobRegistry.Contract.UPGRADEINTERFACEVERSION(&_JobRegistry.CallOpts)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address user, address delegate) view returns(uint256)
func (_JobRegistry *JobRegistryCaller) DelegateAllowance(opts *bind.CallOpts, user common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "delegateAllowance", user, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address user, address delegate) view returns(uint256)
func (_JobRegistry *JobRegistrySession) DelegateAllowance(user common.Address, delegate common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.DelegateAllowance(&_JobRegistry.CallOpts, user, delegate)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address user, address delegate) view returns(uint256)
func (_JobRegistry *JobRegistryCallerSession) DelegateAllowance(user common.Address, delegate common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.DelegateAllowance(&_JobRegistry.CallOpts, user, delegate)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,uint8,uint256,bytes32,bytes32,uint256,uint256,uint256,uint256,address,uint256,bytes32,uint256,uint256,bytes32,uint256,uint256))
func (_JobRegistry *JobRegistryCaller) GetJob(opts *bind.CallOpts, jobId *big.Int) (IJobRegistryJob, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "getJob", jobId)

	if err != nil {
		return *new(IJobRegistryJob), err
	}

	out0 := *abi.ConvertType(out[0], new(IJobRegistryJob)).(*IJobRegistryJob)

	return out0, err

}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,uint8,uint256,bytes32,bytes32,uint256,uint256,uint256,uint256,address,uint256,bytes32,uint256,uint256,bytes32,uint256,uint256))
func (_JobRegistry *JobRegistrySession) GetJob(jobId *big.Int) (IJobRegistryJob, error) {
	return _JobRegistry.Contract.GetJob(&_JobRegistry.CallOpts, jobId)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,uint8,uint256,bytes32,bytes32,uint256,uint256,uint256,uint256,address,uint256,bytes32,uint256,uint256,bytes32,uint256,uint256))
func (_JobRegistry *JobRegistryCallerSession) GetJob(jobId *big.Int) (IJobRegistryJob, error) {
	return _JobRegistry.Contract.GetJob(&_JobRegistry.CallOpts, jobId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((address,bytes32,address,uint8,bytes,bytes,uint256,uint256,uint256,uint256,address[]))
func (_JobRegistry *JobRegistryCaller) GetSession(opts *bind.CallOpts, sessionId *big.Int) (IJobRegistrySession, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "getSession", sessionId)

	if err != nil {
		return *new(IJobRegistrySession), err
	}

	out0 := *abi.ConvertType(out[0], new(IJobRegistrySession)).(*IJobRegistrySession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((address,bytes32,address,uint8,bytes,bytes,uint256,uint256,uint256,uint256,address[]))
func (_JobRegistry *JobRegistrySession) GetSession(sessionId *big.Int) (IJobRegistrySession, error) {
	return _JobRegistry.Contract.GetSession(&_JobRegistry.CallOpts, sessionId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((address,bytes32,address,uint8,bytes,bytes,uint256,uint256,uint256,uint256,address[]))
func (_JobRegistry *JobRegistryCallerSession) GetSession(sessionId *big.Int) (IJobRegistrySession, error) {
	return _JobRegistry.Contract.GetSession(&_JobRegistry.CallOpts, sessionId)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_JobRegistry *JobRegistryCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_JobRegistry *JobRegistrySession) Guardian() (common.Address, error) {
	return _JobRegistry.Contract.Guardian(&_JobRegistry.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_JobRegistry *JobRegistryCallerSession) Guardian() (common.Address, error) {
	return _JobRegistry.Contract.Guardian(&_JobRegistry.CallOpts)
}

// IsDelegateAuthorized is a free data retrieval call binding the contract method 0xed45975d.
//
// Solidity: function isDelegateAuthorized(address user, address delegate) view returns(bool)
func (_JobRegistry *JobRegistryCaller) IsDelegateAuthorized(opts *bind.CallOpts, user common.Address, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "isDelegateAuthorized", user, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegateAuthorized is a free data retrieval call binding the contract method 0xed45975d.
//
// Solidity: function isDelegateAuthorized(address user, address delegate) view returns(bool)
func (_JobRegistry *JobRegistrySession) IsDelegateAuthorized(user common.Address, delegate common.Address) (bool, error) {
	return _JobRegistry.Contract.IsDelegateAuthorized(&_JobRegistry.CallOpts, user, delegate)
}

// IsDelegateAuthorized is a free data retrieval call binding the contract method 0xed45975d.
//
// Solidity: function isDelegateAuthorized(address user, address delegate) view returns(bool)
func (_JobRegistry *JobRegistryCallerSession) IsDelegateAuthorized(user common.Address, delegate common.Address) (bool, error) {
	return _JobRegistry.Contract.IsDelegateAuthorized(&_JobRegistry.CallOpts, user, delegate)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address account) view returns(uint256)
func (_JobRegistry *JobRegistryCaller) Nonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "nonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address account) view returns(uint256)
func (_JobRegistry *JobRegistrySession) Nonce(account common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.Nonce(&_JobRegistry.CallOpts, account)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address account) view returns(uint256)
func (_JobRegistry *JobRegistryCallerSession) Nonce(account common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.Nonce(&_JobRegistry.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobRegistry *JobRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobRegistry *JobRegistrySession) Owner() (common.Address, error) {
	return _JobRegistry.Contract.Owner(&_JobRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobRegistry *JobRegistryCallerSession) Owner() (common.Address, error) {
	return _JobRegistry.Contract.Owner(&_JobRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JobRegistry *JobRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JobRegistry *JobRegistrySession) Paused() (bool, error) {
	return _JobRegistry.Contract.Paused(&_JobRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JobRegistry *JobRegistryCallerSession) Paused() (bool, error) {
	return _JobRegistry.Contract.Paused(&_JobRegistry.CallOpts)
}

// PendingRefund is a free data retrieval call binding the contract method 0x99d82c5f.
//
// Solidity: function pendingRefund(address account) view returns(uint256)
func (_JobRegistry *JobRegistryCaller) PendingRefund(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "pendingRefund", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRefund is a free data retrieval call binding the contract method 0x99d82c5f.
//
// Solidity: function pendingRefund(address account) view returns(uint256)
func (_JobRegistry *JobRegistrySession) PendingRefund(account common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.PendingRefund(&_JobRegistry.CallOpts, account)
}

// PendingRefund is a free data retrieval call binding the contract method 0x99d82c5f.
//
// Solidity: function pendingRefund(address account) view returns(uint256)
func (_JobRegistry *JobRegistryCallerSession) PendingRefund(account common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.PendingRefund(&_JobRegistry.CallOpts, account)
}

// PrepaidBalanceOf is a free data retrieval call binding the contract method 0xf990c29b.
//
// Solidity: function prepaidBalanceOf(address user) view returns(uint256)
func (_JobRegistry *JobRegistryCaller) PrepaidBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "prepaidBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrepaidBalanceOf is a free data retrieval call binding the contract method 0xf990c29b.
//
// Solidity: function prepaidBalanceOf(address user) view returns(uint256)
func (_JobRegistry *JobRegistrySession) PrepaidBalanceOf(user common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.PrepaidBalanceOf(&_JobRegistry.CallOpts, user)
}

// PrepaidBalanceOf is a free data retrieval call binding the contract method 0xf990c29b.
//
// Solidity: function prepaidBalanceOf(address user) view returns(uint256)
func (_JobRegistry *JobRegistryCallerSession) PrepaidBalanceOf(user common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.PrepaidBalanceOf(&_JobRegistry.CallOpts, user)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_JobRegistry *JobRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_JobRegistry *JobRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _JobRegistry.Contract.ProxiableUUID(&_JobRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_JobRegistry *JobRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _JobRegistry.Contract.ProxiableUUID(&_JobRegistry.CallOpts)
}

// WorkerBalance is a free data retrieval call binding the contract method 0x78904a35.
//
// Solidity: function workerBalance(address worker) view returns(uint256)
func (_JobRegistry *JobRegistryCaller) WorkerBalance(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobRegistry.contract.Call(opts, &out, "workerBalance", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkerBalance is a free data retrieval call binding the contract method 0x78904a35.
//
// Solidity: function workerBalance(address worker) view returns(uint256)
func (_JobRegistry *JobRegistrySession) WorkerBalance(worker common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.WorkerBalance(&_JobRegistry.CallOpts, worker)
}

// WorkerBalance is a free data retrieval call binding the contract method 0x78904a35.
//
// Solidity: function workerBalance(address worker) view returns(uint256)
func (_JobRegistry *JobRegistryCallerSession) WorkerBalance(worker common.Address) (*big.Int, error) {
	return _JobRegistry.Contract.WorkerBalance(&_JobRegistry.CallOpts, worker)
}

// AcknowledgeJob is a paid mutator transaction binding the contract method 0x9f6c3cd9.
//
// Solidity: function acknowledgeJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactor) AcknowledgeJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "acknowledgeJob", jobId)
}

// AcknowledgeJob is a paid mutator transaction binding the contract method 0x9f6c3cd9.
//
// Solidity: function acknowledgeJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistrySession) AcknowledgeJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.AcknowledgeJob(&_JobRegistry.TransactOpts, jobId)
}

// AcknowledgeJob is a paid mutator transaction binding the contract method 0x9f6c3cd9.
//
// Solidity: function acknowledgeJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactorSession) AcknowledgeJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.AcknowledgeJob(&_JobRegistry.TransactOpts, jobId)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_JobRegistry *JobRegistryTransactor) ClaimRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "claimRefund")
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_JobRegistry *JobRegistrySession) ClaimRefund() (*types.Transaction, error) {
	return _JobRegistry.Contract.ClaimRefund(&_JobRegistry.TransactOpts)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_JobRegistry *JobRegistryTransactorSession) ClaimRefund() (*types.Transaction, error) {
	return _JobRegistry.Contract.ClaimRefund(&_JobRegistry.TransactOpts)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x86e773f1.
//
// Solidity: function claimTimeout(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactor) ClaimTimeout(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "claimTimeout", jobId)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x86e773f1.
//
// Solidity: function claimTimeout(uint256 jobId) returns()
func (_JobRegistry *JobRegistrySession) ClaimTimeout(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ClaimTimeout(&_JobRegistry.TransactOpts, jobId)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x86e773f1.
//
// Solidity: function claimTimeout(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactorSession) ClaimTimeout(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ClaimTimeout(&_JobRegistry.TransactOpts, jobId)
}

// CloseSession is a paid mutator transaction binding the contract method 0xa3bfdf47.
//
// Solidity: function closeSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistryTransactor) CloseSession(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "closeSession", sessionId)
}

// CloseSession is a paid mutator transaction binding the contract method 0xa3bfdf47.
//
// Solidity: function closeSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistrySession) CloseSession(sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.CloseSession(&_JobRegistry.TransactOpts, sessionId)
}

// CloseSession is a paid mutator transaction binding the contract method 0xa3bfdf47.
//
// Solidity: function closeSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistryTransactorSession) CloseSession(sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.CloseSession(&_JobRegistry.TransactOpts, sessionId)
}

// CompleteJob is a paid mutator transaction binding the contract method 0x2c842a15.
//
// Solidity: function completeJob(uint256 jobId, bytes32 responseBlobHash, bytes32 responseCiphertextHash) returns()
func (_JobRegistry *JobRegistryTransactor) CompleteJob(opts *bind.TransactOpts, jobId *big.Int, responseBlobHash [32]byte, responseCiphertextHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "completeJob", jobId, responseBlobHash, responseCiphertextHash)
}

// CompleteJob is a paid mutator transaction binding the contract method 0x2c842a15.
//
// Solidity: function completeJob(uint256 jobId, bytes32 responseBlobHash, bytes32 responseCiphertextHash) returns()
func (_JobRegistry *JobRegistrySession) CompleteJob(jobId *big.Int, responseBlobHash [32]byte, responseCiphertextHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.CompleteJob(&_JobRegistry.TransactOpts, jobId, responseBlobHash, responseCiphertextHash)
}

// CompleteJob is a paid mutator transaction binding the contract method 0x2c842a15.
//
// Solidity: function completeJob(uint256 jobId, bytes32 responseBlobHash, bytes32 responseCiphertextHash) returns()
func (_JobRegistry *JobRegistryTransactorSession) CompleteJob(jobId *big.Int, responseBlobHash [32]byte, responseCiphertextHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.CompleteJob(&_JobRegistry.TransactOpts, jobId, responseBlobHash, responseCiphertextHash)
}

// CreateSession is a paid mutator transaction binding the contract method 0xe80116b4.
//
// Solidity: function createSession(bytes32 modelId, address worker, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) payable returns(uint256 sessionId)
func (_JobRegistry *JobRegistryTransactor) CreateSession(opts *bind.TransactOpts, modelId [32]byte, worker common.Address, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "createSession", modelId, worker, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// CreateSession is a paid mutator transaction binding the contract method 0xe80116b4.
//
// Solidity: function createSession(bytes32 modelId, address worker, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) payable returns(uint256 sessionId)
func (_JobRegistry *JobRegistrySession) CreateSession(modelId [32]byte, worker common.Address, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.CreateSession(&_JobRegistry.TransactOpts, modelId, worker, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// CreateSession is a paid mutator transaction binding the contract method 0xe80116b4.
//
// Solidity: function createSession(bytes32 modelId, address worker, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) payable returns(uint256 sessionId)
func (_JobRegistry *JobRegistryTransactorSession) CreateSession(modelId [32]byte, worker common.Address, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.CreateSession(&_JobRegistry.TransactOpts, modelId, worker, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_JobRegistry *JobRegistryTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_JobRegistry *JobRegistrySession) Deposit() (*types.Transaction, error) {
	return _JobRegistry.Contract.Deposit(&_JobRegistry.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_JobRegistry *JobRegistryTransactorSession) Deposit() (*types.Transaction, error) {
	return _JobRegistry.Contract.Deposit(&_JobRegistry.TransactOpts)
}

// DepositAndAuthorize is a paid mutator transaction binding the contract method 0x882fe3c9.
//
// Solidity: function depositAndAuthorize(address delegate) payable returns()
func (_JobRegistry *JobRegistryTransactor) DepositAndAuthorize(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "depositAndAuthorize", delegate)
}

// DepositAndAuthorize is a paid mutator transaction binding the contract method 0x882fe3c9.
//
// Solidity: function depositAndAuthorize(address delegate) payable returns()
func (_JobRegistry *JobRegistrySession) DepositAndAuthorize(delegate common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.DepositAndAuthorize(&_JobRegistry.TransactOpts, delegate)
}

// DepositAndAuthorize is a paid mutator transaction binding the contract method 0x882fe3c9.
//
// Solidity: function depositAndAuthorize(address delegate) payable returns()
func (_JobRegistry *JobRegistryTransactorSession) DepositAndAuthorize(delegate common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.DepositAndAuthorize(&_JobRegistry.TransactOpts, delegate)
}

// DisputeJob is a paid mutator transaction binding the contract method 0xd93d9beb.
//
// Solidity: function disputeJob(uint256 jobId) payable returns()
func (_JobRegistry *JobRegistryTransactor) DisputeJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "disputeJob", jobId)
}

// DisputeJob is a paid mutator transaction binding the contract method 0xd93d9beb.
//
// Solidity: function disputeJob(uint256 jobId) payable returns()
func (_JobRegistry *JobRegistrySession) DisputeJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.DisputeJob(&_JobRegistry.TransactOpts, jobId)
}

// DisputeJob is a paid mutator transaction binding the contract method 0xd93d9beb.
//
// Solidity: function disputeJob(uint256 jobId) payable returns()
func (_JobRegistry *JobRegistryTransactorSession) DisputeJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.DisputeJob(&_JobRegistry.TransactOpts, jobId)
}

// DisputeResponseMismatch is a paid mutator transaction binding the contract method 0x6d3bf718.
//
// Solidity: function disputeResponseMismatch(uint256 jobId, bytes ciphertext, bytes signature) returns()
func (_JobRegistry *JobRegistryTransactor) DisputeResponseMismatch(opts *bind.TransactOpts, jobId *big.Int, ciphertext []byte, signature []byte) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "disputeResponseMismatch", jobId, ciphertext, signature)
}

// DisputeResponseMismatch is a paid mutator transaction binding the contract method 0x6d3bf718.
//
// Solidity: function disputeResponseMismatch(uint256 jobId, bytes ciphertext, bytes signature) returns()
func (_JobRegistry *JobRegistrySession) DisputeResponseMismatch(jobId *big.Int, ciphertext []byte, signature []byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.DisputeResponseMismatch(&_JobRegistry.TransactOpts, jobId, ciphertext, signature)
}

// DisputeResponseMismatch is a paid mutator transaction binding the contract method 0x6d3bf718.
//
// Solidity: function disputeResponseMismatch(uint256 jobId, bytes ciphertext, bytes signature) returns()
func (_JobRegistry *JobRegistryTransactorSession) DisputeResponseMismatch(jobId *big.Int, ciphertext []byte, signature []byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.DisputeResponseMismatch(&_JobRegistry.TransactOpts, jobId, ciphertext, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _initialOwner, address _workerRegistry, address _aiConfig, address _treasury, address _feePool, address _guardian) returns()
func (_JobRegistry *JobRegistryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _workerRegistry common.Address, _aiConfig common.Address, _treasury common.Address, _feePool common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "initialize", _initialOwner, _workerRegistry, _aiConfig, _treasury, _feePool, _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _initialOwner, address _workerRegistry, address _aiConfig, address _treasury, address _feePool, address _guardian) returns()
func (_JobRegistry *JobRegistrySession) Initialize(_initialOwner common.Address, _workerRegistry common.Address, _aiConfig common.Address, _treasury common.Address, _feePool common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.Initialize(&_JobRegistry.TransactOpts, _initialOwner, _workerRegistry, _aiConfig, _treasury, _feePool, _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _initialOwner, address _workerRegistry, address _aiConfig, address _treasury, address _feePool, address _guardian) returns()
func (_JobRegistry *JobRegistryTransactorSession) Initialize(_initialOwner common.Address, _workerRegistry common.Address, _aiConfig common.Address, _treasury common.Address, _feePool common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.Initialize(&_JobRegistry.TransactOpts, _initialOwner, _workerRegistry, _aiConfig, _treasury, _feePool, _guardian)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JobRegistry *JobRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JobRegistry *JobRegistrySession) Pause() (*types.Transaction, error) {
	return _JobRegistry.Contract.Pause(&_JobRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_JobRegistry *JobRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _JobRegistry.Contract.Pause(&_JobRegistry.TransactOpts)
}

// ReassignSession is a paid mutator transaction binding the contract method 0x62ee75f5.
//
// Solidity: function reassignSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistryTransactor) ReassignSession(opts *bind.TransactOpts, sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "reassignSession", sessionId)
}

// ReassignSession is a paid mutator transaction binding the contract method 0x62ee75f5.
//
// Solidity: function reassignSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistrySession) ReassignSession(sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReassignSession(&_JobRegistry.TransactOpts, sessionId)
}

// ReassignSession is a paid mutator transaction binding the contract method 0x62ee75f5.
//
// Solidity: function reassignSession(uint256 sessionId) returns()
func (_JobRegistry *JobRegistryTransactorSession) ReassignSession(sessionId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReassignSession(&_JobRegistry.TransactOpts, sessionId)
}

// ReleaseJob is a paid mutator transaction binding the contract method 0x585fd57f.
//
// Solidity: function releaseJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactor) ReleaseJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "releaseJob", jobId)
}

// ReleaseJob is a paid mutator transaction binding the contract method 0x585fd57f.
//
// Solidity: function releaseJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistrySession) ReleaseJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReleaseJob(&_JobRegistry.TransactOpts, jobId)
}

// ReleaseJob is a paid mutator transaction binding the contract method 0x585fd57f.
//
// Solidity: function releaseJob(uint256 jobId) returns()
func (_JobRegistry *JobRegistryTransactorSession) ReleaseJob(jobId *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReleaseJob(&_JobRegistry.TransactOpts, jobId)
}

// ReleaseJobs is a paid mutator transaction binding the contract method 0x64dc01d5.
//
// Solidity: function releaseJobs(uint256[] jobIds) returns()
func (_JobRegistry *JobRegistryTransactor) ReleaseJobs(opts *bind.TransactOpts, jobIds []*big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "releaseJobs", jobIds)
}

// ReleaseJobs is a paid mutator transaction binding the contract method 0x64dc01d5.
//
// Solidity: function releaseJobs(uint256[] jobIds) returns()
func (_JobRegistry *JobRegistrySession) ReleaseJobs(jobIds []*big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReleaseJobs(&_JobRegistry.TransactOpts, jobIds)
}

// ReleaseJobs is a paid mutator transaction binding the contract method 0x64dc01d5.
//
// Solidity: function releaseJobs(uint256[] jobIds) returns()
func (_JobRegistry *JobRegistryTransactorSession) ReleaseJobs(jobIds []*big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ReleaseJobs(&_JobRegistry.TransactOpts, jobIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobRegistry *JobRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobRegistry *JobRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _JobRegistry.Contract.RenounceOwnership(&_JobRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobRegistry *JobRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _JobRegistry.Contract.RenounceOwnership(&_JobRegistry.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xe29db78e.
//
// Solidity: function resolveDispute(uint256 jobId, bool workerGuilty, bytes32 reExecutionBlobHash, uint256 _similarityScore) returns()
func (_JobRegistry *JobRegistryTransactor) ResolveDispute(opts *bind.TransactOpts, jobId *big.Int, workerGuilty bool, reExecutionBlobHash [32]byte, _similarityScore *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "resolveDispute", jobId, workerGuilty, reExecutionBlobHash, _similarityScore)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xe29db78e.
//
// Solidity: function resolveDispute(uint256 jobId, bool workerGuilty, bytes32 reExecutionBlobHash, uint256 _similarityScore) returns()
func (_JobRegistry *JobRegistrySession) ResolveDispute(jobId *big.Int, workerGuilty bool, reExecutionBlobHash [32]byte, _similarityScore *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ResolveDispute(&_JobRegistry.TransactOpts, jobId, workerGuilty, reExecutionBlobHash, _similarityScore)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xe29db78e.
//
// Solidity: function resolveDispute(uint256 jobId, bool workerGuilty, bytes32 reExecutionBlobHash, uint256 _similarityScore) returns()
func (_JobRegistry *JobRegistryTransactorSession) ResolveDispute(jobId *big.Int, workerGuilty bool, reExecutionBlobHash [32]byte, _similarityScore *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.ResolveDispute(&_JobRegistry.TransactOpts, jobId, workerGuilty, reExecutionBlobHash, _similarityScore)
}

// SetDelegateAllowance is a paid mutator transaction binding the contract method 0xd0ffe1cb.
//
// Solidity: function setDelegateAllowance(address delegate, uint256 allowance) returns()
func (_JobRegistry *JobRegistryTransactor) SetDelegateAllowance(opts *bind.TransactOpts, delegate common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "setDelegateAllowance", delegate, allowance)
}

// SetDelegateAllowance is a paid mutator transaction binding the contract method 0xd0ffe1cb.
//
// Solidity: function setDelegateAllowance(address delegate, uint256 allowance) returns()
func (_JobRegistry *JobRegistrySession) SetDelegateAllowance(delegate common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.SetDelegateAllowance(&_JobRegistry.TransactOpts, delegate, allowance)
}

// SetDelegateAllowance is a paid mutator transaction binding the contract method 0xd0ffe1cb.
//
// Solidity: function setDelegateAllowance(address delegate, uint256 allowance) returns()
func (_JobRegistry *JobRegistryTransactorSession) SetDelegateAllowance(delegate common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.SetDelegateAllowance(&_JobRegistry.TransactOpts, delegate, allowance)
}

// SetDelegateAuthorization is a paid mutator transaction binding the contract method 0xa6e4e43f.
//
// Solidity: function setDelegateAuthorization(address delegate, bool authorized) returns()
func (_JobRegistry *JobRegistryTransactor) SetDelegateAuthorization(opts *bind.TransactOpts, delegate common.Address, authorized bool) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "setDelegateAuthorization", delegate, authorized)
}

// SetDelegateAuthorization is a paid mutator transaction binding the contract method 0xa6e4e43f.
//
// Solidity: function setDelegateAuthorization(address delegate, bool authorized) returns()
func (_JobRegistry *JobRegistrySession) SetDelegateAuthorization(delegate common.Address, authorized bool) (*types.Transaction, error) {
	return _JobRegistry.Contract.SetDelegateAuthorization(&_JobRegistry.TransactOpts, delegate, authorized)
}

// SetDelegateAuthorization is a paid mutator transaction binding the contract method 0xa6e4e43f.
//
// Solidity: function setDelegateAuthorization(address delegate, bool authorized) returns()
func (_JobRegistry *JobRegistryTransactorSession) SetDelegateAuthorization(delegate common.Address, authorized bool) (*types.Transaction, error) {
	return _JobRegistry.Contract.SetDelegateAuthorization(&_JobRegistry.TransactOpts, delegate, authorized)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe3f4f3e9.
//
// Solidity: function submitJob(uint256 sessionId, bytes32 blobHash) payable returns(uint256 jobId)
func (_JobRegistry *JobRegistryTransactor) SubmitJob(opts *bind.TransactOpts, sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "submitJob", sessionId, blobHash)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe3f4f3e9.
//
// Solidity: function submitJob(uint256 sessionId, bytes32 blobHash) payable returns(uint256 jobId)
func (_JobRegistry *JobRegistrySession) SubmitJob(sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.SubmitJob(&_JobRegistry.TransactOpts, sessionId, blobHash)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe3f4f3e9.
//
// Solidity: function submitJob(uint256 sessionId, bytes32 blobHash) payable returns(uint256 jobId)
func (_JobRegistry *JobRegistryTransactorSession) SubmitJob(sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.SubmitJob(&_JobRegistry.TransactOpts, sessionId, blobHash)
}

// SubmitJobOnBehalf is a paid mutator transaction binding the contract method 0xfb811755.
//
// Solidity: function submitJobOnBehalf(address user, uint256 sessionId, bytes32 blobHash) returns(uint256 jobId)
func (_JobRegistry *JobRegistryTransactor) SubmitJobOnBehalf(opts *bind.TransactOpts, user common.Address, sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "submitJobOnBehalf", user, sessionId, blobHash)
}

// SubmitJobOnBehalf is a paid mutator transaction binding the contract method 0xfb811755.
//
// Solidity: function submitJobOnBehalf(address user, uint256 sessionId, bytes32 blobHash) returns(uint256 jobId)
func (_JobRegistry *JobRegistrySession) SubmitJobOnBehalf(user common.Address, sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.SubmitJobOnBehalf(&_JobRegistry.TransactOpts, user, sessionId, blobHash)
}

// SubmitJobOnBehalf is a paid mutator transaction binding the contract method 0xfb811755.
//
// Solidity: function submitJobOnBehalf(address user, uint256 sessionId, bytes32 blobHash) returns(uint256 jobId)
func (_JobRegistry *JobRegistryTransactorSession) SubmitJobOnBehalf(user common.Address, sessionId *big.Int, blobHash [32]byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.SubmitJobOnBehalf(&_JobRegistry.TransactOpts, user, sessionId, blobHash)
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_JobRegistry *JobRegistryTransactor) SunsetGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "sunsetGuardian")
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_JobRegistry *JobRegistrySession) SunsetGuardian() (*types.Transaction, error) {
	return _JobRegistry.Contract.SunsetGuardian(&_JobRegistry.TransactOpts)
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_JobRegistry *JobRegistryTransactorSession) SunsetGuardian() (*types.Transaction, error) {
	return _JobRegistry.Contract.SunsetGuardian(&_JobRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobRegistry *JobRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobRegistry *JobRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.TransferOwnership(&_JobRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobRegistry *JobRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JobRegistry.Contract.TransferOwnership(&_JobRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JobRegistry *JobRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JobRegistry *JobRegistrySession) Unpause() (*types.Transaction, error) {
	return _JobRegistry.Contract.Unpause(&_JobRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_JobRegistry *JobRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _JobRegistry.Contract.Unpause(&_JobRegistry.TransactOpts)
}

// UpdateSessionKey is a paid mutator transaction binding the contract method 0x0938e5ac.
//
// Solidity: function updateSessionKey(uint256 sessionId, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) returns()
func (_JobRegistry *JobRegistryTransactor) UpdateSessionKey(opts *bind.TransactOpts, sessionId *big.Int, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "updateSessionKey", sessionId, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// UpdateSessionKey is a paid mutator transaction binding the contract method 0x0938e5ac.
//
// Solidity: function updateSessionKey(uint256 sessionId, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) returns()
func (_JobRegistry *JobRegistrySession) UpdateSessionKey(sessionId *big.Int, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.UpdateSessionKey(&_JobRegistry.TransactOpts, sessionId, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// UpdateSessionKey is a paid mutator transaction binding the contract method 0x0938e5ac.
//
// Solidity: function updateSessionKey(uint256 sessionId, bytes encWorkerKey, bytes encDisputerKey, bytes dispatcherSignature, uint256 expiry) returns()
func (_JobRegistry *JobRegistryTransactorSession) UpdateSessionKey(sessionId *big.Int, encWorkerKey []byte, encDisputerKey []byte, dispatcherSignature []byte, expiry *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.UpdateSessionKey(&_JobRegistry.TransactOpts, sessionId, encWorkerKey, encDisputerKey, dispatcherSignature, expiry)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_JobRegistry *JobRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_JobRegistry *JobRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.UpgradeToAndCall(&_JobRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_JobRegistry *JobRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _JobRegistry.Contract.UpgradeToAndCall(&_JobRegistry.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_JobRegistry *JobRegistryTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_JobRegistry *JobRegistrySession) Withdraw() (*types.Transaction, error) {
	return _JobRegistry.Contract.Withdraw(&_JobRegistry.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_JobRegistry *JobRegistryTransactorSession) Withdraw() (*types.Transaction, error) {
	return _JobRegistry.Contract.Withdraw(&_JobRegistry.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_JobRegistry *JobRegistryTransactor) WithdrawBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _JobRegistry.contract.Transact(opts, "withdrawBalance", amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_JobRegistry *JobRegistrySession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.WithdrawBalance(&_JobRegistry.TransactOpts, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_JobRegistry *JobRegistryTransactorSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _JobRegistry.Contract.WithdrawBalance(&_JobRegistry.TransactOpts, amount)
}

// JobRegistryDelegateAllowanceSetIterator is returned from FilterDelegateAllowanceSet and is used to iterate over the raw logs and unpacked data for DelegateAllowanceSet events raised by the JobRegistry contract.
type JobRegistryDelegateAllowanceSetIterator struct {
	Event *JobRegistryDelegateAllowanceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryDelegateAllowanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryDelegateAllowanceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryDelegateAllowanceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryDelegateAllowanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryDelegateAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryDelegateAllowanceSet represents a DelegateAllowanceSet event raised by the JobRegistry contract.
type JobRegistryDelegateAllowanceSet struct {
	User      common.Address
	Delegate  common.Address
	Allowance *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateAllowanceSet is a free log retrieval operation binding the contract event 0x0a92c71d37967e95ab74942290d3af32f58045aa7f0533d676a1bac4a9f839dc.
//
// Solidity: event DelegateAllowanceSet(address indexed user, address indexed delegate, uint256 allowance)
func (_JobRegistry *JobRegistryFilterer) FilterDelegateAllowanceSet(opts *bind.FilterOpts, user []common.Address, delegate []common.Address) (*JobRegistryDelegateAllowanceSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "DelegateAllowanceSet", userRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryDelegateAllowanceSetIterator{contract: _JobRegistry.contract, event: "DelegateAllowanceSet", logs: logs, sub: sub}, nil
}

// WatchDelegateAllowanceSet is a free log subscription operation binding the contract event 0x0a92c71d37967e95ab74942290d3af32f58045aa7f0533d676a1bac4a9f839dc.
//
// Solidity: event DelegateAllowanceSet(address indexed user, address indexed delegate, uint256 allowance)
func (_JobRegistry *JobRegistryFilterer) WatchDelegateAllowanceSet(opts *bind.WatchOpts, sink chan<- *JobRegistryDelegateAllowanceSet, user []common.Address, delegate []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "DelegateAllowanceSet", userRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryDelegateAllowanceSet)
				if err := _JobRegistry.contract.UnpackLog(event, "DelegateAllowanceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateAllowanceSet is a log parse operation binding the contract event 0x0a92c71d37967e95ab74942290d3af32f58045aa7f0533d676a1bac4a9f839dc.
//
// Solidity: event DelegateAllowanceSet(address indexed user, address indexed delegate, uint256 allowance)
func (_JobRegistry *JobRegistryFilterer) ParseDelegateAllowanceSet(log types.Log) (*JobRegistryDelegateAllowanceSet, error) {
	event := new(JobRegistryDelegateAllowanceSet)
	if err := _JobRegistry.contract.UnpackLog(event, "DelegateAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryDelegateAuthorizationSetIterator is returned from FilterDelegateAuthorizationSet and is used to iterate over the raw logs and unpacked data for DelegateAuthorizationSet events raised by the JobRegistry contract.
type JobRegistryDelegateAuthorizationSetIterator struct {
	Event *JobRegistryDelegateAuthorizationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryDelegateAuthorizationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryDelegateAuthorizationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryDelegateAuthorizationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryDelegateAuthorizationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryDelegateAuthorizationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryDelegateAuthorizationSet represents a DelegateAuthorizationSet event raised by the JobRegistry contract.
type JobRegistryDelegateAuthorizationSet struct {
	User       common.Address
	Delegate   common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegateAuthorizationSet is a free log retrieval operation binding the contract event 0x9755c2a3048b679844f90739c00e9c3cb2e66b722ce7a0797e4d5d2efb5981d3.
//
// Solidity: event DelegateAuthorizationSet(address indexed user, address indexed delegate, bool authorized)
func (_JobRegistry *JobRegistryFilterer) FilterDelegateAuthorizationSet(opts *bind.FilterOpts, user []common.Address, delegate []common.Address) (*JobRegistryDelegateAuthorizationSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "DelegateAuthorizationSet", userRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryDelegateAuthorizationSetIterator{contract: _JobRegistry.contract, event: "DelegateAuthorizationSet", logs: logs, sub: sub}, nil
}

// WatchDelegateAuthorizationSet is a free log subscription operation binding the contract event 0x9755c2a3048b679844f90739c00e9c3cb2e66b722ce7a0797e4d5d2efb5981d3.
//
// Solidity: event DelegateAuthorizationSet(address indexed user, address indexed delegate, bool authorized)
func (_JobRegistry *JobRegistryFilterer) WatchDelegateAuthorizationSet(opts *bind.WatchOpts, sink chan<- *JobRegistryDelegateAuthorizationSet, user []common.Address, delegate []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "DelegateAuthorizationSet", userRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryDelegateAuthorizationSet)
				if err := _JobRegistry.contract.UnpackLog(event, "DelegateAuthorizationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateAuthorizationSet is a log parse operation binding the contract event 0x9755c2a3048b679844f90739c00e9c3cb2e66b722ce7a0797e4d5d2efb5981d3.
//
// Solidity: event DelegateAuthorizationSet(address indexed user, address indexed delegate, bool authorized)
func (_JobRegistry *JobRegistryFilterer) ParseDelegateAuthorizationSet(log types.Log) (*JobRegistryDelegateAuthorizationSet, error) {
	event := new(JobRegistryDelegateAuthorizationSet)
	if err := _JobRegistry.contract.UnpackLog(event, "DelegateAuthorizationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the JobRegistry contract.
type JobRegistryDepositedIterator struct {
	Event *JobRegistryDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryDeposited represents a Deposited event raised by the JobRegistry contract.
type JobRegistryDeposited struct {
	User       common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*JobRegistryDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryDepositedIterator{contract: _JobRegistry.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *JobRegistryDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryDeposited)
				if err := _JobRegistry.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) ParseDeposited(log types.Log) (*JobRegistryDeposited, error) {
	event := new(JobRegistryDeposited)
	if err := _JobRegistry.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryDisputeCreatedIterator is returned from FilterDisputeCreated and is used to iterate over the raw logs and unpacked data for DisputeCreated events raised by the JobRegistry contract.
type JobRegistryDisputeCreatedIterator struct {
	Event *JobRegistryDisputeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryDisputeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryDisputeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryDisputeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryDisputeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryDisputeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryDisputeCreated represents a DisputeCreated event raised by the JobRegistry contract.
type JobRegistryDisputeCreated struct {
	JobId    *big.Int
	Disputer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisputeCreated is a free log retrieval operation binding the contract event 0x3c9db11b5587ddca19c48971a5bd5cb36500cd7d156815c6728cbc58ff9ac6bf.
//
// Solidity: event DisputeCreated(uint256 indexed jobId, address indexed disputer)
func (_JobRegistry *JobRegistryFilterer) FilterDisputeCreated(opts *bind.FilterOpts, jobId []*big.Int, disputer []common.Address) (*JobRegistryDisputeCreatedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "DisputeCreated", jobIdRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryDisputeCreatedIterator{contract: _JobRegistry.contract, event: "DisputeCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCreated is a free log subscription operation binding the contract event 0x3c9db11b5587ddca19c48971a5bd5cb36500cd7d156815c6728cbc58ff9ac6bf.
//
// Solidity: event DisputeCreated(uint256 indexed jobId, address indexed disputer)
func (_JobRegistry *JobRegistryFilterer) WatchDisputeCreated(opts *bind.WatchOpts, sink chan<- *JobRegistryDisputeCreated, jobId []*big.Int, disputer []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "DisputeCreated", jobIdRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryDisputeCreated)
				if err := _JobRegistry.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCreated is a log parse operation binding the contract event 0x3c9db11b5587ddca19c48971a5bd5cb36500cd7d156815c6728cbc58ff9ac6bf.
//
// Solidity: event DisputeCreated(uint256 indexed jobId, address indexed disputer)
func (_JobRegistry *JobRegistryFilterer) ParseDisputeCreated(log types.Log) (*JobRegistryDisputeCreated, error) {
	event := new(JobRegistryDisputeCreated)
	if err := _JobRegistry.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the JobRegistry contract.
type JobRegistryDisputeResolvedIterator struct {
	Event *JobRegistryDisputeResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryDisputeResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryDisputeResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryDisputeResolved represents a DisputeResolved event raised by the JobRegistry contract.
type JobRegistryDisputeResolved struct {
	JobId           *big.Int
	WorkerGuilty    bool
	SimilarityScore *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x5add0b1fc2a3dfd95be33d1d5feb6709087e50213b7a73c3c3b2c387001221e7.
//
// Solidity: event DisputeResolved(uint256 indexed jobId, bool workerGuilty, uint256 similarityScore)
func (_JobRegistry *JobRegistryFilterer) FilterDisputeResolved(opts *bind.FilterOpts, jobId []*big.Int) (*JobRegistryDisputeResolvedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "DisputeResolved", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryDisputeResolvedIterator{contract: _JobRegistry.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x5add0b1fc2a3dfd95be33d1d5feb6709087e50213b7a73c3c3b2c387001221e7.
//
// Solidity: event DisputeResolved(uint256 indexed jobId, bool workerGuilty, uint256 similarityScore)
func (_JobRegistry *JobRegistryFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *JobRegistryDisputeResolved, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "DisputeResolved", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryDisputeResolved)
				if err := _JobRegistry.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolved is a log parse operation binding the contract event 0x5add0b1fc2a3dfd95be33d1d5feb6709087e50213b7a73c3c3b2c387001221e7.
//
// Solidity: event DisputeResolved(uint256 indexed jobId, bool workerGuilty, uint256 similarityScore)
func (_JobRegistry *JobRegistryFilterer) ParseDisputeResolved(log types.Log) (*JobRegistryDisputeResolved, error) {
	event := new(JobRegistryDisputeResolved)
	if err := _JobRegistry.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryFeeDistributedIterator is returned from FilterFeeDistributed and is used to iterate over the raw logs and unpacked data for FeeDistributed events raised by the JobRegistry contract.
type JobRegistryFeeDistributedIterator struct {
	Event *JobRegistryFeeDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryFeeDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryFeeDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryFeeDistributed represents a FeeDistributed event raised by the JobRegistry contract.
type JobRegistryFeeDistributed struct {
	JobId         *big.Int
	WorkerShare   *big.Int
	ProtocolShare *big.Int
	FeePoolShare  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributed is a free log retrieval operation binding the contract event 0x5b3735d487e9342fbcf11df79c2146f60ffabadb4d2e31ced107fe4346cd21ef.
//
// Solidity: event FeeDistributed(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) FilterFeeDistributed(opts *bind.FilterOpts, jobId []*big.Int) (*JobRegistryFeeDistributedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "FeeDistributed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryFeeDistributedIterator{contract: _JobRegistry.contract, event: "FeeDistributed", logs: logs, sub: sub}, nil
}

// WatchFeeDistributed is a free log subscription operation binding the contract event 0x5b3735d487e9342fbcf11df79c2146f60ffabadb4d2e31ced107fe4346cd21ef.
//
// Solidity: event FeeDistributed(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) WatchFeeDistributed(opts *bind.WatchOpts, sink chan<- *JobRegistryFeeDistributed, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "FeeDistributed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryFeeDistributed)
				if err := _JobRegistry.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeDistributed is a log parse operation binding the contract event 0x5b3735d487e9342fbcf11df79c2146f60ffabadb4d2e31ced107fe4346cd21ef.
//
// Solidity: event FeeDistributed(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) ParseFeeDistributed(log types.Log) (*JobRegistryFeeDistributed, error) {
	event := new(JobRegistryFeeDistributed)
	if err := _JobRegistry.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryGuardianUpdatedIterator is returned from FilterGuardianUpdated and is used to iterate over the raw logs and unpacked data for GuardianUpdated events raised by the JobRegistry contract.
type JobRegistryGuardianUpdatedIterator struct {
	Event *JobRegistryGuardianUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryGuardianUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryGuardianUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryGuardianUpdated represents a GuardianUpdated event raised by the JobRegistry contract.
type JobRegistryGuardianUpdated struct {
	OldGuardian common.Address
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGuardianUpdated is a free log retrieval operation binding the contract event 0x064d28d3d3071c5cbc271a261c10c2f0f0d9e319390397101aa0eb23c6bad909.
//
// Solidity: event GuardianUpdated(address indexed oldGuardian, address indexed newGuardian)
func (_JobRegistry *JobRegistryFilterer) FilterGuardianUpdated(opts *bind.FilterOpts, oldGuardian []common.Address, newGuardian []common.Address) (*JobRegistryGuardianUpdatedIterator, error) {

	var oldGuardianRule []interface{}
	for _, oldGuardianItem := range oldGuardian {
		oldGuardianRule = append(oldGuardianRule, oldGuardianItem)
	}
	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "GuardianUpdated", oldGuardianRule, newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryGuardianUpdatedIterator{contract: _JobRegistry.contract, event: "GuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchGuardianUpdated is a free log subscription operation binding the contract event 0x064d28d3d3071c5cbc271a261c10c2f0f0d9e319390397101aa0eb23c6bad909.
//
// Solidity: event GuardianUpdated(address indexed oldGuardian, address indexed newGuardian)
func (_JobRegistry *JobRegistryFilterer) WatchGuardianUpdated(opts *bind.WatchOpts, sink chan<- *JobRegistryGuardianUpdated, oldGuardian []common.Address, newGuardian []common.Address) (event.Subscription, error) {

	var oldGuardianRule []interface{}
	for _, oldGuardianItem := range oldGuardian {
		oldGuardianRule = append(oldGuardianRule, oldGuardianItem)
	}
	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "GuardianUpdated", oldGuardianRule, newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryGuardianUpdated)
				if err := _JobRegistry.contract.UnpackLog(event, "GuardianUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGuardianUpdated is a log parse operation binding the contract event 0x064d28d3d3071c5cbc271a261c10c2f0f0d9e319390397101aa0eb23c6bad909.
//
// Solidity: event GuardianUpdated(address indexed oldGuardian, address indexed newGuardian)
func (_JobRegistry *JobRegistryFilterer) ParseGuardianUpdated(log types.Log) (*JobRegistryGuardianUpdated, error) {
	event := new(JobRegistryGuardianUpdated)
	if err := _JobRegistry.contract.UnpackLog(event, "GuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the JobRegistry contract.
type JobRegistryInitializedIterator struct {
	Event *JobRegistryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryInitialized represents a Initialized event raised by the JobRegistry contract.
type JobRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JobRegistry *JobRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobRegistryInitializedIterator, error) {

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobRegistryInitializedIterator{contract: _JobRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JobRegistry *JobRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryInitialized)
				if err := _JobRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JobRegistry *JobRegistryFilterer) ParseInitialized(log types.Log) (*JobRegistryInitialized, error) {
	event := new(JobRegistryInitialized)
	if err := _JobRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryJobAcknowledgedIterator is returned from FilterJobAcknowledged and is used to iterate over the raw logs and unpacked data for JobAcknowledged events raised by the JobRegistry contract.
type JobRegistryJobAcknowledgedIterator struct {
	Event *JobRegistryJobAcknowledged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryJobAcknowledgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryJobAcknowledged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryJobAcknowledged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryJobAcknowledgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryJobAcknowledgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryJobAcknowledged represents a JobAcknowledged event raised by the JobRegistry contract.
type JobRegistryJobAcknowledged struct {
	JobId  *big.Int
	Worker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobAcknowledged is a free log retrieval operation binding the contract event 0xda96050a785b20f9636c63908189212b00fbf5e2f24d7a9d8a707deb1ce84bce.
//
// Solidity: event JobAcknowledged(uint256 indexed jobId, address indexed worker)
func (_JobRegistry *JobRegistryFilterer) FilterJobAcknowledged(opts *bind.FilterOpts, jobId []*big.Int, worker []common.Address) (*JobRegistryJobAcknowledgedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "JobAcknowledged", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryJobAcknowledgedIterator{contract: _JobRegistry.contract, event: "JobAcknowledged", logs: logs, sub: sub}, nil
}

// WatchJobAcknowledged is a free log subscription operation binding the contract event 0xda96050a785b20f9636c63908189212b00fbf5e2f24d7a9d8a707deb1ce84bce.
//
// Solidity: event JobAcknowledged(uint256 indexed jobId, address indexed worker)
func (_JobRegistry *JobRegistryFilterer) WatchJobAcknowledged(opts *bind.WatchOpts, sink chan<- *JobRegistryJobAcknowledged, jobId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "JobAcknowledged", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryJobAcknowledged)
				if err := _JobRegistry.contract.UnpackLog(event, "JobAcknowledged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobAcknowledged is a log parse operation binding the contract event 0xda96050a785b20f9636c63908189212b00fbf5e2f24d7a9d8a707deb1ce84bce.
//
// Solidity: event JobAcknowledged(uint256 indexed jobId, address indexed worker)
func (_JobRegistry *JobRegistryFilterer) ParseJobAcknowledged(log types.Log) (*JobRegistryJobAcknowledged, error) {
	event := new(JobRegistryJobAcknowledged)
	if err := _JobRegistry.contract.UnpackLog(event, "JobAcknowledged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryJobCompletedIterator is returned from FilterJobCompleted and is used to iterate over the raw logs and unpacked data for JobCompleted events raised by the JobRegistry contract.
type JobRegistryJobCompletedIterator struct {
	Event *JobRegistryJobCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryJobCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryJobCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryJobCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryJobCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryJobCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryJobCompleted represents a JobCompleted event raised by the JobRegistry contract.
type JobRegistryJobCompleted struct {
	JobId                  *big.Int
	Worker                 common.Address
	ResponseBlobHash       [32]byte
	ResponseCiphertextHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterJobCompleted is a free log retrieval operation binding the contract event 0xdb545db74bae046337ed01971cf61569fd1a1460ff8ed511ab19ceaac1326377.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed worker, bytes32 responseBlobHash, bytes32 responseCiphertextHash)
func (_JobRegistry *JobRegistryFilterer) FilterJobCompleted(opts *bind.FilterOpts, jobId []*big.Int, worker []common.Address) (*JobRegistryJobCompletedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "JobCompleted", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryJobCompletedIterator{contract: _JobRegistry.contract, event: "JobCompleted", logs: logs, sub: sub}, nil
}

// WatchJobCompleted is a free log subscription operation binding the contract event 0xdb545db74bae046337ed01971cf61569fd1a1460ff8ed511ab19ceaac1326377.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed worker, bytes32 responseBlobHash, bytes32 responseCiphertextHash)
func (_JobRegistry *JobRegistryFilterer) WatchJobCompleted(opts *bind.WatchOpts, sink chan<- *JobRegistryJobCompleted, jobId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "JobCompleted", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryJobCompleted)
				if err := _JobRegistry.contract.UnpackLog(event, "JobCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobCompleted is a log parse operation binding the contract event 0xdb545db74bae046337ed01971cf61569fd1a1460ff8ed511ab19ceaac1326377.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed worker, bytes32 responseBlobHash, bytes32 responseCiphertextHash)
func (_JobRegistry *JobRegistryFilterer) ParseJobCompleted(log types.Log) (*JobRegistryJobCompleted, error) {
	event := new(JobRegistryJobCompleted)
	if err := _JobRegistry.contract.UnpackLog(event, "JobCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryJobReleasedIterator is returned from FilterJobReleased and is used to iterate over the raw logs and unpacked data for JobReleased events raised by the JobRegistry contract.
type JobRegistryJobReleasedIterator struct {
	Event *JobRegistryJobReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryJobReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryJobReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryJobReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryJobReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryJobReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryJobReleased represents a JobReleased event raised by the JobRegistry contract.
type JobRegistryJobReleased struct {
	JobId         *big.Int
	WorkerShare   *big.Int
	ProtocolShare *big.Int
	FeePoolShare  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterJobReleased is a free log retrieval operation binding the contract event 0xa3a19f49236690faa5a4bd840e6640d2c92c54214a2b7e304267e3005a9a9cc0.
//
// Solidity: event JobReleased(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) FilterJobReleased(opts *bind.FilterOpts, jobId []*big.Int) (*JobRegistryJobReleasedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "JobReleased", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryJobReleasedIterator{contract: _JobRegistry.contract, event: "JobReleased", logs: logs, sub: sub}, nil
}

// WatchJobReleased is a free log subscription operation binding the contract event 0xa3a19f49236690faa5a4bd840e6640d2c92c54214a2b7e304267e3005a9a9cc0.
//
// Solidity: event JobReleased(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) WatchJobReleased(opts *bind.WatchOpts, sink chan<- *JobRegistryJobReleased, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "JobReleased", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryJobReleased)
				if err := _JobRegistry.contract.UnpackLog(event, "JobReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobReleased is a log parse operation binding the contract event 0xa3a19f49236690faa5a4bd840e6640d2c92c54214a2b7e304267e3005a9a9cc0.
//
// Solidity: event JobReleased(uint256 indexed jobId, uint256 workerShare, uint256 protocolShare, uint256 feePoolShare)
func (_JobRegistry *JobRegistryFilterer) ParseJobReleased(log types.Log) (*JobRegistryJobReleased, error) {
	event := new(JobRegistryJobReleased)
	if err := _JobRegistry.contract.UnpackLog(event, "JobReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryJobSubmittedIterator is returned from FilterJobSubmitted and is used to iterate over the raw logs and unpacked data for JobSubmitted events raised by the JobRegistry contract.
type JobRegistryJobSubmittedIterator struct {
	Event *JobRegistryJobSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryJobSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryJobSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryJobSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryJobSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryJobSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryJobSubmitted represents a JobSubmitted event raised by the JobRegistry contract.
type JobRegistryJobSubmitted struct {
	JobId     *big.Int
	SessionId *big.Int
	Worker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobSubmitted is a free log retrieval operation binding the contract event 0xfb47370368875d7490803c5653d9496d0a3c5e1b49a17f013ec37abd9d86d356.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed sessionId, address worker)
func (_JobRegistry *JobRegistryFilterer) FilterJobSubmitted(opts *bind.FilterOpts, jobId []*big.Int, sessionId []*big.Int) (*JobRegistryJobSubmittedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "JobSubmitted", jobIdRule, sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryJobSubmittedIterator{contract: _JobRegistry.contract, event: "JobSubmitted", logs: logs, sub: sub}, nil
}

// WatchJobSubmitted is a free log subscription operation binding the contract event 0xfb47370368875d7490803c5653d9496d0a3c5e1b49a17f013ec37abd9d86d356.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed sessionId, address worker)
func (_JobRegistry *JobRegistryFilterer) WatchJobSubmitted(opts *bind.WatchOpts, sink chan<- *JobRegistryJobSubmitted, jobId []*big.Int, sessionId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "JobSubmitted", jobIdRule, sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryJobSubmitted)
				if err := _JobRegistry.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobSubmitted is a log parse operation binding the contract event 0xfb47370368875d7490803c5653d9496d0a3c5e1b49a17f013ec37abd9d86d356.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed sessionId, address worker)
func (_JobRegistry *JobRegistryFilterer) ParseJobSubmitted(log types.Log) (*JobRegistryJobSubmitted, error) {
	event := new(JobRegistryJobSubmitted)
	if err := _JobRegistry.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryJobTimedOutIterator is returned from FilterJobTimedOut and is used to iterate over the raw logs and unpacked data for JobTimedOut events raised by the JobRegistry contract.
type JobRegistryJobTimedOutIterator struct {
	Event *JobRegistryJobTimedOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryJobTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryJobTimedOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryJobTimedOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryJobTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryJobTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryJobTimedOut represents a JobTimedOut event raised by the JobRegistry contract.
type JobRegistryJobTimedOut struct {
	JobId       *big.Int
	Worker      common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJobTimedOut is a free log retrieval operation binding the contract event 0x6c824cc82239382c89711e47345227ceb53137c068153f1392e553362dcd1720.
//
// Solidity: event JobTimedOut(uint256 indexed jobId, address indexed worker, uint256 slashAmount)
func (_JobRegistry *JobRegistryFilterer) FilterJobTimedOut(opts *bind.FilterOpts, jobId []*big.Int, worker []common.Address) (*JobRegistryJobTimedOutIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "JobTimedOut", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryJobTimedOutIterator{contract: _JobRegistry.contract, event: "JobTimedOut", logs: logs, sub: sub}, nil
}

// WatchJobTimedOut is a free log subscription operation binding the contract event 0x6c824cc82239382c89711e47345227ceb53137c068153f1392e553362dcd1720.
//
// Solidity: event JobTimedOut(uint256 indexed jobId, address indexed worker, uint256 slashAmount)
func (_JobRegistry *JobRegistryFilterer) WatchJobTimedOut(opts *bind.WatchOpts, sink chan<- *JobRegistryJobTimedOut, jobId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "JobTimedOut", jobIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryJobTimedOut)
				if err := _JobRegistry.contract.UnpackLog(event, "JobTimedOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobTimedOut is a log parse operation binding the contract event 0x6c824cc82239382c89711e47345227ceb53137c068153f1392e553362dcd1720.
//
// Solidity: event JobTimedOut(uint256 indexed jobId, address indexed worker, uint256 slashAmount)
func (_JobRegistry *JobRegistryFilterer) ParseJobTimedOut(log types.Log) (*JobRegistryJobTimedOut, error) {
	event := new(JobRegistryJobTimedOut)
	if err := _JobRegistry.contract.UnpackLog(event, "JobTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the JobRegistry contract.
type JobRegistryOwnershipTransferredIterator struct {
	Event *JobRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the JobRegistry contract.
type JobRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JobRegistry *JobRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryOwnershipTransferredIterator{contract: _JobRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JobRegistry *JobRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryOwnershipTransferred)
				if err := _JobRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JobRegistry *JobRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*JobRegistryOwnershipTransferred, error) {
	event := new(JobRegistryOwnershipTransferred)
	if err := _JobRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the JobRegistry contract.
type JobRegistryPausedIterator struct {
	Event *JobRegistryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryPaused represents a Paused event raised by the JobRegistry contract.
type JobRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JobRegistry *JobRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*JobRegistryPausedIterator, error) {

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &JobRegistryPausedIterator{contract: _JobRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JobRegistry *JobRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *JobRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryPaused)
				if err := _JobRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JobRegistry *JobRegistryFilterer) ParsePaused(log types.Log) (*JobRegistryPaused, error) {
	event := new(JobRegistryPaused)
	if err := _JobRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryRefundClaimedIterator is returned from FilterRefundClaimed and is used to iterate over the raw logs and unpacked data for RefundClaimed events raised by the JobRegistry contract.
type JobRegistryRefundClaimedIterator struct {
	Event *JobRegistryRefundClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryRefundClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryRefundClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryRefundClaimed represents a RefundClaimed event raised by the JobRegistry contract.
type JobRegistryRefundClaimed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundClaimed is a free log retrieval operation binding the contract event 0x358fe4192934d3bf28ae181feda1f4bd08ca67f5e2fad55582cce5eb67304ae9.
//
// Solidity: event RefundClaimed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) FilterRefundClaimed(opts *bind.FilterOpts, recipient []common.Address) (*JobRegistryRefundClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "RefundClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryRefundClaimedIterator{contract: _JobRegistry.contract, event: "RefundClaimed", logs: logs, sub: sub}, nil
}

// WatchRefundClaimed is a free log subscription operation binding the contract event 0x358fe4192934d3bf28ae181feda1f4bd08ca67f5e2fad55582cce5eb67304ae9.
//
// Solidity: event RefundClaimed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) WatchRefundClaimed(opts *bind.WatchOpts, sink chan<- *JobRegistryRefundClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "RefundClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryRefundClaimed)
				if err := _JobRegistry.contract.UnpackLog(event, "RefundClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundClaimed is a log parse operation binding the contract event 0x358fe4192934d3bf28ae181feda1f4bd08ca67f5e2fad55582cce5eb67304ae9.
//
// Solidity: event RefundClaimed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) ParseRefundClaimed(log types.Log) (*JobRegistryRefundClaimed, error) {
	event := new(JobRegistryRefundClaimed)
	if err := _JobRegistry.contract.UnpackLog(event, "RefundClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryRefundCreditedToBalanceIterator is returned from FilterRefundCreditedToBalance and is used to iterate over the raw logs and unpacked data for RefundCreditedToBalance events raised by the JobRegistry contract.
type JobRegistryRefundCreditedToBalanceIterator struct {
	Event *JobRegistryRefundCreditedToBalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryRefundCreditedToBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryRefundCreditedToBalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryRefundCreditedToBalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryRefundCreditedToBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryRefundCreditedToBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryRefundCreditedToBalance represents a RefundCreditedToBalance event raised by the JobRegistry contract.
type JobRegistryRefundCreditedToBalance struct {
	User   common.Address
	JobId  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundCreditedToBalance is a free log retrieval operation binding the contract event 0x05fb0485772e72ba2f031e5e64c7502f931bb9579a4093d7734d85c5f0b05617.
//
// Solidity: event RefundCreditedToBalance(address indexed user, uint256 indexed jobId, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) FilterRefundCreditedToBalance(opts *bind.FilterOpts, user []common.Address, jobId []*big.Int) (*JobRegistryRefundCreditedToBalanceIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "RefundCreditedToBalance", userRule, jobIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryRefundCreditedToBalanceIterator{contract: _JobRegistry.contract, event: "RefundCreditedToBalance", logs: logs, sub: sub}, nil
}

// WatchRefundCreditedToBalance is a free log subscription operation binding the contract event 0x05fb0485772e72ba2f031e5e64c7502f931bb9579a4093d7734d85c5f0b05617.
//
// Solidity: event RefundCreditedToBalance(address indexed user, uint256 indexed jobId, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) WatchRefundCreditedToBalance(opts *bind.WatchOpts, sink chan<- *JobRegistryRefundCreditedToBalance, user []common.Address, jobId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "RefundCreditedToBalance", userRule, jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryRefundCreditedToBalance)
				if err := _JobRegistry.contract.UnpackLog(event, "RefundCreditedToBalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundCreditedToBalance is a log parse operation binding the contract event 0x05fb0485772e72ba2f031e5e64c7502f931bb9579a4093d7734d85c5f0b05617.
//
// Solidity: event RefundCreditedToBalance(address indexed user, uint256 indexed jobId, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) ParseRefundCreditedToBalance(log types.Log) (*JobRegistryRefundCreditedToBalance, error) {
	event := new(JobRegistryRefundCreditedToBalance)
	if err := _JobRegistry.contract.UnpackLog(event, "RefundCreditedToBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryRefundEscrowedIterator is returned from FilterRefundEscrowed and is used to iterate over the raw logs and unpacked data for RefundEscrowed events raised by the JobRegistry contract.
type JobRegistryRefundEscrowedIterator struct {
	Event *JobRegistryRefundEscrowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryRefundEscrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryRefundEscrowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryRefundEscrowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryRefundEscrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryRefundEscrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryRefundEscrowed represents a RefundEscrowed event raised by the JobRegistry contract.
type JobRegistryRefundEscrowed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundEscrowed is a free log retrieval operation binding the contract event 0xe5f59195cd6bc7dcd4258d037bb2e7138cda6d6752f09a26c1166a2d21897181.
//
// Solidity: event RefundEscrowed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) FilterRefundEscrowed(opts *bind.FilterOpts, recipient []common.Address) (*JobRegistryRefundEscrowedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "RefundEscrowed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryRefundEscrowedIterator{contract: _JobRegistry.contract, event: "RefundEscrowed", logs: logs, sub: sub}, nil
}

// WatchRefundEscrowed is a free log subscription operation binding the contract event 0xe5f59195cd6bc7dcd4258d037bb2e7138cda6d6752f09a26c1166a2d21897181.
//
// Solidity: event RefundEscrowed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) WatchRefundEscrowed(opts *bind.WatchOpts, sink chan<- *JobRegistryRefundEscrowed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "RefundEscrowed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryRefundEscrowed)
				if err := _JobRegistry.contract.UnpackLog(event, "RefundEscrowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundEscrowed is a log parse operation binding the contract event 0xe5f59195cd6bc7dcd4258d037bb2e7138cda6d6752f09a26c1166a2d21897181.
//
// Solidity: event RefundEscrowed(address indexed recipient, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) ParseRefundEscrowed(log types.Log) (*JobRegistryRefundEscrowed, error) {
	event := new(JobRegistryRefundEscrowed)
	if err := _JobRegistry.contract.UnpackLog(event, "RefundEscrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistrySessionClosedIterator is returned from FilterSessionClosed and is used to iterate over the raw logs and unpacked data for SessionClosed events raised by the JobRegistry contract.
type JobRegistrySessionClosedIterator struct {
	Event *JobRegistrySessionClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistrySessionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistrySessionClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistrySessionClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistrySessionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistrySessionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistrySessionClosed represents a SessionClosed event raised by the JobRegistry contract.
type JobRegistrySessionClosed struct {
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSessionClosed is a free log retrieval operation binding the contract event 0x737ea773789756155ce4ec81b73f88f21c918501c8bbb06721a6c6c46beb11c2.
//
// Solidity: event SessionClosed(uint256 indexed sessionId)
func (_JobRegistry *JobRegistryFilterer) FilterSessionClosed(opts *bind.FilterOpts, sessionId []*big.Int) (*JobRegistrySessionClosedIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "SessionClosed", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistrySessionClosedIterator{contract: _JobRegistry.contract, event: "SessionClosed", logs: logs, sub: sub}, nil
}

// WatchSessionClosed is a free log subscription operation binding the contract event 0x737ea773789756155ce4ec81b73f88f21c918501c8bbb06721a6c6c46beb11c2.
//
// Solidity: event SessionClosed(uint256 indexed sessionId)
func (_JobRegistry *JobRegistryFilterer) WatchSessionClosed(opts *bind.WatchOpts, sink chan<- *JobRegistrySessionClosed, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "SessionClosed", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistrySessionClosed)
				if err := _JobRegistry.contract.UnpackLog(event, "SessionClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionClosed is a log parse operation binding the contract event 0x737ea773789756155ce4ec81b73f88f21c918501c8bbb06721a6c6c46beb11c2.
//
// Solidity: event SessionClosed(uint256 indexed sessionId)
func (_JobRegistry *JobRegistryFilterer) ParseSessionClosed(log types.Log) (*JobRegistrySessionClosed, error) {
	event := new(JobRegistrySessionClosed)
	if err := _JobRegistry.contract.UnpackLog(event, "SessionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistrySessionCreatedIterator is returned from FilterSessionCreated and is used to iterate over the raw logs and unpacked data for SessionCreated events raised by the JobRegistry contract.
type JobRegistrySessionCreatedIterator struct {
	Event *JobRegistrySessionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistrySessionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistrySessionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistrySessionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistrySessionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistrySessionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistrySessionCreated represents a SessionCreated event raised by the JobRegistry contract.
type JobRegistrySessionCreated struct {
	SessionId      *big.Int
	User           common.Address
	ModelId        [32]byte
	Worker         common.Address
	EncWorkerKey   []byte
	EncDisputerKey []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSessionCreated is a free log retrieval operation binding the contract event 0xedf9fab204f0bb366f5b33ff07f441f4e387a833e86bfe1364a42ae2c7e05d73.
//
// Solidity: event SessionCreated(uint256 indexed sessionId, address indexed user, bytes32 indexed modelId, address worker, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) FilterSessionCreated(opts *bind.FilterOpts, sessionId []*big.Int, user []common.Address, modelId [][32]byte) (*JobRegistrySessionCreatedIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "SessionCreated", sessionIdRule, userRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistrySessionCreatedIterator{contract: _JobRegistry.contract, event: "SessionCreated", logs: logs, sub: sub}, nil
}

// WatchSessionCreated is a free log subscription operation binding the contract event 0xedf9fab204f0bb366f5b33ff07f441f4e387a833e86bfe1364a42ae2c7e05d73.
//
// Solidity: event SessionCreated(uint256 indexed sessionId, address indexed user, bytes32 indexed modelId, address worker, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) WatchSessionCreated(opts *bind.WatchOpts, sink chan<- *JobRegistrySessionCreated, sessionId []*big.Int, user []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "SessionCreated", sessionIdRule, userRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistrySessionCreated)
				if err := _JobRegistry.contract.UnpackLog(event, "SessionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionCreated is a log parse operation binding the contract event 0xedf9fab204f0bb366f5b33ff07f441f4e387a833e86bfe1364a42ae2c7e05d73.
//
// Solidity: event SessionCreated(uint256 indexed sessionId, address indexed user, bytes32 indexed modelId, address worker, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) ParseSessionCreated(log types.Log) (*JobRegistrySessionCreated, error) {
	event := new(JobRegistrySessionCreated)
	if err := _JobRegistry.contract.UnpackLog(event, "SessionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistrySessionKeyUpdatedIterator is returned from FilterSessionKeyUpdated and is used to iterate over the raw logs and unpacked data for SessionKeyUpdated events raised by the JobRegistry contract.
type JobRegistrySessionKeyUpdatedIterator struct {
	Event *JobRegistrySessionKeyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistrySessionKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistrySessionKeyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistrySessionKeyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistrySessionKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistrySessionKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistrySessionKeyUpdated represents a SessionKeyUpdated event raised by the JobRegistry contract.
type JobRegistrySessionKeyUpdated struct {
	SessionId      *big.Int
	EncWorkerKey   []byte
	EncDisputerKey []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyUpdated is a free log retrieval operation binding the contract event 0x563032dd5900bf903bcd1eee162d0cb3ef4182646d0ee3f1cfff45a7f5713e5b.
//
// Solidity: event SessionKeyUpdated(uint256 indexed sessionId, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) FilterSessionKeyUpdated(opts *bind.FilterOpts, sessionId []*big.Int) (*JobRegistrySessionKeyUpdatedIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "SessionKeyUpdated", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistrySessionKeyUpdatedIterator{contract: _JobRegistry.contract, event: "SessionKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchSessionKeyUpdated is a free log subscription operation binding the contract event 0x563032dd5900bf903bcd1eee162d0cb3ef4182646d0ee3f1cfff45a7f5713e5b.
//
// Solidity: event SessionKeyUpdated(uint256 indexed sessionId, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) WatchSessionKeyUpdated(opts *bind.WatchOpts, sink chan<- *JobRegistrySessionKeyUpdated, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "SessionKeyUpdated", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistrySessionKeyUpdated)
				if err := _JobRegistry.contract.UnpackLog(event, "SessionKeyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionKeyUpdated is a log parse operation binding the contract event 0x563032dd5900bf903bcd1eee162d0cb3ef4182646d0ee3f1cfff45a7f5713e5b.
//
// Solidity: event SessionKeyUpdated(uint256 indexed sessionId, bytes encWorkerKey, bytes encDisputerKey)
func (_JobRegistry *JobRegistryFilterer) ParseSessionKeyUpdated(log types.Log) (*JobRegistrySessionKeyUpdated, error) {
	event := new(JobRegistrySessionKeyUpdated)
	if err := _JobRegistry.contract.UnpackLog(event, "SessionKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistrySessionReassignedIterator is returned from FilterSessionReassigned and is used to iterate over the raw logs and unpacked data for SessionReassigned events raised by the JobRegistry contract.
type JobRegistrySessionReassignedIterator struct {
	Event *JobRegistrySessionReassigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistrySessionReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistrySessionReassigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistrySessionReassigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistrySessionReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistrySessionReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistrySessionReassigned represents a SessionReassigned event raised by the JobRegistry contract.
type JobRegistrySessionReassigned struct {
	SessionId *big.Int
	NewWorker common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSessionReassigned is a free log retrieval operation binding the contract event 0x11fe9d1166fa89599df5ec1b45c61dbc4aa119faaa664b11d57f38d9931072ab.
//
// Solidity: event SessionReassigned(uint256 indexed sessionId, address newWorker)
func (_JobRegistry *JobRegistryFilterer) FilterSessionReassigned(opts *bind.FilterOpts, sessionId []*big.Int) (*JobRegistrySessionReassignedIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "SessionReassigned", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistrySessionReassignedIterator{contract: _JobRegistry.contract, event: "SessionReassigned", logs: logs, sub: sub}, nil
}

// WatchSessionReassigned is a free log subscription operation binding the contract event 0x11fe9d1166fa89599df5ec1b45c61dbc4aa119faaa664b11d57f38d9931072ab.
//
// Solidity: event SessionReassigned(uint256 indexed sessionId, address newWorker)
func (_JobRegistry *JobRegistryFilterer) WatchSessionReassigned(opts *bind.WatchOpts, sink chan<- *JobRegistrySessionReassigned, sessionId []*big.Int) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "SessionReassigned", sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistrySessionReassigned)
				if err := _JobRegistry.contract.UnpackLog(event, "SessionReassigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionReassigned is a log parse operation binding the contract event 0x11fe9d1166fa89599df5ec1b45c61dbc4aa119faaa664b11d57f38d9931072ab.
//
// Solidity: event SessionReassigned(uint256 indexed sessionId, address newWorker)
func (_JobRegistry *JobRegistryFilterer) ParseSessionReassigned(log types.Log) (*JobRegistrySessionReassigned, error) {
	event := new(JobRegistrySessionReassigned)
	if err := _JobRegistry.contract.UnpackLog(event, "SessionReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryTransferFailedIterator is returned from FilterTransferFailed and is used to iterate over the raw logs and unpacked data for TransferFailed events raised by the JobRegistry contract.
type JobRegistryTransferFailedIterator struct {
	Event *JobRegistryTransferFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryTransferFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryTransferFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryTransferFailed represents a TransferFailed event raised by the JobRegistry contract.
type JobRegistryTransferFailed struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferFailed is a free log retrieval operation binding the contract event 0x1c43b9761b3fba5321ca8212bfc231945f668ccc0c446f333999eea9ce8fda81.
//
// Solidity: event TransferFailed(address indexed to, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) FilterTransferFailed(opts *bind.FilterOpts, to []common.Address) (*JobRegistryTransferFailedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "TransferFailed", toRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryTransferFailedIterator{contract: _JobRegistry.contract, event: "TransferFailed", logs: logs, sub: sub}, nil
}

// WatchTransferFailed is a free log subscription operation binding the contract event 0x1c43b9761b3fba5321ca8212bfc231945f668ccc0c446f333999eea9ce8fda81.
//
// Solidity: event TransferFailed(address indexed to, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) WatchTransferFailed(opts *bind.WatchOpts, sink chan<- *JobRegistryTransferFailed, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "TransferFailed", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryTransferFailed)
				if err := _JobRegistry.contract.UnpackLog(event, "TransferFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFailed is a log parse operation binding the contract event 0x1c43b9761b3fba5321ca8212bfc231945f668ccc0c446f333999eea9ce8fda81.
//
// Solidity: event TransferFailed(address indexed to, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) ParseTransferFailed(log types.Log) (*JobRegistryTransferFailed, error) {
	event := new(JobRegistryTransferFailed)
	if err := _JobRegistry.contract.UnpackLog(event, "TransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the JobRegistry contract.
type JobRegistryUnpausedIterator struct {
	Event *JobRegistryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryUnpaused represents a Unpaused event raised by the JobRegistry contract.
type JobRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JobRegistry *JobRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*JobRegistryUnpausedIterator, error) {

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &JobRegistryUnpausedIterator{contract: _JobRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JobRegistry *JobRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *JobRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryUnpaused)
				if err := _JobRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JobRegistry *JobRegistryFilterer) ParseUnpaused(log types.Log) (*JobRegistryUnpaused, error) {
	event := new(JobRegistryUnpaused)
	if err := _JobRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the JobRegistry contract.
type JobRegistryUpgradedIterator struct {
	Event *JobRegistryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryUpgraded represents a Upgraded event raised by the JobRegistry contract.
type JobRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_JobRegistry *JobRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*JobRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryUpgradedIterator{contract: _JobRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_JobRegistry *JobRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *JobRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryUpgraded)
				if err := _JobRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_JobRegistry *JobRegistryFilterer) ParseUpgraded(log types.Log) (*JobRegistryUpgraded, error) {
	event := new(JobRegistryUpgraded)
	if err := _JobRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the JobRegistry contract.
type JobRegistryWithdrewIterator struct {
	Event *JobRegistryWithdrew // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryWithdrew)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryWithdrew)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryWithdrew represents a Withdrew event raised by the JobRegistry contract.
type JobRegistryWithdrew struct {
	User       common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0xadec52fcd1408589179b85e44b434374db078b4eaf793e7d1a1bb0ae4ecfeee5.
//
// Solidity: event Withdrew(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) FilterWithdrew(opts *bind.FilterOpts, user []common.Address) (*JobRegistryWithdrewIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "Withdrew", userRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryWithdrewIterator{contract: _JobRegistry.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0xadec52fcd1408589179b85e44b434374db078b4eaf793e7d1a1bb0ae4ecfeee5.
//
// Solidity: event Withdrew(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *JobRegistryWithdrew, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "Withdrew", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryWithdrew)
				if err := _JobRegistry.contract.UnpackLog(event, "Withdrew", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrew is a log parse operation binding the contract event 0xadec52fcd1408589179b85e44b434374db078b4eaf793e7d1a1bb0ae4ecfeee5.
//
// Solidity: event Withdrew(address indexed user, uint256 amount, uint256 newBalance)
func (_JobRegistry *JobRegistryFilterer) ParseWithdrew(log types.Log) (*JobRegistryWithdrew, error) {
	event := new(JobRegistryWithdrew)
	if err := _JobRegistry.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobRegistryWorkerWithdrawalIterator is returned from FilterWorkerWithdrawal and is used to iterate over the raw logs and unpacked data for WorkerWithdrawal events raised by the JobRegistry contract.
type JobRegistryWorkerWithdrawalIterator struct {
	Event *JobRegistryWorkerWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobRegistryWorkerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobRegistryWorkerWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobRegistryWorkerWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobRegistryWorkerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobRegistryWorkerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobRegistryWorkerWithdrawal represents a WorkerWithdrawal event raised by the JobRegistry contract.
type JobRegistryWorkerWithdrawal struct {
	Worker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerWithdrawal is a free log retrieval operation binding the contract event 0x4a45698c7a2010a925f4f516fc9c004979abcf0a044581b2d17bee34c858a90c.
//
// Solidity: event WorkerWithdrawal(address indexed worker, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) FilterWorkerWithdrawal(opts *bind.FilterOpts, worker []common.Address) (*JobRegistryWorkerWithdrawalIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.FilterLogs(opts, "WorkerWithdrawal", workerRule)
	if err != nil {
		return nil, err
	}
	return &JobRegistryWorkerWithdrawalIterator{contract: _JobRegistry.contract, event: "WorkerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchWorkerWithdrawal is a free log subscription operation binding the contract event 0x4a45698c7a2010a925f4f516fc9c004979abcf0a044581b2d17bee34c858a90c.
//
// Solidity: event WorkerWithdrawal(address indexed worker, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) WatchWorkerWithdrawal(opts *bind.WatchOpts, sink chan<- *JobRegistryWorkerWithdrawal, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _JobRegistry.contract.WatchLogs(opts, "WorkerWithdrawal", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobRegistryWorkerWithdrawal)
				if err := _JobRegistry.contract.UnpackLog(event, "WorkerWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWorkerWithdrawal is a log parse operation binding the contract event 0x4a45698c7a2010a925f4f516fc9c004979abcf0a044581b2d17bee34c858a90c.
//
// Solidity: event WorkerWithdrawal(address indexed worker, uint256 amount)
func (_JobRegistry *JobRegistryFilterer) ParseWorkerWithdrawal(log types.Log) (*JobRegistryWorkerWithdrawal, error) {
	event := new(JobRegistryWorkerWithdrawal)
	if err := _JobRegistry.contract.UnpackLog(event, "WorkerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
