package model

import (
  "github.com/jinzhu/gorm"
)

// MnemonicVersion for organization
type MnemonicVersion struct {
  gorm.Model
  Version  string  `gorm:"type:varchar(500);index"`
  Xpubs []Xpub
}
