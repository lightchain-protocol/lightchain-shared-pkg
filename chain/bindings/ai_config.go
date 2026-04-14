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

// AIConfigMetaData contains all meta data concerning the AIConfig contract.
var AIConfigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateJobFee\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAckTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobRetentionPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCanarySimilarityThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCompletionTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCompletionTimeoutSlashBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDispatcherAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisputeBondMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisputeSlashBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisputeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisputerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxReassignments\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxSlashBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinAckTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinDisputeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinResolutionTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinWorkerStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getModelFee\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getModelMaxOutputTokens\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProtocolFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResolutionTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSamplingRateBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSessionInactivityTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSimilarityThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSuspensionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSuspensionThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimeoutSlashBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkerFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_disputer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isModelEnabled\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerModel\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxOutputTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAckTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlobRetentionPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCanarySimilarityThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCompletionTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCompletionTimeoutSlashBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDispatcherAddress\",\"inputs\":[{\"name\":\"dispatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDisputeBondMultiplier\",\"inputs\":[{\"name\":\"multiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDisputeSlashBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDisputeWindow\",\"inputs\":[{\"name\":\"window\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDisputerAddress\",\"inputs\":[{\"name\":\"disputer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeDistribution\",\"inputs\":[{\"name\":\"workerBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"protocolBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"burnBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxReassignments\",\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSlashBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinAckTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinDisputeWindow\",\"inputs\":[{\"name\":\"window\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinResolutionTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWorkerStake\",\"inputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setModelFee\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setModelMaxOutputTokens\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setResolutionTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSamplingRateBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSessionInactivityTimeout\",\"inputs\":[{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSimilarityThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSuspensionCooldown\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSuspensionThreshold\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTimeoutSlashBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DispatcherAddressUpdated\",\"inputs\":[{\"name\":\"newDispatcher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeParamsUpdated\",\"inputs\":[{\"name\":\"bondMultiplier\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"window\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"resolutionTimeout\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputerAddressUpdated\",\"inputs\":[{\"name\":\"newDisputer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeDistributionUpdated\",\"inputs\":[{\"name\":\"workerBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InfraParamsUpdated\",\"inputs\":[{\"name\":\"param\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWorkerStakeUpdated\",\"inputs\":[{\"name\":\"newStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelConfigUpdated\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxOutputTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelDisabled\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelEnabled\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelRegistered\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxOutputTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingParamsUpdated\",\"inputs\":[{\"name\":\"timeoutSlashBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"completionTimeoutSlashBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"disputeSlashBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuspensionParamsUpdated\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cooldown\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// AIConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use AIConfigMetaData.ABI instead.
var AIConfigABI = AIConfigMetaData.ABI

// AIConfig is an auto generated Go binding around an Ethereum contract.
type AIConfig struct {
	AIConfigCaller     // Read-only binding to the contract
	AIConfigTransactor // Write-only binding to the contract
	AIConfigFilterer   // Log filterer for contract events
}

// AIConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type AIConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AIConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AIConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AIConfigSession struct {
	Contract     *AIConfig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AIConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AIConfigCallerSession struct {
	Contract *AIConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AIConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AIConfigTransactorSession struct {
	Contract     *AIConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AIConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type AIConfigRaw struct {
	Contract *AIConfig // Generic contract binding to access the raw methods on
}

// AIConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AIConfigCallerRaw struct {
	Contract *AIConfigCaller // Generic read-only contract binding to access the raw methods on
}

// AIConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AIConfigTransactorRaw struct {
	Contract *AIConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAIConfig creates a new instance of AIConfig, bound to a specific deployed contract.
func NewAIConfig(address common.Address, backend bind.ContractBackend) (*AIConfig, error) {
	contract, err := bindAIConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AIConfig{AIConfigCaller: AIConfigCaller{contract: contract}, AIConfigTransactor: AIConfigTransactor{contract: contract}, AIConfigFilterer: AIConfigFilterer{contract: contract}}, nil
}

// NewAIConfigCaller creates a new read-only instance of AIConfig, bound to a specific deployed contract.
func NewAIConfigCaller(address common.Address, caller bind.ContractCaller) (*AIConfigCaller, error) {
	contract, err := bindAIConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AIConfigCaller{contract: contract}, nil
}

// NewAIConfigTransactor creates a new write-only instance of AIConfig, bound to a specific deployed contract.
func NewAIConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*AIConfigTransactor, error) {
	contract, err := bindAIConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AIConfigTransactor{contract: contract}, nil
}

// NewAIConfigFilterer creates a new log filterer instance of AIConfig, bound to a specific deployed contract.
func NewAIConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*AIConfigFilterer, error) {
	contract, err := bindAIConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AIConfigFilterer{contract: contract}, nil
}

// bindAIConfig binds a generic wrapper to an already deployed contract.
func bindAIConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AIConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIConfig *AIConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIConfig.Contract.AIConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIConfig *AIConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIConfig.Contract.AIConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIConfig *AIConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIConfig.Contract.AIConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIConfig *AIConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIConfig *AIConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIConfig *AIConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIConfig.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AIConfig *AIConfigCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AIConfig *AIConfigSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AIConfig.Contract.UPGRADEINTERFACEVERSION(&_AIConfig.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AIConfig *AIConfigCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AIConfig.Contract.UPGRADEINTERFACEVERSION(&_AIConfig.CallOpts)
}

// CalculateJobFee is a free data retrieval call binding the contract method 0x33763d83.
//
// Solidity: function calculateJobFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCaller) CalculateJobFee(opts *bind.CallOpts, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "calculateJobFee", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateJobFee is a free data retrieval call binding the contract method 0x33763d83.
//
// Solidity: function calculateJobFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigSession) CalculateJobFee(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.CalculateJobFee(&_AIConfig.CallOpts, modelId)
}

// CalculateJobFee is a free data retrieval call binding the contract method 0x33763d83.
//
// Solidity: function calculateJobFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCallerSession) CalculateJobFee(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.CalculateJobFee(&_AIConfig.CallOpts, modelId)
}

// GetAckTimeout is a free data retrieval call binding the contract method 0x6ef7016d.
//
// Solidity: function getAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetAckTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getAckTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAckTimeout is a free data retrieval call binding the contract method 0x6ef7016d.
//
// Solidity: function getAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetAckTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetAckTimeout(&_AIConfig.CallOpts)
}

// GetAckTimeout is a free data retrieval call binding the contract method 0x6ef7016d.
//
// Solidity: function getAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetAckTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetAckTimeout(&_AIConfig.CallOpts)
}

// GetBlobRetentionPeriod is a free data retrieval call binding the contract method 0xa537e009.
//
// Solidity: function getBlobRetentionPeriod() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetBlobRetentionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getBlobRetentionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlobRetentionPeriod is a free data retrieval call binding the contract method 0xa537e009.
//
// Solidity: function getBlobRetentionPeriod() view returns(uint256)
func (_AIConfig *AIConfigSession) GetBlobRetentionPeriod() (*big.Int, error) {
	return _AIConfig.Contract.GetBlobRetentionPeriod(&_AIConfig.CallOpts)
}

// GetBlobRetentionPeriod is a free data retrieval call binding the contract method 0xa537e009.
//
// Solidity: function getBlobRetentionPeriod() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetBlobRetentionPeriod() (*big.Int, error) {
	return _AIConfig.Contract.GetBlobRetentionPeriod(&_AIConfig.CallOpts)
}

// GetBurnFeeBps is a free data retrieval call binding the contract method 0x5dd2c216.
//
// Solidity: function getBurnFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetBurnFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getBurnFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnFeeBps is a free data retrieval call binding the contract method 0x5dd2c216.
//
// Solidity: function getBurnFeeBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetBurnFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetBurnFeeBps(&_AIConfig.CallOpts)
}

// GetBurnFeeBps is a free data retrieval call binding the contract method 0x5dd2c216.
//
// Solidity: function getBurnFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetBurnFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetBurnFeeBps(&_AIConfig.CallOpts)
}

// GetCanarySimilarityThreshold is a free data retrieval call binding the contract method 0x9a316fd3.
//
// Solidity: function getCanarySimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetCanarySimilarityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getCanarySimilarityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCanarySimilarityThreshold is a free data retrieval call binding the contract method 0x9a316fd3.
//
// Solidity: function getCanarySimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigSession) GetCanarySimilarityThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetCanarySimilarityThreshold(&_AIConfig.CallOpts)
}

// GetCanarySimilarityThreshold is a free data retrieval call binding the contract method 0x9a316fd3.
//
// Solidity: function getCanarySimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetCanarySimilarityThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetCanarySimilarityThreshold(&_AIConfig.CallOpts)
}

// GetCompletionTimeout is a free data retrieval call binding the contract method 0xededfc09.
//
// Solidity: function getCompletionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetCompletionTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getCompletionTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompletionTimeout is a free data retrieval call binding the contract method 0xededfc09.
//
// Solidity: function getCompletionTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetCompletionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetCompletionTimeout(&_AIConfig.CallOpts)
}

// GetCompletionTimeout is a free data retrieval call binding the contract method 0xededfc09.
//
// Solidity: function getCompletionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetCompletionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetCompletionTimeout(&_AIConfig.CallOpts)
}

// GetCompletionTimeoutSlashBps is a free data retrieval call binding the contract method 0xbea04965.
//
// Solidity: function getCompletionTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetCompletionTimeoutSlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getCompletionTimeoutSlashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompletionTimeoutSlashBps is a free data retrieval call binding the contract method 0xbea04965.
//
// Solidity: function getCompletionTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetCompletionTimeoutSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetCompletionTimeoutSlashBps(&_AIConfig.CallOpts)
}

// GetCompletionTimeoutSlashBps is a free data retrieval call binding the contract method 0xbea04965.
//
// Solidity: function getCompletionTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetCompletionTimeoutSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetCompletionTimeoutSlashBps(&_AIConfig.CallOpts)
}

