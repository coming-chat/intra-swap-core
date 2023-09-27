package v3

import (
	"context"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type AmountQuote struct {
	Amount                      *entities.CurrencyAmount
	Quote                       *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

type RouteWithQuotes struct {
	Route        *base_entities.MRoute
	AmountQuotes []AmountQuote
}

type QuoteProvider interface {
	GetQuotesMany(
		amounts []*entities.CurrencyAmount,
		routes []*base_entities.MRoute,
		tradeType entities.TradeType,
		providerConfig *config.Config,
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
	ChainId            base_entities.ChainId
	Provider           rpc.Provider
	MultiCall2Provider rpc.MultiCallProviderCore
	BlockNumberConfig  BlockNumberConfig
	MultiCallConfig    *rpc.MultiCallConfig
	Offline            bool
	Ctx                context.Context
}

func NewBaseQuoteProvider(
	ctx context.Context,
	chainId base_entities.ChainId,
	baseProvider rpc.Provider,
	multiCallCore rpc.MultiCallProviderCore,
	blockNumberConfig *BlockNumberConfig,
	offline bool,
	config *rpc.MultiCallConfig,
) *BaseQuoteProvider {
	quoteProvider := &BaseQuoteProvider{
		ChainId:            chainId,
		Ctx:                ctx,
		Provider:           baseProvider,
		Offline:            offline,
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
				MultiCallChunk:  1000,
				GasLimitPerCall: 1_000_000,
				MinSuccessRate:  0.2,
			},
			MaxConcurrentNum: 3,
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

func (b *BaseQuoteProvider) GetQuotesMany(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	tradeType entities.TradeType,
	providerConfig *config.Config,
) ([]RouteWithQuotes, int64, error) {
	if b.Offline {
		// tmp not support offline
	}
	return b.getQuotesManyDataOnline(amounts, routes, tradeType, providerConfig)
}

func (b *BaseQuoteProvider) getQuotesManyDataOnline(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	tradeType entities.TradeType,
	providerConfig *config.Config,
) ([]RouteWithQuotes, int64, error) {
	client, err := b.Provider.GetEthRpc(b.ChainId)
	if err != nil {
		return nil, 0, err
	}
	originalBlockNumber, err := client.BlockNumber(b.Ctx)
	if err != nil {
		return nil, 0, err
	}
	if providerConfig == nil {
		providerConfig = &config.Config{
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

	syncAmounts := make([][]*entities.CurrencyAmount, len(routes))
	for i := range syncAmounts {
		syncAmounts[i] = make([]*entities.CurrencyAmount, len(amounts))
		copy(syncAmounts[i], amounts)
	}

	for i := 0; i < maxHop; i++ {
		var multiCallParams []rpc.MultiCallSingleParam
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for _, a := range syncAmounts[ri] {
				if a.EqualTo(base_constant.ZeroFraction) {
					continue
				}
				call, err := QuoteMultiCall(route, i, tradeType, a)
				if err != nil {
					return nil, 0, err
				}
				multiCallParams = append(multiCallParams, call)
				if i == 0 {
					result[ri].AmountQuotes = append(result[ri].AmountQuotes, AmountQuote{
						Amount:      entities.FromRawAmount(a.Currency, a.Quotient()),
						GasEstimate: big.NewInt(0),
					})
				}
			}
		}
		if len(multiCallParams) == 0 {
			continue
		}
		callResult, err := rpc.ConcurrentMultiCall[QuoteData](b.ChainId, b.MultiCall2Provider, multiCallParams, b.MultiCallConfig)
		if err != nil {
			return nil, 0, err
		}
		callDataIndex := 0
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for ra, a := range syncAmounts[ri] {
				if a.EqualTo(base_constant.ZeroFraction) {
					continue
				}
				quoteData := callResult.ReturnData[callDataIndex]
				if !quoteData.Success {
					result[ri].AmountQuotes[ra].Quote = big.NewInt(0)
				} else {
					syncAmounts[ri][ra] = entities.FromRawAmount(syncAmounts[ri][ra].Currency, quoteData.Data.AmountOut)
					result[ri].AmountQuotes[ra].Quote = quoteData.Data.AmountOut
					result[ri].AmountQuotes[ra].SqrtPriceX96AfterList = append(result[ri].AmountQuotes[ra].SqrtPriceX96AfterList, quoteData.Data.SqrtPriceX96AfterList...)
					result[ri].AmountQuotes[ra].InitializedTicksCrossedList = append(result[ri].AmountQuotes[ra].InitializedTicksCrossedList, quoteData.Data.InitializedTicksCrossedList...)
					result[ri].AmountQuotes[ra].GasEstimate.Add(result[ri].AmountQuotes[ra].GasEstimate, quoteData.Data.GasEstimate)
				}
				callDataIndex++
			}
		}
	}

	return result, 0, nil
}
