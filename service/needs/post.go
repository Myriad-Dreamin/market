package needsservice

import (
	"github.com/Myriad-Dreamin/market/auth"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"time"
)


type PostReply struct {
	Code  int          `json:"code"`
	Needs *model.Needs `json:"needs"`
}

func NeedsToPostReply(obj *model.Needs) *PostReply {
	return &PostReply{
		Code:  types.CodeOK,
		Needs: obj,
	}
}

type PostRequest struct {
	EndAt       time.Time `json:"end_at" form:"end_at" binding:"required"`
	Type        uint16    `json:"g_type" form:"g_type" binding:"required"`
	Name        string    `json:"name" form:"name" binding:"required"`
	MinPrice    uint64    `json:"min_price" form:"min_price" binding:"required"`
	MaxPrice    uint64    `json:"max_price" form:"max_price" binding:"required"`
	Description string    `json:"description" form:"description"`
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Needs)
	// fill here
	var claims = ginhelper.GetCustomFields(c)
	obj.Buyer = uint(claims.UID)
	obj.EndAt = req.EndAt
	obj.Type = req.Type
	obj.Name = req.Name
	obj.MinPrice = req.MinPrice
	obj.MaxPrice = req.MaxPrice
	obj.Description = req.Description
	obj.Status = types.GoodsStatusUnFinished

	return obj
}


func (srv *Service) AfterPost(reply *PostReply) interface{} {
	if b, err := auth.NeedsEntity.AddReadPolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.Needs.Buyer), reply.Needs.ID);
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

	if b, err := auth.NeedsEntity.AddWritePolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.Needs.Buyer), reply.Needs.ID);
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

