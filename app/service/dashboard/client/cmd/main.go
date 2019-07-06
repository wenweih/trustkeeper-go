package main

import (
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "github.com/Pallinder/go-randomdata"
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
		struct{Roles []string;UID string;NID string}{[]string{"merchant_admin"}, "466361632363970561", "466361632420134913"})

  var namespaceID = "466361632420134913"
  group, err := s.CreateGroup(ctxWithAuthInfo, "55ee782d-4404-435c-b587-e5cf5ecc7da1", randomdata.SillyName(), "deeeeee", namespaceID)
  if err != nil {
    logger.Log("CreateGroup error: ", err.Error())
  }

  if err := s.UpdateGroup(ctxWithAuthInfo, group.ID, "changename", "changedesc"); err != nil {
    logger.Log("change group err: ", err.Error())
  }

  // for _, namespaceID := range []string{namespaceID} {
  //   groups, err := s.GetGroups(ctxWithAuthInfo, namespaceID)
  //   if err != nil {
  //     logger.Log("GetGroups error: ", err.Error())
  //   }
  //   for _, g := range groups {
  //     logger.Log("Get Group:", g.Name)
  //     if err := s.UpdateGroup(ctxWithAuthInfo, g.ID, "changename", "changedesc"); err != nil {
  //       logger.Log("change group err: ", err.Error())
  //     }
  //   }
  // }

}
