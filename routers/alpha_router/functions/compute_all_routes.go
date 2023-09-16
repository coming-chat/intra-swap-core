package functions

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
)

func ComputeAllV3Routes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	pools []*base_entities.V3Pool,
	maxHops int,
) ([]*base_entities.V3Route, error) {
	return computeAllRoutes[*base_entities.V3Pool, *base_entities.V3Route](
		tokenIn,
		tokenOut,
		func(pools []*base_entities.V3Pool, tokenIn *entities.Token, tokenOut *entities.Token) (*base_entities.V3Route, error) {
			var v3Pools []*entitiesV3.Pool
			for _, p := range pools {
				v3Pools = append(v3Pools, p.Pool)
			}
			return entitiesV3.NewRoute(v3Pools, tokenIn, tokenOut)
		},
		pools,
		maxHops,
	)
}

func ComputeAllV2Routes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	pools []*base_entities.V2Pool,
	maxHops int,
) ([]*base_entities.V2Route, error) {
	return computeAllRoutes[*base_entities.V2Pool, *base_entities.V2Route](
		tokenIn,
		tokenOut,
		func(pools []*base_entities.V2Pool, tokenIn *entities.Token, tokenOut *entities.Token) (*entitiesV2.Route, error) {
			var v2Pools []*entitiesV2.Pair
			for _, p := range pools {
				v2Pools = append(v2Pools, p.Pair)
			}
			return entitiesV2.NewRoute(v2Pools, tokenIn, tokenOut)
		},
		pools,
		maxHops,
	)
}

func computeAllRoutes[TPool base_entities.Pool, TRoute base_entities.Route](
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	buildRoute func(route []TPool, tokenIn *entities.Token, tokenOut *entities.Token) (TRoute, error),
	pools []TPool,
	maxHops int,
) ([]TRoute, error) {
	var (
		routes        []TRoute
		poolsUsed     = make([]bool, len(pools))
		computeRoutes func(
			tokenIn *entities.Token,
			tokenOut *entities.Token,
			currentRoute []TPool,
			poolsUsed []bool,
			_previousTokenOut *entities.Token,
		) error
	)

	computeRoutes = func(
		tokenIn *entities.Token,
		tokenOut *entities.Token,
		currentRoute []TPool,
		poolsUsed []bool,
		_previousTokenOut *entities.Token,
	) error {
		if len(currentRoute) > maxHops {
			return nil
		}
		if len(currentRoute) > 0 && currentRoute[len(currentRoute)-1].InvolvesToken(tokenOut) {
			route, err := buildRoute(currentRoute, tokenIn, tokenOut)
			if err != nil {
				return err
			}
			routes = append(routes, route)
			return nil
		}

		for i := 0; i < len(pools); i++ {
			if poolsUsed[i] {
				continue
			}

			var (
				curPool          = pools[i]
				previousTokenOut *entities.Token
			)
			if _previousTokenOut != nil {
				previousTokenOut = _previousTokenOut
			} else {
				previousTokenOut = tokenIn
			}

			if !curPool.InvolvesToken(previousTokenOut) {
				continue
			}

			var currentTokenOut *entities.Token
			if curPool.Token0().Equal(previousTokenOut) {
				currentTokenOut = curPool.Token1()
			} else {
				currentTokenOut = curPool.Token0()
			}

			currentRoute = append(currentRoute, curPool)
			poolsUsed[i] = true
			err := computeRoutes(
				tokenIn,
				tokenOut,
				currentRoute,
				poolsUsed,
				currentTokenOut,
			)
			if err != nil {
				return err
			}
			poolsUsed[i] = false
			currentRoute = currentRoute[:len(currentRoute)-1]
		}
		return nil
	}

	err := computeRoutes(tokenIn, tokenOut, []TPool{}, poolsUsed, nil)
	if err != nil {
		return nil, err
	}
	return routes, nil
}
