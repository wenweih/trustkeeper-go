package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
  "trustkeeper-go/app/service/account/pkg/enforcer"
)

type repo struct {
  db *gorm.DB
  iAccountRepo
  iNamespaceRepo
}

// New new
func New(db *gorm.DB, jwtKey string) IBiz {
  db.AutoMigrate(
    model.Account{},
    model.Namespace{})
  repo := repo{
    db,
    &accountRepo{db: db, Enforcer: enforcer.NewCasbinEnforcer(db)},
    &namespaceRepo{}}
  var biz IBiz = &repo
  return biz
}
