package model

import (
  "github.com/jinzhu/gorm"
)

const (
  StatePending string = "pending"
  StateConfirming string = "confirming"
  StateSuccess string = "success"
  StateFail string = "fail"

  ChainBitcoin string = "Bitcoincore"
  ChainEthereum string = "Ethereum"
)

// Tx balance related with assets for address
type Tx struct {
  gorm.Model
  TxID      string
  TxType    string
  Address   string
  Asset     string
  Amount    string
  Confirmations int64
  BalanceID uint
  Balance   Balance
  ChainName string
  State     string
}
