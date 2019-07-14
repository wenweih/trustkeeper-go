package repository

import (
  "fmt"
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
  ChainID  string  `json:"chainid"`
	Name     string  `json:"name"`
  Coin     string  `json:"desc"`
  Status   bool    `json:"status"`
  SimpleTokens []*SimpleToken `json:"simpletokens"`
}

func (repo *chainAssetRepo) Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (chains []*model.Chain, err error) {
  result := tx.Set("gorm:auto_preload", true).Find(&chains, query)
  if result.Error != nil {
    return nil, fmt.Errorf("Chains Query error: %s", err.Error())
  }
  if len(chains) < 1 {
    return nil, fmt.Errorf("Empty records")
  }
  return chains, nil
}

func (repo *chainAssetRepo) Create(tx *gorm.DB, m *model.Chain) *gorm.DB {
  return tx.Save(m)
}

func (repo *chainAssetRepo) Update(tx *gorm.DB, m *model.Chain) *gorm.DB {
  return tx.Create(m)
}
