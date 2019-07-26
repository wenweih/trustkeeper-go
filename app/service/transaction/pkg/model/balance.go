package model

import (
  "github.com/jinzhu/gorm"
)

// Balance balance related with assets for address
type Balance struct {
  gorm.Model
  Address   string  `gorm:"index;not null"`
  Symbol    string  `gorm:"not null"`
  Identify  string
}
