package main

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var blockMonitor = &cobra.Command{
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

			// maintain orderHeight and increase 1 each subscribe callback, because head.number would jump blocks
			var orderHeight = new(big.Int)
			defer sub.Unsubscribe()
			for {
				select {
				case err := <-sub.Err():
					logger.Log("sub:", err.Error())
				case head := <-blockCh:
					ordertmp, err := subHandle(orderHeight, head)
					if err != nil {
						logger.Log("Ethereum subscribe handle error", err.Error())
					}
					orderHeight = ordertmp
				}
			}
		case "eosio":
		default:
			logger.Log("Only support bitcoincore, ethereum, eosio")
			os.Exit(1)
		}
	},
}
