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

// IVelodromeRoute is an auto generated low-level Go binding around an user-defined struct.
type IVelodromeRoute struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}

// IVelodromeZap is an auto generated low-level Go binding around an user-defined struct.
type IVelodromeZap struct {
	TokenA        common.Address
	TokenB        common.Address
	Stable        bool
	Factory       common.Address
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}

// IVelodromeMetaData contains all meta data concerning the IVelodrome contract.
var IVelodromeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ConversionFromV2ToV1VeloProhibited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountADesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountAOptimal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountBDesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmountInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolFactoryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"UNSAFE_swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapInParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapOutParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"poolFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"quoteStableLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIVelodrome.Zap\",\"name\":\"zapInPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stake\",\"type\":\"bool\"}],\"name\":\"zapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIVelodrome.Zap\",\"name\":\"zapOutPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodrome.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"zapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IVelodromeABI is the input ABI used to generate the binding from.
// Deprecated: Use IVelodromeMetaData.ABI instead.
var IVelodromeABI = IVelodromeMetaData.ABI

// IVelodrome is an auto generated Go binding around an Ethereum contract.
type IVelodrome struct {
	IVelodromeCaller     // Read-only binding to the contract
	IVelodromeTransactor // Write-only binding to the contract
	IVelodromeFilterer   // Log filterer for contract events
}

// IVelodromeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVelodromeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVelodromeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVelodromeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVelodromeSession struct {
	Contract     *IVelodrome       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVelodromeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVelodromeCallerSession struct {
	Contract *IVelodromeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IVelodromeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVelodromeTransactorSession struct {
	Contract     *IVelodromeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IVelodromeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVelodromeRaw struct {
	Contract *IVelodrome // Generic contract binding to access the raw methods on
}

// IVelodromeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVelodromeCallerRaw struct {
	Contract *IVelodromeCaller // Generic read-only contract binding to access the raw methods on
}

// IVelodromeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVelodromeTransactorRaw struct {
	Contract *IVelodromeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVelodrome creates a new instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodrome(address common.Address, backend bind.ContractBackend) (*IVelodrome, error) {
	contract, err := bindIVelodrome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVelodrome{IVelodromeCaller: IVelodromeCaller{contract: contract}, IVelodromeTransactor: IVelodromeTransactor{contract: contract}, IVelodromeFilterer: IVelodromeFilterer{contract: contract}}, nil
}

// NewIVelodromeCaller creates a new read-only instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeCaller(address common.Address, caller bind.ContractCaller) (*IVelodromeCaller, error) {
	contract, err := bindIVelodrome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVelodromeCaller{contract: contract}, nil
}

// NewIVelodromeTransactor creates a new write-only instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeTransactor(address common.Address, transactor bind.ContractTransactor) (*IVelodromeTransactor, error) {
	contract, err := bindIVelodrome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVelodromeTransactor{contract: contract}, nil
}

// NewIVelodromeFilterer creates a new log filterer instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeFilterer(address common.Address, filterer bind.ContractFilterer) (*IVelodromeFilterer, error) {
	contract, err := bindIVelodrome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVelodromeFilterer{contract: contract}, nil
}

// bindIVelodrome binds a generic wrapper to an already deployed contract.
func bindIVelodrome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVelodromeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVelodrome *IVelodromeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVelodrome.Contract.IVelodromeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVelodrome *IVelodromeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVelodrome.Contract.IVelodromeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVelodrome *IVelodromeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVelodrome.Contract.IVelodromeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVelodrome *IVelodromeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVelodrome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVelodrome *IVelodromeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVelodrome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVelodrome *IVelodromeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVelodrome.Contract.contract.Transact(opts, method, params...)
}

// GenerateZapInParams is a free data retrieval call binding the contract method 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IVelodrome *IVelodromeCaller) GenerateZapInParams(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "generateZapInParams", tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)

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
func (_IVelodrome *IVelodromeSession) GenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IVelodrome.Contract.GenerateZapInParams(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)
}

