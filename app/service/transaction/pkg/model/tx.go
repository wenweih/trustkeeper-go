package model

import (
  "github.com/jinzhu/gorm"
)

// Balance balance related with assets for address
type Tx struct {
  gorm.Model
  TxID      string
  TxType    string
  Address   string
  Asset     string
  Amount    uint64  `gorm:"type:bigint"`
  Confirmations uint16
  BalanceID uint
  Balance   Balance
}
