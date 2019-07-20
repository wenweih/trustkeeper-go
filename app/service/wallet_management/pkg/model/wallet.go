package model

import (
  "github.com/jinzhu/gorm"
)

// UNIQUE constraint across multiple keys?
// https://github.com/jinzhu/gorm/issues/961
type Wallet struct {
  gorm.Model
  Bip44Change int     `gorm:"unique_index:idx_bip44_change_bip44_index_xpub_uid;not null"`
  Address     string  `gorm:"unique_index;not null"`
  Bip44Index  uint32  `gorm:"unique_index:idx_bip44_change_bip44_index_xpub_uid;not null"`
  XpubUID     uint    `gorm:"unique_index:idx_bip44_change_bip44_index_xpub_uid;not null"`
  Xpub        Xpub    `gorm:"foreignkey:XpubUID"`
  Status      bool
}
