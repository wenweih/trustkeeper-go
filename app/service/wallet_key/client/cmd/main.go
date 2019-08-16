package main

import (
  "os"
  "fmt"
  "context"
  log "github.com/go-kit/kit/log"
  "github.com/caarlos0/env"
  "trustkeeper-go/app/service/wallet_key/client"
  "trustkeeper-go/app/service/wallet_key/pkg/repository"
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

  // asset := []int32{int32(0), int32(60), int32(66)}
  // chainsWithXpubs, version, err := s.GenerateMnemonic(context.Background(), "aaaaa", asset, 10)
  // if err != nil {
  //   logger.Log("GenerateMnemonic:", err.Error())
  // }
  //
  // for _, chainwithXpubs := range chainsWithXpubs {
  //   for _, xpub := range chainwithXpubs.Xpubs {
  //     xpubTmp := xpub
  //    logger.Log("chain:", chainwithXpubs.Chain, " xpub account:", xpubTmp.Account, " xpub key:", xpubTmp.Key, "version:", version)
  //   }
  // }
  walletHD := repository.WalletHD{
    CoinType: 0,
    Account: 0,
    Change: 0,
    AddressIndex: 46,
    MnemonicVersion: "471192264474624001/2019-07-23 15:27:51.530679",
  }
  signedTxHex, err := s.SignedBitcoincoreTx(context.Background(), walletHD,
    "01000000016ba0feb402a9ac1cb1c65c58d6da59815130abdbfbc211512014ff46ff4f1de70100000000ffffffff0280969800000000001976a9144e1eff3ae7f5d38921b238469137c4f4f6f14f0488ac50d55c05000000001976a9144e1eff3ae7f5d38921b238469137c4f4f6f14f0488ac00000000",
    100000000)
  if err != nil {
    fmt.Println("SignedBitcoincoreTx error:", err.Error())
  }
  fmt.Println("SignedBitcoincoreTx:", signedTxHex)
}
