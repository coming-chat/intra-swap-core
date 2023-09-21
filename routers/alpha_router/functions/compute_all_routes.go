package functions

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

func ComputeAllV3Routes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	pools []*base_entities.V3Pool,
	maxHops int,
) ([]*base_entities.MRoute, error) {
	var ipool []base_entities.Pool
	for _, pool := range pools {
		ipool = append(ipool, pool)
	}
	return computeAllRoutes(
		tokenIn,
		tokenOut,
		func(pools []base_entities.Pool, tokenIn *entities.Token, tokenOut *entities.Token) (*base_entities.MRoute, error) {
			return base_entities.NewMRoute(pools, tokenIn, tokenOut)
		},
		ipool,
		maxHops,
	)
}

func ComputeAllV2Routes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	pools []*base_entities.V2Pool,
	maxHops int,
) ([]*base_entities.MRoute, error) {
	var ipool []base_entities.Pool
	for _, pool := range pools {
		ipool = append(ipool, pool)
	}
	return computeAllRoutes(
		tokenIn,
		tokenOut,
		func(pools []base_entities.Pool, tokenIn *entities.Token, tokenOut *entities.Token) (*base_entities.MRoute, error) {
			return base_entities.NewMRoute(pools, tokenIn, tokenOut)
		},
		ipool,
		maxHops,
	)
}

func computeAllRoutes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	buildRoute func(route []base_entities.Pool, tokenIn *entities.Token, tokenOut *entities.Token) (*base_entities.MRoute, error),
	pools []base_entities.Pool,
	maxHops int,
) ([]*base_entities.MRoute, error) {
	var (
		routes        []*base_entities.MRoute
		poolsUsed     = make([]bool, len(pools))
		computeRoutes func(
			tokenIn *entities.Token,
			tokenOut *entities.Token,
			currentRoute []base_entities.Pool,
			poolsUsed []bool,
			_previousTokenOut *entities.Token,
		) error
	)

	computeRoutes = func(
		tokenIn *entities.Token,
		tokenOut *entities.Token,
		currentRoute []base_entities.Pool,
		poolsUsed []bool,
		_previousTokenOut *entities.Token,
	) error {
		if len(currentRoute) > maxHops {
			return nil
		}
		if len(currentRoute) > 0 && currentRoute[len(currentRoute)-1].InvolvesToken(tokenOut) {
			foundRoute := make([]base_entities.Pool, len(currentRoute))
			copy(foundRoute, currentRoute)
			route, err := buildRoute(foundRoute, tokenIn, tokenOut)
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

	err := computeRoutes(tokenIn, tokenOut, []base_entities.Pool{}, poolsUsed, nil)
	if err != nil {
		return nil, err
	}
	return routes, nil
}
