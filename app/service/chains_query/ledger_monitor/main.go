package main

import (
  "os"
  "trustkeeper-go/app/service/chains_query/pkg/configure"
  "github.com/spf13/cobra"
  "github.com/ethereum/go-ethereum/ethclient"
  log "github.com/go-kit/kit/log"
  service "trustkeeper-go/app/service/chains_query/pkg/service"
)

var (
  err error
  chain	string
  ethereumClient *ethclient.Client
  conf   *configure.Conf
  logger log.Logger
  svc    service.LedgerMonitorService
)

var rootCmd = &cobra.Command {
	Use:   "ledger_monitor",
	Short: "Blockchain ledger monitor",
}

func main() {
  logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

  conf, err = configure.New()
  if err != nil {
    logger.Log("configure err: ", err.Error())
    os.Exit(1)
  }
  svc, err = service.NewLedgerMonitorService(*conf, logger)
  if err != nil {
    logger.Log("svc error: ", err.Error())
    os.Exit(1)
  }
  if err := rootCmd.Execute(); err != nil {
    logger.Log("Command execute error:", err)
    os.Exit(1)
  }
}

func init()  {
  rootCmd.AddCommand(blockMonitor)
  rootCmd.AddCommand(traceTx)
  rootCmd.AddCommand(updateTx)
  rootCmd.PersistentFlags().StringVarP(&chain, "chain", "c", "", "Support bitcoincore, ethereum")
  rootCmd.MarkFlagRequired("chain")
}
