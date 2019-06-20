package model

import (
  "github.com/jinzhu/gorm"
)

// Account account table
type Account struct {
  gorm.Model
  Email     string    `gorm:"type:varchar(100);unique_index;not null"`
  Password  string    `gorm:"not null"`
  UUID      string    `gorm:"unique_index;not null"`
  TokenID   string    `gorm:"index"`
  // https://gorm.io/docs/has_one.html#Association-ForeignKey
  Namespace Namespace `gorm:"foreignkey:creator_uid;association_foreignkey:uuid"`
}
