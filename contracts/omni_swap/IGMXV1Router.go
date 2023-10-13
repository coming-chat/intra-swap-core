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

// IGMXV1RouterMetaData contains all meta data concerning the IGMXV1Router contract.
var IGMXV1RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swapETHToTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swapTokensToETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGMXV1RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IGMXV1RouterMetaData.ABI instead.
var IGMXV1RouterABI = IGMXV1RouterMetaData.ABI

// IGMXV1Router is an auto generated Go binding around an Ethereum contract.
type IGMXV1Router struct {
	IGMXV1RouterCaller     // Read-only binding to the contract
	IGMXV1RouterTransactor // Write-only binding to the contract
	IGMXV1RouterFilterer   // Log filterer for contract events
}

// IGMXV1RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGMXV1RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGMXV1RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGMXV1RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGMXV1RouterSession struct {
	Contract     *IGMXV1Router     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGMXV1RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGMXV1RouterCallerSession struct {
	Contract *IGMXV1RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGMXV1RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGMXV1RouterTransactorSession struct {
	Contract     *IGMXV1RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGMXV1RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGMXV1RouterRaw struct {
	Contract *IGMXV1Router // Generic contract binding to access the raw methods on
}

// IGMXV1RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGMXV1RouterCallerRaw struct {
	Contract *IGMXV1RouterCaller // Generic read-only contract binding to access the raw methods on
}

// IGMXV1RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGMXV1RouterTransactorRaw struct {
	Contract *IGMXV1RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGMXV1Router creates a new instance of IGMXV1Router, bound to a specific deployed contract.
func NewIGMXV1Router(address common.Address, backend bind.ContractBackend) (*IGMXV1Router, error) {
	contract, err := bindIGMXV1Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGMXV1Router{IGMXV1RouterCaller: IGMXV1RouterCaller{contract: contract}, IGMXV1RouterTransactor: IGMXV1RouterTransactor{contract: contract}, IGMXV1RouterFilterer: IGMXV1RouterFilterer{contract: contract}}, nil
}

// NewIGMXV1RouterCaller creates a new read-only instance of IGMXV1Router, bound to a specific deployed contract.
func NewIGMXV1RouterCaller(address common.Address, caller bind.ContractCaller) (*IGMXV1RouterCaller, error) {
	contract, err := bindIGMXV1Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1RouterCaller{contract: contract}, nil
}

// NewIGMXV1RouterTransactor creates a new write-only instance of IGMXV1Router, bound to a specific deployed contract.
func NewIGMXV1RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGMXV1RouterTransactor, error) {
	contract, err := bindIGMXV1Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1RouterTransactor{contract: contract}, nil
}

// NewIGMXV1RouterFilterer creates a new log filterer instance of IGMXV1Router, bound to a specific deployed contract.
func NewIGMXV1RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGMXV1RouterFilterer, error) {
	contract, err := bindIGMXV1Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGMXV1RouterFilterer{contract: contract}, nil
}

// bindIGMXV1Router binds a generic wrapper to an already deployed contract.
func bindIGMXV1Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGMXV1RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Router *IGMXV1RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Router.Contract.IGMXV1RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Router *IGMXV1RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.IGMXV1RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Router *IGMXV1RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.IGMXV1RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Router *IGMXV1RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Router *IGMXV1RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Router *IGMXV1RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterTransactor) Swap(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.contract.Transact(opts, "swap", _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.Swap(&_IGMXV1Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterTransactorSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.Swap(&_IGMXV1Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_IGMXV1Router *IGMXV1RouterTransactor) SwapETHToTokens(opts *bind.TransactOpts, _path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.contract.Transact(opts, "swapETHToTokens", _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_IGMXV1Router *IGMXV1RouterSession) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.SwapETHToTokens(&_IGMXV1Router.TransactOpts, _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_IGMXV1Router *IGMXV1RouterTransactorSession) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.SwapETHToTokens(&_IGMXV1Router.TransactOpts, _path, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterTransactor) SwapTokensToETH(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.contract.Transact(opts, "swapTokensToETH", _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterSession) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.SwapTokensToETH(&_IGMXV1Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IGMXV1Router *IGMXV1RouterTransactorSession) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Router.Contract.SwapTokensToETH(&_IGMXV1Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}
