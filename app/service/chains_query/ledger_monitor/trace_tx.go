package main

import (
  "context"
  "os"
  "bytes"
  "sync"
  "encoding/gob"
  "encoding/json"
  "github.com/spf13/cobra"
  "github.com/streadway/amqp"
  common "trustkeeper-go/library/util"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "trustkeeper-go/app/service/chains_query/pkg/model"
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
      // EthereumBestBlock
      ctx := context.Background()
      bestBlock, err := svc.EthereumBestBlock(ctx)
      if err != nil {
        logger.Log("EthereumBestBlock", err.Error())
      }
      if err := svc.CreateETHBlockWithTx(ctx, bestBlock.Number().Int64()); err != nil {
        logger.Log("CreateETHBlockWithTx", err.Error())
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
  err := svc.MQSubscribe("bestblock", "direct", "ethereum_best_block_queue",
    "ethereum", "eth", onEthMessage)
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
  logger.Log("blockhash", mqdata.Hash.String(), " height:", mqdata.Header.Number.String())
  for _, tx := range mqdata.Tx {
    v, err := hexutil.DecodeBig(tx.ValueHex)
    if err != nil {
      logger.Log("err:", err.Error())
    }
    logger.Log("tx", tx.THash, " value:", v.Int64(), " input data:", tx.Data)
  }
}

func bitcoinMQ(wg *sync.WaitGroup)  {
  defer wg.Done()
  forever := make(chan bool)
  err := svc.MQSubscribe("bestblock", "direct", "bitcoincore_best_block_queue",
    "bitcoincore", "btc", onBitcoinMessage)
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
