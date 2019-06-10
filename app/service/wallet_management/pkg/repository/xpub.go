package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type xpubRepo struct {
  db *gorm.DB
}

type iXpubRepo interface {
  Create(m *model.Xpub) error
}

// Create save repo
func (repo xpubRepo) Create(m *model.Xpub) error {
  return repo.db.Create(m).Error
}
