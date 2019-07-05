package model

import (
  "github.com/jinzhu/gorm"
  "github.com/qor/transition"
)

// Xpub for organization
type Xpub struct {
  gorm.Model
  Key           string    `gorm:"type:varchar(500);index"`
  Bip44ChainID  int
  BIP44Account  int
  GroupID       string    `gorm:"index"`
  MnemonicVersionID uint
  MnemonicVersion MnemonicVersion
  transition.Transition
}
