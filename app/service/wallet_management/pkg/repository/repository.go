package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

// Repo repo obj
type repo struct {
  iChainRepo
  iWalletRepo
  iXpubRepo
}

// New new repo
func New(db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Chain{},
    model.Wallet{},
    model.Xpub{})
  repo := repo{
    &chainRepo{db},
    &walletRepo{db},
    &xpubRepo{db}}
  var biz IBiz = &repo
  return biz
}