// GenerateZapInParams is a free data retrieval call binding the contract method 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IVelodrome *IVelodromeCallerSession) GenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IVelodrome.Contract.GenerateZapInParams(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, amountInA, amountInB, routesA, routesB)
}

// GenerateZapOutParams is a free data retrieval call binding the contract method 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IVelodrome *IVelodromeCaller) GenerateZapOutParams(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "generateZapOutParams", tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)

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
func (_IVelodrome *IVelodromeSession) GenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IVelodrome.Contract.GenerateZapOutParams(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)
}

// GenerateZapOutParams is a free data retrieval call binding the contract method 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (_IVelodrome *IVelodromeCallerSession) GenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}, error) {
	return _IVelodrome.Contract.GenerateZapOutParams(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity, routesA, routesB)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IVelodrome *IVelodromeCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, routes []IVelodromeRoute) ([]*big.Int, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "getAmountsOut", amountIn, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IVelodrome *IVelodromeSession) GetAmountsOut(amountIn *big.Int, routes []IVelodromeRoute) ([]*big.Int, error) {
	return _IVelodrome.Contract.GetAmountsOut(&_IVelodrome.CallOpts, amountIn, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (_IVelodrome *IVelodromeCallerSession) GetAmountsOut(amountIn *big.Int, routes []IVelodromeRoute) ([]*big.Int, error) {
	return _IVelodrome.Contract.GetAmountsOut(&_IVelodrome.CallOpts, amountIn, routes)
}

// GetReserves is a free data retrieval call binding the contract method 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (_IVelodrome *IVelodromeCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable, _factory)

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
func (_IVelodrome *IVelodromeSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _IVelodrome.Contract.GetReserves(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// GetReserves is a free data retrieval call binding the contract method 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (_IVelodrome *IVelodromeCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _IVelodrome.Contract.GetReserves(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PairFor is a free data retrieval call binding the contract method 0xe4ea9d74.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable, _factory)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0xe4ea9d74.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IVelodrome.Contract.PairFor(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PairFor is a free data retrieval call binding the contract method 0xe4ea9d74.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeCallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IVelodrome.Contract.PairFor(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeCaller) PoolFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "poolFor", tokenA, tokenB, stable, _factory)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeSession) PoolFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IVelodrome.Contract.PoolFor(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// PoolFor is a free data retrieval call binding the contract method 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (_IVelodrome *IVelodromeCallerSession) PoolFor(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address) (common.Address, error) {
	return _IVelodrome.Contract.PoolFor(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IVelodrome *IVelodromeCaller) QuoteAddLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "quoteAddLiquidity", tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)

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
func (_IVelodrome *IVelodromeSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _IVelodrome.Contract.QuoteAddLiquidity(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IVelodrome *IVelodromeCallerSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _IVelodrome.Contract.QuoteAddLiquidity(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, amountADesired, amountBDesired)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_IVelodrome *IVelodromeCaller) QuoteRemoveLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "quoteRemoveLiquidity", tokenA, tokenB, stable, _factory, liquidity)

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
func (_IVelodrome *IVelodromeSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _IVelodrome.Contract.QuoteRemoveLiquidity(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_IVelodrome *IVelodromeCallerSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, _factory common.Address, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _IVelodrome.Contract.QuoteRemoveLiquidity(&_IVelodrome.CallOpts, tokenA, tokenB, stable, _factory, liquidity)
}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IVelodrome *IVelodromeCaller) QuoteStableLiquidityRatio(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "quoteStableLiquidityRatio", tokenA, tokenB, factory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IVelodrome *IVelodromeSession) QuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	return _IVelodrome.Contract.QuoteStableLiquidityRatio(&_IVelodrome.CallOpts, tokenA, tokenB, factory)
}

// QuoteStableLiquidityRatio is a free data retrieval call binding the contract method 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address factory) view returns(uint256 ratio)
func (_IVelodrome *IVelodromeCallerSession) QuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) (*big.Int, error) {
	return _IVelodrome.Contract.QuoteStableLiquidityRatio(&_IVelodrome.CallOpts, tokenA, tokenB, factory)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_IVelodrome *IVelodromeCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

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
func (_IVelodrome *IVelodromeSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _IVelodrome.Contract.SortTokens(&_IVelodrome.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_IVelodrome *IVelodromeCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _IVelodrome.Contract.SortTokens(&_IVelodrome.CallOpts, tokenA, tokenB)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IVelodrome *IVelodromeTransactor) UNSAFESwapExactTokensForTokens(opts *bind.TransactOpts, amounts []*big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IVelodrome *IVelodromeSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.UNSAFESwapExactTokensForTokens(&_IVelodrome.TransactOpts, amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (_IVelodrome *IVelodromeTransactorSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.UNSAFESwapExactTokensForTokens(&_IVelodrome.TransactOpts, amounts, routes, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IVelodrome *IVelodromeTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IVelodrome *IVelodromeSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.AddLiquidity(&_IVelodrome.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IVelodrome *IVelodromeTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.AddLiquidity(&_IVelodrome.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IVelodrome *IVelodromeTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IVelodrome *IVelodromeSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.AddLiquidityETH(&_IVelodrome.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IVelodrome *IVelodromeTransactorSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.AddLiquidityETH(&_IVelodrome.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IVelodrome *IVelodromeTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IVelodrome *IVelodromeSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidity(&_IVelodrome.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IVelodrome *IVelodromeTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidity(&_IVelodrome.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IVelodrome *IVelodromeTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IVelodrome *IVelodromeSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidityETH(&_IVelodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IVelodrome *IVelodromeTransactorSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidityETH(&_IVelodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IVelodrome *IVelodromeTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IVelodrome *IVelodromeSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IVelodrome *IVelodromeTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IVelodrome *IVelodromeSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactETHForTokens(&_IVelodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactETHForTokens(&_IVelodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IVelodrome *IVelodromeTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IVelodrome *IVelodromeSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3da5acba.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (_IVelodrome *IVelodromeTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForETH(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForETH(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x12bc3aca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_IVelodrome *IVelodromeTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x88cd821e.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (_IVelodrome *IVelodromeTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IVelodromeRoute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IVelodrome.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IVelodrome.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IVelodrome *IVelodromeTransactor) ZapIn(opts *bind.TransactOpts, tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "zapIn", tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IVelodrome *IVelodromeSession) ZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IVelodrome.Contract.ZapIn(&_IVelodrome.TransactOpts, tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapIn is a paid mutator transaction binding the contract method 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (_IVelodrome *IVelodromeTransactorSession) ZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute, to common.Address, stake bool) (*types.Transaction, error) {
	return _IVelodrome.Contract.ZapIn(&_IVelodrome.TransactOpts, tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IVelodrome *IVelodromeTransactor) ZapOut(opts *bind.TransactOpts, tokenOut common.Address, liquidity *big.Int, zapOutPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "zapOut", tokenOut, liquidity, zapOutPool, routesA, routesB)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IVelodrome *IVelodromeSession) ZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (*types.Transaction, error) {
	return _IVelodrome.Contract.ZapOut(&_IVelodrome.TransactOpts, tokenOut, liquidity, zapOutPool, routesA, routesB)
}

// ZapOut is a paid mutator transaction binding the contract method 0xa81b9159.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (_IVelodrome *IVelodromeTransactorSession) ZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IVelodromeZap, routesA []IVelodromeRoute, routesB []IVelodromeRoute) (*types.Transaction, error) {
	return _IVelodrome.Contract.ZapOut(&_IVelodrome.TransactOpts, tokenOut, liquidity, zapOutPool, routesA, routesB)
}
