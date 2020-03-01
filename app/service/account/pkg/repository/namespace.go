package repository

import (
	"trustkeeper-go/app/service/account/pkg/model"

	"github.com/jinzhu/gorm"
)

type namespaceRepo struct{}

type iNamespaceRepo interface {
	Create(tx *gorm.DB, m *model.Namespace) *gorm.DB
}

func (repo *namespaceRepo) Create(tx *gorm.DB, m *model.Namespace) *gorm.DB {
	return tx.Create(&m)
}
