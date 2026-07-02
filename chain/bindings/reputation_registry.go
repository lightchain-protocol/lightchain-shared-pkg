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

// IReputationRegistryWorkerReputation is an auto generated low-level Go binding around an user-defined struct.
type IReputationRegistryWorkerReputation struct {
	Score        uint64
	Completions  uint32
	DisputesLost uint32
	Timeouts     uint32
	Initialized  bool
}

// ReputationRegistryMetaData contains all meta data concerning the ReputationRegistry contract.
var ReputationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_SCORE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"START_SCORE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReputation\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReputationRegistry.WorkerReputation\",\"components\":[{\"name\":\"score\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"completions\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disputesLost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeouts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getScore\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_jobRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"jobRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordCompletion\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordDisputeLoss\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordDisputeWin\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordTimeout\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"completionReward\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disputeWinReward\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timeoutPenalty\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disputeLossPenalty\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReputationParamsUpdated\",\"inputs\":[{\"name\":\"completionReward\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"disputeWinReward\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutPenalty\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"disputeLossPenalty\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReputationUpdated\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newScore\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ReputationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationRegistryMetaData.ABI instead.
var ReputationRegistryABI = ReputationRegistryMetaData.ABI

// ReputationRegistry is an auto generated Go binding around an Ethereum contract.
type ReputationRegistry struct {
	ReputationRegistryCaller     // Read-only binding to the contract
	ReputationRegistryTransactor // Write-only binding to the contract
	ReputationRegistryFilterer   // Log filterer for contract events
}

// ReputationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationRegistrySession struct {
	Contract     *ReputationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReputationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationRegistryCallerSession struct {
	Contract *ReputationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ReputationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationRegistryTransactorSession struct {
	Contract     *ReputationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ReputationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRegistryRaw struct {
	Contract *ReputationRegistry // Generic contract binding to access the raw methods on
}

// ReputationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationRegistryCallerRaw struct {
	Contract *ReputationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationRegistryTransactorRaw struct {
	Contract *ReputationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationRegistry creates a new instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistry(address common.Address, backend bind.ContractBackend) (*ReputationRegistry, error) {
	contract, err := bindReputationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistry{ReputationRegistryCaller: ReputationRegistryCaller{contract: contract}, ReputationRegistryTransactor: ReputationRegistryTransactor{contract: contract}, ReputationRegistryFilterer: ReputationRegistryFilterer{contract: contract}}, nil
}

// NewReputationRegistryCaller creates a new read-only instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryCaller(address common.Address, caller bind.ContractCaller) (*ReputationRegistryCaller, error) {
	contract, err := bindReputationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryCaller{contract: contract}, nil
}

// NewReputationRegistryTransactor creates a new write-only instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationRegistryTransactor, error) {
	contract, err := bindReputationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryTransactor{contract: contract}, nil
}

// NewReputationRegistryFilterer creates a new log filterer instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationRegistryFilterer, error) {
	contract, err := bindReputationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryFilterer{contract: contract}, nil
}

// bindReputationRegistry binds a generic wrapper to an already deployed contract.
func bindReputationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReputationRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationRegistry *ReputationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationRegistry.Contract.ReputationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationRegistry *ReputationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.ReputationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationRegistry *ReputationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.ReputationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationRegistry *ReputationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationRegistry *ReputationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationRegistry *ReputationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXSCORE is a free data retrieval call binding the contract method 0x27ff6223.
//
// Solidity: function MAX_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistryCaller) MAXSCORE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "MAX_SCORE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXSCORE is a free data retrieval call binding the contract method 0x27ff6223.
//
// Solidity: function MAX_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistrySession) MAXSCORE() (uint64, error) {
	return _ReputationRegistry.Contract.MAXSCORE(&_ReputationRegistry.CallOpts)
}

