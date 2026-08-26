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

// WorkerRegistryMetaData contains all meta data concerning the WorkerRegistry contract.
var WorkerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSupportedModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addWhitelistedModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aiConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decrementActiveJobs\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterWorker\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveJobCount\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityMask\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityName\",\"inputs\":[{\"name\":\"bit\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEligibleWorkers\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOffenseCount\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedFunds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSuspendedUntil\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStaked\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTreasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkerCapabilities\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkerEncryptionKey\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkerStake\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incrementActiveJobs\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aiConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_jobRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_guardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isEligible\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isModelWhitelisted\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWorkerRegistered\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWorkerSuspended\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jobRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerCapability\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"bit\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWorker\",\"inputs\":[{\"name\":\"encryptionPubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"reinstate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSupportedModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeWhitelistedModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetOffenses\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"selectEligibleWorker\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"excluded\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCapabilities\",\"inputs\":[{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGuardian\",\"inputs\":[{\"name\":\"newGuardian\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashWithBounty\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bountyRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"slashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bountyPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashWithBountyBps\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bountyBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bountyRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"slashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bountyPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sunsetGuardian\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topUpStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawSlashedFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"workerSupportsModel\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CapabilityRegistered\",\"inputs\":[{\"name\":\"bit\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GuardianUpdated\",\"inputs\":[{\"name\":\"oldGuardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGuardian\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelAdded\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelDelisted\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelRemoved\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelWhitelisted\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OffensesReset\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashedFundsWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeTopUp\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerCapabilitiesSet\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerDeactivated\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerDeregistered\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerRegistered\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"encryptionPubKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerReinstated\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerSlashed\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerSuspended\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"until\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// WorkerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerRegistryMetaData.ABI instead.
var WorkerRegistryABI = WorkerRegistryMetaData.ABI

// WorkerRegistry is an auto generated Go binding around an Ethereum contract.
type WorkerRegistry struct {
	WorkerRegistryCaller     // Read-only binding to the contract
	WorkerRegistryTransactor // Write-only binding to the contract
	WorkerRegistryFilterer   // Log filterer for contract events
}

// WorkerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerRegistrySession struct {
	Contract     *WorkerRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerRegistryCallerSession struct {
	Contract *WorkerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WorkerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerRegistryTransactorSession struct {
	Contract     *WorkerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WorkerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerRegistryRaw struct {
	Contract *WorkerRegistry // Generic contract binding to access the raw methods on
}

// WorkerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerRegistryCallerRaw struct {
	Contract *WorkerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerRegistryTransactorRaw struct {
	Contract *WorkerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerRegistry creates a new instance of WorkerRegistry, bound to a specific deployed contract.
func NewWorkerRegistry(address common.Address, backend bind.ContractBackend) (*WorkerRegistry, error) {
	contract, err := bindWorkerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistry{WorkerRegistryCaller: WorkerRegistryCaller{contract: contract}, WorkerRegistryTransactor: WorkerRegistryTransactor{contract: contract}, WorkerRegistryFilterer: WorkerRegistryFilterer{contract: contract}}, nil
}

// NewWorkerRegistryCaller creates a new read-only instance of WorkerRegistry, bound to a specific deployed contract.
func NewWorkerRegistryCaller(address common.Address, caller bind.ContractCaller) (*WorkerRegistryCaller, error) {
	contract, err := bindWorkerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryCaller{contract: contract}, nil
}

// NewWorkerRegistryTransactor creates a new write-only instance of WorkerRegistry, bound to a specific deployed contract.
func NewWorkerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerRegistryTransactor, error) {
	contract, err := bindWorkerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryTransactor{contract: contract}, nil
}

// NewWorkerRegistryFilterer creates a new log filterer instance of WorkerRegistry, bound to a specific deployed contract.
func NewWorkerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerRegistryFilterer, error) {
	contract, err := bindWorkerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryFilterer{contract: contract}, nil
}

// bindWorkerRegistry binds a generic wrapper to an already deployed contract.
func bindWorkerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerRegistry *WorkerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerRegistry.Contract.WorkerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerRegistry *WorkerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WorkerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerRegistry *WorkerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WorkerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerRegistry *WorkerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerRegistry *WorkerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerRegistry *WorkerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WorkerRegistry *WorkerRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WorkerRegistry *WorkerRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WorkerRegistry.Contract.UPGRADEINTERFACEVERSION(&_WorkerRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WorkerRegistry *WorkerRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WorkerRegistry.Contract.UPGRADEINTERFACEVERSION(&_WorkerRegistry.CallOpts)
}

// AiConfig is a free data retrieval call binding the contract method 0x85ff4862.
//
// Solidity: function aiConfig() view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) AiConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "aiConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AiConfig is a free data retrieval call binding the contract method 0x85ff4862.
//
// Solidity: function aiConfig() view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) AiConfig() (common.Address, error) {
	return _WorkerRegistry.Contract.AiConfig(&_WorkerRegistry.CallOpts)
}

// AiConfig is a free data retrieval call binding the contract method 0x85ff4862.
//
// Solidity: function aiConfig() view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) AiConfig() (common.Address, error) {
	return _WorkerRegistry.Contract.AiConfig(&_WorkerRegistry.CallOpts)
}

// GetActiveJobCount is a free data retrieval call binding the contract method 0xfc532ccd.
//
// Solidity: function getActiveJobCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetActiveJobCount(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getActiveJobCount", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveJobCount is a free data retrieval call binding the contract method 0xfc532ccd.
//
// Solidity: function getActiveJobCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetActiveJobCount(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetActiveJobCount(&_WorkerRegistry.CallOpts, worker)
}

// GetActiveJobCount is a free data retrieval call binding the contract method 0xfc532ccd.
//
// Solidity: function getActiveJobCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetActiveJobCount(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetActiveJobCount(&_WorkerRegistry.CallOpts, worker)
}

// GetCapabilityCount is a free data retrieval call binding the contract method 0xb036fe7e.
//
// Solidity: function getCapabilityCount() view returns(uint16)
func (_WorkerRegistry *WorkerRegistryCaller) GetCapabilityCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getCapabilityCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetCapabilityCount is a free data retrieval call binding the contract method 0xb036fe7e.
//
// Solidity: function getCapabilityCount() view returns(uint16)
func (_WorkerRegistry *WorkerRegistrySession) GetCapabilityCount() (uint16, error) {
	return _WorkerRegistry.Contract.GetCapabilityCount(&_WorkerRegistry.CallOpts)
}

// GetCapabilityCount is a free data retrieval call binding the contract method 0xb036fe7e.
//
// Solidity: function getCapabilityCount() view returns(uint16)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetCapabilityCount() (uint16, error) {
	return _WorkerRegistry.Contract.GetCapabilityCount(&_WorkerRegistry.CallOpts)
}

// GetCapabilityMask is a free data retrieval call binding the contract method 0xd3d6402a.
//
// Solidity: function getCapabilityMask(string name) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetCapabilityMask(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getCapabilityMask", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCapabilityMask is a free data retrieval call binding the contract method 0xd3d6402a.
//
// Solidity: function getCapabilityMask(string name) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetCapabilityMask(name string) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetCapabilityMask(&_WorkerRegistry.CallOpts, name)
}

// GetCapabilityMask is a free data retrieval call binding the contract method 0xd3d6402a.
//
// Solidity: function getCapabilityMask(string name) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetCapabilityMask(name string) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetCapabilityMask(&_WorkerRegistry.CallOpts, name)
}

// GetCapabilityName is a free data retrieval call binding the contract method 0x61568e08.
//
// Solidity: function getCapabilityName(uint8 bit) view returns(string)
func (_WorkerRegistry *WorkerRegistryCaller) GetCapabilityName(opts *bind.CallOpts, bit uint8) (string, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getCapabilityName", bit)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCapabilityName is a free data retrieval call binding the contract method 0x61568e08.
//
// Solidity: function getCapabilityName(uint8 bit) view returns(string)
func (_WorkerRegistry *WorkerRegistrySession) GetCapabilityName(bit uint8) (string, error) {
	return _WorkerRegistry.Contract.GetCapabilityName(&_WorkerRegistry.CallOpts, bit)
}

// GetCapabilityName is a free data retrieval call binding the contract method 0x61568e08.
//
// Solidity: function getCapabilityName(uint8 bit) view returns(string)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetCapabilityName(bit uint8) (string, error) {
	return _WorkerRegistry.Contract.GetCapabilityName(&_WorkerRegistry.CallOpts, bit)
}

// GetEligibleWorkers is a free data retrieval call binding the contract method 0xf88cca87.
//
// Solidity: function getEligibleWorkers(bytes32 modelId) view returns(address[])
func (_WorkerRegistry *WorkerRegistryCaller) GetEligibleWorkers(opts *bind.CallOpts, modelId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getEligibleWorkers", modelId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEligibleWorkers is a free data retrieval call binding the contract method 0xf88cca87.
//
// Solidity: function getEligibleWorkers(bytes32 modelId) view returns(address[])
func (_WorkerRegistry *WorkerRegistrySession) GetEligibleWorkers(modelId [32]byte) ([]common.Address, error) {
	return _WorkerRegistry.Contract.GetEligibleWorkers(&_WorkerRegistry.CallOpts, modelId)
}

// GetEligibleWorkers is a free data retrieval call binding the contract method 0xf88cca87.
//
// Solidity: function getEligibleWorkers(bytes32 modelId) view returns(address[])
func (_WorkerRegistry *WorkerRegistryCallerSession) GetEligibleWorkers(modelId [32]byte) ([]common.Address, error) {
	return _WorkerRegistry.Contract.GetEligibleWorkers(&_WorkerRegistry.CallOpts, modelId)
}

// GetOffenseCount is a free data retrieval call binding the contract method 0x3292a1a8.
//
// Solidity: function getOffenseCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetOffenseCount(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getOffenseCount", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOffenseCount is a free data retrieval call binding the contract method 0x3292a1a8.
//
// Solidity: function getOffenseCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetOffenseCount(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetOffenseCount(&_WorkerRegistry.CallOpts, worker)
}

