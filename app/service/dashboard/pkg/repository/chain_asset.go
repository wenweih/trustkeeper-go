package repository

import (
	"errors"
	"trustkeeper-go/app/service/dashboard/pkg/model"

	"github.com/jinzhu/gorm"
)

const chainAssetResource = "chain_asset"

type chainAssetRepo struct{}

type iChainAssetRepo interface {
	Create(tx *gorm.DB, m *model.Chain) *gorm.DB
	Update(tx *gorm.DB, m *model.Chain) *gorm.DB
	Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (chains []*model.Chain, err error)
}

// SimpleAsset token resp
type SimpleAsset struct {
	AssetID  string `json:"AssetID"`
	Symbol   string `json:"Symbol"`
	Status   bool   `json:"Status"`
	Identify string `json:"Identify"`
	Decimal  uint64 `json:"Decimal"`
}

// ChainAsset tokens correspond with chain resp
type ChainAsset struct {
	ChainID      string         `json:"ChainID"`
	Name         string         `json:"Name"`
	Coin         string         `json:"Coin"`
	Status       bool           `json:"Status"`
	Decimal      uint64         `json:"Decimal"`
	SimpleAssets []*SimpleAsset `json:"SimpleAsset"`
}

func (repo *chainAssetRepo) Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (chains []*model.Chain, err error) {
	err = tx.Preload("Assets").Where(query).Where("id in (?)", ids).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	if len(chains) < 1 {
		return nil, errors.New("Empty records")
	}
	return
}

func (repo *chainAssetRepo) Create(tx *gorm.DB, m *model.Chain) *gorm.DB {
	return tx.Create(m)
}

func (repo *chainAssetRepo) Update(tx *gorm.DB, m *model.Chain) *gorm.DB {
	record := model.Chain{}
	tx.Where(model.Chain{Name: m.Name, GroupID: m.GroupID, Coin: m.Coin}).Assign(model.Chain{Status: m.Status}).First(&record)
	m.ID = record.ID
	return tx.Save(m)
}
