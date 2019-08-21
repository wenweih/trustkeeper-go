package repository

import (
  "fmt"
  "math"
  "context"
  "strconv"
  "math/big"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
)

const (
  // DepositEthereumComfirmation more than DepositEthereumComfirmation mean deposit successfully
  DepositEthereumComfirmation int64 = 12

  ChainNameEthereum string = "Ethereum"
)

// ERC20Token Ethereum erc20 token info
type ERC20Token struct {
  Symbol   string `json:"Symbol"`
  Name     string `json:"Name"`
  Decimals uint64 `json:"Decimals"`
  Address  string `json:"Address"`
}

func (repo *repo) EthereumSubscribeNewHead(
  ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
  return repo.ethClient.SubscribeNewHead(ctx, ch)
}

func (repo *repo) EthereumBlock(ctx context.Context, number *big.Int) (*types.Block, error) {
  return repo.ethClient.BlockByNumber(ctx, number)
}

func (repo *repo) EthereumBestBlock(ctx context.Context) (*types.Block, error) {
  head, err := repo.ethClient.HeaderByNumber(ctx, nil)
  if err != nil {
    return nil, err
  }
  return repo.EthereumBlock(ctx, head.Number)
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
  decimalsStr := fmt.Sprintf("%f", math.Pow10(int(decimals)))
  decimalsUint64, err := strconv.ParseUint(decimalsStr[:len(decimalsStr)-7], 10, 64)
  if err != nil {
    return nil, err
  }

  return &ERC20Token{
    Symbol: symbol,
    Name: name,
    Address: tokenHex,
    Decimals: decimalsUint64}, nil
}

// QueryEthereumTx query ethereum tx
func (repo *repo) QueryEthereumTxReceipt(ctx context.Context, txid string) (*types.Receipt, error) {
  return repo.ethClient.TransactionReceipt(ctx, common.HexToHash(txid))
}