// GetDispatcherAddress is a free data retrieval call binding the contract method 0x75766e25.
//
// Solidity: function getDispatcherAddress() view returns(address)
func (_AIConfig *AIConfigCaller) GetDispatcherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getDispatcherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDispatcherAddress is a free data retrieval call binding the contract method 0x75766e25.
//
// Solidity: function getDispatcherAddress() view returns(address)
func (_AIConfig *AIConfigSession) GetDispatcherAddress() (common.Address, error) {
	return _AIConfig.Contract.GetDispatcherAddress(&_AIConfig.CallOpts)
}

// GetDispatcherAddress is a free data retrieval call binding the contract method 0x75766e25.
//
// Solidity: function getDispatcherAddress() view returns(address)
func (_AIConfig *AIConfigCallerSession) GetDispatcherAddress() (common.Address, error) {
	return _AIConfig.Contract.GetDispatcherAddress(&_AIConfig.CallOpts)
}

// GetDisputeBondMultiplier is a free data retrieval call binding the contract method 0xc86ba3fa.
//
// Solidity: function getDisputeBondMultiplier() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetDisputeBondMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getDisputeBondMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeBondMultiplier is a free data retrieval call binding the contract method 0xc86ba3fa.
//
// Solidity: function getDisputeBondMultiplier() view returns(uint256)
func (_AIConfig *AIConfigSession) GetDisputeBondMultiplier() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeBondMultiplier(&_AIConfig.CallOpts)
}

// GetDisputeBondMultiplier is a free data retrieval call binding the contract method 0xc86ba3fa.
//
// Solidity: function getDisputeBondMultiplier() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetDisputeBondMultiplier() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeBondMultiplier(&_AIConfig.CallOpts)
}

// GetDisputeSlashBps is a free data retrieval call binding the contract method 0xcd47c55f.
//
// Solidity: function getDisputeSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetDisputeSlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getDisputeSlashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeSlashBps is a free data retrieval call binding the contract method 0xcd47c55f.
//
// Solidity: function getDisputeSlashBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetDisputeSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeSlashBps(&_AIConfig.CallOpts)
}

// GetDisputeSlashBps is a free data retrieval call binding the contract method 0xcd47c55f.
//
// Solidity: function getDisputeSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetDisputeSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeSlashBps(&_AIConfig.CallOpts)
}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x1e19082a.
//
// Solidity: function getDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetDisputeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getDisputeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x1e19082a.
//
// Solidity: function getDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigSession) GetDisputeWindow() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeWindow(&_AIConfig.CallOpts)
}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x1e19082a.
//
// Solidity: function getDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetDisputeWindow() (*big.Int, error) {
	return _AIConfig.Contract.GetDisputeWindow(&_AIConfig.CallOpts)
}

// GetDisputerAddress is a free data retrieval call binding the contract method 0xc973b66b.
//
// Solidity: function getDisputerAddress() view returns(address)
func (_AIConfig *AIConfigCaller) GetDisputerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getDisputerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDisputerAddress is a free data retrieval call binding the contract method 0xc973b66b.
//
// Solidity: function getDisputerAddress() view returns(address)
func (_AIConfig *AIConfigSession) GetDisputerAddress() (common.Address, error) {
	return _AIConfig.Contract.GetDisputerAddress(&_AIConfig.CallOpts)
}

// GetDisputerAddress is a free data retrieval call binding the contract method 0xc973b66b.
//
// Solidity: function getDisputerAddress() view returns(address)
func (_AIConfig *AIConfigCallerSession) GetDisputerAddress() (common.Address, error) {
	return _AIConfig.Contract.GetDisputerAddress(&_AIConfig.CallOpts)
}

// GetMaxReassignments is a free data retrieval call binding the contract method 0x27ad5f34.
//
// Solidity: function getMaxReassignments() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMaxReassignments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMaxReassignments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxReassignments is a free data retrieval call binding the contract method 0x27ad5f34.
//
// Solidity: function getMaxReassignments() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMaxReassignments() (*big.Int, error) {
	return _AIConfig.Contract.GetMaxReassignments(&_AIConfig.CallOpts)
}

// GetMaxReassignments is a free data retrieval call binding the contract method 0x27ad5f34.
//
// Solidity: function getMaxReassignments() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMaxReassignments() (*big.Int, error) {
	return _AIConfig.Contract.GetMaxReassignments(&_AIConfig.CallOpts)
}

// GetMaxSlashBps is a free data retrieval call binding the contract method 0xb78f9b8c.
//
// Solidity: function getMaxSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMaxSlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMaxSlashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSlashBps is a free data retrieval call binding the contract method 0xb78f9b8c.
//
// Solidity: function getMaxSlashBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMaxSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetMaxSlashBps(&_AIConfig.CallOpts)
}

// GetMaxSlashBps is a free data retrieval call binding the contract method 0xb78f9b8c.
//
// Solidity: function getMaxSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMaxSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetMaxSlashBps(&_AIConfig.CallOpts)
}

// GetMinAckTimeout is a free data retrieval call binding the contract method 0x398b902d.
//
// Solidity: function getMinAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMinAckTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMinAckTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinAckTimeout is a free data retrieval call binding the contract method 0x398b902d.
//
// Solidity: function getMinAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMinAckTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetMinAckTimeout(&_AIConfig.CallOpts)
}

// GetMinAckTimeout is a free data retrieval call binding the contract method 0x398b902d.
//
// Solidity: function getMinAckTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMinAckTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetMinAckTimeout(&_AIConfig.CallOpts)
}

// GetMinDisputeWindow is a free data retrieval call binding the contract method 0x0cf6e55d.
//
// Solidity: function getMinDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMinDisputeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMinDisputeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDisputeWindow is a free data retrieval call binding the contract method 0x0cf6e55d.
//
// Solidity: function getMinDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMinDisputeWindow() (*big.Int, error) {
	return _AIConfig.Contract.GetMinDisputeWindow(&_AIConfig.CallOpts)
}

// GetMinDisputeWindow is a free data retrieval call binding the contract method 0x0cf6e55d.
//
// Solidity: function getMinDisputeWindow() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMinDisputeWindow() (*big.Int, error) {
	return _AIConfig.Contract.GetMinDisputeWindow(&_AIConfig.CallOpts)
}

// GetMinResolutionTimeout is a free data retrieval call binding the contract method 0x227545cc.
//
// Solidity: function getMinResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMinResolutionTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMinResolutionTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinResolutionTimeout is a free data retrieval call binding the contract method 0x227545cc.
//
// Solidity: function getMinResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMinResolutionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetMinResolutionTimeout(&_AIConfig.CallOpts)
}

// GetMinResolutionTimeout is a free data retrieval call binding the contract method 0x227545cc.
//
// Solidity: function getMinResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMinResolutionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetMinResolutionTimeout(&_AIConfig.CallOpts)
}

// GetMinWorkerStake is a free data retrieval call binding the contract method 0xca22dfd1.
//
// Solidity: function getMinWorkerStake() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetMinWorkerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getMinWorkerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinWorkerStake is a free data retrieval call binding the contract method 0xca22dfd1.
//
// Solidity: function getMinWorkerStake() view returns(uint256)
func (_AIConfig *AIConfigSession) GetMinWorkerStake() (*big.Int, error) {
	return _AIConfig.Contract.GetMinWorkerStake(&_AIConfig.CallOpts)
}

// GetMinWorkerStake is a free data retrieval call binding the contract method 0xca22dfd1.
//
// Solidity: function getMinWorkerStake() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetMinWorkerStake() (*big.Int, error) {
	return _AIConfig.Contract.GetMinWorkerStake(&_AIConfig.CallOpts)
}

// GetModelFee is a free data retrieval call binding the contract method 0xcbee2004.
//
// Solidity: function getModelFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCaller) GetModelFee(opts *bind.CallOpts, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getModelFee", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModelFee is a free data retrieval call binding the contract method 0xcbee2004.
//
// Solidity: function getModelFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigSession) GetModelFee(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.GetModelFee(&_AIConfig.CallOpts, modelId)
}

// GetModelFee is a free data retrieval call binding the contract method 0xcbee2004.
//
// Solidity: function getModelFee(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetModelFee(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.GetModelFee(&_AIConfig.CallOpts, modelId)
}

// GetModelMaxOutputTokens is a free data retrieval call binding the contract method 0x1678f85b.
//
// Solidity: function getModelMaxOutputTokens(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCaller) GetModelMaxOutputTokens(opts *bind.CallOpts, modelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getModelMaxOutputTokens", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModelMaxOutputTokens is a free data retrieval call binding the contract method 0x1678f85b.
//
// Solidity: function getModelMaxOutputTokens(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigSession) GetModelMaxOutputTokens(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.GetModelMaxOutputTokens(&_AIConfig.CallOpts, modelId)
}

// GetModelMaxOutputTokens is a free data retrieval call binding the contract method 0x1678f85b.
//
// Solidity: function getModelMaxOutputTokens(bytes32 modelId) view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetModelMaxOutputTokens(modelId [32]byte) (*big.Int, error) {
	return _AIConfig.Contract.GetModelMaxOutputTokens(&_AIConfig.CallOpts, modelId)
}

// GetProtocolFeeBps is a free data retrieval call binding the contract method 0xe590cea5.
//
// Solidity: function getProtocolFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetProtocolFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getProtocolFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeeBps is a free data retrieval call binding the contract method 0xe590cea5.
//
// Solidity: function getProtocolFeeBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetProtocolFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetProtocolFeeBps(&_AIConfig.CallOpts)
}

// GetProtocolFeeBps is a free data retrieval call binding the contract method 0xe590cea5.
//
// Solidity: function getProtocolFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetProtocolFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetProtocolFeeBps(&_AIConfig.CallOpts)
}

// GetResolutionTimeout is a free data retrieval call binding the contract method 0xa908ce4a.
//
// Solidity: function getResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetResolutionTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getResolutionTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResolutionTimeout is a free data retrieval call binding the contract method 0xa908ce4a.
//
// Solidity: function getResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetResolutionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetResolutionTimeout(&_AIConfig.CallOpts)
}

// GetResolutionTimeout is a free data retrieval call binding the contract method 0xa908ce4a.
//
// Solidity: function getResolutionTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetResolutionTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetResolutionTimeout(&_AIConfig.CallOpts)
}

