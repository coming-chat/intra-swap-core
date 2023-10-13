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

// IPearlRouterroute is an auto generated low-level Go binding around an user-defined struct.
type IPearlRouterroute struct {
	From   common.Address
	To     common.Address
	Stable bool
}

// IPearlRouterMetaData contains all meta data concerning the IPearlRouter contract.
var IPearlRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIPearlRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIPearlRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIPearlRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPearlRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IPearlRouterMetaData.ABI instead.
var IPearlRouterABI = IPearlRouterMetaData.ABI

// IPearlRouter is an auto generated Go binding around an Ethereum contract.
type IPearlRouter struct {
	IPearlRouterCaller     // Read-only binding to the contract
	IPearlRouterTransactor // Write-only binding to the contract
	IPearlRouterFilterer   // Log filterer for contract events
}

// IPearlRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPearlRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPearlRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPearlRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPearlRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPearlRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPearlRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPearlRouterSession struct {
	Contract     *IPearlRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPearlRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPearlRouterCallerSession struct {
	Contract *IPearlRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPearlRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPearlRouterTransactorSession struct {
	Contract     *IPearlRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPearlRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPearlRouterRaw struct {
	Contract *IPearlRouter // Generic contract binding to access the raw methods on
}

// IPearlRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPearlRouterCallerRaw struct {
	Contract *IPearlRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IPearlRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPearlRouterTransactorRaw struct {
	Contract *IPearlRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPearlRouter creates a new instance of IPearlRouter, bound to a specific deployed contract.
func NewIPearlRouter(address common.Address, backend bind.ContractBackend) (*IPearlRouter, error) {
	contract, err := bindIPearlRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPearlRouter{IPearlRouterCaller: IPearlRouterCaller{contract: contract}, IPearlRouterTransactor: IPearlRouterTransactor{contract: contract}, IPearlRouterFilterer: IPearlRouterFilterer{contract: contract}}, nil
}

// NewIPearlRouterCaller creates a new read-only instance of IPearlRouter, bound to a specific deployed contract.
func NewIPearlRouterCaller(address common.Address, caller bind.ContractCaller) (*IPearlRouterCaller, error) {
	contract, err := bindIPearlRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPearlRouterCaller{contract: contract}, nil
}

// NewIPearlRouterTransactor creates a new write-only instance of IPearlRouter, bound to a specific deployed contract.
func NewIPearlRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPearlRouterTransactor, error) {
	contract, err := bindIPearlRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPearlRouterTransactor{contract: contract}, nil
}

// NewIPearlRouterFilterer creates a new log filterer instance of IPearlRouter, bound to a specific deployed contract.
func NewIPearlRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPearlRouterFilterer, error) {
	contract, err := bindIPearlRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPearlRouterFilterer{contract: contract}, nil
}

// bindIPearlRouter binds a generic wrapper to an already deployed contract.
func bindIPearlRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPearlRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPearlRouter *IPearlRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPearlRouter.Contract.IPearlRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPearlRouter *IPearlRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPearlRouter.Contract.IPearlRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPearlRouter *IPearlRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPearlRouter.Contract.IPearlRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPearlRouter *IPearlRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPearlRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPearlRouter *IPearlRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPearlRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPearlRouter *IPearlRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPearlRouter.Contract.contract.Transact(opts, method, params...)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactETHForTokens(&_IPearlRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactETHForTokens(&_IPearlRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactTokensForETH(&_IPearlRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactTokensForETH(&_IPearlRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactTokensForTokens(&_IPearlRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IPearlRouter *IPearlRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IPearlRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPearlRouter.Contract.SwapExactTokensForTokens(&_IPearlRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}
