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

// IMuteRouterMetaData contains all meta data concerning the IMuteRouter contract.
var IMuteRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_stable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOutExpanded\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMuteRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMuteRouterMetaData.ABI instead.
var IMuteRouterABI = IMuteRouterMetaData.ABI

// IMuteRouter is an auto generated Go binding around an Ethereum contract.
type IMuteRouter struct {
	IMuteRouterCaller     // Read-only binding to the contract
	IMuteRouterTransactor // Write-only binding to the contract
	IMuteRouterFilterer   // Log filterer for contract events
}

// IMuteRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMuteRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuteRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMuteRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuteRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMuteRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuteRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMuteRouterSession struct {
	Contract     *IMuteRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMuteRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMuteRouterCallerSession struct {
	Contract *IMuteRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMuteRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMuteRouterTransactorSession struct {
	Contract     *IMuteRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMuteRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMuteRouterRaw struct {
	Contract *IMuteRouter // Generic contract binding to access the raw methods on
}

// IMuteRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMuteRouterCallerRaw struct {
	Contract *IMuteRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IMuteRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMuteRouterTransactorRaw struct {
	Contract *IMuteRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMuteRouter creates a new instance of IMuteRouter, bound to a specific deployed contract.
func NewIMuteRouter(address common.Address, backend bind.ContractBackend) (*IMuteRouter, error) {
	contract, err := bindIMuteRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMuteRouter{IMuteRouterCaller: IMuteRouterCaller{contract: contract}, IMuteRouterTransactor: IMuteRouterTransactor{contract: contract}, IMuteRouterFilterer: IMuteRouterFilterer{contract: contract}}, nil
}

// NewIMuteRouterCaller creates a new read-only instance of IMuteRouter, bound to a specific deployed contract.
func NewIMuteRouterCaller(address common.Address, caller bind.ContractCaller) (*IMuteRouterCaller, error) {
	contract, err := bindIMuteRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMuteRouterCaller{contract: contract}, nil
}

// NewIMuteRouterTransactor creates a new write-only instance of IMuteRouter, bound to a specific deployed contract.
func NewIMuteRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMuteRouterTransactor, error) {
	contract, err := bindIMuteRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMuteRouterTransactor{contract: contract}, nil
}

// NewIMuteRouterFilterer creates a new log filterer instance of IMuteRouter, bound to a specific deployed contract.
func NewIMuteRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMuteRouterFilterer, error) {
	contract, err := bindIMuteRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMuteRouterFilterer{contract: contract}, nil
}

// bindIMuteRouter binds a generic wrapper to an already deployed contract.
func bindIMuteRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMuteRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuteRouter *IMuteRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuteRouter.Contract.IMuteRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuteRouter *IMuteRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuteRouter.Contract.IMuteRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuteRouter *IMuteRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuteRouter.Contract.IMuteRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuteRouter *IMuteRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuteRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuteRouter *IMuteRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuteRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuteRouter *IMuteRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuteRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_IMuteRouter *IMuteRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_IMuteRouter *IMuteRouterSession) WETH() (common.Address, error) {
	return _IMuteRouter.Contract.WETH(&_IMuteRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_IMuteRouter *IMuteRouterCallerSession) WETH() (common.Address, error) {
	return _IMuteRouter.Contract.WETH(&_IMuteRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IMuteRouter *IMuteRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IMuteRouter *IMuteRouterSession) Factory() (common.Address, error) {
	return _IMuteRouter.Contract.Factory(&_IMuteRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IMuteRouter *IMuteRouterCallerSession) Factory() (common.Address, error) {
	return _IMuteRouter.Contract.Factory(&_IMuteRouter.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_IMuteRouter *IMuteRouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

	outstruct := new(struct {
		AmountOut *big.Int
		Stable    bool
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_IMuteRouter *IMuteRouterSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountOut(&_IMuteRouter.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_IMuteRouter *IMuteRouterCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountOut(&_IMuteRouter.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts, bool[] _stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, stable []bool) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, path, stable)

	outstruct := new(struct {
		Amounts []*big.Int
		Stable  []bool
		Fees    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.Fees = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts, bool[] _stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountsOut(&_IMuteRouter.CallOpts, amountIn, path, stable)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts, bool[] _stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountsOut(&_IMuteRouter.CallOpts, amountIn, path, stable)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterCaller) GetAmountsOutExpanded(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "getAmountsOutExpanded", amountIn, path)

	outstruct := new(struct {
		Amounts []*big.Int
		Stable  []bool
		Fees    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.Fees = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountsOutExpanded(&_IMuteRouter.CallOpts, amountIn, path)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_IMuteRouter *IMuteRouterCallerSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _IMuteRouter.Contract.GetAmountsOutExpanded(&_IMuteRouter.CallOpts, amountIn, path)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_IMuteRouter *IMuteRouterCaller) GetPairInfo(opts *bind.CallOpts, path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "getPairInfo", path, stable)

	outstruct := new(struct {
		TokenA   common.Address
		TokenB   common.Address
		Pair     common.Address
		ReserveA *big.Int
		ReserveB *big.Int
		Fee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Pair = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReserveA = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_IMuteRouter *IMuteRouterSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _IMuteRouter.Contract.GetPairInfo(&_IMuteRouter.CallOpts, path, stable)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_IMuteRouter *IMuteRouterCallerSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _IMuteRouter.Contract.GetPairInfo(&_IMuteRouter.CallOpts, path, stable)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IMuteRouter *IMuteRouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IMuteRouter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IMuteRouter *IMuteRouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IMuteRouter.Contract.Quote(&_IMuteRouter.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IMuteRouter *IMuteRouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IMuteRouter.Contract.Quote(&_IMuteRouter.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IMuteRouter *IMuteRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IMuteRouter *IMuteRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.AddLiquidity(&_IMuteRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IMuteRouter *IMuteRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.AddLiquidity(&_IMuteRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IMuteRouter *IMuteRouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IMuteRouter *IMuteRouterSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.AddLiquidityETH(&_IMuteRouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IMuteRouter *IMuteRouterTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.AddLiquidityETH(&_IMuteRouter.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_IMuteRouter *IMuteRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_IMuteRouter *IMuteRouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidity(&_IMuteRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_IMuteRouter *IMuteRouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidity(&_IMuteRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_IMuteRouter *IMuteRouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_IMuteRouter *IMuteRouterSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidityETH(&_IMuteRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_IMuteRouter *IMuteRouterTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidityETH(&_IMuteRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_IMuteRouter *IMuteRouterTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_IMuteRouter *IMuteRouterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_IMuteRouter *IMuteRouterTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactETHForTokens(&_IMuteRouter.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactETHForTokens(&_IMuteRouter.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_IMuteRouter *IMuteRouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_IMuteRouter *IMuteRouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForETH(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForETH(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_IMuteRouter *IMuteRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _IMuteRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IMuteRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}
