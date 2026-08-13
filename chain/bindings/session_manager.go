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

// ISessionManagerRequest is an auto generated low-level Go binding around an user-defined struct.
type ISessionManagerRequest struct {
	User         common.Address
	ModelId      [32]byte
	RequestBlock uint64
	Expiry       uint64
	Worker       common.Address
	Status       uint8
}

// SessionManagerMetaData contains all meta data concerning the SessionManager contract.
var SessionManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSession\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eligibleNow\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISessionManager.Request\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISessionManager.ReqStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCapabilities\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workerRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aiConfig_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reputation_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideSessionKey\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSessionKeyFor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSession\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionFor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionForWithCapabilities\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionWithCapabilities\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionClaimed\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionReady\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionRequested\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// SessionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionManagerMetaData.ABI instead.
var SessionManagerABI = SessionManagerMetaData.ABI

// SessionManager is an auto generated Go binding around an Ethereum contract.
type SessionManager struct {
	SessionManagerCaller     // Read-only binding to the contract
	SessionManagerTransactor // Write-only binding to the contract
	SessionManagerFilterer   // Log filterer for contract events
}

// SessionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionManagerSession struct {
	Contract     *SessionManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SessionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionManagerCallerSession struct {
	Contract *SessionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SessionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionManagerTransactorSession struct {
	Contract     *SessionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SessionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionManagerRaw struct {
	Contract *SessionManager // Generic contract binding to access the raw methods on
}

// SessionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionManagerCallerRaw struct {
	Contract *SessionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SessionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionManagerTransactorRaw struct {
	Contract *SessionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionManager creates a new instance of SessionManager, bound to a specific deployed contract.
func NewSessionManager(address common.Address, backend bind.ContractBackend) (*SessionManager, error) {
	contract, err := bindSessionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionManager{SessionManagerCaller: SessionManagerCaller{contract: contract}, SessionManagerTransactor: SessionManagerTransactor{contract: contract}, SessionManagerFilterer: SessionManagerFilterer{contract: contract}}, nil
}

// NewSessionManagerCaller creates a new read-only instance of SessionManager, bound to a specific deployed contract.
func NewSessionManagerCaller(address common.Address, caller bind.ContractCaller) (*SessionManagerCaller, error) {
	contract, err := bindSessionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionManagerCaller{contract: contract}, nil
}

// NewSessionManagerTransactor creates a new write-only instance of SessionManager, bound to a specific deployed contract.
func NewSessionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionManagerTransactor, error) {
	contract, err := bindSessionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionManagerTransactor{contract: contract}, nil
}

// NewSessionManagerFilterer creates a new log filterer instance of SessionManager, bound to a specific deployed contract.
func NewSessionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionManagerFilterer, error) {
	contract, err := bindSessionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionManagerFilterer{contract: contract}, nil
}

// bindSessionManager binds a generic wrapper to an already deployed contract.
func bindSessionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionManager *SessionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionManager.Contract.SessionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionManager *SessionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionManager.Contract.SessionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionManager *SessionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionManager.Contract.SessionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionManager *SessionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionManager *SessionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionManager *SessionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionManager *SessionManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionManager *SessionManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SessionManager.Contract.UPGRADEINTERFACEVERSION(&_SessionManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionManager *SessionManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SessionManager.Contract.UPGRADEINTERFACEVERSION(&_SessionManager.CallOpts)
}

// EligibleNow is a free data retrieval call binding the contract method 0xa5c6f7e8.
//
// Solidity: function eligibleNow(uint256 reqId, address worker) view returns(bool)
func (_SessionManager *SessionManagerCaller) EligibleNow(opts *bind.CallOpts, reqId *big.Int, worker common.Address) (bool, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "eligibleNow", reqId, worker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EligibleNow is a free data retrieval call binding the contract method 0xa5c6f7e8.
//
// Solidity: function eligibleNow(uint256 reqId, address worker) view returns(bool)
func (_SessionManager *SessionManagerSession) EligibleNow(reqId *big.Int, worker common.Address) (bool, error) {
	return _SessionManager.Contract.EligibleNow(&_SessionManager.CallOpts, reqId, worker)
}

// EligibleNow is a free data retrieval call binding the contract method 0xa5c6f7e8.
//
// Solidity: function eligibleNow(uint256 reqId, address worker) view returns(bool)
func (_SessionManager *SessionManagerCallerSession) EligibleNow(reqId *big.Int, worker common.Address) (bool, error) {
	return _SessionManager.Contract.EligibleNow(&_SessionManager.CallOpts, reqId, worker)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 reqId) view returns((address,bytes32,uint64,uint64,address,uint8))
func (_SessionManager *SessionManagerCaller) GetRequest(opts *bind.CallOpts, reqId *big.Int) (ISessionManagerRequest, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getRequest", reqId)

	if err != nil {
		return *new(ISessionManagerRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(ISessionManagerRequest)).(*ISessionManagerRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 reqId) view returns((address,bytes32,uint64,uint64,address,uint8))
func (_SessionManager *SessionManagerSession) GetRequest(reqId *big.Int) (ISessionManagerRequest, error) {
	return _SessionManager.Contract.GetRequest(&_SessionManager.CallOpts, reqId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 reqId) view returns((address,bytes32,uint64,uint64,address,uint8))
func (_SessionManager *SessionManagerCallerSession) GetRequest(reqId *big.Int) (ISessionManagerRequest, error) {
	return _SessionManager.Contract.GetRequest(&_SessionManager.CallOpts, reqId)
}

// GetRequiredCapabilities is a free data retrieval call binding the contract method 0xa5f56fd6.
//
// Solidity: function getRequiredCapabilities(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetRequiredCapabilities(opts *bind.CallOpts, reqId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getRequiredCapabilities", reqId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredCapabilities is a free data retrieval call binding the contract method 0xa5f56fd6.
//
// Solidity: function getRequiredCapabilities(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerSession) GetRequiredCapabilities(reqId *big.Int) (*big.Int, error) {
	return _SessionManager.Contract.GetRequiredCapabilities(&_SessionManager.CallOpts, reqId)
}

// GetRequiredCapabilities is a free data retrieval call binding the contract method 0xa5f56fd6.
//
// Solidity: function getRequiredCapabilities(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetRequiredCapabilities(reqId *big.Int) (*big.Int, error) {
	return _SessionManager.Contract.GetRequiredCapabilities(&_SessionManager.CallOpts, reqId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionManager *SessionManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionManager *SessionManagerSession) Owner() (common.Address, error) {
	return _SessionManager.Contract.Owner(&_SessionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionManager *SessionManagerCallerSession) Owner() (common.Address, error) {
	return _SessionManager.Contract.Owner(&_SessionManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionManager *SessionManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionManager *SessionManagerSession) ProxiableUUID() ([32]byte, error) {
	return _SessionManager.Contract.ProxiableUUID(&_SessionManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionManager *SessionManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SessionManager.Contract.ProxiableUUID(&_SessionManager.CallOpts)
}

// ClaimSession is a paid mutator transaction binding the contract method 0xf446092d.
//
// Solidity: function claimSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactor) ClaimSession(opts *bind.TransactOpts, reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "claimSession", reqId)
}

// ClaimSession is a paid mutator transaction binding the contract method 0xf446092d.
//
// Solidity: function claimSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerSession) ClaimSession(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.ClaimSession(&_SessionManager.TransactOpts, reqId)
}

// ClaimSession is a paid mutator transaction binding the contract method 0xf446092d.
//
// Solidity: function claimSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactorSession) ClaimSession(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.ClaimSession(&_SessionManager.TransactOpts, reqId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address initialOwner, address jobRegistry_, address workerRegistry_, address aiConfig_, address reputation_) returns()
func (_SessionManager *SessionManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, jobRegistry_ common.Address, workerRegistry_ common.Address, aiConfig_ common.Address, reputation_ common.Address) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "initialize", initialOwner, jobRegistry_, workerRegistry_, aiConfig_, reputation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address initialOwner, address jobRegistry_, address workerRegistry_, address aiConfig_, address reputation_) returns()
func (_SessionManager *SessionManagerSession) Initialize(initialOwner common.Address, jobRegistry_ common.Address, workerRegistry_ common.Address, aiConfig_ common.Address, reputation_ common.Address) (*types.Transaction, error) {
	return _SessionManager.Contract.Initialize(&_SessionManager.TransactOpts, initialOwner, jobRegistry_, workerRegistry_, aiConfig_, reputation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address initialOwner, address jobRegistry_, address workerRegistry_, address aiConfig_, address reputation_) returns()
func (_SessionManager *SessionManagerTransactorSession) Initialize(initialOwner common.Address, jobRegistry_ common.Address, workerRegistry_ common.Address, aiConfig_ common.Address, reputation_ common.Address) (*types.Transaction, error) {
	return _SessionManager.Contract.Initialize(&_SessionManager.TransactOpts, initialOwner, jobRegistry_, workerRegistry_, aiConfig_, reputation_)
}

// ProvideSessionKey is a paid mutator transaction binding the contract method 0x5ed55de7.
//
// Solidity: function provideSessionKey(uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerTransactor) ProvideSessionKey(opts *bind.TransactOpts, reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "provideSessionKey", reqId, encWorkerKey, encDisputerKey)
}

// ProvideSessionKey is a paid mutator transaction binding the contract method 0x5ed55de7.
//
// Solidity: function provideSessionKey(uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerSession) ProvideSessionKey(reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.ProvideSessionKey(&_SessionManager.TransactOpts, reqId, encWorkerKey, encDisputerKey)
}

// ProvideSessionKey is a paid mutator transaction binding the contract method 0x5ed55de7.
//
// Solidity: function provideSessionKey(uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerTransactorSession) ProvideSessionKey(reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.ProvideSessionKey(&_SessionManager.TransactOpts, reqId, encWorkerKey, encDisputerKey)
}

// ProvideSessionKeyFor is a paid mutator transaction binding the contract method 0xa468b41c.
//
// Solidity: function provideSessionKeyFor(address user, uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerTransactor) ProvideSessionKeyFor(opts *bind.TransactOpts, user common.Address, reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "provideSessionKeyFor", user, reqId, encWorkerKey, encDisputerKey)
}

// ProvideSessionKeyFor is a paid mutator transaction binding the contract method 0xa468b41c.
//
// Solidity: function provideSessionKeyFor(address user, uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerSession) ProvideSessionKeyFor(user common.Address, reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.ProvideSessionKeyFor(&_SessionManager.TransactOpts, user, reqId, encWorkerKey, encDisputerKey)
}

