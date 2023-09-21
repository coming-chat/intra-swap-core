package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
)

//type (
//	V3Route = entitiesV3.Route
//	V2Route = entitiesV2.Route
//)

type Protocol int

const (
	_ Protocol = iota
	_
	V2
	V3
)

//type Route interface {
//	*V3Route | *V2Route
//}

func NewMRoute(pools []Pool, input, output entities.Currency) (*MRoute, error) {
	if len(pools) == 0 {
		return nil, entitiesV3.ErrRouteNoPools
	}
	chainID := pools[0].ChainID()
	for _, p := range pools {
		if p.ChainID() != chainID {
			return nil, entitiesV3.ErrAllOnSameChain
		}
	}
	if input == nil || !pools[0].InvolvesToken(input.Wrapped()) {
		return nil, entitiesV3.ErrInputNotInvolved
	}

	// Normalizes token0-token1 order and selects the next token/fee step to add to the path

	path := []*entities.Token{input.Wrapped()}
	for i, p := range pools {
		currentInputToken := path[i]
		if !(currentInputToken.Equal(p.Token0()) || currentInputToken.Equal(p.Token1())) {
			return nil, entitiesV3.ErrPathNotContinuous
		}
		nextToken := p.Token0()
		if currentInputToken.Equal(p.Token0()) {
			nextToken = p.Token1()
		}
		path = append(path, nextToken)
	}

	if output == nil {
		output = path[len(path)-1]
	} else {
		if !pools[len(pools)-1].InvolvesToken(output.Wrapped()) {
			return nil, entitiesV3.ErrOutputNotInvolved
		}
	}
	return &MRoute{
		Pools:  pools,
		Path:   path,
		Input:  input,
		Output: output,
	}, nil
}

type MRoute struct {
	Pools  []Pool
	Path   []*entities.Token
	Input  entities.Currency
	Output entities.Currency

	midPrice *entities.Price
}

// MidPrice Returns the mid price of the route
func (r *MRoute) MidPrice() (*entities.Price, error) {
	if r.midPrice != nil {
		return r.midPrice, nil
	}
	var (
		nextInput *entities.Token
		price     *entities.Price
	)
	if r.Pools[0].Token0().Equal(r.Input) {
		nextInput = r.Pools[0].Token1()
		price = r.Pools[0].Token0Price()
	} else {
		nextInput = r.Pools[0].Token0()
		price = r.Pools[0].Token1Price()
	}
	price, err := reducePrice(nextInput, price, r.Pools[1:])
	if err != nil {
		return nil, err
	}
	r.midPrice = entities.NewPrice(r.Input, r.Output, price.Denominator, price.Numerator)
	return r.midPrice, nil
}

// reducePrice reduces the price of the route by the given amount
func reducePrice(nextInput *entities.Token, price *entities.Price, pools []Pool) (*entities.Price, error) {
	var err error
	for _, p := range pools {
		if nextInput.Equal(p.Token0()) {
			nextInput = p.Token1()
			price, err = price.Multiply(p.Token0Price())
			if err != nil {
				return nil, err
			}
		} else {
			nextInput = p.Token0()
			price, err = price.Multiply(p.Token1Price())
			if err != nil {
				return nil, err
			}
		}
	}
	return price, nil
}

type SwapAndAddParameters struct {
	// starting balance for tokenIn which will inform the tokenIn position amount
	InitialBalanceTokenIn *entities.CurrencyAmount
	// starting balance for tokenOut which will inform the tokenOut position amount
	InitialBalanceTokenOut *entities.CurrencyAmount
	// position details needed to create a new Position with the known liquidity amounts
	PreLiquidityPosition *entitiesV3.Position
}
