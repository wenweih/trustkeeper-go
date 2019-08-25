package repository

import (
  "context"
  "math/big"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum"
  "github.com/streadway/amqp"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

type IBiz interface {
  MQPublish(msg []byte, exchangeName, exchangeType, bindingKey string) error
  MQSubscribe(
    exchangeName, exchangeType, queueName,
    bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error
  QueryBitcoincoreBlock(
    ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
  QueryBTCBlockCH(ctx context.Context, height int64) (<-chan UTXOBlockResult)
  TrackBlock(
    ctx context.Context, bestBlockHeight int64, isTracking bool, queryBlockResultCh <- chan UTXOBlockResult) (bool, int64)
  GetBTCBlockHashByHeight(ctx context.Context, height int64) (*chainhash.Hash, error)
  QueryBTCLedgerInfo(ctx context.Context) (*btcjson.GetBlockChainInfoResult, error)
  EthereumSubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
  EthereumBlock(ctx context.Context, number *big.Int) (*types.Block, error)
  CreateBtcBlockWithUtxoPipeline(ctx context.Context, height int64) (<-chan CreateBlockResult)
  CreateBTCBlockWithUTXOs(ctx context.Context, queryBlockResultCh <- chan UTXOBlockResult) (<-chan CreateBlockResult)
  QueryOmniProperty(propertyid int64) (*OmniProperty, error)

  ERC20TokenInfo(ctx context.Context, tokenHex string) (*ERC20Token, error)
  CreateETHBlockWithTx(ctx context.Context, height int64) (*model.EthBlock, error)
  EthereumBestBlock(ctx context.Context) (*types.Block, error)
  EthererumDBBestBlock(ctx context.Context) (*model.EthBlock, error)

  UpdateEthereumTx(ctx context.Context)
  UpdateBitcoincoreTx(ctx context.Context)

  ConstructTxBTC(ctx context.Context, from, to, amount string) (unsignedTxHex string, vinAmount int64, err error)
  SendBTCTx(ctx context.Context, signedTxHex string) (txID string, err error)

  ConstructTxETH(ctx context.Context, from, to, amount string) (unsignedTxHex, chainID string, err error)

  WalletValidate(ctx context.Context, chainName, address string) (err error)
  QueryBalance(ctx context.Context, symbol, address string) (balance string, err error)
}
