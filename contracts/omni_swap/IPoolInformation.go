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

// IPoolInformationBinInfo is an auto generated low-level Go binding around an user-defined struct.
type IPoolInformationBinInfo struct {
	Id        *big.Int
	Kind      uint8
	LowerTick int32
	ReserveA  *big.Int
	ReserveB  *big.Int
	MergeId   *big.Int
}

// IPoolInformationBinState is an auto generated low-level Go binding around an user-defined struct.
type IPoolInformationBinState struct {
	ReserveA        *big.Int
	ReserveB        *big.Int
	MergeBinBalance *big.Int
	MergeId         *big.Int
	TotalSupply     *big.Int
	Kind            uint8
	LowerTick       int32
}

// IPoolInformationMetaData contains all meta data concerning the IPoolInformation contract.
var IPoolInformationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"activeTickLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sqrtPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exactOutput\",\"type\":\"bool\"}],\"name\":\"calculateMultihopSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"tokenAIn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exactOutput\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"startBinIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endBinIndex\",\"type\":\"uint128\"}],\"name\":\"getActiveBins\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"uint128\",\"name\":\"reserveA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reserveB\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"mergeId\",\"type\":\"uint128\"}],\"internalType\":\"structIPoolInformation.BinInfo[]\",\"name\":\"bins\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"}],\"name\":\"getBinDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"tick\",\"type\":\"int32\"}],\"name\":\"getBinsAtTick\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"reserveA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reserveB\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"mergeBinBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"mergeId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"}],\"internalType\":\"structIPoolInformation.BinState[]\",\"name\":\"bins\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getSqrtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sqrtPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"tick\",\"type\":\"int32\"}],\"name\":\"tickLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sqrtPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPoolInformationABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolInformationMetaData.ABI instead.
var IPoolInformationABI = IPoolInformationMetaData.ABI

// IPoolInformation is an auto generated Go binding around an Ethereum contract.
type IPoolInformation struct {
	IPoolInformationCaller     // Read-only binding to the contract
	IPoolInformationTransactor // Write-only binding to the contract
	IPoolInformationFilterer   // Log filterer for contract events
}

// IPoolInformationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolInformationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInformationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolInformationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInformationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolInformationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInformationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolInformationSession struct {
	Contract     *IPoolInformation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolInformationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolInformationCallerSession struct {
	Contract *IPoolInformationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPoolInformationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolInformationTransactorSession struct {
	Contract     *IPoolInformationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPoolInformationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolInformationRaw struct {
	Contract *IPoolInformation // Generic contract binding to access the raw methods on
}

// IPoolInformationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolInformationCallerRaw struct {
	Contract *IPoolInformationCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolInformationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolInformationTransactorRaw struct {
	Contract *IPoolInformationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolInformation creates a new instance of IPoolInformation, bound to a specific deployed contract.
func NewIPoolInformation(address common.Address, backend bind.ContractBackend) (*IPoolInformation, error) {
	contract, err := bindIPoolInformation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolInformation{IPoolInformationCaller: IPoolInformationCaller{contract: contract}, IPoolInformationTransactor: IPoolInformationTransactor{contract: contract}, IPoolInformationFilterer: IPoolInformationFilterer{contract: contract}}, nil
}

// NewIPoolInformationCaller creates a new read-only instance of IPoolInformation, bound to a specific deployed contract.
func NewIPoolInformationCaller(address common.Address, caller bind.ContractCaller) (*IPoolInformationCaller, error) {
	contract, err := bindIPoolInformation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInformationCaller{contract: contract}, nil
}

// NewIPoolInformationTransactor creates a new write-only instance of IPoolInformation, bound to a specific deployed contract.
func NewIPoolInformationTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolInformationTransactor, error) {
	contract, err := bindIPoolInformation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInformationTransactor{contract: contract}, nil
}

// NewIPoolInformationFilterer creates a new log filterer instance of IPoolInformation, bound to a specific deployed contract.
func NewIPoolInformationFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolInformationFilterer, error) {
	contract, err := bindIPoolInformation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolInformationFilterer{contract: contract}, nil
}

// bindIPoolInformation binds a generic wrapper to an already deployed contract.
func bindIPoolInformation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoolInformationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInformation *IPoolInformationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInformation.Contract.IPoolInformationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInformation *IPoolInformationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInformation.Contract.IPoolInformationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInformation *IPoolInformationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInformation.Contract.IPoolInformationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInformation *IPoolInformationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInformation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInformation *IPoolInformationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInformation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInformation *IPoolInformationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInformation.Contract.contract.Transact(opts, method, params...)
}