// GetSamplingRateBps is a free data retrieval call binding the contract method 0x63c3f809.
//
// Solidity: function getSamplingRateBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetSamplingRateBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getSamplingRateBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSamplingRateBps is a free data retrieval call binding the contract method 0x63c3f809.
//
// Solidity: function getSamplingRateBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetSamplingRateBps() (*big.Int, error) {
	return _AIConfig.Contract.GetSamplingRateBps(&_AIConfig.CallOpts)
}

// GetSamplingRateBps is a free data retrieval call binding the contract method 0x63c3f809.
//
// Solidity: function getSamplingRateBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetSamplingRateBps() (*big.Int, error) {
	return _AIConfig.Contract.GetSamplingRateBps(&_AIConfig.CallOpts)
}

// GetSessionInactivityTimeout is a free data retrieval call binding the contract method 0xa388cdb4.
//
// Solidity: function getSessionInactivityTimeout() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetSessionInactivityTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getSessionInactivityTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSessionInactivityTimeout is a free data retrieval call binding the contract method 0xa388cdb4.
//
// Solidity: function getSessionInactivityTimeout() view returns(uint256)
func (_AIConfig *AIConfigSession) GetSessionInactivityTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetSessionInactivityTimeout(&_AIConfig.CallOpts)
}

// GetSessionInactivityTimeout is a free data retrieval call binding the contract method 0xa388cdb4.
//
// Solidity: function getSessionInactivityTimeout() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetSessionInactivityTimeout() (*big.Int, error) {
	return _AIConfig.Contract.GetSessionInactivityTimeout(&_AIConfig.CallOpts)
}

// GetSimilarityThreshold is a free data retrieval call binding the contract method 0xcaf94095.
//
// Solidity: function getSimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetSimilarityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getSimilarityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSimilarityThreshold is a free data retrieval call binding the contract method 0xcaf94095.
//
// Solidity: function getSimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigSession) GetSimilarityThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetSimilarityThreshold(&_AIConfig.CallOpts)
}

// GetSimilarityThreshold is a free data retrieval call binding the contract method 0xcaf94095.
//
// Solidity: function getSimilarityThreshold() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetSimilarityThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetSimilarityThreshold(&_AIConfig.CallOpts)
}

// GetSuspensionCooldown is a free data retrieval call binding the contract method 0x7b368ecb.
//
// Solidity: function getSuspensionCooldown() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetSuspensionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getSuspensionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuspensionCooldown is a free data retrieval call binding the contract method 0x7b368ecb.
//
// Solidity: function getSuspensionCooldown() view returns(uint256)
func (_AIConfig *AIConfigSession) GetSuspensionCooldown() (*big.Int, error) {
	return _AIConfig.Contract.GetSuspensionCooldown(&_AIConfig.CallOpts)
}

// GetSuspensionCooldown is a free data retrieval call binding the contract method 0x7b368ecb.
//
// Solidity: function getSuspensionCooldown() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetSuspensionCooldown() (*big.Int, error) {
	return _AIConfig.Contract.GetSuspensionCooldown(&_AIConfig.CallOpts)
}

// GetSuspensionThreshold is a free data retrieval call binding the contract method 0x1498389b.
//
// Solidity: function getSuspensionThreshold() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetSuspensionThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getSuspensionThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuspensionThreshold is a free data retrieval call binding the contract method 0x1498389b.
//
// Solidity: function getSuspensionThreshold() view returns(uint256)
func (_AIConfig *AIConfigSession) GetSuspensionThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetSuspensionThreshold(&_AIConfig.CallOpts)
}

// GetSuspensionThreshold is a free data retrieval call binding the contract method 0x1498389b.
//
// Solidity: function getSuspensionThreshold() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetSuspensionThreshold() (*big.Int, error) {
	return _AIConfig.Contract.GetSuspensionThreshold(&_AIConfig.CallOpts)
}

// GetTimeoutSlashBps is a free data retrieval call binding the contract method 0x01018a05.
//
// Solidity: function getTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetTimeoutSlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getTimeoutSlashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeoutSlashBps is a free data retrieval call binding the contract method 0x01018a05.
//
// Solidity: function getTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetTimeoutSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetTimeoutSlashBps(&_AIConfig.CallOpts)
}

// GetTimeoutSlashBps is a free data retrieval call binding the contract method 0x01018a05.
//
// Solidity: function getTimeoutSlashBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetTimeoutSlashBps() (*big.Int, error) {
	return _AIConfig.Contract.GetTimeoutSlashBps(&_AIConfig.CallOpts)
}

// GetWorkerFeeBps is a free data retrieval call binding the contract method 0x549787e2.
//
// Solidity: function getWorkerFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCaller) GetWorkerFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "getWorkerFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkerFeeBps is a free data retrieval call binding the contract method 0x549787e2.
//
// Solidity: function getWorkerFeeBps() view returns(uint256)
func (_AIConfig *AIConfigSession) GetWorkerFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetWorkerFeeBps(&_AIConfig.CallOpts)
}

// GetWorkerFeeBps is a free data retrieval call binding the contract method 0x549787e2.
//
// Solidity: function getWorkerFeeBps() view returns(uint256)
func (_AIConfig *AIConfigCallerSession) GetWorkerFeeBps() (*big.Int, error) {
	return _AIConfig.Contract.GetWorkerFeeBps(&_AIConfig.CallOpts)
}

// IsModelEnabled is a free data retrieval call binding the contract method 0x2b99ec4d.
//
// Solidity: function isModelEnabled(bytes32 modelId) view returns(bool)
func (_AIConfig *AIConfigCaller) IsModelEnabled(opts *bind.CallOpts, modelId [32]byte) (bool, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "isModelEnabled", modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModelEnabled is a free data retrieval call binding the contract method 0x2b99ec4d.
//
// Solidity: function isModelEnabled(bytes32 modelId) view returns(bool)
func (_AIConfig *AIConfigSession) IsModelEnabled(modelId [32]byte) (bool, error) {
	return _AIConfig.Contract.IsModelEnabled(&_AIConfig.CallOpts, modelId)
}

// IsModelEnabled is a free data retrieval call binding the contract method 0x2b99ec4d.
//
// Solidity: function isModelEnabled(bytes32 modelId) view returns(bool)
func (_AIConfig *AIConfigCallerSession) IsModelEnabled(modelId [32]byte) (bool, error) {
	return _AIConfig.Contract.IsModelEnabled(&_AIConfig.CallOpts, modelId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIConfig *AIConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIConfig *AIConfigSession) Owner() (common.Address, error) {
	return _AIConfig.Contract.Owner(&_AIConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIConfig *AIConfigCallerSession) Owner() (common.Address, error) {
	return _AIConfig.Contract.Owner(&_AIConfig.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AIConfig *AIConfigCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AIConfig.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AIConfig *AIConfigSession) ProxiableUUID() ([32]byte, error) {
	return _AIConfig.Contract.ProxiableUUID(&_AIConfig.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AIConfig *AIConfigCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AIConfig.Contract.ProxiableUUID(&_AIConfig.CallOpts)
}

// DisableModel is a paid mutator transaction binding the contract method 0x70bca450.
//
// Solidity: function disableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigTransactor) DisableModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "disableModel", modelId)
}

// DisableModel is a paid mutator transaction binding the contract method 0x70bca450.
//
// Solidity: function disableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigSession) DisableModel(modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.Contract.DisableModel(&_AIConfig.TransactOpts, modelId)
}

// DisableModel is a paid mutator transaction binding the contract method 0x70bca450.
//
// Solidity: function disableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigTransactorSession) DisableModel(modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.Contract.DisableModel(&_AIConfig.TransactOpts, modelId)
}

// EnableModel is a paid mutator transaction binding the contract method 0xe631bdb1.
//
// Solidity: function enableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigTransactor) EnableModel(opts *bind.TransactOpts, modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "enableModel", modelId)
}

// EnableModel is a paid mutator transaction binding the contract method 0xe631bdb1.
//
// Solidity: function enableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigSession) EnableModel(modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.Contract.EnableModel(&_AIConfig.TransactOpts, modelId)
}

