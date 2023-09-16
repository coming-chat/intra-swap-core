package v3

import (
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"math/big"
)

type OptimismGasData struct {
	L1BaseFee *big.Int
	Scalar    *big.Int
	Decimals  *big.Int
	Overhead  *big.Int
}

type ArbitrumGasData struct {
	PerL2TxFee       *big.Int
	PerL1CalldataFee *big.Int
	PerArbGasTotal   *big.Int
}

func GetOptimismL2GasData(multiCall2Provider rpc.MultiCallProviderCore, gasPriceAddress string) (OptimismGasData, error) {
	//funcNames := []string{"l1BaseFee", "scalar", "decimals", "overhead"}
	//var multiCallData []rpc.MultiCallSingleParam
	//for _, v := range funcNames {
	//	multiCallData = append(multiCallData, rpc.MultiCallSingleParam{
	//
	//	})
	//}
	//result, err := rpc.NewUniswapMultiCallProvider[*big.Int](*o.multiCall2Provider).MultiCall(
	//	rpc.MultiCallParams{
	//		MultiCallData:,
	//	},
	//)
	//if err != nil {
	//	return OptimismGasData{}
	//}
	return OptimismGasData{}, nil
}

func GetArbitrumL2GasData(multiCall2Provider rpc.MultiCallProviderCore, gasPriceAddress string) (ArbitrumGasData, error) {
	//funcNames := []string{"l1BaseFee", "scalar", "decimals", "overhead"}
	//var multiCallData []rpc.MultiCallSingleParam
	//for _, v := range funcNames {
	//	multiCallData = append(multiCallData, rpc.MultiCallSingleParam{
	//
	//	})
	//}
	//result, err := rpc.NewUniswapMultiCallProvider[*big.Int](*o.multiCall2Provider).MultiCall(
	//	rpc.MultiCallParams{
	//		MultiCallData:,
	//	},
	//)
	//if err != nil {
	//	return OptimismGasData{}
	//}
	return ArbitrumGasData{}, nil
}
