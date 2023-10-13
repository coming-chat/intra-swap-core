// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package omni_swap

import (
	"errors"
	common2 "github.com/gkirito/go-ethereum/common"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/gkirito/go-ethereum/accounts/abi"
	"github.com/gkirito/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gkirito/go-ethereum/core/types"
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

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
//type IVaultBatchSwapStep struct {
//	PoolId        [32]byte
//	AssetInIndex  *big.Int
//	AssetOutIndex *big.Int
//	Amount        *big.Int
//	UserData      []byte
//}

// IVaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
//type IVaultFundManagement struct {
//	Sender              common.Address
//	FromInternalBalance bool
//	Recipient           common.Address
//	ToInternalBalance   bool
//}

// IVaultSingleSwap is an auto generated low-level Go binding around an user-defined struct.
//type IVaultSingleSwap struct {
//	PoolId   [32]byte
//	Kind     uint8
//	AssetIn  common.Address
//	AssetOut common.Address
//	Amount   *big.Int
//	UserData []byte
//}

// IBalancerQueriesMetaData contains all meta data concerning the IBalancerQueries contract.
var IBalancerQueriesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"queryBatchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"assetDeltas\",\"type\":\"int256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.SingleSwap\",\"name\":\"singleSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"querySwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBalancerQueriesABI is the input ABI used to generate the binding from.
// Deprecated: Use IBalancerQueriesMetaData.ABI instead.
var IBalancerQueriesABI = IBalancerQueriesMetaData.ABI

// IBalancerQueries is an auto generated Go binding around an Ethereum contract.
type IBalancerQueries struct {
	IBalancerQueriesCaller     // Read-only binding to the contract
	IBalancerQueriesTransactor // Write-only binding to the contract
	IBalancerQueriesFilterer   // Log filterer for contract events
}

// IBalancerQueriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBalancerQueriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBalancerQueriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBalancerQueriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBalancerQueriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBalancerQueriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBalancerQueriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBalancerQueriesSession struct {
	Contract     *IBalancerQueries // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBalancerQueriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBalancerQueriesCallerSession struct {
	Contract *IBalancerQueriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IBalancerQueriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBalancerQueriesTransactorSession struct {
	Contract     *IBalancerQueriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IBalancerQueriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBalancerQueriesRaw struct {
	Contract *IBalancerQueries // Generic contract binding to access the raw methods on
}

// IBalancerQueriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBalancerQueriesCallerRaw struct {
	Contract *IBalancerQueriesCaller // Generic read-only contract binding to access the raw methods on
}

// IBalancerQueriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBalancerQueriesTransactorRaw struct {
	Contract *IBalancerQueriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBalancerQueries creates a new instance of IBalancerQueries, bound to a specific deployed contract.
func NewIBalancerQueries(address common.Address, backend bind.ContractBackend) (*IBalancerQueries, error) {
	contract, err := bindIBalancerQueries(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBalancerQueries{IBalancerQueriesCaller: IBalancerQueriesCaller{contract: contract}, IBalancerQueriesTransactor: IBalancerQueriesTransactor{contract: contract}, IBalancerQueriesFilterer: IBalancerQueriesFilterer{contract: contract}}, nil
}

// NewIBalancerQueriesCaller creates a new read-only instance of IBalancerQueries, bound to a specific deployed contract.
func NewIBalancerQueriesCaller(address common.Address, caller bind.ContractCaller) (*IBalancerQueriesCaller, error) {
	contract, err := bindIBalancerQueries(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBalancerQueriesCaller{contract: contract}, nil
}

// NewIBalancerQueriesTransactor creates a new write-only instance of IBalancerQueries, bound to a specific deployed contract.
func NewIBalancerQueriesTransactor(address common.Address, transactor bind.ContractTransactor) (*IBalancerQueriesTransactor, error) {
	contract, err := bindIBalancerQueries(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBalancerQueriesTransactor{contract: contract}, nil
}

// NewIBalancerQueriesFilterer creates a new log filterer instance of IBalancerQueries, bound to a specific deployed contract.
func NewIBalancerQueriesFilterer(address common.Address, filterer bind.ContractFilterer) (*IBalancerQueriesFilterer, error) {
	contract, err := bindIBalancerQueries(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBalancerQueriesFilterer{contract: contract}, nil
}

// bindIBalancerQueries binds a generic wrapper to an already deployed contract.
func bindIBalancerQueries(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBalancerQueriesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBalancerQueries *IBalancerQueriesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBalancerQueries.Contract.IBalancerQueriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBalancerQueries *IBalancerQueriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.IBalancerQueriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBalancerQueries *IBalancerQueriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.IBalancerQueriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBalancerQueries *IBalancerQueriesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBalancerQueries.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBalancerQueries *IBalancerQueriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBalancerQueries *IBalancerQueriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.contract.Transact(opts, method, params...)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[] assetDeltas)
func (_IBalancerQueries *IBalancerQueriesTransactor) QueryBatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.contract.Transact(opts, "queryBatchSwap", kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[] assetDeltas)
func (_IBalancerQueries *IBalancerQueriesSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.QueryBatchSwap(&_IBalancerQueries.TransactOpts, kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[] assetDeltas)
func (_IBalancerQueries *IBalancerQueriesTransactorSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.QueryBatchSwap(&_IBalancerQueries.TransactOpts, kind, swaps, assets, funds)
}

// QuerySwap is a paid mutator transaction binding the contract method 0xe969f6b3.
//
// Solidity: function querySwap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds) returns(uint256)
func (_IBalancerQueries *IBalancerQueriesTransactor) QuerySwap(opts *bind.TransactOpts, singleSwap IVaultSingleSwap, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.contract.Transact(opts, "querySwap", singleSwap, funds)
}

// QuerySwap is a paid mutator transaction binding the contract method 0xe969f6b3.
//
// Solidity: function querySwap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds) returns(uint256)
func (_IBalancerQueries *IBalancerQueriesSession) QuerySwap(singleSwap IVaultSingleSwap, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.QuerySwap(&_IBalancerQueries.TransactOpts, singleSwap, funds)
}

// QuerySwap is a paid mutator transaction binding the contract method 0xe969f6b3.
//
// Solidity: function querySwap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds) returns(uint256)
func (_IBalancerQueries *IBalancerQueriesTransactorSession) QuerySwap(singleSwap IVaultSingleSwap, funds IVaultFundManagement) (*types.Transaction, error) {
	return _IBalancerQueries.Contract.QuerySwap(&_IBalancerQueries.TransactOpts, singleSwap, funds)
}
