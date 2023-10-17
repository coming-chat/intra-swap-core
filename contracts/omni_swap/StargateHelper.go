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

// StargateHelperMetaData contains all meta data concerning the StargateHelper contract.
var StargateHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stargatePayload\",\"type\":\"bytes\"}],\"name\":\"findStargatePayload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inputData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"soDiamond\",\"type\":\"address\"}],\"name\":\"tryFindStargatePayload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StargateHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use StargateHelperMetaData.ABI instead.
var StargateHelperABI = StargateHelperMetaData.ABI

// StargateHelper is an auto generated Go binding around an Ethereum contract.
type StargateHelper struct {
	StargateHelperCaller     // Read-only binding to the contract
	StargateHelperTransactor // Write-only binding to the contract
	StargateHelperFilterer   // Log filterer for contract events
}

// StargateHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type StargateHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StargateHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StargateHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StargateHelperSession struct {
	Contract     *StargateHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StargateHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StargateHelperCallerSession struct {
	Contract *StargateHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StargateHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StargateHelperTransactorSession struct {
	Contract     *StargateHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StargateHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type StargateHelperRaw struct {
	Contract *StargateHelper // Generic contract binding to access the raw methods on
}

// StargateHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StargateHelperCallerRaw struct {
	Contract *StargateHelperCaller // Generic read-only contract binding to access the raw methods on
}

// StargateHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StargateHelperTransactorRaw struct {
	Contract *StargateHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStargateHelper creates a new instance of StargateHelper, bound to a specific deployed contract.
func NewStargateHelper(address common.Address, backend bind.ContractBackend) (*StargateHelper, error) {
	contract, err := bindStargateHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StargateHelper{StargateHelperCaller: StargateHelperCaller{contract: contract}, StargateHelperTransactor: StargateHelperTransactor{contract: contract}, StargateHelperFilterer: StargateHelperFilterer{contract: contract}}, nil
}

// NewStargateHelperCaller creates a new read-only instance of StargateHelper, bound to a specific deployed contract.
func NewStargateHelperCaller(address common.Address, caller bind.ContractCaller) (*StargateHelperCaller, error) {
	contract, err := bindStargateHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StargateHelperCaller{contract: contract}, nil
}

// NewStargateHelperTransactor creates a new write-only instance of StargateHelper, bound to a specific deployed contract.
func NewStargateHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*StargateHelperTransactor, error) {
	contract, err := bindStargateHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StargateHelperTransactor{contract: contract}, nil
}

// NewStargateHelperFilterer creates a new log filterer instance of StargateHelper, bound to a specific deployed contract.
func NewStargateHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*StargateHelperFilterer, error) {
	contract, err := bindStargateHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StargateHelperFilterer{contract: contract}, nil
}

// bindStargateHelper binds a generic wrapper to an already deployed contract.
func bindStargateHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StargateHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargateHelper *StargateHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargateHelper.Contract.StargateHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargateHelper *StargateHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateHelper.Contract.StargateHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargateHelper *StargateHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargateHelper.Contract.StargateHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargateHelper *StargateHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargateHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargateHelper *StargateHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargateHelper *StargateHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargateHelper.Contract.contract.Transact(opts, method, params...)
}

// FindStargatePayload is a free data retrieval call binding the contract method 0x02521f32.
//
// Solidity: function findStargatePayload(bytes stargatePayload) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperCaller) FindStargatePayload(opts *bind.CallOpts, stargatePayload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _StargateHelper.contract.Call(opts, &out, "findStargatePayload", stargatePayload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// FindStargatePayload is a free data retrieval call binding the contract method 0x02521f32.
//
// Solidity: function findStargatePayload(bytes stargatePayload) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperSession) FindStargatePayload(stargatePayload []byte) (*big.Int, []byte, error) {
	return _StargateHelper.Contract.FindStargatePayload(&_StargateHelper.CallOpts, stargatePayload)
}

// FindStargatePayload is a free data retrieval call binding the contract method 0x02521f32.
//
// Solidity: function findStargatePayload(bytes stargatePayload) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperCallerSession) FindStargatePayload(stargatePayload []byte) (*big.Int, []byte, error) {
	return _StargateHelper.Contract.FindStargatePayload(&_StargateHelper.CallOpts, stargatePayload)
}

// TryFindStargatePayload is a free data retrieval call binding the contract method 0xa9121afa.
//
// Solidity: function tryFindStargatePayload(bytes inputData, address soDiamond) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperCaller) TryFindStargatePayload(opts *bind.CallOpts, inputData []byte, soDiamond common.Address) (*big.Int, []byte, error) {
	var out []interface{}
	err := _StargateHelper.contract.Call(opts, &out, "tryFindStargatePayload", inputData, soDiamond)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// TryFindStargatePayload is a free data retrieval call binding the contract method 0xa9121afa.
//
// Solidity: function tryFindStargatePayload(bytes inputData, address soDiamond) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperSession) TryFindStargatePayload(inputData []byte, soDiamond common.Address) (*big.Int, []byte, error) {
	return _StargateHelper.Contract.TryFindStargatePayload(&_StargateHelper.CallOpts, inputData, soDiamond)
}

// TryFindStargatePayload is a free data retrieval call binding the contract method 0xa9121afa.
//
// Solidity: function tryFindStargatePayload(bytes inputData, address soDiamond) view returns(uint256, bytes)
func (_StargateHelper *StargateHelperCallerSession) TryFindStargatePayload(inputData []byte, soDiamond common.Address) (*big.Int, []byte, error) {
	return _StargateHelper.Contract.TryFindStargatePayload(&_StargateHelper.CallOpts, inputData, soDiamond)
}
