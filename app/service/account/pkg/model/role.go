package model

import (
  "github.com/jinzhu/gorm"
)

type Role struct {
  gorm.Model
  Name        string  `gorm:"not null;unique_index"`
  Memo        string  `gorm:"size:200"`
}
