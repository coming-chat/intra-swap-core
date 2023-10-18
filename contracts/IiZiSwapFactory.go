// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IiZiSwapFactoryMetaData contains all meta data concerning the IiZiSwapFactory contract.
var IiZiSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"pointDelta\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"NewPool\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"chargeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFeeChargePercent\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployPoolParams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"pointDelta\",\"type\":\"int24\"},{\"internalType\":\"uint24\",\"name\":\"feeChargePercent\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"pointDelta\",\"type\":\"uint24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"fee2pointDelta\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"pointDelta\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitOrderModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chargeReceiver\",\"type\":\"address\"}],\"name\":\"modifyChargeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_defaultFeeChargePercent\",\"type\":\"uint24\"}],\"name\":\"modifyDefaultFeeChargePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"}],\"name\":\"newPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapX2YModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapY2XModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IiZiSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IiZiSwapFactoryMetaData.ABI instead.
var IiZiSwapFactoryABI = IiZiSwapFactoryMetaData.ABI

// IiZiSwapFactory is an auto generated Go binding around an Ethereum contract.
type IiZiSwapFactory struct {
	IiZiSwapFactoryCaller     // Read-only binding to the contract
	IiZiSwapFactoryTransactor // Write-only binding to the contract
	IiZiSwapFactoryFilterer   // Log filterer for contract events
}

// IiZiSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IiZiSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IiZiSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IiZiSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IiZiSwapFactorySession struct {
	Contract     *IiZiSwapFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IiZiSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IiZiSwapFactoryCallerSession struct {
	Contract *IiZiSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IiZiSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IiZiSwapFactoryTransactorSession struct {
	Contract     *IiZiSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IiZiSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IiZiSwapFactoryRaw struct {
	Contract *IiZiSwapFactory // Generic contract binding to access the raw methods on
}

// IiZiSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IiZiSwapFactoryCallerRaw struct {
	Contract *IiZiSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IiZiSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IiZiSwapFactoryTransactorRaw struct {
	Contract *IiZiSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIiZiSwapFactory creates a new instance of IiZiSwapFactory, bound to a specific deployed contract.
func NewIiZiSwapFactory(address common.Address, backend bind.ContractBackend) (*IiZiSwapFactory, error) {
	contract, err := bindIiZiSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFactory{IiZiSwapFactoryCaller: IiZiSwapFactoryCaller{contract: contract}, IiZiSwapFactoryTransactor: IiZiSwapFactoryTransactor{contract: contract}, IiZiSwapFactoryFilterer: IiZiSwapFactoryFilterer{contract: contract}}, nil
}

// NewIiZiSwapFactoryCaller creates a new read-only instance of IiZiSwapFactory, bound to a specific deployed contract.
func NewIiZiSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*IiZiSwapFactoryCaller, error) {
	contract, err := bindIiZiSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFactoryCaller{contract: contract}, nil
}

// NewIiZiSwapFactoryTransactor creates a new write-only instance of IiZiSwapFactory, bound to a specific deployed contract.
func NewIiZiSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IiZiSwapFactoryTransactor, error) {
	contract, err := bindIiZiSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFactoryTransactor{contract: contract}, nil
}

// NewIiZiSwapFactoryFilterer creates a new log filterer instance of IiZiSwapFactory, bound to a specific deployed contract.
func NewIiZiSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IiZiSwapFactoryFilterer, error) {
	contract, err := bindIiZiSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFactoryFilterer{contract: contract}, nil
}

// bindIiZiSwapFactory binds a generic wrapper to an already deployed contract.
func bindIiZiSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IiZiSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapFactory *IiZiSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapFactory.Contract.IiZiSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapFactory *IiZiSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.IiZiSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapFactory *IiZiSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.IiZiSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapFactory *IiZiSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// ChargeReceiver is a free data retrieval call binding the contract method 0xd8cd50e2.
//
// Solidity: function chargeReceiver() view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryCaller) ChargeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IiZiSwapFactory.contract.Call(opts, &out, "chargeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChargeReceiver is a free data retrieval call binding the contract method 0xd8cd50e2.
//
// Solidity: function chargeReceiver() view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) ChargeReceiver() (common.Address, error) {
	return _IiZiSwapFactory.Contract.ChargeReceiver(&_IiZiSwapFactory.CallOpts)
}

