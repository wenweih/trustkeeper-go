package main

import (
  "os"
  "context"
  // "math/big"
  // "trustkeeper-go/app/service/chains_query/pkg/configure"
  "github.com/spf13/cobra"
  "github.com/gin-gonic/gin"
  "github.com/ethereum/go-ethereum/core/types"
  // "github.com/ethereum/go-ethereum/ethclient"
)

var blockMonitor = &cobra.Command {
  Use:   "best-block",
  Short: "Best Block monitor",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {
    case "bitcoincore":
      gin.SetMode(gin.ReleaseMode)
      r := gin.Default()
      r.GET("/btc-best-block-notify", btcBestBlockNotifyHandle)
      if err := r.Run(":3001"); err != nil {
        logger.Log(err.Error())
        os.Exit(1)
      }
    case "ethereum":
      blockCh := make(chan *types.Header, 16)
      sub, err := svc.EthereumSubscribeNewHead(context.Background(), blockCh)
      if err != nil {
        logger.Log(err.Error())
        os.Exit(1)
      }
      defer sub.Unsubscribe()
      for {
        select {
        case err := <-sub.Err():
          logger.Log("sub:", err.Error())
        case head := <-blockCh:
          logger.Log(head.Hash().String())
        }
      }
    case "eosio":
    default:
      logger.Log("Only support bitcoincore, ethereum, eosio")
      os.Exit(1)
    }
	},
}
