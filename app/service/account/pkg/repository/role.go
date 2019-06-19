package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
)

type roleRepo struct {
  db *gorm.DB
}

type iRoleRepo interface {
  Create(m *model.Role) error
}

func (repo *roleRepo) Create(m *model.Role) error {
  return repo.db.Create(&m).Error
}
