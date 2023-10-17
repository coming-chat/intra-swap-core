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

// ISwapRouter02ExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter02ExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouter02ExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter02ExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouter02ExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter02ExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouter02ExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter02ExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouter02MetaData contains all meta data concerning the ISwapRouter02 contract.
var ISwapRouter02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter02.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter02.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter02.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter02.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ISwapRouter02ABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapRouter02MetaData.ABI instead.
var ISwapRouter02ABI = ISwapRouter02MetaData.ABI

// ISwapRouter02 is an auto generated Go binding around an Ethereum contract.
type ISwapRouter02 struct {
	ISwapRouter02Caller     // Read-only binding to the contract
	ISwapRouter02Transactor // Write-only binding to the contract
	ISwapRouter02Filterer   // Log filterer for contract events
}

// ISwapRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapRouter02Session struct {
	Contract     *ISwapRouter02    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapRouter02CallerSession struct {
	Contract *ISwapRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISwapRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapRouter02TransactorSession struct {
	Contract     *ISwapRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISwapRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRouter02Raw struct {
	Contract *ISwapRouter02 // Generic contract binding to access the raw methods on
}

// ISwapRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapRouter02CallerRaw struct {
	Contract *ISwapRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// ISwapRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapRouter02TransactorRaw struct {
	Contract *ISwapRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapRouter02 creates a new instance of ISwapRouter02, bound to a specific deployed contract.
func NewISwapRouter02(address common.Address, backend bind.ContractBackend) (*ISwapRouter02, error) {
	contract, err := bindISwapRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapRouter02{ISwapRouter02Caller: ISwapRouter02Caller{contract: contract}, ISwapRouter02Transactor: ISwapRouter02Transactor{contract: contract}, ISwapRouter02Filterer: ISwapRouter02Filterer{contract: contract}}, nil
}

// NewISwapRouter02Caller creates a new read-only instance of ISwapRouter02, bound to a specific deployed contract.
func NewISwapRouter02Caller(address common.Address, caller bind.ContractCaller) (*ISwapRouter02Caller, error) {
	contract, err := bindISwapRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouter02Caller{contract: contract}, nil
}

// NewISwapRouter02Transactor creates a new write-only instance of ISwapRouter02, bound to a specific deployed contract.
func NewISwapRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*ISwapRouter02Transactor, error) {
	contract, err := bindISwapRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouter02Transactor{contract: contract}, nil
}

// NewISwapRouter02Filterer creates a new log filterer instance of ISwapRouter02, bound to a specific deployed contract.
func NewISwapRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*ISwapRouter02Filterer, error) {
	contract, err := bindISwapRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapRouter02Filterer{contract: contract}, nil
}

// bindISwapRouter02 binds a generic wrapper to an already deployed contract.
func bindISwapRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapRouter02MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouter02 *ISwapRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouter02.Contract.ISwapRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouter02 *ISwapRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ISwapRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouter02 *ISwapRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ISwapRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouter02 *ISwapRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouter02 *ISwapRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouter02 *ISwapRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.contract.Transact(opts, method, params...)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02Transactor) ExactInput(opts *bind.TransactOpts, params ISwapRouter02ExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter02.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02Session) ExactInput(params ISwapRouter02ExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactInput(&_ISwapRouter02.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02TransactorSession) ExactInput(params ISwapRouter02ExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactInput(&_ISwapRouter02.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02Transactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouter02ExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02Session) ExactInputSingle(params ISwapRouter02ExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactInputSingle(&_ISwapRouter02.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter02 *ISwapRouter02TransactorSession) ExactInputSingle(params ISwapRouter02ExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactInputSingle(&_ISwapRouter02.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02Transactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouter02ExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter02.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02Session) ExactOutput(params ISwapRouter02ExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactOutput(&_ISwapRouter02.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02TransactorSession) ExactOutput(params ISwapRouter02ExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactOutput(&_ISwapRouter02.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02Transactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouter02ExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02Session) ExactOutputSingle(params ISwapRouter02ExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactOutputSingle(&_ISwapRouter02.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter02 *ISwapRouter02TransactorSession) ExactOutputSingle(params ISwapRouter02ExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.ExactOutputSingle(&_ISwapRouter02.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[] results)
func (_ISwapRouter02 *ISwapRouter02Transactor) Multicall(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _ISwapRouter02.contract.Transact(opts, "multicall", deadline, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[] results)
func (_ISwapRouter02 *ISwapRouter02Session) Multicall(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.Multicall(&_ISwapRouter02.TransactOpts, deadline, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[] results)
func (_ISwapRouter02 *ISwapRouter02TransactorSession) Multicall(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _ISwapRouter02.Contract.Multicall(&_ISwapRouter02.TransactOpts, deadline, data)
}
