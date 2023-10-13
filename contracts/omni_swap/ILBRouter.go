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

// ILBRouterLiquidityParameters is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterLiquidityParameters struct {
	TokenX          common.Address
	TokenY          common.Address
	BinStep         *big.Int
	AmountX         *big.Int
	AmountY         *big.Int
	AmountXMin      *big.Int
	AmountYMin      *big.Int
	ActiveIdDesired *big.Int
	IdSlippage      *big.Int
	DeltaIds        []*big.Int
	DistributionX   []*big.Int
	DistributionY   []*big.Int
	To              common.Address
	RefundTo        common.Address
	Deadline        *big.Int
}

// ILBRouterPath is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterPath struct {
	PairBinSteps []*big.Int
	Versions     []uint8
	TokenPath    []common.Address
}

// ILBRouterMetaData contains all meta data concerning the ILBRouter contract.
var ILBRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageBPTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__BinReserveOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__BrokenSwapSafetyCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"LBRouter__DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBRouter__FailedToSendNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdDesiredOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"LBRouter__IdOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrongToken\",\"type\":\"address\"}],\"name\":\"LBRouter__InvalidTokenPath\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"LBRouter__MaxAmountInExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__NotFactoryOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBRouter__PairNotCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__SenderIsNotWNATIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__SwapOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"LBRouter__TooMuchTokensIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongNativeLiquidityParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__WrongTokenOrder\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidityNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getIdFromPrice\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyFactory\",\"outputs\":[{\"internalType\":\"contractILBLegacyFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyRouter\",\"outputs\":[{\"internalType\":\"contractILBLegacyRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"getPriceFromId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapIn\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOutLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapOut\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountInLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getV1Factory\",\"outputs\":[{\"internalType\":\"contractIJoeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWNATIVE\",\"outputs\":[{\"internalType\":\"contractIWNATIVE\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNATIVEMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNATIVE\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactNATIVEForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinNATIVE\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForNATIVESupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapNATIVEForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactNATIVE\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"enumILBRouter.Version[]\",\"name\":\"versions\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"internalType\":\"structILBRouter.Path\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBToken\",\"name\":\"_lbToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"sweepLBToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	return bind.NewBoundContract(common2.Address(address), *parsed, caller, transactor, filterer), nil
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

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_ILBRouter *ILBRouterCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_ILBRouter *ILBRouterSession) GetFactory() (common.Address, error) {
	return _ILBRouter.Contract.GetFactory(&_ILBRouter.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_ILBRouter *ILBRouterCallerSession) GetFactory() (common.Address, error) {
	return _ILBRouter.Contract.GetFactory(&_ILBRouter.CallOpts)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address LBPair, uint256 price) view returns(uint24)
func (_ILBRouter *ILBRouterCaller) GetIdFromPrice(opts *bind.CallOpts, LBPair common.Address, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getIdFromPrice", LBPair, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address LBPair, uint256 price) view returns(uint24)
func (_ILBRouter *ILBRouterSession) GetIdFromPrice(LBPair common.Address, price *big.Int) (*big.Int, error) {
	return _ILBRouter.Contract.GetIdFromPrice(&_ILBRouter.CallOpts, LBPair, price)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address LBPair, uint256 price) view returns(uint24)
func (_ILBRouter *ILBRouterCallerSession) GetIdFromPrice(LBPair common.Address, price *big.Int) (*big.Int, error) {
	return _ILBRouter.Contract.GetIdFromPrice(&_ILBRouter.CallOpts, LBPair, price)
}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address)
func (_ILBRouter *ILBRouterCaller) GetLegacyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getLegacyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address)
func (_ILBRouter *ILBRouterSession) GetLegacyFactory() (common.Address, error) {
	return _ILBRouter.Contract.GetLegacyFactory(&_ILBRouter.CallOpts)
}

// GetLegacyFactory is a free data retrieval call binding the contract method 0x71d1974a.
//
// Solidity: function getLegacyFactory() view returns(address)
func (_ILBRouter *ILBRouterCallerSession) GetLegacyFactory() (common.Address, error) {
	return _ILBRouter.Contract.GetLegacyFactory(&_ILBRouter.CallOpts)
}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address)
func (_ILBRouter *ILBRouterCaller) GetLegacyRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getLegacyRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address)
func (_ILBRouter *ILBRouterSession) GetLegacyRouter() (common.Address, error) {
	return _ILBRouter.Contract.GetLegacyRouter(&_ILBRouter.CallOpts)
}

