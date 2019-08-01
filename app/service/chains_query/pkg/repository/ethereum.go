package repository

import (
  "context"
  "math/big"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/core/types"
)

func (repo *repo) EthereumSubscribeNewHead(
  ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
  return repo.ethClient.SubscribeNewHead(ctx, ch)
}

func (repo *repo) EthereumBlock(ctx context.Context, number *big.Int) (*types.Block, error) {
  return repo.ethClient.BlockByNumber(ctx, number)
}
