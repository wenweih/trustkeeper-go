package repository

import (
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

const chainAssetResource = "chain_asset"

type chainAssetRepo struct {}

type iChainAssetRepo interface {
  Create(tx *gorm.DB, m *model.Chain) *gorm.DB
  Update(tx *gorm.DB, m *model.Chain) *gorm.DB
  Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (chains []*model.Chain, err error)
}

// SimpleToken token resp
type SimpleToken struct {
  TokenID  string  `json:"tokenid"`
  Symbol   string  `json:"symbol"`
  Status   bool    `json:"status"`
}

// ChainAsset tokens correspond with chain resp
type ChainAsset struct {
  Chainid  string  `json:"chainid"`
	Name     string  `json:"name"`
  Coin     string  `json:"desc"`
  Status   bool    `json:"status"`
  SimpleTokens []*SimpleToken `json:"simpletokens"`
}

func (repo *chainAssetRepo) Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (chains []*model.Chain, err error) {
  err = tx.Preload("Tokens").Where(query).Where("id in (?)", ids).Find(&chains).Error
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
  return tx.Save(m)
}
