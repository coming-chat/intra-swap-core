// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package omni_swap

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

// IQuickSwapCallbackMetaData contains all meta data concerning the IQuickSwapCallback contract.
var IQuickSwapCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IQuickSwapCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IQuickSwapCallbackMetaData.ABI instead.
var IQuickSwapCallbackABI = IQuickSwapCallbackMetaData.ABI

// IQuickSwapCallback is an auto generated Go binding around an Ethereum contract.
type IQuickSwapCallback struct {
	IQuickSwapCallbackCaller     // Read-only binding to the contract
	IQuickSwapCallbackTransactor // Write-only binding to the contract
	IQuickSwapCallbackFilterer   // Log filterer for contract events
}

// IQuickSwapCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuickSwapCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuickSwapCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuickSwapCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuickSwapCallbackSession struct {
	Contract     *IQuickSwapCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IQuickSwapCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuickSwapCallbackCallerSession struct {
	Contract *IQuickSwapCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IQuickSwapCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuickSwapCallbackTransactorSession struct {
	Contract     *IQuickSwapCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IQuickSwapCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuickSwapCallbackRaw struct {
	Contract *IQuickSwapCallback // Generic contract binding to access the raw methods on
}

// IQuickSwapCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuickSwapCallbackCallerRaw struct {
	Contract *IQuickSwapCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IQuickSwapCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuickSwapCallbackTransactorRaw struct {
	Contract *IQuickSwapCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuickSwapCallback creates a new instance of IQuickSwapCallback, bound to a specific deployed contract.
func NewIQuickSwapCallback(address common.Address, backend bind.ContractBackend) (*IQuickSwapCallback, error) {
	contract, err := bindIQuickSwapCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapCallback{IQuickSwapCallbackCaller: IQuickSwapCallbackCaller{contract: contract}, IQuickSwapCallbackTransactor: IQuickSwapCallbackTransactor{contract: contract}, IQuickSwapCallbackFilterer: IQuickSwapCallbackFilterer{contract: contract}}, nil
}

// NewIQuickSwapCallbackCaller creates a new read-only instance of IQuickSwapCallback, bound to a specific deployed contract.
func NewIQuickSwapCallbackCaller(address common.Address, caller bind.ContractCaller) (*IQuickSwapCallbackCaller, error) {
	contract, err := bindIQuickSwapCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapCallbackCaller{contract: contract}, nil
}

// NewIQuickSwapCallbackTransactor creates a new write-only instance of IQuickSwapCallback, bound to a specific deployed contract.
func NewIQuickSwapCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuickSwapCallbackTransactor, error) {
	contract, err := bindIQuickSwapCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapCallbackTransactor{contract: contract}, nil
}

// NewIQuickSwapCallbackFilterer creates a new log filterer instance of IQuickSwapCallback, bound to a specific deployed contract.
func NewIQuickSwapCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuickSwapCallbackFilterer, error) {
	contract, err := bindIQuickSwapCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapCallbackFilterer{contract: contract}, nil
}

// bindIQuickSwapCallback binds a generic wrapper to an already deployed contract.
func bindIQuickSwapCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IQuickSwapCallbackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickSwapCallback *IQuickSwapCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickSwapCallback.Contract.IQuickSwapCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickSwapCallback *IQuickSwapCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.IQuickSwapCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickSwapCallback *IQuickSwapCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.IQuickSwapCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickSwapCallback *IQuickSwapCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickSwapCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickSwapCallback *IQuickSwapCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickSwapCallback *IQuickSwapCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.contract.Transact(opts, method, params...)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapCallback *IQuickSwapCallbackTransactor) AlgebraSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapCallback.contract.Transact(opts, "algebraSwapCallback", amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapCallback *IQuickSwapCallbackSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.AlgebraSwapCallback(&_IQuickSwapCallback.TransactOpts, amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapCallback *IQuickSwapCallbackTransactorSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapCallback.Contract.AlgebraSwapCallback(&_IQuickSwapCallback.TransactOpts, amount0Delta, amount1Delta, data)
}