// MAXSCORE is a free data retrieval call binding the contract method 0x27ff6223.
//
// Solidity: function MAX_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistryCallerSession) MAXSCORE() (uint64, error) {
	return _ReputationRegistry.Contract.MAXSCORE(&_ReputationRegistry.CallOpts)
}

// STARTSCORE is a free data retrieval call binding the contract method 0x8acf2773.
//
// Solidity: function START_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistryCaller) STARTSCORE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "START_SCORE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// STARTSCORE is a free data retrieval call binding the contract method 0x8acf2773.
//
// Solidity: function START_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistrySession) STARTSCORE() (uint64, error) {
	return _ReputationRegistry.Contract.STARTSCORE(&_ReputationRegistry.CallOpts)
}

// STARTSCORE is a free data retrieval call binding the contract method 0x8acf2773.
//
// Solidity: function START_SCORE() view returns(uint64)
func (_ReputationRegistry *ReputationRegistryCallerSession) STARTSCORE() (uint64, error) {
	return _ReputationRegistry.Contract.STARTSCORE(&_ReputationRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ReputationRegistry *ReputationRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ReputationRegistry *ReputationRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ReputationRegistry.Contract.UPGRADEINTERFACEVERSION(&_ReputationRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ReputationRegistry *ReputationRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ReputationRegistry.Contract.UPGRADEINTERFACEVERSION(&_ReputationRegistry.CallOpts)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address worker) view returns((uint64,uint32,uint32,uint32,bool))
func (_ReputationRegistry *ReputationRegistryCaller) GetReputation(opts *bind.CallOpts, worker common.Address) (IReputationRegistryWorkerReputation, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getReputation", worker)

	if err != nil {
		return *new(IReputationRegistryWorkerReputation), err
	}

	out0 := *abi.ConvertType(out[0], new(IReputationRegistryWorkerReputation)).(*IReputationRegistryWorkerReputation)

	return out0, err

}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address worker) view returns((uint64,uint32,uint32,uint32,bool))
func (_ReputationRegistry *ReputationRegistrySession) GetReputation(worker common.Address) (IReputationRegistryWorkerReputation, error) {
	return _ReputationRegistry.Contract.GetReputation(&_ReputationRegistry.CallOpts, worker)
}

// GetReputation is a free data retrieval call binding the contract method 0x9c89a0e2.
//
// Solidity: function getReputation(address worker) view returns((uint64,uint32,uint32,uint32,bool))
func (_ReputationRegistry *ReputationRegistryCallerSession) GetReputation(worker common.Address) (IReputationRegistryWorkerReputation, error) {
	return _ReputationRegistry.Contract.GetReputation(&_ReputationRegistry.CallOpts, worker)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address worker) view returns(uint256)
func (_ReputationRegistry *ReputationRegistryCaller) GetScore(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getScore", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address worker) view returns(uint256)
func (_ReputationRegistry *ReputationRegistrySession) GetScore(worker common.Address) (*big.Int, error) {
	return _ReputationRegistry.Contract.GetScore(&_ReputationRegistry.CallOpts, worker)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address worker) view returns(uint256)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetScore(worker common.Address) (*big.Int, error) {
	return _ReputationRegistry.Contract.GetScore(&_ReputationRegistry.CallOpts, worker)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistryCaller) JobRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "jobRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistrySession) JobRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.JobRegistry(&_ReputationRegistry.CallOpts)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistryCallerSession) JobRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.JobRegistry(&_ReputationRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationRegistry *ReputationRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationRegistry *ReputationRegistrySession) Owner() (common.Address, error) {
	return _ReputationRegistry.Contract.Owner(&_ReputationRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationRegistry *ReputationRegistryCallerSession) Owner() (common.Address, error) {
	return _ReputationRegistry.Contract.Owner(&_ReputationRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ReputationRegistry *ReputationRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ReputationRegistry *ReputationRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ReputationRegistry.Contract.ProxiableUUID(&_ReputationRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ReputationRegistry *ReputationRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ReputationRegistry.Contract.ProxiableUUID(&_ReputationRegistry.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _jobRegistry) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _jobRegistry common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "initialize", _initialOwner, _jobRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _jobRegistry) returns()
func (_ReputationRegistry *ReputationRegistrySession) Initialize(_initialOwner common.Address, _jobRegistry common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.Initialize(&_ReputationRegistry.TransactOpts, _initialOwner, _jobRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _jobRegistry) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) Initialize(_initialOwner common.Address, _jobRegistry common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.Initialize(&_ReputationRegistry.TransactOpts, _initialOwner, _jobRegistry)
}

// RecordCompletion is a paid mutator transaction binding the contract method 0x8a9e9dec.
//
// Solidity: function recordCompletion(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RecordCompletion(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "recordCompletion", worker)
}

// RecordCompletion is a paid mutator transaction binding the contract method 0x8a9e9dec.
//
// Solidity: function recordCompletion(address worker) returns()
func (_ReputationRegistry *ReputationRegistrySession) RecordCompletion(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordCompletion(&_ReputationRegistry.TransactOpts, worker)
}

// RecordCompletion is a paid mutator transaction binding the contract method 0x8a9e9dec.
//
// Solidity: function recordCompletion(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RecordCompletion(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordCompletion(&_ReputationRegistry.TransactOpts, worker)
}

// RecordDisputeLoss is a paid mutator transaction binding the contract method 0x9325ae86.
//
// Solidity: function recordDisputeLoss(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RecordDisputeLoss(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "recordDisputeLoss", worker)
}

// RecordDisputeLoss is a paid mutator transaction binding the contract method 0x9325ae86.
//
// Solidity: function recordDisputeLoss(address worker) returns()
func (_ReputationRegistry *ReputationRegistrySession) RecordDisputeLoss(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordDisputeLoss(&_ReputationRegistry.TransactOpts, worker)
}

// RecordDisputeLoss is a paid mutator transaction binding the contract method 0x9325ae86.
//
// Solidity: function recordDisputeLoss(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RecordDisputeLoss(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordDisputeLoss(&_ReputationRegistry.TransactOpts, worker)
}

