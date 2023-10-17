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

// IGMXV1VaultMetaData contains all meta data concerning the IGMXV1Vault contract.
var IGMXV1VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allWhitelistedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelistedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"approvedRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"bufferAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"buyUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"cumulativeFundingRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"decreasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"directPoolDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"feeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMaxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMinPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getRedemptionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTargetUsdgAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"globalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"guaranteedUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasDynamicFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inManagerMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateLiquidationMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"increasePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeverageEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isLiquidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"lastFundingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFeeUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxUsdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minProfitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProfitTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBurnFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"poolAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reservedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sellUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBufferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_errorCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_error\",\"type\":\"string\"}],\"name\":\"setError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasDynamicFees\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inManagerMode\",\"type\":\"bool\"}],\"name\":\"setInManagerMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateLiquidationMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isManager\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMaxGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_redemptionBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"shortableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableTaxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"stableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"tokenToUsdMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"usdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGMXV1VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IGMXV1VaultMetaData.ABI instead.
var IGMXV1VaultABI = IGMXV1VaultMetaData.ABI

// IGMXV1Vault is an auto generated Go binding around an Ethereum contract.
type IGMXV1Vault struct {
	IGMXV1VaultCaller     // Read-only binding to the contract
	IGMXV1VaultTransactor // Write-only binding to the contract
	IGMXV1VaultFilterer   // Log filterer for contract events
}

// IGMXV1VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGMXV1VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGMXV1VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGMXV1VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMXV1VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGMXV1VaultSession struct {
	Contract     *IGMXV1Vault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGMXV1VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGMXV1VaultCallerSession struct {
	Contract *IGMXV1VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IGMXV1VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGMXV1VaultTransactorSession struct {
	Contract     *IGMXV1VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IGMXV1VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGMXV1VaultRaw struct {
	Contract *IGMXV1Vault // Generic contract binding to access the raw methods on
}

// IGMXV1VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGMXV1VaultCallerRaw struct {
	Contract *IGMXV1VaultCaller // Generic read-only contract binding to access the raw methods on
}

// IGMXV1VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGMXV1VaultTransactorRaw struct {
	Contract *IGMXV1VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGMXV1Vault creates a new instance of IGMXV1Vault, bound to a specific deployed contract.
func NewIGMXV1Vault(address common.Address, backend bind.ContractBackend) (*IGMXV1Vault, error) {
	contract, err := bindIGMXV1Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGMXV1Vault{IGMXV1VaultCaller: IGMXV1VaultCaller{contract: contract}, IGMXV1VaultTransactor: IGMXV1VaultTransactor{contract: contract}, IGMXV1VaultFilterer: IGMXV1VaultFilterer{contract: contract}}, nil
}

// NewIGMXV1VaultCaller creates a new read-only instance of IGMXV1Vault, bound to a specific deployed contract.
func NewIGMXV1VaultCaller(address common.Address, caller bind.ContractCaller) (*IGMXV1VaultCaller, error) {
	contract, err := bindIGMXV1Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1VaultCaller{contract: contract}, nil
}

// NewIGMXV1VaultTransactor creates a new write-only instance of IGMXV1Vault, bound to a specific deployed contract.
func NewIGMXV1VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IGMXV1VaultTransactor, error) {
	contract, err := bindIGMXV1Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGMXV1VaultTransactor{contract: contract}, nil
}

// NewIGMXV1VaultFilterer creates a new log filterer instance of IGMXV1Vault, bound to a specific deployed contract.
func NewIGMXV1VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IGMXV1VaultFilterer, error) {
	contract, err := bindIGMXV1Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGMXV1VaultFilterer{contract: contract}, nil
}

