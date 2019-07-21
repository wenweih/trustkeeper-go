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

type ChainWithWallets struct {
  ChainName string  `json:"ChainName"`
  TotalSize int32   `json:"TotalSize"`
  Wallets []*Wallet `json:"Wallets"`
}

type Wallet struct {
  ID         string  `json:"ID"`
  Address    string  `json:"Address"`
  Status     bool    `json:"Status"`
  ChainName  string  `json:"ChainName"`
}
