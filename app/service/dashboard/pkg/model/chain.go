package model

import (
  "github.com/jinzhu/gorm"
)

// Chain chain
type Chain struct {
  gorm.Model
  GroupID string  `gorm:"unique_index:idx_group_id_name;not null"`
  Group   Group
  Name    string  `gorm:"unique_index:idx_group_id_name;not null"`
  Coin    string
  Status  bool
  Tokens  []Token
}
