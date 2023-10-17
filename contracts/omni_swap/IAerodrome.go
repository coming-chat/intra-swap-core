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

// IAerodromeRoute is an auto generated low-level Go binding around an user-defined struct.
type IAerodromeRoute struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}

// IAerodromeZap is an auto generated low-level Go binding around an user-defined struct.
type IAerodromeZap struct {
	TokenA        common.Address
	TokenB        common.Address
	Stable        bool
	Factory       common.Address
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}

// IAerodromeMetaData contains all meta data concerning the IAerodrome contract.
var IAerodromeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountADesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountAOptimal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountBDesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmountInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolFactoryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"UNSAFE_swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapInParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapOutParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"poolFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"quoteStableLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIAerodrome.Zap\",\"name\":\"zapInPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stake\",\"type\":\"bool\"}],\"name\":\"zapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIAerodrome.Zap\",\"name\":\"zapOutPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIAerodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"zapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAerodromeABI is the input ABI used to generate the binding from.
// Deprecated: Use IAerodromeMetaData.ABI instead.
var IAerodromeABI = IAerodromeMetaData.ABI

// IAerodrome is an auto generated Go binding around an Ethereum contract.
type IAerodrome struct {
	IAerodromeCaller     // Read-only binding to the contract
	IAerodromeTransactor // Write-only binding to the contract
	IAerodromeFilterer   // Log filterer for contract events
}

// IAerodromeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAerodromeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAerodromeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAerodromeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAerodromeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAerodromeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAerodromeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAerodromeSession struct {
	Contract     *IAerodrome       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAerodromeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAerodromeCallerSession struct {
	Contract *IAerodromeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IAerodromeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAerodromeTransactorSession struct {
	Contract     *IAerodromeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IAerodromeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAerodromeRaw struct {
	Contract *IAerodrome // Generic contract binding to access the raw methods on
}

// IAerodromeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAerodromeCallerRaw struct {
	Contract *IAerodromeCaller // Generic read-only contract binding to access the raw methods on
}

// IAerodromeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAerodromeTransactorRaw struct {
	Contract *IAerodromeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAerodrome creates a new instance of IAerodrome, bound to a specific deployed contract.
func NewIAerodrome(address common.Address, backend bind.ContractBackend) (*IAerodrome, error) {
	contract, err := bindIAerodrome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAerodrome{IAerodromeCaller: IAerodromeCaller{contract: contract}, IAerodromeTransactor: IAerodromeTransactor{contract: contract}, IAerodromeFilterer: IAerodromeFilterer{contract: contract}}, nil
}

// NewIAerodromeCaller creates a new read-only instance of IAerodrome, bound to a specific deployed contract.
func NewIAerodromeCaller(address common.Address, caller bind.ContractCaller) (*IAerodromeCaller, error) {
	contract, err := bindIAerodrome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAerodromeCaller{contract: contract}, nil
}

// NewIAerodromeTransactor creates a new write-only instance of IAerodrome, bound to a specific deployed contract.
func NewIAerodromeTransactor(address common.Address, transactor bind.ContractTransactor) (*IAerodromeTransactor, error) {
	contract, err := bindIAerodrome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAerodromeTransactor{contract: contract}, nil
}

// NewIAerodromeFilterer creates a new log filterer instance of IAerodrome, bound to a specific deployed contract.
func NewIAerodromeFilterer(address common.Address, filterer bind.ContractFilterer) (*IAerodromeFilterer, error) {
	contract, err := bindIAerodrome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAerodromeFilterer{contract: contract}, nil
}

// bindIAerodrome binds a generic wrapper to an already deployed contract.
func bindIAerodrome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAerodromeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAerodrome *IAerodromeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAerodrome.Contract.IAerodromeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAerodrome *IAerodromeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAerodrome.Contract.IAerodromeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAerodrome *IAerodromeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAerodrome.Contract.IAerodromeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAerodrome *IAerodromeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAerodrome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAerodrome *IAerodromeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAerodrome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAerodrome *IAerodromeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAerodrome.Contract.contract.Transact(opts, method, params...)
}

// ETHER is a free data retrieval call binding the contract method 0x42cb1fbc.
//
// Solidity: function ETHER() view returns(address)
func (_IAerodrome *IAerodromeCaller) ETHER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "ETHER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHER is a free data retrieval call binding the contract method 0x42cb1fbc.
//
// Solidity: function ETHER() view returns(address)
func (_IAerodrome *IAerodromeSession) ETHER() (common.Address, error) {
	return _IAerodrome.Contract.ETHER(&_IAerodrome.CallOpts)
}

