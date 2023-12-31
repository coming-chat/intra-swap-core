package v3

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/dex"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"log/slog"
	"math/big"
	"sync"
)

const cacheKeyFormat = "%d:%s-%s/%d"

type TokenPairs struct {
	Token0         *entities.Token
	Token1         *entities.Token
	FeeAmount      constants.FeeAmount
	PairAddress    string
	FactoryAddress string
	QuoteAddress   string
	RouterAddress  string
}

type PoolProvider interface {
	GetPools(tokenPairs []TokenPairs, providerConfig *config.Config) (PoolAccessor, error)

	GetPoolAddress(
		tokenA,
		tokenB *entities.Token,
		//factoryAddress common.Address,
		feeAmount constants.FeeAmount,
	) (poolAddress string, token0, token1 *entities.Token, err error)
}

type PoolAccessor interface {
	GetPool(
		tokenA,
		tokenB *entities.Token,
		//factoryAddress common.Address,
		feeAmount constants.FeeAmount,
	) (*base_entities.V3Pool, error)
	GetPoolByAddress(address string) *base_entities.V3Pool
	GetAllPools() []*base_entities.V3Pool
	GetBasePool(subPool *entitiesV3.Pool) *base_entities.V3Pool
}

type BasePoolAccessor struct {
	poolAddressToPool map[string]*base_entities.V3Pool
	pools             []*base_entities.V3Pool
	subPoolMap        map[*entitiesV3.Pool]*base_entities.V3Pool
	getPoolAddress    func(tokenA, tokenB *entities.Token,
		//factoryAddress common.Address,
		feeAmount constants.FeeAmount,
	) (poolAddress string, token0, token1 *entities.Token, err error)
}

func (b BasePoolAccessor) GetBasePool(subPool *entitiesV3.Pool) *base_entities.V3Pool {
	return b.subPoolMap[subPool]
}

func (b BasePoolAccessor) GetPool(
	tokenA, tokenB *entities.Token,
	//factoryAddress common.Address,
	feeAmount constants.FeeAmount,
) (*base_entities.V3Pool, error) {
	address, _, _, err := b.getPoolAddress(tokenA, tokenB, feeAmount)
	if err != nil {
		return nil, err
	}
	return b.poolAddressToPool[address], nil
}

func (b BasePoolAccessor) GetPoolByAddress(address string) *base_entities.V3Pool {
	return b.poolAddressToPool[address]
}

func (b BasePoolAccessor) GetAllPools() []*base_entities.V3Pool {
	return b.pools
}

type PoolRetryOptions struct {
	Retries    int
	MinTimeout int
	MaxTimeout int
}

type BasePoolProvider struct {
	PoolAddressCache  map[string]string
	ChainId           base_entities.ChainId
	MultiCallProvider rpc.MultiCallProviderCore
	RetryOptions      PoolRetryOptions
}

func NewBasePoolProvider(
	chainId base_entities.ChainId,
	poolAddressCache map[string]string,
	multiCallCore rpc.MultiCallProviderCore,
	retryOptions *PoolRetryOptions,
) *BasePoolProvider {
	basePool := &BasePoolProvider{
		ChainId:           chainId,
		PoolAddressCache:  poolAddressCache,
		MultiCallProvider: multiCallCore,
	}
	if retryOptions == nil {
		retryOptions = &PoolRetryOptions{
			Retries:    2,
			MinTimeout: 50,
			MaxTimeout: 500,
		}
	}
	basePool.RetryOptions = *retryOptions
	return basePool
}

