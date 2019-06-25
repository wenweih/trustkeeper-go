package repository

import (
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type groupRepo struct {}

type iGroupRepo interface {
  Create(tx *gorm.DB, m *model.Group) *gorm.DB
  Query(tx *gorm.DB, query map[string]interface{}) ([]*model.Group, error)
}

// Create save repo
func (repo groupRepo) Create(tx *gorm.DB, m *model.Group) *gorm.DB {
  return tx.Create(m)
}

func (repo *groupRepo) Query(tx *gorm.DB, query map[string]interface{}) (chains []*model.Group, err error) {
  err = tx.Where(query).Find(&chains).Error
  if len(chains) < 1 {
    return nil, errors.New("Empty records")
  }
  return
}
