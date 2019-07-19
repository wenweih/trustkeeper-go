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
		struct{Roles []string;UID string;NID string}{[]string{"merchant_admin"}, "d7fd20c1-cfb8-461f-9a95-a3d028f20e35", "469763787409293313"})

  // for _, str := range []string{"aa", "bb", "ccc", "ee", "dd", "ff", "gg", "hh"} {
  //   if err := s.CreateChain(ctxWithAuthInfo, str, "bit44ID", true); err != nil {
  //     logger.Log("CreateChain BTC error: ", err.Error())
  //   }
  // }
  //
  // if err := s.AssignedXpubToGroup(ctxWithAuthInfo, "466126082655944705"); err != nil {
  //   logger.Log("err:", err.Error())
  // }

  chains, err := s.GetChains(ctxWithAuthInfo)
  if err != nil {
    logger.Log("Fail to get chains", err.Error())
  }
  for _, c := range chains {
   logger.Log("chain: id", c.ID, "Name", c.Name, "Coin", c.Coin, "Bip44id", c.Bip44id, "status", c.Status)
  }
  if err := s.CreateWallet(ctxWithAuthInfo, "469764006120751105", "Ethereum", int(1)); err != nil {
    logger.Log("CreateWallet error: ", err.Error())
  }
}
