package model

import (
  "github.com/jinzhu/gorm"
)

// Namespace for organization
type Namespace struct {
  gorm.Model
  CreatorID   string   `gorm:"index;not null"`
  Groups      []Group
  Name        string    `gorm:"index;unique_index"`
  DefaultKey  string    `gorm:"type:varchar(500);index"`
  Xpubs       []Xpub    `gorm:"foreignkey:UUID"`
}
