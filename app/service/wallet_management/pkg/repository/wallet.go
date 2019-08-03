package repository

import (
  "context"
  "github.com/jinzhu/gorm"
  "github.com/jinzhu/copier"
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


func (repo *repo) QueryWalletsForGroupByChainName(ctx context.Context, groupid, chainName string) ([]*Wallet, error) {
  chain := model.Chain{}
  err := repo.db.Where("name = ?", chainName).First(&chain).Error
  if err != nil {
    return nil, err
  }
  xpub := model.Xpub{}
  repo.db.Preload("Wallets").Where("group_id = ? AND bip44_chain_id = ? AND state = ?", groupid, chain.Bip44id, Assigned).First(&xpub)

  wallets :=[]*Wallet{}
  if err := copier.Copy(&wallets, &xpub.Wallets); err != nil {
    return nil, err
  }
  return wallets, nil
}
