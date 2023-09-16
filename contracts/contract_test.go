package contracts

import (
	"context"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestQuoter2(t *testing.T) {
	abi, err := omni_swap.IQuoterV2MetaData.GetAbi()
	if err != nil {
		return
	}
	client, err := ethclient.Dial("https://endpoints.omniatech.io/v1/base/mainnet/public")
	if err != nil {
		return
	}
	data, err := abi.Pack("quoteExactInput", []byte{
		217,
		170,
		236,
		134,
		182,
		93,
		134,
		246,
		167,
		181,
		177,
		176,
		196,
		47,
		250,
		83,
		23,
		16,
		182,
		202,
		0,
		0,
		0,
		66,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		6,
	}, big.NewInt(1500000000))
	if err != nil {
		return
	}
	result, err := client.CallContract(context.TODO(), ethereum.CallMsg{
		To:   &util.UniswapV3Quoter,
		Data: data,
	}, nil)
	if err != nil {
		return
	}

	t.Log(result)
}
