//go:generate package-attach-to-path -generate_register_map
package authservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control/auth"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	cfg        *config.ServerConfig
	logger     types.Logger
	middleware *jwt.Middleware
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

func (svc *Service) GrantGroup(c controller.MContext) {
	panic("implement me")
}

func (svc *Service) RevokeGroup(c controller.MContext) {
	panic("implement me")
}

func (svc *Service) CheckGroup(c controller.MContext) {
	panic("implement me")
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
	Code  int    `json:"code"`
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
	a.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	return
}
