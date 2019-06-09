package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type groupRepo struct {
  db *gorm.DB
}

type iGroupRepo interface {
  Create(m *model.Group) error
}

// Create save repo
func (repo groupRepo) Create(m *model.Group) error {
  return repo.db.Create(m).Error
}
