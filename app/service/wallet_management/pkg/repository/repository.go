package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/library/database/orm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

// Repo repo obj
type repo struct {
  iChainRepo
  iWalletRepo
  iXpubRepo
}

// DB database connect
func DB(dbInfo string) *gorm.DB {
  db, err := orm.Connect(dbInfo)
  if err != nil {
    panic(err.Error())
  }
  return db
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