// ActiveTickLiquidity is a free data retrieval call binding the contract method 0xa3a5d06d.
//
// Solidity: function activeTickLiquidity(address pool) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationCaller) ActiveTickLiquidity(opts *bind.CallOpts, pool common.Address) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "activeTickLiquidity", pool)

	outstruct := new(struct {
		SqrtPrice *big.Int
		Liquidity *big.Int
		ReserveA  *big.Int
		ReserveB  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReserveA = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActiveTickLiquidity is a free data retrieval call binding the contract method 0xa3a5d06d.
//
// Solidity: function activeTickLiquidity(address pool) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationSession) ActiveTickLiquidity(pool common.Address) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	return _IPoolInformation.Contract.ActiveTickLiquidity(&_IPoolInformation.CallOpts, pool)
}

// ActiveTickLiquidity is a free data retrieval call binding the contract method 0xa3a5d06d.
//
// Solidity: function activeTickLiquidity(address pool) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationCallerSession) ActiveTickLiquidity(pool common.Address) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	return _IPoolInformation.Contract.ActiveTickLiquidity(&_IPoolInformation.CallOpts, pool)
}

// GetActiveBins is a free data retrieval call binding the contract method 0x2d4c5a38.
//
// Solidity: function getActiveBins(address pool, uint128 startBinIndex, uint128 endBinIndex) view returns((uint128,uint8,int32,uint128,uint128,uint128)[] bins)
func (_IPoolInformation *IPoolInformationCaller) GetActiveBins(opts *bind.CallOpts, pool common.Address, startBinIndex *big.Int, endBinIndex *big.Int) ([]IPoolInformationBinInfo, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "getActiveBins", pool, startBinIndex, endBinIndex)

	if err != nil {
		return *new([]IPoolInformationBinInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolInformationBinInfo)).(*[]IPoolInformationBinInfo)

	return out0, err

}

// GetActiveBins is a free data retrieval call binding the contract method 0x2d4c5a38.
//
// Solidity: function getActiveBins(address pool, uint128 startBinIndex, uint128 endBinIndex) view returns((uint128,uint8,int32,uint128,uint128,uint128)[] bins)
func (_IPoolInformation *IPoolInformationSession) GetActiveBins(pool common.Address, startBinIndex *big.Int, endBinIndex *big.Int) ([]IPoolInformationBinInfo, error) {
	return _IPoolInformation.Contract.GetActiveBins(&_IPoolInformation.CallOpts, pool, startBinIndex, endBinIndex)
}

// GetActiveBins is a free data retrieval call binding the contract method 0x2d4c5a38.
//
// Solidity: function getActiveBins(address pool, uint128 startBinIndex, uint128 endBinIndex) view returns((uint128,uint8,int32,uint128,uint128,uint128)[] bins)
func (_IPoolInformation *IPoolInformationCallerSession) GetActiveBins(pool common.Address, startBinIndex *big.Int, endBinIndex *big.Int) ([]IPoolInformationBinInfo, error) {
	return _IPoolInformation.Contract.GetActiveBins(&_IPoolInformation.CallOpts, pool, startBinIndex, endBinIndex)
}