// bindIGMXV1Vault binds a generic wrapper to an already deployed contract.
func bindIGMXV1Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGMXV1VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Vault *IGMXV1VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Vault.Contract.IGMXV1VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Vault *IGMXV1VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.IGMXV1VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Vault *IGMXV1VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.IGMXV1VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMXV1Vault *IGMXV1VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMXV1Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMXV1Vault *IGMXV1VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMXV1Vault *IGMXV1VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.contract.Transact(opts, method, params...)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IGMXV1Vault *IGMXV1VaultSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _IGMXV1Vault.Contract.AllWhitelistedTokens(&_IGMXV1Vault.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _IGMXV1Vault.Contract.AllWhitelistedTokens(&_IGMXV1Vault.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _IGMXV1Vault.Contract.AllWhitelistedTokensLength(&_IGMXV1Vault.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _IGMXV1Vault.Contract.AllWhitelistedTokensLength(&_IGMXV1Vault.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) ApprovedRouters(opts *bind.CallOpts, _account common.Address, _router common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "approvedRouters", _account, _router)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) ApprovedRouters(_account common.Address, _router common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.ApprovedRouters(&_IGMXV1Vault.CallOpts, _account, _router)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) ApprovedRouters(_account common.Address, _router common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.ApprovedRouters(&_IGMXV1Vault.CallOpts, _account, _router)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) BufferAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "bufferAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) BufferAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.BufferAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) BufferAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.BufferAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) CumulativeFundingRates(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "cumulativeFundingRates", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) CumulativeFundingRates(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.CumulativeFundingRates(&_IGMXV1Vault.CallOpts, _token)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) CumulativeFundingRates(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.CumulativeFundingRates(&_IGMXV1Vault.CallOpts, _token)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) FeeReserves(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "feeReserves", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) FeeReserves(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.FeeReserves(&_IGMXV1Vault.CallOpts, _token)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) FeeReserves(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.FeeReserves(&_IGMXV1Vault.CallOpts, _token)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) FundingInterval() (*big.Int, error) {
	return _IGMXV1Vault.Contract.FundingInterval(&_IGMXV1Vault.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) FundingInterval() (*big.Int, error) {
	return _IGMXV1Vault.Contract.FundingInterval(&_IGMXV1Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) FundingRateFactor() (*big.Int, error) {
	return _IGMXV1Vault.Contract.FundingRateFactor(&_IGMXV1Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) FundingRateFactor() (*big.Int, error) {
	return _IGMXV1Vault.Contract.FundingRateFactor(&_IGMXV1Vault.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _IGMXV1Vault.Contract.GetDelta(&_IGMXV1Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _IGMXV1Vault.Contract.GetDelta(&_IGMXV1Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetFeeBasisPoints(&_IGMXV1Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetFeeBasisPoints(&_IGMXV1Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetMaxPrice(&_IGMXV1Vault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetMaxPrice(&_IGMXV1Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetMinPrice(&_IGMXV1Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetMinPrice(&_IGMXV1Vault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetNextFundingRate(&_IGMXV1Vault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetNextFundingRate(&_IGMXV1Vault.CallOpts, _token)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _IGMXV1Vault.Contract.GetPosition(&_IGMXV1Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _IGMXV1Vault.Contract.GetPosition(&_IGMXV1Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetRedemptionAmount(&_IGMXV1Vault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetRedemptionAmount(&_IGMXV1Vault.CallOpts, _token, _usdgAmount)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetTargetUsdgAmount(&_IGMXV1Vault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GetTargetUsdgAmount(&_IGMXV1Vault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GlobalShortAveragePrices(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "globalShortAveragePrices", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GlobalShortAveragePrices(&_IGMXV1Vault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GlobalShortAveragePrices(&_IGMXV1Vault.CallOpts, _token)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GlobalShortSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "globalShortSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GlobalShortSizes(&_IGMXV1Vault.CallOpts, _token)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GlobalShortSizes(&_IGMXV1Vault.CallOpts, _token)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultSession) Gov() (common.Address, error) {
	return _IGMXV1Vault.Contract.Gov(&_IGMXV1Vault.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) Gov() (common.Address, error) {
	return _IGMXV1Vault.Contract.Gov(&_IGMXV1Vault.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) GuaranteedUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "guaranteedUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) GuaranteedUsd(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GuaranteedUsd(&_IGMXV1Vault.CallOpts, _token)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) GuaranteedUsd(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.GuaranteedUsd(&_IGMXV1Vault.CallOpts, _token)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) HasDynamicFees() (bool, error) {
	return _IGMXV1Vault.Contract.HasDynamicFees(&_IGMXV1Vault.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) HasDynamicFees() (bool, error) {
	return _IGMXV1Vault.Contract.HasDynamicFees(&_IGMXV1Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) InManagerMode() (bool, error) {
	return _IGMXV1Vault.Contract.InManagerMode(&_IGMXV1Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) InManagerMode() (bool, error) {
	return _IGMXV1Vault.Contract.InManagerMode(&_IGMXV1Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) InPrivateLiquidationMode() (bool, error) {
	return _IGMXV1Vault.Contract.InPrivateLiquidationMode(&_IGMXV1Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _IGMXV1Vault.Contract.InPrivateLiquidationMode(&_IGMXV1Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) IsInitialized() (bool, error) {
	return _IGMXV1Vault.Contract.IsInitialized(&_IGMXV1Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) IsInitialized() (bool, error) {
	return _IGMXV1Vault.Contract.IsInitialized(&_IGMXV1Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) IsLeverageEnabled() (bool, error) {
	return _IGMXV1Vault.Contract.IsLeverageEnabled(&_IGMXV1Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) IsLeverageEnabled() (bool, error) {
	return _IGMXV1Vault.Contract.IsLeverageEnabled(&_IGMXV1Vault.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) IsLiquidator(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "isLiquidator", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) IsLiquidator(_account common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.IsLiquidator(&_IGMXV1Vault.CallOpts, _account)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) IsLiquidator(_account common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.IsLiquidator(&_IGMXV1Vault.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) IsManager(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "isManager", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) IsManager(_account common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.IsManager(&_IGMXV1Vault.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) IsManager(_account common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.IsManager(&_IGMXV1Vault.CallOpts, _account)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) IsSwapEnabled() (bool, error) {
	return _IGMXV1Vault.Contract.IsSwapEnabled(&_IGMXV1Vault.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) IsSwapEnabled() (bool, error) {
	return _IGMXV1Vault.Contract.IsSwapEnabled(&_IGMXV1Vault.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) LastFundingTimes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "lastFundingTimes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) LastFundingTimes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.LastFundingTimes(&_IGMXV1Vault.CallOpts, _token)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) LastFundingTimes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.LastFundingTimes(&_IGMXV1Vault.CallOpts, _token)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) LiquidationFeeUsd() (*big.Int, error) {
	return _IGMXV1Vault.Contract.LiquidationFeeUsd(&_IGMXV1Vault.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _IGMXV1Vault.Contract.LiquidationFeeUsd(&_IGMXV1Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MarginFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MarginFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MaxGasPrice() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxGasPrice(&_IGMXV1Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MaxGasPrice() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxGasPrice(&_IGMXV1Vault.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MaxGlobalShortSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "maxGlobalShortSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxGlobalShortSizes(&_IGMXV1Vault.CallOpts, _token)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxGlobalShortSizes(&_IGMXV1Vault.CallOpts, _token)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MaxLeverage() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxLeverage(&_IGMXV1Vault.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MaxLeverage() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxLeverage(&_IGMXV1Vault.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MaxUsdgAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "maxUsdgAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MaxUsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxUsdgAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MaxUsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MaxUsdgAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MinProfitBasisPoints(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "minProfitBasisPoints", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MinProfitBasisPoints(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MinProfitBasisPoints(&_IGMXV1Vault.CallOpts, _token)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MinProfitBasisPoints(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.MinProfitBasisPoints(&_IGMXV1Vault.CallOpts, _token)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MinProfitTime() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MinProfitTime(&_IGMXV1Vault.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MinProfitTime() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MinProfitTime(&_IGMXV1Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MintBurnFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.MintBurnFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) PoolAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "poolAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) PoolAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.PoolAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) PoolAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.PoolAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultSession) PriceFeed() (common.Address, error) {
	return _IGMXV1Vault.Contract.PriceFeed(&_IGMXV1Vault.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) PriceFeed() (common.Address, error) {
	return _IGMXV1Vault.Contract.PriceFeed(&_IGMXV1Vault.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) ReservedAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "reservedAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) ReservedAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.ReservedAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) ReservedAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.ReservedAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultSession) Router() (common.Address, error) {
	return _IGMXV1Vault.Contract.Router(&_IGMXV1Vault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) Router() (common.Address, error) {
	return _IGMXV1Vault.Contract.Router(&_IGMXV1Vault.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) ShortableTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "shortableTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) ShortableTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.ShortableTokens(&_IGMXV1Vault.CallOpts, _token)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) ShortableTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.ShortableTokens(&_IGMXV1Vault.CallOpts, _token)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) StableFundingRateFactor() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableFundingRateFactor(&_IGMXV1Vault.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableFundingRateFactor(&_IGMXV1Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableSwapFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableSwapFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) StableTaxBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableTaxBasisPoints(&_IGMXV1Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.StableTaxBasisPoints(&_IGMXV1Vault.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) StableTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "stableTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) StableTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.StableTokens(&_IGMXV1Vault.CallOpts, _token)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) StableTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.StableTokens(&_IGMXV1Vault.CallOpts, _token)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.SwapFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.SwapFeeBasisPoints(&_IGMXV1Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TaxBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.TaxBasisPoints(&_IGMXV1Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _IGMXV1Vault.Contract.TaxBasisPoints(&_IGMXV1Vault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TokenBalances(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "tokenBalances", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TokenBalances(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenBalances(&_IGMXV1Vault.CallOpts, _token)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TokenBalances(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenBalances(&_IGMXV1Vault.CallOpts, _token)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TokenDecimals(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "tokenDecimals", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TokenDecimals(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenDecimals(&_IGMXV1Vault.CallOpts, _token)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TokenDecimals(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenDecimals(&_IGMXV1Vault.CallOpts, _token)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenToUsdMin(&_IGMXV1Vault.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenToUsdMin(&_IGMXV1Vault.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TokenWeights(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "tokenWeights", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TokenWeights(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenWeights(&_IGMXV1Vault.CallOpts, _token)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TokenWeights(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.TokenWeights(&_IGMXV1Vault.CallOpts, _token)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) TotalTokenWeights() (*big.Int, error) {
	return _IGMXV1Vault.Contract.TotalTokenWeights(&_IGMXV1Vault.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _IGMXV1Vault.Contract.TotalTokenWeights(&_IGMXV1Vault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultSession) Usdg() (common.Address, error) {
	return _IGMXV1Vault.Contract.Usdg(&_IGMXV1Vault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) Usdg() (common.Address, error) {
	return _IGMXV1Vault.Contract.Usdg(&_IGMXV1Vault.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) UsdgAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "usdgAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) UsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.UsdgAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) UsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IGMXV1Vault.Contract.UsdgAmounts(&_IGMXV1Vault.CallOpts, _token)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) WhitelistedTokenCount() (*big.Int, error) {
	return _IGMXV1Vault.Contract.WhitelistedTokenCount(&_IGMXV1Vault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _IGMXV1Vault.Contract.WhitelistedTokenCount(&_IGMXV1Vault.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCaller) WhitelistedTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IGMXV1Vault.contract.Call(opts, &out, "whitelistedTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultSession) WhitelistedTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.WhitelistedTokens(&_IGMXV1Vault.CallOpts, _token)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IGMXV1Vault *IGMXV1VaultCallerSession) WhitelistedTokens(_token common.Address) (bool, error) {
	return _IGMXV1Vault.Contract.WhitelistedTokens(&_IGMXV1Vault.CallOpts, _token)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.BuyUSDG(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.BuyUSDG(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.DecreasePosition(&_IGMXV1Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.DecreasePosition(&_IGMXV1Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.DirectPoolDeposit(&_IGMXV1Vault.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.DirectPoolDeposit(&_IGMXV1Vault.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.IncreasePosition(&_IGMXV1Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.IncreasePosition(&_IGMXV1Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SellUSDG(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SellUSDG(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetBufferAmount(&_IGMXV1Vault.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetBufferAmount(&_IGMXV1Vault.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetError(&_IGMXV1Vault.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetError(&_IGMXV1Vault.TransactOpts, _errorCode, _error)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetFees(&_IGMXV1Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetFees(&_IGMXV1Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetFundingRate(&_IGMXV1Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetFundingRate(&_IGMXV1Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetInManagerMode(&_IGMXV1Vault.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetInManagerMode(&_IGMXV1Vault.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetInPrivateLiquidationMode(&_IGMXV1Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetInPrivateLiquidationMode(&_IGMXV1Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetIsLeverageEnabled(&_IGMXV1Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetIsLeverageEnabled(&_IGMXV1Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetIsSwapEnabled(&_IGMXV1Vault.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetIsSwapEnabled(&_IGMXV1Vault.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetLiquidator(&_IGMXV1Vault.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetLiquidator(&_IGMXV1Vault.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetManager(&_IGMXV1Vault.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetManager(&_IGMXV1Vault.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxGasPrice(&_IGMXV1Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxGasPrice(&_IGMXV1Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxGlobalShortSize(&_IGMXV1Vault.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxGlobalShortSize(&_IGMXV1Vault.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxLeverage(&_IGMXV1Vault.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetMaxLeverage(&_IGMXV1Vault.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetPriceFeed(&_IGMXV1Vault.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetPriceFeed(&_IGMXV1Vault.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IGMXV1Vault *IGMXV1VaultSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetTokenConfig(&_IGMXV1Vault.TransactOpts, _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.SetTokenConfig(&_IGMXV1Vault.TransactOpts, _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.Swap(&_IGMXV1Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.Swap(&_IGMXV1Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.WithdrawFees(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IGMXV1Vault *IGMXV1VaultTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IGMXV1Vault.Contract.WithdrawFees(&_IGMXV1Vault.TransactOpts, _token, _receiver)
}