// ChargeReceiver is a free data retrieval call binding the contract method 0xd8cd50e2.
//
// Solidity: function chargeReceiver() view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryCallerSession) ChargeReceiver() (common.Address, error) {
	return _IiZiSwapFactory.Contract.ChargeReceiver(&_IiZiSwapFactory.CallOpts)
}

// DeployPoolParams is a free data retrieval call binding the contract method 0xcece24fe.
//
// Solidity: function deployPoolParams() view returns(address tokenX, address tokenY, uint24 fee, int24 currentPoint, int24 pointDelta, uint24 feeChargePercent)
func (_IiZiSwapFactory *IiZiSwapFactoryCaller) DeployPoolParams(opts *bind.CallOpts) (struct {
	TokenX           common.Address
	TokenY           common.Address
	Fee              *big.Int
	CurrentPoint     *big.Int
	PointDelta       *big.Int
	FeeChargePercent *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapFactory.contract.Call(opts, &out, "deployPoolParams")

	outstruct := new(struct {
		TokenX           common.Address
		TokenY           common.Address
		Fee              *big.Int
		CurrentPoint     *big.Int
		PointDelta       *big.Int
		FeeChargePercent *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenX = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenY = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PointDelta = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FeeChargePercent = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DeployPoolParams is a free data retrieval call binding the contract method 0xcece24fe.
//
// Solidity: function deployPoolParams() view returns(address tokenX, address tokenY, uint24 fee, int24 currentPoint, int24 pointDelta, uint24 feeChargePercent)
func (_IiZiSwapFactory *IiZiSwapFactorySession) DeployPoolParams() (struct {
	TokenX           common.Address
	TokenY           common.Address
	Fee              *big.Int
	CurrentPoint     *big.Int
	PointDelta       *big.Int
	FeeChargePercent *big.Int
}, error) {
	return _IiZiSwapFactory.Contract.DeployPoolParams(&_IiZiSwapFactory.CallOpts)
}

// DeployPoolParams is a free data retrieval call binding the contract method 0xcece24fe.
//
// Solidity: function deployPoolParams() view returns(address tokenX, address tokenY, uint24 fee, int24 currentPoint, int24 pointDelta, uint24 feeChargePercent)
func (_IiZiSwapFactory *IiZiSwapFactoryCallerSession) DeployPoolParams() (struct {
	TokenX           common.Address
	TokenY           common.Address
	Fee              *big.Int
	CurrentPoint     *big.Int
	PointDelta       *big.Int
	FeeChargePercent *big.Int
}, error) {
	return _IiZiSwapFactory.Contract.DeployPoolParams(&_IiZiSwapFactory.CallOpts)
}

// Fee2pointDelta is a free data retrieval call binding the contract method 0x3ce8e8db.
//
// Solidity: function fee2pointDelta(uint24 fee) view returns(int24 pointDelta)
func (_IiZiSwapFactory *IiZiSwapFactoryCaller) Fee2pointDelta(opts *bind.CallOpts, fee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapFactory.contract.Call(opts, &out, "fee2pointDelta", fee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee2pointDelta is a free data retrieval call binding the contract method 0x3ce8e8db.
//
// Solidity: function fee2pointDelta(uint24 fee) view returns(int24 pointDelta)
func (_IiZiSwapFactory *IiZiSwapFactorySession) Fee2pointDelta(fee *big.Int) (*big.Int, error) {
	return _IiZiSwapFactory.Contract.Fee2pointDelta(&_IiZiSwapFactory.CallOpts, fee)
}

// Fee2pointDelta is a free data retrieval call binding the contract method 0x3ce8e8db.
//
// Solidity: function fee2pointDelta(uint24 fee) view returns(int24 pointDelta)
func (_IiZiSwapFactory *IiZiSwapFactoryCallerSession) Fee2pointDelta(fee *big.Int) (*big.Int, error) {
	return _IiZiSwapFactory.Contract.Fee2pointDelta(&_IiZiSwapFactory.CallOpts, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IiZiSwapFactory.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _IiZiSwapFactory.Contract.Pool(&_IiZiSwapFactory.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _IiZiSwapFactory.Contract.Pool(&_IiZiSwapFactory.CallOpts, tokenX, tokenY, fee)
}

// DefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x59950c86.
//
// Solidity: function defaultFeeChargePercent() returns(uint24)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) DefaultFeeChargePercent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "defaultFeeChargePercent")
}

// DefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x59950c86.
//
// Solidity: function defaultFeeChargePercent() returns(uint24)
func (_IiZiSwapFactory *IiZiSwapFactorySession) DefaultFeeChargePercent() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.DefaultFeeChargePercent(&_IiZiSwapFactory.TransactOpts)
}

// DefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x59950c86.
//
// Solidity: function defaultFeeChargePercent() returns(uint24)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) DefaultFeeChargePercent() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.DefaultFeeChargePercent(&_IiZiSwapFactory.TransactOpts)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x10a17ee8.
//
// Solidity: function enableFeeAmount(uint24 fee, uint24 pointDelta) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, pointDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "enableFeeAmount", fee, pointDelta)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x10a17ee8.
//
// Solidity: function enableFeeAmount(uint24 fee, uint24 pointDelta) returns()
func (_IiZiSwapFactory *IiZiSwapFactorySession) EnableFeeAmount(fee *big.Int, pointDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.EnableFeeAmount(&_IiZiSwapFactory.TransactOpts, fee, pointDelta)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x10a17ee8.
//
// Solidity: function enableFeeAmount(uint24 fee, uint24 pointDelta) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) EnableFeeAmount(fee *big.Int, pointDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.EnableFeeAmount(&_IiZiSwapFactory.TransactOpts, fee, pointDelta)
}

// FlashModule is a paid mutator transaction binding the contract method 0x5deef20a.
//
// Solidity: function flashModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) FlashModule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "flashModule")
}

// FlashModule is a paid mutator transaction binding the contract method 0x5deef20a.
//
// Solidity: function flashModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) FlashModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.FlashModule(&_IiZiSwapFactory.TransactOpts)
}

// FlashModule is a paid mutator transaction binding the contract method 0x5deef20a.
//
// Solidity: function flashModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) FlashModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.FlashModule(&_IiZiSwapFactory.TransactOpts)
}

// LimitOrderModule is a paid mutator transaction binding the contract method 0x476476e0.
//
// Solidity: function limitOrderModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) LimitOrderModule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "limitOrderModule")
}

// LimitOrderModule is a paid mutator transaction binding the contract method 0x476476e0.
//
// Solidity: function limitOrderModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) LimitOrderModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.LimitOrderModule(&_IiZiSwapFactory.TransactOpts)
}

// LimitOrderModule is a paid mutator transaction binding the contract method 0x476476e0.
//
// Solidity: function limitOrderModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) LimitOrderModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.LimitOrderModule(&_IiZiSwapFactory.TransactOpts)
}

