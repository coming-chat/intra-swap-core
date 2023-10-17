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

// ICelerBridgeMetaData contains all meta data concerning the ICelerBridge contract.
var ICelerBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"sendNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICelerBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ICelerBridgeMetaData.ABI instead.
var ICelerBridgeABI = ICelerBridgeMetaData.ABI

// ICelerBridge is an auto generated Go binding around an Ethereum contract.
type ICelerBridge struct {
	ICelerBridgeCaller     // Read-only binding to the contract
	ICelerBridgeTransactor // Write-only binding to the contract
	ICelerBridgeFilterer   // Log filterer for contract events
}

// ICelerBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICelerBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICelerBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICelerBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICelerBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICelerBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICelerBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICelerBridgeSession struct {
	Contract     *ICelerBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICelerBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICelerBridgeCallerSession struct {
	Contract *ICelerBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ICelerBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICelerBridgeTransactorSession struct {
	Contract     *ICelerBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ICelerBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICelerBridgeRaw struct {
	Contract *ICelerBridge // Generic contract binding to access the raw methods on
}

// ICelerBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICelerBridgeCallerRaw struct {
	Contract *ICelerBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ICelerBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICelerBridgeTransactorRaw struct {
	Contract *ICelerBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICelerBridge creates a new instance of ICelerBridge, bound to a specific deployed contract.
func NewICelerBridge(address common.Address, backend bind.ContractBackend) (*ICelerBridge, error) {
	contract, err := bindICelerBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICelerBridge{ICelerBridgeCaller: ICelerBridgeCaller{contract: contract}, ICelerBridgeTransactor: ICelerBridgeTransactor{contract: contract}, ICelerBridgeFilterer: ICelerBridgeFilterer{contract: contract}}, nil
}

// NewICelerBridgeCaller creates a new read-only instance of ICelerBridge, bound to a specific deployed contract.
func NewICelerBridgeCaller(address common.Address, caller bind.ContractCaller) (*ICelerBridgeCaller, error) {
	contract, err := bindICelerBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICelerBridgeCaller{contract: contract}, nil
}

// NewICelerBridgeTransactor creates a new write-only instance of ICelerBridge, bound to a specific deployed contract.
func NewICelerBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ICelerBridgeTransactor, error) {
	contract, err := bindICelerBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICelerBridgeTransactor{contract: contract}, nil
}

// NewICelerBridgeFilterer creates a new log filterer instance of ICelerBridge, bound to a specific deployed contract.
func NewICelerBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ICelerBridgeFilterer, error) {
	contract, err := bindICelerBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICelerBridgeFilterer{contract: contract}, nil
}

// bindICelerBridge binds a generic wrapper to an already deployed contract.
func bindICelerBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICelerBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICelerBridge *ICelerBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICelerBridge.Contract.ICelerBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICelerBridge *ICelerBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICelerBridge.Contract.ICelerBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICelerBridge *ICelerBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICelerBridge.Contract.ICelerBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICelerBridge *ICelerBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICelerBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICelerBridge *ICelerBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICelerBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICelerBridge *ICelerBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICelerBridge.Contract.contract.Transact(opts, method, params...)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeCaller) MaxSend(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "maxSend", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeSession) MaxSend(_token common.Address) (*big.Int, error) {
	return _ICelerBridge.Contract.MaxSend(&_ICelerBridge.CallOpts, _token)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeCallerSession) MaxSend(_token common.Address) (*big.Int, error) {
	return _ICelerBridge.Contract.MaxSend(&_ICelerBridge.CallOpts, _token)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeCaller) MinSend(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "minSend", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeSession) MinSend(_token common.Address) (*big.Int, error) {
	return _ICelerBridge.Contract.MinSend(&_ICelerBridge.CallOpts, _token)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address _token) view returns(uint256)
