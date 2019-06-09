package repository

import (
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type xpubRepo struct {
  db *gorm.DB
}

type iXpubRepo interface {
  Create(m *model.Xpub) error
}

// Create save repo
func (repo xpubRepo) Create(m *model.Xpub) error {
  return repo.db.Create(m).Error
}