// ProvideSessionKeyFor is a paid mutator transaction binding the contract method 0xa468b41c.
//
// Solidity: function provideSessionKeyFor(address user, uint256 reqId, bytes encWorkerKey, bytes encDisputerKey) returns()
func (_SessionManager *SessionManagerTransactorSession) ProvideSessionKeyFor(user common.Address, reqId *big.Int, encWorkerKey []byte, encDisputerKey []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.ProvideSessionKeyFor(&_SessionManager.TransactOpts, user, reqId, encWorkerKey, encDisputerKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionManager *SessionManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionManager *SessionManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _SessionManager.Contract.RenounceOwnership(&_SessionManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionManager *SessionManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SessionManager.Contract.RenounceOwnership(&_SessionManager.TransactOpts)
}

// RequestSession is a paid mutator transaction binding the contract method 0x3686764b.
//
// Solidity: function requestSession(bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSession(opts *bind.TransactOpts, modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSession", modelId, expiry)
}

// RequestSession is a paid mutator transaction binding the contract method 0x3686764b.
//
// Solidity: function requestSession(bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSession(modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSession(&_SessionManager.TransactOpts, modelId, expiry)
}

// RequestSession is a paid mutator transaction binding the contract method 0x3686764b.
//
// Solidity: function requestSession(bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSession(modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSession(&_SessionManager.TransactOpts, modelId, expiry)
}

