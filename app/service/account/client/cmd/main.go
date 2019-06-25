package main

import (
  "os"
  "context"
  "strings"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/account/client"
  "github.com/Pallinder/go-randomdata"
  stdjwt "github.com/go-kit/kit/auth/jwt"
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

  email := randomdata.Email()
  password := "111111"
  orgName := randomdata.SillyName()

  s, err := client.New(cfg.ConsulAddr, logger)
  if err != nil {
    logger.Log("service client error: ", err.Error())
  }

  uuid, err := s.Create(context.Background(), email, password, orgName)
  if err != nil {
    logger.Log("Create error: ", err.Error())
  }

  logger.Log("uuid: ", uuid)

  token, err := s.Signin(context.Background(), email, password)
  if err != nil || len(token) == 0{
    logger.Log("Signin error: ", err.Error())
  }
  logger.Log("token: ", token)

  roles, err := s.Roles(context.WithValue(context.Background(), stdjwt.JWTTokenContextKey , token))
  if err != nil {
    logger.Log("Roles error: ", err.Error())
  }
  logger.Log("Roles:", strings.Join(roles," "))

  _, _, err = s.Auth(context.WithValue(context.Background(), stdjwt.JWTTokenContextKey , token))
  if err != nil {
    logger.Log("Auth error: ", err.Error())
  }
  //
  if err := s.Signout(context.WithValue(context.Background(), stdjwt.JWTTokenContextKey , token)); err != nil {
    logger.Log("Signout error: ", err.Error())
  }

}