func (b *BasePoolProvider) GetPools(tokenPairs []TokenPairs, providerConfig *config.Config) (PoolAccessor, error) {
	var (
		poolAddressSet = make(map[string]struct{})
		sortedPool     []struct {
			FeeAmount      constants.FeeAmount
			Token0         *entities.Token
			Token1         *entities.Token
			PairAddress    string
			FactoryAddress string
			QuoteAddress   string
			RouterAddress  string
		}
		sortedPoolAddresses []string
		slot0s              []rpc.MultiCallSingle[dex.IPoolState]
		liquiditys          []rpc.MultiCallSingle[dex.ILiquidity]
	)

	var err error
	for _, pair := range tokenPairs {
		var address string
		token0, token1 := pair.Token0, pair.Token1
		address, token0, token1, err = b.GetPoolAddress(pair.Token0, pair.Token1, pair.FeeAmount)
		if err != nil {
			return nil, err
		}
		if pair.PairAddress == "" {
			key := fmt.Sprintf(cacheKeyFormat, b.ChainId, token0.Address.String(), token1.Address.String(), pair.FeeAmount)
			if cacheAddr, has := b.PoolAddressCache[key]; !has {
				pair.PairAddress = address
			} else {
				pair.PairAddress = cacheAddr
			}
		}

		if _, ok := poolAddressSet[pair.PairAddress]; ok {
			continue
		}
		poolAddressSet[pair.PairAddress] = struct{}{}
		sortedPool = append(sortedPool, struct {
			FeeAmount      constants.FeeAmount
			Token0         *entities.Token
			Token1         *entities.Token
			PairAddress    string
			FactoryAddress string
			QuoteAddress   string
			RouterAddress  string
		}{
			FeeAmount:      pair.FeeAmount,
			Token0:         token0,
			Token1:         token1,
			PairAddress:    pair.PairAddress,
			QuoteAddress:   pair.QuoteAddress,
			FactoryAddress: pair.FactoryAddress,
			RouterAddress:  pair.RouterAddress,
		})
		sortedPoolAddresses = append(sortedPoolAddresses, pair.PairAddress)
		slot0s = append(slot0s, dex.RouterAddrDexMap[common.HexToAddress(pair.RouterAddress)].GetPoolInfo().GetSlot0(common.HexToAddress(pair.PairAddress)))
		liquiditys = append(liquiditys, dex.RouterAddrDexMap[common.HexToAddress(pair.RouterAddress)].GetPoolInfo().GetLiquidity(common.HexToAddress(pair.PairAddress)))
	}
	var (
		syncGroup              = sync.WaitGroup{}
		errSlot0, errLiquidity error
	)

	syncGroup.Add(2)
	go func() {
		defer syncGroup.Done()
		for i := 0; i < b.RetryOptions.Retries; i++ {
			_, _, errSlot0 = rpc.GetMultiCallProvider[dex.IPoolState](b.MultiCallProvider).MultiCall(
				b.ChainId,
				slot0s,
				0,
				false,
				providerConfig,
			)
			if errSlot0 != nil {
				continue
			}
			break
		}
	}()

	go func() {
		defer syncGroup.Done()
		for i := 0; i < b.RetryOptions.Retries; i++ {
			_, _, errLiquidity = rpc.GetMultiCallProvider[dex.ILiquidity](b.MultiCallProvider).MultiCall(
				b.ChainId,
				liquiditys,
				0,
				false,
				providerConfig,
			)
			if errLiquidity != nil {
				continue
			}
			break
		}
	}()

	syncGroup.Wait()

	if errSlot0 != nil {
		return nil, errSlot0
	}

	if errLiquidity != nil {
		return nil, errLiquidity
	}

	poolAccessor := BasePoolAccessor{
		poolAddressToPool: make(map[string]*base_entities.V3Pool),
		subPoolMap:        make(map[*entitiesV3.Pool]*base_entities.V3Pool),
		getPoolAddress:    b.GetPoolAddress,
	}

	//var (
	//	tickInfoParams []rpc.MultiCallSingle[dex.PoolTick]
	//	ticks          [][]entitiesV3.Tick
	//)
	//
	//for i, address := range sortedPoolAddresses {
	//	if !slot0s[i].CallResult.Success ||
	//		!liquiditys[i].CallResult.Success ||
	//		slot0s[i].CallResult.Data.GetPrice().Cmp(big.NewInt(0)) == 0 {
	//		continue
	//	}
	//
	//	ticks = append(ticks, []entitiesV3.Tick{
	//		{
	//			Index: entitiesV3.NearestUsableTick(utils.MinTick, constants.TickSpacings[sortedPool[i].FeeAmount]),
	//		},
	//		{
	//			Index: entitiesV3.NearestUsableTick(utils.MinTick, constants.TickSpacings[sortedPool[i].FeeAmount]),
	//		},
	//	})
	//
	//	tickInfoParams = append(tickInfoParams, dex.RouterAddrDexMap[common.HexToAddress(sortedPool[i].RouterAddress)].GetPoolInfo().GetTicks(common.HexToAddress(address), slot0s[i].CallResult.Data.GetTick()))
	//}

	//_, _, err = rpc.GetMultiCallProvider[dex.PoolTick](b.MultiCallProvider).MultiCall(
	//	b.ChainId,
	//	tickInfoParams,
	//	0,
	//	false,
	//	providerConfig,
	//)
	//if err != nil {
	//	return nil, err
	//}

	//tickInfoResultIndex := 0
	for i, address := range sortedPoolAddresses {
		if !slot0s[i].CallResult.Success ||
			!liquiditys[i].CallResult.Success ||
			slot0s[i].CallResult.Data.GetPrice().Cmp(big.NewInt(0)) == 0 { // || !tickInfoParams[tickInfoResultIndex].CallResult.Success
			continue
		}
		//	// create tick data provider
		//	ticks[tickInfoResultIndex][0].LiquidityNet = tickInfoParams[tickInfoResultIndex].CallResult.Data.GetLiquidityNet()
		//	ticks[tickInfoResultIndex][1].LiquidityNet = new(big.Int).Neg(tickInfoParams[tickInfoResultIndex].CallResult.Data.GetLiquidityNet())
		//	ticks[tickInfoResultIndex][0].LiquidityGross = tickInfoParams[tickInfoResultIndex].CallResult.Data.GetLiquidityGross()
		//	ticks[tickInfoResultIndex][1].LiquidityGross = tickInfoParams[tickInfoResultIndex].CallResult.Data.GetLiquidityGross()
		p, err := entitiesV3.NewTickListDataProvider([]entitiesV3.Tick{
			{
				Index:          entitiesV3.NearestUsableTick(utils.MinTick, constants.TickSpacings[sortedPool[i].FeeAmount]),
				LiquidityNet:   big.NewInt(0),
				LiquidityGross: big.NewInt(0),
			},
			{
				Index:          entitiesV3.NearestUsableTick(utils.MaxTick, constants.TickSpacings[sortedPool[i].FeeAmount]),
				LiquidityNet:   big.NewInt(0),
				LiquidityGross: big.NewInt(0),
			},
		}, constants.TickSpacings[sortedPool[i].FeeAmount])
		if err != nil {
			slog.Error("init v3 TickListDataProvider err: %v", err)
			continue
		}
		sqrtPriceX96 := slot0s[i].CallResult.Data.GetPrice()
		liquidity := liquiditys[i].CallResult.Data.GetLiquidity()
		if sqrtPriceX96 == nil {
			sqrtPriceX96 = big.NewInt(0)
		}
		if liquidity == nil {
			liquidity = big.NewInt(0)
		}

		v3pool, err := entitiesV3.NewPool(
			sortedPool[i].Token0,
			sortedPool[i].Token1,
			sortedPool[i].FeeAmount,
			sqrtPriceX96,
			liquidity,
			int(slot0s[i].CallResult.Data.GetTick().Int64()),
			p,
		)
		if err != nil {
			slog.Error("new v3Pool err: %v", err)
			continue
		}
		pool := base_entities.NewV3Pool(v3pool, sortedPool[i].PairAddress, sortedPool[i].QuoteAddress, sortedPool[i].RouterAddress, sortedPool[i].FactoryAddress, false)
		poolAccessor.poolAddressToPool[common.HexToAddress(address).String()] = pool
		poolAccessor.subPoolMap[pool.Pool] = pool
		b.PoolAddressCache[fmt.Sprintf(cacheKeyFormat, b.ChainId, sortedPool[i].Token0.Address.String(), sortedPool[i].Token1.Address.String(), sortedPool[i].FeeAmount)] = common.HexToAddress(address).String()
		poolAccessor.pools = append(poolAccessor.pools, pool)
		//tickInfoResultIndex++
	}
	return poolAccessor, nil
}

