package repository

import (
	"errors"
	"trustkeeper-go/app/service/dashboard/pkg/model"

	"github.com/jinzhu/gorm"
)

const groupResource = "group"

type groupRepo struct{}

type iGroupRepo interface {
	Create(tx *gorm.DB, m *model.Group) *gorm.DB
	Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) ([]*model.Group, error)
	Update(tx *gorm.DB, m *model.Group) *gorm.DB
}

type GetGroupsResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Create save repo
func (repo groupRepo) Create(tx *gorm.DB, m *model.Group) *gorm.DB {
	return tx.Create(m)
}

func (repo *groupRepo) Query(tx *gorm.DB, ids []interface{}, query map[string]interface{}) (groups []*model.Group, err error) {
	err = tx.Order("created_at desc").Where(query).Where("id in (?)", ids).Find(&groups).Error
	if len(groups) < 1 {
		return nil, errors.New("Empty records")
	}
	return
}

func (repo *groupRepo) Update(tx *gorm.DB, m *model.Group) *gorm.DB {
	return tx.Save(m)
}
