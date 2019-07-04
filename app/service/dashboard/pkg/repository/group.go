package repository

import (
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)
const groupResource = "group"

type groupRepo struct {}

type iGroupRepo interface {
  Create(tx *gorm.DB, m *model.Group) *gorm.DB
  Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) ([]*model.Group, error)
}

// Create save repo
func (repo groupRepo) Create(tx *gorm.DB, m *model.Group) *gorm.DB {
  return tx.Create(m)
}

func (repo *groupRepo) Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (groups []*model.Group, err error) {
  err = tx.Where(query).Where("id in (?)", ids).Find(&groups).Error
  if len(groups) < 1 {
    return nil, errors.New("Empty records")
  }
  return
}
