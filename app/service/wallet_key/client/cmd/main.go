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
    CoinType: 60,
    Account: 0,
    Change: 0,
    AddressIndex: 17,
    MnemonicVersion: "471192264474624001/2019-07-23 15:27:51.530679",
  }
  // signedTxHex, err := s.SignedBitcoincoreTx(context.Background(), walletHD,
  //   "01000000023428ed5680af8b5766573665a45b1af0b9dfa91837ca9024dd6c8b6c2e9458390100000000ffffffff7d41315bfca7eca7ac8be2c3dbd1ef315c8b11baa1b85fb585f7faf0defd1def0000000000ffffffff0280d1f008000000001976a914289cafa2615c1bc74369e86ea5d08427a063abcf88ac9cb9c304000000001976a9144e1eff3ae7f5d38921b238469137c4f4f6f14f0488ac00000000",
  //   100000000)
  // if err != nil {
  //   fmt.Println("SignedBitcoincoreTx error:", err.Error())
  // }
  // fmt.Println("SignedBitcoincoreTx:", signedTxHex)
  signedETHTXHex, err := s.SignedEthereumTx(context.Background(), walletHD, "0xe78001825208942cf0bbc1244f7957627c3df49d5b9e79f7b95e9688016345785d8a000080808080", "1337")
  if err != nil {
    fmt.Println("SignedEthereumTx error:", err.Error())
  }
  fmt.Println("SignedEthereumTx:", signedETHTXHex)
}
