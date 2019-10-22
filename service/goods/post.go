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

/*
我要卖功能模块：默认显示当前自己已发布的所有物品信息；发布新的物品信息；修改自己已发布未成交物品信息；删除自己已发布未成交物品信息；
我要买功能模块：默认显示当前已发布的所有物品信息；发布自己新的物品购买需求信息；修改自己已发布未成交的购买需求信息；删除已发布未成交的购买需求信息；
*/

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