// RequestSessionFor is a paid mutator transaction binding the contract method 0x4cf70891.
//
// Solidity: function requestSessionFor(address user, bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSessionFor(opts *bind.TransactOpts, user common.Address, modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSessionFor", user, modelId, expiry)
}

// RequestSessionFor is a paid mutator transaction binding the contract method 0x4cf70891.
//
// Solidity: function requestSessionFor(address user, bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSessionFor(user common.Address, modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionFor(&_SessionManager.TransactOpts, user, modelId, expiry)
}

// RequestSessionFor is a paid mutator transaction binding the contract method 0x4cf70891.
//
// Solidity: function requestSessionFor(address user, bytes32 modelId, uint256 expiry) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSessionFor(user common.Address, modelId [32]byte, expiry *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionFor(&_SessionManager.TransactOpts, user, modelId, expiry)
}

// RequestSessionForWithCapabilities is a paid mutator transaction binding the contract method 0xb5a0f052.
//
// Solidity: function requestSessionForWithCapabilities(address user, bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSessionForWithCapabilities(opts *bind.TransactOpts, user common.Address, modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSessionForWithCapabilities", user, modelId, expiry, requiredCaps)
}

// RequestSessionForWithCapabilities is a paid mutator transaction binding the contract method 0xb5a0f052.
//
// Solidity: function requestSessionForWithCapabilities(address user, bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSessionForWithCapabilities(user common.Address, modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionForWithCapabilities(&_SessionManager.TransactOpts, user, modelId, expiry, requiredCaps)
}

