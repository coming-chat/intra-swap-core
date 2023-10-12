package v2

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
	"math/big"
)

const cacheKeyFormat = "%d:%s-%s"

type TokenPairs struct {
	Token0         *entities.Token
	Token1         *entities.Token
	PairAddress    string
	FactoryAddress string
	RouterAddress  string
}

type IReserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

type PoolProvider interface {
	GetPools(
		tokenPairs []TokenPairs,
		providerConfig *config.Config,
	) (PoolAccessor, error)

	GetPoolAddress(
		tokenA *entities.Token,
		tokenB *entities.Token,
		//factoryAddress common.Address,
	) (poolAddress string, token0, token1 *entities.Token, err error)
}

type PoolAccessor interface {
	GetPool(
		tokenA, tokenB *entities.Token,
	) (*base_entities.V2Pool, error)
	GetPoolByAddress(address string) *base_entities.V2Pool
	GetAllPools() []*base_entities.V2Pool
	GetBasePool(subPool *entitiesV2.Pair) *base_entities.V2Pool
}

type BasePoolAccessor struct {
	poolAddressToPool map[string]*base_entities.V2Pool
	pools             []*base_entities.V2Pool
	subPoolMap        map[*entitiesV2.Pair]*base_entities.V2Pool
	getPool           func(
		tokenA, tokenB *entities.Token,
		//factoryAddress common.Address,
	) (poolAddress string, token0, token1 *entities.Token, err error)
}

func (b *BasePoolAccessor) GetBasePool(subPool *entitiesV2.Pair) *base_entities.V2Pool {
	return b.subPoolMap[subPool]
}

func (b *BasePoolAccessor) GetPool(
	tokenA, tokenB *entities.Token,
	// factoryAddress common.Address,
) (*base_entities.V2Pool, error) {
	poolAddress, _, _, err := b.getPool(tokenA, tokenB)
	if err != nil {
		return nil, err
	}
	return b.poolAddressToPool[poolAddress], nil
}

func (b *BasePoolAccessor) GetPoolByAddress(address string) *base_entities.V2Pool {
	return b.poolAddressToPool[address]
}

func (b *BasePoolAccessor) GetAllPools() []*base_entities.V2Pool {
	return b.pools
}

type PoolRetryOptions struct {
	Retries    int
	MinTimeout int
	MaxTimeout int
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

type BasePoolProvider struct {
	PoolAddressCache  map[string]string
	ChainId           base_entities.ChainId
	MultiCallProvider rpc.MultiCallProviderCore
	RetryOptions      PoolRetryOptions
}

func (b *BasePoolProvider) GetPools(tokenPairs []TokenPairs, providerConfig *config.Config) (PoolAccessor, error) {
	var (
		poolAddressSet   = make(map[string]struct{})
		sortedTokenPairs []struct {
			Token0 *entities.Token
			Token1 *entities.Token
		}
		sortedPoolAddresses []string
		multiCallData       []rpc.MultiCallSingle[IReserves]
		err                 error
	)
	for _, pair := range tokenPairs {
		var address string
		token0, token1 := pair.Token0, pair.Token1
		address, token0, token1, err = b.GetPoolAddress(pair.Token0, pair.Token1)
		if err != nil {
			return nil, err
		}
		if pair.PairAddress == "" {
			key := fmt.Sprintf(cacheKeyFormat, b.ChainId, token0.Address.String(), token1.Address.String())
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
		sortedTokenPairs = append(sortedTokenPairs, struct {
			Token0 *entities.Token
			Token1 *entities.Token
		}{Token0: token0, Token1: token1})
		sortedPoolAddresses = append(sortedPoolAddresses, pair.PairAddress)

		multiCallData = append(multiCallData, rpc.MultiCallSingle[IReserves]{
			FunctionName:    "getReserves",
			Contract:        contracts.IUniswapV2PoolAbi,
			ContractAddress: common.HexToAddress(pair.PairAddress),
		})
	}

	_, _, err = rpc.GetMultiCallProvider[IReserves](b.MultiCallProvider).MultiCall(b.ChainId, multiCallData, 0, false, nil)
	if err != nil {
		return nil, err
	}
	poolAccessor := &BasePoolAccessor{
		poolAddressToPool: make(map[string]*base_entities.V2Pool),
		subPoolMap:        make(map[*entitiesV2.Pair]*base_entities.V2Pool),
		getPool:           b.GetPoolAddress,
	}
	for i, result := range multiCallData {
		if !result.CallResult.Success {
			continue
		}
		poolAddress := common.HexToAddress(sortedPoolAddresses[i])
		pair, err := entitiesV2.NewPair(
			entities.FromRawAmount(sortedTokenPairs[i].Token0, result.CallResult.Data.Reserve0),
			entities.FromRawAmount(sortedTokenPairs[i].Token1, result.CallResult.Data.Reserve1),
			&entitiesV2.PairOptions{
				Address: &poolAddress,
			})
		if err != nil {
			return nil, err
		}
		pool := base_entities.NewV2Pool(pair, tokenPairs[i].PairAddress, tokenPairs[i].RouterAddress, tokenPairs[i].RouterAddress, tokenPairs[i].FactoryAddress, false)
		poolAccessor.poolAddressToPool[poolAddress.String()] = pool
		poolAccessor.subPoolMap[pool.Pair] = pool
		b.PoolAddressCache[fmt.Sprintf(cacheKeyFormat, b.ChainId, sortedTokenPairs[i].Token0.Address.String(), sortedTokenPairs[i].Token1.Address.String())] = poolAddress.String()
		poolAccessor.pools = append(poolAccessor.pools, pool)
	}
	return poolAccessor, nil
}

func (b *BasePoolProvider) GetPoolAddress(
	tokenA, tokenB *entities.Token,
	// factoryAddress common.Address,
) (string, *entities.Token, *entities.Token, error) {
	token0, token1 := tokenA, tokenB
	before, err := tokenA.SortsBefore(tokenB)
	if err != nil {
		return "", nil, nil, err
	}
	if !before {
		token0, token1 = tokenB, tokenA
	}

	key := fmt.Sprintf(cacheKeyFormat, b.ChainId, token0.Address.String(), token1.Address.String())
	cachePoolCache, ok := b.PoolAddressCache[key]
	if ok {
		return cachePoolCache, token0, token1, nil
	}
	// offline
	address, err := entitiesV2.GetAddress(token0, token1, base_constant.BaseAlienbaseV2Factory, common.FromHex("0x60faccfea37b2ca2a11c3f92823f437251b05755a3d924bd97259ab624057faa"))
	//if err != nil {
	//abi, err := uniswap_v2.PairFactoryMetaData.GetAbi()
	//if err != nil {
	//	return "", nil, nil, err
	//}
	//result, err := rpc.NewUniswapMultiCallProvider[common.Address](b.MultiCallProvider).MultiCall([]rpc.MultiCallSingleParam{
	//	{
	//		Contract:     abi,
	//		FunctionName: "getPair",
	//		FunctionParams: []any{
	//			token0.Address,
	//			token1.Address,
	//		},
	//		ContractAddress: util.AlienbaseV2Factory,
	//	},
	//}, nil)
	if err != nil {
		return "", nil, nil, err
	}
	//if len(result.Data) == 0 || !result.Data[0].Success {
	//	return "", nil, nil, errors.New("not found the pair")
	//}
	//b.PoolAddressCache[key] = address.String()
	return address.String(), token0, token1, nil
}
