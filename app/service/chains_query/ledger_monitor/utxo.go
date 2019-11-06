package main

import (
  "os"
  "github.com/spf13/cobra"
)

var utxoSet = &cobra.Command {
  Use:   "utxo",
  Short: "maintain utxo set",
  Run: func(cmd *cobra.Command, args []string) {
    switch chain {
    case "bitcoincore":
    default:
      logger.Log("Only support utxo base blockchain:", "bitcoincore")
      os.Exit(1)
    }
	},
}
