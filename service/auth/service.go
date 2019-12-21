//go:generate package-attach-to-path -generate_register_map
package authservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control/auth"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/model/sp-layer"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	cfg        *config.ServerConfig
	logger     types.Logger
	middleware *jwt.Middleware
	enforcer   *splayer.Enforcer
}

type ApiGroup struct {
	*Service
	groupName string
}

// @Title Group Grants
// @Description Group Grants
func (svc ApiGroup) Grant(c controller.MContext) {
	id , ok := ginhelper.ParseUint(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User)
	if !ok {
		return
	}
	if added, err := svc.enforcer.AddGroupingPolicy(auth.UserEntity.CreateObj(id), svc.groupName); !added {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantNoEffect,
			Error: "",
		})
	} else if err != nil {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}

}

// @Title Group Revokes
// @Description Group Revokes
func (svc ApiGroup) Revoke(c controller.MContext) {
	id , ok := ginhelper.ParseUint(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User)
	if !ok {
		return
	}
	if removed, err := svc.enforcer.RemoveGroupingPolicy(auth.UserEntity.CreateObj(id), svc.groupName); !removed {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantNoEffect,
			Error: "",
		})
	} else if err != nil {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}

// @Title Group Checks
// @Description Group Checks
func (svc ApiGroup) Check(c controller.MContext) {
	id , ok := ginhelper.ParseUint(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, CheckReply{Has:svc.enforcer.HasGroupingPolicy(auth.UserEntity.CreateObj(id), svc.groupName)})
}

func (svc *Service) GroupGranter(groupName string) controller.HandlerFunc {
	return ApiGroup{svc, groupName}.Grant
}

func (svc *Service) GroupRevoker(groupName string) controller.HandlerFunc {
	return ApiGroup{svc, groupName}.Revoke
}

func (svc *Service) GroupChecker(groupName string) controller.HandlerFunc {
	return ApiGroup{svc, groupName}.Check
}

func (svc *Service) Grant(c controller.MContext) {
	panic("implement me")
}

func (svc *Service) Revoke(c controller.MContext) {
	panic("implement me")
}

func (svc *Service) Check(c controller.MContext) {
	panic("implement me")
}

// /auth/group/user/1

type GroupRequest struct {
	GroupName string `json:"group" form:"group"`
}

func (svc *Service) GrantGroup(c controller.MContext) {
	var req GroupRequest
	id , ok := ginhelper.ParseUintAndBind(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User, &req)
	if !ok {
		return
	}
	if added, err := svc.enforcer.AddGroupingPolicy(auth.UserEntity.CreateObj(id), req.GroupName); !added {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantNoEffect,
			Error: "",
		})
	} else if err != nil {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}

func (svc *Service) RevokeGroup(c controller.MContext) {
	var req GroupRequest
	id , ok := ginhelper.ParseUintAndBind(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User, &req)
	if !ok {
		return
	}
	if removed, err := svc.enforcer.RemoveGroupingPolicy(auth.UserEntity.CreateObj(id), req.GroupName); !removed {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantNoEffect,
			Error: "",
		})
	} else if err != nil {
		c.JSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGrantError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}

type CheckReply struct {
	Has bool `json:"has"`
}

func (svc *Service) CheckGroup(c controller.MContext) {
	var req GroupRequest
	id , ok := ginhelper.ParseUintAndBind(c, svc.cfg.BaseParametersConfig.PathPlaceholder.User, &req)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, CheckReply{Has:svc.enforcer.HasGroupingPolicy(auth.UserEntity.CreateObj(id), req.GroupName)})
}

func (svc *Service) ReadGranter(entityInterface auth.ReadEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) ReadRevoker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) ReadChecker(entityInterface auth.ReadEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) WriteGranter(entityInterface auth.WriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) WriteRevoker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) WriteChecker(entityInterface auth.WriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) JustSimpleWriteGranter(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) JustSimpleWriteRevoker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) JustSimpleWriteChecker(entityInterface auth.JustSimpleWriteEntityInterface) controller.HandlerFunc {
	panic("implement me")
}

func (svc *Service) AuthSignatureXXX() interface{} { return svc }

type RefreshTokenReply struct {
	Code  types.CodeType    `json:"code"`
	Token string `json:"token"`
}

func (svc *Service) RefreshToken(c controller.MContext) {
	newToken, err := svc.middleware.RefreshToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, RefreshTokenReply{
		Code:  types.CodeOK,
		Token: newToken,
	})
}

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.enforcer = m.Require(config.ModulePath.Provider.Model).(*model.Provider).Enforcer()
	a.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	return
}
