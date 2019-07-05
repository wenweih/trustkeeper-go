package repository

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

type mnemonicVersionRepo struct {}

type imnemonicVersionRepo interface {
  Create(tx *gorm.DB, m *model.MnemonicVersion) *gorm.DB
  VersionLikeQuery(tx *gorm.DB, version string) ([]*model.MnemonicVersion, error)
}

// Create save repo
func (repo mnemonicVersionRepo) Create(tx *gorm.DB, m *model.MnemonicVersion) *gorm.DB {
  return tx.Create(m)
}

func (repo *mnemonicVersionRepo) VersionLikeQuery(tx *gorm.DB, version string) (mnemonicVs []*model.MnemonicVersion, err error) {
  err = tx.Where("version LIKE ?", "%" + version + "%").Find(&mnemonicVs).Error
  if len(mnemonicVs) < 1 {
    return nil, fmt.Errorf("Empty records")
  }
  return
}
