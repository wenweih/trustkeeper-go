package repository

import (
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type chainRepo struct {}

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
