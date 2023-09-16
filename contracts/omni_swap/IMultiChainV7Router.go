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

// IMultiChainV7RouterProxyInfo is an auto generated low-level Go binding around an user-defined struct.
type IMultiChainV7RouterProxyInfo struct {
	Supported      bool
	AcceptAnyToken bool
}

// IMultiChainV7RouterMetaData contains all meta data concerning the IMultiChainV7Router contract.
var IMultiChainV7RouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"LogAnySwapInAndExec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogAnySwapOutAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"anycallProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogRetryExecRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"dontExec\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"LogRetrySwapInAndExec\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapOutUnderlyingAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"anycallExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"anycallProxyInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"acceptAnyToken\",\"type\":\"bool\"}],\"internalType\":\"structIMultiChainV7Router.ProxyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMultiChainV7RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMultiChainV7RouterMetaData.ABI instead.
var IMultiChainV7RouterABI = IMultiChainV7RouterMetaData.ABI

// IMultiChainV7Router is an auto generated Go binding around an Ethereum contract.
type IMultiChainV7Router struct {
	IMultiChainV7RouterCaller     // Read-only binding to the contract
	IMultiChainV7RouterTransactor // Write-only binding to the contract
	IMultiChainV7RouterFilterer   // Log filterer for contract events
}

// IMultiChainV7RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMultiChainV7RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiChainV7RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMultiChainV7RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiChainV7RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMultiChainV7RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiChainV7RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMultiChainV7RouterSession struct {
	Contract     *IMultiChainV7Router // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMultiChainV7RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMultiChainV7RouterCallerSession struct {
	Contract *IMultiChainV7RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMultiChainV7RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMultiChainV7RouterTransactorSession struct {
	Contract     *IMultiChainV7RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMultiChainV7RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMultiChainV7RouterRaw struct {
	Contract *IMultiChainV7Router // Generic contract binding to access the raw methods on
}

// IMultiChainV7RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMultiChainV7RouterCallerRaw struct {
	Contract *IMultiChainV7RouterCaller // Generic read-only contract binding to access the raw methods on
}

// IMultiChainV7RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMultiChainV7RouterTransactorRaw struct {
	Contract *IMultiChainV7RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMultiChainV7Router creates a new instance of IMultiChainV7Router, bound to a specific deployed contract.
func NewIMultiChainV7Router(address common.Address, backend bind.ContractBackend) (*IMultiChainV7Router, error) {
	contract, err := bindIMultiChainV7Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7Router{IMultiChainV7RouterCaller: IMultiChainV7RouterCaller{contract: contract}, IMultiChainV7RouterTransactor: IMultiChainV7RouterTransactor{contract: contract}, IMultiChainV7RouterFilterer: IMultiChainV7RouterFilterer{contract: contract}}, nil
}

// NewIMultiChainV7RouterCaller creates a new read-only instance of IMultiChainV7Router, bound to a specific deployed contract.
func NewIMultiChainV7RouterCaller(address common.Address, caller bind.ContractCaller) (*IMultiChainV7RouterCaller, error) {
	contract, err := bindIMultiChainV7Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterCaller{contract: contract}, nil
}

// NewIMultiChainV7RouterTransactor creates a new write-only instance of IMultiChainV7Router, bound to a specific deployed contract.
func NewIMultiChainV7RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMultiChainV7RouterTransactor, error) {
	contract, err := bindIMultiChainV7Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterTransactor{contract: contract}, nil
}

// NewIMultiChainV7RouterFilterer creates a new log filterer instance of IMultiChainV7Router, bound to a specific deployed contract.
func NewIMultiChainV7RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMultiChainV7RouterFilterer, error) {
	contract, err := bindIMultiChainV7Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterFilterer{contract: contract}, nil
}

