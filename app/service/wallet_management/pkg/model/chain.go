package model

import (
  "github.com/jinzhu/gorm"
)

type Chain struct {
  gorm.Model
  Name  string
  Coin  string
  Bip44id int // BIP44 公链编号 https://github.com/satoshilabs/slips/blob/master/slip-0044.md
  Status  bool
}
