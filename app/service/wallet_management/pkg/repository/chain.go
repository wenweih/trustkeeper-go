package repository

import (
	"errors"
	"trustkeeper-go/app/service/wallet_management/pkg/model"

	"github.com/jinzhu/gorm"
)

// SimpleChain simple chain
type SimpleChain struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Coin    string `json:"desc"`
	Bip44id uint   `json:"bip44id"`
	Status  bool   `json:"status"`
	Decimal uint64
}

type chainRepo struct{}

type iChainRepo interface {
	Create(tx *gorm.DB, m *model.Chain) *gorm.DB
	Query(tx *gorm.DB, query map[string]interface{}) ([]*model.Chain, error)
}

// Create save repo
func (repo chainRepo) Create(tx *gorm.DB, m *model.Chain) *gorm.DB {
	return tx.Create(m)
}

func (repo *chainRepo) Query(tx *gorm.DB, query map[string]interface{}) (chains []*model.Chain, err error) {
	err = tx.Where(query).Find(&chains).Error
	if len(chains) < 1 {
		return nil, errors.New("Empty records")
	}
	return
}
