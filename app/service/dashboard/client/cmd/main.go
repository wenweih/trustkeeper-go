package main

import (
  "fmt"
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  // "github.com/Pallinder/go-randomdata"
  "trustkeeper-go/app/service/dashboard/client"
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
		struct{Roles []string;UID string;NID string}{[]string{"merchant_admin"}, "dab3452a-defe-461d-ae52-c31bced94f7a", "471192264474624001"})
  //
  // var namespaceID = "466361632420134913"
  // group, err := s.CreateGroup(ctxWithAuthInfo, "55ee782d-4404-435c-b587-e5cf5ecc7da1", randomdata.SillyName(), "deeeeee", namespaceID)
  // if err != nil {
  //   logger.Log("CreateGroup error: ", err.Error())
  // }
  //
  // if err := s.UpdateGroup(ctxWithAuthInfo, group.ID, randomdata.SillyName(), "changedesc"); err != nil {
  //   logger.Log("change group err: ", err.Error())
  // }

  // 468348259016146945  468348353832910849
  chainAssets, err := s.GetGroupAssets(ctxWithAuthInfo, "471192325366218753")
  if err != nil {
    logger.Log("fail to GetGroupAsset", err.Error())
  }
  for _, ca := range chainAssets {
    logger.Log("chainid: ", ca.ChainID, "coin: ", ca.Coin, "Name", ca.Name, "Status", ca.Status)
    for _, asset := range ca.SimpleAssets {
      logger.Log("status", asset.Status, "Symbol", asset.Symbol, "TokenID", asset.AssetID)
    }
  }

  asset, err := s.AddAsset(ctxWithAuthInfo, "471192325366218753", "471868166221627393", "Omni tokens", "1", "100000000")
  if err != nil {
    logger.Log("AddAssetError", err.Error())
  }
  fmt.Println("AddedAsset: ", asset)

  // if _, err := s.ChangeGroupAssets(ctxWithAuthInfo, chainAssets, "468348259016146945"); err != nil {
  //   logger.Log("fail to ChangeGroupAssets", err.Error())
  // }
}
