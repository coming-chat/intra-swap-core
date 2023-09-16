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

// ICurveFiMetaData contains all meta data concerning the ICurveFi contract.
var ICurveFiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dy\",\"type\":\"uint256\"}],\"name\":\"get_dx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dy\",\"type\":\"uint256\"}],\"name\":\"get_dx_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICurveFiABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurveFiMetaData.ABI instead.
var ICurveFiABI = ICurveFiMetaData.ABI

// ICurveFi is an auto generated Go binding around an Ethereum contract.
type ICurveFi struct {
	ICurveFiCaller     // Read-only binding to the contract
	ICurveFiTransactor // Write-only binding to the contract
	ICurveFiFilterer   // Log filterer for contract events
}

// ICurveFiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurveFiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurveFiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurveFiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurveFiSession struct {
	Contract     *ICurveFi         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurveFiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurveFiCallerSession struct {
	Contract *ICurveFiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ICurveFiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurveFiTransactorSession struct {
	Contract     *ICurveFiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICurveFiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurveFiRaw struct {
	Contract *ICurveFi // Generic contract binding to access the raw methods on
}

// ICurveFiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurveFiCallerRaw struct {
	Contract *ICurveFiCaller // Generic read-only contract binding to access the raw methods on
}

// ICurveFiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurveFiTransactorRaw struct {
	Contract *ICurveFiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurveFi creates a new instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFi(address common.Address, backend bind.ContractBackend) (*ICurveFi, error) {
	contract, err := bindICurveFi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurveFi{ICurveFiCaller: ICurveFiCaller{contract: contract}, ICurveFiTransactor: ICurveFiTransactor{contract: contract}, ICurveFiFilterer: ICurveFiFilterer{contract: contract}}, nil
}

// NewICurveFiCaller creates a new read-only instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiCaller(address common.Address, caller bind.ContractCaller) (*ICurveFiCaller, error) {
	contract, err := bindICurveFi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveFiCaller{contract: contract}, nil
}

// NewICurveFiTransactor creates a new write-only instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurveFiTransactor, error) {
	contract, err := bindICurveFi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveFiTransactor{contract: contract}, nil
}

// NewICurveFiFilterer creates a new log filterer instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurveFiFilterer, error) {
	contract, err := bindICurveFi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurveFiFilterer{contract: contract}, nil
}

// bindICurveFi binds a generic wrapper to an already deployed contract.
func bindICurveFi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICurveFiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveFi *ICurveFiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveFi.Contract.ICurveFiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveFi *ICurveFiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveFi.Contract.ICurveFiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveFi *ICurveFiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveFi.Contract.ICurveFiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveFi *ICurveFiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveFi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveFi *ICurveFiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveFi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveFi *ICurveFiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveFi.Contract.contract.Transact(opts, method, params...)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDx(&_ICurveFi.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDx(&_ICurveFi.CallOpts, i, j, dy)
}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetDxUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_dx_underlying", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetDxUnderlying(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDxUnderlying(&_ICurveFi.CallOpts, i, j, dy)
}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetDxUnderlying(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDxUnderlying(&_ICurveFi.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDy(&_ICurveFi.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDy(&_ICurveFi.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDyUnderlying(&_ICurveFi.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDyUnderlying(&_ICurveFi.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurveFi.Contract.GetVirtualPrice(&_ICurveFi.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurveFi.Contract.GetVirtualPrice(&_ICurveFi.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange(&_ICurveFi.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange(&_ICurveFi.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.ExchangeUnderlying(&_ICurveFi.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurveFi *ICurveFiTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.ExchangeUnderlying(&_ICurveFi.TransactOpts, i, j, dx, min_dy)
}
