package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type mnemonicVersionRepo struct {}

type imnemonicVersionRepo interface {
  Create(tx *gorm.DB, m *model.MnemonicVersion) *gorm.DB
}

// Create save repo
func (repo mnemonicVersionRepo) Create(tx *gorm.DB, m *model.MnemonicVersion) *gorm.DB {
  return tx.Create(m)
}