// GetBinDepth is a free data retrieval call binding the contract method 0x136c7589.
//
// Solidity: function getBinDepth(address pool, uint128 binId) view returns(uint256 depth)
func (_IPoolInformation *IPoolInformationCaller) GetBinDepth(opts *bind.CallOpts, pool common.Address, binId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "getBinDepth", pool, binId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBinDepth is a free data retrieval call binding the contract method 0x136c7589.
//
// Solidity: function getBinDepth(address pool, uint128 binId) view returns(uint256 depth)
func (_IPoolInformation *IPoolInformationSession) GetBinDepth(pool common.Address, binId *big.Int) (*big.Int, error) {
	return _IPoolInformation.Contract.GetBinDepth(&_IPoolInformation.CallOpts, pool, binId)
}

// GetBinDepth is a free data retrieval call binding the contract method 0x136c7589.
//
// Solidity: function getBinDepth(address pool, uint128 binId) view returns(uint256 depth)
func (_IPoolInformation *IPoolInformationCallerSession) GetBinDepth(pool common.Address, binId *big.Int) (*big.Int, error) {
	return _IPoolInformation.Contract.GetBinDepth(&_IPoolInformation.CallOpts, pool, binId)
}

// GetBinsAtTick is a free data retrieval call binding the contract method 0xd12e2bb2.
//
// Solidity: function getBinsAtTick(address pool, int32 tick) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32)[] bins)
func (_IPoolInformation *IPoolInformationCaller) GetBinsAtTick(opts *bind.CallOpts, pool common.Address, tick int32) ([]IPoolInformationBinState, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "getBinsAtTick", pool, tick)

	if err != nil {
		return *new([]IPoolInformationBinState), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolInformationBinState)).(*[]IPoolInformationBinState)

	return out0, err

}

// GetBinsAtTick is a free data retrieval call binding the contract method 0xd12e2bb2.
//
// Solidity: function getBinsAtTick(address pool, int32 tick) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32)[] bins)
func (_IPoolInformation *IPoolInformationSession) GetBinsAtTick(pool common.Address, tick int32) ([]IPoolInformationBinState, error) {
	return _IPoolInformation.Contract.GetBinsAtTick(&_IPoolInformation.CallOpts, pool, tick)
}

// GetBinsAtTick is a free data retrieval call binding the contract method 0xd12e2bb2.
//
// Solidity: function getBinsAtTick(address pool, int32 tick) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32)[] bins)
func (_IPoolInformation *IPoolInformationCallerSession) GetBinsAtTick(pool common.Address, tick int32) ([]IPoolInformationBinState, error) {
	return _IPoolInformation.Contract.GetBinsAtTick(&_IPoolInformation.CallOpts, pool, tick)
}

