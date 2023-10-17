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

// IiZiSwapSwapAmountParams is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapSwapAmountParams struct {
	Path        []byte
	Recipient   common.Address
	Amount      *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// IiZiSwapSwapDesireParams is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapSwapDesireParams struct {
	Path      []byte
	Recipient common.Address
	Desire    *big.Int
	MaxPayed  *big.Int
	Deadline  *big.Int
}

// IiZiSwapSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapSwapParams struct {
	TokenX      common.Address
	TokenY      common.Address
	Fee         *big.Int
	BoundaryPt  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	MaxPayed    *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// IiZiSwapMetaData contains all meta data concerning the IiZiSwap contract.
var IiZiSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapAmountParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desire\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapDesireParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapDesire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2Y\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2X\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IiZiSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use IiZiSwapMetaData.ABI instead.
var IiZiSwapABI = IiZiSwapMetaData.ABI

// IiZiSwap is an auto generated Go binding around an Ethereum contract.
type IiZiSwap struct {
	IiZiSwapCaller     // Read-only binding to the contract
	IiZiSwapTransactor // Write-only binding to the contract
	IiZiSwapFilterer   // Log filterer for contract events
}

// IiZiSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IiZiSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IiZiSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IiZiSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IiZiSwapSession struct {
	Contract     *IiZiSwap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IiZiSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IiZiSwapCallerSession struct {
	Contract *IiZiSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IiZiSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IiZiSwapTransactorSession struct {
	Contract     *IiZiSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IiZiSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IiZiSwapRaw struct {
	Contract *IiZiSwap // Generic contract binding to access the raw methods on
}

// IiZiSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IiZiSwapCallerRaw struct {
	Contract *IiZiSwapCaller // Generic read-only contract binding to access the raw methods on
}

// IiZiSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IiZiSwapTransactorRaw struct {
	Contract *IiZiSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIiZiSwap creates a new instance of IiZiSwap, bound to a specific deployed contract.
func NewIiZiSwap(address common.Address, backend bind.ContractBackend) (*IiZiSwap, error) {
	contract, err := bindIiZiSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IiZiSwap{IiZiSwapCaller: IiZiSwapCaller{contract: contract}, IiZiSwapTransactor: IiZiSwapTransactor{contract: contract}, IiZiSwapFilterer: IiZiSwapFilterer{contract: contract}}, nil
}

// NewIiZiSwapCaller creates a new read-only instance of IiZiSwap, bound to a specific deployed contract.
func NewIiZiSwapCaller(address common.Address, caller bind.ContractCaller) (*IiZiSwapCaller, error) {
	contract, err := bindIiZiSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapCaller{contract: contract}, nil
}

// NewIiZiSwapTransactor creates a new write-only instance of IiZiSwap, bound to a specific deployed contract.
func NewIiZiSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*IiZiSwapTransactor, error) {
	contract, err := bindIiZiSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapTransactor{contract: contract}, nil
}

// NewIiZiSwapFilterer creates a new log filterer instance of IiZiSwap, bound to a specific deployed contract.
func NewIiZiSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*IiZiSwapFilterer, error) {
	contract, err := bindIiZiSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapFilterer{contract: contract}, nil
}

// bindIiZiSwap binds a generic wrapper to an already deployed contract.
func bindIiZiSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IiZiSwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwap *IiZiSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwap.Contract.IiZiSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwap *IiZiSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwap.Contract.IiZiSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwap *IiZiSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwap.Contract.IiZiSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwap *IiZiSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwap *IiZiSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwap *IiZiSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwap.Contract.contract.Transact(opts, method, params...)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapTransactor) SwapAmount(opts *bind.TransactOpts, params IiZiSwapSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapAmount", params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapSession) SwapAmount(params IiZiSwapSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapAmount(&_IiZiSwap.TransactOpts, params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapTransactorSession) SwapAmount(params IiZiSwapSwapAmountParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapAmount(&_IiZiSwap.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapTransactor) SwapDesire(opts *bind.TransactOpts, params IiZiSwapSwapDesireParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapDesire", params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapSession) SwapDesire(params IiZiSwapSwapDesireParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapDesire(&_IiZiSwap.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_IiZiSwap *IiZiSwapTransactorSession) SwapDesire(params IiZiSwapSwapDesireParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapDesire(&_IiZiSwap.TransactOpts, params)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactor) SwapX2Y(opts *bind.TransactOpts, swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapX2Y", swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapSession) SwapX2Y(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapX2Y(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactorSession) SwapX2Y(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapX2Y(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactor) SwapX2YDesireY(opts *bind.TransactOpts, swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapX2YDesireY", swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapSession) SwapX2YDesireY(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapX2YDesireY(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactorSession) SwapX2YDesireY(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapX2YDesireY(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactor) SwapY2X(opts *bind.TransactOpts, swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapY2X", swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapSession) SwapY2X(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapY2X(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactorSession) SwapY2X(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapY2X(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactor) SwapY2XDesireX(opts *bind.TransactOpts, swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.contract.Transact(opts, "swapY2XDesireX", swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapSession) SwapY2XDesireX(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapY2XDesireX(&_IiZiSwap.TransactOpts, swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_IiZiSwap *IiZiSwapTransactorSession) SwapY2XDesireX(swapParams IiZiSwapSwapParams) (*types.Transaction, error) {
	return _IiZiSwap.Contract.SwapY2XDesireX(&_IiZiSwap.TransactOpts, swapParams)
}
