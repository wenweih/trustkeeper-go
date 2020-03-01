package model

import (
	"github.com/jinzhu/gorm"
)

// BalanceLog balance loog
type BalanceLog struct {
	gorm.Model
	TxID      string
	From      string
	To        string
	BalanceID uint
	Balance   Balance
}
