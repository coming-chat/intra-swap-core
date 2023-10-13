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

// IWombatRouterMetaData contains all meta data concerning the IWombatRouter contract.
var IWombatRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"poolPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"haircuts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"poolPath\",\"type\":\"address[]\"},{\"internalType\":\"int256\",\"name\":\"amountIn\",\"type\":\"int256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"haircuts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"poolPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minimumamountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNativeForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"poolPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumamountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"poolPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWombatRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IWombatRouterMetaData.ABI instead.
var IWombatRouterABI = IWombatRouterMetaData.ABI

// IWombatRouter is an auto generated Go binding around an Ethereum contract.
type IWombatRouter struct {
	IWombatRouterCaller     // Read-only binding to the contract
	IWombatRouterTransactor // Write-only binding to the contract
	IWombatRouterFilterer   // Log filterer for contract events
}

// IWombatRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWombatRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWombatRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWombatRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWombatRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWombatRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWombatRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWombatRouterSession struct {
	Contract     *IWombatRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWombatRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWombatRouterCallerSession struct {
	Contract *IWombatRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IWombatRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWombatRouterTransactorSession struct {
	Contract     *IWombatRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IWombatRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWombatRouterRaw struct {
	Contract *IWombatRouter // Generic contract binding to access the raw methods on
}

// IWombatRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWombatRouterCallerRaw struct {
	Contract *IWombatRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IWombatRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWombatRouterTransactorRaw struct {
	Contract *IWombatRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWombatRouter creates a new instance of IWombatRouter, bound to a specific deployed contract.
func NewIWombatRouter(address common.Address, backend bind.ContractBackend) (*IWombatRouter, error) {
	contract, err := bindIWombatRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWombatRouter{IWombatRouterCaller: IWombatRouterCaller{contract: contract}, IWombatRouterTransactor: IWombatRouterTransactor{contract: contract}, IWombatRouterFilterer: IWombatRouterFilterer{contract: contract}}, nil
}

// NewIWombatRouterCaller creates a new read-only instance of IWombatRouter, bound to a specific deployed contract.
func NewIWombatRouterCaller(address common.Address, caller bind.ContractCaller) (*IWombatRouterCaller, error) {
	contract, err := bindIWombatRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWombatRouterCaller{contract: contract}, nil
}

// NewIWombatRouterTransactor creates a new write-only instance of IWombatRouter, bound to a specific deployed contract.
func NewIWombatRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IWombatRouterTransactor, error) {
	contract, err := bindIWombatRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWombatRouterTransactor{contract: contract}, nil
}

// NewIWombatRouterFilterer creates a new log filterer instance of IWombatRouter, bound to a specific deployed contract.
func NewIWombatRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IWombatRouterFilterer, error) {
	contract, err := bindIWombatRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWombatRouterFilterer{contract: contract}, nil
}

// bindIWombatRouter binds a generic wrapper to an already deployed contract.
func bindIWombatRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWombatRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWombatRouter *IWombatRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWombatRouter.Contract.IWombatRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWombatRouter *IWombatRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWombatRouter.Contract.IWombatRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWombatRouter *IWombatRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWombatRouter.Contract.IWombatRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWombatRouter *IWombatRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWombatRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWombatRouter *IWombatRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWombatRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWombatRouter *IWombatRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWombatRouter.Contract.contract.Transact(opts, method, params...)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x3226288e.
//
// Solidity: function getAmountIn(address[] tokenPath, address[] poolPath, uint256 amountOut) view returns(uint256 amountIn, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterCaller) GetAmountIn(opts *bind.CallOpts, tokenPath []common.Address, poolPath []common.Address, amountOut *big.Int) (struct {
	AmountIn *big.Int
	Haircuts []*big.Int
}, error) {
	var out []interface{}
	err := _IWombatRouter.contract.Call(opts, &out, "getAmountIn", tokenPath, poolPath, amountOut)

	outstruct := new(struct {
		AmountIn *big.Int
		Haircuts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Haircuts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x3226288e.
//
// Solidity: function getAmountIn(address[] tokenPath, address[] poolPath, uint256 amountOut) view returns(uint256 amountIn, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterSession) GetAmountIn(tokenPath []common.Address, poolPath []common.Address, amountOut *big.Int) (struct {
	AmountIn *big.Int
	Haircuts []*big.Int
}, error) {
	return _IWombatRouter.Contract.GetAmountIn(&_IWombatRouter.CallOpts, tokenPath, poolPath, amountOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x3226288e.
//
// Solidity: function getAmountIn(address[] tokenPath, address[] poolPath, uint256 amountOut) view returns(uint256 amountIn, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterCallerSession) GetAmountIn(tokenPath []common.Address, poolPath []common.Address, amountOut *big.Int) (struct {
	AmountIn *big.Int
	Haircuts []*big.Int
}, error) {
	return _IWombatRouter.Contract.GetAmountIn(&_IWombatRouter.CallOpts, tokenPath, poolPath, amountOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xb9e598c1.
//
// Solidity: function getAmountOut(address[] tokenPath, address[] poolPath, int256 amountIn) view returns(uint256 amountOut, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterCaller) GetAmountOut(opts *bind.CallOpts, tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int) (struct {
	AmountOut *big.Int
	Haircuts  []*big.Int
}, error) {
	var out []interface{}
	err := _IWombatRouter.contract.Call(opts, &out, "getAmountOut", tokenPath, poolPath, amountIn)

	outstruct := new(struct {
		AmountOut *big.Int
		Haircuts  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Haircuts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xb9e598c1.
//
// Solidity: function getAmountOut(address[] tokenPath, address[] poolPath, int256 amountIn) view returns(uint256 amountOut, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterSession) GetAmountOut(tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int) (struct {
	AmountOut *big.Int
	Haircuts  []*big.Int
}, error) {
	return _IWombatRouter.Contract.GetAmountOut(&_IWombatRouter.CallOpts, tokenPath, poolPath, amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xb9e598c1.
//
// Solidity: function getAmountOut(address[] tokenPath, address[] poolPath, int256 amountIn) view returns(uint256 amountOut, uint256[] haircuts)
func (_IWombatRouter *IWombatRouterCallerSession) GetAmountOut(tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int) (struct {
	AmountOut *big.Int
	Haircuts  []*big.Int
}, error) {
	return _IWombatRouter.Contract.GetAmountOut(&_IWombatRouter.CallOpts, tokenPath, poolPath, amountIn)
}

// SwapExactNativeForTokens is a paid mutator transaction binding the contract method 0x018ad001.
//
// Solidity: function swapExactNativeForTokens(address[] tokenPath, address[] poolPath, uint256 minimumamountOut, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactor) SwapExactNativeForTokens(opts *bind.TransactOpts, tokenPath []common.Address, poolPath []common.Address, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.contract.Transact(opts, "swapExactNativeForTokens", tokenPath, poolPath, minimumamountOut, to, deadline)
}

// SwapExactNativeForTokens is a paid mutator transaction binding the contract method 0x018ad001.
//
// Solidity: function swapExactNativeForTokens(address[] tokenPath, address[] poolPath, uint256 minimumamountOut, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterSession) SwapExactNativeForTokens(tokenPath []common.Address, poolPath []common.Address, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactNativeForTokens(&_IWombatRouter.TransactOpts, tokenPath, poolPath, minimumamountOut, to, deadline)
}

// SwapExactNativeForTokens is a paid mutator transaction binding the contract method 0x018ad001.
//
// Solidity: function swapExactNativeForTokens(address[] tokenPath, address[] poolPath, uint256 minimumamountOut, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactorSession) SwapExactNativeForTokens(tokenPath []common.Address, poolPath []common.Address, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactNativeForTokens(&_IWombatRouter.TransactOpts, tokenPath, poolPath, minimumamountOut, to, deadline)
}

// SwapExactTokensForNative is a paid mutator transaction binding the contract method 0x85335cc0.
//
// Solidity: function swapExactTokensForNative(address[] tokenPath, address[] poolPath, uint256 amountIn, uint256 minimumamountOut, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactor) SwapExactTokensForNative(opts *bind.TransactOpts, tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.contract.Transact(opts, "swapExactTokensForNative", tokenPath, poolPath, amountIn, minimumamountOut, to, deadline)
}