func (_ICelerBridge *ICelerBridgeCallerSession) MinSend(_token common.Address) (*big.Int, error) {
	return _ICelerBridge.Contract.MinSend(&_ICelerBridge.CallOpts, _token)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_ICelerBridge *ICelerBridgeCaller) MinimalMaxSlippage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "minimalMaxSlippage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_ICelerBridge *ICelerBridgeSession) MinimalMaxSlippage() (uint32, error) {
	return _ICelerBridge.Contract.MinimalMaxSlippage(&_ICelerBridge.CallOpts)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_ICelerBridge *ICelerBridgeCallerSession) MinimalMaxSlippage() (uint32, error) {
	return _ICelerBridge.Contract.MinimalMaxSlippage(&_ICelerBridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_ICelerBridge *ICelerBridgeCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_ICelerBridge *ICelerBridgeSession) NativeWrap() (common.Address, error) {
	return _ICelerBridge.Contract.NativeWrap(&_ICelerBridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_ICelerBridge *ICelerBridgeCallerSession) NativeWrap() (common.Address, error) {
	return _ICelerBridge.Contract.NativeWrap(&_ICelerBridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 transferId) view returns(bool)
func (_ICelerBridge *ICelerBridgeCaller) Transfers(opts *bind.CallOpts, transferId [32]byte) (bool, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "transfers", transferId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 transferId) view returns(bool)
func (_ICelerBridge *ICelerBridgeSession) Transfers(transferId [32]byte) (bool, error) {
	return _ICelerBridge.Contract.Transfers(&_ICelerBridge.CallOpts, transferId)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 transferId) view returns(bool)
func (_ICelerBridge *ICelerBridgeCallerSession) Transfers(transferId [32]byte) (bool, error) {
	return _ICelerBridge.Contract.Transfers(&_ICelerBridge.CallOpts, transferId)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ICelerBridge *ICelerBridgeCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ICelerBridge *ICelerBridgeSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ICelerBridge.Contract.VerifySigs(&_ICelerBridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ICelerBridge *ICelerBridgeCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ICelerBridge.Contract.VerifySigs(&_ICelerBridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 withdrawId) view returns(bool)
func (_ICelerBridge *ICelerBridgeCaller) Withdraws(opts *bind.CallOpts, withdrawId [32]byte) (bool, error) {
	var out []interface{}
	err := _ICelerBridge.contract.Call(opts, &out, "withdraws", withdrawId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 withdrawId) view returns(bool)
func (_ICelerBridge *ICelerBridgeSession) Withdraws(withdrawId [32]byte) (bool, error) {
	return _ICelerBridge.Contract.Withdraws(&_ICelerBridge.CallOpts, withdrawId)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 withdrawId) view returns(bool)
func (_ICelerBridge *ICelerBridgeCallerSession) Withdraws(withdrawId [32]byte) (bool, error) {
	return _ICelerBridge.Contract.Withdraws(&_ICelerBridge.CallOpts, withdrawId)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeTransactor) Relay(opts *bind.TransactOpts, _relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.contract.Transact(opts, "relay", _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Relay(&_ICelerBridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeTransactorSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Relay(&_ICelerBridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_ICelerBridge *ICelerBridgeTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.contract.Transact(opts, "send", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_ICelerBridge *ICelerBridgeSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Send(&_ICelerBridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_ICelerBridge *ICelerBridgeTransactorSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Send(&_ICelerBridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_ICelerBridge *ICelerBridgeTransactor) SendNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.contract.Transact(opts, "sendNative", _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_ICelerBridge *ICelerBridgeSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.Contract.SendNative(&_ICelerBridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_ICelerBridge *ICelerBridgeTransactorSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _ICelerBridge.Contract.SendNative(&_ICelerBridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeTransactor) Withdraw(opts *bind.TransactOpts, _wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.contract.Transact(opts, "withdraw", _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Withdraw(&_ICelerBridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_ICelerBridge *ICelerBridgeTransactorSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _ICelerBridge.Contract.Withdraw(&_ICelerBridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// ICelerBridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the ICelerBridge contract.
type ICelerBridgeSendIterator struct {
	Event *ICelerBridgeSend // Event containing the contract specifics and raw log

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
func (it *ICelerBridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICelerBridgeSend)
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
		it.Event = new(ICelerBridgeSend)
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
func (it *ICelerBridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICelerBridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICelerBridgeSend represents a Send event raised by the ICelerBridge contract.
type ICelerBridgeSend struct {
	TransferId  [32]byte
	Sender      common.Address
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	DstChainId  uint64
	Nonce       uint64
	MaxSlippage uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_ICelerBridge *ICelerBridgeFilterer) FilterSend(opts *bind.FilterOpts) (*ICelerBridgeSendIterator, error) {

	logs, sub, err := _ICelerBridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &ICelerBridgeSendIterator{contract: _ICelerBridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_ICelerBridge *ICelerBridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *ICelerBridgeSend) (event.Subscription, error) {

	logs, sub, err := _ICelerBridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICelerBridgeSend)
				if err := _ICelerBridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_ICelerBridge *ICelerBridgeFilterer) ParseSend(log types.Log) (*ICelerBridgeSend, error) {
	event := new(ICelerBridgeSend)
	if err := _ICelerBridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
