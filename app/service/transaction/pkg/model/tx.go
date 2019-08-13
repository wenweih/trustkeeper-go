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
  TxID      string      `gorm:"index;not null"`
  TxType    string
  Address   string
  Asset     string
  Amount    string       `gorm:"not null;default:'0'"`
  Confirmations int64
  BalanceID uint        `gorm:"not null"`
  Balance   Balance
  ChainName string
  State     string
}
