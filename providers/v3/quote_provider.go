package v3

import (
	"context"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"
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
	Route       *entitiesV3.Route
	AmountQuote []AmountQuote
}

type functionName = string

const (
	QuoteExactInput  functionName = "quoteExactInput"
	QuoteExactOutput functionName = "quoteExactOutput"
)

type QuoteProvider interface {
	GetQuotesManyExactIn(
		amountIns []*entities.CurrencyAmount,
		routes []*base_entities.V3Route,
	) ([]RouteWithQuotes, error)

	GetQuotesManyExactOut(
		amountOuts []*entities.CurrencyAmount,
		routes []*base_entities.V3Route,
	) ([]RouteWithQuotes, error)
}

type BatchParams struct {
	MultiCallChunk      int
	GasLimitPerCall     int64
	QuoteMinSuccessRate float64
}

type BlockNumberConfig struct {
	BaseBlockOffset uint64
	Rollback        bool
}

type QuoteRetryOptions struct {
	Retries    int
	MinTimeout int
	MaxTimeout int
}

type BaseQuoteProvider struct {
	QuoterAddress      common.Address
	ChainId            base_entities.ChainId
	Provider           rpc.BaseProvider
	MultiCall2Provider rpc.MultiCallProviderCore
	BatchParams        BatchParams
	BlockNumberConfig  BlockNumberConfig
	RetryOptions       QuoteRetryOptions
	Ctx                context.Context
}

func NewBaseQuoteProvider(
	ctx context.Context,
	address common.Address,
	chainId base_entities.ChainId,
	baseProvider rpc.BaseProvider,
	multiCallCore rpc.MultiCallProviderCore,
	batchParams *BatchParams,
	blockNumberConfig *BlockNumberConfig,
	retryOptions *QuoteRetryOptions,
) *BaseQuoteProvider {
	quoteProvider := &BaseQuoteProvider{
		ChainId:            chainId,
		Ctx:                ctx,
		QuoterAddress:      address,
		Provider:           baseProvider,
		MultiCall2Provider: multiCallCore,
	}
	if batchParams == nil {
		batchParams = &BatchParams{
			MultiCallChunk:      150,
			GasLimitPerCall:     1_000_000,
			QuoteMinSuccessRate: 0.2,
		}
	}
	if blockNumberConfig == nil {
		blockNumberConfig = &BlockNumberConfig{
			BaseBlockOffset: 0,
			Rollback:        false,
		}
	}
	if retryOptions == nil {
		retryOptions = &QuoteRetryOptions{
			Retries:    2,
			MinTimeout: 25,
			MaxTimeout: 250,
		}
	}
	quoteProvider.BatchParams = *batchParams
	quoteProvider.BlockNumberConfig = *blockNumberConfig
	quoteProvider.RetryOptions = *retryOptions
	return quoteProvider
}

func (b *BaseQuoteProvider) GetQuotesManyExactIn(
	amountIns []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
) ([]RouteWithQuotes, error) {
	return b.getQuotesManyData(amountIns, routes, entities.ExactInput)
}

func (b *BaseQuoteProvider) GetQuotesManyExactOut(
	amountOuts []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
) ([]RouteWithQuotes, error) {
	return b.getQuotesManyData(amountOuts, routes, entities.ExactOutput)
}

func (b *BaseQuoteProvider) getQuotesManyData(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
	tradeType entities.TradeType,
) ([]RouteWithQuotes, error) {
	var (
		routesQuotes []RouteWithQuotes
		err          error
	)
	for _, route := range routes {
		var amountQuotes []AmountQuote
		for _, amount := range amounts {
			cycleAmount := amount.Wrapped()
			switch tradeType {
			case entities.ExactInput:
				for _, pool := range route.Pools {
					cycleAmount, _, err = pool.GetOutputAmount(cycleAmount, nil)
					if err != nil {
						return nil, err
					}
				}
			case entities.ExactOutput:
				for i := len(route.Pools) - 1; i >= 0; i-- {
					cycleAmount, _, err = route.Pools[i].GetInputAmount(cycleAmount, nil)
					if err != nil {
						return nil, err
					}
				}
			default:
				return nil, fmt.Errorf("unsupported tradeTyep[%#v]", tradeType)
			}
			amountQuotes = append(amountQuotes, AmountQuote{
				Amount: amount,
				Quote:  cycleAmount.Quotient(),
			})
		}
		routesQuotes = append(routesQuotes, RouteWithQuotes{
			Route:       route,
			AmountQuote: amountQuotes,
		})
	}

	return routesQuotes, nil
}
