package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type chainRepo struct {
  db *gorm.DB
}

type iChainRepo interface {
  Create(m *model.Chain) error
}

// Create save repo
func (repo chainRepo) Create(m *model.Chain) error {
  return repo.db.Create(m).Error
}