// EnableModel is a paid mutator transaction binding the contract method 0xe631bdb1.
//
// Solidity: function enableModel(bytes32 modelId) returns()
func (_AIConfig *AIConfigTransactorSession) EnableModel(modelId [32]byte) (*types.Transaction, error) {
	return _AIConfig.Contract.EnableModel(&_AIConfig.TransactOpts, modelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _dispatcher, address _disputer) returns()
func (_AIConfig *AIConfigTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _dispatcher common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "initialize", _initialOwner, _dispatcher, _disputer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _dispatcher, address _disputer) returns()
func (_AIConfig *AIConfigSession) Initialize(_initialOwner common.Address, _dispatcher common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.Initialize(&_AIConfig.TransactOpts, _initialOwner, _dispatcher, _disputer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _dispatcher, address _disputer) returns()
func (_AIConfig *AIConfigTransactorSession) Initialize(_initialOwner common.Address, _dispatcher common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.Initialize(&_AIConfig.TransactOpts, _initialOwner, _dispatcher, _disputer)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xed842e5d.
//
// Solidity: function registerModel(bytes32 modelId, uint256 fee, uint256 maxOutputTokens) returns()
func (_AIConfig *AIConfigTransactor) RegisterModel(opts *bind.TransactOpts, modelId [32]byte, fee *big.Int, maxOutputTokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "registerModel", modelId, fee, maxOutputTokens)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xed842e5d.
//
// Solidity: function registerModel(bytes32 modelId, uint256 fee, uint256 maxOutputTokens) returns()
func (_AIConfig *AIConfigSession) RegisterModel(modelId [32]byte, fee *big.Int, maxOutputTokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.RegisterModel(&_AIConfig.TransactOpts, modelId, fee, maxOutputTokens)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xed842e5d.
//
// Solidity: function registerModel(bytes32 modelId, uint256 fee, uint256 maxOutputTokens) returns()
func (_AIConfig *AIConfigTransactorSession) RegisterModel(modelId [32]byte, fee *big.Int, maxOutputTokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.RegisterModel(&_AIConfig.TransactOpts, modelId, fee, maxOutputTokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIConfig *AIConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIConfig *AIConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _AIConfig.Contract.RenounceOwnership(&_AIConfig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIConfig *AIConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AIConfig.Contract.RenounceOwnership(&_AIConfig.TransactOpts)
}

// SetAckTimeout is a paid mutator transaction binding the contract method 0xc8be8ab3.
//
// Solidity: function setAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetAckTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setAckTimeout", timeout)
}

// SetAckTimeout is a paid mutator transaction binding the contract method 0xc8be8ab3.
//
// Solidity: function setAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetAckTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetAckTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetAckTimeout is a paid mutator transaction binding the contract method 0xc8be8ab3.
//
// Solidity: function setAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetAckTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetAckTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetBlobRetentionPeriod is a paid mutator transaction binding the contract method 0x9855a76f.
//
// Solidity: function setBlobRetentionPeriod(uint256 period) returns()
func (_AIConfig *AIConfigTransactor) SetBlobRetentionPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setBlobRetentionPeriod", period)
}

// SetBlobRetentionPeriod is a paid mutator transaction binding the contract method 0x9855a76f.
//
// Solidity: function setBlobRetentionPeriod(uint256 period) returns()
func (_AIConfig *AIConfigSession) SetBlobRetentionPeriod(period *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetBlobRetentionPeriod(&_AIConfig.TransactOpts, period)
}

// SetBlobRetentionPeriod is a paid mutator transaction binding the contract method 0x9855a76f.
//
// Solidity: function setBlobRetentionPeriod(uint256 period) returns()
func (_AIConfig *AIConfigTransactorSession) SetBlobRetentionPeriod(period *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetBlobRetentionPeriod(&_AIConfig.TransactOpts, period)
}

// SetCanarySimilarityThreshold is a paid mutator transaction binding the contract method 0x6116a049.
//
// Solidity: function setCanarySimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetCanarySimilarityThreshold(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setCanarySimilarityThreshold", bps)
}

// SetCanarySimilarityThreshold is a paid mutator transaction binding the contract method 0x6116a049.
//
// Solidity: function setCanarySimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetCanarySimilarityThreshold(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCanarySimilarityThreshold(&_AIConfig.TransactOpts, bps)
}

// SetCanarySimilarityThreshold is a paid mutator transaction binding the contract method 0x6116a049.
//
// Solidity: function setCanarySimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetCanarySimilarityThreshold(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCanarySimilarityThreshold(&_AIConfig.TransactOpts, bps)
}

// SetCompletionTimeout is a paid mutator transaction binding the contract method 0x5bf54104.
//
// Solidity: function setCompletionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetCompletionTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setCompletionTimeout", timeout)
}

// SetCompletionTimeout is a paid mutator transaction binding the contract method 0x5bf54104.
//
// Solidity: function setCompletionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetCompletionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCompletionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetCompletionTimeout is a paid mutator transaction binding the contract method 0x5bf54104.
//
// Solidity: function setCompletionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetCompletionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCompletionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetCompletionTimeoutSlashBps is a paid mutator transaction binding the contract method 0x200ea958.
//
// Solidity: function setCompletionTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetCompletionTimeoutSlashBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setCompletionTimeoutSlashBps", bps)
}

// SetCompletionTimeoutSlashBps is a paid mutator transaction binding the contract method 0x200ea958.
//
// Solidity: function setCompletionTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetCompletionTimeoutSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCompletionTimeoutSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetCompletionTimeoutSlashBps is a paid mutator transaction binding the contract method 0x200ea958.
//
// Solidity: function setCompletionTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetCompletionTimeoutSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetCompletionTimeoutSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetDispatcherAddress is a paid mutator transaction binding the contract method 0xafaebf3e.
//
// Solidity: function setDispatcherAddress(address dispatcher) returns()
func (_AIConfig *AIConfigTransactor) SetDispatcherAddress(opts *bind.TransactOpts, dispatcher common.Address) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setDispatcherAddress", dispatcher)
}

// SetDispatcherAddress is a paid mutator transaction binding the contract method 0xafaebf3e.
//
// Solidity: function setDispatcherAddress(address dispatcher) returns()
func (_AIConfig *AIConfigSession) SetDispatcherAddress(dispatcher common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDispatcherAddress(&_AIConfig.TransactOpts, dispatcher)
}

// SetDispatcherAddress is a paid mutator transaction binding the contract method 0xafaebf3e.
//
// Solidity: function setDispatcherAddress(address dispatcher) returns()
func (_AIConfig *AIConfigTransactorSession) SetDispatcherAddress(dispatcher common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDispatcherAddress(&_AIConfig.TransactOpts, dispatcher)
}

// SetDisputeBondMultiplier is a paid mutator transaction binding the contract method 0x7f50ff5d.
//
// Solidity: function setDisputeBondMultiplier(uint256 multiplier) returns()
func (_AIConfig *AIConfigTransactor) SetDisputeBondMultiplier(opts *bind.TransactOpts, multiplier *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setDisputeBondMultiplier", multiplier)
}

// SetDisputeBondMultiplier is a paid mutator transaction binding the contract method 0x7f50ff5d.
//
// Solidity: function setDisputeBondMultiplier(uint256 multiplier) returns()
func (_AIConfig *AIConfigSession) SetDisputeBondMultiplier(multiplier *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeBondMultiplier(&_AIConfig.TransactOpts, multiplier)
}

// SetDisputeBondMultiplier is a paid mutator transaction binding the contract method 0x7f50ff5d.
//
// Solidity: function setDisputeBondMultiplier(uint256 multiplier) returns()
func (_AIConfig *AIConfigTransactorSession) SetDisputeBondMultiplier(multiplier *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeBondMultiplier(&_AIConfig.TransactOpts, multiplier)
}

// SetDisputeSlashBps is a paid mutator transaction binding the contract method 0x293a3ba3.
//
// Solidity: function setDisputeSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetDisputeSlashBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setDisputeSlashBps", bps)
}

// SetDisputeSlashBps is a paid mutator transaction binding the contract method 0x293a3ba3.
//
// Solidity: function setDisputeSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetDisputeSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetDisputeSlashBps is a paid mutator transaction binding the contract method 0x293a3ba3.
//
// Solidity: function setDisputeSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetDisputeSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigTransactor) SetDisputeWindow(opts *bind.TransactOpts, window *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setDisputeWindow", window)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigSession) SetDisputeWindow(window *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeWindow(&_AIConfig.TransactOpts, window)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigTransactorSession) SetDisputeWindow(window *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputeWindow(&_AIConfig.TransactOpts, window)
}

// SetDisputerAddress is a paid mutator transaction binding the contract method 0x6d2baf84.
//
// Solidity: function setDisputerAddress(address disputer) returns()
func (_AIConfig *AIConfigTransactor) SetDisputerAddress(opts *bind.TransactOpts, disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setDisputerAddress", disputer)
}

// SetDisputerAddress is a paid mutator transaction binding the contract method 0x6d2baf84.
//
// Solidity: function setDisputerAddress(address disputer) returns()
func (_AIConfig *AIConfigSession) SetDisputerAddress(disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputerAddress(&_AIConfig.TransactOpts, disputer)
}

// SetDisputerAddress is a paid mutator transaction binding the contract method 0x6d2baf84.
//
// Solidity: function setDisputerAddress(address disputer) returns()
func (_AIConfig *AIConfigTransactorSession) SetDisputerAddress(disputer common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.SetDisputerAddress(&_AIConfig.TransactOpts, disputer)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xba8e568f.
//
// Solidity: function setFeeDistribution(uint256 workerBps, uint256 protocolBps, uint256 burnBps) returns()
func (_AIConfig *AIConfigTransactor) SetFeeDistribution(opts *bind.TransactOpts, workerBps *big.Int, protocolBps *big.Int, burnBps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setFeeDistribution", workerBps, protocolBps, burnBps)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xba8e568f.
//
// Solidity: function setFeeDistribution(uint256 workerBps, uint256 protocolBps, uint256 burnBps) returns()
func (_AIConfig *AIConfigSession) SetFeeDistribution(workerBps *big.Int, protocolBps *big.Int, burnBps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetFeeDistribution(&_AIConfig.TransactOpts, workerBps, protocolBps, burnBps)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0xba8e568f.
//
// Solidity: function setFeeDistribution(uint256 workerBps, uint256 protocolBps, uint256 burnBps) returns()
func (_AIConfig *AIConfigTransactorSession) SetFeeDistribution(workerBps *big.Int, protocolBps *big.Int, burnBps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetFeeDistribution(&_AIConfig.TransactOpts, workerBps, protocolBps, burnBps)
}

// SetMaxReassignments is a paid mutator transaction binding the contract method 0xc35400f1.
//
// Solidity: function setMaxReassignments(uint256 max) returns()
func (_AIConfig *AIConfigTransactor) SetMaxReassignments(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMaxReassignments", max)
}

// SetMaxReassignments is a paid mutator transaction binding the contract method 0xc35400f1.
//
// Solidity: function setMaxReassignments(uint256 max) returns()
func (_AIConfig *AIConfigSession) SetMaxReassignments(max *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMaxReassignments(&_AIConfig.TransactOpts, max)
}

// SetMaxReassignments is a paid mutator transaction binding the contract method 0xc35400f1.
//
// Solidity: function setMaxReassignments(uint256 max) returns()
func (_AIConfig *AIConfigTransactorSession) SetMaxReassignments(max *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMaxReassignments(&_AIConfig.TransactOpts, max)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetMaxSlashBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMaxSlashBps", bps)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetMaxSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMaxSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetMaxSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMaxSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetMinAckTimeout is a paid mutator transaction binding the contract method 0x9d47c6c8.
//
// Solidity: function setMinAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetMinAckTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMinAckTimeout", timeout)
}

// SetMinAckTimeout is a paid mutator transaction binding the contract method 0x9d47c6c8.
//
// Solidity: function setMinAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetMinAckTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinAckTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetMinAckTimeout is a paid mutator transaction binding the contract method 0x9d47c6c8.
//
// Solidity: function setMinAckTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetMinAckTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinAckTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetMinDisputeWindow is a paid mutator transaction binding the contract method 0x6059283a.
//
// Solidity: function setMinDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigTransactor) SetMinDisputeWindow(opts *bind.TransactOpts, window *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMinDisputeWindow", window)
}

// SetMinDisputeWindow is a paid mutator transaction binding the contract method 0x6059283a.
//
// Solidity: function setMinDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigSession) SetMinDisputeWindow(window *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinDisputeWindow(&_AIConfig.TransactOpts, window)
}

// SetMinDisputeWindow is a paid mutator transaction binding the contract method 0x6059283a.
//
// Solidity: function setMinDisputeWindow(uint256 window) returns()
func (_AIConfig *AIConfigTransactorSession) SetMinDisputeWindow(window *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinDisputeWindow(&_AIConfig.TransactOpts, window)
}

// SetMinResolutionTimeout is a paid mutator transaction binding the contract method 0xf1189366.
//
// Solidity: function setMinResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetMinResolutionTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMinResolutionTimeout", timeout)
}

