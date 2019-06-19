package model

import (
  "github.com/jinzhu/gorm"
)

// Namespace for organization
type Namespace struct {
  gorm.Model
  CreatorID   string   `gorm:"index;not null"`
  Name        string    `gorm:"index;unique_index"`
  DefaultKey  string    `gorm:"type:varchar(500);index"`
}
