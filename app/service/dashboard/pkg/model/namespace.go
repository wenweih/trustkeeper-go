package model

import (
  "github.com/jinzhu/gorm"
)

// Namespace for organization
type Namespace struct {
  gorm.Model
  CreatorID   string   `gorm:"index;not null"`
  Groups      []Group
  Name        string     `gorm:"index;unique_index"`
  Key         string    `gorm:"type:varchar(255);index"`
}
