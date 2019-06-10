package model

import (
  "github.com/jinzhu/gorm"
)

type Chain struct {
  gorm.Model
  Symbol  string
  BIP44ID int
  Status  bool
}