// SwapExactTokensForNative is a paid mutator transaction binding the contract method 0x85335cc0.
//
// Solidity: function swapExactTokensForNative(address[] tokenPath, address[] poolPath, uint256 amountIn, uint256 minimumamountOut, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterSession) SwapExactTokensForNative(tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactTokensForNative(&_IWombatRouter.TransactOpts, tokenPath, poolPath, amountIn, minimumamountOut, to, deadline)
}

// SwapExactTokensForNative is a paid mutator transaction binding the contract method 0x85335cc0.
//
// Solidity: function swapExactTokensForNative(address[] tokenPath, address[] poolPath, uint256 amountIn, uint256 minimumamountOut, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactorSession) SwapExactTokensForNative(tokenPath []common.Address, poolPath []common.Address, amountIn *big.Int, minimumamountOut *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactTokensForNative(&_IWombatRouter.TransactOpts, tokenPath, poolPath, amountIn, minimumamountOut, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xd44107a4.
//
// Solidity: function swapExactTokensForTokens(address[] tokenPath, address[] poolPath, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, tokenPath []common.Address, poolPath []common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.contract.Transact(opts, "swapExactTokensForTokens", tokenPath, poolPath, fromAmount, minimumToAmount, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xd44107a4.
//
// Solidity: function swapExactTokensForTokens(address[] tokenPath, address[] poolPath, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterSession) SwapExactTokensForTokens(tokenPath []common.Address, poolPath []common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactTokensForTokens(&_IWombatRouter.TransactOpts, tokenPath, poolPath, fromAmount, minimumToAmount, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xd44107a4.
//
// Solidity: function swapExactTokensForTokens(address[] tokenPath, address[] poolPath, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 amountOut)
func (_IWombatRouter *IWombatRouterTransactorSession) SwapExactTokensForTokens(tokenPath []common.Address, poolPath []common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IWombatRouter.Contract.SwapExactTokensForTokens(&_IWombatRouter.TransactOpts, tokenPath, poolPath, fromAmount, minimumToAmount, to, deadline)
}
