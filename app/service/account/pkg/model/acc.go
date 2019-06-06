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
}
