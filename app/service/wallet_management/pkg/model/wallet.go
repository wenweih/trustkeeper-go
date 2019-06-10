package model

import (
  "github.com/jinzhu/gorm"
)

type Wallet struct {
  gorm.Model
  BIP44Change int
  Address     string
  BIP44Index  string
  XpubUID     string
  Status      bool
}
