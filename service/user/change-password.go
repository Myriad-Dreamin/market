package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

type ChangePasswordRequest struct {
	OldPassword string `form:"old-password" json:"old-password" binding:"required"`
	NewPassword string `form:"new-password" json:"new-password" binding:"required"`
}


func (srv *Service) ChangePassword(c controller.MContext) {
	var req = new(ChangePasswordRequest)
	id, ok := ginhelper.ParseUintAndBind(c, "id", req)
	if !ok {
		return
	}
	if sug := CheckStrongPassword(req.NewPassword); len(sug) != 0 {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeWeakPassword,
			Error: sug,
		})
		return
	}

	user, err := srv.userDB.Query(id)
	if ginhelper.MaybeSelectError(c, user, err) {
		return
	}

	if !ginhelper.AuthenticatePassword(c, user, req.OldPassword) {
		return
	}

	if ginhelper.ResetPassword(c, user, req.NewPassword) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}
