package repository

import (
  "github.com/jinzhu/gorm"
  "github.com/qor/transition"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

const (
  Idle = "idle"
  Assigned = "assigned"
  Abandon = "abandon"
)

type xpubRepo struct {}

type iXpubRepo interface {
  Create(tx *gorm.DB, m *model.Xpub) *gorm.DB
}

// Create save repo
func (repo xpubRepo) Create(tx *gorm.DB, m *model.Xpub) *gorm.DB {
  xpubStateMachine := transition.New(m)
  xpubStateMachine.Initial("idle")
  xpubStateMachine.Event("init")
  xpubStateMachine.Trigger("init", m, tx, "init default xpub status")
  return tx.Create(m)
}