// LiquidityModule is a paid mutator transaction binding the contract method 0x400b6cdc.
//
// Solidity: function liquidityModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) LiquidityModule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "liquidityModule")
}

// LiquidityModule is a paid mutator transaction binding the contract method 0x400b6cdc.
//
// Solidity: function liquidityModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) LiquidityModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.LiquidityModule(&_IiZiSwapFactory.TransactOpts)
}

// LiquidityModule is a paid mutator transaction binding the contract method 0x400b6cdc.
//
// Solidity: function liquidityModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) LiquidityModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.LiquidityModule(&_IiZiSwapFactory.TransactOpts)
}

// ModifyChargeReceiver is a paid mutator transaction binding the contract method 0x3a6edcce.
//
// Solidity: function modifyChargeReceiver(address _chargeReceiver) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) ModifyChargeReceiver(opts *bind.TransactOpts, _chargeReceiver common.Address) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "modifyChargeReceiver", _chargeReceiver)
}

// ModifyChargeReceiver is a paid mutator transaction binding the contract method 0x3a6edcce.
//
// Solidity: function modifyChargeReceiver(address _chargeReceiver) returns()
func (_IiZiSwapFactory *IiZiSwapFactorySession) ModifyChargeReceiver(_chargeReceiver common.Address) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.ModifyChargeReceiver(&_IiZiSwapFactory.TransactOpts, _chargeReceiver)
}

