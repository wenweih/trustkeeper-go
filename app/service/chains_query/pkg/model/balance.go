package model

import (
  "github.com/jinzhu/gorm"
)

// Balance balance related with assets for address
type Balance struct {
  gorm.Model
  Address   string  `gorm:"unique_index:idx_address_symbol;not null"`
  Symbol    string  `gorm:"unique_index:idx_address_symbol;not null"`
  Identify  string
  Decimal   uint64
  Amount    uint64  `gorm:"type:bigint"`
  Txes      []Tx
  BalanceLog []BalanceLog
}