// SetMinResolutionTimeout is a paid mutator transaction binding the contract method 0xf1189366.
//
// Solidity: function setMinResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetMinResolutionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinResolutionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetMinResolutionTimeout is a paid mutator transaction binding the contract method 0xf1189366.
//
// Solidity: function setMinResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetMinResolutionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinResolutionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetMinWorkerStake is a paid mutator transaction binding the contract method 0x3cf62d1b.
//
// Solidity: function setMinWorkerStake(uint256 stake) returns()
func (_AIConfig *AIConfigTransactor) SetMinWorkerStake(opts *bind.TransactOpts, stake *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setMinWorkerStake", stake)
}

// SetMinWorkerStake is a paid mutator transaction binding the contract method 0x3cf62d1b.
//
// Solidity: function setMinWorkerStake(uint256 stake) returns()
func (_AIConfig *AIConfigSession) SetMinWorkerStake(stake *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinWorkerStake(&_AIConfig.TransactOpts, stake)
}

// SetMinWorkerStake is a paid mutator transaction binding the contract method 0x3cf62d1b.
//
// Solidity: function setMinWorkerStake(uint256 stake) returns()
func (_AIConfig *AIConfigTransactorSession) SetMinWorkerStake(stake *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetMinWorkerStake(&_AIConfig.TransactOpts, stake)
}

// SetModelFee is a paid mutator transaction binding the contract method 0xf029ac3e.
//
// Solidity: function setModelFee(bytes32 modelId, uint256 fee) returns()
func (_AIConfig *AIConfigTransactor) SetModelFee(opts *bind.TransactOpts, modelId [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setModelFee", modelId, fee)
}

// SetModelFee is a paid mutator transaction binding the contract method 0xf029ac3e.
//
// Solidity: function setModelFee(bytes32 modelId, uint256 fee) returns()
func (_AIConfig *AIConfigSession) SetModelFee(modelId [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetModelFee(&_AIConfig.TransactOpts, modelId, fee)
}

// SetModelFee is a paid mutator transaction binding the contract method 0xf029ac3e.
//
// Solidity: function setModelFee(bytes32 modelId, uint256 fee) returns()
func (_AIConfig *AIConfigTransactorSession) SetModelFee(modelId [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetModelFee(&_AIConfig.TransactOpts, modelId, fee)
}

// SetModelMaxOutputTokens is a paid mutator transaction binding the contract method 0x14953714.
//
// Solidity: function setModelMaxOutputTokens(bytes32 modelId, uint256 tokens) returns()
func (_AIConfig *AIConfigTransactor) SetModelMaxOutputTokens(opts *bind.TransactOpts, modelId [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setModelMaxOutputTokens", modelId, tokens)
}

// SetModelMaxOutputTokens is a paid mutator transaction binding the contract method 0x14953714.
//
// Solidity: function setModelMaxOutputTokens(bytes32 modelId, uint256 tokens) returns()
func (_AIConfig *AIConfigSession) SetModelMaxOutputTokens(modelId [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetModelMaxOutputTokens(&_AIConfig.TransactOpts, modelId, tokens)
}

// SetModelMaxOutputTokens is a paid mutator transaction binding the contract method 0x14953714.
//
// Solidity: function setModelMaxOutputTokens(bytes32 modelId, uint256 tokens) returns()
func (_AIConfig *AIConfigTransactorSession) SetModelMaxOutputTokens(modelId [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetModelMaxOutputTokens(&_AIConfig.TransactOpts, modelId, tokens)
}

// SetResolutionTimeout is a paid mutator transaction binding the contract method 0xdb625232.
//
// Solidity: function setResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetResolutionTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setResolutionTimeout", timeout)
}

// SetResolutionTimeout is a paid mutator transaction binding the contract method 0xdb625232.
//
// Solidity: function setResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetResolutionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetResolutionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetResolutionTimeout is a paid mutator transaction binding the contract method 0xdb625232.
//
// Solidity: function setResolutionTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetResolutionTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetResolutionTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetSamplingRateBps is a paid mutator transaction binding the contract method 0x9269681c.
//
// Solidity: function setSamplingRateBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetSamplingRateBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setSamplingRateBps", bps)
}

// SetSamplingRateBps is a paid mutator transaction binding the contract method 0x9269681c.
//
// Solidity: function setSamplingRateBps(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetSamplingRateBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSamplingRateBps(&_AIConfig.TransactOpts, bps)
}

// SetSamplingRateBps is a paid mutator transaction binding the contract method 0x9269681c.
//
// Solidity: function setSamplingRateBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetSamplingRateBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSamplingRateBps(&_AIConfig.TransactOpts, bps)
}

// SetSessionInactivityTimeout is a paid mutator transaction binding the contract method 0xba758a25.
//
// Solidity: function setSessionInactivityTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactor) SetSessionInactivityTimeout(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setSessionInactivityTimeout", timeout)
}

// SetSessionInactivityTimeout is a paid mutator transaction binding the contract method 0xba758a25.
//
// Solidity: function setSessionInactivityTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigSession) SetSessionInactivityTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSessionInactivityTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetSessionInactivityTimeout is a paid mutator transaction binding the contract method 0xba758a25.
//
// Solidity: function setSessionInactivityTimeout(uint256 timeout) returns()
func (_AIConfig *AIConfigTransactorSession) SetSessionInactivityTimeout(timeout *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSessionInactivityTimeout(&_AIConfig.TransactOpts, timeout)
}

// SetSimilarityThreshold is a paid mutator transaction binding the contract method 0xd3a64c63.
//
// Solidity: function setSimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetSimilarityThreshold(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setSimilarityThreshold", bps)
}

// SetSimilarityThreshold is a paid mutator transaction binding the contract method 0xd3a64c63.
//
// Solidity: function setSimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetSimilarityThreshold(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSimilarityThreshold(&_AIConfig.TransactOpts, bps)
}

// SetSimilarityThreshold is a paid mutator transaction binding the contract method 0xd3a64c63.
//
// Solidity: function setSimilarityThreshold(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetSimilarityThreshold(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSimilarityThreshold(&_AIConfig.TransactOpts, bps)
}

// SetSuspensionCooldown is a paid mutator transaction binding the contract method 0x0a623eaf.
//
// Solidity: function setSuspensionCooldown(uint256 duration) returns()
func (_AIConfig *AIConfigTransactor) SetSuspensionCooldown(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setSuspensionCooldown", duration)
}

// SetSuspensionCooldown is a paid mutator transaction binding the contract method 0x0a623eaf.
//
// Solidity: function setSuspensionCooldown(uint256 duration) returns()
func (_AIConfig *AIConfigSession) SetSuspensionCooldown(duration *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSuspensionCooldown(&_AIConfig.TransactOpts, duration)
}

// SetSuspensionCooldown is a paid mutator transaction binding the contract method 0x0a623eaf.
//
// Solidity: function setSuspensionCooldown(uint256 duration) returns()
func (_AIConfig *AIConfigTransactorSession) SetSuspensionCooldown(duration *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSuspensionCooldown(&_AIConfig.TransactOpts, duration)
}

// SetSuspensionThreshold is a paid mutator transaction binding the contract method 0xe52e2415.
//
// Solidity: function setSuspensionThreshold(uint256 count) returns()
func (_AIConfig *AIConfigTransactor) SetSuspensionThreshold(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setSuspensionThreshold", count)
}

// SetSuspensionThreshold is a paid mutator transaction binding the contract method 0xe52e2415.
//
// Solidity: function setSuspensionThreshold(uint256 count) returns()
func (_AIConfig *AIConfigSession) SetSuspensionThreshold(count *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSuspensionThreshold(&_AIConfig.TransactOpts, count)
}

// SetSuspensionThreshold is a paid mutator transaction binding the contract method 0xe52e2415.
//
// Solidity: function setSuspensionThreshold(uint256 count) returns()
func (_AIConfig *AIConfigTransactorSession) SetSuspensionThreshold(count *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetSuspensionThreshold(&_AIConfig.TransactOpts, count)
}

// SetTimeoutSlashBps is a paid mutator transaction binding the contract method 0xa1d38253.
//
// Solidity: function setTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactor) SetTimeoutSlashBps(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "setTimeoutSlashBps", bps)
}

// SetTimeoutSlashBps is a paid mutator transaction binding the contract method 0xa1d38253.
//
// Solidity: function setTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigSession) SetTimeoutSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetTimeoutSlashBps(&_AIConfig.TransactOpts, bps)
}

