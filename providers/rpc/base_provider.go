package rpc

import (
	"context"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/rand"
	"sync"
)

type IntraClient struct {
	*ethclient.Client
	Score int64
}

type Provider interface {
	GetEthRpc(chainId base_entities.ChainId) (*ethclient.Client, error)
}

func NewBaseProvider(ctx context.Context, rpcAddresses []string, maxScore int64) (*BaseProvider, error) {
	p := &BaseProvider{
		RWMutex:  sync.RWMutex{},
		rpcs:     make(map[base_entities.ChainId][]*IntraClient),
		maxScore: maxScore,
		ctx:      ctx,
	}
	for _, address := range rpcAddresses {
		client, err := ethclient.Dial(address)
		if err != nil {
			logrus.Errorf("init ethclient to rpc [%s] err: %v", address, err)
			continue
		}
		id, err := client.ChainID(ctx)
		if err != nil {
			logrus.Errorf("get eth chain id failed: %v", err)
			continue
		}
		//TODO make the score useful
		p.rpcs[base_entities.ChainId(id.Uint64())] = append(p.rpcs[base_entities.ChainId(id.Uint64())], &IntraClient{Client: client})
	}
	return p, nil
}

type BaseProvider struct {
	sync.RWMutex
	ctx      context.Context
	rpcs     map[base_entities.ChainId][]*IntraClient
	maxScore int64
}

func (b *BaseProvider) GetEthRpc(chanId base_entities.ChainId) (*ethclient.Client, error) {
	b.RLock()
	client, has := b.rpcs[chanId]
	b.RUnlock()
	if !has || len(client) == 0 || client[0].Score > b.maxScore {
		return nil, fmt.Errorf("cannot not found the eth client for chain id [%d]", chanId)
	}
	i := rand.Intn(len(client))
	return client[i].Client, nil
}
