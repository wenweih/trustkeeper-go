package model

import (
	"github.com/jinzhu/gorm"
)

// Balance balance related with assets for address
type Balance struct {
	gorm.Model
	Address      string `gorm:"unique_index:idx_address_symbol;not null"`
	Symbol       string `gorm:"unique_index:idx_address_symbol;not null"`
	Identify     string
	Decimal      uint64
	Amount       string `gorm:"default:'0'"`
	WithdrawLock string `gorm:"default:'0'"`
	Txes         []Tx
	BalanceLog   []BalanceLog
}