// GetOffenseCount is a free data retrieval call binding the contract method 0x3292a1a8.
//
// Solidity: function getOffenseCount(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetOffenseCount(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetOffenseCount(&_WorkerRegistry.CallOpts, worker)
}

// GetSlashedFunds is a free data retrieval call binding the contract method 0x4d7b03f9.
//
// Solidity: function getSlashedFunds() view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetSlashedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getSlashedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashedFunds is a free data retrieval call binding the contract method 0x4d7b03f9.
//
// Solidity: function getSlashedFunds() view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetSlashedFunds() (*big.Int, error) {
	return _WorkerRegistry.Contract.GetSlashedFunds(&_WorkerRegistry.CallOpts)
}

// GetSlashedFunds is a free data retrieval call binding the contract method 0x4d7b03f9.
//
// Solidity: function getSlashedFunds() view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetSlashedFunds() (*big.Int, error) {
	return _WorkerRegistry.Contract.GetSlashedFunds(&_WorkerRegistry.CallOpts)
}

// GetSuspendedUntil is a free data retrieval call binding the contract method 0x2368c52c.
//
// Solidity: function getSuspendedUntil(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetSuspendedUntil(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getSuspendedUntil", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuspendedUntil is a free data retrieval call binding the contract method 0x2368c52c.
//
// Solidity: function getSuspendedUntil(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetSuspendedUntil(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetSuspendedUntil(&_WorkerRegistry.CallOpts, worker)
}

// GetSuspendedUntil is a free data retrieval call binding the contract method 0x2368c52c.
//
// Solidity: function getSuspendedUntil(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetSuspendedUntil(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetSuspendedUntil(&_WorkerRegistry.CallOpts, worker)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2f8899f3.
//
// Solidity: function getTotalStaked(uint256 blockNumber) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetTotalStaked(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getTotalStaked", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2f8899f3.
//
// Solidity: function getTotalStaked(uint256 blockNumber) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetTotalStaked(blockNumber *big.Int) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetTotalStaked(&_WorkerRegistry.CallOpts, blockNumber)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2f8899f3.
//
// Solidity: function getTotalStaked(uint256 blockNumber) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetTotalStaked(blockNumber *big.Int) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetTotalStaked(&_WorkerRegistry.CallOpts, blockNumber)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) GetTreasury() (common.Address, error) {
	return _WorkerRegistry.Contract.GetTreasury(&_WorkerRegistry.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetTreasury() (common.Address, error) {
	return _WorkerRegistry.Contract.GetTreasury(&_WorkerRegistry.CallOpts)
}

// GetWorkerCapabilities is a free data retrieval call binding the contract method 0x0dd1e8e0.
//
// Solidity: function getWorkerCapabilities(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetWorkerCapabilities(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getWorkerCapabilities", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkerCapabilities is a free data retrieval call binding the contract method 0x0dd1e8e0.
//
// Solidity: function getWorkerCapabilities(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetWorkerCapabilities(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetWorkerCapabilities(&_WorkerRegistry.CallOpts, worker)
}

// GetWorkerCapabilities is a free data retrieval call binding the contract method 0x0dd1e8e0.
//
// Solidity: function getWorkerCapabilities(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetWorkerCapabilities(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetWorkerCapabilities(&_WorkerRegistry.CallOpts, worker)
}

// GetWorkerEncryptionKey is a free data retrieval call binding the contract method 0x80034cc7.
//
// Solidity: function getWorkerEncryptionKey(address worker) view returns(bytes)
func (_WorkerRegistry *WorkerRegistryCaller) GetWorkerEncryptionKey(opts *bind.CallOpts, worker common.Address) ([]byte, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getWorkerEncryptionKey", worker)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWorkerEncryptionKey is a free data retrieval call binding the contract method 0x80034cc7.
//
// Solidity: function getWorkerEncryptionKey(address worker) view returns(bytes)
func (_WorkerRegistry *WorkerRegistrySession) GetWorkerEncryptionKey(worker common.Address) ([]byte, error) {
	return _WorkerRegistry.Contract.GetWorkerEncryptionKey(&_WorkerRegistry.CallOpts, worker)
}

// GetWorkerEncryptionKey is a free data retrieval call binding the contract method 0x80034cc7.
//
// Solidity: function getWorkerEncryptionKey(address worker) view returns(bytes)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetWorkerEncryptionKey(worker common.Address) ([]byte, error) {
	return _WorkerRegistry.Contract.GetWorkerEncryptionKey(&_WorkerRegistry.CallOpts, worker)
}

// GetWorkerStake is a free data retrieval call binding the contract method 0x7a55f24d.
//
// Solidity: function getWorkerStake(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCaller) GetWorkerStake(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "getWorkerStake", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkerStake is a free data retrieval call binding the contract method 0x7a55f24d.
//
// Solidity: function getWorkerStake(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistrySession) GetWorkerStake(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetWorkerStake(&_WorkerRegistry.CallOpts, worker)
}

// GetWorkerStake is a free data retrieval call binding the contract method 0x7a55f24d.
//
// Solidity: function getWorkerStake(address worker) view returns(uint256)
func (_WorkerRegistry *WorkerRegistryCallerSession) GetWorkerStake(worker common.Address) (*big.Int, error) {
	return _WorkerRegistry.Contract.GetWorkerStake(&_WorkerRegistry.CallOpts, worker)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) Guardian() (common.Address, error) {
	return _WorkerRegistry.Contract.Guardian(&_WorkerRegistry.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) Guardian() (common.Address, error) {
	return _WorkerRegistry.Contract.Guardian(&_WorkerRegistry.CallOpts)
}

// IsEligible is a free data retrieval call binding the contract method 0xdb3ebef1.
//
// Solidity: function isEligible(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) IsEligible(opts *bind.CallOpts, worker common.Address, modelId [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "isEligible", worker, modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEligible is a free data retrieval call binding the contract method 0xdb3ebef1.
//
// Solidity: function isEligible(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) IsEligible(worker common.Address, modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.IsEligible(&_WorkerRegistry.CallOpts, worker, modelId)
}

// IsEligible is a free data retrieval call binding the contract method 0xdb3ebef1.
//
// Solidity: function isEligible(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) IsEligible(worker common.Address, modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.IsEligible(&_WorkerRegistry.CallOpts, worker, modelId)
}

// IsModelWhitelisted is a free data retrieval call binding the contract method 0xf42e0c2e.
//
// Solidity: function isModelWhitelisted(bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) IsModelWhitelisted(opts *bind.CallOpts, modelId [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "isModelWhitelisted", modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModelWhitelisted is a free data retrieval call binding the contract method 0xf42e0c2e.
//
// Solidity: function isModelWhitelisted(bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) IsModelWhitelisted(modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.IsModelWhitelisted(&_WorkerRegistry.CallOpts, modelId)
}

// IsModelWhitelisted is a free data retrieval call binding the contract method 0xf42e0c2e.
//
// Solidity: function isModelWhitelisted(bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) IsModelWhitelisted(modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.IsModelWhitelisted(&_WorkerRegistry.CallOpts, modelId)
}

// IsWorkerRegistered is a free data retrieval call binding the contract method 0xe798a7da.
//
// Solidity: function isWorkerRegistered(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) IsWorkerRegistered(opts *bind.CallOpts, worker common.Address) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "isWorkerRegistered", worker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWorkerRegistered is a free data retrieval call binding the contract method 0xe798a7da.
//
// Solidity: function isWorkerRegistered(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) IsWorkerRegistered(worker common.Address) (bool, error) {
	return _WorkerRegistry.Contract.IsWorkerRegistered(&_WorkerRegistry.CallOpts, worker)
}

// IsWorkerRegistered is a free data retrieval call binding the contract method 0xe798a7da.
//
// Solidity: function isWorkerRegistered(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) IsWorkerRegistered(worker common.Address) (bool, error) {
	return _WorkerRegistry.Contract.IsWorkerRegistered(&_WorkerRegistry.CallOpts, worker)
}

// IsWorkerSuspended is a free data retrieval call binding the contract method 0x2808ae43.
//
// Solidity: function isWorkerSuspended(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) IsWorkerSuspended(opts *bind.CallOpts, worker common.Address) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "isWorkerSuspended", worker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWorkerSuspended is a free data retrieval call binding the contract method 0x2808ae43.
//
// Solidity: function isWorkerSuspended(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) IsWorkerSuspended(worker common.Address) (bool, error) {
	return _WorkerRegistry.Contract.IsWorkerSuspended(&_WorkerRegistry.CallOpts, worker)
}

