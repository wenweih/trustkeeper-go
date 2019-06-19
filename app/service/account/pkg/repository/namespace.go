package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
)

type namespaceRepo struct {
  db *gorm.DB
}

type iNamespaceRepo interface {
  Create(m *model.Namespace) error
}

func (repo *namespaceRepo) Create(m *model.Namespace) error {
  return repo.db.Create(&m).Error
}
