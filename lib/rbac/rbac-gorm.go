package rbac

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
)


func InitGorm(x *gorm.DB) error {
	a, err := gormadapter.NewAdapterByDB(x)
	if err != nil {
		return err
	}
	e, err := casbin.NewSyncedEnforcer(rbacModel, a)
	if err != nil {
		return err
	}
	return registerEnforcer(e)
}
