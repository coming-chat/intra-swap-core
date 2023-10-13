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

// IQuickQuoterMetaData contains all meta data concerning the IQuickQuoter contract.
var IQuickQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"fees\",\"type\":\"uint16[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint16[]\",\"name\":\"fees\",\"type\":\"uint16[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IQuickQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use IQuickQuoterMetaData.ABI instead.
var IQuickQuoterABI = IQuickQuoterMetaData.ABI

// IQuickQuoter is an auto generated Go binding around an Ethereum contract.
type IQuickQuoter struct {
	IQuickQuoterCaller     // Read-only binding to the contract
	IQuickQuoterTransactor // Write-only binding to the contract
	IQuickQuoterFilterer   // Log filterer for contract events
}

// IQuickQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuickQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuickQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuickQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuickQuoterSession struct {
	Contract     *IQuickQuoter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQuickQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuickQuoterCallerSession struct {
	Contract *IQuickQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IQuickQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuickQuoterTransactorSession struct {
	Contract     *IQuickQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IQuickQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuickQuoterRaw struct {
	Contract *IQuickQuoter // Generic contract binding to access the raw methods on
}

// IQuickQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuickQuoterCallerRaw struct {
	Contract *IQuickQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// IQuickQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuickQuoterTransactorRaw struct {
	Contract *IQuickQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuickQuoter creates a new instance of IQuickQuoter, bound to a specific deployed contract.
func NewIQuickQuoter(address common.Address, backend bind.ContractBackend) (*IQuickQuoter, error) {
	contract, err := bindIQuickQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuickQuoter{IQuickQuoterCaller: IQuickQuoterCaller{contract: contract}, IQuickQuoterTransactor: IQuickQuoterTransactor{contract: contract}, IQuickQuoterFilterer: IQuickQuoterFilterer{contract: contract}}, nil
}

// NewIQuickQuoterCaller creates a new read-only instance of IQuickQuoter, bound to a specific deployed contract.
func NewIQuickQuoterCaller(address common.Address, caller bind.ContractCaller) (*IQuickQuoterCaller, error) {
	contract, err := bindIQuickQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickQuoterCaller{contract: contract}, nil
}

// NewIQuickQuoterTransactor creates a new write-only instance of IQuickQuoter, bound to a specific deployed contract.
func NewIQuickQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuickQuoterTransactor, error) {
	contract, err := bindIQuickQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickQuoterTransactor{contract: contract}, nil
}

// NewIQuickQuoterFilterer creates a new log filterer instance of IQuickQuoter, bound to a specific deployed contract.
func NewIQuickQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuickQuoterFilterer, error) {
	contract, err := bindIQuickQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuickQuoterFilterer{contract: contract}, nil
}

// bindIQuickQuoter binds a generic wrapper to an already deployed contract.
func bindIQuickQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IQuickQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickQuoter *IQuickQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickQuoter.Contract.IQuickQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickQuoter *IQuickQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.IQuickQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickQuoter *IQuickQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.IQuickQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickQuoter *IQuickQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickQuoter *IQuickQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickQuoter *IQuickQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactInput(&_IQuickQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactInput(&_IQuickQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_IQuickQuoter *IQuickQuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.contract.Transact(opts, "quoteExactInputSingle", tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_IQuickQuoter *IQuickQuoterSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactInputSingle(&_IQuickQuoter.TransactOpts, tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x2d9ebd1d.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint256 amountIn, uint160 limitSqrtPrice) returns(uint256 amountOut, uint16 fee)
func (_IQuickQuoter *IQuickQuoterTransactorSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactInputSingle(&_IQuickQuoter.TransactOpts, tokenIn, tokenOut, amountIn, limitSqrtPrice)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactOutput(&_IQuickQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint16[] fees)
func (_IQuickQuoter *IQuickQuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactOutput(&_IQuickQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_IQuickQuoter *IQuickQuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.contract.Transact(opts, "quoteExactOutputSingle", tokenIn, tokenOut, amountOut, limitSqrtPrice)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_IQuickQuoter *IQuickQuoterSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactOutputSingle(&_IQuickQuoter.TransactOpts, tokenIn, tokenOut, amountOut, limitSqrtPrice)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x9e73c81d.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint256 amountOut, uint160 limitSqrtPrice) returns(uint256 amountIn, uint16 fee)
func (_IQuickQuoter *IQuickQuoterTransactorSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, amountOut *big.Int, limitSqrtPrice *big.Int) (*types.Transaction, error) {
	return _IQuickQuoter.Contract.QuoteExactOutputSingle(&_IQuickQuoter.TransactOpts, tokenIn, tokenOut, amountOut, limitSqrtPrice)
}
