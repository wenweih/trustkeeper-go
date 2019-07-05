package main

import (
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/wallet_management/client"
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
  ctxWithAuthInfo := context.WithValue(context.Background(), "auth",
		struct{Roles []string;UID string;NID string}{[]string{"merchant_admin"}, "466054237439852545", "466054237547266049"})

  for _, str := range []string{"aa", "bb", "ccc", "ee", "dd", "ff", "gg", "hh"} {
    if err := s.CreateChain(ctxWithAuthInfo, str, "bit44ID", true); err != nil {
      logger.Log("CreateChain BTC error: ", err.Error())
    }
  }

  if err := s.AssignedXpubToGroup(ctxWithAuthInfo, "466126082655944705"); err != nil {
    logger.Log("err:", err.Error())
  }
}
