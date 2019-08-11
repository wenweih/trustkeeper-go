package main

import (
  "os"
  "bytes"
  "sync"
  "strings"
  "context"
  "encoding/gob"
  "encoding/json"
  "github.com/spf13/cobra"
  "github.com/streadway/amqp"
  common "trustkeeper-go/library/util"
  "github.com/btcsuite/btcd/btcjson"
  "trustkeeper-go/app/service/chains_query/pkg/model"
  "trustkeeper-go/app/service/chains_query/pkg/repository"
)

var traceTx = &cobra.Command {
  Use:   "trace-tx",
  Short: "Trace blockchain tx",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {
    case "bitcoincore":
      // query ledger info
      ctx := context.Background()
      ledgerInfo, err := svc.QueryBTCLedgerInfo(ctx)
      if err != nil {
        logger.Log("QueryBTCLedgerInfo", err.Error())
      }

      createBlockResul := <- svc.CreateBtcBlockWithUtxoPipeline(ctx, int64(ledgerInfo.Headers - 5))
      if createBlockResul.Error != nil{
        logger.Log("CreateBtcBlockWithUtxoPipeline", createBlockResul.Error.Error())
      }

      bestBlock := createBlockResul.Block
      logger.Log("CreateBlock", bestBlock.Height, "Hash", bestBlock.Hash)

      isTracking := true
      trackHeight := bestBlock.Height - 1
      for isTracking {
        isTracking, trackHeight = svc.TrackBtcBlockPipeline(
          ctx, trackHeight, bestBlock.Height, isTracking)
      }

      dbBestHeight := bestBlock.Height
      for height := (dbBestHeight + 1); height <= int64(ledgerInfo.Headers); height++ {
        createBlockResul := <- svc.CreateBtcBlockWithUtxoPipeline(ctx, height)
        if createBlockResul.Error != nil{
          logger.Log("createBlockResulError", createBlockResul.Error.Error())
        }
      }

      var wg sync.WaitGroup
      wg.Add(1)
      go bitcoinMQ(&wg)
      wg.Wait()
    case "ethereum":
      ctx := context.Background()
      bestBlock, err := svc.EthereumBestBlock(ctx)
      if err != nil {
        logger.Log("EthereumBestBlock", err.Error())
      }
      bestHeight := bestBlock.Number().Int64()

      dbBestBlock, err := svc.EthereumDBBestBlock(ctx)
      if err != nil && !strings.Contains(err.Error(), "record not found"){
        logger.Log("EthereumDBBestBlock", err.Error())
        os.Exit(0)
      }

      minHeight := int64(0)
      if dbBestBlock == nil {
        minHeight = bestHeight - 12
      } else {
        minHeight = dbBestBlock.Height - 12
      }

      for traceHeight := bestHeight; traceHeight > minHeight; traceHeight-- {
        block, err := svc.CreateETHBlockWithTx(ctx, traceHeight)
        if err != nil {
          logger.Log("CreateETHBlockWithTx", err.Error())
        }else {
          logger.Log("TraceBlock", traceHeight, "Hash", block.Hash)
        }
      }
      var wg sync.WaitGroup
    	wg.Add(1)
    	go ethReceive(&wg)
    	wg.Wait()
    	// Makes sure connection is closed when service exits.
    	common.HandleSigterm(func() {
    		// if messageClient != nil {
    		// 	messageClient.Close()
    		// }
    	})
    case "eosio":
    default:
      logger.Log("Only support:", "bitcoincore, ethereum, eosio")
      os.Exit(1)
    }
	},
}

func ethReceive(wg *sync.WaitGroup) {
	defer wg.Done()
	forever := make(chan bool)
  err := svc.MQSubscribe(repository.ExchangeNameForEthereumBestBlock,
    "fanout",
    repository.QueueNameForEthereumTraceTx,
    repository.BindKeyEthereum, "trace_tx_eth", onEthMessage)
  if err != nil {
    logger.Log("ethReceiveFail", err.Error())
  }
	<-forever
}

func onEthMessage(d amqp.Delivery) {
	mqdata := model.EthereumBlock{}
  buf := bytes.NewBuffer(d.Body)
  dc := gob.NewDecoder(buf)
  err := dc.Decode(&mqdata)
  if err != nil {
    logger.Log("EthereumBlockReadError", err.Error())
    return
  }
  blockHeight := mqdata.Header.Number.Int64()
  for traceHeight := blockHeight; traceHeight > blockHeight -12; traceHeight-- {
    block, err := svc.CreateETHBlockWithTx(context.Background(), traceHeight)
    if err != nil {
      logger.Log("CreateETHBlockWithTx", err.Error())
    }else {
      logger.Log("TraceBlock", traceHeight, "Hash", block.Hash)
    }
  }
}

func bitcoinMQ(wg *sync.WaitGroup)  {
  defer wg.Done()
  forever := make(chan bool)
  err := svc.MQSubscribe(
    repository.ExchangeNameForBitcoincoreBestBlock,
    "fanout",
    repository.QueueNameForBitcoincoreTraceTx,
    repository.BindKeyBitcoincore, "trace_tx_btc", onBitcoinMessage)
  if err != nil {
    logger.Log("bitcoinMQ", err.Error())
  }
  <-forever
}

func onBitcoinMessage(d amqp.Delivery) {
  ctx := context.Background()
	var mqdata *btcjson.GetBlockVerboseResult
	if err := json.Unmarshal(d.Body, &mqdata); err != nil {
    logger.Log("GetBlockVerboseResultUnmarshalError", err.Error())
    return
  }
  logger.Log("Consumer Bitcoin New Block", mqdata.Hash)

  createBlockResult := <- svc.CreateBTCBlockWithUTXOs(ctx, mqdata)
  if createBlockResult.Error != nil{
    logger.Log("CreateBTCBlockWithUTXOsError", createBlockResult.Error.Error())
  }

  isTracking := true
  trackHeight := mqdata.Height - 1
  for isTracking {
    isTracking, trackHeight = svc.TrackBtcBlockPipeline(
      ctx, trackHeight, mqdata.Height, isTracking)
  }
}
