package model

import (
  "github.com/jinzhu/gorm"
)

// Token token
type Token struct {
  gorm.Model
  GroupID    string  `gorm:"unique_index:idx_group_id_symbol;not null"`
  Group      Group
  Symbol     string  `gorm:"unique_index:idx_group_id_symbol;not null"`
  ChainID    string
  Chain      Chain
  Status  bool
}
