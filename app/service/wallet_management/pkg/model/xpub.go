package model

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

// Xpub for organization
// Bip44ChainID value would be 0 (Bitcoincore)
// NOTE When query with struct, GORM will only query with those fields has non-zero value,
// that means if your field’s value is 0, '', false or other zero values, it won’t be used to build query conditions
// http://gorm.io/docs/query.html
type Xpub struct {
	gorm.Model
	Key               string `gorm:"type:varchar(500);index"`
	Bip44ChainID      *uint
	Chain             Chain `gorm:"foreignkey:Bip44ChainID;association_foreignkey:Bip44id"`
	Bip44Account      int
	GroupID           string `gorm:"index"`
	MnemonicVersionID uint
	MnemonicVersion   MnemonicVersion
	transition.Transition
	Wallets []Wallet `gorm:"foreignkey:XpubUID"`
}
