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

// IQuickPoolMetaData contains all meta data concerning the IQuickPool contract.
var IQuickPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee1New\",\"type\":\"uint8\"}],\"name\":\"CommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"Incentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"liquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"LiquidityCooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeIncentive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"virtualPool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorageOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"}],\"name\":\"getInnerCumulatives\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"innerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"innerSecondsSpentPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"innerSecondsSpent\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"getTimepoints\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulatives\",\"type\":\"uint160[]\"},{\"internalType\":\"uint112[]\",\"name\":\"volatilityCumulatives\",\"type\":\"uint112[]\"},{\"internalType\":\"uint256[]\",\"name\":\"volumePerAvgLiquiditys\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalState\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timepointIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken1\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"cooldownInSeconds\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityActual\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastLiquidityAddTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"fees0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fees1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapSupportingFeeOnInputTokens\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityTotal\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"outerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"outerSecondsPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"outerSecondsSpent\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"timepoints\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulative\",\"type\":\"uint160\"},{\"internalType\":\"uint88\",\"name\":\"volatilityCumulative\",\"type\":\"uint88\"},{\"internalType\":\"int24\",\"name\":\"averageTick\",\"type\":\"int24\"},{\"internalType\":\"uint144\",\"name\":\"volumePerLiquidityCumulative\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth0Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth1Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IQuickPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IQuickPoolMetaData.ABI instead.
var IQuickPoolABI = IQuickPoolMetaData.ABI

// IQuickPool is an auto generated Go binding around an Ethereum contract.
type IQuickPool struct {
	IQuickPoolCaller     // Read-only binding to the contract
	IQuickPoolTransactor // Write-only binding to the contract
	IQuickPoolFilterer   // Log filterer for contract events
}

// IQuickPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuickPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuickPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuickPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuickPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuickPoolSession struct {
	Contract     *IQuickPool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQuickPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuickPoolCallerSession struct {
	Contract *IQuickPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IQuickPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuickPoolTransactorSession struct {
	Contract     *IQuickPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IQuickPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuickPoolRaw struct {
	Contract *IQuickPool // Generic contract binding to access the raw methods on
}

// IQuickPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuickPoolCallerRaw struct {
	Contract *IQuickPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IQuickPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuickPoolTransactorRaw struct {
	Contract *IQuickPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuickPool creates a new instance of IQuickPool, bound to a specific deployed contract.
func NewIQuickPool(address common.Address, backend bind.ContractBackend) (*IQuickPool, error) {
	contract, err := bindIQuickPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuickPool{IQuickPoolCaller: IQuickPoolCaller{contract: contract}, IQuickPoolTransactor: IQuickPoolTransactor{contract: contract}, IQuickPoolFilterer: IQuickPoolFilterer{contract: contract}}, nil
}

// NewIQuickPoolCaller creates a new read-only instance of IQuickPool, bound to a specific deployed contract.
func NewIQuickPoolCaller(address common.Address, caller bind.ContractCaller) (*IQuickPoolCaller, error) {
	contract, err := bindIQuickPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolCaller{contract: contract}, nil
}

// NewIQuickPoolTransactor creates a new write-only instance of IQuickPool, bound to a specific deployed contract.
func NewIQuickPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuickPoolTransactor, error) {
	contract, err := bindIQuickPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolTransactor{contract: contract}, nil
}

// NewIQuickPoolFilterer creates a new log filterer instance of IQuickPool, bound to a specific deployed contract.
func NewIQuickPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuickPoolFilterer, error) {
	contract, err := bindIQuickPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolFilterer{contract: contract}, nil
}

// bindIQuickPool binds a generic wrapper to an already deployed contract.
func bindIQuickPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IQuickPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickPool *IQuickPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickPool.Contract.IQuickPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickPool *IQuickPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickPool.Contract.IQuickPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickPool *IQuickPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickPool.Contract.IQuickPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuickPool *IQuickPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuickPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuickPool *IQuickPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuickPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuickPool *IQuickPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuickPool.Contract.contract.Transact(opts, method, params...)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_IQuickPool *IQuickPoolCaller) ActiveIncentive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "activeIncentive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_IQuickPool *IQuickPoolSession) ActiveIncentive() (common.Address, error) {
	return _IQuickPool.Contract.ActiveIncentive(&_IQuickPool.CallOpts)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_IQuickPool *IQuickPoolCallerSession) ActiveIncentive() (common.Address, error) {
	return _IQuickPool.Contract.ActiveIncentive(&_IQuickPool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_IQuickPool *IQuickPoolCaller) DataStorageOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "dataStorageOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_IQuickPool *IQuickPoolSession) DataStorageOperator() (common.Address, error) {
	return _IQuickPool.Contract.DataStorageOperator(&_IQuickPool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_IQuickPool *IQuickPoolCallerSession) DataStorageOperator() (common.Address, error) {
	return _IQuickPool.Contract.DataStorageOperator(&_IQuickPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IQuickPool *IQuickPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IQuickPool *IQuickPoolSession) Factory() (common.Address, error) {
	return _IQuickPool.Contract.Factory(&_IQuickPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IQuickPool *IQuickPoolCallerSession) Factory() (common.Address, error) {
	return _IQuickPool.Contract.Factory(&_IQuickPool.CallOpts)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_IQuickPool *IQuickPoolCaller) GetInnerCumulatives(opts *bind.CallOpts, bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "getInnerCumulatives", bottomTick, topTick)

	outstruct := new(struct {
		InnerTickCumulative           *big.Int
		InnerSecondsSpentPerLiquidity *big.Int
		InnerSecondsSpent             uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnerTickCumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpentPerLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpent = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_IQuickPool *IQuickPoolSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _IQuickPool.Contract.GetInnerCumulatives(&_IQuickPool.CallOpts, bottomTick, topTick)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_IQuickPool *IQuickPoolCallerSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _IQuickPool.Contract.GetInnerCumulatives(&_IQuickPool.CallOpts, bottomTick, topTick)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_IQuickPool *IQuickPoolCaller) GetTimepoints(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "getTimepoints", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                []*big.Int
		SecondsPerLiquidityCumulatives []*big.Int
		VolatilityCumulatives          []*big.Int
		VolumePerAvgLiquiditys         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulatives = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolatilityCumulatives = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolumePerAvgLiquiditys = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_IQuickPool *IQuickPoolSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _IQuickPool.Contract.GetTimepoints(&_IQuickPool.CallOpts, secondsAgos)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_IQuickPool *IQuickPoolCallerSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _IQuickPool.Contract.GetTimepoints(&_IQuickPool.CallOpts, secondsAgos)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_IQuickPool *IQuickPoolCaller) GlobalState(opts *bind.CallOpts) (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "globalState")

	outstruct := new(struct {
		Price              *big.Int
		Tick               *big.Int
		Fee                uint16
		TimepointIndex     uint16
		CommunityFeeToken0 uint8
		CommunityFeeToken1 uint8
		Unlocked           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TimepointIndex = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken0 = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.CommunityFeeToken1 = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_IQuickPool *IQuickPoolSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _IQuickPool.Contract.GlobalState(&_IQuickPool.CallOpts)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_IQuickPool *IQuickPoolCallerSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _IQuickPool.Contract.GlobalState(&_IQuickPool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IQuickPool *IQuickPoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IQuickPool *IQuickPoolSession) Liquidity() (*big.Int, error) {
	return _IQuickPool.Contract.Liquidity(&_IQuickPool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IQuickPool *IQuickPoolCallerSession) Liquidity() (*big.Int, error) {
	return _IQuickPool.Contract.Liquidity(&_IQuickPool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_IQuickPool *IQuickPoolCaller) LiquidityCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "liquidityCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_IQuickPool *IQuickPoolSession) LiquidityCooldown() (uint32, error) {
	return _IQuickPool.Contract.LiquidityCooldown(&_IQuickPool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_IQuickPool *IQuickPoolCallerSession) LiquidityCooldown() (uint32, error) {
	return _IQuickPool.Contract.LiquidityCooldown(&_IQuickPool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IQuickPool *IQuickPoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IQuickPool *IQuickPoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IQuickPool.Contract.MaxLiquidityPerTick(&_IQuickPool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IQuickPool *IQuickPoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IQuickPool.Contract.MaxLiquidityPerTick(&_IQuickPool.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_IQuickPool *IQuickPoolCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		LiquidityAmount           *big.Int
		LastLiquidityAddTimestamp uint32
		InnerFeeGrowth0Token      *big.Int
		InnerFeeGrowth1Token      *big.Int
		Fees0                     *big.Int
		Fees1                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLiquidityAddTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.InnerFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InnerFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_IQuickPool *IQuickPoolSession) Positions(key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _IQuickPool.Contract.Positions(&_IQuickPool.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_IQuickPool *IQuickPoolCallerSession) Positions(key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _IQuickPool.Contract.Positions(&_IQuickPool.CallOpts, key)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IQuickPool *IQuickPoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IQuickPool *IQuickPoolSession) TickSpacing() (*big.Int, error) {
	return _IQuickPool.Contract.TickSpacing(&_IQuickPool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IQuickPool *IQuickPoolCallerSession) TickSpacing() (*big.Int, error) {
	return _IQuickPool.Contract.TickSpacing(&_IQuickPool.CallOpts)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_IQuickPool *IQuickPoolCaller) TickTable(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "tickTable", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_IQuickPool *IQuickPoolSession) TickTable(wordPosition int16) (*big.Int, error) {
	return _IQuickPool.Contract.TickTable(&_IQuickPool.CallOpts, wordPosition)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_IQuickPool *IQuickPoolCallerSession) TickTable(wordPosition int16) (*big.Int, error) {
	return _IQuickPool.Contract.TickTable(&_IQuickPool.CallOpts, wordPosition)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_IQuickPool *IQuickPoolCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityTotal           *big.Int
		LiquidityDelta           *big.Int
		OuterFeeGrowth0Token     *big.Int
		OuterFeeGrowth1Token     *big.Int
		OuterTickCumulative      *big.Int
		OuterSecondsPerLiquidity *big.Int
		OuterSecondsSpent        uint32
		Initialized              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OuterTickCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsPerLiquidity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsSpent = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_IQuickPool *IQuickPoolSession) Ticks(tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _IQuickPool.Contract.Ticks(&_IQuickPool.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_IQuickPool *IQuickPoolCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _IQuickPool.Contract.Ticks(&_IQuickPool.CallOpts, tick)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_IQuickPool *IQuickPoolCaller) Timepoints(opts *bind.CallOpts, index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "timepoints", index)

	outstruct := new(struct {
		Initialized                   bool
		BlockTimestamp                uint32
		TickCumulative                *big.Int
		SecondsPerLiquidityCumulative *big.Int
		VolatilityCumulative          *big.Int
		AverageTick                   *big.Int
		VolumePerLiquidityCumulative  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulative = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VolatilityCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AverageTick = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VolumePerLiquidityCumulative = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_IQuickPool *IQuickPoolSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _IQuickPool.Contract.Timepoints(&_IQuickPool.CallOpts, index)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_IQuickPool *IQuickPoolCallerSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _IQuickPool.Contract.Timepoints(&_IQuickPool.CallOpts, index)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IQuickPool *IQuickPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IQuickPool *IQuickPoolSession) Token0() (common.Address, error) {
	return _IQuickPool.Contract.Token0(&_IQuickPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IQuickPool *IQuickPoolCallerSession) Token0() (common.Address, error) {
	return _IQuickPool.Contract.Token0(&_IQuickPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IQuickPool *IQuickPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IQuickPool *IQuickPoolSession) Token1() (common.Address, error) {
	return _IQuickPool.Contract.Token1(&_IQuickPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IQuickPool *IQuickPoolCallerSession) Token1() (common.Address, error) {
	return _IQuickPool.Contract.Token1(&_IQuickPool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_IQuickPool *IQuickPoolCaller) TotalFeeGrowth0Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "totalFeeGrowth0Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_IQuickPool *IQuickPoolSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _IQuickPool.Contract.TotalFeeGrowth0Token(&_IQuickPool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_IQuickPool *IQuickPoolCallerSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _IQuickPool.Contract.TotalFeeGrowth0Token(&_IQuickPool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_IQuickPool *IQuickPoolCaller) TotalFeeGrowth1Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IQuickPool.contract.Call(opts, &out, "totalFeeGrowth1Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_IQuickPool *IQuickPoolSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _IQuickPool.Contract.TotalFeeGrowth1Token(&_IQuickPool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_IQuickPool *IQuickPoolCallerSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _IQuickPool.Contract.TotalFeeGrowth1Token(&_IQuickPool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolTransactor) Burn(opts *bind.TransactOpts, bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "burn", bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Burn(&_IQuickPool.TransactOpts, bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolTransactorSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Burn(&_IQuickPool.TransactOpts, bottomTick, topTick, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "collect", recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Collect(&_IQuickPool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolTransactorSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Collect(&_IQuickPool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IQuickPool *IQuickPoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IQuickPool *IQuickPoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Flash(&_IQuickPool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IQuickPool *IQuickPoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Flash(&_IQuickPool.TransactOpts, recipient, amount0, amount1, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_IQuickPool *IQuickPoolTransactor) Initialize(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "initialize", price)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_IQuickPool *IQuickPoolSession) Initialize(price *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Initialize(&_IQuickPool.TransactOpts, price)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_IQuickPool *IQuickPoolTransactorSession) Initialize(price *big.Int) (*types.Transaction, error) {
	return _IQuickPool.Contract.Initialize(&_IQuickPool.TransactOpts, price)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_IQuickPool *IQuickPoolTransactor) Mint(opts *bind.TransactOpts, sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "mint", sender, recipient, bottomTick, topTick, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_IQuickPool *IQuickPoolSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Mint(&_IQuickPool.TransactOpts, sender, recipient, bottomTick, topTick, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_IQuickPool *IQuickPoolTransactorSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Mint(&_IQuickPool.TransactOpts, sender, recipient, bottomTick, topTick, amount, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "swap", recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolSession) Swap(recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Swap(&_IQuickPool.TransactOpts, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolTransactorSession) Swap(recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.Swap(&_IQuickPool.TransactOpts, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolTransactor) SwapSupportingFeeOnInputTokens(opts *bind.TransactOpts, sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.contract.Transact(opts, "swapSupportingFeeOnInputTokens", sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.SwapSupportingFeeOnInputTokens(&_IQuickPool.TransactOpts, sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_IQuickPool *IQuickPoolTransactorSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _IQuickPool.Contract.SwapSupportingFeeOnInputTokens(&_IQuickPool.TransactOpts, sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// IQuickPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IQuickPool contract.
type IQuickPoolBurnIterator struct {
	Event *IQuickPoolBurn // Event containing the contract specifics and raw log

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
func (it *IQuickPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolBurn)
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
		it.Event = new(IQuickPoolBurn)
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
func (it *IQuickPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolBurn represents a Burn event raised by the IQuickPool contract.
type IQuickPoolBurn struct {
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*IQuickPoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolBurnIterator{contract: _IQuickPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IQuickPoolBurn, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolBurn)
				if err := _IQuickPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) ParseBurn(log types.Log) (*IQuickPoolBurn, error) {
	event := new(IQuickPoolBurn)
	if err := _IQuickPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the IQuickPool contract.
type IQuickPoolCollectIterator struct {
	Event *IQuickPoolCollect // Event containing the contract specifics and raw log

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
func (it *IQuickPoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolCollect)
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
		it.Event = new(IQuickPoolCollect)
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
func (it *IQuickPoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolCollect represents a Collect event raised by the IQuickPool contract.
type IQuickPoolCollect struct {
	Owner      common.Address
	Recipient  common.Address
	BottomTick *big.Int
	TopTick    *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*IQuickPoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolCollectIterator{contract: _IQuickPool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *IQuickPoolCollect, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolCollect)
				if err := _IQuickPool.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_IQuickPool *IQuickPoolFilterer) ParseCollect(log types.Log) (*IQuickPoolCollect, error) {
	event := new(IQuickPoolCollect)
	if err := _IQuickPool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolCommunityFeeIterator is returned from FilterCommunityFee and is used to iterate over the raw logs and unpacked data for CommunityFee events raised by the IQuickPool contract.
type IQuickPoolCommunityFeeIterator struct {
	Event *IQuickPoolCommunityFee // Event containing the contract specifics and raw log

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
func (it *IQuickPoolCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolCommunityFee)
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
		it.Event = new(IQuickPoolCommunityFee)
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
func (it *IQuickPoolCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolCommunityFee represents a CommunityFee event raised by the IQuickPool contract.
type IQuickPoolCommunityFee struct {
	CommunityFee0New uint8
	CommunityFee1New uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommunityFee is a free log retrieval operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_IQuickPool *IQuickPoolFilterer) FilterCommunityFee(opts *bind.FilterOpts) (*IQuickPoolCommunityFeeIterator, error) {

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return &IQuickPoolCommunityFeeIterator{contract: _IQuickPool.contract, event: "CommunityFee", logs: logs, sub: sub}, nil
}

// WatchCommunityFee is a free log subscription operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_IQuickPool *IQuickPoolFilterer) WatchCommunityFee(opts *bind.WatchOpts, sink chan<- *IQuickPoolCommunityFee) (event.Subscription, error) {

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolCommunityFee)
				if err := _IQuickPool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
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

// ParseCommunityFee is a log parse operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_IQuickPool *IQuickPoolFilterer) ParseCommunityFee(log types.Log) (*IQuickPoolCommunityFee, error) {
	event := new(IQuickPoolCommunityFee)
	if err := _IQuickPool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the IQuickPool contract.
type IQuickPoolFeeIterator struct {
	Event *IQuickPoolFee // Event containing the contract specifics and raw log

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
func (it *IQuickPoolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolFee)
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
		it.Event = new(IQuickPoolFee)
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
func (it *IQuickPoolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolFee represents a Fee event raised by the IQuickPool contract.
type IQuickPoolFee struct {
	Fee uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_IQuickPool *IQuickPoolFilterer) FilterFee(opts *bind.FilterOpts) (*IQuickPoolFeeIterator, error) {

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &IQuickPoolFeeIterator{contract: _IQuickPool.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_IQuickPool *IQuickPoolFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *IQuickPoolFee) (event.Subscription, error) {

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolFee)
				if err := _IQuickPool.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_IQuickPool *IQuickPoolFilterer) ParseFee(log types.Log) (*IQuickPoolFee, error) {
	event := new(IQuickPoolFee)
	if err := _IQuickPool.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the IQuickPool contract.
type IQuickPoolFlashIterator struct {
	Event *IQuickPoolFlash // Event containing the contract specifics and raw log

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
func (it *IQuickPoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolFlash)
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
		it.Event = new(IQuickPoolFlash)
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
func (it *IQuickPoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolFlash represents a Flash event raised by the IQuickPool contract.
type IQuickPoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IQuickPool *IQuickPoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IQuickPoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolFlashIterator{contract: _IQuickPool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IQuickPool *IQuickPoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *IQuickPoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolFlash)
				if err := _IQuickPool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IQuickPool *IQuickPoolFilterer) ParseFlash(log types.Log) (*IQuickPoolFlash, error) {
	event := new(IQuickPoolFlash)
	if err := _IQuickPool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolIncentiveIterator is returned from FilterIncentive and is used to iterate over the raw logs and unpacked data for Incentive events raised by the IQuickPool contract.
type IQuickPoolIncentiveIterator struct {
	Event *IQuickPoolIncentive // Event containing the contract specifics and raw log

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
func (it *IQuickPoolIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolIncentive)
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
		it.Event = new(IQuickPoolIncentive)
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
func (it *IQuickPoolIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolIncentive represents a Incentive event raised by the IQuickPool contract.
type IQuickPoolIncentive struct {
	VirtualPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentive is a free log retrieval operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_IQuickPool *IQuickPoolFilterer) FilterIncentive(opts *bind.FilterOpts, virtualPoolAddress []common.Address) (*IQuickPoolIncentiveIterator, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolIncentiveIterator{contract: _IQuickPool.contract, event: "Incentive", logs: logs, sub: sub}, nil
}

// WatchIncentive is a free log subscription operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_IQuickPool *IQuickPoolFilterer) WatchIncentive(opts *bind.WatchOpts, sink chan<- *IQuickPoolIncentive, virtualPoolAddress []common.Address) (event.Subscription, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolIncentive)
				if err := _IQuickPool.contract.UnpackLog(event, "Incentive", log); err != nil {
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

// ParseIncentive is a log parse operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_IQuickPool *IQuickPoolFilterer) ParseIncentive(log types.Log) (*IQuickPoolIncentive, error) {
	event := new(IQuickPoolIncentive)
	if err := _IQuickPool.contract.UnpackLog(event, "Incentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the IQuickPool contract.
type IQuickPoolInitializeIterator struct {
	Event *IQuickPoolInitialize // Event containing the contract specifics and raw log

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
func (it *IQuickPoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolInitialize)
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
		it.Event = new(IQuickPoolInitialize)
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
func (it *IQuickPoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolInitialize represents a Initialize event raised by the IQuickPool contract.
type IQuickPoolInitialize struct {
	Price *big.Int
	Tick  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*IQuickPoolInitializeIterator, error) {

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &IQuickPoolInitializeIterator{contract: _IQuickPool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *IQuickPoolInitialize) (event.Subscription, error) {

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolInitialize)
				if err := _IQuickPool.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) ParseInitialize(log types.Log) (*IQuickPoolInitialize, error) {
	event := new(IQuickPoolInitialize)
	if err := _IQuickPool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolLiquidityCooldownIterator is returned from FilterLiquidityCooldown and is used to iterate over the raw logs and unpacked data for LiquidityCooldown events raised by the IQuickPool contract.
type IQuickPoolLiquidityCooldownIterator struct {
	Event *IQuickPoolLiquidityCooldown // Event containing the contract specifics and raw log

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
func (it *IQuickPoolLiquidityCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolLiquidityCooldown)
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
		it.Event = new(IQuickPoolLiquidityCooldown)
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
func (it *IQuickPoolLiquidityCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolLiquidityCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolLiquidityCooldown represents a LiquidityCooldown event raised by the IQuickPool contract.
type IQuickPoolLiquidityCooldown struct {
	LiquidityCooldown uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCooldown is a free log retrieval operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_IQuickPool *IQuickPoolFilterer) FilterLiquidityCooldown(opts *bind.FilterOpts) (*IQuickPoolLiquidityCooldownIterator, error) {

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return &IQuickPoolLiquidityCooldownIterator{contract: _IQuickPool.contract, event: "LiquidityCooldown", logs: logs, sub: sub}, nil
}

// WatchLiquidityCooldown is a free log subscription operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_IQuickPool *IQuickPoolFilterer) WatchLiquidityCooldown(opts *bind.WatchOpts, sink chan<- *IQuickPoolLiquidityCooldown) (event.Subscription, error) {

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolLiquidityCooldown)
				if err := _IQuickPool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
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

// ParseLiquidityCooldown is a log parse operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_IQuickPool *IQuickPoolFilterer) ParseLiquidityCooldown(log types.Log) (*IQuickPoolLiquidityCooldown, error) {
	event := new(IQuickPoolLiquidityCooldown)
	if err := _IQuickPool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IQuickPool contract.
type IQuickPoolMintIterator struct {
	Event *IQuickPoolMint // Event containing the contract specifics and raw log

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
func (it *IQuickPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolMint)
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
		it.Event = new(IQuickPoolMint)
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
func (it *IQuickPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolMint represents a Mint event raised by the IQuickPool contract.
type IQuickPoolMint struct {
	Sender          common.Address
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*IQuickPoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolMintIterator{contract: _IQuickPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IQuickPoolMint, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolMint)
				if err := _IQuickPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_IQuickPool *IQuickPoolFilterer) ParseMint(log types.Log) (*IQuickPoolMint, error) {
	event := new(IQuickPoolMint)
	if err := _IQuickPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IQuickPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IQuickPool contract.
type IQuickPoolSwapIterator struct {
	Event *IQuickPoolSwap // Event containing the contract specifics and raw log

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
func (it *IQuickPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IQuickPoolSwap)
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
		it.Event = new(IQuickPoolSwap)
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
func (it *IQuickPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IQuickPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IQuickPoolSwap represents a Swap event raised by the IQuickPool contract.
type IQuickPoolSwap struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Price     *big.Int
	Liquidity *big.Int
	Tick      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IQuickPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IQuickPool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IQuickPoolSwapIterator{contract: _IQuickPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IQuickPoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IQuickPool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IQuickPoolSwap)
				if err := _IQuickPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_IQuickPool *IQuickPoolFilterer) ParseSwap(log types.Log) (*IQuickPoolSwap, error) {
	event := new(IQuickPoolSwap)
	if err := _IQuickPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
