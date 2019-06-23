package model

import (
  "github.com/jinzhu/gorm"
)

// Xpub for organization
type Xpub struct {
  gorm.Model
  Key           string    `gorm:"type:varchar(500);index"`
  Status        bool      `gorm:"default:true"`
  Bip44ChainID  int
  BIP44Account  int
  GroupID       int8
  MnemonicVersionID uint
  MnemonicVersion MnemonicVersion
}
