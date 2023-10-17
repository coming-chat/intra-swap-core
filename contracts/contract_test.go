package contracts

import (
	"context"
	"github.com/gkirito/go-ethereum"
	"github.com/gkirito/go-ethereum/common"
	"github.com/gkirito/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestForUniswapV2Router02(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.base.org")
	if err != nil {
		t.Fatal(err)
	}
	packed, err := IUniswapV2Router02Abi.Pack("getAmountsOut", big.NewInt(4041330), []common.Address{
		common.HexToAddress("0x9F00d101A9E06b76F388A4A5Cd1Bf8f542f02bBc"),
		common.HexToAddress("0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"),
	})
	if err != nil {
		t.Fatal(err)
	}
	contractAdd := common.HexToAddress("0xaaa3b1F1bd7BCc97fD1917c18ADE665C5D31F066")
	callResp, err := client.CallContract(context.TODO(), ethereum.CallMsg{
		From: common.Address{},
		To:   &contractAdd,
		Data: packed,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	unpack, err := IUniswapV2Router02Abi.Unpack("getAmountsOut", callResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(unpack)
}
