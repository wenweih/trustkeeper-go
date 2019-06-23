package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type xpubRepo struct {}

type iXpubRepo interface {
  Create(tx *gorm.DB, m *model.Xpub) *gorm.DB
}

// Create save repo
func (repo xpubRepo) Create(tx *gorm.DB, m *model.Xpub) *gorm.DB {
  return tx.Create(m)
}
