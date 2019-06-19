package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
  "trustkeeper-go/app/service/account/pkg/enforcer"
)

type repo struct {
  iAccountRepo
  iNamespaceRepo
  iRoleRepo
}

// New new
func New(db *gorm.DB, jwtKey string) IBiz {
  db.AutoMigrate(
    model.Account{},
    model.Role{})
  repo := repo{
    &accountRepo{db: db, Enforcer: enforcer.NewCasbinEnforcer(db)},
    &namespaceRepo{db: db},
    &roleRepo{db: db}}
  var biz IBiz = &repo
  return biz
}
