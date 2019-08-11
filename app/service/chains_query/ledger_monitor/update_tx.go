package main

import (
  "os"
  "bytes"
  "sync"
  "encoding/gob"
  "encoding/json"
  "github.com/spf13/cobra"
  "github.com/streadway/amqp"
  common "trustkeeper-go/library/util"
  "github.com/btcsuite/btcd/btcjson"
  "trustkeeper-go/app/service/chains_query/pkg/model"
  "trustkeeper-go/app/service/chains_query/pkg/repository"
)

var updateTx = &cobra.Command {
  Use:   "update-tx",
  Short: "Update system tx",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {
    case "bitcoincore":
      var wg sync.WaitGroup
      wg.Add(1)
      go bitcoinUpdateTxMQ(&wg)
      wg.Wait()
    case "ethereum":
      var wg sync.WaitGroup
    	wg.Add(1)
    	go ethUpdateTxReceive(&wg)
    	wg.Wait()
    	common.HandleSigterm(func() {
    	})
    case "eosio":
    default:
      logger.Log("Only support:", "bitcoincore, ethereum, eosio")
      os.Exit(1)
    }
	},
}

func ethUpdateTxReceive(wg *sync.WaitGroup) {
	defer wg.Done()
	forever := make(chan bool)
  err := svc.MQSubscribe(
    repository.ExchangeNameForEthereumBestBlock,
    "fanout",
    repository.QueueNameForEthereumUpdateTx,
    repository.BindKeyEthereum, "update_tx_eth", onEthUpdateTxMessage)
  if err != nil {
    logger.Log("ethReceiveFail", err.Error())
  }
	<-forever
}

func onEthUpdateTxMessage(d amqp.Delivery) {
	mqdata := model.EthereumBlock{}
  buf := bytes.NewBuffer(d.Body)
  dc := gob.NewDecoder(buf)
  err := dc.Decode(&mqdata)
  if err != nil {
    logger.Log("EthereumBlockReadError", err.Error())
    return
  }
  logger.Log("Ethereum", mqdata)
}

func bitcoinUpdateTxMQ(wg *sync.WaitGroup)  {
  defer wg.Done()
  forever := make(chan bool)
  err := svc.MQSubscribe(
    repository.ExchangeNameForBitcoincoreBestBlock,
    "fanout",
    repository.QueueNameForBitcoincoreUpdateTx,
    repository.BindKeyBitcoincore, "update_tx_btc", onBitcoinUpdateTxMessage)
  if err != nil {
    logger.Log("bitcoinMQ", err.Error())
  }
  <-forever
}

func onBitcoinUpdateTxMessage(d amqp.Delivery) {
	var mqdata *btcjson.GetBlockVerboseResult
	if err := json.Unmarshal(d.Body, &mqdata); err != nil {
    logger.Log("GetBlockVerboseResultUnmarshalError", err.Error())
    return
  }
  logger.Log("Consumer Bitcoin New Block", mqdata.Hash)
}
