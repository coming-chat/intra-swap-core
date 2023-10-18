// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IiZiSwapPoolLimitOrderStruct is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapPoolLimitOrderStruct struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}

// IiZiSwapPoolMetaData contains all meta data concerning the IiZiSwapPool contract.
var IiZiSwapPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"addAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"acquireAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimSold\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"AddLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"collectDec\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"collectEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"CollectLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"CollectLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"decreaseAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimSold\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"DecLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidY\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountX\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountY\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignX\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"fromLegacy\",\"type\":\"bool\"}],\"name\":\"assignLimOrderEarnX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignY\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"fromLegacy\",\"type\":\"bool\"}],\"name\":\"assignLimOrderEarnY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amountXLim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLim\",\"type\":\"uint256\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAmountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualAmountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFeeCharged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"collectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"collectEarn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isEarnY\",\"type\":\"bool\"}],\"name\":\"collectLimOrder\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualCollectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"actualCollectEarn\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaX\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaY\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newNextQueueLen\",\"type\":\"uint16\"}],\"name\":\"expandObservationQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeChargePercent\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"}],\"name\":\"limitOrderData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarnY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarnX\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"limitOrderSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwapPool.LimitOrderStruct[]\",\"name\":\"limitOrders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"liquiditySnapshot\",\"outputs\":[{\"internalType\":\"int128[]\",\"name\":\"deltaLiquidities\",\"type\":\"int128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newFeeChargePercent\",\"type\":\"uint24\"}],\"name\":\"modifyFeeChargePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"accPoint\",\"type\":\"int56\"},{\"internalType\":\"bool\",\"name\":\"init\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"accPoints\",\"type\":\"int56[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"}],\"name\":\"orderOrEndpoint\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"val\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"pointBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"}],\"name\":\"points\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidSum\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"accFeeXOut_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accFeeYOut_128\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEndpt\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sqrtRate_96\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPrice_96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationCurrentIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationNextQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityX\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireY\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2X\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireX\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeXCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeYCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"userEarnX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"userEarnY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IiZiSwapPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IiZiSwapPoolMetaData.ABI instead.
var IiZiSwapPoolABI = IiZiSwapPoolMetaData.ABI

// IiZiSwapPool is an auto generated Go binding around an Ethereum contract.
type IiZiSwapPool struct {
	IiZiSwapPoolCaller     // Read-only binding to the contract
	IiZiSwapPoolTransactor // Write-only binding to the contract
	IiZiSwapPoolFilterer   // Log filterer for contract events
}

// IiZiSwapPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IiZiSwapPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IiZiSwapPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IiZiSwapPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IiZiSwapPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IiZiSwapPoolSession struct {
	Contract     *IiZiSwapPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IiZiSwapPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IiZiSwapPoolCallerSession struct {
	Contract *IiZiSwapPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IiZiSwapPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IiZiSwapPoolTransactorSession struct {
	Contract     *IiZiSwapPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IiZiSwapPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IiZiSwapPoolRaw struct {
	Contract *IiZiSwapPool // Generic contract binding to access the raw methods on
}

// IiZiSwapPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IiZiSwapPoolCallerRaw struct {
	Contract *IiZiSwapPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IiZiSwapPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IiZiSwapPoolTransactorRaw struct {
	Contract *IiZiSwapPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIiZiSwapPool creates a new instance of IiZiSwapPool, bound to a specific deployed contract.
func NewIiZiSwapPool(address common.Address, backend bind.ContractBackend) (*IiZiSwapPool, error) {
	contract, err := bindIiZiSwapPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPool{IiZiSwapPoolCaller: IiZiSwapPoolCaller{contract: contract}, IiZiSwapPoolTransactor: IiZiSwapPoolTransactor{contract: contract}, IiZiSwapPoolFilterer: IiZiSwapPoolFilterer{contract: contract}}, nil
}

// NewIiZiSwapPoolCaller creates a new read-only instance of IiZiSwapPool, bound to a specific deployed contract.
func NewIiZiSwapPoolCaller(address common.Address, caller bind.ContractCaller) (*IiZiSwapPoolCaller, error) {
	contract, err := bindIiZiSwapPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolCaller{contract: contract}, nil
}

// NewIiZiSwapPoolTransactor creates a new write-only instance of IiZiSwapPool, bound to a specific deployed contract.
func NewIiZiSwapPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IiZiSwapPoolTransactor, error) {
	contract, err := bindIiZiSwapPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolTransactor{contract: contract}, nil
}

// NewIiZiSwapPoolFilterer creates a new log filterer instance of IiZiSwapPool, bound to a specific deployed contract.
func NewIiZiSwapPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IiZiSwapPoolFilterer, error) {
	contract, err := bindIiZiSwapPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolFilterer{contract: contract}, nil
}

// bindIiZiSwapPool binds a generic wrapper to an already deployed contract.
func bindIiZiSwapPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IiZiSwapPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapPool *IiZiSwapPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapPool.Contract.IiZiSwapPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapPool *IiZiSwapPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.IiZiSwapPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapPool *IiZiSwapPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.IiZiSwapPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IiZiSwapPool *IiZiSwapPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IiZiSwapPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IiZiSwapPool *IiZiSwapPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IiZiSwapPool *IiZiSwapPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.contract.Transact(opts, method, params...)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_IiZiSwapPool *IiZiSwapPoolCaller) FeeChargePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "feeChargePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_IiZiSwapPool *IiZiSwapPoolSession) FeeChargePercent() (*big.Int, error) {
	return _IiZiSwapPool.Contract.FeeChargePercent(&_IiZiSwapPool.CallOpts)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) FeeChargePercent() (*big.Int, error) {
	return _IiZiSwapPool.Contract.FeeChargePercent(&_IiZiSwapPool.CallOpts)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 point) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_IiZiSwapPool *IiZiSwapPoolCaller) LimitOrderData(opts *bind.CallOpts, point *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "limitOrderData", point)

	outstruct := new(struct {
		SellingX       *big.Int
		EarnY          *big.Int
		AccEarnY       *big.Int
		LegacyAccEarnY *big.Int
		LegacyEarnY    *big.Int
		SellingY       *big.Int
		EarnX          *big.Int
		LegacyEarnX    *big.Int
		AccEarnX       *big.Int
		LegacyAccEarnX *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellingX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EarnY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccEarnY = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LegacyAccEarnY = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarnY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SellingY = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EarnX = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarnX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.AccEarnX = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LegacyAccEarnX = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 point) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_IiZiSwapPool *IiZiSwapPoolSession) LimitOrderData(point *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	return _IiZiSwapPool.Contract.LimitOrderData(&_IiZiSwapPool.CallOpts, point)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 point) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) LimitOrderData(point *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	return _IiZiSwapPool.Contract.LimitOrderData(&_IiZiSwapPool.CallOpts, point)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_IiZiSwapPool *IiZiSwapPoolCaller) LimitOrderSnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "limitOrderSnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]IiZiSwapPoolLimitOrderStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]IiZiSwapPoolLimitOrderStruct)).(*[]IiZiSwapPoolLimitOrderStruct)

	return out0, err

}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_IiZiSwapPool *IiZiSwapPoolSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _IiZiSwapPool.Contract.LimitOrderSnapshot(&_IiZiSwapPool.CallOpts, leftPoint, rightPoint)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _IiZiSwapPool.Contract.LimitOrderSnapshot(&_IiZiSwapPool.CallOpts, leftPoint, rightPoint)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IiZiSwapPool *IiZiSwapPoolCaller) Liquidity(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "liquidity", key)

	outstruct := new(struct {
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		TokenOwedX       *big.Int
		TokenOwedY       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedX = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IiZiSwapPool *IiZiSwapPoolSession) Liquidity(key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _IiZiSwapPool.Contract.Liquidity(&_IiZiSwapPool.CallOpts, key)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) Liquidity(key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _IiZiSwapPool.Contract.Liquidity(&_IiZiSwapPool.CallOpts, key)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_IiZiSwapPool *IiZiSwapPoolCaller) LiquiditySnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "liquiditySnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_IiZiSwapPool *IiZiSwapPoolSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _IiZiSwapPool.Contract.LiquiditySnapshot(&_IiZiSwapPool.CallOpts, leftPoint, rightPoint)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _IiZiSwapPool.Contract.LiquiditySnapshot(&_IiZiSwapPool.CallOpts, leftPoint, rightPoint)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_IiZiSwapPool *IiZiSwapPoolCaller) Observations(opts *bind.CallOpts, index *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "observations", index)

	outstruct := new(struct {
		Timestamp uint32
		AccPoint  *big.Int
		Init      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AccPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Init = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_IiZiSwapPool *IiZiSwapPoolSession) Observations(index *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _IiZiSwapPool.Contract.Observations(&_IiZiSwapPool.CallOpts, index)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) Observations(index *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _IiZiSwapPool.Contract.Observations(&_IiZiSwapPool.CallOpts, index)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_IiZiSwapPool *IiZiSwapPoolCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "observe", secondsAgos)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_IiZiSwapPool *IiZiSwapPoolSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _IiZiSwapPool.Contract.Observe(&_IiZiSwapPool.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _IiZiSwapPool.Contract.Observe(&_IiZiSwapPool.CallOpts, secondsAgos)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 wordPosition) view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCaller) PointBitmap(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "pointBitmap", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 wordPosition) view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolSession) PointBitmap(wordPosition int16) (*big.Int, error) {
	return _IiZiSwapPool.Contract.PointBitmap(&_IiZiSwapPool.CallOpts, wordPosition)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 wordPosition) view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) PointBitmap(wordPosition int16) (*big.Int, error) {
	return _IiZiSwapPool.Contract.PointBitmap(&_IiZiSwapPool.CallOpts, wordPosition)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 point) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_IiZiSwapPool *IiZiSwapPoolCaller) Points(opts *bind.CallOpts, point *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "points", point)

	outstruct := new(struct {
		LiquidSum     *big.Int
		LiquidDelta   *big.Int
		AccFeeXOut128 *big.Int
		AccFeeYOut128 *big.Int
		IsEndpt       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidSum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccFeeXOut128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccFeeYOut128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsEndpt = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 point) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_IiZiSwapPool *IiZiSwapPoolSession) Points(point *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _IiZiSwapPool.Contract.Points(&_IiZiSwapPool.CallOpts, point)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 point) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) Points(point *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _IiZiSwapPool.Contract.Points(&_IiZiSwapPool.CallOpts, point)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_IiZiSwapPool *IiZiSwapPoolCaller) SqrtRate96(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "sqrtRate_96")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_IiZiSwapPool *IiZiSwapPoolSession) SqrtRate96() (*big.Int, error) {
	return _IiZiSwapPool.Contract.SqrtRate96(&_IiZiSwapPool.CallOpts)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) SqrtRate96() (*big.Int, error) {
	return _IiZiSwapPool.Contract.SqrtRate96(&_IiZiSwapPool.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IiZiSwapPool *IiZiSwapPoolCaller) State(opts *bind.CallOpts) (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SqrtPrice96             *big.Int
		CurrentPoint            *big.Int
		ObservationCurrentIndex uint16
		ObservationQueueLen     uint16
		ObservationNextQueueLen uint16
		Locked                  bool
		Liquidity               *big.Int
		LiquidityX              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationCurrentIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationQueueLen = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationNextQueueLen = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Locked = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Liquidity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IiZiSwapPool *IiZiSwapPoolSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _IiZiSwapPool.Contract.State(&_IiZiSwapPool.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _IiZiSwapPool.Contract.State(&_IiZiSwapPool.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCaller) TotalFeeXCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "totalFeeXCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolSession) TotalFeeXCharged() (*big.Int, error) {
	return _IiZiSwapPool.Contract.TotalFeeXCharged(&_IiZiSwapPool.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) TotalFeeXCharged() (*big.Int, error) {
	return _IiZiSwapPool.Contract.TotalFeeXCharged(&_IiZiSwapPool.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCaller) TotalFeeYCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "totalFeeYCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolSession) TotalFeeYCharged() (*big.Int, error) {
	return _IiZiSwapPool.Contract.TotalFeeYCharged(&_IiZiSwapPool.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) TotalFeeYCharged() (*big.Int, error) {
	return _IiZiSwapPool.Contract.TotalFeeYCharged(&_IiZiSwapPool.CallOpts)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolCaller) UserEarnX(opts *bind.CallOpts, key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "userEarnX", key)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		LegacyEarn    *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarn = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolSession) UserEarnX(key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _IiZiSwapPool.Contract.UserEarnX(&_IiZiSwapPool.CallOpts, key)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) UserEarnX(key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _IiZiSwapPool.Contract.UserEarnX(&_IiZiSwapPool.CallOpts, key)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolCaller) UserEarnY(opts *bind.CallOpts, key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _IiZiSwapPool.contract.Call(opts, &out, "userEarnY", key)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		LegacyEarn    *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarn = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolSession) UserEarnY(key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _IiZiSwapPool.Contract.UserEarnY(&_IiZiSwapPool.CallOpts, key)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 key) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_IiZiSwapPool *IiZiSwapPoolCallerSession) UserEarnY(key [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _IiZiSwapPool.Contract.UserEarnY(&_IiZiSwapPool.CallOpts, key)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) AddLimOrderWithX(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "addLimOrderWithX", recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_IiZiSwapPool *IiZiSwapPoolSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AddLimOrderWithX(&_IiZiSwapPool.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AddLimOrderWithX(&_IiZiSwapPool.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) AddLimOrderWithY(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "addLimOrderWithY", recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_IiZiSwapPool *IiZiSwapPoolSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AddLimOrderWithY(&_IiZiSwapPool.TransactOpts, recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AddLimOrderWithY(&_IiZiSwapPool.TransactOpts, recipient, point, amountY, data)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) AssignLimOrderEarnX(opts *bind.TransactOpts, point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "assignLimOrderEarnX", point, assignX, fromLegacy)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_IiZiSwapPool *IiZiSwapPoolSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AssignLimOrderEarnX(&_IiZiSwapPool.TransactOpts, point, assignX, fromLegacy)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AssignLimOrderEarnX(&_IiZiSwapPool.TransactOpts, point, assignX, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) AssignLimOrderEarnY(opts *bind.TransactOpts, point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "assignLimOrderEarnY", point, assignY, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_IiZiSwapPool *IiZiSwapPoolSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AssignLimOrderEarnY(&_IiZiSwapPool.TransactOpts, point, assignY, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.AssignLimOrderEarnY(&_IiZiSwapPool.TransactOpts, point, assignY, fromLegacy)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) Burn(opts *bind.TransactOpts, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "burn", leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Burn(&_IiZiSwapPool.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Burn(&_IiZiSwapPool.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "collect", recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Collect(&_IiZiSwapPool.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Collect(&_IiZiSwapPool.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactor) CollectFeeCharged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "collectFeeCharged")
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_IiZiSwapPool *IiZiSwapPoolSession) CollectFeeCharged() (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.CollectFeeCharged(&_IiZiSwapPool.TransactOpts)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) CollectFeeCharged() (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.CollectFeeCharged(&_IiZiSwapPool.TransactOpts)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) CollectLimOrder(opts *bind.TransactOpts, recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "collectLimOrder", recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_IiZiSwapPool *IiZiSwapPoolSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.CollectLimOrder(&_IiZiSwapPool.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.CollectLimOrder(&_IiZiSwapPool.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) DecLimOrderWithX(opts *bind.TransactOpts, point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "decLimOrderWithX", point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.DecLimOrderWithX(&_IiZiSwapPool.TransactOpts, point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.DecLimOrderWithX(&_IiZiSwapPool.TransactOpts, point, deltaX)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) DecLimOrderWithY(opts *bind.TransactOpts, point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "decLimOrderWithY", point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.DecLimOrderWithY(&_IiZiSwapPool.TransactOpts, point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.DecLimOrderWithY(&_IiZiSwapPool.TransactOpts, point, deltaY)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactor) ExpandObservationQueue(opts *bind.TransactOpts, newNextQueueLen uint16) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "expandObservationQueue", newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_IiZiSwapPool *IiZiSwapPoolSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.ExpandObservationQueue(&_IiZiSwapPool.TransactOpts, newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.ExpandObservationQueue(&_IiZiSwapPool.TransactOpts, newNextQueueLen)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "flash", recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_IiZiSwapPool *IiZiSwapPoolSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Flash(&_IiZiSwapPool.TransactOpts, recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Flash(&_IiZiSwapPool.TransactOpts, recipient, amountX, amountY, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "mint", recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Mint(&_IiZiSwapPool.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.Mint(&_IiZiSwapPool.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactor) ModifyFeeChargePercent(opts *bind.TransactOpts, newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "modifyFeeChargePercent", newFeeChargePercent)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_IiZiSwapPool *IiZiSwapPoolSession) ModifyFeeChargePercent(newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.ModifyFeeChargePercent(&_IiZiSwapPool.TransactOpts, newFeeChargePercent)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) ModifyFeeChargePercent(newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.ModifyFeeChargePercent(&_IiZiSwapPool.TransactOpts, newFeeChargePercent)
}

// OrderOrEndpoint is a paid mutator transaction binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 point) returns(int24 val)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) OrderOrEndpoint(opts *bind.TransactOpts, point *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "orderOrEndpoint", point)
}

// OrderOrEndpoint is a paid mutator transaction binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 point) returns(int24 val)
func (_IiZiSwapPool *IiZiSwapPoolSession) OrderOrEndpoint(point *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.OrderOrEndpoint(&_IiZiSwapPool.TransactOpts, point)
}

// OrderOrEndpoint is a paid mutator transaction binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 point) returns(int24 val)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) OrderOrEndpoint(point *big.Int) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.OrderOrEndpoint(&_IiZiSwapPool.TransactOpts, point)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) SwapX2Y(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "swapX2Y", recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapX2Y(&_IiZiSwapPool.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapX2Y(&_IiZiSwapPool.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) SwapX2YDesireY(opts *bind.TransactOpts, recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "swapX2YDesireY", recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapX2YDesireY(&_IiZiSwapPool.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapX2YDesireY(&_IiZiSwapPool.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) SwapY2X(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "swapY2X", recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapY2X(&_IiZiSwapPool.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapY2X(&_IiZiSwapPool.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactor) SwapY2XDesireX(opts *bind.TransactOpts, recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.contract.Transact(opts, "swapY2XDesireX", recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapY2XDesireX(&_IiZiSwapPool.TransactOpts, recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolTransactorSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _IiZiSwapPool.Contract.SwapY2XDesireX(&_IiZiSwapPool.TransactOpts, recipient, desireX, highPt, data)
}

// IiZiSwapPoolAddLimitOrderIterator is returned from FilterAddLimitOrder and is used to iterate over the raw logs and unpacked data for AddLimitOrder events raised by the IiZiSwapPool contract.
type IiZiSwapPoolAddLimitOrderIterator struct {
	Event *IiZiSwapPoolAddLimitOrder // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolAddLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolAddLimitOrder)
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
		it.Event = new(IiZiSwapPoolAddLimitOrder)
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
func (it *IiZiSwapPoolAddLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolAddLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolAddLimitOrder represents a AddLimitOrder event raised by the IiZiSwapPool contract.
type IiZiSwapPoolAddLimitOrder struct {
	Owner         common.Address
	AddAmount     *big.Int
	AcquireAmount *big.Int
	Point         *big.Int
	ClaimSold     *big.Int
	ClaimEarn     *big.Int
	SellXEarnY    bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLimitOrder is a free log retrieval operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterAddLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*IiZiSwapPoolAddLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "AddLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolAddLimitOrderIterator{contract: _IiZiSwapPool.contract, event: "AddLimitOrder", logs: logs, sub: sub}, nil
}

// WatchAddLimitOrder is a free log subscription operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchAddLimitOrder(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolAddLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "AddLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolAddLimitOrder)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
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

// ParseAddLimitOrder is a log parse operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseAddLimitOrder(log types.Log) (*IiZiSwapPoolAddLimitOrder, error) {
	event := new(IiZiSwapPoolAddLimitOrder)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IiZiSwapPool contract.
type IiZiSwapPoolBurnIterator struct {
	Event *IiZiSwapPoolBurn // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolBurn)
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
		it.Event = new(IiZiSwapPoolBurn)
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
func (it *IiZiSwapPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolBurn represents a Burn event raised by the IiZiSwapPool contract.
type IiZiSwapPoolBurn struct {
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*IiZiSwapPoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolBurnIterator{contract: _IiZiSwapPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolBurn, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolBurn)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseBurn(log types.Log) (*IiZiSwapPoolBurn, error) {
	event := new(IiZiSwapPoolBurn)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolCollectLimitOrderIterator is returned from FilterCollectLimitOrder and is used to iterate over the raw logs and unpacked data for CollectLimitOrder events raised by the IiZiSwapPool contract.
type IiZiSwapPoolCollectLimitOrderIterator struct {
	Event *IiZiSwapPoolCollectLimitOrder // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolCollectLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolCollectLimitOrder)
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
		it.Event = new(IiZiSwapPoolCollectLimitOrder)
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
func (it *IiZiSwapPoolCollectLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolCollectLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolCollectLimitOrder represents a CollectLimitOrder event raised by the IiZiSwapPool contract.
type IiZiSwapPoolCollectLimitOrder struct {
	Owner       common.Address
	Recipient   common.Address
	Point       *big.Int
	CollectDec  *big.Int
	CollectEarn *big.Int
	SellXEarnY  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCollectLimitOrder is a free log retrieval operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterCollectLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*IiZiSwapPoolCollectLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "CollectLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolCollectLimitOrderIterator{contract: _IiZiSwapPool.contract, event: "CollectLimitOrder", logs: logs, sub: sub}, nil
}

// WatchCollectLimitOrder is a free log subscription operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchCollectLimitOrder(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolCollectLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "CollectLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolCollectLimitOrder)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "CollectLimitOrder", log); err != nil {
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

// ParseCollectLimitOrder is a log parse operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseCollectLimitOrder(log types.Log) (*IiZiSwapPoolCollectLimitOrder, error) {
	event := new(IiZiSwapPoolCollectLimitOrder)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "CollectLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolCollectLiquidityIterator is returned from FilterCollectLiquidity and is used to iterate over the raw logs and unpacked data for CollectLiquidity events raised by the IiZiSwapPool contract.
type IiZiSwapPoolCollectLiquidityIterator struct {
	Event *IiZiSwapPoolCollectLiquidity // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolCollectLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolCollectLiquidity)
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
		it.Event = new(IiZiSwapPoolCollectLiquidity)
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
func (it *IiZiSwapPoolCollectLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolCollectLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolCollectLiquidity represents a CollectLiquidity event raised by the IiZiSwapPool contract.
type IiZiSwapPoolCollectLiquidity struct {
	Owner      common.Address
	Recipient  common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectLiquidity is a free log retrieval operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterCollectLiquidity(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*IiZiSwapPoolCollectLiquidityIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "CollectLiquidity", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolCollectLiquidityIterator{contract: _IiZiSwapPool.contract, event: "CollectLiquidity", logs: logs, sub: sub}, nil
}

// WatchCollectLiquidity is a free log subscription operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchCollectLiquidity(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolCollectLiquidity, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "CollectLiquidity", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolCollectLiquidity)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "CollectLiquidity", log); err != nil {
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

// ParseCollectLiquidity is a log parse operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseCollectLiquidity(log types.Log) (*IiZiSwapPoolCollectLiquidity, error) {
	event := new(IiZiSwapPoolCollectLiquidity)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "CollectLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolDecLimitOrderIterator is returned from FilterDecLimitOrder and is used to iterate over the raw logs and unpacked data for DecLimitOrder events raised by the IiZiSwapPool contract.
type IiZiSwapPoolDecLimitOrderIterator struct {
	Event *IiZiSwapPoolDecLimitOrder // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolDecLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolDecLimitOrder)
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
		it.Event = new(IiZiSwapPoolDecLimitOrder)
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
func (it *IiZiSwapPoolDecLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolDecLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolDecLimitOrder represents a DecLimitOrder event raised by the IiZiSwapPool contract.
type IiZiSwapPoolDecLimitOrder struct {
	Owner          common.Address
	DecreaseAmount *big.Int
	Point          *big.Int
	ClaimSold      *big.Int
	ClaimEarn      *big.Int
	SellXEarnY     bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDecLimitOrder is a free log retrieval operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterDecLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*IiZiSwapPoolDecLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "DecLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolDecLimitOrderIterator{contract: _IiZiSwapPool.contract, event: "DecLimitOrder", logs: logs, sub: sub}, nil
}

// WatchDecLimitOrder is a free log subscription operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchDecLimitOrder(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolDecLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "DecLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolDecLimitOrder)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
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

// ParseDecLimitOrder is a log parse operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseDecLimitOrder(log types.Log) (*IiZiSwapPoolDecLimitOrder, error) {
	event := new(IiZiSwapPoolDecLimitOrder)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the IiZiSwapPool contract.
type IiZiSwapPoolFlashIterator struct {
	Event *IiZiSwapPoolFlash // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolFlash)
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
		it.Event = new(IiZiSwapPoolFlash)
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
func (it *IiZiSwapPoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolFlash represents a Flash event raised by the IiZiSwapPool contract.
type IiZiSwapPoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	AmountX   *big.Int
	AmountY   *big.Int
	PaidX     *big.Int
	PaidY     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IiZiSwapPoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolFlashIterator{contract: _IiZiSwapPool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolFlash)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseFlash(log types.Log) (*IiZiSwapPoolFlash, error) {
	event := new(IiZiSwapPoolFlash)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IiZiSwapPool contract.
type IiZiSwapPoolMintIterator struct {
	Event *IiZiSwapPoolMint // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolMint)
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
		it.Event = new(IiZiSwapPoolMint)
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
func (it *IiZiSwapPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolMint represents a Mint event raised by the IiZiSwapPool contract.
type IiZiSwapPoolMint struct {
	Sender     common.Address
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*IiZiSwapPoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolMintIterator{contract: _IiZiSwapPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolMint, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolMint)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseMint(log types.Log) (*IiZiSwapPoolMint, error) {
	event := new(IiZiSwapPoolMint)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IiZiSwapPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IiZiSwapPool contract.
type IiZiSwapPoolSwapIterator struct {
	Event *IiZiSwapPoolSwap // Event containing the contract specifics and raw log

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
func (it *IiZiSwapPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IiZiSwapPoolSwap)
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
		it.Event = new(IiZiSwapPoolSwap)
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
func (it *IiZiSwapPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IiZiSwapPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IiZiSwapPoolSwap represents a Swap event raised by the IiZiSwapPool contract.
type IiZiSwapPoolSwap struct {
	TokenX       common.Address
	TokenY       common.Address
	Fee          *big.Int
	SellXEarnY   bool
	AmountX      *big.Int
	AmountY      *big.Int
	CurrentPoint *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) FilterSwap(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (*IiZiSwapPoolSwapIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.FilterLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &IiZiSwapPoolSwapIterator{contract: _IiZiSwapPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IiZiSwapPoolSwap, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IiZiSwapPool.contract.WatchLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IiZiSwapPoolSwap)
				if err := _IiZiSwapPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_IiZiSwapPool *IiZiSwapPoolFilterer) ParseSwap(log types.Log) (*IiZiSwapPoolSwap, error) {
	event := new(IiZiSwapPoolSwap)
	if err := _IiZiSwapPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
