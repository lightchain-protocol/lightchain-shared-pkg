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

// ISessionManagerRequestTerms is an auto generated low-level Go binding around an user-defined struct.
type ISessionManagerRequestTerms struct {
	RequiredCaps   *big.Int
	ExcludedCaps   *big.Int
	MaxFee         *big.Int
	Worker         common.Address
	FallbackBlocks uint64
}

// SessionManagerMetaData contains all meta data concerning the SessionManager contract.
var SessionManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_PRICE_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelRequest\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimOpensAt\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSession\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"declineSession\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eligibleNow\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChosenWorker\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimPrice\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinPriceBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingClaims\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISessionManager.Request\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISessionManager.ReqStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestTerms\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISessionManager.RequestTerms\",\"components\":[{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"excludedCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCapabilities\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkerPrice\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workerRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aiConfig_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reputation_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minWorkerPrice\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideSessionKey\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSessionKeyFor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encWorkerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encDisputerKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteWorkerFee\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseClaim\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSession\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionFor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionForWithCapabilities\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionForWithTerms\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terms\",\"type\":\"tuple\",\"internalType\":\"structISessionManager.RequestTerms\",\"components\":[{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"excludedCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionWithCapabilities\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSessionWithTerms\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terms\",\"type\":\"tuple\",\"internalType\":\"structISessionManager.RequestTerms\",\"components\":[{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"excludedCaps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinPriceBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceDelayBlocks\",\"inputs\":[{\"name\":\"blocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkerPrice\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ClaimReleased\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinPriceBpsSet\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceDelayBlocksSet\",\"inputs\":[{\"name\":\"blocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestCancelled\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestDeclined\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reopened\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionClaimPrice\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionClaimed\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionReady\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sessionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionRequestTerms\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requiredCaps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"excludedCaps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionRequested\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SessionWorkerChosen\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fallbackBlocks\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkerPriceSet\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
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

// MAXPRICEDELAYBLOCKS is a free data retrieval call binding the contract method 0xfe283c40.
//
// Solidity: function MAX_PRICE_DELAY_BLOCKS() view returns(uint256)
func (_SessionManager *SessionManagerCaller) MAXPRICEDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "MAX_PRICE_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEDELAYBLOCKS is a free data retrieval call binding the contract method 0xfe283c40.
//
// Solidity: function MAX_PRICE_DELAY_BLOCKS() view returns(uint256)
func (_SessionManager *SessionManagerSession) MAXPRICEDELAYBLOCKS() (*big.Int, error) {
	return _SessionManager.Contract.MAXPRICEDELAYBLOCKS(&_SessionManager.CallOpts)
}

// MAXPRICEDELAYBLOCKS is a free data retrieval call binding the contract method 0xfe283c40.
//
// Solidity: function MAX_PRICE_DELAY_BLOCKS() view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) MAXPRICEDELAYBLOCKS() (*big.Int, error) {
	return _SessionManager.Contract.MAXPRICEDELAYBLOCKS(&_SessionManager.CallOpts)
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

// ClaimOpensAt is a free data retrieval call binding the contract method 0x4f2f4a8f.
//
// Solidity: function claimOpensAt(uint256 reqId, address worker) view returns(uint256)
func (_SessionManager *SessionManagerCaller) ClaimOpensAt(opts *bind.CallOpts, reqId *big.Int, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "claimOpensAt", reqId, worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimOpensAt is a free data retrieval call binding the contract method 0x4f2f4a8f.
//
// Solidity: function claimOpensAt(uint256 reqId, address worker) view returns(uint256)
func (_SessionManager *SessionManagerSession) ClaimOpensAt(reqId *big.Int, worker common.Address) (*big.Int, error) {
	return _SessionManager.Contract.ClaimOpensAt(&_SessionManager.CallOpts, reqId, worker)
}

// ClaimOpensAt is a free data retrieval call binding the contract method 0x4f2f4a8f.
//
// Solidity: function claimOpensAt(uint256 reqId, address worker) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) ClaimOpensAt(reqId *big.Int, worker common.Address) (*big.Int, error) {
	return _SessionManager.Contract.ClaimOpensAt(&_SessionManager.CallOpts, reqId, worker)
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

// GetChosenWorker is a free data retrieval call binding the contract method 0x8517e326.
//
// Solidity: function getChosenWorker(uint256 reqId) view returns(address worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerCaller) GetChosenWorker(opts *bind.CallOpts, reqId *big.Int) (struct {
	Worker         common.Address
	FallbackBlocks uint64
}, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getChosenWorker", reqId)

	outstruct := new(struct {
		Worker         common.Address
		FallbackBlocks uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Worker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FallbackBlocks = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetChosenWorker is a free data retrieval call binding the contract method 0x8517e326.
//
// Solidity: function getChosenWorker(uint256 reqId) view returns(address worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerSession) GetChosenWorker(reqId *big.Int) (struct {
	Worker         common.Address
	FallbackBlocks uint64
}, error) {
	return _SessionManager.Contract.GetChosenWorker(&_SessionManager.CallOpts, reqId)
}

// GetChosenWorker is a free data retrieval call binding the contract method 0x8517e326.
//
// Solidity: function getChosenWorker(uint256 reqId) view returns(address worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerCallerSession) GetChosenWorker(reqId *big.Int) (struct {
	Worker         common.Address
	FallbackBlocks uint64
}, error) {
	return _SessionManager.Contract.GetChosenWorker(&_SessionManager.CallOpts, reqId)
}