// RequestSessionForWithCapabilities is a paid mutator transaction binding the contract method 0xb5a0f052.
//
// Solidity: function requestSessionForWithCapabilities(address user, bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSessionForWithCapabilities(user common.Address, modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionForWithCapabilities(&_SessionManager.TransactOpts, user, modelId, expiry, requiredCaps)
}

// RequestSessionWithCapabilities is a paid mutator transaction binding the contract method 0x46a268c5.
//
// Solidity: function requestSessionWithCapabilities(bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSessionWithCapabilities(opts *bind.TransactOpts, modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSessionWithCapabilities", modelId, expiry, requiredCaps)
}

// RequestSessionWithCapabilities is a paid mutator transaction binding the contract method 0x46a268c5.
//
// Solidity: function requestSessionWithCapabilities(bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSessionWithCapabilities(modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionWithCapabilities(&_SessionManager.TransactOpts, modelId, expiry, requiredCaps)
}

// RequestSessionWithCapabilities is a paid mutator transaction binding the contract method 0x46a268c5.
//
// Solidity: function requestSessionWithCapabilities(bytes32 modelId, uint256 expiry, uint256 requiredCaps) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSessionWithCapabilities(modelId [32]byte, expiry *big.Int, requiredCaps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionWithCapabilities(&_SessionManager.TransactOpts, modelId, expiry, requiredCaps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionManager *SessionManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionManager *SessionManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SessionManager.Contract.TransferOwnership(&_SessionManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionManager *SessionManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SessionManager.Contract.TransferOwnership(&_SessionManager.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionManager *SessionManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionManager *SessionManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.UpgradeToAndCall(&_SessionManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionManager *SessionManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionManager.Contract.UpgradeToAndCall(&_SessionManager.TransactOpts, newImplementation, data)
}

// SessionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SessionManager contract.
type SessionManagerInitializedIterator struct {
	Event *SessionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *SessionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerInitialized)
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
		it.Event = new(SessionManagerInitialized)
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
func (it *SessionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerInitialized represents a Initialized event raised by the SessionManager contract.
type SessionManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SessionManager *SessionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SessionManagerInitializedIterator, error) {

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SessionManagerInitializedIterator{contract: _SessionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SessionManager *SessionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SessionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerInitialized)
				if err := _SessionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SessionManager *SessionManagerFilterer) ParseInitialized(log types.Log) (*SessionManagerInitialized, error) {
	event := new(SessionManagerInitialized)
	if err := _SessionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SessionManager contract.
type SessionManagerOwnershipTransferredIterator struct {
	Event *SessionManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SessionManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerOwnershipTransferred)
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
		it.Event = new(SessionManagerOwnershipTransferred)
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
func (it *SessionManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerOwnershipTransferred represents a OwnershipTransferred event raised by the SessionManager contract.
type SessionManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SessionManager *SessionManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SessionManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerOwnershipTransferredIterator{contract: _SessionManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SessionManager *SessionManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SessionManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerOwnershipTransferred)
				if err := _SessionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SessionManager *SessionManagerFilterer) ParseOwnershipTransferred(log types.Log) (*SessionManagerOwnershipTransferred, error) {
	event := new(SessionManagerOwnershipTransferred)
	if err := _SessionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerSessionClaimedIterator is returned from FilterSessionClaimed and is used to iterate over the raw logs and unpacked data for SessionClaimed events raised by the SessionManager contract.
type SessionManagerSessionClaimedIterator struct {
	Event *SessionManagerSessionClaimed // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionClaimed)
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
		it.Event = new(SessionManagerSessionClaimed)
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
func (it *SessionManagerSessionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionClaimed represents a SessionClaimed event raised by the SessionManager contract.
type SessionManagerSessionClaimed struct {
	ReqId  *big.Int
	Worker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSessionClaimed is a free log retrieval operation binding the contract event 0xa0b6ed9ccbcfe601c67a69f1cde215af082a77c2af5f16fa2c658beb364b56d1.
//
// Solidity: event SessionClaimed(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) FilterSessionClaimed(opts *bind.FilterOpts, reqId []*big.Int, worker []common.Address) (*SessionManagerSessionClaimedIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionClaimed", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionClaimedIterator{contract: _SessionManager.contract, event: "SessionClaimed", logs: logs, sub: sub}, nil
}

// WatchSessionClaimed is a free log subscription operation binding the contract event 0xa0b6ed9ccbcfe601c67a69f1cde215af082a77c2af5f16fa2c658beb364b56d1.
//
// Solidity: event SessionClaimed(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) WatchSessionClaimed(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionClaimed, reqId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionClaimed", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionClaimed)
				if err := _SessionManager.contract.UnpackLog(event, "SessionClaimed", log); err != nil {
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

// ParseSessionClaimed is a log parse operation binding the contract event 0xa0b6ed9ccbcfe601c67a69f1cde215af082a77c2af5f16fa2c658beb364b56d1.
//
// Solidity: event SessionClaimed(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) ParseSessionClaimed(log types.Log) (*SessionManagerSessionClaimed, error) {
	event := new(SessionManagerSessionClaimed)
	if err := _SessionManager.contract.UnpackLog(event, "SessionClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerSessionReadyIterator is returned from FilterSessionReady and is used to iterate over the raw logs and unpacked data for SessionReady events raised by the SessionManager contract.
type SessionManagerSessionReadyIterator struct {
	Event *SessionManagerSessionReady // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionReady)
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
		it.Event = new(SessionManagerSessionReady)
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
func (it *SessionManagerSessionReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionReady represents a SessionReady event raised by the SessionManager contract.
type SessionManagerSessionReady struct {
	ReqId     *big.Int
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSessionReady is a free log retrieval operation binding the contract event 0xbd06d2615567f6a9673a4c804e1737a4f8039a009152dd72a817c9e8f7d91c25.
//
// Solidity: event SessionReady(uint256 indexed reqId, uint256 indexed sessionId)
func (_SessionManager *SessionManagerFilterer) FilterSessionReady(opts *bind.FilterOpts, reqId []*big.Int, sessionId []*big.Int) (*SessionManagerSessionReadyIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionReady", reqIdRule, sessionIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionReadyIterator{contract: _SessionManager.contract, event: "SessionReady", logs: logs, sub: sub}, nil
}

// WatchSessionReady is a free log subscription operation binding the contract event 0xbd06d2615567f6a9673a4c804e1737a4f8039a009152dd72a817c9e8f7d91c25.
//
// Solidity: event SessionReady(uint256 indexed reqId, uint256 indexed sessionId)
func (_SessionManager *SessionManagerFilterer) WatchSessionReady(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionReady, reqId []*big.Int, sessionId []*big.Int) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionReady", reqIdRule, sessionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionReady)
				if err := _SessionManager.contract.UnpackLog(event, "SessionReady", log); err != nil {
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

// ParseSessionReady is a log parse operation binding the contract event 0xbd06d2615567f6a9673a4c804e1737a4f8039a009152dd72a817c9e8f7d91c25.
//
// Solidity: event SessionReady(uint256 indexed reqId, uint256 indexed sessionId)
func (_SessionManager *SessionManagerFilterer) ParseSessionReady(log types.Log) (*SessionManagerSessionReady, error) {
	event := new(SessionManagerSessionReady)
	if err := _SessionManager.contract.UnpackLog(event, "SessionReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerSessionRequestedIterator is returned from FilterSessionRequested and is used to iterate over the raw logs and unpacked data for SessionRequested events raised by the SessionManager contract.
type SessionManagerSessionRequestedIterator struct {
	Event *SessionManagerSessionRequested // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionRequested)
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
		it.Event = new(SessionManagerSessionRequested)
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
func (it *SessionManagerSessionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionRequested represents a SessionRequested event raised by the SessionManager contract.
type SessionManagerSessionRequested struct {
	ReqId        *big.Int
	User         common.Address
	ModelId      [32]byte
	RequestBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSessionRequested is a free log retrieval operation binding the contract event 0x2bea7050fbb8aacbf4e8fae6ea4568546406d1150ad899b0425ce002620eba22.
//
// Solidity: event SessionRequested(uint256 indexed reqId, address indexed user, bytes32 indexed modelId, uint256 requestBlock)
func (_SessionManager *SessionManagerFilterer) FilterSessionRequested(opts *bind.FilterOpts, reqId []*big.Int, user []common.Address, modelId [][32]byte) (*SessionManagerSessionRequestedIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionRequested", reqIdRule, userRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionRequestedIterator{contract: _SessionManager.contract, event: "SessionRequested", logs: logs, sub: sub}, nil
}

// WatchSessionRequested is a free log subscription operation binding the contract event 0x2bea7050fbb8aacbf4e8fae6ea4568546406d1150ad899b0425ce002620eba22.
//
// Solidity: event SessionRequested(uint256 indexed reqId, address indexed user, bytes32 indexed modelId, uint256 requestBlock)
func (_SessionManager *SessionManagerFilterer) WatchSessionRequested(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionRequested, reqId []*big.Int, user []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionRequested", reqIdRule, userRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionRequested)
				if err := _SessionManager.contract.UnpackLog(event, "SessionRequested", log); err != nil {
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

// ParseSessionRequested is a log parse operation binding the contract event 0x2bea7050fbb8aacbf4e8fae6ea4568546406d1150ad899b0425ce002620eba22.
//
// Solidity: event SessionRequested(uint256 indexed reqId, address indexed user, bytes32 indexed modelId, uint256 requestBlock)
func (_SessionManager *SessionManagerFilterer) ParseSessionRequested(log types.Log) (*SessionManagerSessionRequested, error) {
	event := new(SessionManagerSessionRequested)
	if err := _SessionManager.contract.UnpackLog(event, "SessionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SessionManager contract.
type SessionManagerUpgradedIterator struct {
	Event *SessionManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *SessionManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerUpgraded)
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
		it.Event = new(SessionManagerUpgraded)
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
func (it *SessionManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerUpgraded represents a Upgraded event raised by the SessionManager contract.
type SessionManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SessionManager *SessionManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SessionManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerUpgradedIterator{contract: _SessionManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SessionManager *SessionManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SessionManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerUpgraded)
				if err := _SessionManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SessionManager *SessionManagerFilterer) ParseUpgraded(log types.Log) (*SessionManagerUpgraded, error) {
	event := new(SessionManagerUpgraded)
	if err := _SessionManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
