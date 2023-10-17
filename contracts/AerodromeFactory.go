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

// AerodromeFactoryMetaData contains all meta data concerning the AerodromeFactory contract.
var AerodromeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AerodromeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AerodromeFactoryMetaData.ABI instead.
var AerodromeFactoryABI = AerodromeFactoryMetaData.ABI

// AerodromeFactory is an auto generated Go binding around an Ethereum contract.
type AerodromeFactory struct {
	AerodromeFactoryCaller     // Read-only binding to the contract
	AerodromeFactoryTransactor // Write-only binding to the contract
	AerodromeFactoryFilterer   // Log filterer for contract events
}

// AerodromeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AerodromeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AerodromeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AerodromeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AerodromeFactorySession struct {
	Contract     *AerodromeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AerodromeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AerodromeFactoryCallerSession struct {
	Contract *AerodromeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AerodromeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AerodromeFactoryTransactorSession struct {
	Contract     *AerodromeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AerodromeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AerodromeFactoryRaw struct {
	Contract *AerodromeFactory // Generic contract binding to access the raw methods on
}

// AerodromeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AerodromeFactoryCallerRaw struct {
	Contract *AerodromeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AerodromeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AerodromeFactoryTransactorRaw struct {
	Contract *AerodromeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAerodromeFactory creates a new instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactory(address common.Address, backend bind.ContractBackend) (*AerodromeFactory, error) {
	contract, err := bindAerodromeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactory{AerodromeFactoryCaller: AerodromeFactoryCaller{contract: contract}, AerodromeFactoryTransactor: AerodromeFactoryTransactor{contract: contract}, AerodromeFactoryFilterer: AerodromeFactoryFilterer{contract: contract}}, nil
}

// NewAerodromeFactoryCaller creates a new read-only instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryCaller(address common.Address, caller bind.ContractCaller) (*AerodromeFactoryCaller, error) {
	contract, err := bindAerodromeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryCaller{contract: contract}, nil
}

// NewAerodromeFactoryTransactor creates a new write-only instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AerodromeFactoryTransactor, error) {
	contract, err := bindAerodromeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryTransactor{contract: contract}, nil
}

// NewAerodromeFactoryFilterer creates a new log filterer instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AerodromeFactoryFilterer, error) {
	contract, err := bindAerodromeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryFilterer{contract: contract}, nil
}

// bindAerodromeFactory binds a generic wrapper to an already deployed contract.
func bindAerodromeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AerodromeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeFactory *AerodromeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeFactory.Contract.AerodromeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeFactory *AerodromeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.AerodromeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeFactory *AerodromeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.AerodromeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeFactory *AerodromeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeFactory *AerodromeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeFactory *AerodromeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_AerodromeFactory *AerodromeFactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "getPool", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_AerodromeFactory *AerodromeFactorySession) GetPool(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _AerodromeFactory.Contract.GetPool(&_AerodromeFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPool is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_AerodromeFactory *AerodromeFactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _AerodromeFactory.Contract.GetPool(&_AerodromeFactory.CallOpts, tokenA, tokenB, stable)
}
