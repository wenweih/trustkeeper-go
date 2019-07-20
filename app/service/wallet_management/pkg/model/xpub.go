package model

import (
  "github.com/jinzhu/gorm"
  "github.com/qor/transition"
)

// Xpub for organization
type Xpub struct {
  gorm.Model
  Key           string    `gorm:"type:varchar(500);index"`
  Bip44ChainID  uint
  Chain         Chain  `gorm:"foreignkey:Bip44ChainID;association_foreignkey:Bip44id"`
  Bip44Account  int
  GroupID       string    `gorm:"index"`
  MnemonicVersionID uint
  MnemonicVersion MnemonicVersion
  transition.Transition
  Wallets []Wallet  `gorm:"foreignkey:XpubUID"`
}