// GetClaimPrice is a free data retrieval call binding the contract method 0x245e5aac.
//
// Solidity: function getClaimPrice(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetClaimPrice(opts *bind.CallOpts, reqId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getClaimPrice", reqId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimPrice is a free data retrieval call binding the contract method 0x245e5aac.
//
// Solidity: function getClaimPrice(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerSession) GetClaimPrice(reqId *big.Int) (*big.Int, error) {
	return _SessionManager.Contract.GetClaimPrice(&_SessionManager.CallOpts, reqId)
}

// GetClaimPrice is a free data retrieval call binding the contract method 0x245e5aac.
//
// Solidity: function getClaimPrice(uint256 reqId) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetClaimPrice(reqId *big.Int) (*big.Int, error) {
	return _SessionManager.Contract.GetClaimPrice(&_SessionManager.CallOpts, reqId)
}

// GetMinPriceBps is a free data retrieval call binding the contract method 0xefbcbea5.
//
// Solidity: function getMinPriceBps() view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetMinPriceBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getMinPriceBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPriceBps is a free data retrieval call binding the contract method 0xefbcbea5.
//
// Solidity: function getMinPriceBps() view returns(uint256)
func (_SessionManager *SessionManagerSession) GetMinPriceBps() (*big.Int, error) {
	return _SessionManager.Contract.GetMinPriceBps(&_SessionManager.CallOpts)
}

// GetMinPriceBps is a free data retrieval call binding the contract method 0xefbcbea5.
//
// Solidity: function getMinPriceBps() view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetMinPriceBps() (*big.Int, error) {
	return _SessionManager.Contract.GetMinPriceBps(&_SessionManager.CallOpts)
}

// GetPendingClaims is a free data retrieval call binding the contract method 0x99f91405.
//
// Solidity: function getPendingClaims(address worker) view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetPendingClaims(opts *bind.CallOpts, worker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getPendingClaims", worker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingClaims is a free data retrieval call binding the contract method 0x99f91405.
//
// Solidity: function getPendingClaims(address worker) view returns(uint256)
func (_SessionManager *SessionManagerSession) GetPendingClaims(worker common.Address) (*big.Int, error) {
	return _SessionManager.Contract.GetPendingClaims(&_SessionManager.CallOpts, worker)
}

// GetPendingClaims is a free data retrieval call binding the contract method 0x99f91405.
//
// Solidity: function getPendingClaims(address worker) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetPendingClaims(worker common.Address) (*big.Int, error) {
	return _SessionManager.Contract.GetPendingClaims(&_SessionManager.CallOpts, worker)
}

// GetPriceDelayBlocks is a free data retrieval call binding the contract method 0xfb18cd35.
//
// Solidity: function getPriceDelayBlocks() view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetPriceDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getPriceDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceDelayBlocks is a free data retrieval call binding the contract method 0xfb18cd35.
//
// Solidity: function getPriceDelayBlocks() view returns(uint256)
func (_SessionManager *SessionManagerSession) GetPriceDelayBlocks() (*big.Int, error) {
	return _SessionManager.Contract.GetPriceDelayBlocks(&_SessionManager.CallOpts)
}

// GetPriceDelayBlocks is a free data retrieval call binding the contract method 0xfb18cd35.
//
// Solidity: function getPriceDelayBlocks() view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetPriceDelayBlocks() (*big.Int, error) {
	return _SessionManager.Contract.GetPriceDelayBlocks(&_SessionManager.CallOpts)
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

// GetRequestTerms is a free data retrieval call binding the contract method 0xeaff8f58.
//
// Solidity: function getRequestTerms(uint256 reqId) view returns((uint256,uint256,uint256,address,uint64))
func (_SessionManager *SessionManagerCaller) GetRequestTerms(opts *bind.CallOpts, reqId *big.Int) (ISessionManagerRequestTerms, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getRequestTerms", reqId)

	if err != nil {
		return *new(ISessionManagerRequestTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(ISessionManagerRequestTerms)).(*ISessionManagerRequestTerms)

	return out0, err

}

// GetRequestTerms is a free data retrieval call binding the contract method 0xeaff8f58.
//
// Solidity: function getRequestTerms(uint256 reqId) view returns((uint256,uint256,uint256,address,uint64))
func (_SessionManager *SessionManagerSession) GetRequestTerms(reqId *big.Int) (ISessionManagerRequestTerms, error) {
	return _SessionManager.Contract.GetRequestTerms(&_SessionManager.CallOpts, reqId)
}

// GetRequestTerms is a free data retrieval call binding the contract method 0xeaff8f58.
//
// Solidity: function getRequestTerms(uint256 reqId) view returns((uint256,uint256,uint256,address,uint64))
func (_SessionManager *SessionManagerCallerSession) GetRequestTerms(reqId *big.Int) (ISessionManagerRequestTerms, error) {
	return _SessionManager.Contract.GetRequestTerms(&_SessionManager.CallOpts, reqId)
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

// GetWorkerPrice is a free data retrieval call binding the contract method 0x58f116f6.
//
// Solidity: function getWorkerPrice(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCaller) GetWorkerPrice(opts *bind.CallOpts, worker common.Address, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "getWorkerPrice", worker, modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkerPrice is a free data retrieval call binding the contract method 0x58f116f6.
//
// Solidity: function getWorkerPrice(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerSession) GetWorkerPrice(worker common.Address, modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.GetWorkerPrice(&_SessionManager.CallOpts, worker, modelId)
}

// GetWorkerPrice is a free data retrieval call binding the contract method 0x58f116f6.
//
// Solidity: function getWorkerPrice(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) GetWorkerPrice(worker common.Address, modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.GetWorkerPrice(&_SessionManager.CallOpts, worker, modelId)
}

