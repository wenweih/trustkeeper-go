package model

import (
  "github.com/jinzhu/gorm"
)


type Role struct {
  gorm.Model
  GroupID     int
  Group       Group
  Name        string  `gorm:"not null;unique_index"`
  Memo        string  `gorm:"size:200"`
}
