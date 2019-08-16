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

  rawTxHex, vinAmount, err := s.ConstructTxBTC(ctxWithAuthInfo, "mne28j3A5ht5yp8LtptHwuFRQfzhDS1YiH", "mjDh2U12TioqP7hUXU2vBagU6Z5R4y9Pbj", "1.5")
  if err != nil {
    fmt.Println("ConstructTxBTC:", err.Error())
  }
  fmt.Println("ConstructTxBTC:", rawTxHex, vinAmount)

  txid, err := s.SendBTCTx(ctxWithAuthInfo, "01000000023428ed5680af8b5766573665a45b1af0b9dfa91837ca9024dd6c8b6c2e945839010000006b483045022100e3680c1f97b7313710b46b5e2929674013ae71c993e3b7b4dfd12d64a1399a3002206e830b45f595f224901d59083f00c248d05c7ca990aef5a6d34ea0896254c53b012102b6ca80a3a74bbe371c816fda2fbd3ee31962418660ac7014a8e0a3813e1f4de4ffffffff7d41315bfca7eca7ac8be2c3dbd1ef315c8b11baa1b85fb585f7faf0defd1def000000006b483045022100f0e13bf1017b272f40fc2771ef04db8ae9f01ea906911d398b925d3c6a08d56f02200211f8259646963cf5d60cfb50587957bf784ea2693c7179868a77e809b952a3012102b6ca80a3a74bbe371c816fda2fbd3ee31962418660ac7014a8e0a3813e1f4de4ffffffff0280d1f008000000001976a914289cafa2615c1bc74369e86ea5d08427a063abcf88ac9cb9c304000000001976a9144e1eff3ae7f5d38921b238469137c4f4f6f14f0488ac00000000")
  if err != nil {
    fmt.Println("SendBTCTx err:", err.Error())
  }
  fmt.Println("SendBTCTx:", txid)
}
