package casbin

import (
	"github.com/jinzhu/gorm"
	stdcasbin "github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CasbinRepo struct {
  *stdcasbin.SyncedEnforcer  // authorization service
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	Adapter := gormadapter.NewAdapterByDB(db)
	enforcer := stdcasbin.NewSyncedEnforcer(stdcasbin.NewModel(CasbinConf), Adapter)
	enforcer.EnableAutoSave(true)
	return &CasbinRepo{enforcer}
}
