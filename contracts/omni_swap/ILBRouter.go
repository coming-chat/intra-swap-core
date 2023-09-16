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

// ILBRouterPath is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterPath struct {
	PairBinSteps []*big.Int
	Versions     []uint8
	TokenPath    []common.Address
}

// ILBRouterMetaData contains all meta data concerning the ILBRouter contract.
var ILBRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageBPTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__BinReserveOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__BrokenSwapSafetyCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"LBRouter__DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBRouter__FailedToSendNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdDesiredOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"LBRouter__IdOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrongToken\",\"type\":\"address\"}],\"name\":\"LBRouter__InvalidTokenPath\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"LBRouter__MaxAmountInExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__NotFactoryOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBRouter__PairNotCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__SenderIsNotWNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__SwapOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"LBRouter__TooMuchTokensIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongNativeLiquidityParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__WrongTokenOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVESupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapNATIVEForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactNATIVE\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILBRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ILBRouterMetaData.ABI instead.
var ILBRouterABI = ILBRouterMetaData.ABI

// ILBRouter is an auto generated Go binding around an Ethereum contract.
type ILBRouter struct {
	ILBRouterCaller     // Read-only binding to the contract
	ILBRouterTransactor // Write-only binding to the contract
	ILBRouterFilterer   // Log filterer for contract events
}

// ILBRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILBRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILBRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILBRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILBRouterSession struct {
	Contract     *ILBRouter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILBRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILBRouterCallerSession struct {
	Contract *ILBRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ILBRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILBRouterTransactorSession struct {
	Contract     *ILBRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ILBRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILBRouterRaw struct {
	Contract *ILBRouter // Generic contract binding to access the raw methods on
}

// ILBRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILBRouterCallerRaw struct {
	Contract *ILBRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ILBRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILBRouterTransactorRaw struct {
	Contract *ILBRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILBRouter creates a new instance of ILBRouter, bound to a specific deployed contract.
func NewILBRouter(address common.Address, backend bind.ContractBackend) (*ILBRouter, error) {
	contract, err := bindILBRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILBRouter{ILBRouterCaller: ILBRouterCaller{contract: contract}, ILBRouterTransactor: ILBRouterTransactor{contract: contract}, ILBRouterFilterer: ILBRouterFilterer{contract: contract}}, nil
}

// NewILBRouterCaller creates a new read-only instance of ILBRouter, bound to a specific deployed contract.
func NewILBRouterCaller(address common.Address, caller bind.ContractCaller) (*ILBRouterCaller, error) {
	contract, err := bindILBRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILBRouterCaller{contract: contract}, nil
}

// NewILBRouterTransactor creates a new write-only instance of ILBRouter, bound to a specific deployed contract.
func NewILBRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ILBRouterTransactor, error) {
	contract, err := bindILBRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILBRouterTransactor{contract: contract}, nil
}

// NewILBRouterFilterer creates a new log filterer instance of ILBRouter, bound to a specific deployed contract.
func NewILBRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ILBRouterFilterer, error) {
	contract, err := bindILBRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILBRouterFilterer{contract: contract}, nil
}

// bindILBRouter binds a generic wrapper to an already deployed contract.
func bindILBRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILBRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILBRouter *ILBRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILBRouter.Contract.ILBRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILBRouter *ILBRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBRouter.Contract.ILBRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILBRouter *ILBRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILBRouter.Contract.ILBRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILBRouter *ILBRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILBRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILBRouter *ILBRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILBRouter *ILBRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILBRouter.Contract.contract.Transact(opts, method, params...)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactNATIVEForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactNATIVEForTokens", amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactNATIVEForTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactNATIVEForTokens(&_ILBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokens is a paid mutator transaction binding the contract method 0xb066ea7c.
//
// Solidity: function swapExactNATIVEForTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactNATIVEForTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactNATIVEForTokens(&_ILBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactNATIVEForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactNATIVEForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe038e6dc.
//
// Solidity: function swapExactNATIVEForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactNATIVEForTokensSupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactTokensForNATIVE(opts *bind.TransactOpts, amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactTokensForNATIVE", amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactTokensForNATIVE(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForNATIVE(&_ILBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVE is a paid mutator transaction binding the contract method 0x9ab6156b.
//
// Solidity: function swapExactTokensForNATIVE(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactTokensForNATIVE(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForNATIVE(&_ILBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactTokensForNATIVESupportingFeeOnTransferTokens", amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForNATIVESupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForNATIVESupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x1a24f9a9.
//
// Solidity: function swapExactTokensForNATIVESupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMinNATIVE, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactTokensForNATIVESupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMinNATIVE *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForNATIVESupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMinNATIVE, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x2a443fae.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x4b801870.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256 amountOut)
func (_ILBRouter *ILBRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_ILBRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactor) SwapNATIVEForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapNATIVEForExactTokens", amountOut, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterSession) SwapNATIVEForExactTokens(amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapNATIVEForExactTokens(&_ILBRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapNATIVEForExactTokens is a paid mutator transaction binding the contract method 0x2075ad22.
//
// Solidity: function swapNATIVEForExactTokens(uint256 amountOut, (uint256[],uint8[],address[]) path, address to, uint256 deadline) payable returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactorSession) SwapNATIVEForExactTokens(amountOut *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapNATIVEForExactTokens(&_ILBRouter.TransactOpts, amountOut, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactor) SwapTokensForExactNATIVE(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapTokensForExactNATIVE", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterSession) SwapTokensForExactNATIVE(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapTokensForExactNATIVE(&_ILBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactNATIVE is a paid mutator transaction binding the contract method 0x3dc8f8ec.
//
// Solidity: function swapTokensForExactNATIVE(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactorSession) SwapTokensForExactNATIVE(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapTokensForExactNATIVE(&_ILBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapTokensForExactTokens(&_ILBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x92fe8e70.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (uint256[],uint8[],address[]) path, address to, uint256 deadline) returns(uint256[] amountsIn)
func (_ILBRouter *ILBRouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path ILBRouterPath, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SwapTokensForExactTokens(&_ILBRouter.TransactOpts, amountOut, amountInMax, path, to, deadline)
}
