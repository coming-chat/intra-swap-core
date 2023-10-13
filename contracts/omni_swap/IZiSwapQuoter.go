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

// IZiSwapQuoterSwapAmountParams is an auto generated low-level Go binding around an user-defined struct.
type IZiSwapQuoterSwapAmountParams struct {
	Path        []byte
	Recipient   common.Address
	Amount      *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// IZiSwapQuoterMetaData contains all meta data concerning the IZiSwapQuoter contract.
var IZiSwapQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIZiSwapQuoter.SwapAmountParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IZiSwapQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use IZiSwapQuoterMetaData.ABI instead.
var IZiSwapQuoterABI = IZiSwapQuoterMetaData.ABI

// IZiSwapQuoter is an auto generated Go binding around an Ethereum contract.
type IZiSwapQuoter struct {
	IZiSwapQuoterCaller     // Read-only binding to the contract
	IZiSwapQuoterTransactor // Write-only binding to the contract
	IZiSwapQuoterFilterer   // Log filterer for contract events
}

// IZiSwapQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZiSwapQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZiSwapQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZiSwapQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZiSwapQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZiSwapQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZiSwapQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZiSwapQuoterSession struct {
	Contract     *IZiSwapQuoter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZiSwapQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZiSwapQuoterCallerSession struct {
	Contract *IZiSwapQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IZiSwapQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZiSwapQuoterTransactorSession struct {
	Contract     *IZiSwapQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IZiSwapQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZiSwapQuoterRaw struct {
	Contract *IZiSwapQuoter // Generic contract binding to access the raw methods on
}

// IZiSwapQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZiSwapQuoterCallerRaw struct {
	Contract *IZiSwapQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// IZiSwapQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZiSwapQuoterTransactorRaw struct {
	Contract *IZiSwapQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZiSwapQuoter creates a new instance of IZiSwapQuoter, bound to a specific deployed contract.
func NewIZiSwapQuoter(address common.Address, backend bind.ContractBackend) (*IZiSwapQuoter, error) {
	contract, err := bindIZiSwapQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZiSwapQuoter{IZiSwapQuoterCaller: IZiSwapQuoterCaller{contract: contract}, IZiSwapQuoterTransactor: IZiSwapQuoterTransactor{contract: contract}, IZiSwapQuoterFilterer: IZiSwapQuoterFilterer{contract: contract}}, nil
}

// NewIZiSwapQuoterCaller creates a new read-only instance of IZiSwapQuoter, bound to a specific deployed contract.
func NewIZiSwapQuoterCaller(address common.Address, caller bind.ContractCaller) (*IZiSwapQuoterCaller, error) {
	contract, err := bindIZiSwapQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZiSwapQuoterCaller{contract: contract}, nil
}

// NewIZiSwapQuoterTransactor creates a new write-only instance of IZiSwapQuoter, bound to a specific deployed contract.
func NewIZiSwapQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*IZiSwapQuoterTransactor, error) {
	contract, err := bindIZiSwapQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZiSwapQuoterTransactor{contract: contract}, nil
}

// NewIZiSwapQuoterFilterer creates a new log filterer instance of IZiSwapQuoter, bound to a specific deployed contract.
func NewIZiSwapQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*IZiSwapQuoterFilterer, error) {
	contract, err := bindIZiSwapQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZiSwapQuoterFilterer{contract: contract}, nil
}

// bindIZiSwapQuoter binds a generic wrapper to an already deployed contract.
func bindIZiSwapQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZiSwapQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZiSwapQuoter *IZiSwapQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZiSwapQuoter.Contract.IZiSwapQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZiSwapQuoter *IZiSwapQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.IZiSwapQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZiSwapQuoter *IZiSwapQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.IZiSwapQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZiSwapQuoter *IZiSwapQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZiSwapQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZiSwapQuoter *IZiSwapQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZiSwapQuoter *IZiSwapQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.contract.Transact(opts, method, params...)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IZiSwapQuoter *IZiSwapQuoterTransactor) SwapAmount(opts *bind.TransactOpts, params IZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IZiSwapQuoter.contract.Transact(opts, "swapAmount", params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IZiSwapQuoter *IZiSwapQuoterSession) SwapAmount(params IZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.SwapAmount(&_IZiSwapQuoter.TransactOpts, params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IZiSwapQuoter *IZiSwapQuoterTransactorSession) SwapAmount(params IZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IZiSwapQuoter.Contract.SwapAmount(&_IZiSwapQuoter.TransactOpts, params)
}
