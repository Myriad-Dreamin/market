package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code   int           `json:"code"`
	Needs *model.Needs `json:"needs"`
}

func NeedsToPostReply(obj *model.Needs) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		Needs: obj,
	}
}

type PostRequest struct {
	Type uint8 `json:"g_type" form:"g_type"`
	Name string `json:"name" form:"name"`
	MinPrice uint64 `json:"min_price" form:"min_price"`
	MaxPrice uint64 `json:"max_price" form:"max_price"`
	IsFixed bool `json:"is_fixed" form:"is_fixed"`
	Description string `json:"description" form:"description"`
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
	obj.Type = req.Type
	obj.Name = req.Name
	obj.MinPrice = req.MinPrice
	obj.MaxPrice = req.MaxPrice
	obj.IsFixed = req.IsFixed
	obj.Description = req.Description

	return obj
}

