package model

import (
	"github.com/jinzhu/gorm"
)

// Asset token
type Asset struct {
	gorm.Model
	GroupID  string `gorm:"unique_index:idx_group_id_symbol;not null"`
	Symbol   string `gorm:"unique_index:idx_group_id_symbol;not null"`
	Group    Group
	ChainID  string
	Chain    Chain
	Status   bool
	Identify string
	Decimal  uint64
}
