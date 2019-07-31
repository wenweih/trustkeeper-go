package main

import (
  "os"
  "sync"
  "encoding/json"
  "github.com/spf13/cobra"
  "github.com/streadway/amqp"
  common "trustkeeper-go/library/util"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/ethereum/go-ethereum/core/types"
)

var traceTx = &cobra.Command {
  Use:   "trace-tx",
  Short: "Trace blockchain tx",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {

    case "bitcoincore":
      var wg sync.WaitGroup
      wg.Add(1)
      go bitcoinMQ(&wg)
      wg.Wait()
    case "ethereum":
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
	var mqdata *types.Block
	err := json.Unmarshal(d.Body, &mqdata)
  if err != nil {
    logger.Log("EthereumBlockUnmarshalError", err.Error())
  }
  logger.Log("blockhash", mqdata.Hash().String())
  for _, tx := range mqdata.Body().Transactions {
    logger.Log("tx", tx.Hash().String())
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
	var mqdata *btcjson.GetBlockVerboseResult
	if err := json.Unmarshal(d.Body, &mqdata); err != nil {
    logger.Log("GetBlockVerboseResultUnmarshalError", err.Error())
  }
  logger.Log("blockhash", mqdata.Hash)
}