// MinWorkerPrice is a free data retrieval call binding the contract method 0x4f349166.
//
// Solidity: function minWorkerPrice(bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCaller) MinWorkerPrice(opts *bind.CallOpts, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "minWorkerPrice", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWorkerPrice is a free data retrieval call binding the contract method 0x4f349166.
//
// Solidity: function minWorkerPrice(bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerSession) MinWorkerPrice(modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.MinWorkerPrice(&_SessionManager.CallOpts, modelId)
}

// MinWorkerPrice is a free data retrieval call binding the contract method 0x4f349166.
//
// Solidity: function minWorkerPrice(bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) MinWorkerPrice(modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.MinWorkerPrice(&_SessionManager.CallOpts, modelId)
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

// QuoteWorkerFee is a free data retrieval call binding the contract method 0xabf087c2.
//
// Solidity: function quoteWorkerFee(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCaller) QuoteWorkerFee(opts *bind.CallOpts, worker common.Address, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SessionManager.contract.Call(opts, &out, "quoteWorkerFee", worker, modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteWorkerFee is a free data retrieval call binding the contract method 0xabf087c2.
//
// Solidity: function quoteWorkerFee(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerSession) QuoteWorkerFee(worker common.Address, modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.QuoteWorkerFee(&_SessionManager.CallOpts, worker, modelId)
}

// QuoteWorkerFee is a free data retrieval call binding the contract method 0xabf087c2.
//
// Solidity: function quoteWorkerFee(address worker, bytes32 modelId) view returns(uint256)
func (_SessionManager *SessionManagerCallerSession) QuoteWorkerFee(worker common.Address, modelId [32]byte) (*big.Int, error) {
	return _SessionManager.Contract.QuoteWorkerFee(&_SessionManager.CallOpts, worker, modelId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactor) CancelRequest(opts *bind.TransactOpts, reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "cancelRequest", reqId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 reqId) returns()
func (_SessionManager *SessionManagerSession) CancelRequest(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.CancelRequest(&_SessionManager.TransactOpts, reqId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactorSession) CancelRequest(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.CancelRequest(&_SessionManager.TransactOpts, reqId)
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

// DeclineSession is a paid mutator transaction binding the contract method 0xd31b7a24.
//
// Solidity: function declineSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactor) DeclineSession(opts *bind.TransactOpts, reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "declineSession", reqId)
}

// DeclineSession is a paid mutator transaction binding the contract method 0xd31b7a24.
//
// Solidity: function declineSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerSession) DeclineSession(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.DeclineSession(&_SessionManager.TransactOpts, reqId)
}

// DeclineSession is a paid mutator transaction binding the contract method 0xd31b7a24.
//
// Solidity: function declineSession(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactorSession) DeclineSession(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.DeclineSession(&_SessionManager.TransactOpts, reqId)
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

// ReleaseClaim is a paid mutator transaction binding the contract method 0xa8bde4eb.
//
// Solidity: function releaseClaim(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactor) ReleaseClaim(opts *bind.TransactOpts, reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "releaseClaim", reqId)
}

// ReleaseClaim is a paid mutator transaction binding the contract method 0xa8bde4eb.
//
// Solidity: function releaseClaim(uint256 reqId) returns()
func (_SessionManager *SessionManagerSession) ReleaseClaim(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.ReleaseClaim(&_SessionManager.TransactOpts, reqId)
}

// ReleaseClaim is a paid mutator transaction binding the contract method 0xa8bde4eb.
//
// Solidity: function releaseClaim(uint256 reqId) returns()
func (_SessionManager *SessionManagerTransactorSession) ReleaseClaim(reqId *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.ReleaseClaim(&_SessionManager.TransactOpts, reqId)
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

// RequestSessionForWithTerms is a paid mutator transaction binding the contract method 0xbfcebd47.
//
// Solidity: function requestSessionForWithTerms(address user, bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSessionForWithTerms(opts *bind.TransactOpts, user common.Address, modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSessionForWithTerms", user, modelId, expiry, terms)
}

// RequestSessionForWithTerms is a paid mutator transaction binding the contract method 0xbfcebd47.
//
// Solidity: function requestSessionForWithTerms(address user, bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSessionForWithTerms(user common.Address, modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionForWithTerms(&_SessionManager.TransactOpts, user, modelId, expiry, terms)
}

// RequestSessionForWithTerms is a paid mutator transaction binding the contract method 0xbfcebd47.
//
// Solidity: function requestSessionForWithTerms(address user, bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSessionForWithTerms(user common.Address, modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionForWithTerms(&_SessionManager.TransactOpts, user, modelId, expiry, terms)
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

// RequestSessionWithTerms is a paid mutator transaction binding the contract method 0x4f5e2db2.
//
// Solidity: function requestSessionWithTerms(bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactor) RequestSessionWithTerms(opts *bind.TransactOpts, modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "requestSessionWithTerms", modelId, expiry, terms)
}

