package v2

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"strings"
)

type IndexerPool struct {
	Id             string
	RouterAddress  string
	QuoteAddress   string
	FactoryAddress string
	Token0         struct {
		Id string
	}
	Token1 struct {
		Id string
	}
	Supply  decimal.Decimal
	Reserve decimal.Decimal
	Dex     string
}

type IndexerProvider interface {
	GetPools(tokenIn *entities.Token, tokenOut *entities.Token, providerConfig *config.Config) ([]IndexerPool, error)
}

func NewGeckoTerminalProvider(chaiId base_entities.ChainId) *GeckoTerminalProvider {
	return &GeckoTerminalProvider{
		api:        "https://api.geckoterminal.com/api/v2",
		poolSearch: "/networks/base/dexes/alien-base/pools",
		headers: map[string]string{
			"Accept":     "application/json;version=20230302",
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		},
		ChainId: chaiId,
	}
}

type GeckoTerminalProvider struct {
	ChainId    base_entities.ChainId
	api        string
	headers    map[string]string
	poolSearch string
}

type GeckoTerminalRespData struct {
	Data []struct {
		Id         string
		Type       string
		Attributes struct {
			Address      string
			Name         string
			ReserveInUsd string `json:"reserve_in_usd"`
		}
		Relationships struct {
			Dex struct {
				Data struct {
					Id   string
					Type string
				}
			}
			BaseToken struct {
				Data struct {
					Id   string
					Type string
				}
			} `json:"base_token"`
			QuoteToken struct {
				Data struct {
					Id   string
					Type string
				}
			} `json:"quote_token"`
		}
	}
}

func (g *GeckoTerminalProvider) GetPools(tokenIn *entities.Token, tokenOut *entities.Token, providerConfig *config.Config) (pools []IndexerPool, err error) {
	params := map[string]string{}
	geckoResp, err := util.GetReq[GeckoTerminalRespData](g.api+g.poolSearch, g.headers, params)
	if err != nil {
		return nil, err
	}
	for _, poolData := range geckoResp.Data {
		reserve, err := decimal.NewFromString(poolData.Attributes.ReserveInUsd)
		if err != nil {
			return nil, err
		}
		pools = append(pools, IndexerPool{
			Id:             common.HexToAddress(poolData.Attributes.Address).String(),
			Token0:         struct{ Id string }{Id: common.HexToAddress(strings.TrimPrefix(poolData.Relationships.BaseToken.Data.Id, "base_")).String()},
			Token1:         struct{ Id string }{Id: common.HexToAddress(strings.TrimPrefix(poolData.Relationships.QuoteToken.Data.Id, "base_")).String()},
			Reserve:        reserve,
			Dex:            poolData.Relationships.Dex.Data.Id,
			RouterAddress:  "0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7",
			FactoryAddress: "",
			QuoteAddress:   "0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7",
		})
	}
	return
}
