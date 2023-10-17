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

// ISwapBasedQuoterMetaData contains all meta data concerning the ISwapBasedQuoter contract.
var ISwapBasedQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"fees\",\"type\":\"uint16[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"fees\",\"type\":\"uint16[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapBasedQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapBasedQuoterMetaData.ABI instead.
var ISwapBasedQuoterABI = ISwapBasedQuoterMetaData.ABI

// ISwapBasedQuoter is an auto generated Go binding around an Ethereum contract.
type ISwapBasedQuoter struct {
	ISwapBasedQuoterCaller     // Read-only binding to the contract
	ISwapBasedQuoterTransactor // Write-only binding to the contract
	ISwapBasedQuoterFilterer   // Log filterer for contract events
}

// ISwapBasedQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapBasedQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapBasedQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapBasedQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapBasedQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapBasedQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapBasedQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapBasedQuoterSession struct {
	Contract     *ISwapBasedQuoter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapBasedQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapBasedQuoterCallerSession struct {
	Contract *ISwapBasedQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISwapBasedQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapBasedQuoterTransactorSession struct {
	Contract     *ISwapBasedQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISwapBasedQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapBasedQuoterRaw struct {
	Contract *ISwapBasedQuoter // Generic contract binding to access the raw methods on
}

// ISwapBasedQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapBasedQuoterCallerRaw struct {
	Contract *ISwapBasedQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapBasedQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapBasedQuoterTransactorRaw struct {
	Contract *ISwapBasedQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapBasedQuoter creates a new instance of ISwapBasedQuoter, bound to a specific deployed contract.
func NewISwapBasedQuoter(address common.Address, backend bind.ContractBackend) (*ISwapBasedQuoter, error) {
	contract, err := bindISwapBasedQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapBasedQuoter{ISwapBasedQuoterCaller: ISwapBasedQuoterCaller{contract: contract}, ISwapBasedQuoterTransactor: ISwapBasedQuoterTransactor{contract: contract}, ISwapBasedQuoterFilterer: ISwapBasedQuoterFilterer{contract: contract}}, nil
}

// NewISwapBasedQuoterCaller creates a new read-only instance of ISwapBasedQuoter, bound to a specific deployed contract.
func NewISwapBasedQuoterCaller(address common.Address, caller bind.ContractCaller) (*ISwapBasedQuoterCaller, error) {
	contract, err := bindISwapBasedQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapBasedQuoterCaller{contract: contract}, nil
}

// NewISwapBasedQuoterTransactor creates a new write-only instance of ISwapBasedQuoter, bound to a specific deployed contract.
func NewISwapBasedQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapBasedQuoterTransactor, error) {
	contract, err := bindISwapBasedQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapBasedQuoterTransactor{contract: contract}, nil
}

// NewISwapBasedQuoterFilterer creates a new log filterer instance of ISwapBasedQuoter, bound to a specific deployed contract.
func NewISwapBasedQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapBasedQuoterFilterer, error) {
	contract, err := bindISwapBasedQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapBasedQuoterFilterer{contract: contract}, nil
}

// bindISwapBasedQuoter binds a generic wrapper to an already deployed contract.
func bindISwapBasedQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapBasedQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapBasedQuoter *ISwapBasedQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapBasedQuoter.Contract.ISwapBasedQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapBasedQuoter *ISwapBasedQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.ISwapBasedQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapBasedQuoter *ISwapBasedQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.ISwapBasedQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapBasedQuoter *ISwapBasedQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapBasedQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactInput(&_ISwapBasedQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactInput(&_ISwapBasedQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.contract.Transact(opts, "quoteExactInputSingle", tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactInputSingle(&_ISwapBasedQuoter.TransactOpts, tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactInputSingle(&_ISwapBasedQuoter.TransactOpts, tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactOutput(&_ISwapBasedQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactOutput(&_ISwapBasedQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.contract.Transact(opts, "quoteExactOutputSingle", tokenIn, tokenOut, amountOut, limitSqrtPrice)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactOutputSingle(&_ISwapBasedQuoter.TransactOpts, tokenIn, tokenOut, amountOut, limitSqrtPrice)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_ISwapBasedQuoter *ISwapBasedQuoterTransactorSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _ISwapBasedQuoter.Contract.QuoteExactOutputSingle(&_ISwapBasedQuoter.TransactOpts, tokenIn, tokenOut, amountOut, limitSqrtPrice)
}