func (b *BasePoolProvider) GetPoolAddress(
	tokenA, tokenB *entities.Token,
	//factoryAddress common.Address,
	feeAmount constants.FeeAmount,
) (string, *entities.Token, *entities.Token, error) {
	before, err := tokenA.SortsBefore(tokenB)
	if err != nil {
		return "", nil, nil, err
	}
	token0, token1 := tokenA, tokenB
	if !before {
		token0, token1 = tokenB, tokenA
	}
	key := fmt.Sprintf(cacheKeyFormat, b.ChainId, token0.Address.String(), token1.Address.String(), feeAmount)
	cacheAddress, ok := b.PoolAddressCache[key]
	if ok {
		return cacheAddress, token0, token1, nil
	}
	// offline
	poolAddress, err := utils.ComputePoolAddress(base_constant.BaseUniswapV3Factory, token0, token1, feeAmount, "")
	//abi, err := uniswap_v3.FactoryMetaData.GetAbi()
	//if err != nil {
	//	return "", nil, nil, err
	//}
	//result, err := rpc.NewUniswapMultiCallProvider[common.Address](b.MultiCallProvider).MultiCall([]rpc.MultiCallSingleParam{
	//	{
	//		Contract:     abi,
	//		FunctionName: "getPool",
	//		FunctionParams: []any{
	//			token0.Address,
	//			token1.Address,
	//			big.NewInt(int64(feeAmount)),
	//		},
	//		ContractAddress: util.BaseUniswapV3Factory,
	//	},
	//}, nil)
	if err != nil {
		return "", nil, nil, err
	}
	//if len(result.Data) == 0 || !result.Data[0].Success {
	//	return "", nil, nil, errors.New("not found the pair")
	//}
	b.PoolAddressCache[key] = poolAddress.String()
	return poolAddress.String(), token0, token1, nil
}
