package rpc

import "github.com/ethereum/go-ethereum/ethclient"

type BaseProvider struct {
	Rpc *ethclient.Client
}
