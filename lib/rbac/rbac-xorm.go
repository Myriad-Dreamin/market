package rbac

import "github.com/casbin/casbin/v2"

import xormadapter "github.com/casbin/xorm-adapter"
import "github.com/go-xorm/xorm"

func InitXorm(x *xorm.Engine) error {

	a, err := xormadapter.NewAdapterByEngine(x)
	if err != nil {
		return err
	}
	e, err := casbin.NewSyncedEnforcer(rbacModel, a)
	if err != nil {
		return err
	}
	return registerEnforcer(e)
}