// RequestSessionWithTerms is a paid mutator transaction binding the contract method 0x4f5e2db2.
//
// Solidity: function requestSessionWithTerms(bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerSession) RequestSessionWithTerms(modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionWithTerms(&_SessionManager.TransactOpts, modelId, expiry, terms)
}

// RequestSessionWithTerms is a paid mutator transaction binding the contract method 0x4f5e2db2.
//
// Solidity: function requestSessionWithTerms(bytes32 modelId, uint256 expiry, (uint256,uint256,uint256,address,uint64) terms) returns(uint256 reqId)
func (_SessionManager *SessionManagerTransactorSession) RequestSessionWithTerms(modelId [32]byte, expiry *big.Int, terms ISessionManagerRequestTerms) (*types.Transaction, error) {
	return _SessionManager.Contract.RequestSessionWithTerms(&_SessionManager.TransactOpts, modelId, expiry, terms)
}

// SetMinPriceBps is a paid mutator transaction binding the contract method 0x4555a357.
//
// Solidity: function setMinPriceBps(uint256 bps) returns()
func (_SessionManager *SessionManagerTransactor) SetMinPriceBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "setMinPriceBps", bps)
}

// SetMinPriceBps is a paid mutator transaction binding the contract method 0x4555a357.
//
// Solidity: function setMinPriceBps(uint256 bps) returns()
func (_SessionManager *SessionManagerSession) SetMinPriceBps(bps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetMinPriceBps(&_SessionManager.TransactOpts, bps)
}

// SetMinPriceBps is a paid mutator transaction binding the contract method 0x4555a357.
//
// Solidity: function setMinPriceBps(uint256 bps) returns()
func (_SessionManager *SessionManagerTransactorSession) SetMinPriceBps(bps *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetMinPriceBps(&_SessionManager.TransactOpts, bps)
}

// SetPriceDelayBlocks is a paid mutator transaction binding the contract method 0x477c11ad.
//
// Solidity: function setPriceDelayBlocks(uint256 blocks) returns()
func (_SessionManager *SessionManagerTransactor) SetPriceDelayBlocks(opts *bind.TransactOpts, blocks *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "setPriceDelayBlocks", blocks)
}

// SetPriceDelayBlocks is a paid mutator transaction binding the contract method 0x477c11ad.
//
// Solidity: function setPriceDelayBlocks(uint256 blocks) returns()
func (_SessionManager *SessionManagerSession) SetPriceDelayBlocks(blocks *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetPriceDelayBlocks(&_SessionManager.TransactOpts, blocks)
}

// SetPriceDelayBlocks is a paid mutator transaction binding the contract method 0x477c11ad.
//
// Solidity: function setPriceDelayBlocks(uint256 blocks) returns()
func (_SessionManager *SessionManagerTransactorSession) SetPriceDelayBlocks(blocks *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetPriceDelayBlocks(&_SessionManager.TransactOpts, blocks)
}

// SetWorkerPrice is a paid mutator transaction binding the contract method 0x56b740f9.
//
// Solidity: function setWorkerPrice(bytes32 modelId, uint256 price) returns()
func (_SessionManager *SessionManagerTransactor) SetWorkerPrice(opts *bind.TransactOpts, modelId [32]byte, price *big.Int) (*types.Transaction, error) {
	return _SessionManager.contract.Transact(opts, "setWorkerPrice", modelId, price)
}

// SetWorkerPrice is a paid mutator transaction binding the contract method 0x56b740f9.
//
// Solidity: function setWorkerPrice(bytes32 modelId, uint256 price) returns()
func (_SessionManager *SessionManagerSession) SetWorkerPrice(modelId [32]byte, price *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetWorkerPrice(&_SessionManager.TransactOpts, modelId, price)
}

