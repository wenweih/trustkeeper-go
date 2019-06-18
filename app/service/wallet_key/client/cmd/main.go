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

  for _, str := range []string{"aa", "bb", "ccc", "ee", "dd", "ff", "gg", "hh"} {
    xpub, err := s.GenerateMnemonic(context.Background(), str)
    if err != nil {
      logger.Log("GenerateMnemonic:", err.Error())
    }
    logger.Log("xpub:", xpub)
  }
}
