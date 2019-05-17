package enforcer

import (
	"github.com/jinzhu/gorm"
	stdcasbin "github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewCasbinEnforcer(db *gorm.DB) *stdcasbin.Enforcer {
	Adapter := gormadapter.NewAdapterByDB(db)
	enforcer := stdcasbin.NewEnforcer(stdcasbin.NewModel(CasbinConf), Adapter)
	return enforcer
}
