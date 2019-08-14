package main

import (
  "fmt"
  "os"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/chains_query/client"
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

  property, err := s.QueryOmniProperty(ctxWithAuthInfo, int64(2147483651))
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println("property:", property)

  token, err := s.ERC20TokenInfo(ctxWithAuthInfo, "0xf0680d66aac362b1e42e21d3098ad61e92c6f43f")
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println("token:", token)

  rawTxHex, err := s.ConstructTxBTC(ctxWithAuthInfo, "mne28j3A5ht5yp8LtptHwuFRQfzhDS1YiH", "mjDh2U12TioqP7hUXU2vBagU6Z5R4y9Pbj", "0.1")
  if err != nil {
    fmt.Println("ConstructTxBTC:", err.Error())
  }
  fmt.Println("ConstructTxBTC:", rawTxHex)
}
