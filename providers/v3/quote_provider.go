package v3

import (
	"context"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/provider"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type AmountQuote struct {
	Amount                      *entities.CurrencyAmount
	Quote                       *big.Int
	QuoteList                   []*big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

type RouteWithQuotes struct {
	Route       *base_entities.MRoute
	AmountQuote []AmountQuote
}

type QuoteProvider interface {
	GetQuotesManyExactIn(
		amountIns []*entities.CurrencyAmount,
		routes []*base_entities.MRoute,
		providerConfig *provider.Config,
	) ([]RouteWithQuotes, int64, error)

	GetQuotesManyExactOut(
		amountOuts []*entities.CurrencyAmount,
		routes []*base_entities.MRoute,
		providerConfig *provider.Config,
	) ([]RouteWithQuotes, int64, error)
}

type BlockNumberConfig struct {
	BaseBlockOffset uint64
	Rollback        bool
}

type QuoteData struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

type BaseQuoteProvider struct {
	QuoterAddress      common.Address
	ChainId            base_entities.ChainId
	Provider           rpc.BaseProvider
	MultiCall2Provider rpc.MultiCallProviderCore
	BlockNumberConfig  BlockNumberConfig
	MultiCallConfig    *rpc.MultiCallConfig
	Ctx                context.Context
}

func NewBaseQuoteProvider(
	ctx context.Context,
	address common.Address,
	chainId base_entities.ChainId,
	baseProvider rpc.BaseProvider,
	multiCallCore rpc.MultiCallProviderCore,
	blockNumberConfig *BlockNumberConfig,
	config *rpc.MultiCallConfig,
) *BaseQuoteProvider {
	quoteProvider := &BaseQuoteProvider{
		ChainId:            chainId,
		Ctx:                ctx,
		QuoterAddress:      address,
		Provider:           baseProvider,
		MultiCall2Provider: multiCallCore,
	}
	if config == nil {
		config = &rpc.MultiCallConfig{
			RetryOptions: rpc.RetryOptions{
				Retries:    2,
				MinTimeout: 25,
				MaxTimeout: 250,
			},
			BatchParams: rpc.BatchParams{
				MultiCallChunk:  300,
				GasLimitPerCall: 1_000_000,
				MinSuccessRate:  0.2,
			},
		}
	}
	if blockNumberConfig == nil {
		blockNumberConfig = &BlockNumberConfig{
			BaseBlockOffset: 0,
			Rollback:        false,
		}
	}
	quoteProvider.BlockNumberConfig = *blockNumberConfig
	quoteProvider.MultiCallConfig = config
	return quoteProvider
}

func (b *BaseQuoteProvider) GetQuotesManyExactIn(
	amountIns []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	return b.getQuotesManyData(amountIns, routes, QuoteExactInput, providerConfig)
}

func (b *BaseQuoteProvider) GetQuotesManyExactOut(
	amountOuts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	return b.getQuotesManyData(amountOuts, routes, QuoteExactOutput, providerConfig)
}

func (b *BaseQuoteProvider) getQuotesManyData(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	functionName functionName,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	originalBlockNumber, err := b.Provider.Rpc.BlockNumber(b.Ctx)
	if err != nil {
		return nil, 0, err
	}
	if providerConfig == nil {
		providerConfig = &provider.Config{
			BlockNumber: originalBlockNumber + b.BlockNumberConfig.BaseBlockOffset,
		}
	}
	var (
		maxHop = 0
		result []RouteWithQuotes
	)
	for _, route := range routes {
		if len(route.Pools) > maxHop {
			maxHop = len(route.Pools)
		}
		result = append(result, RouteWithQuotes{
			Route: route,
		})
	}

	syncAmounts := make([]*entities.CurrencyAmount, len(amounts))
	copy(syncAmounts, amounts)

	for i := 0; i < maxHop; i++ {
		var multiCallParams []rpc.MultiCallSingleParam
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for _, a := range syncAmounts {
				call, err := QuoteMultiCall(route.Pools[i], functionName, a)
				if err != nil {
					return nil, 0, err
				}
				multiCallParams = append(multiCallParams, call)
				if i == 0 {
					result[ri].AmountQuote = append(result[ri].AmountQuote, AmountQuote{
						Amount:      a,
						GasEstimate: big.NewInt(0),
					})
				}
			}
		}
		callResult, err := rpc.ConcurrentMultiCall[QuoteData](b.MultiCall2Provider, multiCallParams, b.MultiCallConfig)
		if err != nil {
			return nil, 0, err
		}

		callRouteIndex := 0
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for ra, _ := range syncAmounts {
				quoteData := callResult.ReturnData[callRouteIndex*len(syncAmounts)+ra]
				syncAmounts[ra] = entities.FromRawAmount(syncAmounts[ra].Currency, quoteData.Data.AmountOut)
				result[ri].AmountQuote[ra].Quote = quoteData.Data.AmountOut
				result[ri].AmountQuote[ra].QuoteList = append(result[ri].AmountQuote[ra].QuoteList, quoteData.Data.AmountOut)
				result[ri].AmountQuote[ra].SqrtPriceX96AfterList = append(result[ri].AmountQuote[ra].SqrtPriceX96AfterList, quoteData.Data.SqrtPriceX96AfterList...)
				result[ri].AmountQuote[ra].InitializedTicksCrossedList = append(result[ri].AmountQuote[ra].InitializedTicksCrossedList, quoteData.Data.InitializedTicksCrossedList...)
				result[ri].AmountQuote[ra].GasEstimate.Add(result[ri].AmountQuote[ra].GasEstimate, quoteData.Data.GasEstimate)
			}
			callRouteIndex++
		}
	}

	fmt.Println("aaa: " + result[0].AmountQuote[19].Quote.String())
	return result, 0, nil
}
