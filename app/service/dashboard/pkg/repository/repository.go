package repository

import (
	"trustkeeper-go/library/casbin"

	"github.com/jinzhu/gorm"

	// trustkeeper-go/library/database/orm
	"trustkeeper-go/app/service/dashboard/pkg/model"
	account_const "trustkeeper-go/library/const/account"
)

// Repo repo obj
type repo struct {
	db *gorm.DB
	// redisPool *redis.Pool
	iCasbinRepo
	iGroupRepo
	iChainAssetRepo
}

// New new repo
func New(db *gorm.DB) IBiz {
	db.AutoMigrate(
		model.Group{},
		model.Chain{},
		model.Asset{})
	repo := repo{
		db,
		&casbinRepo{casbin.NewCasbinRepo(db)},
		&groupRepo{},
		&chainAssetRepo{}}
	repo.iCasbinRepo.AddResoucreCreatePolicyForRole(account_const.MerchantAdmin,
		"", []string{groupResource, chainAssetResource}...)
	var biz IBiz = &repo
	return biz
}

func (repo *repo) close() error {
	if err := repo.db.Close(); err != nil {
		return err
	}
	// if err := repo.redisPool.Close(); err != nil {
	//   return err
	// }
	return nil
}