// RecordDisputeWin is a paid mutator transaction binding the contract method 0x47f24036.
//
// Solidity: function recordDisputeWin(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RecordDisputeWin(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "recordDisputeWin", worker)
}

// RecordDisputeWin is a paid mutator transaction binding the contract method 0x47f24036.
//
// Solidity: function recordDisputeWin(address worker) returns()
func (_ReputationRegistry *ReputationRegistrySession) RecordDisputeWin(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordDisputeWin(&_ReputationRegistry.TransactOpts, worker)
}

// RecordDisputeWin is a paid mutator transaction binding the contract method 0x47f24036.
//
// Solidity: function recordDisputeWin(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RecordDisputeWin(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordDisputeWin(&_ReputationRegistry.TransactOpts, worker)
}

// RecordTimeout is a paid mutator transaction binding the contract method 0x0260f984.
//
// Solidity: function recordTimeout(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RecordTimeout(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "recordTimeout", worker)
}

// RecordTimeout is a paid mutator transaction binding the contract method 0x0260f984.
//
// Solidity: function recordTimeout(address worker) returns()
func (_ReputationRegistry *ReputationRegistrySession) RecordTimeout(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordTimeout(&_ReputationRegistry.TransactOpts, worker)
}

// RecordTimeout is a paid mutator transaction binding the contract method 0x0260f984.
//
// Solidity: function recordTimeout(address worker) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RecordTimeout(worker common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RecordTimeout(&_ReputationRegistry.TransactOpts, worker)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReputationRegistry *ReputationRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RenounceOwnership(&_ReputationRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RenounceOwnership(&_ReputationRegistry.TransactOpts)
}

// SetParams is a paid mutator transaction binding the contract method 0x6f3931f4.
//
// Solidity: function setParams(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) SetParams(opts *bind.TransactOpts, completionReward uint64, disputeWinReward uint64, timeoutPenalty uint64, disputeLossPenalty uint64) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "setParams", completionReward, disputeWinReward, timeoutPenalty, disputeLossPenalty)
}