// ModifyChargeReceiver is a paid mutator transaction binding the contract method 0x3a6edcce.
//
// Solidity: function modifyChargeReceiver(address _chargeReceiver) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) ModifyChargeReceiver(_chargeReceiver common.Address) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.ModifyChargeReceiver(&_IiZiSwapFactory.TransactOpts, _chargeReceiver)
}

// ModifyDefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x66f629f0.
//
// Solidity: function modifyDefaultFeeChargePercent(uint24 _defaultFeeChargePercent) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) ModifyDefaultFeeChargePercent(opts *bind.TransactOpts, _defaultFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "modifyDefaultFeeChargePercent", _defaultFeeChargePercent)
}

// ModifyDefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x66f629f0.
//
// Solidity: function modifyDefaultFeeChargePercent(uint24 _defaultFeeChargePercent) returns()
func (_IiZiSwapFactory *IiZiSwapFactorySession) ModifyDefaultFeeChargePercent(_defaultFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.ModifyDefaultFeeChargePercent(&_IiZiSwapFactory.TransactOpts, _defaultFeeChargePercent)
}

// ModifyDefaultFeeChargePercent is a paid mutator transaction binding the contract method 0x66f629f0.
//
// Solidity: function modifyDefaultFeeChargePercent(uint24 _defaultFeeChargePercent) returns()
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) ModifyDefaultFeeChargePercent(_defaultFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.ModifyDefaultFeeChargePercent(&_IiZiSwapFactory.TransactOpts, _defaultFeeChargePercent)
}

// NewPool is a paid mutator transaction binding the contract method 0x78eda67b.
//
// Solidity: function newPool(address tokenX, address tokenY, uint24 fee, int24 currentPoint) returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) NewPool(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, currentPoint *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "newPool", tokenX, tokenY, fee, currentPoint)
}

// NewPool is a paid mutator transaction binding the contract method 0x78eda67b.
//
// Solidity: function newPool(address tokenX, address tokenY, uint24 fee, int24 currentPoint) returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) NewPool(tokenX common.Address, tokenY common.Address, fee *big.Int, currentPoint *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.NewPool(&_IiZiSwapFactory.TransactOpts, tokenX, tokenY, fee, currentPoint)
}

// NewPool is a paid mutator transaction binding the contract method 0x78eda67b.
//
// Solidity: function newPool(address tokenX, address tokenY, uint24 fee, int24 currentPoint) returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) NewPool(tokenX common.Address, tokenY common.Address, fee *big.Int, currentPoint *big.Int) (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.NewPool(&_IiZiSwapFactory.TransactOpts, tokenX, tokenY, fee, currentPoint)
}

// SwapX2YModule is a paid mutator transaction binding the contract method 0x254ace8f.
//
// Solidity: function swapX2YModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) SwapX2YModule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "swapX2YModule")
}

