package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

const walletResource = "wallet"

type walletRepo struct {
  db *gorm.DB
}

type iWalletRepo interface {
  Create(m *model.Wallet) error
}

func (repo *walletRepo) Create(m *model.Wallet) error {
  return repo.db.Create(&m).Error
}
