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

// IGMXV1ReaderMetaData contains all meta data concerning the IGMXV1Reader contract.
var IGMXV1ReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGMXV1Vault\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IGMXV1ReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use IGMXV1ReaderMetaData.ABI instead.
var IGMXV1ReaderABI = IGMXV1ReaderMetaData.ABI

// IGMXV1Reader is an auto generated Go binding around an Ethereum contract.
type IGMXV1Reader struct {
	IGMXV1ReaderCaller     // Read-only binding to the contract
	IGMXV1ReaderTransactor // Write-only binding to the contract
	IGMXV1ReaderFilterer   // Log filterer for contract events
}

// IGMXV1ReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGMXV1ReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1ReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGMXV1ReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1ReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGMXV1ReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1ReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGMXV1ReaderSession struct {
	Contract     *IGMXV1Reader     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGMXV1ReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGMXV1ReaderCallerSession struct {
	Contract *IGMXV1ReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGMXV1ReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGMXV1ReaderTransactorSession struct {
	Contract     *IGMXV1ReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGMXV1ReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGMXV1ReaderRaw struct {
	Contract *IGMXV1Reader // Generic contract binding to access the raw methods on
}

// IGMXV1ReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGMXV1ReaderCallerRaw struct {
	Contract *IGMXV1ReaderCaller // Generic read-only contract binding to access the raw methods on
}

// IGMXV1ReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGMXV1ReaderTransactorRaw struct {
	Contract *IGMXV1ReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGMXV1Reader creates a new instance of IGMXV1Reader, bound to a specific deployed contract.
func NewIGMXV1Reader(address common.Address, backend bind.ContractBackend) (*IGMXV1Reader, error) {
	contract, err := bindIGMXV1Reader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGMXV1Reader{IGMXV1ReaderCaller: IGMXV1ReaderCaller{contract: contract}, IGMXV1ReaderTransactor: IGMXV1ReaderTransactor{contract: contract}, IGMXV1ReaderFilterer: IGMXV1ReaderFilterer{contract: contract}}, nil
}

// NewIGMXV1ReaderCaller creates a new read-only instance of IGMXV1Reader, bound to a specific deployed contract.
func NewIGMXV1ReaderCaller(address common.Address, caller bind.ContractCaller) (*IGMXV1ReaderCaller, error) {
	contract, err := bindIGMXV1Reader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1ReaderCaller{contract: contract}, nil
}

// NewIGMXV1ReaderTransactor creates a new write-only instance of IGMXV1Reader, bound to a specific deployed contract.
func NewIGMXV1ReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*IGMXV1ReaderTransactor, error) {
	contract, err := bindIGMXV1Reader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1ReaderTransactor{contract: contract}, nil
}

// NewIGMXV1ReaderFilterer creates a new log filterer instance of IGMXV1Reader, bound to a specific deployed contract.
func NewIGMXV1ReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*IGMXV1ReaderFilterer, error) {
	contract, err := bindIGMXV1Reader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGMXV1ReaderFilterer{contract: contract}, nil
}

// bindIGMXV1Reader binds a generic wrapper to an already deployed contract.
func bindIGMXV1Reader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGMXV1ReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Reader *IGMXV1ReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Reader.Contract.IGMXV1ReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Reader *IGMXV1ReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Reader.Contract.IGMXV1ReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Reader *IGMXV1ReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Reader.Contract.IGMXV1ReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Reader *IGMXV1ReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Reader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Reader *IGMXV1ReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Reader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Reader *IGMXV1ReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Reader.Contract.contract.Transact(opts, method, params...)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_IGMXV1Reader *IGMXV1ReaderCaller) GetAmountOut(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGMXV1Reader.contract.Call(opts, &out, "getAmountOut", _vault, _tokenIn, _tokenOut, _amountIn)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_IGMXV1Reader *IGMXV1ReaderSession) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _IGMXV1Reader.Contract.GetAmountOut(&_IGMXV1Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_IGMXV1Reader *IGMXV1ReaderCallerSession) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _IGMXV1Reader.Contract.GetAmountOut(&_IGMXV1Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}
