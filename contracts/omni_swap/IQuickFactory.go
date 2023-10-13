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

// IQuickFactoryMetaData contains all meta data concerning the IQuickFactory contract.
var IQuickFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFarmingAddress\",\"type\":\"address\"}],\"name\":\"FarmingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"FeeConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Owner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAddress\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"setBaseFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmingAddress\",\"type\":\"address\"}],\"name\":\"setFarmingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IQuickFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IQuickFactoryMetaData.ABI instead.
var IQuickFactoryABI = IQuickFactoryMetaData.ABI

// IQuickFactory is an auto generated Go binding around an Ethereum contract.
type IQuickFactory struct {
	IQuickFactoryCaller     // Read-only binding to the contract
	IQuickFactoryTransactor // Write-only binding to the contract
	IQuickFactoryFilterer   // Log filterer for contract events
}

// IQuickFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuickFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuickFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuickFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuickFactorySession struct {
	Contract     *IQuickFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQuickFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuickFactoryCallerSession struct {
	Contract *IQuickFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IQuickFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuickFactoryTransactorSession struct {
	Contract     *IQuickFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IQuickFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuickFactoryRaw struct {
	Contract *IQuickFactory // Generic contract binding to access the raw methods on
}

// IQuickFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuickFactoryCallerRaw struct {
	Contract *IQuickFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IQuickFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuickFactoryTransactorRaw struct {
	Contract *IQuickFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuickFactory creates a new instance of IQuickFactory, bound to a specific deployed contract.
func NewIQuickFactory(address common.Address, backend bind.ContractBackend) (*IQuickFactory, error) {
	contract, err := bindIQuickFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuickFactory{IQuickFactoryCaller: IQuickFactoryCaller{contract: contract}, IQuickFactoryTransactor: IQuickFactoryTransactor{contract: contract}, IQuickFactoryFilterer: IQuickFactoryFilterer{contract: contract}}, nil
}

// NewIQuickFactoryCaller creates a new read-only instance of IQuickFactory, bound to a specific deployed contract.
func NewIQuickFactoryCaller(address common.Address, caller bind.ContractCaller) (*IQuickFactoryCaller, error) {
	contract, err := bindIQuickFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryCaller{contract: contract}, nil
}

// NewIQuickFactoryTransactor creates a new write-only instance of IQuickFactory, bound to a specific deployed contract.
func NewIQuickFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuickFactoryTransactor, error) {
	contract, err := bindIQuickFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryTransactor{contract: contract}, nil
}

// NewIQuickFactoryFilterer creates a new log filterer instance of IQuickFactory, bound to a specific deployed contract.
func NewIQuickFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuickFactoryFilterer, error) {
	contract, err := bindIQuickFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryFilterer{contract: contract}, nil
}

// bindIQuickFactory binds a generic wrapper to an already deployed contract.
func bindIQuickFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IQuickFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickFactory *IQuickFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickFactory.Contract.IQuickFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickFactory *IQuickFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickFactory.Contract.IQuickFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickFactory *IQuickFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickFactory.Contract.IQuickFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickFactory *IQuickFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickFactory *IQuickFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickFactory *IQuickFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickFactory.Contract.contract.Transact(opts, method, params...)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_IQuickFactory *IQuickFactoryCaller) FarmingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickFactory.contract.Call(opts, &out, "farmingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_IQuickFactory *IQuickFactorySession) FarmingAddress() (common.Address, error) {
	return _IQuickFactory.Contract.FarmingAddress(&_IQuickFactory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_IQuickFactory *IQuickFactoryCallerSession) FarmingAddress() (common.Address, error) {
	return _IQuickFactory.Contract.FarmingAddress(&_IQuickFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IQuickFactory *IQuickFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IQuickFactory *IQuickFactorySession) Owner() (common.Address, error) {
	return _IQuickFactory.Contract.Owner(&_IQuickFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IQuickFactory *IQuickFactoryCallerSession) Owner() (common.Address, error) {
	return _IQuickFactory.Contract.Owner(&_IQuickFactory.CallOpts)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_IQuickFactory *IQuickFactoryCaller) PoolByPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IQuickFactory.contract.Call(opts, &out, "poolByPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_IQuickFactory *IQuickFactorySession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IQuickFactory.Contract.PoolByPair(&_IQuickFactory.CallOpts, tokenA, tokenB)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_IQuickFactory *IQuickFactoryCallerSession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IQuickFactory.Contract.PoolByPair(&_IQuickFactory.CallOpts, tokenA, tokenB)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_IQuickFactory *IQuickFactoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickFactory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_IQuickFactory *IQuickFactorySession) PoolDeployer() (common.Address, error) {
	return _IQuickFactory.Contract.PoolDeployer(&_IQuickFactory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_IQuickFactory *IQuickFactoryCallerSession) PoolDeployer() (common.Address, error) {
	return _IQuickFactory.Contract.PoolDeployer(&_IQuickFactory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_IQuickFactory *IQuickFactoryCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickFactory.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_IQuickFactory *IQuickFactorySession) VaultAddress() (common.Address, error) {
	return _IQuickFactory.Contract.VaultAddress(&_IQuickFactory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_IQuickFactory *IQuickFactoryCallerSession) VaultAddress() (common.Address, error) {
	return _IQuickFactory.Contract.VaultAddress(&_IQuickFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_IQuickFactory *IQuickFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IQuickFactory.contract.Transact(opts, "createPool", tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_IQuickFactory *IQuickFactorySession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.CreatePool(&_IQuickFactory.TransactOpts, tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_IQuickFactory *IQuickFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.CreatePool(&_IQuickFactory.TransactOpts, tokenA, tokenB)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_IQuickFactory *IQuickFactoryTransactor) SetBaseFeeConfiguration(opts *bind.TransactOpts, alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _IQuickFactory.contract.Transact(opts, "setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_IQuickFactory *IQuickFactorySession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetBaseFeeConfiguration(&_IQuickFactory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_IQuickFactory *IQuickFactoryTransactorSession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetBaseFeeConfiguration(&_IQuickFactory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_IQuickFactory *IQuickFactoryTransactor) SetFarmingAddress(opts *bind.TransactOpts, _farmingAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.contract.Transact(opts, "setFarmingAddress", _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_IQuickFactory *IQuickFactorySession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetFarmingAddress(&_IQuickFactory.TransactOpts, _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_IQuickFactory *IQuickFactoryTransactorSession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetFarmingAddress(&_IQuickFactory.TransactOpts, _farmingAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IQuickFactory *IQuickFactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IQuickFactory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IQuickFactory *IQuickFactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetOwner(&_IQuickFactory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IQuickFactory *IQuickFactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetOwner(&_IQuickFactory.TransactOpts, _owner)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_IQuickFactory *IQuickFactoryTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_IQuickFactory *IQuickFactorySession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetVaultAddress(&_IQuickFactory.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_IQuickFactory *IQuickFactoryTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _IQuickFactory.Contract.SetVaultAddress(&_IQuickFactory.TransactOpts, _vaultAddress)
}

// IQuickFactoryFarmingAddressIterator is returned from FilterFarmingAddress and is used to iterate over the raw logs and unpacked data for FarmingAddress events raised by the IQuickFactory contract.
type IQuickFactoryFarmingAddressIterator struct {
	Event *IQuickFactoryFarmingAddress // Event containing the contract specifics and raw log

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
func (it *IQuickFactoryFarmingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickFactoryFarmingAddress)
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
		it.Event = new(IQuickFactoryFarmingAddress)
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
func (it *IQuickFactoryFarmingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickFactoryFarmingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickFactoryFarmingAddress represents a FarmingAddress event raised by the IQuickFactory contract.
type IQuickFactoryFarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFarmingAddress is a free log retrieval operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_IQuickFactory *IQuickFactoryFilterer) FilterFarmingAddress(opts *bind.FilterOpts, newFarmingAddress []common.Address) (*IQuickFactoryFarmingAddressIterator, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _IQuickFactory.contract.FilterLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryFarmingAddressIterator{contract: _IQuickFactory.contract, event: "FarmingAddress", logs: logs, sub: sub}, nil
}

// WatchFarmingAddress is a free log subscription operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_IQuickFactory *IQuickFactoryFilterer) WatchFarmingAddress(opts *bind.WatchOpts, sink chan<- *IQuickFactoryFarmingAddress, newFarmingAddress []common.Address) (event.Subscription, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _IQuickFactory.contract.WatchLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickFactoryFarmingAddress)
				if err := _IQuickFactory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
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

// ParseFarmingAddress is a log parse operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_IQuickFactory *IQuickFactoryFilterer) ParseFarmingAddress(log types.Log) (*IQuickFactoryFarmingAddress, error) {
	event := new(IQuickFactoryFarmingAddress)
	if err := _IQuickFactory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickFactoryFeeConfigurationIterator is returned from FilterFeeConfiguration and is used to iterate over the raw logs and unpacked data for FeeConfiguration events raised by the IQuickFactory contract.
type IQuickFactoryFeeConfigurationIterator struct {
	Event *IQuickFactoryFeeConfiguration // Event containing the contract specifics and raw log

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
func (it *IQuickFactoryFeeConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickFactoryFeeConfiguration)
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
		it.Event = new(IQuickFactoryFeeConfiguration)
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
func (it *IQuickFactoryFeeConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickFactoryFeeConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickFactoryFeeConfiguration represents a FeeConfiguration event raised by the IQuickFactory contract.
type IQuickFactoryFeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeConfiguration is a free log retrieval operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_IQuickFactory *IQuickFactoryFilterer) FilterFeeConfiguration(opts *bind.FilterOpts) (*IQuickFactoryFeeConfigurationIterator, error) {

	logs, sub, err := _IQuickFactory.contract.FilterLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryFeeConfigurationIterator{contract: _IQuickFactory.contract, event: "FeeConfiguration", logs: logs, sub: sub}, nil
}

// WatchFeeConfiguration is a free log subscription operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_IQuickFactory *IQuickFactoryFilterer) WatchFeeConfiguration(opts *bind.WatchOpts, sink chan<- *IQuickFactoryFeeConfiguration) (event.Subscription, error) {

	logs, sub, err := _IQuickFactory.contract.WatchLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickFactoryFeeConfiguration)
				if err := _IQuickFactory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
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

// ParseFeeConfiguration is a log parse operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_IQuickFactory *IQuickFactoryFilterer) ParseFeeConfiguration(log types.Log) (*IQuickFactoryFeeConfiguration, error) {
	event := new(IQuickFactoryFeeConfiguration)
	if err := _IQuickFactory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickFactoryOwnerIterator is returned from FilterOwner and is used to iterate over the raw logs and unpacked data for Owner events raised by the IQuickFactory contract.
type IQuickFactoryOwnerIterator struct {
	Event *IQuickFactoryOwner // Event containing the contract specifics and raw log

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
func (it *IQuickFactoryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickFactoryOwner)
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
		it.Event = new(IQuickFactoryOwner)
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
func (it *IQuickFactoryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickFactoryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickFactoryOwner represents a Owner event raised by the IQuickFactory contract.
type IQuickFactoryOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwner is a free log retrieval operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_IQuickFactory *IQuickFactoryFilterer) FilterOwner(opts *bind.FilterOpts, newOwner []common.Address) (*IQuickFactoryOwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IQuickFactory.contract.FilterLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryOwnerIterator{contract: _IQuickFactory.contract, event: "Owner", logs: logs, sub: sub}, nil
}

// WatchOwner is a free log subscription operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_IQuickFactory *IQuickFactoryFilterer) WatchOwner(opts *bind.WatchOpts, sink chan<- *IQuickFactoryOwner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IQuickFactory.contract.WatchLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickFactoryOwner)
				if err := _IQuickFactory.contract.UnpackLog(event, "Owner", log); err != nil {
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

// ParseOwner is a log parse operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_IQuickFactory *IQuickFactoryFilterer) ParseOwner(log types.Log) (*IQuickFactoryOwner, error) {
	event := new(IQuickFactoryOwner)
	if err := _IQuickFactory.contract.UnpackLog(event, "Owner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickFactoryPoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the IQuickFactory contract.
type IQuickFactoryPoolIterator struct {
	Event *IQuickFactoryPool // Event containing the contract specifics and raw log

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
func (it *IQuickFactoryPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickFactoryPool)
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
		it.Event = new(IQuickFactoryPool)
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
func (it *IQuickFactoryPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickFactoryPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickFactoryPool represents a Pool event raised by the IQuickFactory contract.
type IQuickFactoryPool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_IQuickFactory *IQuickFactoryFilterer) FilterPool(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*IQuickFactoryPoolIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IQuickFactory.contract.FilterLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryPoolIterator{contract: _IQuickFactory.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_IQuickFactory *IQuickFactoryFilterer) WatchPool(opts *bind.WatchOpts, sink chan<- *IQuickFactoryPool, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IQuickFactory.contract.WatchLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickFactoryPool)
				if err := _IQuickFactory.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_IQuickFactory *IQuickFactoryFilterer) ParsePool(log types.Log) (*IQuickFactoryPool, error) {
	event := new(IQuickFactoryPool)
	if err := _IQuickFactory.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickFactoryVaultAddressIterator is returned from FilterVaultAddress and is used to iterate over the raw logs and unpacked data for VaultAddress events raised by the IQuickFactory contract.
type IQuickFactoryVaultAddressIterator struct {
	Event *IQuickFactoryVaultAddress // Event containing the contract specifics and raw log

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
func (it *IQuickFactoryVaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickFactoryVaultAddress)
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
		it.Event = new(IQuickFactoryVaultAddress)
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
func (it *IQuickFactoryVaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickFactoryVaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickFactoryVaultAddress represents a VaultAddress event raised by the IQuickFactory contract.
type IQuickFactoryVaultAddress struct {
	NewVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultAddress is a free log retrieval operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_IQuickFactory *IQuickFactoryFilterer) FilterVaultAddress(opts *bind.FilterOpts, newVaultAddress []common.Address) (*IQuickFactoryVaultAddressIterator, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _IQuickFactory.contract.FilterLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &IQuickFactoryVaultAddressIterator{contract: _IQuickFactory.contract, event: "VaultAddress", logs: logs, sub: sub}, nil
}

// WatchVaultAddress is a free log subscription operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_IQuickFactory *IQuickFactoryFilterer) WatchVaultAddress(opts *bind.WatchOpts, sink chan<- *IQuickFactoryVaultAddress, newVaultAddress []common.Address) (event.Subscription, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _IQuickFactory.contract.WatchLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickFactoryVaultAddress)
				if err := _IQuickFactory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
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

// ParseVaultAddress is a log parse operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_IQuickFactory *IQuickFactoryFilterer) ParseVaultAddress(log types.Log) (*IQuickFactoryVaultAddress, error) {
	event := new(IQuickFactoryVaultAddress)
	if err := _IQuickFactory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
