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

// ISyncSwapRouterSplitPermitParams is an auto generated low-level Go binding around an user-defined struct.
type ISyncSwapRouterSplitPermitParams struct {
	Token         common.Address
	ApproveAmount *big.Int
	Deadline      *big.Int
	V             uint8
	R             [32]byte
	S             [32]byte
}

// ISyncSwapRouterSwapPath is an auto generated low-level Go binding around an user-defined struct.
type ISyncSwapRouterSwapPath struct {
	Steps    []ISyncSwapRouterSwapStep
	TokenIn  common.Address
	AmountIn *big.Int
}

// ISyncSwapRouterSwapStep is an auto generated low-level Go binding around an user-defined struct.
type ISyncSwapRouterSwapStep struct {
	Pool         common.Address
	Data         []byte
	Callback     common.Address
	CallbackData []byte
}

// ISyncSwapRouterMetaData contains all meta data concerning the ISyncSwapRouter contract.
var ISyncSwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"internalType\":\"structISyncSwapRouter.SwapStep[]\",\"name\":\"steps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structISyncSwapRouter.SwapPath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"internalType\":\"structISyncSwapRouter.SwapStep[]\",\"name\":\"steps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structISyncSwapRouter.SwapPath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structISyncSwapRouter.SplitPermitParams\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"swapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISyncSwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISyncSwapRouterMetaData.ABI instead.
var ISyncSwapRouterABI = ISyncSwapRouterMetaData.ABI

// ISyncSwapRouter is an auto generated Go binding around an Ethereum contract.
type ISyncSwapRouter struct {
	ISyncSwapRouterCaller     // Read-only binding to the contract
	ISyncSwapRouterTransactor // Write-only binding to the contract
	ISyncSwapRouterFilterer   // Log filterer for contract events
}

// ISyncSwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISyncSwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISyncSwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISyncSwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISyncSwapRouterSession struct {
	Contract     *ISyncSwapRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISyncSwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISyncSwapRouterCallerSession struct {
	Contract *ISyncSwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISyncSwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISyncSwapRouterTransactorSession struct {
	Contract     *ISyncSwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISyncSwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISyncSwapRouterRaw struct {
	Contract *ISyncSwapRouter // Generic contract binding to access the raw methods on
}

// ISyncSwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISyncSwapRouterCallerRaw struct {
	Contract *ISyncSwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ISyncSwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISyncSwapRouterTransactorRaw struct {
	Contract *ISyncSwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISyncSwapRouter creates a new instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouter(address common.Address, backend bind.ContractBackend) (*ISyncSwapRouter, error) {
	contract, err := bindISyncSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouter{ISyncSwapRouterCaller: ISyncSwapRouterCaller{contract: contract}, ISyncSwapRouterTransactor: ISyncSwapRouterTransactor{contract: contract}, ISyncSwapRouterFilterer: ISyncSwapRouterFilterer{contract: contract}}, nil
}

// NewISyncSwapRouterCaller creates a new read-only instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*ISyncSwapRouterCaller, error) {
	contract, err := bindISyncSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterCaller{contract: contract}, nil
}

// NewISyncSwapRouterTransactor creates a new write-only instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISyncSwapRouterTransactor, error) {
	contract, err := bindISyncSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterTransactor{contract: contract}, nil
}

// NewISyncSwapRouterFilterer creates a new log filterer instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISyncSwapRouterFilterer, error) {
	contract, err := bindISyncSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterFilterer{contract: contract}, nil
}

// bindISyncSwapRouter binds a generic wrapper to an already deployed contract.
func bindISyncSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISyncSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapRouter *ISyncSwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapRouter *ISyncSwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapRouter *ISyncSwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.contract.Transact(opts, method, params...)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapRouter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterSession) Vault() (common.Address, error) {
	return _ISyncSwapRouter.Contract.Vault(&_ISyncSwapRouter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterCallerSession) Vault() (common.Address, error) {
	return _ISyncSwapRouter.Contract.Vault(&_ISyncSwapRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapRouter.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterSession) WETH() (common.Address, error) {
	return _ISyncSwapRouter.Contract.WETH(&_ISyncSwapRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterCallerSession) WETH() (common.Address, error) {
	return _ISyncSwapRouter.Contract.WETH(&_ISyncSwapRouter.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) payable returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterTransactor) CreatePool(opts *bind.TransactOpts, factory common.Address, data []byte) (*types.Transaction, error) {
	return _ISyncSwapRouter.contract.Transact(opts, "createPool", factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) payable returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterSession) CreatePool(factory common.Address, data []byte) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.CreatePool(&_ISyncSwapRouter.TransactOpts, factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) payable returns(address)
func (_ISyncSwapRouter *ISyncSwapRouterTransactorSession) CreatePool(factory common.Address, data []byte) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.CreatePool(&_ISyncSwapRouter.TransactOpts, factory, data)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterTransactor) Swap(opts *bind.TransactOpts, paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISyncSwapRouter.contract.Transact(opts, "swap", paths, amountOutMin, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterSession) Swap(paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.Swap(&_ISyncSwapRouter.TransactOpts, paths, amountOutMin, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterTransactorSession) Swap(paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.Swap(&_ISyncSwapRouter.TransactOpts, paths, amountOutMin, deadline)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterTransactor) SwapWithPermit(opts *bind.TransactOpts, paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit ISyncSwapRouterSplitPermitParams) (*types.Transaction, error) {
	return _ISyncSwapRouter.contract.Transact(opts, "swapWithPermit", paths, amountOutMin, deadline, permit)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterSession) SwapWithPermit(paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit ISyncSwapRouterSplitPermitParams) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.SwapWithPermit(&_ISyncSwapRouter.TransactOpts, paths, amountOutMin, deadline, permit)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns(uint256 amountOut)
func (_ISyncSwapRouter *ISyncSwapRouterTransactorSession) SwapWithPermit(paths []ISyncSwapRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit ISyncSwapRouterSplitPermitParams) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.SwapWithPermit(&_ISyncSwapRouter.TransactOpts, paths, amountOutMin, deadline, permit)
}
