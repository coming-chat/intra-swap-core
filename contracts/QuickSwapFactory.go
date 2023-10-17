// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/gkirito/go-ethereum"
	"github.com/gkirito/go-ethereum/accounts/abi"
	"github.com/gkirito/go-ethereum/accounts/abi/bind"
	"github.com/gkirito/go-ethereum/common"
	"github.com/gkirito/go-ethereum/core/types"
	"github.com/gkirito/go-ethereum/event"
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

// QuickSwapFactoryMetaData contains all meta data concerning the QuickSwapFactory contract.
var QuickSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QuickSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use QuickSwapFactoryMetaData.ABI instead.
var QuickSwapFactoryABI = QuickSwapFactoryMetaData.ABI

// QuickSwapFactory is an auto generated Go binding around an Ethereum contract.
type QuickSwapFactory struct {
	QuickSwapFactoryCaller     // Read-only binding to the contract
	QuickSwapFactoryTransactor // Write-only binding to the contract
	QuickSwapFactoryFilterer   // Log filterer for contract events
}

// QuickSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuickSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuickSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuickSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuickSwapFactorySession struct {
	Contract     *QuickSwapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuickSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuickSwapFactoryCallerSession struct {
	Contract *QuickSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// QuickSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuickSwapFactoryTransactorSession struct {
	Contract     *QuickSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QuickSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuickSwapFactoryRaw struct {
	Contract *QuickSwapFactory // Generic contract binding to access the raw methods on
}

// QuickSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuickSwapFactoryCallerRaw struct {
	Contract *QuickSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// QuickSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuickSwapFactoryTransactorRaw struct {
	Contract *QuickSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuickSwapFactory creates a new instance of QuickSwapFactory, bound to a specific deployed contract.
func NewQuickSwapFactory(address common.Address, backend bind.ContractBackend) (*QuickSwapFactory, error) {
	contract, err := bindQuickSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuickSwapFactory{QuickSwapFactoryCaller: QuickSwapFactoryCaller{contract: contract}, QuickSwapFactoryTransactor: QuickSwapFactoryTransactor{contract: contract}, QuickSwapFactoryFilterer: QuickSwapFactoryFilterer{contract: contract}}, nil
}

// NewQuickSwapFactoryCaller creates a new read-only instance of QuickSwapFactory, bound to a specific deployed contract.
func NewQuickSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*QuickSwapFactoryCaller, error) {
	contract, err := bindQuickSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuickSwapFactoryCaller{contract: contract}, nil
}

// NewQuickSwapFactoryTransactor creates a new write-only instance of QuickSwapFactory, bound to a specific deployed contract.
func NewQuickSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*QuickSwapFactoryTransactor, error) {
	contract, err := bindQuickSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuickSwapFactoryTransactor{contract: contract}, nil
}

// NewQuickSwapFactoryFilterer creates a new log filterer instance of QuickSwapFactory, bound to a specific deployed contract.
func NewQuickSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*QuickSwapFactoryFilterer, error) {
	contract, err := bindQuickSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuickSwapFactoryFilterer{contract: contract}, nil
}

// bindQuickSwapFactory binds a generic wrapper to an already deployed contract.
func bindQuickSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuickSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickSwapFactory *QuickSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickSwapFactory.Contract.QuickSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickSwapFactory *QuickSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickSwapFactory.Contract.QuickSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickSwapFactory *QuickSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickSwapFactory.Contract.QuickSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickSwapFactory *QuickSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickSwapFactory *QuickSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickSwapFactory *QuickSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_QuickSwapFactory *QuickSwapFactoryCaller) PoolByPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _QuickSwapFactory.contract.Call(opts, &out, "poolByPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_QuickSwapFactory *QuickSwapFactorySession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _QuickSwapFactory.Contract.PoolByPair(&_QuickSwapFactory.CallOpts, tokenA, tokenB)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_QuickSwapFactory *QuickSwapFactoryCallerSession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _QuickSwapFactory.Contract.PoolByPair(&_QuickSwapFactory.CallOpts, tokenA, tokenB)
}
