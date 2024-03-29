package model

import (
	"github.com/jinzhu/gorm"
)

const (
	BTCSymbol string = "BTC"

	UTXOStateLocked     = "locked"
	UTXOStateSelected   = "selected"
	UTXOStateUnSelected = "unselected"
)

// BtcUtxo utxo model
type BtcUtxo struct {
	gorm.Model
	Txid       string  `gorm:"not null"`
	Amount     float64 `gorm:"not null"`
	Height     int64   `gorm:"not null"`
	VoutIndex  uint32  `gorm:"not null"`
	ReOrg      bool    `gorm:"not null;default:false"`
	UsedBy     string
	BalanceID  uint `gorm:"not null"`
	Balance    Balance
	BtcBlock   BtcBlock
	BtcBlockID uint
	State      string `gorm:"default:'unselected';not null"`
}

// BtcBlock notify block info
type BtcBlock struct {
	gorm.Model
	Hash   string    `gorm:"not null;unique_index:idx_hash_height"`
	Height int64     `gorm:"not null;unique_index:idx_hash_height"`
	Utxos  []BtcUtxo `gorm:"foreignkey:BtcBlockID;association_foreignkey:Refer"`
	ReOrg  bool      `gorm:"default:false"`
}
