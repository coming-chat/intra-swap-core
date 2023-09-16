// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package omni_swap

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

// IQuickSwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IQuickSwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterExactInputSingleParams struct {
	TokenIn          common.Address
	TokenOut         common.Address
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	LimitSqrtPrice   *big.Int
}

// IQuickSwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IQuickSwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterExactOutputSingleParams struct {
	TokenIn         common.Address
	TokenOut        common.Address
	Fee             *big.Int
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	LimitSqrtPrice  *big.Int
}

// IQuickSwapRouterMetaData contains all meta data concerning the IQuickSwapRouter contract.
var IQuickSwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIQuickSwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingleSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIQuickSwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IQuickSwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IQuickSwapRouterMetaData.ABI instead.
var IQuickSwapRouterABI = IQuickSwapRouterMetaData.ABI

// IQuickSwapRouter is an auto generated Go binding around an Ethereum contract.
type IQuickSwapRouter struct {
	IQuickSwapRouterCaller     // Read-only binding to the contract
	IQuickSwapRouterTransactor // Write-only binding to the contract
	IQuickSwapRouterFilterer   // Log filterer for contract events
}

// IQuickSwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuickSwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuickSwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuickSwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickSwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuickSwapRouterSession struct {
	Contract     *IQuickSwapRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQuickSwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuickSwapRouterCallerSession struct {
	Contract *IQuickSwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IQuickSwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuickSwapRouterTransactorSession struct {
	Contract     *IQuickSwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IQuickSwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuickSwapRouterRaw struct {
	Contract *IQuickSwapRouter // Generic contract binding to access the raw methods on
}

// IQuickSwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuickSwapRouterCallerRaw struct {
	Contract *IQuickSwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IQuickSwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuickSwapRouterTransactorRaw struct {
	Contract *IQuickSwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuickSwapRouter creates a new instance of IQuickSwapRouter, bound to a specific deployed contract.
func NewIQuickSwapRouter(address common.Address, backend bind.ContractBackend) (*IQuickSwapRouter, error) {
	contract, err := bindIQuickSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapRouter{IQuickSwapRouterCaller: IQuickSwapRouterCaller{contract: contract}, IQuickSwapRouterTransactor: IQuickSwapRouterTransactor{contract: contract}, IQuickSwapRouterFilterer: IQuickSwapRouterFilterer{contract: contract}}, nil
}

// NewIQuickSwapRouterCaller creates a new read-only instance of IQuickSwapRouter, bound to a specific deployed contract.
func NewIQuickSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*IQuickSwapRouterCaller, error) {
	contract, err := bindIQuickSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapRouterCaller{contract: contract}, nil
}

// NewIQuickSwapRouterTransactor creates a new write-only instance of IQuickSwapRouter, bound to a specific deployed contract.
func NewIQuickSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuickSwapRouterTransactor, error) {
	contract, err := bindIQuickSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapRouterTransactor{contract: contract}, nil
}

// NewIQuickSwapRouterFilterer creates a new log filterer instance of IQuickSwapRouter, bound to a specific deployed contract.
func NewIQuickSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuickSwapRouterFilterer, error) {
	contract, err := bindIQuickSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuickSwapRouterFilterer{contract: contract}, nil
}

// bindIQuickSwapRouter binds a generic wrapper to an already deployed contract.
func bindIQuickSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IQuickSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickSwapRouter *IQuickSwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickSwapRouter.Contract.IQuickSwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickSwapRouter *IQuickSwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.IQuickSwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickSwapRouter *IQuickSwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.IQuickSwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickSwapRouter *IQuickSwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickSwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickSwapRouter *IQuickSwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickSwapRouter *IQuickSwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.contract.Transact(opts, method, params...)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) AlgebraSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "algebraSwapCallback", amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapRouter *IQuickSwapRouterSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.AlgebraSwapCallback(&_IQuickSwapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.AlgebraSwapCallback(&_IQuickSwapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) ExactInput(opts *bind.TransactOpts, params IQuickSwapRouterExactInputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterSession) ExactInput(params IQuickSwapRouterExactInputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInput(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) ExactInput(params IQuickSwapRouterExactInputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInput(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterSession) ExactInputSingle(params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInputSingle(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) ExactInputSingle(params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInputSingle(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) ExactInputSingleSupportingFeeOnTransferTokens(opts *bind.TransactOpts, params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "exactInputSingleSupportingFeeOnTransferTokens", params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterSession) ExactInputSingleSupportingFeeOnTransferTokens(params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) ExactInputSingleSupportingFeeOnTransferTokens(params IQuickSwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params IQuickSwapRouterExactOutputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterSession) ExactOutput(params IQuickSwapRouterExactOutputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactOutput(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) ExactOutput(params IQuickSwapRouterExactOutputParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactOutput(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params IQuickSwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterSession) ExactOutputSingle(params IQuickSwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactOutputSingle(&_IQuickSwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_IQuickSwapRouter *IQuickSwapRouterTransactorSession) ExactOutputSingle(params IQuickSwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IQuickSwapRouter.Contract.ExactOutputSingle(&_IQuickSwapRouter.TransactOpts, params)
}
