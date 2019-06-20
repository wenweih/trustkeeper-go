package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
)

type namespaceRepo struct {}

type iNamespaceRepo interface {
  Create(tx *gorm.DB, m *model.Namespace) *gorm.DB
}

func (repo *namespaceRepo) Create(tx *gorm.DB, m *model.Namespace) *gorm.DB {
  return tx.Create(&m)
}
