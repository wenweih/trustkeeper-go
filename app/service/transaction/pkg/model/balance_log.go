package model

import (
  "github.com/jinzhu/gorm"
)

// Balance balance related with assets for address
type BalanceLog struct {
  gorm.Model
  TxID      string
  From    uint64  `gorm:"type:bigint"`
  To      uint64  `gorm:"type:bigint"`
  BalanceID uint
  Balance   Balance
}
