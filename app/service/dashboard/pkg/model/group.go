package model

import (
  "github.com/jinzhu/gorm"
)

// Group group
type Group struct {
  gorm.Model
  NamespaceID string
  Name        string    `gorm:"unique_index;not null"`
  Desc        string    `sql:"type:text;"`
  CreatorID   string    `gorm:"index;not null"`
}
