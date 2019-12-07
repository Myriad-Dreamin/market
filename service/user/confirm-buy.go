package userservice

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfirmBuyRequest struct {
	ConfirmOrCancel bool `json:"cc" form:"cc" binding:"exists"`
}


func (srv *Service) ConfirmBuy(c *gin.Context) {
	var req ConfirmBuyRequest
	id, ok := ginhelper.ParseUintAndBind(c, "goid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := srv.goodsDB.ConfirmBuy(id, req.ConfirmOrCancel, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}

}