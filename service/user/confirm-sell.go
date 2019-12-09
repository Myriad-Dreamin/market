package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

type ConfirmSellRequest struct {
	ConfirmOrCancel bool `json:"cc" form:"cc" validate:"exists"`
}

func (srv *Service) ConfirmSell(c controller.MContext) {
	var req ConfirmSellRequest
	id, ok := ginhelper.ParseUintAndBind(c, "nid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := srv.needsDB.ConfirmSell(id, req.ConfirmOrCancel, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}