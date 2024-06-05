package chain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Sync struct {
	ChainId int64
	Client  *ethclient.Client
	Status  syncStatus
	Url     string
}

type syncStatus struct {
	Status       bool
	LatestNumber *big.Int
	Msg          string
}

func InitSyncer(url string, chainId int64) (*Sync, error) {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Sync{Client: ethClient, ChainId: chainId, Status: syncStatus{Status: false, Msg: "暂未初始化同步状态"}, Url: url}, nil
}
