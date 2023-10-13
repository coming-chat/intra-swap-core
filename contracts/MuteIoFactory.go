// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// MuteIoFactoryMetaData contains all meta data concerning the MuteIoFactory contract.
var MuteIoFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MuteIoFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MuteIoFactoryMetaData.ABI instead.
var MuteIoFactoryABI = MuteIoFactoryMetaData.ABI

// MuteIoFactory is an auto generated Go binding around an Ethereum contract.
type MuteIoFactory struct {
	MuteIoFactoryCaller     // Read-only binding to the contract
	MuteIoFactoryTransactor // Write-only binding to the contract
	MuteIoFactoryFilterer   // Log filterer for contract events
}

// MuteIoFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuteIoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuteIoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuteIoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuteIoFactorySession struct {
	Contract     *MuteIoFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuteIoFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuteIoFactoryCallerSession struct {
	Contract *MuteIoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MuteIoFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuteIoFactoryTransactorSession struct {
	Contract     *MuteIoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MuteIoFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuteIoFactoryRaw struct {
	Contract *MuteIoFactory // Generic contract binding to access the raw methods on
}

// MuteIoFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuteIoFactoryCallerRaw struct {
	Contract *MuteIoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MuteIoFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuteIoFactoryTransactorRaw struct {
	Contract *MuteIoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuteIoFactory creates a new instance of MuteIoFactory, bound to a specific deployed contract.
func NewMuteIoFactory(address common.Address, backend bind.ContractBackend) (*MuteIoFactory, error) {
	contract, err := bindMuteIoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuteIoFactory{MuteIoFactoryCaller: MuteIoFactoryCaller{contract: contract}, MuteIoFactoryTransactor: MuteIoFactoryTransactor{contract: contract}, MuteIoFactoryFilterer: MuteIoFactoryFilterer{contract: contract}}, nil
}

// NewMuteIoFactoryCaller creates a new read-only instance of MuteIoFactory, bound to a specific deployed contract.
func NewMuteIoFactoryCaller(address common.Address, caller bind.ContractCaller) (*MuteIoFactoryCaller, error) {
	contract, err := bindMuteIoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuteIoFactoryCaller{contract: contract}, nil
}

// NewMuteIoFactoryTransactor creates a new write-only instance of MuteIoFactory, bound to a specific deployed contract.
func NewMuteIoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MuteIoFactoryTransactor, error) {
	contract, err := bindMuteIoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuteIoFactoryTransactor{contract: contract}, nil
}

// NewMuteIoFactoryFilterer creates a new log filterer instance of MuteIoFactory, bound to a specific deployed contract.
func NewMuteIoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MuteIoFactoryFilterer, error) {
	contract, err := bindMuteIoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuteIoFactoryFilterer{contract: contract}, nil
}

// bindMuteIoFactory binds a generic wrapper to an already deployed contract.
func bindMuteIoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuteIoFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuteIoFactory *MuteIoFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuteIoFactory.Contract.MuteIoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuteIoFactory *MuteIoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuteIoFactory.Contract.MuteIoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuteIoFactory *MuteIoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuteIoFactory.Contract.MuteIoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuteIoFactory *MuteIoFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuteIoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuteIoFactory *MuteIoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuteIoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuteIoFactory *MuteIoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuteIoFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_MuteIoFactory *MuteIoFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _MuteIoFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_MuteIoFactory *MuteIoFactorySession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _MuteIoFactory.Contract.GetPair(&_MuteIoFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address pool)
func (_MuteIoFactory *MuteIoFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _MuteIoFactory.Contract.GetPair(&_MuteIoFactory.CallOpts, tokenA, tokenB, stable)
}