// SwapX2YModule is a paid mutator transaction binding the contract method 0x254ace8f.
//
// Solidity: function swapX2YModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) SwapX2YModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.SwapX2YModule(&_IiZiSwapFactory.TransactOpts)
}

// SwapX2YModule is a paid mutator transaction binding the contract method 0x254ace8f.
//
// Solidity: function swapX2YModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) SwapX2YModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.SwapX2YModule(&_IiZiSwapFactory.TransactOpts)
}

// SwapY2XModule is a paid mutator transaction binding the contract method 0x86df77de.
//
// Solidity: function swapY2XModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactor) SwapY2XModule(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapFactory.contract.Transact(opts, "swapY2XModule")
}

// SwapY2XModule is a paid mutator transaction binding the contract method 0x86df77de.
//
// Solidity: function swapY2XModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactorySession) SwapY2XModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.SwapY2XModule(&_IiZiSwapFactory.TransactOpts)
}

// SwapY2XModule is a paid mutator transaction binding the contract method 0x86df77de.
//
// Solidity: function swapY2XModule() returns(address)
func (_IiZiSwapFactory *IiZiSwapFactoryTransactorSession) SwapY2XModule() (*types.Transaction, error) {
	return _IiZiSwapFactory.Contract.SwapY2XModule(&_IiZiSwapFactory.TransactOpts)
}

// IiZiSwapFactoryNewPoolIterator is returned from FilterNewPool and is used to iterate over the raw logs and unpacked data for NewPool events raised by the IiZiSwapFactory contract.
type IiZiSwapFactoryNewPoolIterator struct {
	Event *IiZiSwapFactoryNewPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IiZiSwapFactoryNewPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapFactoryNewPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IiZiSwapFactoryNewPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IiZiSwapFactoryNewPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapFactoryNewPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapFactoryNewPool represents a NewPool event raised by the IiZiSwapFactory contract.
type IiZiSwapFactoryNewPool struct {
	TokenX     common.Address
	TokenY     common.Address
	Fee        *big.Int
	PointDelta *big.Int
	Pool       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewPool is a free log retrieval operation binding the contract event 0xf04da67755adf58739649e2fb9949a6328518141b7ac9e44aa10320688b04900.
//
// Solidity: event NewPool(address indexed tokenX, address indexed tokenY, uint24 indexed fee, uint24 pointDelta, address pool)
func (_IiZiSwapFactory *IiZiSwapFactoryFilterer) FilterNewPool(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (*IiZiSwapFactoryNewPoolIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IiZiSwapFactory.contract.FilterLogs(opts, "NewPool", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFactoryNewPoolIterator{contract: _IiZiSwapFactory.contract, event: "NewPool", logs: logs, sub: sub}, nil
}

// WatchNewPool is a free log subscription operation binding the contract event 0xf04da67755adf58739649e2fb9949a6328518141b7ac9e44aa10320688b04900.
//
// Solidity: event NewPool(address indexed tokenX, address indexed tokenY, uint24 indexed fee, uint24 pointDelta, address pool)
func (_IiZiSwapFactory *IiZiSwapFactoryFilterer) WatchNewPool(opts *bind.WatchOpts, sink chan<- *IiZiSwapFactoryNewPool, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IiZiSwapFactory.contract.WatchLogs(opts, "NewPool", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapFactoryNewPool)
				if err := _IiZiSwapFactory.contract.UnpackLog(event, "NewPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPool is a log parse operation binding the contract event 0xf04da67755adf58739649e2fb9949a6328518141b7ac9e44aa10320688b04900.
//
// Solidity: event NewPool(address indexed tokenX, address indexed tokenY, uint24 indexed fee, uint24 pointDelta, address pool)
func (_IiZiSwapFactory *IiZiSwapFactoryFilterer) ParseNewPool(log types.Log) (*IiZiSwapFactoryNewPool, error) {
	event := new(IiZiSwapFactoryNewPool)
	if err := _IiZiSwapFactory.contract.UnpackLog(event, "NewPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