// GetSqrtPrice is a free data retrieval call binding the contract method 0x91c0914e.
//
// Solidity: function getSqrtPrice(address pool) view returns(uint256 sqrtPrice)
func (_IPoolInformation *IPoolInformationCaller) GetSqrtPrice(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "getSqrtPrice", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSqrtPrice is a free data retrieval call binding the contract method 0x91c0914e.
//
// Solidity: function getSqrtPrice(address pool) view returns(uint256 sqrtPrice)
func (_IPoolInformation *IPoolInformationSession) GetSqrtPrice(pool common.Address) (*big.Int, error) {
	return _IPoolInformation.Contract.GetSqrtPrice(&_IPoolInformation.CallOpts, pool)
}

// GetSqrtPrice is a free data retrieval call binding the contract method 0x91c0914e.
//
// Solidity: function getSqrtPrice(address pool) view returns(uint256 sqrtPrice)
func (_IPoolInformation *IPoolInformationCallerSession) GetSqrtPrice(pool common.Address) (*big.Int, error) {
	return _IPoolInformation.Contract.GetSqrtPrice(&_IPoolInformation.CallOpts, pool)
}

// TickLiquidity is a free data retrieval call binding the contract method 0xd388a5a6.
//
// Solidity: function tickLiquidity(address pool, int32 tick) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationCaller) TickLiquidity(opts *bind.CallOpts, pool common.Address, tick int32) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	var out []interface{}
	err := _IPoolInformation.contract.Call(opts, &out, "tickLiquidity", pool, tick)

	outstruct := new(struct {
		SqrtPrice *big.Int
		Liquidity *big.Int
		ReserveA  *big.Int
		ReserveB  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReserveA = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TickLiquidity is a free data retrieval call binding the contract method 0xd388a5a6.
//
// Solidity: function tickLiquidity(address pool, int32 tick) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationSession) TickLiquidity(pool common.Address, tick int32) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	return _IPoolInformation.Contract.TickLiquidity(&_IPoolInformation.CallOpts, pool, tick)
}

// TickLiquidity is a free data retrieval call binding the contract method 0xd388a5a6.
//
// Solidity: function tickLiquidity(address pool, int32 tick) view returns(uint256 sqrtPrice, uint256 liquidity, uint256 reserveA, uint256 reserveB)
func (_IPoolInformation *IPoolInformationCallerSession) TickLiquidity(pool common.Address, tick int32) (struct {
	SqrtPrice *big.Int
	Liquidity *big.Int
	ReserveA  *big.Int
	ReserveB  *big.Int
}, error) {
	return _IPoolInformation.Contract.TickLiquidity(&_IPoolInformation.CallOpts, pool, tick)
}

// CalculateMultihopSwap is a paid mutator transaction binding the contract method 0x21f8ce41.
//
// Solidity: function calculateMultihopSwap(bytes path, uint256 amount, bool exactOutput) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationTransactor) CalculateMultihopSwap(opts *bind.TransactOpts, path []byte, amount *big.Int, exactOutput bool) (*types.Transaction, error) {
	return _IPoolInformation.contract.Transact(opts, "calculateMultihopSwap", path, amount, exactOutput)
}

// CalculateMultihopSwap is a paid mutator transaction binding the contract method 0x21f8ce41.
//
// Solidity: function calculateMultihopSwap(bytes path, uint256 amount, bool exactOutput) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationSession) CalculateMultihopSwap(path []byte, amount *big.Int, exactOutput bool) (*types.Transaction, error) {
	return _IPoolInformation.Contract.CalculateMultihopSwap(&_IPoolInformation.TransactOpts, path, amount, exactOutput)
}

// CalculateMultihopSwap is a paid mutator transaction binding the contract method 0x21f8ce41.
//
// Solidity: function calculateMultihopSwap(bytes path, uint256 amount, bool exactOutput) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationTransactorSession) CalculateMultihopSwap(path []byte, amount *big.Int, exactOutput bool) (*types.Transaction, error) {
	return _IPoolInformation.Contract.CalculateMultihopSwap(&_IPoolInformation.TransactOpts, path, amount, exactOutput)
}

// CalculateSwap is a paid mutator transaction binding the contract method 0x2764cd0b.
//
// Solidity: function calculateSwap(address pool, uint128 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationTransactor) CalculateSwap(opts *bind.TransactOpts, pool common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int) (*types.Transaction, error) {
	return _IPoolInformation.contract.Transact(opts, "calculateSwap", pool, amount, tokenAIn, exactOutput, sqrtPriceLimit)
}

// CalculateSwap is a paid mutator transaction binding the contract method 0x2764cd0b.
//
// Solidity: function calculateSwap(address pool, uint128 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationSession) CalculateSwap(pool common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int) (*types.Transaction, error) {
	return _IPoolInformation.Contract.CalculateSwap(&_IPoolInformation.TransactOpts, pool, amount, tokenAIn, exactOutput, sqrtPriceLimit)
}

// CalculateSwap is a paid mutator transaction binding the contract method 0x2764cd0b.
//
// Solidity: function calculateSwap(address pool, uint128 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit) returns(uint256 returnAmount)
func (_IPoolInformation *IPoolInformationTransactorSession) CalculateSwap(pool common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int) (*types.Transaction, error) {
	return _IPoolInformation.Contract.CalculateSwap(&_IPoolInformation.TransactOpts, pool, amount, tokenAIn, exactOutput, sqrtPriceLimit)
}