// GetLegacyRouter is a free data retrieval call binding the contract method 0xba846523.
//
// Solidity: function getLegacyRouter() view returns(address)
func (_ILBRouter *ILBRouterCallerSession) GetLegacyRouter() (common.Address, error) {
	return _ILBRouter.Contract.GetLegacyRouter(&_ILBRouter.CallOpts)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address LBPair, uint24 id) view returns(uint256)
func (_ILBRouter *ILBRouterCaller) GetPriceFromId(opts *bind.CallOpts, LBPair common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getPriceFromId", LBPair, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address LBPair, uint24 id) view returns(uint256)
func (_ILBRouter *ILBRouterSession) GetPriceFromId(LBPair common.Address, id *big.Int) (*big.Int, error) {
	return _ILBRouter.Contract.GetPriceFromId(&_ILBRouter.CallOpts, LBPair, id)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address LBPair, uint24 id) view returns(uint256)
func (_ILBRouter *ILBRouterCallerSession) GetPriceFromId(LBPair common.Address, id *big.Int) (*big.Int, error) {
	return _ILBRouter.Contract.GetPriceFromId(&_ILBRouter.CallOpts, LBPair, id)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address LBPair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBRouter *ILBRouterCaller) GetSwapIn(opts *bind.CallOpts, LBPair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getSwapIn", LBPair, amountOut, swapForY)

	outstruct := new(struct {
		AmountIn      *big.Int
		AmountOutLeft *big.Int
		Fee           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOutLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address LBPair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBRouter *ILBRouterSession) GetSwapIn(LBPair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _ILBRouter.Contract.GetSwapIn(&_ILBRouter.CallOpts, LBPair, amountOut, swapForY)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x964f987c.
//
// Solidity: function getSwapIn(address LBPair, uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBRouter *ILBRouterCallerSession) GetSwapIn(LBPair common.Address, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _ILBRouter.Contract.GetSwapIn(&_ILBRouter.CallOpts, LBPair, amountOut, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address LBPair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBRouter *ILBRouterCaller) GetSwapOut(opts *bind.CallOpts, LBPair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getSwapOut", LBPair, amountIn, swapForY)

	outstruct := new(struct {
		AmountInLeft *big.Int
		AmountOut    *big.Int
		Fee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInLeft = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address LBPair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBRouter *ILBRouterSession) GetSwapOut(LBPair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _ILBRouter.Contract.GetSwapOut(&_ILBRouter.CallOpts, LBPair, amountIn, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xa0d376cf.
//
// Solidity: function getSwapOut(address LBPair, uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBRouter *ILBRouterCallerSession) GetSwapOut(LBPair common.Address, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _ILBRouter.Contract.GetSwapOut(&_ILBRouter.CallOpts, LBPair, amountIn, swapForY)
}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address)
func (_ILBRouter *ILBRouterCaller) GetV1Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getV1Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address)
func (_ILBRouter *ILBRouterSession) GetV1Factory() (common.Address, error) {
	return _ILBRouter.Contract.GetV1Factory(&_ILBRouter.CallOpts)
}

// GetV1Factory is a free data retrieval call binding the contract method 0xbb558a9f.
//
// Solidity: function getV1Factory() view returns(address)
func (_ILBRouter *ILBRouterCallerSession) GetV1Factory() (common.Address, error) {
	return _ILBRouter.Contract.GetV1Factory(&_ILBRouter.CallOpts)
}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address)
func (_ILBRouter *ILBRouterCaller) GetWNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBRouter.contract.Call(opts, &out, "getWNATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address)
func (_ILBRouter *ILBRouterSession) GetWNATIVE() (common.Address, error) {
	return _ILBRouter.Contract.GetWNATIVE(&_ILBRouter.CallOpts)
}

// GetWNATIVE is a free data retrieval call binding the contract method 0x6c9c0078.
//
// Solidity: function getWNATIVE() view returns(address)
func (_ILBRouter *ILBRouterCallerSession) GetWNATIVE() (common.Address, error) {
	return _ILBRouter.Contract.GetWNATIVE(&_ILBRouter.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterTransactor) AddLiquidity(opts *bind.TransactOpts, liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "addLiquidity", liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterSession) AddLiquidity(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.Contract.AddLiquidity(&_ILBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xa3c7271a.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterTransactorSession) AddLiquidity(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.Contract.AddLiquidity(&_ILBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterTransactor) AddLiquidityNATIVE(opts *bind.TransactOpts, liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "addLiquidityNATIVE", liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterSession) AddLiquidityNATIVE(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.Contract.AddLiquidityNATIVE(&_ILBRouter.TransactOpts, liquidityParameters)
}

// AddLiquidityNATIVE is a paid mutator transaction binding the contract method 0x8efc2b2c.
//
// Solidity: function addLiquidityNATIVE((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,address,uint256) liquidityParameters) payable returns(uint256 amountXAdded, uint256 amountYAdded, uint256 amountXLeft, uint256 amountYLeft, uint256[] depositIds, uint256[] liquidityMinted)
func (_ILBRouter *ILBRouterTransactorSession) AddLiquidityNATIVE(liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _ILBRouter.Contract.AddLiquidityNATIVE(&_ILBRouter.TransactOpts, liquidityParameters)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_ILBRouter *ILBRouterTransactor) CreateLBPair(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "createLBPair", tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_ILBRouter *ILBRouterSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _ILBRouter.Contract.CreateLBPair(&_ILBRouter.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_ILBRouter *ILBRouterTransactorSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _ILBRouter.Contract.CreateLBPair(&_ILBRouter.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_ILBRouter *ILBRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "removeLiquidity", tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_ILBRouter *ILBRouterSession) RemoveLiquidity(tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.RemoveLiquidity(&_ILBRouter.TransactOpts, tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address tokenX, address tokenY, uint16 binStep, uint256 amountXMin, uint256 amountYMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_ILBRouter *ILBRouterTransactorSession) RemoveLiquidity(tokenX common.Address, tokenY common.Address, binStep uint16, amountXMin *big.Int, amountYMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.RemoveLiquidity(&_ILBRouter.TransactOpts, tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_ILBRouter *ILBRouterTransactor) RemoveLiquidityNATIVE(opts *bind.TransactOpts, token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "removeLiquidityNATIVE", token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_ILBRouter *ILBRouterSession) RemoveLiquidityNATIVE(token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.RemoveLiquidityNATIVE(&_ILBRouter.TransactOpts, token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
}

// RemoveLiquidityNATIVE is a paid mutator transaction binding the contract method 0x81c2fdfb.
//
// Solidity: function removeLiquidityNATIVE(address token, uint16 binStep, uint256 amountTokenMin, uint256 amountNATIVEMin, uint256[] ids, uint256[] amounts, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountNATIVE)
func (_ILBRouter *ILBRouterTransactorSession) RemoveLiquidityNATIVE(token common.Address, binStep uint16, amountTokenMin *big.Int, amountNATIVEMin *big.Int, ids []*big.Int, amounts []*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.RemoveLiquidityNATIVE(&_ILBRouter.TransactOpts, token, binStep, amountTokenMin, amountNATIVEMin, ids, amounts, to, deadline)
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

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_ILBRouter *ILBRouterTransactor) Sweep(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "sweep", token, to, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_ILBRouter *ILBRouterSession) Sweep(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.Sweep(&_ILBRouter.TransactOpts, token, to, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address token, address to, uint256 amount) returns()
func (_ILBRouter *ILBRouterTransactorSession) Sweep(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.Sweep(&_ILBRouter.TransactOpts, token, to, amount)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_ILBRouter *ILBRouterTransactor) SweepLBToken(opts *bind.TransactOpts, _lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ILBRouter.contract.Transact(opts, "sweepLBToken", _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_ILBRouter *ILBRouterSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SweepLBToken(&_ILBRouter.TransactOpts, _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_ILBRouter *ILBRouterTransactorSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ILBRouter.Contract.SweepLBToken(&_ILBRouter.TransactOpts, _lbToken, _to, _ids, _amounts)
}