// SetTimeoutSlashBps is a paid mutator transaction binding the contract method 0xa1d38253.
//
// Solidity: function setTimeoutSlashBps(uint256 bps) returns()
func (_AIConfig *AIConfigTransactorSession) SetTimeoutSlashBps(bps *big.Int) (*types.Transaction, error) {
	return _AIConfig.Contract.SetTimeoutSlashBps(&_AIConfig.TransactOpts, bps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIConfig *AIConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIConfig *AIConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.TransferOwnership(&_AIConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIConfig *AIConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AIConfig.Contract.TransferOwnership(&_AIConfig.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AIConfig *AIConfigTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AIConfig.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AIConfig *AIConfigSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AIConfig.Contract.UpgradeToAndCall(&_AIConfig.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AIConfig *AIConfigTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AIConfig.Contract.UpgradeToAndCall(&_AIConfig.TransactOpts, newImplementation, data)
}

// AIConfigDispatcherAddressUpdatedIterator is returned from FilterDispatcherAddressUpdated and is used to iterate over the raw logs and unpacked data for DispatcherAddressUpdated events raised by the AIConfig contract.
type AIConfigDispatcherAddressUpdatedIterator struct {
	Event *AIConfigDispatcherAddressUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigDispatcherAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigDispatcherAddressUpdated)
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
		it.Event = new(AIConfigDispatcherAddressUpdated)
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
func (it *AIConfigDispatcherAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigDispatcherAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigDispatcherAddressUpdated represents a DispatcherAddressUpdated event raised by the AIConfig contract.
type AIConfigDispatcherAddressUpdated struct {
	NewDispatcher common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDispatcherAddressUpdated is a free log retrieval operation binding the contract event 0x20d0ae5feff505aaacb84857d5436350cdb9c9c0ff57bb2fe8611e298913df7f.
//
// Solidity: event DispatcherAddressUpdated(address indexed newDispatcher)
func (_AIConfig *AIConfigFilterer) FilterDispatcherAddressUpdated(opts *bind.FilterOpts, newDispatcher []common.Address) (*AIConfigDispatcherAddressUpdatedIterator, error) {

	var newDispatcherRule []interface{}
	for _, newDispatcherItem := range newDispatcher {
		newDispatcherRule = append(newDispatcherRule, newDispatcherItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "DispatcherAddressUpdated", newDispatcherRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigDispatcherAddressUpdatedIterator{contract: _AIConfig.contract, event: "DispatcherAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchDispatcherAddressUpdated is a free log subscription operation binding the contract event 0x20d0ae5feff505aaacb84857d5436350cdb9c9c0ff57bb2fe8611e298913df7f.
//
// Solidity: event DispatcherAddressUpdated(address indexed newDispatcher)
func (_AIConfig *AIConfigFilterer) WatchDispatcherAddressUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigDispatcherAddressUpdated, newDispatcher []common.Address) (event.Subscription, error) {

	var newDispatcherRule []interface{}
	for _, newDispatcherItem := range newDispatcher {
		newDispatcherRule = append(newDispatcherRule, newDispatcherItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "DispatcherAddressUpdated", newDispatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigDispatcherAddressUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "DispatcherAddressUpdated", log); err != nil {
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

// ParseDispatcherAddressUpdated is a log parse operation binding the contract event 0x20d0ae5feff505aaacb84857d5436350cdb9c9c0ff57bb2fe8611e298913df7f.
//
// Solidity: event DispatcherAddressUpdated(address indexed newDispatcher)
func (_AIConfig *AIConfigFilterer) ParseDispatcherAddressUpdated(log types.Log) (*AIConfigDispatcherAddressUpdated, error) {
	event := new(AIConfigDispatcherAddressUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "DispatcherAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigDisputeParamsUpdatedIterator is returned from FilterDisputeParamsUpdated and is used to iterate over the raw logs and unpacked data for DisputeParamsUpdated events raised by the AIConfig contract.
type AIConfigDisputeParamsUpdatedIterator struct {
	Event *AIConfigDisputeParamsUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigDisputeParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigDisputeParamsUpdated)
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
		it.Event = new(AIConfigDisputeParamsUpdated)
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
func (it *AIConfigDisputeParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigDisputeParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigDisputeParamsUpdated represents a DisputeParamsUpdated event raised by the AIConfig contract.
type AIConfigDisputeParamsUpdated struct {
	BondMultiplier    *big.Int
	Window            *big.Int
	ResolutionTimeout *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDisputeParamsUpdated is a free log retrieval operation binding the contract event 0x4a4c041415cbab9f7eb0c14aafcaf2ffcb7faac7936c5d695eef82504359c1b3.
//
// Solidity: event DisputeParamsUpdated(uint256 bondMultiplier, uint256 window, uint256 resolutionTimeout)
func (_AIConfig *AIConfigFilterer) FilterDisputeParamsUpdated(opts *bind.FilterOpts) (*AIConfigDisputeParamsUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "DisputeParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigDisputeParamsUpdatedIterator{contract: _AIConfig.contract, event: "DisputeParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeParamsUpdated is a free log subscription operation binding the contract event 0x4a4c041415cbab9f7eb0c14aafcaf2ffcb7faac7936c5d695eef82504359c1b3.
//
// Solidity: event DisputeParamsUpdated(uint256 bondMultiplier, uint256 window, uint256 resolutionTimeout)
func (_AIConfig *AIConfigFilterer) WatchDisputeParamsUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigDisputeParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "DisputeParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigDisputeParamsUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "DisputeParamsUpdated", log); err != nil {
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

// ParseDisputeParamsUpdated is a log parse operation binding the contract event 0x4a4c041415cbab9f7eb0c14aafcaf2ffcb7faac7936c5d695eef82504359c1b3.
//
// Solidity: event DisputeParamsUpdated(uint256 bondMultiplier, uint256 window, uint256 resolutionTimeout)
func (_AIConfig *AIConfigFilterer) ParseDisputeParamsUpdated(log types.Log) (*AIConfigDisputeParamsUpdated, error) {
	event := new(AIConfigDisputeParamsUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "DisputeParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigDisputerAddressUpdatedIterator is returned from FilterDisputerAddressUpdated and is used to iterate over the raw logs and unpacked data for DisputerAddressUpdated events raised by the AIConfig contract.
type AIConfigDisputerAddressUpdatedIterator struct {
	Event *AIConfigDisputerAddressUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigDisputerAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigDisputerAddressUpdated)
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
		it.Event = new(AIConfigDisputerAddressUpdated)
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
func (it *AIConfigDisputerAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigDisputerAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigDisputerAddressUpdated represents a DisputerAddressUpdated event raised by the AIConfig contract.
type AIConfigDisputerAddressUpdated struct {
	NewDisputer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDisputerAddressUpdated is a free log retrieval operation binding the contract event 0xdb6c994f490543a0786e32e55707cbdba4ced511c54b42ef571b510e78c07798.
//
// Solidity: event DisputerAddressUpdated(address indexed newDisputer)
func (_AIConfig *AIConfigFilterer) FilterDisputerAddressUpdated(opts *bind.FilterOpts, newDisputer []common.Address) (*AIConfigDisputerAddressUpdatedIterator, error) {

	var newDisputerRule []interface{}
	for _, newDisputerItem := range newDisputer {
		newDisputerRule = append(newDisputerRule, newDisputerItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "DisputerAddressUpdated", newDisputerRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigDisputerAddressUpdatedIterator{contract: _AIConfig.contract, event: "DisputerAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputerAddressUpdated is a free log subscription operation binding the contract event 0xdb6c994f490543a0786e32e55707cbdba4ced511c54b42ef571b510e78c07798.
//
// Solidity: event DisputerAddressUpdated(address indexed newDisputer)
func (_AIConfig *AIConfigFilterer) WatchDisputerAddressUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigDisputerAddressUpdated, newDisputer []common.Address) (event.Subscription, error) {

	var newDisputerRule []interface{}
	for _, newDisputerItem := range newDisputer {
		newDisputerRule = append(newDisputerRule, newDisputerItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "DisputerAddressUpdated", newDisputerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigDisputerAddressUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "DisputerAddressUpdated", log); err != nil {
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

// ParseDisputerAddressUpdated is a log parse operation binding the contract event 0xdb6c994f490543a0786e32e55707cbdba4ced511c54b42ef571b510e78c07798.
//
// Solidity: event DisputerAddressUpdated(address indexed newDisputer)
func (_AIConfig *AIConfigFilterer) ParseDisputerAddressUpdated(log types.Log) (*AIConfigDisputerAddressUpdated, error) {
	event := new(AIConfigDisputerAddressUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "DisputerAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigFeeDistributionUpdatedIterator is returned from FilterFeeDistributionUpdated and is used to iterate over the raw logs and unpacked data for FeeDistributionUpdated events raised by the AIConfig contract.
type AIConfigFeeDistributionUpdatedIterator struct {
	Event *AIConfigFeeDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigFeeDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigFeeDistributionUpdated)
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
		it.Event = new(AIConfigFeeDistributionUpdated)
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
func (it *AIConfigFeeDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigFeeDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigFeeDistributionUpdated represents a FeeDistributionUpdated event raised by the AIConfig contract.
type AIConfigFeeDistributionUpdated struct {
	WorkerBps   *big.Int
	ProtocolBps *big.Int
	BurnBps     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionUpdated is a free log retrieval operation binding the contract event 0xc2598ee1064d25a35cc42ae49954377c76f68eb4289b1825844a9f751bcae152.
//
// Solidity: event FeeDistributionUpdated(uint256 workerBps, uint256 protocolBps, uint256 burnBps)
func (_AIConfig *AIConfigFilterer) FilterFeeDistributionUpdated(opts *bind.FilterOpts) (*AIConfigFeeDistributionUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigFeeDistributionUpdatedIterator{contract: _AIConfig.contract, event: "FeeDistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionUpdated is a free log subscription operation binding the contract event 0xc2598ee1064d25a35cc42ae49954377c76f68eb4289b1825844a9f751bcae152.
//
// Solidity: event FeeDistributionUpdated(uint256 workerBps, uint256 protocolBps, uint256 burnBps)
func (_AIConfig *AIConfigFilterer) WatchFeeDistributionUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigFeeDistributionUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "FeeDistributionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigFeeDistributionUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
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

// ParseFeeDistributionUpdated is a log parse operation binding the contract event 0xc2598ee1064d25a35cc42ae49954377c76f68eb4289b1825844a9f751bcae152.
//
// Solidity: event FeeDistributionUpdated(uint256 workerBps, uint256 protocolBps, uint256 burnBps)
func (_AIConfig *AIConfigFilterer) ParseFeeDistributionUpdated(log types.Log) (*AIConfigFeeDistributionUpdated, error) {
	event := new(AIConfigFeeDistributionUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "FeeDistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigInfraParamsUpdatedIterator is returned from FilterInfraParamsUpdated and is used to iterate over the raw logs and unpacked data for InfraParamsUpdated events raised by the AIConfig contract.
type AIConfigInfraParamsUpdatedIterator struct {
	Event *AIConfigInfraParamsUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigInfraParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigInfraParamsUpdated)
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
		it.Event = new(AIConfigInfraParamsUpdated)
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
func (it *AIConfigInfraParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigInfraParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigInfraParamsUpdated represents a InfraParamsUpdated event raised by the AIConfig contract.
type AIConfigInfraParamsUpdated struct {
	Param string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInfraParamsUpdated is a free log retrieval operation binding the contract event 0xa0caaa6ae2842c710862de05ec5df8a156f1fc0a7fc03d0d3724f4f44b6aeffd.
//
// Solidity: event InfraParamsUpdated(string param, uint256 value)
func (_AIConfig *AIConfigFilterer) FilterInfraParamsUpdated(opts *bind.FilterOpts) (*AIConfigInfraParamsUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "InfraParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigInfraParamsUpdatedIterator{contract: _AIConfig.contract, event: "InfraParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchInfraParamsUpdated is a free log subscription operation binding the contract event 0xa0caaa6ae2842c710862de05ec5df8a156f1fc0a7fc03d0d3724f4f44b6aeffd.
//
// Solidity: event InfraParamsUpdated(string param, uint256 value)
func (_AIConfig *AIConfigFilterer) WatchInfraParamsUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigInfraParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "InfraParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigInfraParamsUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "InfraParamsUpdated", log); err != nil {
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

// ParseInfraParamsUpdated is a log parse operation binding the contract event 0xa0caaa6ae2842c710862de05ec5df8a156f1fc0a7fc03d0d3724f4f44b6aeffd.
//
// Solidity: event InfraParamsUpdated(string param, uint256 value)
func (_AIConfig *AIConfigFilterer) ParseInfraParamsUpdated(log types.Log) (*AIConfigInfraParamsUpdated, error) {
	event := new(AIConfigInfraParamsUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "InfraParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AIConfig contract.
type AIConfigInitializedIterator struct {
	Event *AIConfigInitialized // Event containing the contract specifics and raw log

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
func (it *AIConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigInitialized)
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
		it.Event = new(AIConfigInitialized)
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
func (it *AIConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigInitialized represents a Initialized event raised by the AIConfig contract.
type AIConfigInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AIConfig *AIConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*AIConfigInitializedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AIConfigInitializedIterator{contract: _AIConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AIConfig *AIConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AIConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigInitialized)
				if err := _AIConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AIConfig *AIConfigFilterer) ParseInitialized(log types.Log) (*AIConfigInitialized, error) {
	event := new(AIConfigInitialized)
	if err := _AIConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigMinWorkerStakeUpdatedIterator is returned from FilterMinWorkerStakeUpdated and is used to iterate over the raw logs and unpacked data for MinWorkerStakeUpdated events raised by the AIConfig contract.
type AIConfigMinWorkerStakeUpdatedIterator struct {
	Event *AIConfigMinWorkerStakeUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigMinWorkerStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigMinWorkerStakeUpdated)
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
		it.Event = new(AIConfigMinWorkerStakeUpdated)
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
func (it *AIConfigMinWorkerStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigMinWorkerStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigMinWorkerStakeUpdated represents a MinWorkerStakeUpdated event raised by the AIConfig contract.
type AIConfigMinWorkerStakeUpdated struct {
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinWorkerStakeUpdated is a free log retrieval operation binding the contract event 0x8994f374e952201580bec496d78026715d829676b3264f7028c532bce1150b01.
//
// Solidity: event MinWorkerStakeUpdated(uint256 newStake)
func (_AIConfig *AIConfigFilterer) FilterMinWorkerStakeUpdated(opts *bind.FilterOpts) (*AIConfigMinWorkerStakeUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "MinWorkerStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigMinWorkerStakeUpdatedIterator{contract: _AIConfig.contract, event: "MinWorkerStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinWorkerStakeUpdated is a free log subscription operation binding the contract event 0x8994f374e952201580bec496d78026715d829676b3264f7028c532bce1150b01.
//
// Solidity: event MinWorkerStakeUpdated(uint256 newStake)
func (_AIConfig *AIConfigFilterer) WatchMinWorkerStakeUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigMinWorkerStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "MinWorkerStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigMinWorkerStakeUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "MinWorkerStakeUpdated", log); err != nil {
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

// ParseMinWorkerStakeUpdated is a log parse operation binding the contract event 0x8994f374e952201580bec496d78026715d829676b3264f7028c532bce1150b01.
//
// Solidity: event MinWorkerStakeUpdated(uint256 newStake)
func (_AIConfig *AIConfigFilterer) ParseMinWorkerStakeUpdated(log types.Log) (*AIConfigMinWorkerStakeUpdated, error) {
	event := new(AIConfigMinWorkerStakeUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "MinWorkerStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigModelConfigUpdatedIterator is returned from FilterModelConfigUpdated and is used to iterate over the raw logs and unpacked data for ModelConfigUpdated events raised by the AIConfig contract.
type AIConfigModelConfigUpdatedIterator struct {
	Event *AIConfigModelConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigModelConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigModelConfigUpdated)
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
		it.Event = new(AIConfigModelConfigUpdated)
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
func (it *AIConfigModelConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigModelConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigModelConfigUpdated represents a ModelConfigUpdated event raised by the AIConfig contract.
type AIConfigModelConfigUpdated struct {
	ModelId         [32]byte
	Fee             *big.Int
	MaxOutputTokens *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModelConfigUpdated is a free log retrieval operation binding the contract event 0xb50c3aa7ade2043984058e30b86b019e7e60ba9a147fd7405b7845093bb5d9c9.
//
// Solidity: event ModelConfigUpdated(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) FilterModelConfigUpdated(opts *bind.FilterOpts, modelId [][32]byte) (*AIConfigModelConfigUpdatedIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "ModelConfigUpdated", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigModelConfigUpdatedIterator{contract: _AIConfig.contract, event: "ModelConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchModelConfigUpdated is a free log subscription operation binding the contract event 0xb50c3aa7ade2043984058e30b86b019e7e60ba9a147fd7405b7845093bb5d9c9.
//
// Solidity: event ModelConfigUpdated(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) WatchModelConfigUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigModelConfigUpdated, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "ModelConfigUpdated", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigModelConfigUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "ModelConfigUpdated", log); err != nil {
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

// ParseModelConfigUpdated is a log parse operation binding the contract event 0xb50c3aa7ade2043984058e30b86b019e7e60ba9a147fd7405b7845093bb5d9c9.
//
// Solidity: event ModelConfigUpdated(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) ParseModelConfigUpdated(log types.Log) (*AIConfigModelConfigUpdated, error) {
	event := new(AIConfigModelConfigUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "ModelConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigModelDisabledIterator is returned from FilterModelDisabled and is used to iterate over the raw logs and unpacked data for ModelDisabled events raised by the AIConfig contract.
type AIConfigModelDisabledIterator struct {
	Event *AIConfigModelDisabled // Event containing the contract specifics and raw log

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
func (it *AIConfigModelDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigModelDisabled)
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
		it.Event = new(AIConfigModelDisabled)
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
func (it *AIConfigModelDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigModelDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigModelDisabled represents a ModelDisabled event raised by the AIConfig contract.
type AIConfigModelDisabled struct {
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelDisabled is a free log retrieval operation binding the contract event 0x989461451824043ad73d7d959e3e2938fb67dda7cbd6ea10c3e40c3909951556.
//
// Solidity: event ModelDisabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) FilterModelDisabled(opts *bind.FilterOpts, modelId [][32]byte) (*AIConfigModelDisabledIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "ModelDisabled", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigModelDisabledIterator{contract: _AIConfig.contract, event: "ModelDisabled", logs: logs, sub: sub}, nil
}

// WatchModelDisabled is a free log subscription operation binding the contract event 0x989461451824043ad73d7d959e3e2938fb67dda7cbd6ea10c3e40c3909951556.
//
// Solidity: event ModelDisabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) WatchModelDisabled(opts *bind.WatchOpts, sink chan<- *AIConfigModelDisabled, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "ModelDisabled", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigModelDisabled)
				if err := _AIConfig.contract.UnpackLog(event, "ModelDisabled", log); err != nil {
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

// ParseModelDisabled is a log parse operation binding the contract event 0x989461451824043ad73d7d959e3e2938fb67dda7cbd6ea10c3e40c3909951556.
//
// Solidity: event ModelDisabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) ParseModelDisabled(log types.Log) (*AIConfigModelDisabled, error) {
	event := new(AIConfigModelDisabled)
	if err := _AIConfig.contract.UnpackLog(event, "ModelDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigModelEnabledIterator is returned from FilterModelEnabled and is used to iterate over the raw logs and unpacked data for ModelEnabled events raised by the AIConfig contract.
type AIConfigModelEnabledIterator struct {
	Event *AIConfigModelEnabled // Event containing the contract specifics and raw log

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
func (it *AIConfigModelEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigModelEnabled)
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
		it.Event = new(AIConfigModelEnabled)
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
func (it *AIConfigModelEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigModelEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigModelEnabled represents a ModelEnabled event raised by the AIConfig contract.
type AIConfigModelEnabled struct {
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelEnabled is a free log retrieval operation binding the contract event 0xf3396b477d45d29d661046d7eb2d75c0eb25a8d8727095a37fe8e5e5bbc1c6e9.
//
// Solidity: event ModelEnabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) FilterModelEnabled(opts *bind.FilterOpts, modelId [][32]byte) (*AIConfigModelEnabledIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "ModelEnabled", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigModelEnabledIterator{contract: _AIConfig.contract, event: "ModelEnabled", logs: logs, sub: sub}, nil
}

// WatchModelEnabled is a free log subscription operation binding the contract event 0xf3396b477d45d29d661046d7eb2d75c0eb25a8d8727095a37fe8e5e5bbc1c6e9.
//
// Solidity: event ModelEnabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) WatchModelEnabled(opts *bind.WatchOpts, sink chan<- *AIConfigModelEnabled, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "ModelEnabled", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigModelEnabled)
				if err := _AIConfig.contract.UnpackLog(event, "ModelEnabled", log); err != nil {
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

// ParseModelEnabled is a log parse operation binding the contract event 0xf3396b477d45d29d661046d7eb2d75c0eb25a8d8727095a37fe8e5e5bbc1c6e9.
//
// Solidity: event ModelEnabled(bytes32 indexed modelId)
func (_AIConfig *AIConfigFilterer) ParseModelEnabled(log types.Log) (*AIConfigModelEnabled, error) {
	event := new(AIConfigModelEnabled)
	if err := _AIConfig.contract.UnpackLog(event, "ModelEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigModelRegisteredIterator is returned from FilterModelRegistered and is used to iterate over the raw logs and unpacked data for ModelRegistered events raised by the AIConfig contract.
type AIConfigModelRegisteredIterator struct {
	Event *AIConfigModelRegistered // Event containing the contract specifics and raw log

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
func (it *AIConfigModelRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigModelRegistered)
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
		it.Event = new(AIConfigModelRegistered)
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
func (it *AIConfigModelRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigModelRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigModelRegistered represents a ModelRegistered event raised by the AIConfig contract.
type AIConfigModelRegistered struct {
	ModelId         [32]byte
	Fee             *big.Int
	MaxOutputTokens *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModelRegistered is a free log retrieval operation binding the contract event 0xae5a0d3a38bdeab8d8f9d4543bc9c4d389acc41ea1a7f9aa5d9effbfcc024462.
//
// Solidity: event ModelRegistered(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) FilterModelRegistered(opts *bind.FilterOpts, modelId [][32]byte) (*AIConfigModelRegisteredIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "ModelRegistered", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigModelRegisteredIterator{contract: _AIConfig.contract, event: "ModelRegistered", logs: logs, sub: sub}, nil
}

// WatchModelRegistered is a free log subscription operation binding the contract event 0xae5a0d3a38bdeab8d8f9d4543bc9c4d389acc41ea1a7f9aa5d9effbfcc024462.
//
// Solidity: event ModelRegistered(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) WatchModelRegistered(opts *bind.WatchOpts, sink chan<- *AIConfigModelRegistered, modelId [][32]byte) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "ModelRegistered", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigModelRegistered)
				if err := _AIConfig.contract.UnpackLog(event, "ModelRegistered", log); err != nil {
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

// ParseModelRegistered is a log parse operation binding the contract event 0xae5a0d3a38bdeab8d8f9d4543bc9c4d389acc41ea1a7f9aa5d9effbfcc024462.
//
// Solidity: event ModelRegistered(bytes32 indexed modelId, uint256 fee, uint256 maxOutputTokens)
func (_AIConfig *AIConfigFilterer) ParseModelRegistered(log types.Log) (*AIConfigModelRegistered, error) {
	event := new(AIConfigModelRegistered)
	if err := _AIConfig.contract.UnpackLog(event, "ModelRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AIConfig contract.
type AIConfigOwnershipTransferredIterator struct {
	Event *AIConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AIConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigOwnershipTransferred)
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
		it.Event = new(AIConfigOwnershipTransferred)
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
func (it *AIConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigOwnershipTransferred represents a OwnershipTransferred event raised by the AIConfig contract.
type AIConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AIConfig *AIConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AIConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigOwnershipTransferredIterator{contract: _AIConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AIConfig *AIConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AIConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigOwnershipTransferred)
				if err := _AIConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AIConfig *AIConfigFilterer) ParseOwnershipTransferred(log types.Log) (*AIConfigOwnershipTransferred, error) {
	event := new(AIConfigOwnershipTransferred)
	if err := _AIConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigSlashingParamsUpdatedIterator is returned from FilterSlashingParamsUpdated and is used to iterate over the raw logs and unpacked data for SlashingParamsUpdated events raised by the AIConfig contract.
type AIConfigSlashingParamsUpdatedIterator struct {
	Event *AIConfigSlashingParamsUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigSlashingParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigSlashingParamsUpdated)
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
		it.Event = new(AIConfigSlashingParamsUpdated)
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
func (it *AIConfigSlashingParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigSlashingParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigSlashingParamsUpdated represents a SlashingParamsUpdated event raised by the AIConfig contract.
type AIConfigSlashingParamsUpdated struct {
	TimeoutSlashBps           *big.Int
	CompletionTimeoutSlashBps *big.Int
	DisputeSlashBps           *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSlashingParamsUpdated is a free log retrieval operation binding the contract event 0x7517133c9bd7cc6c2096a9e359644fac973818bb14170f966190f95aab32e20c.
//
// Solidity: event SlashingParamsUpdated(uint256 timeoutSlashBps, uint256 completionTimeoutSlashBps, uint256 disputeSlashBps)
func (_AIConfig *AIConfigFilterer) FilterSlashingParamsUpdated(opts *bind.FilterOpts) (*AIConfigSlashingParamsUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "SlashingParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigSlashingParamsUpdatedIterator{contract: _AIConfig.contract, event: "SlashingParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashingParamsUpdated is a free log subscription operation binding the contract event 0x7517133c9bd7cc6c2096a9e359644fac973818bb14170f966190f95aab32e20c.
//
// Solidity: event SlashingParamsUpdated(uint256 timeoutSlashBps, uint256 completionTimeoutSlashBps, uint256 disputeSlashBps)
func (_AIConfig *AIConfigFilterer) WatchSlashingParamsUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigSlashingParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "SlashingParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigSlashingParamsUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "SlashingParamsUpdated", log); err != nil {
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

// ParseSlashingParamsUpdated is a log parse operation binding the contract event 0x7517133c9bd7cc6c2096a9e359644fac973818bb14170f966190f95aab32e20c.
//
// Solidity: event SlashingParamsUpdated(uint256 timeoutSlashBps, uint256 completionTimeoutSlashBps, uint256 disputeSlashBps)
func (_AIConfig *AIConfigFilterer) ParseSlashingParamsUpdated(log types.Log) (*AIConfigSlashingParamsUpdated, error) {
	event := new(AIConfigSlashingParamsUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "SlashingParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigSuspensionParamsUpdatedIterator is returned from FilterSuspensionParamsUpdated and is used to iterate over the raw logs and unpacked data for SuspensionParamsUpdated events raised by the AIConfig contract.
type AIConfigSuspensionParamsUpdatedIterator struct {
	Event *AIConfigSuspensionParamsUpdated // Event containing the contract specifics and raw log

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
func (it *AIConfigSuspensionParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigSuspensionParamsUpdated)
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
		it.Event = new(AIConfigSuspensionParamsUpdated)
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
func (it *AIConfigSuspensionParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigSuspensionParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigSuspensionParamsUpdated represents a SuspensionParamsUpdated event raised by the AIConfig contract.
type AIConfigSuspensionParamsUpdated struct {
	Threshold *big.Int
	Cooldown  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSuspensionParamsUpdated is a free log retrieval operation binding the contract event 0x8a1d0a4bdbc7c616db708e70b2dc58f8f9444a207ed26944e48a02fa37beaa1a.
//
// Solidity: event SuspensionParamsUpdated(uint256 threshold, uint256 cooldown)
func (_AIConfig *AIConfigFilterer) FilterSuspensionParamsUpdated(opts *bind.FilterOpts) (*AIConfigSuspensionParamsUpdatedIterator, error) {

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "SuspensionParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &AIConfigSuspensionParamsUpdatedIterator{contract: _AIConfig.contract, event: "SuspensionParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSuspensionParamsUpdated is a free log subscription operation binding the contract event 0x8a1d0a4bdbc7c616db708e70b2dc58f8f9444a207ed26944e48a02fa37beaa1a.
//
// Solidity: event SuspensionParamsUpdated(uint256 threshold, uint256 cooldown)
func (_AIConfig *AIConfigFilterer) WatchSuspensionParamsUpdated(opts *bind.WatchOpts, sink chan<- *AIConfigSuspensionParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "SuspensionParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigSuspensionParamsUpdated)
				if err := _AIConfig.contract.UnpackLog(event, "SuspensionParamsUpdated", log); err != nil {
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

// ParseSuspensionParamsUpdated is a log parse operation binding the contract event 0x8a1d0a4bdbc7c616db708e70b2dc58f8f9444a207ed26944e48a02fa37beaa1a.
//
// Solidity: event SuspensionParamsUpdated(uint256 threshold, uint256 cooldown)
func (_AIConfig *AIConfigFilterer) ParseSuspensionParamsUpdated(log types.Log) (*AIConfigSuspensionParamsUpdated, error) {
	event := new(AIConfigSuspensionParamsUpdated)
	if err := _AIConfig.contract.UnpackLog(event, "SuspensionParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIConfigUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AIConfig contract.
type AIConfigUpgradedIterator struct {
	Event *AIConfigUpgraded // Event containing the contract specifics and raw log

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
func (it *AIConfigUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIConfigUpgraded)
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
		it.Event = new(AIConfigUpgraded)
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
func (it *AIConfigUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIConfigUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIConfigUpgraded represents a Upgraded event raised by the AIConfig contract.
type AIConfigUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AIConfig *AIConfigFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AIConfigUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AIConfig.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AIConfigUpgradedIterator{contract: _AIConfig.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AIConfig *AIConfigFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AIConfigUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AIConfig.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIConfigUpgraded)
				if err := _AIConfig.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AIConfig *AIConfigFilterer) ParseUpgraded(log types.Log) (*AIConfigUpgraded, error) {
	event := new(AIConfigUpgraded)
	if err := _AIConfig.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
