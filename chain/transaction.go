package chain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
)

func (s Sync) TransactionReceipt(txid string) (*types.Receipt, error) {
	return s.Client.TransactionReceipt(context.Background(), common.HexToHash(txid))
}

func (s Sync) TransactionByHash(txid string) (tx *types.Transaction, isPending bool, err error) {
	return s.Client.TransactionByHash(context.Background(), common.HexToHash(txid))
}

func (s Sync) GetBlockByHash(blockHash string) (*types.Block, error) {
	return s.Client.BlockByHash(context.Background(), common.HexToHash(blockHash))
}

func (s Sync) GetBlockByNumber(blockNumber string) (*types.Block, error) {
	fromString, err := decimal.NewFromString(blockNumber)
	if err != nil {
		return nil, err
	}
	return s.Client.BlockByNumber(context.Background(), fromString.BigInt())
}

func (s Sync) EthCall(from, to, data string) ([]byte, error) {
	toAddr := common.HexToAddress(to)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(from),
		To:   &toAddr,
		Data: []byte(data),
	}
	return s.Client.CallContract(context.Background(), msg, nil)
}