// SetWorkerPrice is a paid mutator transaction binding the contract method 0x56b740f9.
//
// Solidity: function setWorkerPrice(bytes32 modelId, uint256 price) returns()
func (_SessionManager *SessionManagerTransactorSession) SetWorkerPrice(modelId [32]byte, price *big.Int) (*types.Transaction, error) {
	return _SessionManager.Contract.SetWorkerPrice(&_SessionManager.TransactOpts, modelId, price)
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

// SessionManagerClaimReleasedIterator is returned from FilterClaimReleased and is used to iterate over the raw logs and unpacked data for ClaimReleased events raised by the SessionManager contract.
type SessionManagerClaimReleasedIterator struct {
	Event *SessionManagerClaimReleased // Event containing the contract specifics and raw log

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
func (it *SessionManagerClaimReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerClaimReleased)
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
		it.Event = new(SessionManagerClaimReleased)
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
func (it *SessionManagerClaimReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerClaimReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerClaimReleased represents a ClaimReleased event raised by the SessionManager contract.
type SessionManagerClaimReleased struct {
	ReqId  *big.Int
	Worker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimReleased is a free log retrieval operation binding the contract event 0x909d0309e792d4f17972cef6c4067bfd0b5fa112dd3aa00415fbac21afa3eb3e.
//
// Solidity: event ClaimReleased(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) FilterClaimReleased(opts *bind.FilterOpts, reqId []*big.Int, worker []common.Address) (*SessionManagerClaimReleasedIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "ClaimReleased", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerClaimReleasedIterator{contract: _SessionManager.contract, event: "ClaimReleased", logs: logs, sub: sub}, nil
}

// WatchClaimReleased is a free log subscription operation binding the contract event 0x909d0309e792d4f17972cef6c4067bfd0b5fa112dd3aa00415fbac21afa3eb3e.
//
// Solidity: event ClaimReleased(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) WatchClaimReleased(opts *bind.WatchOpts, sink chan<- *SessionManagerClaimReleased, reqId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "ClaimReleased", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerClaimReleased)
				if err := _SessionManager.contract.UnpackLog(event, "ClaimReleased", log); err != nil {
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

// ParseClaimReleased is a log parse operation binding the contract event 0x909d0309e792d4f17972cef6c4067bfd0b5fa112dd3aa00415fbac21afa3eb3e.
//
// Solidity: event ClaimReleased(uint256 indexed reqId, address indexed worker)
func (_SessionManager *SessionManagerFilterer) ParseClaimReleased(log types.Log) (*SessionManagerClaimReleased, error) {
	event := new(SessionManagerClaimReleased)
	if err := _SessionManager.contract.UnpackLog(event, "ClaimReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SessionManagerMinPriceBpsSetIterator is returned from FilterMinPriceBpsSet and is used to iterate over the raw logs and unpacked data for MinPriceBpsSet events raised by the SessionManager contract.
type SessionManagerMinPriceBpsSetIterator struct {
	Event *SessionManagerMinPriceBpsSet // Event containing the contract specifics and raw log

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
func (it *SessionManagerMinPriceBpsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerMinPriceBpsSet)
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
		it.Event = new(SessionManagerMinPriceBpsSet)
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
func (it *SessionManagerMinPriceBpsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerMinPriceBpsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerMinPriceBpsSet represents a MinPriceBpsSet event raised by the SessionManager contract.
type SessionManagerMinPriceBpsSet struct {
	Bps *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinPriceBpsSet is a free log retrieval operation binding the contract event 0xb22a65aff0eb78e17b2c51dedf7609b5ff7d4db86871f3dedb02fd4745c8b268.
//
// Solidity: event MinPriceBpsSet(uint256 bps)
func (_SessionManager *SessionManagerFilterer) FilterMinPriceBpsSet(opts *bind.FilterOpts) (*SessionManagerMinPriceBpsSetIterator, error) {

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "MinPriceBpsSet")
	if err != nil {
		return nil, err
	}
	return &SessionManagerMinPriceBpsSetIterator{contract: _SessionManager.contract, event: "MinPriceBpsSet", logs: logs, sub: sub}, nil
}

// WatchMinPriceBpsSet is a free log subscription operation binding the contract event 0xb22a65aff0eb78e17b2c51dedf7609b5ff7d4db86871f3dedb02fd4745c8b268.
//
// Solidity: event MinPriceBpsSet(uint256 bps)
func (_SessionManager *SessionManagerFilterer) WatchMinPriceBpsSet(opts *bind.WatchOpts, sink chan<- *SessionManagerMinPriceBpsSet) (event.Subscription, error) {

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "MinPriceBpsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerMinPriceBpsSet)
				if err := _SessionManager.contract.UnpackLog(event, "MinPriceBpsSet", log); err != nil {
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

// ParseMinPriceBpsSet is a log parse operation binding the contract event 0xb22a65aff0eb78e17b2c51dedf7609b5ff7d4db86871f3dedb02fd4745c8b268.
//
// Solidity: event MinPriceBpsSet(uint256 bps)
func (_SessionManager *SessionManagerFilterer) ParseMinPriceBpsSet(log types.Log) (*SessionManagerMinPriceBpsSet, error) {
	event := new(SessionManagerMinPriceBpsSet)
	if err := _SessionManager.contract.UnpackLog(event, "MinPriceBpsSet", log); err != nil {
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

// SessionManagerPriceDelayBlocksSetIterator is returned from FilterPriceDelayBlocksSet and is used to iterate over the raw logs and unpacked data for PriceDelayBlocksSet events raised by the SessionManager contract.
type SessionManagerPriceDelayBlocksSetIterator struct {
	Event *SessionManagerPriceDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *SessionManagerPriceDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerPriceDelayBlocksSet)
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
		it.Event = new(SessionManagerPriceDelayBlocksSet)
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
func (it *SessionManagerPriceDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerPriceDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerPriceDelayBlocksSet represents a PriceDelayBlocksSet event raised by the SessionManager contract.
type SessionManagerPriceDelayBlocksSet struct {
	Blocks *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPriceDelayBlocksSet is a free log retrieval operation binding the contract event 0x49cbb0cf02bf9755553abf81d5a11ecf2f8dc2de3d9d8da35e6565b096e071f2.
//
// Solidity: event PriceDelayBlocksSet(uint256 blocks)
func (_SessionManager *SessionManagerFilterer) FilterPriceDelayBlocksSet(opts *bind.FilterOpts) (*SessionManagerPriceDelayBlocksSetIterator, error) {

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "PriceDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &SessionManagerPriceDelayBlocksSetIterator{contract: _SessionManager.contract, event: "PriceDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchPriceDelayBlocksSet is a free log subscription operation binding the contract event 0x49cbb0cf02bf9755553abf81d5a11ecf2f8dc2de3d9d8da35e6565b096e071f2.
//
// Solidity: event PriceDelayBlocksSet(uint256 blocks)
func (_SessionManager *SessionManagerFilterer) WatchPriceDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *SessionManagerPriceDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "PriceDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerPriceDelayBlocksSet)
				if err := _SessionManager.contract.UnpackLog(event, "PriceDelayBlocksSet", log); err != nil {
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

// ParsePriceDelayBlocksSet is a log parse operation binding the contract event 0x49cbb0cf02bf9755553abf81d5a11ecf2f8dc2de3d9d8da35e6565b096e071f2.
//
// Solidity: event PriceDelayBlocksSet(uint256 blocks)
func (_SessionManager *SessionManagerFilterer) ParsePriceDelayBlocksSet(log types.Log) (*SessionManagerPriceDelayBlocksSet, error) {
	event := new(SessionManagerPriceDelayBlocksSet)
	if err := _SessionManager.contract.UnpackLog(event, "PriceDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerRequestCancelledIterator is returned from FilterRequestCancelled and is used to iterate over the raw logs and unpacked data for RequestCancelled events raised by the SessionManager contract.
type SessionManagerRequestCancelledIterator struct {
	Event *SessionManagerRequestCancelled // Event containing the contract specifics and raw log

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
func (it *SessionManagerRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerRequestCancelled)
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
		it.Event = new(SessionManagerRequestCancelled)
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
func (it *SessionManagerRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerRequestCancelled represents a RequestCancelled event raised by the SessionManager contract.
type SessionManagerRequestCancelled struct {
	ReqId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRequestCancelled is a free log retrieval operation binding the contract event 0xee243f878b7fc2f54e934ca33783d4395d42bc07612e2bd4b8e0e178639f7a28.
//
// Solidity: event RequestCancelled(uint256 indexed reqId)
func (_SessionManager *SessionManagerFilterer) FilterRequestCancelled(opts *bind.FilterOpts, reqId []*big.Int) (*SessionManagerRequestCancelledIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "RequestCancelled", reqIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerRequestCancelledIterator{contract: _SessionManager.contract, event: "RequestCancelled", logs: logs, sub: sub}, nil
}

// WatchRequestCancelled is a free log subscription operation binding the contract event 0xee243f878b7fc2f54e934ca33783d4395d42bc07612e2bd4b8e0e178639f7a28.
//
// Solidity: event RequestCancelled(uint256 indexed reqId)
func (_SessionManager *SessionManagerFilterer) WatchRequestCancelled(opts *bind.WatchOpts, sink chan<- *SessionManagerRequestCancelled, reqId []*big.Int) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "RequestCancelled", reqIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerRequestCancelled)
				if err := _SessionManager.contract.UnpackLog(event, "RequestCancelled", log); err != nil {
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

// ParseRequestCancelled is a log parse operation binding the contract event 0xee243f878b7fc2f54e934ca33783d4395d42bc07612e2bd4b8e0e178639f7a28.
//
// Solidity: event RequestCancelled(uint256 indexed reqId)
func (_SessionManager *SessionManagerFilterer) ParseRequestCancelled(log types.Log) (*SessionManagerRequestCancelled, error) {
	event := new(SessionManagerRequestCancelled)
	if err := _SessionManager.contract.UnpackLog(event, "RequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerRequestDeclinedIterator is returned from FilterRequestDeclined and is used to iterate over the raw logs and unpacked data for RequestDeclined events raised by the SessionManager contract.
type SessionManagerRequestDeclinedIterator struct {
	Event *SessionManagerRequestDeclined // Event containing the contract specifics and raw log

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
func (it *SessionManagerRequestDeclinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerRequestDeclined)
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
		it.Event = new(SessionManagerRequestDeclined)
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
func (it *SessionManagerRequestDeclinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerRequestDeclinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerRequestDeclined represents a RequestDeclined event raised by the SessionManager contract.
type SessionManagerRequestDeclined struct {
	ReqId    *big.Int
	Worker   common.Address
	Reopened bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestDeclined is a free log retrieval operation binding the contract event 0xe010058907f2db92af81a69d9c862f0a98ebe7f459e418151a385e9728fee74f.
//
// Solidity: event RequestDeclined(uint256 indexed reqId, address indexed worker, bool reopened)
func (_SessionManager *SessionManagerFilterer) FilterRequestDeclined(opts *bind.FilterOpts, reqId []*big.Int, worker []common.Address) (*SessionManagerRequestDeclinedIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "RequestDeclined", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerRequestDeclinedIterator{contract: _SessionManager.contract, event: "RequestDeclined", logs: logs, sub: sub}, nil
}

// WatchRequestDeclined is a free log subscription operation binding the contract event 0xe010058907f2db92af81a69d9c862f0a98ebe7f459e418151a385e9728fee74f.
//
// Solidity: event RequestDeclined(uint256 indexed reqId, address indexed worker, bool reopened)
func (_SessionManager *SessionManagerFilterer) WatchRequestDeclined(opts *bind.WatchOpts, sink chan<- *SessionManagerRequestDeclined, reqId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "RequestDeclined", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerRequestDeclined)
				if err := _SessionManager.contract.UnpackLog(event, "RequestDeclined", log); err != nil {
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

// ParseRequestDeclined is a log parse operation binding the contract event 0xe010058907f2db92af81a69d9c862f0a98ebe7f459e418151a385e9728fee74f.
//
// Solidity: event RequestDeclined(uint256 indexed reqId, address indexed worker, bool reopened)
func (_SessionManager *SessionManagerFilterer) ParseRequestDeclined(log types.Log) (*SessionManagerRequestDeclined, error) {
	event := new(SessionManagerRequestDeclined)
	if err := _SessionManager.contract.UnpackLog(event, "RequestDeclined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionManagerSessionClaimPriceIterator is returned from FilterSessionClaimPrice and is used to iterate over the raw logs and unpacked data for SessionClaimPrice events raised by the SessionManager contract.
type SessionManagerSessionClaimPriceIterator struct {
	Event *SessionManagerSessionClaimPrice // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionClaimPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionClaimPrice)
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
		it.Event = new(SessionManagerSessionClaimPrice)
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
func (it *SessionManagerSessionClaimPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionClaimPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionClaimPrice represents a SessionClaimPrice event raised by the SessionManager contract.
type SessionManagerSessionClaimPrice struct {
	ReqId *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSessionClaimPrice is a free log retrieval operation binding the contract event 0x934b676dff9ea29a4b736fb44b5e6f9a4f9f91b2823e83142788c2a6144fd071.
//
// Solidity: event SessionClaimPrice(uint256 indexed reqId, uint256 price)
func (_SessionManager *SessionManagerFilterer) FilterSessionClaimPrice(opts *bind.FilterOpts, reqId []*big.Int) (*SessionManagerSessionClaimPriceIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionClaimPrice", reqIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionClaimPriceIterator{contract: _SessionManager.contract, event: "SessionClaimPrice", logs: logs, sub: sub}, nil
}

// WatchSessionClaimPrice is a free log subscription operation binding the contract event 0x934b676dff9ea29a4b736fb44b5e6f9a4f9f91b2823e83142788c2a6144fd071.
//
// Solidity: event SessionClaimPrice(uint256 indexed reqId, uint256 price)
func (_SessionManager *SessionManagerFilterer) WatchSessionClaimPrice(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionClaimPrice, reqId []*big.Int) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionClaimPrice", reqIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionClaimPrice)
				if err := _SessionManager.contract.UnpackLog(event, "SessionClaimPrice", log); err != nil {
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

// ParseSessionClaimPrice is a log parse operation binding the contract event 0x934b676dff9ea29a4b736fb44b5e6f9a4f9f91b2823e83142788c2a6144fd071.
//
// Solidity: event SessionClaimPrice(uint256 indexed reqId, uint256 price)
func (_SessionManager *SessionManagerFilterer) ParseSessionClaimPrice(log types.Log) (*SessionManagerSessionClaimPrice, error) {
	event := new(SessionManagerSessionClaimPrice)
	if err := _SessionManager.contract.UnpackLog(event, "SessionClaimPrice", log); err != nil {
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

// SessionManagerSessionRequestTermsIterator is returned from FilterSessionRequestTerms and is used to iterate over the raw logs and unpacked data for SessionRequestTerms events raised by the SessionManager contract.
type SessionManagerSessionRequestTermsIterator struct {
	Event *SessionManagerSessionRequestTerms // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionRequestTermsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionRequestTerms)
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
		it.Event = new(SessionManagerSessionRequestTerms)
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
func (it *SessionManagerSessionRequestTermsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionRequestTermsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionRequestTerms represents a SessionRequestTerms event raised by the SessionManager contract.
type SessionManagerSessionRequestTerms struct {
	ReqId        *big.Int
	RequiredCaps *big.Int
	ExcludedCaps *big.Int
	MaxFee       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSessionRequestTerms is a free log retrieval operation binding the contract event 0xe84c1c9a41ce81de97a0bb6d76789bf3e1480fd8276932c2edca37dca75491bd.
//
// Solidity: event SessionRequestTerms(uint256 indexed reqId, uint256 requiredCaps, uint256 excludedCaps, uint256 maxFee)
func (_SessionManager *SessionManagerFilterer) FilterSessionRequestTerms(opts *bind.FilterOpts, reqId []*big.Int) (*SessionManagerSessionRequestTermsIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionRequestTerms", reqIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionRequestTermsIterator{contract: _SessionManager.contract, event: "SessionRequestTerms", logs: logs, sub: sub}, nil
}

// WatchSessionRequestTerms is a free log subscription operation binding the contract event 0xe84c1c9a41ce81de97a0bb6d76789bf3e1480fd8276932c2edca37dca75491bd.
//
// Solidity: event SessionRequestTerms(uint256 indexed reqId, uint256 requiredCaps, uint256 excludedCaps, uint256 maxFee)
func (_SessionManager *SessionManagerFilterer) WatchSessionRequestTerms(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionRequestTerms, reqId []*big.Int) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionRequestTerms", reqIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionRequestTerms)
				if err := _SessionManager.contract.UnpackLog(event, "SessionRequestTerms", log); err != nil {
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

// ParseSessionRequestTerms is a log parse operation binding the contract event 0xe84c1c9a41ce81de97a0bb6d76789bf3e1480fd8276932c2edca37dca75491bd.
//
// Solidity: event SessionRequestTerms(uint256 indexed reqId, uint256 requiredCaps, uint256 excludedCaps, uint256 maxFee)
func (_SessionManager *SessionManagerFilterer) ParseSessionRequestTerms(log types.Log) (*SessionManagerSessionRequestTerms, error) {
	event := new(SessionManagerSessionRequestTerms)
	if err := _SessionManager.contract.UnpackLog(event, "SessionRequestTerms", log); err != nil {
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

// SessionManagerSessionWorkerChosenIterator is returned from FilterSessionWorkerChosen and is used to iterate over the raw logs and unpacked data for SessionWorkerChosen events raised by the SessionManager contract.
type SessionManagerSessionWorkerChosenIterator struct {
	Event *SessionManagerSessionWorkerChosen // Event containing the contract specifics and raw log

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
func (it *SessionManagerSessionWorkerChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerSessionWorkerChosen)
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
		it.Event = new(SessionManagerSessionWorkerChosen)
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
func (it *SessionManagerSessionWorkerChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerSessionWorkerChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerSessionWorkerChosen represents a SessionWorkerChosen event raised by the SessionManager contract.
type SessionManagerSessionWorkerChosen struct {
	ReqId          *big.Int
	Worker         common.Address
	FallbackBlocks uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSessionWorkerChosen is a free log retrieval operation binding the contract event 0x91d7b588bdfeb6e0a6991171cf775c7689e4926db7c269e430b858927cef6753.
//
// Solidity: event SessionWorkerChosen(uint256 indexed reqId, address indexed worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerFilterer) FilterSessionWorkerChosen(opts *bind.FilterOpts, reqId []*big.Int, worker []common.Address) (*SessionManagerSessionWorkerChosenIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "SessionWorkerChosen", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerSessionWorkerChosenIterator{contract: _SessionManager.contract, event: "SessionWorkerChosen", logs: logs, sub: sub}, nil
}

// WatchSessionWorkerChosen is a free log subscription operation binding the contract event 0x91d7b588bdfeb6e0a6991171cf775c7689e4926db7c269e430b858927cef6753.
//
// Solidity: event SessionWorkerChosen(uint256 indexed reqId, address indexed worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerFilterer) WatchSessionWorkerChosen(opts *bind.WatchOpts, sink chan<- *SessionManagerSessionWorkerChosen, reqId []*big.Int, worker []common.Address) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "SessionWorkerChosen", reqIdRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerSessionWorkerChosen)
				if err := _SessionManager.contract.UnpackLog(event, "SessionWorkerChosen", log); err != nil {
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

// ParseSessionWorkerChosen is a log parse operation binding the contract event 0x91d7b588bdfeb6e0a6991171cf775c7689e4926db7c269e430b858927cef6753.
//
// Solidity: event SessionWorkerChosen(uint256 indexed reqId, address indexed worker, uint64 fallbackBlocks)
func (_SessionManager *SessionManagerFilterer) ParseSessionWorkerChosen(log types.Log) (*SessionManagerSessionWorkerChosen, error) {
	event := new(SessionManagerSessionWorkerChosen)
	if err := _SessionManager.contract.UnpackLog(event, "SessionWorkerChosen", log); err != nil {
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

// SessionManagerWorkerPriceSetIterator is returned from FilterWorkerPriceSet and is used to iterate over the raw logs and unpacked data for WorkerPriceSet events raised by the SessionManager contract.
type SessionManagerWorkerPriceSetIterator struct {
	Event *SessionManagerWorkerPriceSet // Event containing the contract specifics and raw log

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
func (it *SessionManagerWorkerPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionManagerWorkerPriceSet)
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
		it.Event = new(SessionManagerWorkerPriceSet)
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
func (it *SessionManagerWorkerPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionManagerWorkerPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionManagerWorkerPriceSet represents a WorkerPriceSet event raised by the SessionManager contract.
type SessionManagerWorkerPriceSet struct {
	Worker  common.Address
	ModelId [32]byte
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWorkerPriceSet is a free log retrieval operation binding the contract event 0x0880ce51c56473edccdcd9b2164501f3e83618512e1c42f7807a2eefff3ce6ba.
//
// Solidity: event WorkerPriceSet(address indexed worker, bytes32 indexed modelId, uint256 price)
func (_SessionManager *SessionManagerFilterer) FilterWorkerPriceSet(opts *bind.FilterOpts, worker []common.Address, modelId [][32]byte) (*SessionManagerWorkerPriceSetIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _SessionManager.contract.FilterLogs(opts, "WorkerPriceSet", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionManagerWorkerPriceSetIterator{contract: _SessionManager.contract, event: "WorkerPriceSet", logs: logs, sub: sub}, nil
}

// WatchWorkerPriceSet is a free log subscription operation binding the contract event 0x0880ce51c56473edccdcd9b2164501f3e83618512e1c42f7807a2eefff3ce6ba.
//
// Solidity: event WorkerPriceSet(address indexed worker, bytes32 indexed modelId, uint256 price)
func (_SessionManager *SessionManagerFilterer) WatchWorkerPriceSet(opts *bind.WatchOpts, sink chan<- *SessionManagerWorkerPriceSet, worker []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _SessionManager.contract.WatchLogs(opts, "WorkerPriceSet", workerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionManagerWorkerPriceSet)
				if err := _SessionManager.contract.UnpackLog(event, "WorkerPriceSet", log); err != nil {
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

// ParseWorkerPriceSet is a log parse operation binding the contract event 0x0880ce51c56473edccdcd9b2164501f3e83618512e1c42f7807a2eefff3ce6ba.
//
// Solidity: event WorkerPriceSet(address indexed worker, bytes32 indexed modelId, uint256 price)
func (_SessionManager *SessionManagerFilterer) ParseWorkerPriceSet(log types.Log) (*SessionManagerWorkerPriceSet, error) {
	event := new(SessionManagerWorkerPriceSet)
	if err := _SessionManager.contract.UnpackLog(event, "WorkerPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
