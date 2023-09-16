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

// IStargateCachedSwap is an auto generated low-level Go binding around an user-defined struct.
type IStargateCachedSwap struct {
	Token    common.Address
	AmountLD *big.Int
	To       common.Address
	Payload  []byte
}

// IStargatelzTxObj is an auto generated low-level Go binding around an user-defined struct.
type IStargatelzTxObj struct {
	DstGasForCall   *big.Int
	DstNativeAmount *big.Int
	DstNativeAddr   []byte
}

// IStargateMetaData contains all meta data concerning the IStargate contract.
var IStargateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"CachedSwapSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmountSD\",\"type\":\"uint256\"}],\"name\":\"RedeemLocalCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"bridgeFunctionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Revert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmountSD\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"}],\"name\":\"RevertRedeemLocal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"cachedSwapLookup\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structIStargate.CachedSwap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"clearCachedSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcPoolId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"instantRedeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_functionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_transferAndCallPayload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargate.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"quoteLayerZeroFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargate.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"redeemLocal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargate.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"redeemRemote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendCredits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargate.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IStargateABI is the input ABI used to generate the binding from.
// Deprecated: Use IStargateMetaData.ABI instead.
var IStargateABI = IStargateMetaData.ABI

// IStargate is an auto generated Go binding around an Ethereum contract.
type IStargate struct {
	IStargateCaller     // Read-only binding to the contract
	IStargateTransactor // Write-only binding to the contract
	IStargateFilterer   // Log filterer for contract events
}

// IStargateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStargateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStargateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStargateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStargateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStargateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStargateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStargateSession struct {
	Contract     *IStargate        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStargateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStargateCallerSession struct {
	Contract *IStargateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IStargateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStargateTransactorSession struct {
	Contract     *IStargateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStargateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStargateRaw struct {
	Contract *IStargate // Generic contract binding to access the raw methods on
}

// IStargateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStargateCallerRaw struct {
	Contract *IStargateCaller // Generic read-only contract binding to access the raw methods on
}

// IStargateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStargateTransactorRaw struct {
	Contract *IStargateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStargate creates a new instance of IStargate, bound to a specific deployed contract.
func NewIStargate(address common.Address, backend bind.ContractBackend) (*IStargate, error) {
	contract, err := bindIStargate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStargate{IStargateCaller: IStargateCaller{contract: contract}, IStargateTransactor: IStargateTransactor{contract: contract}, IStargateFilterer: IStargateFilterer{contract: contract}}, nil
}

// NewIStargateCaller creates a new read-only instance of IStargate, bound to a specific deployed contract.
func NewIStargateCaller(address common.Address, caller bind.ContractCaller) (*IStargateCaller, error) {
	contract, err := bindIStargate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStargateCaller{contract: contract}, nil
}

// NewIStargateTransactor creates a new write-only instance of IStargate, bound to a specific deployed contract.
func NewIStargateTransactor(address common.Address, transactor bind.ContractTransactor) (*IStargateTransactor, error) {
	contract, err := bindIStargate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStargateTransactor{contract: contract}, nil
}

// NewIStargateFilterer creates a new log filterer instance of IStargate, bound to a specific deployed contract.
func NewIStargateFilterer(address common.Address, filterer bind.ContractFilterer) (*IStargateFilterer, error) {
	contract, err := bindIStargate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStargateFilterer{contract: contract}, nil
}

// bindIStargate binds a generic wrapper to an already deployed contract.
func bindIStargate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStargateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStargate *IStargateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStargate.Contract.IStargateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStargate *IStargateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStargate.Contract.IStargateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStargate *IStargateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStargate.Contract.IStargateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStargate *IStargateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStargate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStargate *IStargateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStargate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStargate *IStargateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStargate.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IStargate *IStargateCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStargate.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IStargate *IStargateSession) Bridge() (common.Address, error) {
	return _IStargate.Contract.Bridge(&_IStargate.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IStargate *IStargateCallerSession) Bridge() (common.Address, error) {
	return _IStargate.Contract.Bridge(&_IStargate.CallOpts)
}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) view returns((address,uint256,address,bytes))
func (_IStargate *IStargateCaller) CachedSwapLookup(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (IStargateCachedSwap, error) {
	var out []interface{}
	err := _IStargate.contract.Call(opts, &out, "cachedSwapLookup", _srcChainId, _srcAddress, _nonce)

	if err != nil {
		return *new(IStargateCachedSwap), err
	}

	out0 := *abi.ConvertType(out[0], new(IStargateCachedSwap)).(*IStargateCachedSwap)

	return out0, err

}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) view returns((address,uint256,address,bytes))
func (_IStargate *IStargateSession) CachedSwapLookup(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (IStargateCachedSwap, error) {
	return _IStargate.Contract.CachedSwapLookup(&_IStargate.CallOpts, _srcChainId, _srcAddress, _nonce)
}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) view returns((address,uint256,address,bytes))
func (_IStargate *IStargateCallerSession) CachedSwapLookup(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (IStargateCachedSwap, error) {
	return _IStargate.Contract.CachedSwapLookup(&_IStargate.CallOpts, _srcChainId, _srcAddress, _nonce)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IStargate *IStargateCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStargate.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IStargate *IStargateSession) Factory() (common.Address, error) {
	return _IStargate.Contract.Factory(&_IStargate.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IStargate *IStargateCallerSession) Factory() (common.Address, error) {
	return _IStargate.Contract.Factory(&_IStargate.CallOpts)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_IStargate *IStargateCaller) QuoteLayerZeroFee(opts *bind.CallOpts, _dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargatelzTxObj) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IStargate.contract.Call(opts, &out, "quoteLayerZeroFee", _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_IStargate *IStargateSession) QuoteLayerZeroFee(_dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargatelzTxObj) (*big.Int, *big.Int, error) {
	return _IStargate.Contract.QuoteLayerZeroFee(&_IStargate.CallOpts, _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_IStargate *IStargateCallerSession) QuoteLayerZeroFee(_dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargatelzTxObj) (*big.Int, *big.Int, error) {
	return _IStargate.Contract.QuoteLayerZeroFee(&_IStargate.CallOpts, _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_IStargate *IStargateTransactor) AddLiquidity(opts *bind.TransactOpts, _poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "addLiquidity", _poolId, _amountLD, _to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_IStargate *IStargateSession) AddLiquidity(_poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.AddLiquidity(&_IStargate.TransactOpts, _poolId, _amountLD, _to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_IStargate *IStargateTransactorSession) AddLiquidity(_poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.AddLiquidity(&_IStargate.TransactOpts, _poolId, _amountLD, _to)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_IStargate *IStargateTransactor) ClearCachedSwap(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "clearCachedSwap", _srcChainId, _srcAddress, _nonce)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_IStargate *IStargateSession) ClearCachedSwap(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _IStargate.Contract.ClearCachedSwap(&_IStargate.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_IStargate *IStargateTransactorSession) ClearCachedSwap(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _IStargate.Contract.ClearCachedSwap(&_IStargate.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256)
func (_IStargate *IStargateTransactor) InstantRedeemLocal(opts *bind.TransactOpts, _srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "instantRedeemLocal", _srcPoolId, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256)
func (_IStargate *IStargateSession) InstantRedeemLocal(_srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.InstantRedeemLocal(&_IStargate.TransactOpts, _srcPoolId, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256)
func (_IStargate *IStargateTransactorSession) InstantRedeemLocal(_srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.InstantRedeemLocal(&_IStargate.TransactOpts, _srcPoolId, _amountLP, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateTransactor) RedeemLocal(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "redeemLocal", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateSession) RedeemLocal(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.Contract.RedeemLocal(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateTransactorSession) RedeemLocal(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.Contract.RedeemLocal(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateTransactor) RedeemRemote(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "redeemRemote", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateSession) RedeemRemote(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.Contract.RedeemRemote(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_IStargate *IStargateTransactorSession) RedeemRemote(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargatelzTxObj) (*types.Transaction, error) {
	return _IStargate.Contract.RedeemRemote(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_IStargate *IStargateTransactor) SendCredits(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "sendCredits", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_IStargate *IStargateSession) SendCredits(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.SendCredits(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_IStargate *IStargateTransactorSession) SendCredits(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _IStargate.Contract.SendCredits(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_IStargate *IStargateTransactor) Swap(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargatelzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _IStargate.contract.Transact(opts, "swap", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_IStargate *IStargateSession) Swap(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargatelzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _IStargate.Contract.Swap(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_IStargate *IStargateTransactorSession) Swap(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargatelzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _IStargate.Contract.Swap(&_IStargate.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// IStargateCachedSwapSavedIterator is returned from FilterCachedSwapSaved and is used to iterate over the raw logs and unpacked data for CachedSwapSaved events raised by the IStargate contract.
type IStargateCachedSwapSavedIterator struct {
	Event *IStargateCachedSwapSaved // Event containing the contract specifics and raw log

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
func (it *IStargateCachedSwapSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStargateCachedSwapSaved)
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
		it.Event = new(IStargateCachedSwapSaved)
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
func (it *IStargateCachedSwapSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStargateCachedSwapSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStargateCachedSwapSaved represents a CachedSwapSaved event raised by the IStargate contract.
type IStargateCachedSwapSaved struct {
	ChainId    uint16
	SrcAddress []byte
	Nonce      *big.Int
	Token      common.Address
	AmountLD   *big.Int
	To         common.Address
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCachedSwapSaved is a free log retrieval operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_IStargate *IStargateFilterer) FilterCachedSwapSaved(opts *bind.FilterOpts) (*IStargateCachedSwapSavedIterator, error) {

	logs, sub, err := _IStargate.contract.FilterLogs(opts, "CachedSwapSaved")
	if err != nil {
		return nil, err
	}
	return &IStargateCachedSwapSavedIterator{contract: _IStargate.contract, event: "CachedSwapSaved", logs: logs, sub: sub}, nil
}

// WatchCachedSwapSaved is a free log subscription operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_IStargate *IStargateFilterer) WatchCachedSwapSaved(opts *bind.WatchOpts, sink chan<- *IStargateCachedSwapSaved) (event.Subscription, error) {

	logs, sub, err := _IStargate.contract.WatchLogs(opts, "CachedSwapSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStargateCachedSwapSaved)
				if err := _IStargate.contract.UnpackLog(event, "CachedSwapSaved", log); err != nil {
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

// ParseCachedSwapSaved is a log parse operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_IStargate *IStargateFilterer) ParseCachedSwapSaved(log types.Log) (*IStargateCachedSwapSaved, error) {
	event := new(IStargateCachedSwapSaved)
	if err := _IStargate.contract.UnpackLog(event, "CachedSwapSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStargateRedeemLocalCallbackIterator is returned from FilterRedeemLocalCallback and is used to iterate over the raw logs and unpacked data for RedeemLocalCallback events raised by the IStargate contract.
type IStargateRedeemLocalCallbackIterator struct {
	Event *IStargateRedeemLocalCallback // Event containing the contract specifics and raw log

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
func (it *IStargateRedeemLocalCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStargateRedeemLocalCallback)
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
		it.Event = new(IStargateRedeemLocalCallback)
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
func (it *IStargateRedeemLocalCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStargateRedeemLocalCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStargateRedeemLocalCallback represents a RedeemLocalCallback event raised by the IStargate contract.
type IStargateRedeemLocalCallback struct {
	SrcChainId   uint16
	SrcAddress   common.Hash
	Nonce        *big.Int
	SrcPoolId    *big.Int
	DstPoolId    *big.Int
	To           common.Address
	AmountSD     *big.Int
	MintAmountSD *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocalCallback is a free log retrieval operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_IStargate *IStargateFilterer) FilterRedeemLocalCallback(opts *bind.FilterOpts, srcAddress [][]byte, nonce []*big.Int) (*IStargateRedeemLocalCallbackIterator, error) {

	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IStargate.contract.FilterLogs(opts, "RedeemLocalCallback", srcAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &IStargateRedeemLocalCallbackIterator{contract: _IStargate.contract, event: "RedeemLocalCallback", logs: logs, sub: sub}, nil
}

// WatchRedeemLocalCallback is a free log subscription operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_IStargate *IStargateFilterer) WatchRedeemLocalCallback(opts *bind.WatchOpts, sink chan<- *IStargateRedeemLocalCallback, srcAddress [][]byte, nonce []*big.Int) (event.Subscription, error) {

	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IStargate.contract.WatchLogs(opts, "RedeemLocalCallback", srcAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStargateRedeemLocalCallback)
				if err := _IStargate.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
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

// ParseRedeemLocalCallback is a log parse operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_IStargate *IStargateFilterer) ParseRedeemLocalCallback(log types.Log) (*IStargateRedeemLocalCallback, error) {
	event := new(IStargateRedeemLocalCallback)
	if err := _IStargate.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStargateRevertIterator is returned from FilterRevert and is used to iterate over the raw logs and unpacked data for Revert events raised by the IStargate contract.
type IStargateRevertIterator struct {
	Event *IStargateRevert // Event containing the contract specifics and raw log

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
func (it *IStargateRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStargateRevert)
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
		it.Event = new(IStargateRevert)
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
func (it *IStargateRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStargateRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStargateRevert represents a Revert event raised by the IStargate contract.
type IStargateRevert struct {
	BridgeFunctionType uint8
	ChainId            uint16
	SrcAddress         []byte
	Nonce              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRevert is a free log retrieval operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_IStargate *IStargateFilterer) FilterRevert(opts *bind.FilterOpts) (*IStargateRevertIterator, error) {

	logs, sub, err := _IStargate.contract.FilterLogs(opts, "Revert")
	if err != nil {
		return nil, err
	}
	return &IStargateRevertIterator{contract: _IStargate.contract, event: "Revert", logs: logs, sub: sub}, nil
}

// WatchRevert is a free log subscription operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_IStargate *IStargateFilterer) WatchRevert(opts *bind.WatchOpts, sink chan<- *IStargateRevert) (event.Subscription, error) {

	logs, sub, err := _IStargate.contract.WatchLogs(opts, "Revert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStargateRevert)
				if err := _IStargate.contract.UnpackLog(event, "Revert", log); err != nil {
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

// ParseRevert is a log parse operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_IStargate *IStargateFilterer) ParseRevert(log types.Log) (*IStargateRevert, error) {
	event := new(IStargateRevert)
	if err := _IStargate.contract.UnpackLog(event, "Revert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStargateRevertRedeemLocalIterator is returned from FilterRevertRedeemLocal and is used to iterate over the raw logs and unpacked data for RevertRedeemLocal events raised by the IStargate contract.
type IStargateRevertRedeemLocalIterator struct {
	Event *IStargateRevertRedeemLocal // Event containing the contract specifics and raw log

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
func (it *IStargateRevertRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStargateRevertRedeemLocal)
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
		it.Event = new(IStargateRevertRedeemLocal)
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
func (it *IStargateRevertRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStargateRevertRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStargateRevertRedeemLocal represents a RevertRedeemLocal event raised by the IStargate contract.
type IStargateRevertRedeemLocal struct {
	SrcChainId     uint16
	SrcPoolId      *big.Int
	DstPoolId      *big.Int
	To             []byte
	RedeemAmountSD *big.Int
	MintAmountSD   *big.Int
	Nonce          *big.Int
	SrcAddress     common.Hash
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevertRedeemLocal is a free log retrieval operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_IStargate *IStargateFilterer) FilterRevertRedeemLocal(opts *bind.FilterOpts, nonce []*big.Int, srcAddress [][]byte) (*IStargateRevertRedeemLocalIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}

	logs, sub, err := _IStargate.contract.FilterLogs(opts, "RevertRedeemLocal", nonceRule, srcAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStargateRevertRedeemLocalIterator{contract: _IStargate.contract, event: "RevertRedeemLocal", logs: logs, sub: sub}, nil
}

// WatchRevertRedeemLocal is a free log subscription operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_IStargate *IStargateFilterer) WatchRevertRedeemLocal(opts *bind.WatchOpts, sink chan<- *IStargateRevertRedeemLocal, nonce []*big.Int, srcAddress [][]byte) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}

	logs, sub, err := _IStargate.contract.WatchLogs(opts, "RevertRedeemLocal", nonceRule, srcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStargateRevertRedeemLocal)
				if err := _IStargate.contract.UnpackLog(event, "RevertRedeemLocal", log); err != nil {
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

// ParseRevertRedeemLocal is a log parse operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_IStargate *IStargateFilterer) ParseRevertRedeemLocal(log types.Log) (*IStargateRevertRedeemLocal, error) {
	event := new(IStargateRevertRedeemLocal)
	if err := _IStargate.contract.UnpackLog(event, "RevertRedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
