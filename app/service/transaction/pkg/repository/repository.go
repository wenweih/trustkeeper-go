package repository

import (
	"trustkeeper-go/app/service/transaction/pkg/model"

	"github.com/jinzhu/gorm"
)

// Repo repo obj
type repo struct {
	db *gorm.DB
}

// New new repo
func New(db *gorm.DB) IBiz {
	db.AutoMigrate(
		model.Balance{},
		model.Tx{},
		model.EthBlock{},
		model.BtcUtxo{},
		model.BtcBlock{},
		model.BalanceLog{})
	repo := repo{db}
	var biz IBiz = &repo
	return biz
}

func (repo *repo) close() error {
	if err := repo.db.Close(); err != nil {
		return err
	}
	return nil
}
