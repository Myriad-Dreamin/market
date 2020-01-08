package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
	"github.com/Myriad-Dreamin/market/control/auth"

)

type ConfirmBuyRequest struct {
	Confirm bool `json:"cc" form:"cc" validate:"exists"`
}

func (svc *Service) ConfirmBuy(c controller.MContext) {
	var req ConfirmBuyRequest
	id, ok := ginhelper.ParseUintAndBind(c, "goid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := svc.goodsDB.ConfirmBuy(id, req.Confirm, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		goods, err := svc.goodsDB.ID(id)
		if ginhelper.MaybeSelectErrorWithTip(c, goods, err, "fetch to add read property") {
			return
		}

		if b, err := auth.GoodsEntity.AddReadPolicy(svc.enforcer, auth.UserEntity.CreateObj(goods.Buyer), id); err != nil {
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
