package main

import (
  "os"
  "github.com/spf13/cobra"
  "github.com/go-kit/kit/log/level"
)

var utxoSet = &cobra.Command {
  Use:   "utxo",
  Short: "maintain utxo set",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {
    case "btc":
    default:
      level.Warn(logger).Log("msg", "Only support utxo base blockchain", "Support", "btc")
      os.Exit(1)
    }
	},
}