// IsWorkerSuspended is a free data retrieval call binding the contract method 0x2808ae43.
//
// Solidity: function isWorkerSuspended(address worker) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) IsWorkerSuspended(worker common.Address) (bool, error) {
	return _WorkerRegistry.Contract.IsWorkerSuspended(&_WorkerRegistry.CallOpts, worker)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) JobRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "jobRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) JobRegistry() (common.Address, error) {
	return _WorkerRegistry.Contract.JobRegistry(&_WorkerRegistry.CallOpts)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) JobRegistry() (common.Address, error) {
	return _WorkerRegistry.Contract.JobRegistry(&_WorkerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) Owner() (common.Address, error) {
	return _WorkerRegistry.Contract.Owner(&_WorkerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) Owner() (common.Address, error) {
	return _WorkerRegistry.Contract.Owner(&_WorkerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) Paused() (bool, error) {
	return _WorkerRegistry.Contract.Paused(&_WorkerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) Paused() (bool, error) {
	return _WorkerRegistry.Contract.Paused(&_WorkerRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorkerRegistry *WorkerRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorkerRegistry *WorkerRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _WorkerRegistry.Contract.ProxiableUUID(&_WorkerRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorkerRegistry *WorkerRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WorkerRegistry.Contract.ProxiableUUID(&_WorkerRegistry.CallOpts)
}

// SelectEligibleWorker is a free data retrieval call binding the contract method 0xb985528e.
//
// Solidity: function selectEligibleWorker(bytes32 modelId, address[] excluded, uint256 seed) view returns(address)
func (_WorkerRegistry *WorkerRegistryCaller) SelectEligibleWorker(opts *bind.CallOpts, modelId [32]byte, excluded []common.Address, seed *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "selectEligibleWorker", modelId, excluded, seed)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectEligibleWorker is a free data retrieval call binding the contract method 0xb985528e.
//
// Solidity: function selectEligibleWorker(bytes32 modelId, address[] excluded, uint256 seed) view returns(address)
func (_WorkerRegistry *WorkerRegistrySession) SelectEligibleWorker(modelId [32]byte, excluded []common.Address, seed *big.Int) (common.Address, error) {
	return _WorkerRegistry.Contract.SelectEligibleWorker(&_WorkerRegistry.CallOpts, modelId, excluded, seed)
}

// SelectEligibleWorker is a free data retrieval call binding the contract method 0xb985528e.
//
// Solidity: function selectEligibleWorker(bytes32 modelId, address[] excluded, uint256 seed) view returns(address)
func (_WorkerRegistry *WorkerRegistryCallerSession) SelectEligibleWorker(modelId [32]byte, excluded []common.Address, seed *big.Int) (common.Address, error) {
	return _WorkerRegistry.Contract.SelectEligibleWorker(&_WorkerRegistry.CallOpts, modelId, excluded, seed)
}

// WorkerSupportsModel is a free data retrieval call binding the contract method 0xd660efa1.
//
// Solidity: function workerSupportsModel(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCaller) WorkerSupportsModel(opts *bind.CallOpts, worker common.Address, modelId [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkerRegistry.contract.Call(opts, &out, "workerSupportsModel", worker, modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WorkerSupportsModel is a free data retrieval call binding the contract method 0xd660efa1.
//
// Solidity: function workerSupportsModel(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistrySession) WorkerSupportsModel(worker common.Address, modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.WorkerSupportsModel(&_WorkerRegistry.CallOpts, worker, modelId)
}

// WorkerSupportsModel is a free data retrieval call binding the contract method 0xd660efa1.
//
// Solidity: function workerSupportsModel(address worker, bytes32 modelId) view returns(bool)
func (_WorkerRegistry *WorkerRegistryCallerSession) WorkerSupportsModel(worker common.Address, modelId [32]byte) (bool, error) {
	return _WorkerRegistry.Contract.WorkerSupportsModel(&_WorkerRegistry.CallOpts, worker, modelId)
}

// AddSupportedModel is a paid mutator transaction binding the contract method 0x62d657bd.
//
// Solidity: function addSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) AddSupportedModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "addSupportedModel", modelId)
}

// AddSupportedModel is a paid mutator transaction binding the contract method 0x62d657bd.
//
// Solidity: function addSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) AddSupportedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.AddSupportedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// AddSupportedModel is a paid mutator transaction binding the contract method 0x62d657bd.
//
// Solidity: function addSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) AddSupportedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.AddSupportedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// AddWhitelistedModel is a paid mutator transaction binding the contract method 0xc234d50d.
//
// Solidity: function addWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) AddWhitelistedModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "addWhitelistedModel", modelId)
}

// AddWhitelistedModel is a paid mutator transaction binding the contract method 0xc234d50d.
//
// Solidity: function addWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) AddWhitelistedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.AddWhitelistedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// AddWhitelistedModel is a paid mutator transaction binding the contract method 0xc234d50d.
//
// Solidity: function addWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) AddWhitelistedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.AddWhitelistedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// DecrementActiveJobs is a paid mutator transaction binding the contract method 0x0439f300.
//
// Solidity: function decrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) DecrementActiveJobs(opts *bind.TransactOpts, worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "decrementActiveJobs", worker, modelId)
}

// DecrementActiveJobs is a paid mutator transaction binding the contract method 0x0439f300.
//
// Solidity: function decrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) DecrementActiveJobs(worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.DecrementActiveJobs(&_WorkerRegistry.TransactOpts, worker, modelId)
}

// DecrementActiveJobs is a paid mutator transaction binding the contract method 0x0439f300.
//
// Solidity: function decrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) DecrementActiveJobs(worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.DecrementActiveJobs(&_WorkerRegistry.TransactOpts, worker, modelId)
}

// DeregisterWorker is a paid mutator transaction binding the contract method 0x200cd650.
//
// Solidity: function deregisterWorker() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) DeregisterWorker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "deregisterWorker")
}

// DeregisterWorker is a paid mutator transaction binding the contract method 0x200cd650.
//
// Solidity: function deregisterWorker() returns()
func (_WorkerRegistry *WorkerRegistrySession) DeregisterWorker() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.DeregisterWorker(&_WorkerRegistry.TransactOpts)
}

// DeregisterWorker is a paid mutator transaction binding the contract method 0x200cd650.
//
// Solidity: function deregisterWorker() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) DeregisterWorker() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.DeregisterWorker(&_WorkerRegistry.TransactOpts)
}

// IncrementActiveJobs is a paid mutator transaction binding the contract method 0x5c3d7372.
//
// Solidity: function incrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) IncrementActiveJobs(opts *bind.TransactOpts, worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "incrementActiveJobs", worker, modelId)
}

// IncrementActiveJobs is a paid mutator transaction binding the contract method 0x5c3d7372.
//
// Solidity: function incrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) IncrementActiveJobs(worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.IncrementActiveJobs(&_WorkerRegistry.TransactOpts, worker, modelId)
}

// IncrementActiveJobs is a paid mutator transaction binding the contract method 0x5c3d7372.
//
// Solidity: function incrementActiveJobs(address worker, bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) IncrementActiveJobs(worker common.Address, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.IncrementActiveJobs(&_WorkerRegistry.TransactOpts, worker, modelId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _initialOwner, address _aiConfig, address _jobRegistry, address _treasury, address _guardian) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _aiConfig common.Address, _jobRegistry common.Address, _treasury common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "initialize", _initialOwner, _aiConfig, _jobRegistry, _treasury, _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _initialOwner, address _aiConfig, address _jobRegistry, address _treasury, address _guardian) returns()
func (_WorkerRegistry *WorkerRegistrySession) Initialize(_initialOwner common.Address, _aiConfig common.Address, _jobRegistry common.Address, _treasury common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Initialize(&_WorkerRegistry.TransactOpts, _initialOwner, _aiConfig, _jobRegistry, _treasury, _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _initialOwner, address _aiConfig, address _jobRegistry, address _treasury, address _guardian) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) Initialize(_initialOwner common.Address, _aiConfig common.Address, _jobRegistry common.Address, _treasury common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Initialize(&_WorkerRegistry.TransactOpts, _initialOwner, _aiConfig, _jobRegistry, _treasury, _guardian)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerRegistry *WorkerRegistrySession) Pause() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Pause(&_WorkerRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Pause(&_WorkerRegistry.TransactOpts)
}

// RegisterCapability is a paid mutator transaction binding the contract method 0x39441388.
//
// Solidity: function registerCapability(string name) returns(uint8 bit)
func (_WorkerRegistry *WorkerRegistryTransactor) RegisterCapability(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "registerCapability", name)
}

// RegisterCapability is a paid mutator transaction binding the contract method 0x39441388.
//
// Solidity: function registerCapability(string name) returns(uint8 bit)
func (_WorkerRegistry *WorkerRegistrySession) RegisterCapability(name string) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RegisterCapability(&_WorkerRegistry.TransactOpts, name)
}

// RegisterCapability is a paid mutator transaction binding the contract method 0x39441388.
//
// Solidity: function registerCapability(string name) returns(uint8 bit)
func (_WorkerRegistry *WorkerRegistryTransactorSession) RegisterCapability(name string) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RegisterCapability(&_WorkerRegistry.TransactOpts, name)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0xee066b4d.
//
// Solidity: function registerWorker(bytes encryptionPubKey) payable returns()
func (_WorkerRegistry *WorkerRegistryTransactor) RegisterWorker(opts *bind.TransactOpts, encryptionPubKey []byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "registerWorker", encryptionPubKey)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0xee066b4d.
//
// Solidity: function registerWorker(bytes encryptionPubKey) payable returns()
func (_WorkerRegistry *WorkerRegistrySession) RegisterWorker(encryptionPubKey []byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RegisterWorker(&_WorkerRegistry.TransactOpts, encryptionPubKey)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0xee066b4d.
//
// Solidity: function registerWorker(bytes encryptionPubKey) payable returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) RegisterWorker(encryptionPubKey []byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RegisterWorker(&_WorkerRegistry.TransactOpts, encryptionPubKey)
}

