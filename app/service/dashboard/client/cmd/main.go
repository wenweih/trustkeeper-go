package main

import (
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
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
  _, err = s.CreateGroup(context.Background(), "546e1345-4c4c-44c9-9baf-04f3cdc908ec", "testGroup", "desc", uint(462832467565871105))
  if err != nil {
    logger.Log("CreateGroup error: ", err.Error())
  }

  for _, namespaceID := range []uint{uint(462832467565871105)} {
    groups, err := s.GetGroups(context.Background(), namespaceID)
    if err != nil {
      logger.Log("GetGroups error: ", err.Error())
    }
    for _, g := range groups {
      logger.Log("Group:", g.Name)
    }
  }
}
