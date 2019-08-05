package repository

import (
  "context"
  "math/big"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
)

// ERC20Token Ethereum erc20 token info
type ERC20Token struct {
  Symbol   string `json:"Symbol"`
  Name     string `json:"Name"`
  Decimals uint32 `json:"Decimals"`
}

func (repo *repo) EthereumSubscribeNewHead(
  ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
  return repo.ethClient.SubscribeNewHead(ctx, ch)
}

func (repo *repo) EthereumBlock(ctx context.Context, number *big.Int) (*types.Block, error) {
  return repo.ethClient.BlockByNumber(ctx, number)
}

func (repo *repo) ERC20TokenInfo(ctx context.Context, tokenHex string) (*ERC20Token, error) {
  tokenAddress := common.HexToAddress(tokenHex)
  token, err := NewETHToken(tokenAddress, repo.ethClient)
  if err != nil {
    return nil, err
  }

  symbol, err := token.Symbol(nil)
  if err != nil {
    return nil, err
  }

  name, err := token.Name(nil)
  if err != nil {
    return nil, err
  }

  decimals, err := token.Decimals(nil)
  if err != nil {
    return nil, err
  }

  return &ERC20Token{
    Symbol: symbol,
    Name: name,
    Decimals: uint32(decimals)}, nil
}