// Reinstate is a paid mutator transaction binding the contract method 0x28fecd48.
//
// Solidity: function reinstate() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) Reinstate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "reinstate")
}

// Reinstate is a paid mutator transaction binding the contract method 0x28fecd48.
//
// Solidity: function reinstate() returns()
func (_WorkerRegistry *WorkerRegistrySession) Reinstate() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Reinstate(&_WorkerRegistry.TransactOpts)
}

// Reinstate is a paid mutator transaction binding the contract method 0x28fecd48.
//
// Solidity: function reinstate() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) Reinstate() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Reinstate(&_WorkerRegistry.TransactOpts)
}

// RemoveSupportedModel is a paid mutator transaction binding the contract method 0xa64d940a.
//
// Solidity: function removeSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) RemoveSupportedModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "removeSupportedModel", modelId)
}

// RemoveSupportedModel is a paid mutator transaction binding the contract method 0xa64d940a.
//
// Solidity: function removeSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) RemoveSupportedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RemoveSupportedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// RemoveSupportedModel is a paid mutator transaction binding the contract method 0xa64d940a.
//
// Solidity: function removeSupportedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) RemoveSupportedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RemoveSupportedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// RemoveWhitelistedModel is a paid mutator transaction binding the contract method 0x2116ad86.
//
// Solidity: function removeWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) RemoveWhitelistedModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "removeWhitelistedModel", modelId)
}

// RemoveWhitelistedModel is a paid mutator transaction binding the contract method 0x2116ad86.
//
// Solidity: function removeWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistrySession) RemoveWhitelistedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RemoveWhitelistedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// RemoveWhitelistedModel is a paid mutator transaction binding the contract method 0x2116ad86.
//
// Solidity: function removeWhitelistedModel(bytes32 modelId) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) RemoveWhitelistedModel(modelId [32]byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RemoveWhitelistedModel(&_WorkerRegistry.TransactOpts, modelId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerRegistry *WorkerRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RenounceOwnership(&_WorkerRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.RenounceOwnership(&_WorkerRegistry.TransactOpts)
}

// ResetOffenses is a paid mutator transaction binding the contract method 0x1280d02d.
//
// Solidity: function resetOffenses(address worker) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) ResetOffenses(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "resetOffenses", worker)
}

// ResetOffenses is a paid mutator transaction binding the contract method 0x1280d02d.
//
// Solidity: function resetOffenses(address worker) returns()
func (_WorkerRegistry *WorkerRegistrySession) ResetOffenses(worker common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.ResetOffenses(&_WorkerRegistry.TransactOpts, worker)
}

// ResetOffenses is a paid mutator transaction binding the contract method 0x1280d02d.
//
// Solidity: function resetOffenses(address worker) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) ResetOffenses(worker common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.ResetOffenses(&_WorkerRegistry.TransactOpts, worker)
}

// SetCapabilities is a paid mutator transaction binding the contract method 0x38bbddd9.
//
// Solidity: function setCapabilities(uint256 mask) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) SetCapabilities(opts *bind.TransactOpts, mask *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "setCapabilities", mask)
}

// SetCapabilities is a paid mutator transaction binding the contract method 0x38bbddd9.
//
// Solidity: function setCapabilities(uint256 mask) returns()
func (_WorkerRegistry *WorkerRegistrySession) SetCapabilities(mask *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SetCapabilities(&_WorkerRegistry.TransactOpts, mask)
}

// SetCapabilities is a paid mutator transaction binding the contract method 0x38bbddd9.
//
// Solidity: function setCapabilities(uint256 mask) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) SetCapabilities(mask *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SetCapabilities(&_WorkerRegistry.TransactOpts, mask)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address newGuardian) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) SetGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "setGuardian", newGuardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address newGuardian) returns()
func (_WorkerRegistry *WorkerRegistrySession) SetGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SetGuardian(&_WorkerRegistry.TransactOpts, newGuardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address newGuardian) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) SetGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SetGuardian(&_WorkerRegistry.TransactOpts, newGuardian)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address worker, uint256 bps) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) Slash(opts *bind.TransactOpts, worker common.Address, bps *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "slash", worker, bps)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address worker, uint256 bps) returns()
func (_WorkerRegistry *WorkerRegistrySession) Slash(worker common.Address, bps *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Slash(&_WorkerRegistry.TransactOpts, worker, bps)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address worker, uint256 bps) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) Slash(worker common.Address, bps *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Slash(&_WorkerRegistry.TransactOpts, worker, bps)
}

// SlashWithBounty is a paid mutator transaction binding the contract method 0xd11bd4ef.
//
// Solidity: function slashWithBounty(address worker, uint256 bps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistryTransactor) SlashWithBounty(opts *bind.TransactOpts, worker common.Address, bps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "slashWithBounty", worker, bps, bountyRecipient)
}

// SlashWithBounty is a paid mutator transaction binding the contract method 0xd11bd4ef.
//
// Solidity: function slashWithBounty(address worker, uint256 bps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistrySession) SlashWithBounty(worker common.Address, bps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SlashWithBounty(&_WorkerRegistry.TransactOpts, worker, bps, bountyRecipient)
}

// SlashWithBounty is a paid mutator transaction binding the contract method 0xd11bd4ef.
//
// Solidity: function slashWithBounty(address worker, uint256 bps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistryTransactorSession) SlashWithBounty(worker common.Address, bps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SlashWithBounty(&_WorkerRegistry.TransactOpts, worker, bps, bountyRecipient)
}

// SlashWithBountyBps is a paid mutator transaction binding the contract method 0x26778c85.
//
// Solidity: function slashWithBountyBps(address worker, uint256 bps, uint256 bountyBps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistryTransactor) SlashWithBountyBps(opts *bind.TransactOpts, worker common.Address, bps *big.Int, bountyBps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "slashWithBountyBps", worker, bps, bountyBps, bountyRecipient)
}

// SlashWithBountyBps is a paid mutator transaction binding the contract method 0x26778c85.
//
// Solidity: function slashWithBountyBps(address worker, uint256 bps, uint256 bountyBps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistrySession) SlashWithBountyBps(worker common.Address, bps *big.Int, bountyBps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SlashWithBountyBps(&_WorkerRegistry.TransactOpts, worker, bps, bountyBps, bountyRecipient)
}

// SlashWithBountyBps is a paid mutator transaction binding the contract method 0x26778c85.
//
// Solidity: function slashWithBountyBps(address worker, uint256 bps, uint256 bountyBps, address bountyRecipient) returns(uint256 slashAmount, uint256 bountyPaid)
func (_WorkerRegistry *WorkerRegistryTransactorSession) SlashWithBountyBps(worker common.Address, bps *big.Int, bountyBps *big.Int, bountyRecipient common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SlashWithBountyBps(&_WorkerRegistry.TransactOpts, worker, bps, bountyBps, bountyRecipient)
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) SunsetGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "sunsetGuardian")
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_WorkerRegistry *WorkerRegistrySession) SunsetGuardian() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SunsetGuardian(&_WorkerRegistry.TransactOpts)
}

// SunsetGuardian is a paid mutator transaction binding the contract method 0xec096db0.
//
// Solidity: function sunsetGuardian() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) SunsetGuardian() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.SunsetGuardian(&_WorkerRegistry.TransactOpts)
}

// TopUpStake is a paid mutator transaction binding the contract method 0x5a194108.
//
// Solidity: function topUpStake() payable returns()
func (_WorkerRegistry *WorkerRegistryTransactor) TopUpStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "topUpStake")
}

// TopUpStake is a paid mutator transaction binding the contract method 0x5a194108.
//
// Solidity: function topUpStake() payable returns()
func (_WorkerRegistry *WorkerRegistrySession) TopUpStake() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.TopUpStake(&_WorkerRegistry.TransactOpts)
}

// TopUpStake is a paid mutator transaction binding the contract method 0x5a194108.
//
// Solidity: function topUpStake() payable returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) TopUpStake() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.TopUpStake(&_WorkerRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerRegistry *WorkerRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.TransferOwnership(&_WorkerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.TransferOwnership(&_WorkerRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerRegistry *WorkerRegistrySession) Unpause() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Unpause(&_WorkerRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.Unpause(&_WorkerRegistry.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorkerRegistry *WorkerRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorkerRegistry *WorkerRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.UpgradeToAndCall(&_WorkerRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.UpgradeToAndCall(&_WorkerRegistry.TransactOpts, newImplementation, data)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x080da133.
//
// Solidity: function withdrawSlashedFunds() returns()
func (_WorkerRegistry *WorkerRegistryTransactor) WithdrawSlashedFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "withdrawSlashedFunds")
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x080da133.
//
// Solidity: function withdrawSlashedFunds() returns()
func (_WorkerRegistry *WorkerRegistrySession) WithdrawSlashedFunds() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WithdrawSlashedFunds(&_WorkerRegistry.TransactOpts)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x080da133.
//
// Solidity: function withdrawSlashedFunds() returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) WithdrawSlashedFunds() (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WithdrawSlashedFunds(&_WorkerRegistry.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_WorkerRegistry *WorkerRegistryTransactor) WithdrawStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.contract.Transact(opts, "withdrawStake", amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_WorkerRegistry *WorkerRegistrySession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WithdrawStake(&_WorkerRegistry.TransactOpts, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_WorkerRegistry *WorkerRegistryTransactorSession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _WorkerRegistry.Contract.WithdrawStake(&_WorkerRegistry.TransactOpts, amount)
}