// ETHER is a free data retrieval call binding the contract method 0x42cb1fbc.
//
// Solidity: function ETHER() view returns(address)
func (_IAerodrome *IAerodromeCallerSession) ETHER() (common.Address, error) {
	return _IAerodrome.Contract.ETHER(&_IAerodrome.CallOpts)
}

// DefaultFactory is a free data retrieval call binding the contract method 0xd4b6846d.
//
// Solidity: function defaultFactory() view returns(address)
func (_IAerodrome *IAerodromeCaller) DefaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "defaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultFactory is a free data retrieval call binding the contract method 0xd4b6846d.
//
// Solidity: function defaultFactory() view returns(address)
func (_IAerodrome *IAerodromeSession) DefaultFactory() (common.Address, error) {
	return _IAerodrome.Contract.DefaultFactory(&_IAerodrome.CallOpts)
}

// DefaultFactory is a free data retrieval call binding the contract method 0xd4b6846d.
//
// Solidity: function defaultFactory() view returns(address)
func (_IAerodrome *IAerodromeCallerSession) DefaultFactory() (common.Address, error) {
	return _IAerodrome.Contract.DefaultFactory(&_IAerodrome.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_IAerodrome *IAerodromeCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "factoryRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_IAerodrome *IAerodromeSession) FactoryRegistry() (common.Address, error) {
	return _IAerodrome.Contract.FactoryRegistry(&_IAerodrome.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_IAerodrome *IAerodromeCallerSession) FactoryRegistry() (common.Address, error) {
	return _IAerodrome.Contract.FactoryRegistry(&_IAerodrome.CallOpts)
}

// GenerateZapInParams is a free data retrieval call binding the contract method 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeCaller) GenerateZapInParams(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "generateZapInParams", tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)

	outstruct := new(struct {
		AmountOutMinA *big.Int
		AmountOutMinB *big.Int
		AmountAMin    *big.Int
		AmountBMin    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOutMinA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOutMinB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountAMin = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmountBMin = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GenerateZapInParams is a free data retrieval call binding the contract method 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeSession) GenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IAerodrome.Contract.GenerateZapInParams(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)
}

// GenerateZapInParams is a free data retrieval call binding the contract method 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeCallerSession) GenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IAerodrome.Contract.GenerateZapInParams(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)
}

// GenerateZapOutParams is a free data retrieval call binding the contract method 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeCaller) GenerateZapOutParams(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "generateZapOutParams", tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)

	outstruct := new(struct {
		AmountOutMinA *big.Int
		AmountOutMinB *big.Int
		AmountAMin    *big.Int
		AmountBMin    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOutMinA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOutMinB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountAMin = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmountBMin = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GenerateZapOutParams is a free data retrieval call binding the contract method 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeSession) GenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IAerodrome.Contract.GenerateZapOutParams(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)
}

// GenerateZapOutParams is a free data retrieval call binding the contract method 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IAerodrome *IAerodromeCallerSession) GenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IAerodrome.Contract.GenerateZapOutParams(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IAerodrome *IAerodromeCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, routes []IAerodromeRoute) ([]*big.Int, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "getAmountsOut", amountIn, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IAerodrome *IAerodromeSession) GetAmountsOut(amountIn *big.Int, routes []IAerodromeRoute) ([]*big.Int, error) {
	return _IAerodrome.Contract.GetAmountsOut(&_IAerodrome.CallOpts, amountIn, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IAerodrome *IAerodromeCallerSession) GetAmountsOut(amountIn *big.Int, routes []IAerodromeRoute) ([]*big.Int, error) {
	return _IAerodrome.Contract.GetAmountsOut(&_IAerodrome.CallOpts, amountIn, routes)
}

// GetReserves is a free data retrieval call binding the contract method 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (_IAerodrome *IAerodromeCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable, _factory)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (_IAerodrome *IAerodromeSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _IAerodrome.Contract.GetReserves(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// GetReserves is a free data retrieval call binding the contract method 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (_IAerodrome *IAerodromeCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _IAerodrome.Contract.GetReserves(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IAerodrome *IAerodromeCaller) PoolFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "poolFor", tokenA, tokenB, stable, _factory)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IAerodrome *IAerodromeSession) PoolFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IAerodrome.Contract.PoolFor(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IAerodrome *IAerodromeCallerSession) PoolFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IAerodrome.Contract.PoolFor(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeCaller) QuoteAddLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "quoteAddLiquidity", tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)

	outstruct := new(struct {
		AmountA   *big.Int
		AmountB   *big.Int
		Liquidity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _IAerodrome.Contract.QuoteAddLiquidity(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeCallerSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _IAerodrome.Contract.QuoteAddLiquidity(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeCaller) QuoteRemoveLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "quoteRemoveLiquidity", tokenA, tokenB, stable, _factory, liquidity)

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _IAerodrome.Contract.QuoteRemoveLiquidity(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeCallerSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _IAerodrome.Contract.QuoteRemoveLiquidity(&_IAerodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity)
}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IAerodrome *IAerodromeCaller) QuoteStableLiquidityRatio(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "quoteStableLiquidityRatio", tokenA, tokenB, factory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IAerodrome *IAerodromeSession) QuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	return _IAerodrome.Contract.QuoteStableLiquidityRatio(&_IAerodrome.CallOpts, tokenA, tokenB, factory)
}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IAerodrome *IAerodromeCallerSession) QuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	return _IAerodrome.Contract.QuoteStableLiquidityRatio(&_IAerodrome.CallOpts, tokenA, tokenB, factory)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_IAerodrome *IAerodromeCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_IAerodrome *IAerodromeSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _IAerodrome.Contract.SortTokens(&_IAerodrome.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_IAerodrome *IAerodromeCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _IAerodrome.Contract.SortTokens(&_IAerodrome.CallOpts, tokenA, tokenB)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IAerodrome *IAerodromeCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAerodrome.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IAerodrome *IAerodromeSession) Voter() (common.Address, error) {
	return _IAerodrome.Contract.Voter(&_IAerodrome.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IAerodrome *IAerodromeCallerSession) Voter() (common.Address, error) {
	return _IAerodrome.Contract.Voter(&_IAerodrome.CallOpts)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IAerodrome *IAerodromeTransactor) UNSAFESwapExactTokensForTokens(opts *bind.TransactOpts, amounts []*big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IAerodrome *IAerodromeSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.UNSAFESwapExactTokensForTokens(&_IAerodrome.TransactOpts, amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IAerodrome *IAerodromeTransactorSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.UNSAFESwapExactTokensForTokens(&_IAerodrome.TransactOpts, amounts, routes, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.AddLiquidity(&_IAerodrome.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IAerodrome *IAerodromeTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.AddLiquidity(&_IAerodrome.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IAerodrome *IAerodromeTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IAerodrome *IAerodromeSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.AddLiquidityETH(&_IAerodrome.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IAerodrome *IAerodromeTransactorSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.AddLiquidityETH(&_IAerodrome.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidity(&_IAerodrome.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAerodrome *IAerodromeTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidity(&_IAerodrome.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IAerodrome *IAerodromeTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IAerodrome *IAerodromeSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidityETH(&_IAerodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IAerodrome *IAerodromeTransactorSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidityETH(&_IAerodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IAerodrome *IAerodromeTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IAerodrome *IAerodromeSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IAerodrome *IAerodromeTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IAerodrome *IAerodromeSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactETHForTokens(&_IAerodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactETHForTokens(&_IAerodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IAerodrome *IAerodromeTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IAerodrome *IAerodromeSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IAerodrome *IAerodromeTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForETH(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForETH(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAerodrome *IAerodromeTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IAerodrome *IAerodromeTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IAerodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAerodrome.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IAerodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IAerodrome *IAerodromeTransactor) ZapIn(opts *bind.TransactOpts, tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "zapIn", tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IAerodrome *IAerodromeSession) ZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IAerodrome.Contract.ZapIn(&_IAerodrome.TransactOpts, tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IAerodrome *IAerodromeTransactorSession) ZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IAerodrome.Contract.ZapIn(&_IAerodrome.TransactOpts, tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IAerodrome *IAerodromeTransactor) ZapOut(opts *bind.TransactOpts, tokenOut common.Address, liquidity *big.Int, zapOutPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (*types.Transaction, error) {
	return _IAerodrome.contract.Transact(opts, "zapOut", tokenOut, liquidity, zapOutPool, routesA, routesB)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IAerodrome *IAerodromeSession) ZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (*types.Transaction, error) {
	return _IAerodrome.Contract.ZapOut(&_IAerodrome.TransactOpts, tokenOut, liquidity, zapOutPool, routesA, routesB)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IAerodrome *IAerodromeTransactorSession) ZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IAerodromeZap, routesA []IAerodromeRoute, routesB []IAerodromeRoute) (*types.Transaction, error) {
	return _IAerodrome.Contract.ZapOut(&_IAerodrome.TransactOpts, tokenOut, liquidity, zapOutPool, routesA, routesB)
}
