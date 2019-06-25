package model

import (
  "github.com/jinzhu/gorm"
)

// Group group
type Group struct {
  gorm.Model
  NamespaceID uint
  Name        string    `grom:"type:varchar(100);not null;index"`
  Desc        string    `sql:"type:text;"`
  CreatorID   string    `gorm:"index;not null"`
}
