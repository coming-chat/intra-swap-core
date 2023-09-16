package rpc

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"testing"
)

func TestMulticall(t *testing.T) {
	client, err := ethclient.Dial("https://endpoints.omniatech.io/v1/base/mainnet/public")
	if err != nil {
		t.Fatal(err)
	}
	core, err := NewUniswapMultiCallProviderCore(base_entities.BASE, &BaseProvider{client})
	if err != nil {
		t.Fatal(err)
	}
	erc20Abi, err := contracts.Erc20MetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	result, err := NewUniswapMultiCallProvider[uint8](core).MultiCall(
		[]MultiCallSingleParam{
			{
				FunctionName:    "decimals",
				Contract:        erc20Abi,
				ContractAddress: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
			},
		},
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.ReturnData[0].Success)
	t.Log(result.ReturnData[0].ReturnData)
	t.Log(result.BlockNumber)

	//abi, err := contracts.Multicall3MetaData.GetAbi()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//var aa = []MultiCallSingleParam{
	//	{
	//		FunctionName:    "decimals",
	//		Contract:        erc20Abi,
	//		ContractAddress: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
	//	},
	//}
	//var calls []contracts.Multicall3Call
	//for _, param := range aa {
	//	callData, err := param.Contract.Pack(param.FunctionName, param.FunctionParams...)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	calls = append(calls, contracts.Multicall3Call{
	//		Target:   param.ContractAddress,
	//		CallData: callData,
	//	})
	//}
	//pack, err := abi.Pack("blockAndAggregate", calls)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//multi := common.HexToAddress(util.MultiCall3Address)
	//callData, err := client.CallContract(context.TODO(), ethereum.CallMsg{
	//	Data: pack,
	//	To:   &multi,
	//}, nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//var data MultiCallResultWithInfo[uint8]
	//err = abi.UnpackIntoInterface(&data, "blockAndAggregate", callData)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(data)
	b := float64(14) / 6
	a := math.Ceil(b)
	t.Log(a)
}