// bindIMultiChainV7Router binds a generic wrapper to an already deployed contract.
func bindIMultiChainV7Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMultiChainV7RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMultiChainV7Router *IMultiChainV7RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMultiChainV7Router.Contract.IMultiChainV7RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMultiChainV7Router *IMultiChainV7RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.IMultiChainV7RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMultiChainV7Router *IMultiChainV7RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.IMultiChainV7RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMultiChainV7Router *IMultiChainV7RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMultiChainV7Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMultiChainV7Router *IMultiChainV7RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMultiChainV7Router *IMultiChainV7RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.contract.Transact(opts, method, params...)
}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_IMultiChainV7Router *IMultiChainV7RouterCaller) AnycallExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMultiChainV7Router.contract.Call(opts, &out, "anycallExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_IMultiChainV7Router *IMultiChainV7RouterSession) AnycallExecutor() (common.Address, error) {
	return _IMultiChainV7Router.Contract.AnycallExecutor(&_IMultiChainV7Router.CallOpts)
}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_IMultiChainV7Router *IMultiChainV7RouterCallerSession) AnycallExecutor() (common.Address, error) {
	return _IMultiChainV7Router.Contract.AnycallExecutor(&_IMultiChainV7Router.CallOpts)
}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address proxy) view returns((bool,bool))
func (_IMultiChainV7Router *IMultiChainV7RouterCaller) AnycallProxyInfo(opts *bind.CallOpts, proxy common.Address) (IMultiChainV7RouterProxyInfo, error) {
	var out []interface{}
	err := _IMultiChainV7Router.contract.Call(opts, &out, "anycallProxyInfo", proxy)

	if err != nil {
		return *new(IMultiChainV7RouterProxyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMultiChainV7RouterProxyInfo)).(*IMultiChainV7RouterProxyInfo)

	return out0, err

}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address proxy) view returns((bool,bool))
func (_IMultiChainV7Router *IMultiChainV7RouterSession) AnycallProxyInfo(proxy common.Address) (IMultiChainV7RouterProxyInfo, error) {
	return _IMultiChainV7Router.Contract.AnycallProxyInfo(&_IMultiChainV7Router.CallOpts, proxy)
}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address proxy) view returns((bool,bool))
func (_IMultiChainV7Router *IMultiChainV7RouterCallerSession) AnycallProxyInfo(proxy common.Address) (IMultiChainV7RouterProxyInfo, error) {
	return _IMultiChainV7Router.Contract.AnycallProxyInfo(&_IMultiChainV7Router.CallOpts, proxy)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_IMultiChainV7Router *IMultiChainV7RouterTransactor) AnySwapOutUnderlyingAndCall(opts *bind.TransactOpts, token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _IMultiChainV7Router.contract.Transact(opts, "anySwapOutUnderlyingAndCall", token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_IMultiChainV7Router *IMultiChainV7RouterSession) AnySwapOutUnderlyingAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.AnySwapOutUnderlyingAndCall(&_IMultiChainV7Router.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_IMultiChainV7Router *IMultiChainV7RouterTransactorSession) AnySwapOutUnderlyingAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _IMultiChainV7Router.Contract.AnySwapOutUnderlyingAndCall(&_IMultiChainV7Router.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// IMultiChainV7RouterLogAnySwapInAndExecIterator is returned from FilterLogAnySwapInAndExec and is used to iterate over the raw logs and unpacked data for LogAnySwapInAndExec events raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogAnySwapInAndExecIterator struct {
	Event *IMultiChainV7RouterLogAnySwapInAndExec // Event containing the contract specifics and raw log

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
func (it *IMultiChainV7RouterLogAnySwapInAndExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiChainV7RouterLogAnySwapInAndExec)
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
		it.Event = new(IMultiChainV7RouterLogAnySwapInAndExec)
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
func (it *IMultiChainV7RouterLogAnySwapInAndExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiChainV7RouterLogAnySwapInAndExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiChainV7RouterLogAnySwapInAndExec represents a LogAnySwapInAndExec event raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogAnySwapInAndExec struct {
	SwapID      string
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
	Success     bool
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapInAndExec is a free log retrieval operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) FilterLogAnySwapInAndExec(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (*IMultiChainV7RouterLogAnySwapInAndExecIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IMultiChainV7Router.contract.FilterLogs(opts, "LogAnySwapInAndExec", swapoutIDRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterLogAnySwapInAndExecIterator{contract: _IMultiChainV7Router.contract, event: "LogAnySwapInAndExec", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapInAndExec is a free log subscription operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) WatchLogAnySwapInAndExec(opts *bind.WatchOpts, sink chan<- *IMultiChainV7RouterLogAnySwapInAndExec, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IMultiChainV7Router.contract.WatchLogs(opts, "LogAnySwapInAndExec", swapoutIDRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiChainV7RouterLogAnySwapInAndExec)
				if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogAnySwapInAndExec", log); err != nil {
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

// ParseLogAnySwapInAndExec is a log parse operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) ParseLogAnySwapInAndExec(log types.Log) (*IMultiChainV7RouterLogAnySwapInAndExec, error) {
	event := new(IMultiChainV7RouterLogAnySwapInAndExec)
	if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogAnySwapInAndExec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMultiChainV7RouterLogAnySwapOutAndCallIterator is returned from FilterLogAnySwapOutAndCall and is used to iterate over the raw logs and unpacked data for LogAnySwapOutAndCall events raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogAnySwapOutAndCallIterator struct {
	Event *IMultiChainV7RouterLogAnySwapOutAndCall // Event containing the contract specifics and raw log

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
func (it *IMultiChainV7RouterLogAnySwapOutAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiChainV7RouterLogAnySwapOutAndCall)
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
		it.Event = new(IMultiChainV7RouterLogAnySwapOutAndCall)
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
func (it *IMultiChainV7RouterLogAnySwapOutAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiChainV7RouterLogAnySwapOutAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiChainV7RouterLogAnySwapOutAndCall represents a LogAnySwapOutAndCall event raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogAnySwapOutAndCall struct {
	SwapoutID    [32]byte
	Token        common.Address
	From         common.Address
	Receiver     string
	Amount       *big.Int
	ToChainID    *big.Int
	AnycallProxy string
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapOutAndCall is a free log retrieval operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) FilterLogAnySwapOutAndCall(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, from []common.Address) (*IMultiChainV7RouterLogAnySwapOutAndCallIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IMultiChainV7Router.contract.FilterLogs(opts, "LogAnySwapOutAndCall", swapoutIDRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterLogAnySwapOutAndCallIterator{contract: _IMultiChainV7Router.contract, event: "LogAnySwapOutAndCall", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapOutAndCall is a free log subscription operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) WatchLogAnySwapOutAndCall(opts *bind.WatchOpts, sink chan<- *IMultiChainV7RouterLogAnySwapOutAndCall, swapoutID [][32]byte, token []common.Address, from []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IMultiChainV7Router.contract.WatchLogs(opts, "LogAnySwapOutAndCall", swapoutIDRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiChainV7RouterLogAnySwapOutAndCall)
				if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogAnySwapOutAndCall", log); err != nil {
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

// ParseLogAnySwapOutAndCall is a log parse operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) ParseLogAnySwapOutAndCall(log types.Log) (*IMultiChainV7RouterLogAnySwapOutAndCall, error) {
	event := new(IMultiChainV7RouterLogAnySwapOutAndCall)
	if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogAnySwapOutAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMultiChainV7RouterLogRetryExecRecordIterator is returned from FilterLogRetryExecRecord and is used to iterate over the raw logs and unpacked data for LogRetryExecRecord events raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogRetryExecRecordIterator struct {
	Event *IMultiChainV7RouterLogRetryExecRecord // Event containing the contract specifics and raw log

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
func (it *IMultiChainV7RouterLogRetryExecRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiChainV7RouterLogRetryExecRecord)
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
		it.Event = new(IMultiChainV7RouterLogRetryExecRecord)
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
func (it *IMultiChainV7RouterLogRetryExecRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiChainV7RouterLogRetryExecRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiChainV7RouterLogRetryExecRecord represents a LogRetryExecRecord event raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogRetryExecRecord struct {
	SwapID       string
	SwapoutID    [32]byte
	Token        common.Address
	Receiver     common.Address
	Amount       *big.Int
	FromChainID  *big.Int
	AnycallProxy common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogRetryExecRecord is a free log retrieval operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) FilterLogRetryExecRecord(opts *bind.FilterOpts) (*IMultiChainV7RouterLogRetryExecRecordIterator, error) {

	logs, sub, err := _IMultiChainV7Router.contract.FilterLogs(opts, "LogRetryExecRecord")
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterLogRetryExecRecordIterator{contract: _IMultiChainV7Router.contract, event: "LogRetryExecRecord", logs: logs, sub: sub}, nil
}

// WatchLogRetryExecRecord is a free log subscription operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) WatchLogRetryExecRecord(opts *bind.WatchOpts, sink chan<- *IMultiChainV7RouterLogRetryExecRecord) (event.Subscription, error) {

	logs, sub, err := _IMultiChainV7Router.contract.WatchLogs(opts, "LogRetryExecRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiChainV7RouterLogRetryExecRecord)
				if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogRetryExecRecord", log); err != nil {
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

// ParseLogRetryExecRecord is a log parse operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) ParseLogRetryExecRecord(log types.Log) (*IMultiChainV7RouterLogRetryExecRecord, error) {
	event := new(IMultiChainV7RouterLogRetryExecRecord)
	if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogRetryExecRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMultiChainV7RouterLogRetrySwapInAndExecIterator is returned from FilterLogRetrySwapInAndExec and is used to iterate over the raw logs and unpacked data for LogRetrySwapInAndExec events raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogRetrySwapInAndExecIterator struct {
	Event *IMultiChainV7RouterLogRetrySwapInAndExec // Event containing the contract specifics and raw log

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
func (it *IMultiChainV7RouterLogRetrySwapInAndExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiChainV7RouterLogRetrySwapInAndExec)
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
		it.Event = new(IMultiChainV7RouterLogRetrySwapInAndExec)
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
func (it *IMultiChainV7RouterLogRetrySwapInAndExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiChainV7RouterLogRetrySwapInAndExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiChainV7RouterLogRetrySwapInAndExec represents a LogRetrySwapInAndExec event raised by the IMultiChainV7Router contract.
type IMultiChainV7RouterLogRetrySwapInAndExec struct {
	SwapID      string
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
	DontExec    bool
	Success     bool
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRetrySwapInAndExec is a free log retrieval operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) FilterLogRetrySwapInAndExec(opts *bind.FilterOpts) (*IMultiChainV7RouterLogRetrySwapInAndExecIterator, error) {

	logs, sub, err := _IMultiChainV7Router.contract.FilterLogs(opts, "LogRetrySwapInAndExec")
	if err != nil {
		return nil, err
	}
	return &IMultiChainV7RouterLogRetrySwapInAndExecIterator{contract: _IMultiChainV7Router.contract, event: "LogRetrySwapInAndExec", logs: logs, sub: sub}, nil
}

// WatchLogRetrySwapInAndExec is a free log subscription operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) WatchLogRetrySwapInAndExec(opts *bind.WatchOpts, sink chan<- *IMultiChainV7RouterLogRetrySwapInAndExec) (event.Subscription, error) {

	logs, sub, err := _IMultiChainV7Router.contract.WatchLogs(opts, "LogRetrySwapInAndExec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiChainV7RouterLogRetrySwapInAndExec)
				if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogRetrySwapInAndExec", log); err != nil {
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

// ParseLogRetrySwapInAndExec is a log parse operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_IMultiChainV7Router *IMultiChainV7RouterFilterer) ParseLogRetrySwapInAndExec(log types.Log) (*IMultiChainV7RouterLogRetrySwapInAndExec, error) {
	event := new(IMultiChainV7RouterLogRetrySwapInAndExec)
	if err := _IMultiChainV7Router.contract.UnpackLog(event, "LogRetrySwapInAndExec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
