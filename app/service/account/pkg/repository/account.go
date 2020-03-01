package repository

import (
	"context"
	"errors"
	"trustkeeper-go/app/service/account/pkg/model"

	"github.com/jinzhu/gorm"
)

type accountRepo struct{}

type iAccountRepo interface {
	Create(tx *gorm.DB, m *model.Account) *gorm.DB
	Query(ctx context.Context, tx *gorm.DB, query *model.Account) ([]*model.Account, error)
	Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error
}

func (repo *accountRepo) Query(ctx context.Context, tx *gorm.DB, query *model.Account) (accounts []*model.Account, err error) {
	result := tx.Set("gorm:auto_preload", true).Find(&accounts, query)
	if result.Error != nil {
		return nil, errors.New("Account Query error: " + err.Error())
	}
	if len(accounts) < 1 {
		return nil, errors.New("Empty records")
	}
	return accounts, nil
}

func (repo *accountRepo) Create(tx *gorm.DB, m *model.Account) *gorm.DB {
	return tx.Create(m)
}

func (repo *accountRepo) Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error {
	return tx.Model(acc).Update(colums).Error
}
