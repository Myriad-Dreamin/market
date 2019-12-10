package goodsservice

import (
	"github.com/Myriad-Dreamin/market/control/auth"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
	"time"
)

type PostReply struct {
	Code  int          `json:"code"`
	Goods *model.Goods `json:"goods"`
}


func GoodsToPostReply(obj *model.Goods) *PostReply {
	return &PostReply{
		Code:  types.CodeOK,
		Goods: obj,
	}
}

type PostRequest struct {
	EndAt       time.Time `json:"end_at" form:"end_at" binding:"required"`
	Type        uint16    `json:"g_type" form:"g_type" binding:"required"`
	Name        string    `json:"name" form:"name" binding:"required"`
	MinPrice    uint64    `json:"min_price" form:"min_price" binding:"required"`
	IsFixed     bool      `json:"is_fixed" form:"is_fixed" validate:"exists"`
	Description string    `json:"description" form:"description"`
}

func (srv *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	if req.EndAt.Sub(time.Now()) < srv.cfg.BaseParametersConfig.GoodsMinimumEndDuration {
		c.AbortWithStatusJSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "could not set end time before a duration shorter than minimum end duration",
		})
		return nil
	}

	var obj = new(model.Goods)

	var claims = ginhelper.GetCustomFields(c)
	obj.Seller = uint(claims.UID)
	obj.EndAt = req.EndAt
	obj.Type = req.Type
	obj.Name = req.Name
	obj.MinPrice = req.MinPrice
	obj.CurPrice = req.MinPrice
	obj.IsFixed = req.IsFixed
	obj.Description = req.Description
	obj.Status = types.GoodsStatusUnFinished
	return obj
}

func (srv *Service) AfterPost(reply *PostReply) interface{} {
	if b, err := auth.GoodsEntity.AddReadPolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.Goods.Seller), reply.Goods.ID);
		err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return types.ErrorSerializer{
			Code:  types.CodeAddReadPrivilegeError,
			Error: err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}

	if b, err := auth.GoodsEntity.AddWritePolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.Goods.Seller), reply.Goods.ID);
		err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return types.ErrorSerializer{
			Code:  types.CodeAddWritePrivilegeError,
			Error: err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}
	return reply
}
