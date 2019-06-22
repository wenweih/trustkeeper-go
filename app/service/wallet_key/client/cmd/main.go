package main

import (
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/wallet_key/client"
)

type config struct {
  ConsulAddr		string	`env:"consuladdr"`
}

func main()  {

  var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

  cfg := config{}
  if err := env.Parse(&cfg); err != nil {
    logger.Log("fail to parse env: ", err.Error())
    os.Exit(1)
  }

  s, err := client.New(cfg.ConsulAddr, logger)
  if err != nil {
    logger.Log("service client error: ", err.Error())
  }

  asset := []int32{int32(0), int32(60), int32(66)}
  chainsWithXpubs, version, err := s.GenerateMnemonic(context.Background(), "aaaaa", asset, 10)
  if err != nil {
    logger.Log("GenerateMnemonic:", err.Error())
  }

  for _, chainwithXpubs := range chainsWithXpubs {
    for _, xpub := range chainwithXpubs.Xpubs {
      xpubTmp := xpub
     logger.Log("chain:", chainwithXpubs.Chain, " xpub account:", xpubTmp.Account, " xpub key:", xpubTmp.Key, "version:", version)
    }
  }
}
