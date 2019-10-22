package goodsservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code   int           `json:"code"`
	Goods *model.Goods `json:"goods"`
}

func GoodsToPostReply(obj *model.Goods) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		Goods: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req = new(PostRequest)
	if !ginhelper.BindRequest(c, req) {
		return nil
	}

	var obj = new(model.Goods)
	// fill here
	return obj
}

