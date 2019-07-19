package model

import (
  "github.com/jinzhu/gorm"
)

type Chain struct {
  gorm.Model
  Name  string `gorm:"unique;not null"`
  Coin  string `gorm:"unique;not null"`
  Bip44id int  `gorm:"unique_index;not null"`  // BIP44 公链编号 https://github.com/satoshilabs/slips/blob/master/slip-0044.md
  Status  bool
}
