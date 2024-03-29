package main

import (
	"context"
	"fmt"
	"os"
	"trustkeeper-go/app/service/wallet_management/client"

	"github.com/caarlos0/env"
	log "github.com/go-kit/kit/log"
)

type config struct {
	ConsulAddr string `env:"consuladdr"`
}

func main() {

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
		struct {
			Roles []string
			UID   string
			NID   string
		}{[]string{"merchant_admin"}, "dab3452a-defe-461d-ae52-c31bced94f7a", "471192264474624001"})

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
	// wallet, err := s.CreateWallet(ctxWithAuthInfo, "470713253851332609", "Bitcoincore", int(1))
	// if err != nil {
	//   logger.Log("CreateWallet error: ", err.Error())
	// }
	// logger.Log("wallet: address", wallet.Address, "id: ", wallet.ID, " status: ", wallet.Status)

	// wallets, err := s.GetWallets(ctxWithAuthInfo, "", 1, 5, 0)
	// if err != nil {
	//   logger.Log("GetWallets without groupid error: ", err.Error())
	// }
	// for _, wallet := range wallets {
	//   for _, w := range wallet.Wallets {
	//     fmt.Println("ChainName: ", wallet.ChainName, " TotalSize: ", wallet.TotalSize, *w)
	//   }
	// }

	walletsForGroupAndChain, err := s.QueryWalletsForGroupByChainName(ctxWithAuthInfo, "471192380250750977", "Ethereum")
	if err != nil {
		logger.Log("GetWallets without groupid error: ", err.Error())
	}

	for _, wallet := range walletsForGroupAndChain {
		fmt.Println("wallet: ", wallet)
	}
	//
	// groupidwallets, err := s.GetWallets(ctx, "469764006120751105")
	// if err != nil {
	//   logger.Log("GetWallets with groupid error: ", err.Error())
	// }
	// for _, wallet := range groupidwallets {
	//   fmt.Println("wallet with groupid: ", *wallet)
	// }

	walletHD, err := s.QueryWalletHD(ctxWithAuthInfo, "0x034e335B7BcfEadD9f4d6fC3AA3A2fEc7E8364E0")
	if err != nil {
		fmt.Println("QueryWalletHD fail:", err.Error())
	}
	fmt.Println("QueryWalletHD:", walletHD)
}