// SetParams is a paid mutator transaction binding the contract method 0x6f3931f4.
//
// Solidity: function setParams(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty) returns()
func (_ReputationRegistry *ReputationRegistrySession) SetParams(completionReward uint64, disputeWinReward uint64, timeoutPenalty uint64, disputeLossPenalty uint64) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.SetParams(&_ReputationRegistry.TransactOpts, completionReward, disputeWinReward, timeoutPenalty, disputeLossPenalty)
}

// SetParams is a paid mutator transaction binding the contract method 0x6f3931f4.
//
// Solidity: function setParams(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) SetParams(completionReward uint64, disputeWinReward uint64, timeoutPenalty uint64, disputeLossPenalty uint64) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.SetParams(&_ReputationRegistry.TransactOpts, completionReward, disputeWinReward, timeoutPenalty, disputeLossPenalty)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReputationRegistry *ReputationRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.TransferOwnership(&_ReputationRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.TransferOwnership(&_ReputationRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ReputationRegistry *ReputationRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ReputationRegistry *ReputationRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.UpgradeToAndCall(&_ReputationRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.UpgradeToAndCall(&_ReputationRegistry.TransactOpts, newImplementation, data)
}

// ReputationRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReputationRegistry contract.
type ReputationRegistryInitializedIterator struct {
	Event *ReputationRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryInitialized)
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
		it.Event = new(ReputationRegistryInitialized)
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
func (it *ReputationRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryInitialized represents a Initialized event raised by the ReputationRegistry contract.
type ReputationRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReputationRegistryInitializedIterator, error) {

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryInitializedIterator{contract: _ReputationRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReputationRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryInitialized)
				if err := _ReputationRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ReputationRegistry *ReputationRegistryFilterer) ParseInitialized(log types.Log) (*ReputationRegistryInitialized, error) {
	event := new(ReputationRegistryInitialized)
	if err := _ReputationRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReputationRegistry contract.
type ReputationRegistryOwnershipTransferredIterator struct {
	Event *ReputationRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryOwnershipTransferred)
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
		it.Event = new(ReputationRegistryOwnershipTransferred)
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
func (it *ReputationRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ReputationRegistry contract.
type ReputationRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReputationRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryOwnershipTransferredIterator{contract: _ReputationRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReputationRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryOwnershipTransferred)
				if err := _ReputationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ReputationRegistry *ReputationRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ReputationRegistryOwnershipTransferred, error) {
	event := new(ReputationRegistryOwnershipTransferred)
	if err := _ReputationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryReputationParamsUpdatedIterator is returned from FilterReputationParamsUpdated and is used to iterate over the raw logs and unpacked data for ReputationParamsUpdated events raised by the ReputationRegistry contract.
type ReputationRegistryReputationParamsUpdatedIterator struct {
	Event *ReputationRegistryReputationParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryReputationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryReputationParamsUpdated)
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
		it.Event = new(ReputationRegistryReputationParamsUpdated)
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
func (it *ReputationRegistryReputationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryReputationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryReputationParamsUpdated represents a ReputationParamsUpdated event raised by the ReputationRegistry contract.
type ReputationRegistryReputationParamsUpdated struct {
	CompletionReward   uint64
	DisputeWinReward   uint64
	TimeoutPenalty     uint64
	DisputeLossPenalty uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReputationParamsUpdated is a free log retrieval operation binding the contract event 0x187b2829ffd7e1cf6f19f3a2fa25390e7f7ded870c1869ec75e0113b32830f76.
//
// Solidity: event ReputationParamsUpdated(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterReputationParamsUpdated(opts *bind.FilterOpts) (*ReputationRegistryReputationParamsUpdatedIterator, error) {

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "ReputationParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryReputationParamsUpdatedIterator{contract: _ReputationRegistry.contract, event: "ReputationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationParamsUpdated is a free log subscription operation binding the contract event 0x187b2829ffd7e1cf6f19f3a2fa25390e7f7ded870c1869ec75e0113b32830f76.
//
// Solidity: event ReputationParamsUpdated(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchReputationParamsUpdated(opts *bind.WatchOpts, sink chan<- *ReputationRegistryReputationParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "ReputationParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryReputationParamsUpdated)
				if err := _ReputationRegistry.contract.UnpackLog(event, "ReputationParamsUpdated", log); err != nil {
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

// ParseReputationParamsUpdated is a log parse operation binding the contract event 0x187b2829ffd7e1cf6f19f3a2fa25390e7f7ded870c1869ec75e0113b32830f76.
//
// Solidity: event ReputationParamsUpdated(uint64 completionReward, uint64 disputeWinReward, uint64 timeoutPenalty, uint64 disputeLossPenalty)
func (_ReputationRegistry *ReputationRegistryFilterer) ParseReputationParamsUpdated(log types.Log) (*ReputationRegistryReputationParamsUpdated, error) {
	event := new(ReputationRegistryReputationParamsUpdated)
	if err := _ReputationRegistry.contract.UnpackLog(event, "ReputationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryReputationUpdatedIterator is returned from FilterReputationUpdated and is used to iterate over the raw logs and unpacked data for ReputationUpdated events raised by the ReputationRegistry contract.
type ReputationRegistryReputationUpdatedIterator struct {
	Event *ReputationRegistryReputationUpdated // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryReputationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryReputationUpdated)
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
		it.Event = new(ReputationRegistryReputationUpdated)
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
func (it *ReputationRegistryReputationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryReputationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryReputationUpdated represents a ReputationUpdated event raised by the ReputationRegistry contract.
type ReputationRegistryReputationUpdated struct {
	Worker   common.Address
	NewScore uint64
	Reason   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReputationUpdated is a free log retrieval operation binding the contract event 0x43cdc02647bdff478eaf554aeba33fc3bac837c8e43f38276606eab280d0e7f5.
//
// Solidity: event ReputationUpdated(address indexed worker, uint64 newScore, uint8 reason)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterReputationUpdated(opts *bind.FilterOpts, worker []common.Address) (*ReputationRegistryReputationUpdatedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "ReputationUpdated", workerRule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryReputationUpdatedIterator{contract: _ReputationRegistry.contract, event: "ReputationUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationUpdated is a free log subscription operation binding the contract event 0x43cdc02647bdff478eaf554aeba33fc3bac837c8e43f38276606eab280d0e7f5.
//
// Solidity: event ReputationUpdated(address indexed worker, uint64 newScore, uint8 reason)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchReputationUpdated(opts *bind.WatchOpts, sink chan<- *ReputationRegistryReputationUpdated, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "ReputationUpdated", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryReputationUpdated)
				if err := _ReputationRegistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
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

// ParseReputationUpdated is a log parse operation binding the contract event 0x43cdc02647bdff478eaf554aeba33fc3bac837c8e43f38276606eab280d0e7f5.
//
// Solidity: event ReputationUpdated(address indexed worker, uint64 newScore, uint8 reason)
func (_ReputationRegistry *ReputationRegistryFilterer) ParseReputationUpdated(log types.Log) (*ReputationRegistryReputationUpdated, error) {
	event := new(ReputationRegistryReputationUpdated)
	if err := _ReputationRegistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ReputationRegistry contract.
type ReputationRegistryUpgradedIterator struct {
	Event *ReputationRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryUpgraded)
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
		it.Event = new(ReputationRegistryUpgraded)
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
func (it *ReputationRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryUpgraded represents a Upgraded event raised by the ReputationRegistry contract.
type ReputationRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ReputationRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryUpgradedIterator{contract: _ReputationRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ReputationRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryUpgraded)
				if err := _ReputationRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ReputationRegistry *ReputationRegistryFilterer) ParseUpgraded(log types.Log) (*ReputationRegistryUpgraded, error) {
	event := new(ReputationRegistryUpgraded)
	if err := _ReputationRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
