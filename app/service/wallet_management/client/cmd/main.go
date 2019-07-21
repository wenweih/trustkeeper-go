package main

import (
  "time"
  "fmt"
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
		struct{Roles []string;UID string;NID string}{[]string{"merchant_admin"}, "1d30be4e-d61e-42da-9cb2-b0d794e12314", "470630148222189569"})
  ctx, cancel := context.WithTimeout(ctxWithAuthInfo, 5*time.Second)
  defer cancel()

  // for _, str := range []string{"aa", "bb", "ccc", "ee", "dd", "ff", "gg", "hh"} {
  //   if err := s.CreateChain(ctxWithAuthInfo, str, "bit44ID", true); err != nil {
  //     logger.Log("CreateChain BTC error: ", err.Error())
  //   }
  // }
  //
  // if err := s.AssignedXpubToGroup(ctxWithAuthInfo, "466126082655944705"); err != nil {
  //   logger.Log("err:", err.Error())
  // }

  // chains, err := s.GetChains(ctx)
  // if err != nil {
  //   logger.Log("Fail to get chains", err.Error())
  // }
  // for _, c := range chains {
  //  logger.Log("chain: id", c.ID, "Name", c.Name, "Coin", c.Coin, "Bip44id", c.Bip44id, "status", c.Status)
  // }
  wallet, err := s.CreateWallet(ctxWithAuthInfo, "470630307152953345", "Bitcoincore", int(1))
  if err != nil {
    logger.Log("CreateWallet error: ", err.Error())
  }
  logger.Log("wallet: address", wallet.Address, "id: ", wallet.ID, " status: ", wallet.Status)

  wallets, err := s.GetWallets(ctx, "470630307152953345")
  if err != nil {
    logger.Log("GetWallets without groupid error: ", err.Error())
  }
  for _, wallet := range wallets {
    fmt.Println("wallet without groupid: ", *wallet)
  }
  //
  // groupidwallets, err := s.GetWallets(ctx, "469764006120751105")
  // if err != nil {
  //   logger.Log("GetWallets with groupid error: ", err.Error())
  // }
  // for _, wallet := range groupidwallets {
  //   fmt.Println("wallet with groupid: ", *wallet)
  // }
}
