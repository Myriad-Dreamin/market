package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
	"github.com/Myriad-Dreamin/market/control/auth"

)

type ConfirmSellRequest struct {
	ConfirmOrCancel bool `json:"cc" form:"cc" validate:"exists"`
}

func (svc *Service) ConfirmSell(c controller.MContext) {
	var req ConfirmSellRequest
	id, ok := ginhelper.ParseUintAndBind(c, "nid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := svc.needsDB.ConfirmSell(id, req.ConfirmOrCancel, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		needs, err := svc.needsDB.ID(id)
		if ginhelper.MaybeSelectErrorWithTip(c, needs, err, "fetch to add read property") {
			return
		}

		if b, err := auth.NeedsEntity.AddReadPolicy(svc.enforcer, auth.UserEntity.CreateObj(needs.Seller), id); err != nil {
			if !b {
				svc.logger.Debug("add failed")
			}
			c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeAddReadPrivilegeError,
				Error: err.Error(),
			})
			return
		} else {
			if !b {
				svc.logger.Debug("add failed")
			}
		}
	
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}
