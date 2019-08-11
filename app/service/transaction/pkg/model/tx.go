package model

import (
  "github.com/jinzhu/gorm"
)

const (
  Pending string = "pending"
  Success string = "success"
  Fail string = "fail"
)

// Tx balance related with assets for address
type Tx struct {
  gorm.Model
  TxID      string
  TxType    string
  Address   string
  Asset     string
  Amount    string
  Confirmations uint16
  BalanceID uint
  Balance   Balance
  ChainName string
  State     string
}
