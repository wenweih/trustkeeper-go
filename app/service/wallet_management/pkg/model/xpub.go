package model

import (
  "github.com/jinzhu/gorm"
)

// Xpub for organization
type Xpub struct {
  gorm.Model
  UUID          string    `gorm:"unique_index;not null"`
  Key           string    `gorm:"type:varchar(500);index"`
  Status        bool      `gorm:"default:true"`
  ChainID       int8
  BIP44Account  int
  GroupID       int8
}
