package main

import (
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/account/client"
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

  email := "mr.huangwenwei@gmail.com"
  password := "111111"
  orgName := "test"

  s, err := client.New(cfg.ConsulAddr, logger)
  if err != nil {
    logger.Log("service client error: ", err.Error())
  }

  uuid, err := s.Create(context.Background(), email, password, orgName)
  if err != nil {
    logger.Log("Create error: ", err.Error())
  }

  token, err := s.Signin(context.Background(), email, password)
  if err != nil {
    logger.Log("Signin error: ", err.Error())
  }

  logger.Log("uuid: ", uuid, " token: ", token)

}
