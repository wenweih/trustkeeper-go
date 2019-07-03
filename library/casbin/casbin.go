package casbin

import (
	"github.com/jinzhu/gorm"
	stdcasbin "github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CasbinRepo struct {
  *stdcasbin.Enforcer  // authorization service
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	Adapter := gormadapter.NewAdapterByDB(db)
	enforcer := stdcasbin.NewEnforcer(stdcasbin.NewModel(CasbinConf), Adapter)
	return &CasbinRepo{enforcer}
}