// WorkerRegistryCapabilityRegisteredIterator is returned from FilterCapabilityRegistered and is used to iterate over the raw logs and unpacked data for CapabilityRegistered events raised by the WorkerRegistry contract.
type WorkerRegistryCapabilityRegisteredIterator struct {
	Event *WorkerRegistryCapabilityRegistered // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryCapabilityRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryCapabilityRegistered)
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
		it.Event = new(WorkerRegistryCapabilityRegistered)
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
func (it *WorkerRegistryCapabilityRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryCapabilityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryCapabilityRegistered represents a CapabilityRegistered event raised by the WorkerRegistry contract.
type WorkerRegistryCapabilityRegistered struct {
	Bit  uint8
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCapabilityRegistered is a free log retrieval operation binding the contract event 0xda84914123ecd65b01462dd0b3fb0dde2988c2e933595751f7c1c821a182dcc4.
//
// Solidity: event CapabilityRegistered(uint8 indexed bit, string name)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterCapabilityRegistered(opts *bind.FilterOpts, bit []uint8) (*WorkerRegistryCapabilityRegisteredIterator, error) {

	var bitRule []interface{}
	for _, bitItem := range bit {
		bitRule = append(bitRule, bitItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "CapabilityRegistered", bitRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryCapabilityRegisteredIterator{contract: _WorkerRegistry.contract, event: "CapabilityRegistered", logs: logs, sub: sub}, nil
}

// WatchCapabilityRegistered is a free log subscription operation binding the contract event 0xda84914123ecd65b01462dd0b3fb0dde2988c2e933595751f7c1c821a182dcc4.
//
// Solidity: event CapabilityRegistered(uint8 indexed bit, string name)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchCapabilityRegistered(opts *bind.WatchOpts, sink chan<- *WorkerRegistryCapabilityRegistered, bit []uint8) (event.Subscription, error) {

	var bitRule []interface{}
	for _, bitItem := range bit {
		bitRule = append(bitRule, bitItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "CapabilityRegistered", bitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryCapabilityRegistered)
				if err := _WorkerRegistry.contract.UnpackLog(event, "CapabilityRegistered", log); err != nil {
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

// ParseCapabilityRegistered is a log parse operation binding the contract event 0xda84914123ecd65b01462dd0b3fb0dde2988c2e933595751f7c1c821a182dcc4.
//
// Solidity: event CapabilityRegistered(uint8 indexed bit, string name)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseCapabilityRegistered(log types.Log) (*WorkerRegistryCapabilityRegistered, error) {
	event := new(WorkerRegistryCapabilityRegistered)
	if err := _WorkerRegistry.contract.UnpackLog(event, "CapabilityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryGuardianUpdatedIterator is returned from FilterGuardianUpdated and is used to iterate over the raw logs and unpacked data for GuardianUpdated events raised by the WorkerRegistry contract.
type WorkerRegistryGuardianUpdatedIterator struct {
	Event *WorkerRegistryGuardianUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryGuardianUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryGuardianUpdated)
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
		it.Event = new(WorkerRegistryGuardianUpdated)
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
func (it *WorkerRegistryGuardianUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryGuardianUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryGuardianUpdated represents a GuardianUpdated event raised by the WorkerRegistry contract.
type WorkerRegistryGuardianUpdated struct {
	OldGuardian common.Address
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGuardianUpdated is a free log retrieval operation binding the contract event 0x064d28d3d3071c5cbc271a261c10c2f0f0d9e319390397101aa0eb23c6bad909.
//
// Solidity: event GuardianUpdated(address indexed oldGuardian, address indexed newGuardian)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterGuardianUpdated(opts *bind.FilterOpts, oldGuardian []common.Address, newGuardian []common.Address) (*WorkerRegistryGuardianUpdatedIterator, error) {

	var oldGuardianRule []interface{}
	for _, oldGuardianItem := range oldGuardian {
		oldGuardianRule = append(oldGuardianRule, oldGuardianItem)
	}
	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "GuardianUpdated", oldGuardianRule, newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryGuardianUpdatedIterator{contract: _WorkerRegistry.contract, event: "GuardianUpdated", logs: logs, sub: sub}, nil
}

// WatchGuardianUpdated is a free log subscription operation binding the contract event 0x064d28d3d3071c5cbc271a261c10c2f0f0d9e319390397101aa0eb23c6bad909.
//
// Solidity: event GuardianUpdated(address indexed oldGuardian, address indexed newGuardian)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchGuardianUpdated(opts *bind.WatchOpts, sink chan<- *WorkerRegistryGuardianUpdated, oldGuardian []common.Address, newGuardian []common.Address) (event.Subscription, error) {

	var oldGuardianRule []interface{}
	for _, oldGuardianItem := range oldGuardian {
		oldGuardianRule = append(oldGuardianRule, oldGuardianItem)
	}
	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "GuardianUpdated", oldGuardianRule, newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryGuardianUpdated)
				if err := _WorkerRegistry.contract.UnpackLog(event, "GuardianUpdated", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParseGuardianUpdated(log types.Log) (*WorkerRegistryGuardianUpdated, error) {
	event := new(WorkerRegistryGuardianUpdated)
	if err := _WorkerRegistry.contract.UnpackLog(event, "GuardianUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerRegistry contract.
type WorkerRegistryInitializedIterator struct {
	Event *WorkerRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryInitialized)
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
		it.Event = new(WorkerRegistryInitialized)
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
func (it *WorkerRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryInitialized represents a Initialized event raised by the WorkerRegistry contract.
type WorkerRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerRegistryInitializedIterator, error) {

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryInitializedIterator{contract: _WorkerRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryInitialized)
				if err := _WorkerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParseInitialized(log types.Log) (*WorkerRegistryInitialized, error) {
	event := new(WorkerRegistryInitialized)
	if err := _WorkerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryModelAddedIterator is returned from FilterModelAdded and is used to iterate over the raw logs and unpacked data for ModelAdded events raised by the WorkerRegistry contract.
type WorkerRegistryModelAddedIterator struct {
	Event *WorkerRegistryModelAdded // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryModelAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryModelAdded)
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
		it.Event = new(WorkerRegistryModelAdded)
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
func (it *WorkerRegistryModelAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryModelAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryModelAdded represents a ModelAdded event raised by the WorkerRegistry contract.
type WorkerRegistryModelAdded struct {
	Worker  common.Address
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelAdded is a free log retrieval operation binding the contract event 0x609a2a0081c166e57bf456ffdc519b8fca2f7ab2c24f5abc1e3987200a8f263d.
//
// Solidity: event ModelAdded(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterModelAdded(opts *bind.FilterOpts, worker []common.Address, modelId [][32]byte) (*WorkerRegistryModelAddedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "ModelAdded", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryModelAddedIterator{contract: _WorkerRegistry.contract, event: "ModelAdded", logs: logs, sub: sub}, nil
}

// WatchModelAdded is a free log subscription operation binding the contract event 0x609a2a0081c166e57bf456ffdc519b8fca2f7ab2c24f5abc1e3987200a8f263d.
//
// Solidity: event ModelAdded(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchModelAdded(opts *bind.WatchOpts, sink chan<- *WorkerRegistryModelAdded, worker []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "ModelAdded", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryModelAdded)
				if err := _WorkerRegistry.contract.UnpackLog(event, "ModelAdded", log); err != nil {
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

// ParseModelAdded is a log parse operation binding the contract event 0x609a2a0081c166e57bf456ffdc519b8fca2f7ab2c24f5abc1e3987200a8f263d.
//
// Solidity: event ModelAdded(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseModelAdded(log types.Log) (*WorkerRegistryModelAdded, error) {
	event := new(WorkerRegistryModelAdded)
	if err := _WorkerRegistry.contract.UnpackLog(event, "ModelAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryModelDelistedIterator is returned from FilterModelDelisted and is used to iterate over the raw logs and unpacked data for ModelDelisted events raised by the WorkerRegistry contract.
type WorkerRegistryModelDelistedIterator struct {
	Event *WorkerRegistryModelDelisted // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryModelDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryModelDelisted)
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
		it.Event = new(WorkerRegistryModelDelisted)
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
func (it *WorkerRegistryModelDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryModelDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryModelDelisted represents a ModelDelisted event raised by the WorkerRegistry contract.
type WorkerRegistryModelDelisted struct {
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelDelisted is a free log retrieval operation binding the contract event 0xf35cb980118fdf90c1df51a98c9e4595a74f128d617e4dff0b5d7179fba2dd10.
//
// Solidity: event ModelDelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterModelDelisted(opts *bind.FilterOpts, modelId [][32]byte) (*WorkerRegistryModelDelistedIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "ModelDelisted", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryModelDelistedIterator{contract: _WorkerRegistry.contract, event: "ModelDelisted", logs: logs, sub: sub}, nil
}

// WatchModelDelisted is a free log subscription operation binding the contract event 0xf35cb980118fdf90c1df51a98c9e4595a74f128d617e4dff0b5d7179fba2dd10.
//
// Solidity: event ModelDelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchModelDelisted(opts *bind.WatchOpts, sink chan<- *WorkerRegistryModelDelisted, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "ModelDelisted", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryModelDelisted)
				if err := _WorkerRegistry.contract.UnpackLog(event, "ModelDelisted", log); err != nil {
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

// ParseModelDelisted is a log parse operation binding the contract event 0xf35cb980118fdf90c1df51a98c9e4595a74f128d617e4dff0b5d7179fba2dd10.
//
// Solidity: event ModelDelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseModelDelisted(log types.Log) (*WorkerRegistryModelDelisted, error) {
	event := new(WorkerRegistryModelDelisted)
	if err := _WorkerRegistry.contract.UnpackLog(event, "ModelDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryModelRemovedIterator is returned from FilterModelRemoved and is used to iterate over the raw logs and unpacked data for ModelRemoved events raised by the WorkerRegistry contract.
type WorkerRegistryModelRemovedIterator struct {
	Event *WorkerRegistryModelRemoved // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryModelRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryModelRemoved)
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
		it.Event = new(WorkerRegistryModelRemoved)
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
func (it *WorkerRegistryModelRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryModelRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryModelRemoved represents a ModelRemoved event raised by the WorkerRegistry contract.
type WorkerRegistryModelRemoved struct {
	Worker  common.Address
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelRemoved is a free log retrieval operation binding the contract event 0x52c337f918725ebfff17fcddaa7a516d0f0f3181776416ad78c31a568c402b5b.
//
// Solidity: event ModelRemoved(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterModelRemoved(opts *bind.FilterOpts, worker []common.Address, modelId [][32]byte) (*WorkerRegistryModelRemovedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "ModelRemoved", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryModelRemovedIterator{contract: _WorkerRegistry.contract, event: "ModelRemoved", logs: logs, sub: sub}, nil
}

// WatchModelRemoved is a free log subscription operation binding the contract event 0x52c337f918725ebfff17fcddaa7a516d0f0f3181776416ad78c31a568c402b5b.
//
// Solidity: event ModelRemoved(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchModelRemoved(opts *bind.WatchOpts, sink chan<- *WorkerRegistryModelRemoved, worker []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "ModelRemoved", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryModelRemoved)
				if err := _WorkerRegistry.contract.UnpackLog(event, "ModelRemoved", log); err != nil {
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

// ParseModelRemoved is a log parse operation binding the contract event 0x52c337f918725ebfff17fcddaa7a516d0f0f3181776416ad78c31a568c402b5b.
//
// Solidity: event ModelRemoved(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseModelRemoved(log types.Log) (*WorkerRegistryModelRemoved, error) {
	event := new(WorkerRegistryModelRemoved)
	if err := _WorkerRegistry.contract.UnpackLog(event, "ModelRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryModelWhitelistedIterator is returned from FilterModelWhitelisted and is used to iterate over the raw logs and unpacked data for ModelWhitelisted events raised by the WorkerRegistry contract.
type WorkerRegistryModelWhitelistedIterator struct {
	Event *WorkerRegistryModelWhitelisted // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryModelWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryModelWhitelisted)
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
		it.Event = new(WorkerRegistryModelWhitelisted)
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
func (it *WorkerRegistryModelWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryModelWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryModelWhitelisted represents a ModelWhitelisted event raised by the WorkerRegistry contract.
type WorkerRegistryModelWhitelisted struct {
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelWhitelisted is a free log retrieval operation binding the contract event 0x0eb0045c348e61b866083d75ae31ac37eb44d3080753f5be515032b9ce8e82a1.
//
// Solidity: event ModelWhitelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterModelWhitelisted(opts *bind.FilterOpts, modelId [][32]byte) (*WorkerRegistryModelWhitelistedIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "ModelWhitelisted", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryModelWhitelistedIterator{contract: _WorkerRegistry.contract, event: "ModelWhitelisted", logs: logs, sub: sub}, nil
}

// WatchModelWhitelisted is a free log subscription operation binding the contract event 0x0eb0045c348e61b866083d75ae31ac37eb44d3080753f5be515032b9ce8e82a1.
//
// Solidity: event ModelWhitelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchModelWhitelisted(opts *bind.WatchOpts, sink chan<- *WorkerRegistryModelWhitelisted, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "ModelWhitelisted", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryModelWhitelisted)
				if err := _WorkerRegistry.contract.UnpackLog(event, "ModelWhitelisted", log); err != nil {
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

// ParseModelWhitelisted is a log parse operation binding the contract event 0x0eb0045c348e61b866083d75ae31ac37eb44d3080753f5be515032b9ce8e82a1.
//
// Solidity: event ModelWhitelisted(bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseModelWhitelisted(log types.Log) (*WorkerRegistryModelWhitelisted, error) {
	event := new(WorkerRegistryModelWhitelisted)
	if err := _WorkerRegistry.contract.UnpackLog(event, "ModelWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryOffensesResetIterator is returned from FilterOffensesReset and is used to iterate over the raw logs and unpacked data for OffensesReset events raised by the WorkerRegistry contract.
type WorkerRegistryOffensesResetIterator struct {
	Event *WorkerRegistryOffensesReset // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryOffensesResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryOffensesReset)
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
		it.Event = new(WorkerRegistryOffensesReset)
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
func (it *WorkerRegistryOffensesResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryOffensesResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryOffensesReset represents a OffensesReset event raised by the WorkerRegistry contract.
type WorkerRegistryOffensesReset struct {
	Worker        common.Address
	PreviousCount uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOffensesReset is a free log retrieval operation binding the contract event 0x9b0997a029ca4c6f336c817cd0080d29ecd1f4204a8cf5b540d276d195102497.
//
// Solidity: event OffensesReset(address indexed worker, uint32 previousCount)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterOffensesReset(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryOffensesResetIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "OffensesReset", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryOffensesResetIterator{contract: _WorkerRegistry.contract, event: "OffensesReset", logs: logs, sub: sub}, nil
}

// WatchOffensesReset is a free log subscription operation binding the contract event 0x9b0997a029ca4c6f336c817cd0080d29ecd1f4204a8cf5b540d276d195102497.
//
// Solidity: event OffensesReset(address indexed worker, uint32 previousCount)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchOffensesReset(opts *bind.WatchOpts, sink chan<- *WorkerRegistryOffensesReset, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "OffensesReset", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryOffensesReset)
				if err := _WorkerRegistry.contract.UnpackLog(event, "OffensesReset", log); err != nil {
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

// ParseOffensesReset is a log parse operation binding the contract event 0x9b0997a029ca4c6f336c817cd0080d29ecd1f4204a8cf5b540d276d195102497.
//
// Solidity: event OffensesReset(address indexed worker, uint32 previousCount)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseOffensesReset(log types.Log) (*WorkerRegistryOffensesReset, error) {
	event := new(WorkerRegistryOffensesReset)
	if err := _WorkerRegistry.contract.UnpackLog(event, "OffensesReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerRegistry contract.
type WorkerRegistryOwnershipTransferredIterator struct {
	Event *WorkerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryOwnershipTransferred)
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
		it.Event = new(WorkerRegistryOwnershipTransferred)
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
func (it *WorkerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerRegistry contract.
type WorkerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryOwnershipTransferredIterator{contract: _WorkerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryOwnershipTransferred)
				if err := _WorkerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerRegistryOwnershipTransferred, error) {
	event := new(WorkerRegistryOwnershipTransferred)
	if err := _WorkerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerRegistry contract.
type WorkerRegistryPausedIterator struct {
	Event *WorkerRegistryPaused // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryPaused)
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
		it.Event = new(WorkerRegistryPaused)
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
func (it *WorkerRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryPaused represents a Paused event raised by the WorkerRegistry contract.
type WorkerRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerRegistryPausedIterator, error) {

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryPausedIterator{contract: _WorkerRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryPaused)
				if err := _WorkerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParsePaused(log types.Log) (*WorkerRegistryPaused, error) {
	event := new(WorkerRegistryPaused)
	if err := _WorkerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistrySlashedFundsWithdrawnIterator is returned from FilterSlashedFundsWithdrawn and is used to iterate over the raw logs and unpacked data for SlashedFundsWithdrawn events raised by the WorkerRegistry contract.
type WorkerRegistrySlashedFundsWithdrawnIterator struct {
	Event *WorkerRegistrySlashedFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *WorkerRegistrySlashedFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistrySlashedFundsWithdrawn)
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
		it.Event = new(WorkerRegistrySlashedFundsWithdrawn)
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
func (it *WorkerRegistrySlashedFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistrySlashedFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistrySlashedFundsWithdrawn represents a SlashedFundsWithdrawn event raised by the WorkerRegistry contract.
type WorkerRegistrySlashedFundsWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashedFundsWithdrawn is a free log retrieval operation binding the contract event 0xb676ae97062e33f756c764dee2ac7159404c29ad516c4b5dfb707dc7d4e12bd4.
//
// Solidity: event SlashedFundsWithdrawn(address indexed to, uint256 amount)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterSlashedFundsWithdrawn(opts *bind.FilterOpts, to []common.Address) (*WorkerRegistrySlashedFundsWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "SlashedFundsWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistrySlashedFundsWithdrawnIterator{contract: _WorkerRegistry.contract, event: "SlashedFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSlashedFundsWithdrawn is a free log subscription operation binding the contract event 0xb676ae97062e33f756c764dee2ac7159404c29ad516c4b5dfb707dc7d4e12bd4.
//
// Solidity: event SlashedFundsWithdrawn(address indexed to, uint256 amount)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchSlashedFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *WorkerRegistrySlashedFundsWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "SlashedFundsWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistrySlashedFundsWithdrawn)
				if err := _WorkerRegistry.contract.UnpackLog(event, "SlashedFundsWithdrawn", log); err != nil {
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

// ParseSlashedFundsWithdrawn is a log parse operation binding the contract event 0xb676ae97062e33f756c764dee2ac7159404c29ad516c4b5dfb707dc7d4e12bd4.
//
// Solidity: event SlashedFundsWithdrawn(address indexed to, uint256 amount)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseSlashedFundsWithdrawn(log types.Log) (*WorkerRegistrySlashedFundsWithdrawn, error) {
	event := new(WorkerRegistrySlashedFundsWithdrawn)
	if err := _WorkerRegistry.contract.UnpackLog(event, "SlashedFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryStakeTopUpIterator is returned from FilterStakeTopUp and is used to iterate over the raw logs and unpacked data for StakeTopUp events raised by the WorkerRegistry contract.
type WorkerRegistryStakeTopUpIterator struct {
	Event *WorkerRegistryStakeTopUp // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryStakeTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryStakeTopUp)
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
		it.Event = new(WorkerRegistryStakeTopUp)
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
func (it *WorkerRegistryStakeTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryStakeTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryStakeTopUp represents a StakeTopUp event raised by the WorkerRegistry contract.
type WorkerRegistryStakeTopUp struct {
	Worker   common.Address
	Amount   *big.Int
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeTopUp is a free log retrieval operation binding the contract event 0x81249203b9cdd76ddc85c9ac8511dcbd01c78ff941cc7f8f9462564b05c33b64.
//
// Solidity: event StakeTopUp(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterStakeTopUp(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryStakeTopUpIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "StakeTopUp", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryStakeTopUpIterator{contract: _WorkerRegistry.contract, event: "StakeTopUp", logs: logs, sub: sub}, nil
}

// WatchStakeTopUp is a free log subscription operation binding the contract event 0x81249203b9cdd76ddc85c9ac8511dcbd01c78ff941cc7f8f9462564b05c33b64.
//
// Solidity: event StakeTopUp(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchStakeTopUp(opts *bind.WatchOpts, sink chan<- *WorkerRegistryStakeTopUp, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "StakeTopUp", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryStakeTopUp)
				if err := _WorkerRegistry.contract.UnpackLog(event, "StakeTopUp", log); err != nil {
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

// ParseStakeTopUp is a log parse operation binding the contract event 0x81249203b9cdd76ddc85c9ac8511dcbd01c78ff941cc7f8f9462564b05c33b64.
//
// Solidity: event StakeTopUp(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseStakeTopUp(log types.Log) (*WorkerRegistryStakeTopUp, error) {
	event := new(WorkerRegistryStakeTopUp)
	if err := _WorkerRegistry.contract.UnpackLog(event, "StakeTopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the WorkerRegistry contract.
type WorkerRegistryStakeWithdrawnIterator struct {
	Event *WorkerRegistryStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryStakeWithdrawn)
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
		it.Event = new(WorkerRegistryStakeWithdrawn)
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
func (it *WorkerRegistryStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryStakeWithdrawn represents a StakeWithdrawn event raised by the WorkerRegistry contract.
type WorkerRegistryStakeWithdrawn struct {
	Worker   common.Address
	Amount   *big.Int
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x933735aa8de6d7547d0126171b2f31b9c34dd00f3ecd4be85a0ba047db4fafef.
//
// Solidity: event StakeWithdrawn(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryStakeWithdrawnIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "StakeWithdrawn", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryStakeWithdrawnIterator{contract: _WorkerRegistry.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x933735aa8de6d7547d0126171b2f31b9c34dd00f3ecd4be85a0ba047db4fafef.
//
// Solidity: event StakeWithdrawn(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *WorkerRegistryStakeWithdrawn, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "StakeWithdrawn", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryStakeWithdrawn)
				if err := _WorkerRegistry.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x933735aa8de6d7547d0126171b2f31b9c34dd00f3ecd4be85a0ba047db4fafef.
//
// Solidity: event StakeWithdrawn(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseStakeWithdrawn(log types.Log) (*WorkerRegistryStakeWithdrawn, error) {
	event := new(WorkerRegistryStakeWithdrawn)
	if err := _WorkerRegistry.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerRegistry contract.
type WorkerRegistryUnpausedIterator struct {
	Event *WorkerRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryUnpaused)
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
		it.Event = new(WorkerRegistryUnpaused)
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
func (it *WorkerRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryUnpaused represents a Unpaused event raised by the WorkerRegistry contract.
type WorkerRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerRegistryUnpausedIterator, error) {

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryUnpausedIterator{contract: _WorkerRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryUnpaused)
				if err := _WorkerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParseUnpaused(log types.Log) (*WorkerRegistryUnpaused, error) {
	event := new(WorkerRegistryUnpaused)
	if err := _WorkerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WorkerRegistry contract.
type WorkerRegistryUpgradedIterator struct {
	Event *WorkerRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryUpgraded)
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
		it.Event = new(WorkerRegistryUpgraded)
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
func (it *WorkerRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryUpgraded represents a Upgraded event raised by the WorkerRegistry contract.
type WorkerRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WorkerRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryUpgradedIterator{contract: _WorkerRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WorkerRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryUpgraded)
				if err := _WorkerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WorkerRegistry *WorkerRegistryFilterer) ParseUpgraded(log types.Log) (*WorkerRegistryUpgraded, error) {
	event := new(WorkerRegistryUpgraded)
	if err := _WorkerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerCapabilitiesSetIterator is returned from FilterWorkerCapabilitiesSet and is used to iterate over the raw logs and unpacked data for WorkerCapabilitiesSet events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerCapabilitiesSetIterator struct {
	Event *WorkerRegistryWorkerCapabilitiesSet // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerCapabilitiesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerCapabilitiesSet)
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
		it.Event = new(WorkerRegistryWorkerCapabilitiesSet)
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
func (it *WorkerRegistryWorkerCapabilitiesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerCapabilitiesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerCapabilitiesSet represents a WorkerCapabilitiesSet event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerCapabilitiesSet struct {
	Worker common.Address
	Mask   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerCapabilitiesSet is a free log retrieval operation binding the contract event 0xdab16906a519cbc2a53417bb9a1dd8bc2682e01e7d5a615000b83dfd7d9a69e8.
//
// Solidity: event WorkerCapabilitiesSet(address indexed worker, uint256 mask)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerCapabilitiesSet(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerCapabilitiesSetIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerCapabilitiesSet", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerCapabilitiesSetIterator{contract: _WorkerRegistry.contract, event: "WorkerCapabilitiesSet", logs: logs, sub: sub}, nil
}

// WatchWorkerCapabilitiesSet is a free log subscription operation binding the contract event 0xdab16906a519cbc2a53417bb9a1dd8bc2682e01e7d5a615000b83dfd7d9a69e8.
//
// Solidity: event WorkerCapabilitiesSet(address indexed worker, uint256 mask)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerCapabilitiesSet(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerCapabilitiesSet, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerCapabilitiesSet", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerCapabilitiesSet)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerCapabilitiesSet", log); err != nil {
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

// ParseWorkerCapabilitiesSet is a log parse operation binding the contract event 0xdab16906a519cbc2a53417bb9a1dd8bc2682e01e7d5a615000b83dfd7d9a69e8.
//
// Solidity: event WorkerCapabilitiesSet(address indexed worker, uint256 mask)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerCapabilitiesSet(log types.Log) (*WorkerRegistryWorkerCapabilitiesSet, error) {
	event := new(WorkerRegistryWorkerCapabilitiesSet)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerCapabilitiesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerDeactivatedIterator is returned from FilterWorkerDeactivated and is used to iterate over the raw logs and unpacked data for WorkerDeactivated events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerDeactivatedIterator struct {
	Event *WorkerRegistryWorkerDeactivated // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerDeactivated)
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
		it.Event = new(WorkerRegistryWorkerDeactivated)
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
func (it *WorkerRegistryWorkerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerDeactivated represents a WorkerDeactivated event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerDeactivated struct {
	Worker  common.Address
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWorkerDeactivated is a free log retrieval operation binding the contract event 0x9482f9b91134ec1b2f556beea8330bab282c7f5c14befd2e324c23fdedf35138.
//
// Solidity: event WorkerDeactivated(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerDeactivated(opts *bind.FilterOpts, worker []common.Address, modelId [][32]byte) (*WorkerRegistryWorkerDeactivatedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerDeactivated", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerDeactivatedIterator{contract: _WorkerRegistry.contract, event: "WorkerDeactivated", logs: logs, sub: sub}, nil
}

// WatchWorkerDeactivated is a free log subscription operation binding the contract event 0x9482f9b91134ec1b2f556beea8330bab282c7f5c14befd2e324c23fdedf35138.
//
// Solidity: event WorkerDeactivated(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerDeactivated(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerDeactivated, worker []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerDeactivated", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerDeactivated)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerDeactivated", log); err != nil {
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

// ParseWorkerDeactivated is a log parse operation binding the contract event 0x9482f9b91134ec1b2f556beea8330bab282c7f5c14befd2e324c23fdedf35138.
//
// Solidity: event WorkerDeactivated(address indexed worker, bytes32 indexed modelId)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerDeactivated(log types.Log) (*WorkerRegistryWorkerDeactivated, error) {
	event := new(WorkerRegistryWorkerDeactivated)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerDeregisteredIterator is returned from FilterWorkerDeregistered and is used to iterate over the raw logs and unpacked data for WorkerDeregistered events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerDeregisteredIterator struct {
	Event *WorkerRegistryWorkerDeregistered // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerDeregistered)
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
		it.Event = new(WorkerRegistryWorkerDeregistered)
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
func (it *WorkerRegistryWorkerDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerDeregistered represents a WorkerDeregistered event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerDeregistered struct {
	Worker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerDeregistered is a free log retrieval operation binding the contract event 0xde576c51e7828c269f7a259c68554d25364b596a7bd816f01d9b8cdb52e88d43.
//
// Solidity: event WorkerDeregistered(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerDeregistered(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerDeregisteredIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerDeregistered", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerDeregisteredIterator{contract: _WorkerRegistry.contract, event: "WorkerDeregistered", logs: logs, sub: sub}, nil
}

// WatchWorkerDeregistered is a free log subscription operation binding the contract event 0xde576c51e7828c269f7a259c68554d25364b596a7bd816f01d9b8cdb52e88d43.
//
// Solidity: event WorkerDeregistered(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerDeregistered(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerDeregistered, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerDeregistered", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerDeregistered)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerDeregistered", log); err != nil {
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

// ParseWorkerDeregistered is a log parse operation binding the contract event 0xde576c51e7828c269f7a259c68554d25364b596a7bd816f01d9b8cdb52e88d43.
//
// Solidity: event WorkerDeregistered(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerDeregistered(log types.Log) (*WorkerRegistryWorkerDeregistered, error) {
	event := new(WorkerRegistryWorkerDeregistered)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerRegisteredIterator is returned from FilterWorkerRegistered and is used to iterate over the raw logs and unpacked data for WorkerRegistered events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerRegisteredIterator struct {
	Event *WorkerRegistryWorkerRegistered // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerRegistered)
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
		it.Event = new(WorkerRegistryWorkerRegistered)
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
func (it *WorkerRegistryWorkerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerRegistered represents a WorkerRegistered event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerRegistered struct {
	Worker           common.Address
	EncryptionPubKey []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWorkerRegistered is a free log retrieval operation binding the contract event 0x27987c0173113d0f969d0abbf00a8c583fd7f7f44c05af3739f808d2a0afba6f.
//
// Solidity: event WorkerRegistered(address indexed worker, bytes encryptionPubKey)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerRegistered(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerRegisteredIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerRegistered", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerRegisteredIterator{contract: _WorkerRegistry.contract, event: "WorkerRegistered", logs: logs, sub: sub}, nil
}

// WatchWorkerRegistered is a free log subscription operation binding the contract event 0x27987c0173113d0f969d0abbf00a8c583fd7f7f44c05af3739f808d2a0afba6f.
//
// Solidity: event WorkerRegistered(address indexed worker, bytes encryptionPubKey)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerRegistered(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerRegistered, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerRegistered", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerRegistered)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerRegistered", log); err != nil {
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

// ParseWorkerRegistered is a log parse operation binding the contract event 0x27987c0173113d0f969d0abbf00a8c583fd7f7f44c05af3739f808d2a0afba6f.
//
// Solidity: event WorkerRegistered(address indexed worker, bytes encryptionPubKey)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerRegistered(log types.Log) (*WorkerRegistryWorkerRegistered, error) {
	event := new(WorkerRegistryWorkerRegistered)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerReinstatedIterator is returned from FilterWorkerReinstated and is used to iterate over the raw logs and unpacked data for WorkerReinstated events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerReinstatedIterator struct {
	Event *WorkerRegistryWorkerReinstated // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerReinstatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerReinstated)
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
		it.Event = new(WorkerRegistryWorkerReinstated)
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
func (it *WorkerRegistryWorkerReinstatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerReinstatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerReinstated represents a WorkerReinstated event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerReinstated struct {
	Worker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerReinstated is a free log retrieval operation binding the contract event 0x3911afffa47aba3f3a50cb8dd533823d2bbf8c5c44d6dbba306d2cc26da9e584.
//
// Solidity: event WorkerReinstated(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerReinstated(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerReinstatedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerReinstated", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerReinstatedIterator{contract: _WorkerRegistry.contract, event: "WorkerReinstated", logs: logs, sub: sub}, nil
}

// WatchWorkerReinstated is a free log subscription operation binding the contract event 0x3911afffa47aba3f3a50cb8dd533823d2bbf8c5c44d6dbba306d2cc26da9e584.
//
// Solidity: event WorkerReinstated(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerReinstated(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerReinstated, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerReinstated", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerReinstated)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerReinstated", log); err != nil {
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

// ParseWorkerReinstated is a log parse operation binding the contract event 0x3911afffa47aba3f3a50cb8dd533823d2bbf8c5c44d6dbba306d2cc26da9e584.
//
// Solidity: event WorkerReinstated(address indexed worker)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerReinstated(log types.Log) (*WorkerRegistryWorkerReinstated, error) {
	event := new(WorkerRegistryWorkerReinstated)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerReinstated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerSlashedIterator is returned from FilterWorkerSlashed and is used to iterate over the raw logs and unpacked data for WorkerSlashed events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerSlashedIterator struct {
	Event *WorkerRegistryWorkerSlashed // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerSlashed)
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
		it.Event = new(WorkerRegistryWorkerSlashed)
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
func (it *WorkerRegistryWorkerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerSlashed represents a WorkerSlashed event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerSlashed struct {
	Worker   common.Address
	Amount   *big.Int
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWorkerSlashed is a free log retrieval operation binding the contract event 0x35bade04563d8d0b18e6cdc95a7a0eb5b8d9999ac05e48fbef04a65b6cd5ca13.
//
// Solidity: event WorkerSlashed(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerSlashed(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerSlashedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerSlashed", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerSlashedIterator{contract: _WorkerRegistry.contract, event: "WorkerSlashed", logs: logs, sub: sub}, nil
}

// WatchWorkerSlashed is a free log subscription operation binding the contract event 0x35bade04563d8d0b18e6cdc95a7a0eb5b8d9999ac05e48fbef04a65b6cd5ca13.
//
// Solidity: event WorkerSlashed(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerSlashed(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerSlashed, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerSlashed", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerSlashed)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerSlashed", log); err != nil {
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

// ParseWorkerSlashed is a log parse operation binding the contract event 0x35bade04563d8d0b18e6cdc95a7a0eb5b8d9999ac05e48fbef04a65b6cd5ca13.
//
// Solidity: event WorkerSlashed(address indexed worker, uint256 amount, uint256 newStake)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerSlashed(log types.Log) (*WorkerRegistryWorkerSlashed, error) {
	event := new(WorkerRegistryWorkerSlashed)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerRegistryWorkerSuspendedIterator is returned from FilterWorkerSuspended and is used to iterate over the raw logs and unpacked data for WorkerSuspended events raised by the WorkerRegistry contract.
type WorkerRegistryWorkerSuspendedIterator struct {
	Event *WorkerRegistryWorkerSuspended // Event containing the contract specifics and raw log

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
func (it *WorkerRegistryWorkerSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerRegistryWorkerSuspended)
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
		it.Event = new(WorkerRegistryWorkerSuspended)
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
func (it *WorkerRegistryWorkerSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerRegistryWorkerSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerRegistryWorkerSuspended represents a WorkerSuspended event raised by the WorkerRegistry contract.
type WorkerRegistryWorkerSuspended struct {
	Worker common.Address
	Until  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerSuspended is a free log retrieval operation binding the contract event 0x559ba88601c276790cdca8fd3c11feb153784e2cb5785e255374ec6b879eb26f.
//
// Solidity: event WorkerSuspended(address indexed worker, uint256 until)
func (_WorkerRegistry *WorkerRegistryFilterer) FilterWorkerSuspended(opts *bind.FilterOpts, worker []common.Address) (*WorkerRegistryWorkerSuspendedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.FilterLogs(opts, "WorkerSuspended", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerRegistryWorkerSuspendedIterator{contract: _WorkerRegistry.contract, event: "WorkerSuspended", logs: logs, sub: sub}, nil
}

// WatchWorkerSuspended is a free log subscription operation binding the contract event 0x559ba88601c276790cdca8fd3c11feb153784e2cb5785e255374ec6b879eb26f.
//
// Solidity: event WorkerSuspended(address indexed worker, uint256 until)
func (_WorkerRegistry *WorkerRegistryFilterer) WatchWorkerSuspended(opts *bind.WatchOpts, sink chan<- *WorkerRegistryWorkerSuspended, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerRegistry.contract.WatchLogs(opts, "WorkerSuspended", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerRegistryWorkerSuspended)
				if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerSuspended", log); err != nil {
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

// ParseWorkerSuspended is a log parse operation binding the contract event 0x559ba88601c276790cdca8fd3c11feb153784e2cb5785e255374ec6b879eb26f.
//
// Solidity: event WorkerSuspended(address indexed worker, uint256 until)
func (_WorkerRegistry *WorkerRegistryFilterer) ParseWorkerSuspended(log types.Log) (*WorkerRegistryWorkerSuspended, error) {
	event := new(WorkerRegistryWorkerSuspended)
	if err := _WorkerRegistry.contract.UnpackLog(event, "WorkerSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
