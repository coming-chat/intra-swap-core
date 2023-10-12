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

// IiZiSwapQuoterSwapAmountParams is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapQuoterSwapAmountParams struct {
	Path        []byte
	Recipient   common.Address
	Amount      *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// IiZiSwapQuoterMetaData contains all meta data concerning the IiZiSwapQuoter contract.
var IiZiSwapQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwapQuoter.SwapAmountParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IiZiSwapQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use IiZiSwapQuoterMetaData.ABI instead.
var IiZiSwapQuoterABI = IiZiSwapQuoterMetaData.ABI

// IiZiSwapQuoter is an auto generated Go binding around an Ethereum contract.
type IiZiSwapQuoter struct {
	IiZiSwapQuoterCaller     // Read-only binding to the contract
	IiZiSwapQuoterTransactor // Write-only binding to the contract
	IiZiSwapQuoterFilterer   // Log filterer for contract events
}

// IiZiSwapQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IiZiSwapQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IiZiSwapQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IiZiSwapQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IiZiSwapQuoterSession struct {
	Contract     *IiZiSwapQuoter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IiZiSwapQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IiZiSwapQuoterCallerSession struct {
	Contract *IiZiSwapQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IiZiSwapQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IiZiSwapQuoterTransactorSession struct {
	Contract     *IiZiSwapQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IiZiSwapQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IiZiSwapQuoterRaw struct {
	Contract *IiZiSwapQuoter // Generic contract binding to access the raw methods on
}

// IiZiSwapQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IiZiSwapQuoterCallerRaw struct {
	Contract *IiZiSwapQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// IiZiSwapQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IiZiSwapQuoterTransactorRaw struct {
	Contract *IiZiSwapQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIiZiSwapQuoter creates a new instance of IiZiSwapQuoter, bound to a specific deployed contract.
func NewIiZiSwapQuoter(address common.Address, backend bind.ContractBackend) (*IiZiSwapQuoter, error) {
	contract, err := bindIiZiSwapQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapQuoter{IiZiSwapQuoterCaller: IiZiSwapQuoterCaller{contract: contract}, IiZiSwapQuoterTransactor: IiZiSwapQuoterTransactor{contract: contract}, IiZiSwapQuoterFilterer: IiZiSwapQuoterFilterer{contract: contract}}, nil
}

// NewIiZiSwapQuoterCaller creates a new read-only instance of IiZiSwapQuoter, bound to a specific deployed contract.
func NewIiZiSwapQuoterCaller(address common.Address, caller bind.ContractCaller) (*IiZiSwapQuoterCaller, error) {
	contract, err := bindIiZiSwapQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapQuoterCaller{contract: contract}, nil
}

// NewIiZiSwapQuoterTransactor creates a new write-only instance of IiZiSwapQuoter, bound to a specific deployed contract.
func NewIiZiSwapQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*IiZiSwapQuoterTransactor, error) {
	contract, err := bindIiZiSwapQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapQuoterTransactor{contract: contract}, nil
}

// NewIiZiSwapQuoterFilterer creates a new log filterer instance of IiZiSwapQuoter, bound to a specific deployed contract.
func NewIiZiSwapQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*IiZiSwapQuoterFilterer, error) {
	contract, err := bindIiZiSwapQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapQuoterFilterer{contract: contract}, nil
}

// bindIiZiSwapQuoter binds a generic wrapper to an already deployed contract.
func bindIiZiSwapQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IiZiSwapQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapQuoter *IiZiSwapQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapQuoter.Contract.IiZiSwapQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapQuoter *IiZiSwapQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.IiZiSwapQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapQuoter *IiZiSwapQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.IiZiSwapQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapQuoter *IiZiSwapQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapQuoter *IiZiSwapQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapQuoter *IiZiSwapQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.contract.Transact(opts, method, params...)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwapQuoter *IiZiSwapQuoterTransactor) SwapAmount(opts *bind.TransactOpts, params IiZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwapQuoter.contract.Transact(opts, "swapAmount", params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwapQuoter *IiZiSwapQuoterSession) SwapAmount(params IiZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.SwapAmount(&_IiZiSwapQuoter.TransactOpts, params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwapQuoter *IiZiSwapQuoterTransactorSession) SwapAmount(params IiZiSwapQuoterSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwapQuoter.Contract.SwapAmount(&_IiZiSwapQuoter.TransactOpts, params)
}
