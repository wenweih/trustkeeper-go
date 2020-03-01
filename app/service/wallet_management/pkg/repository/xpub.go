package repository

import (
	"trustkeeper-go/app/service/wallet_management/pkg/model"

	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

const (
	Idle     = "idle"
	Assigned = "assigned"
	Abandon  = "abandon"
)

type xpubRepo struct{}

type iXpubRepo interface {
	Create(tx *gorm.DB, m *model.Xpub) *gorm.DB
	UpdateState(tx *gorm.DB, m *model.Xpub, state string) error
}

// Create save repo
func (repo *xpubRepo) Create(tx *gorm.DB, m *model.Xpub) *gorm.DB {
	xpubStateMachine := transition.New(m)
	xpubStateMachine.Initial("idle")
	xpubStateMachine.Event("init")
	xpubStateMachine.Trigger("init", m, tx, "init default xpub status")
	return tx.Create(m)
}

func (repo *xpubRepo) UpdateState(tx *gorm.DB, m *model.Xpub, state string) error {
	xpubStateMachine := transition.New(m)
	xpubStateMachine.State(state)
	event := xpubStateMachine.Event(state)
	event.To(state).From(m.State)
	return xpubStateMachine.Trigger(state, m, tx, "assign xpub to specify group")
}
