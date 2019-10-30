package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"time"
)

type PostReply struct {
	Code  int          `json:"code"`
	Goods *model.Goods `json:"goods"`
}

/*
我要卖功能模块：默认显示当前自己已发布的所有物品信息；发布新的物品信息；修改自己已发布未成交物品信息；删除自己已发布未成交物品信息；
我要买功能模块：默认显示当前已发布的所有物品信息；发布自己新的物品购买需求信息；修改自己已发布未成交的购买需求信息；删除已发布未成交的购买需求信息；
*/

func GoodsToPostReply(obj *model.Goods) *PostReply {
	return &PostReply{
		Code:  types.CodeOK,
		Goods: obj,
	}
}

type PostRequest struct {
	EndAt time.Time `json:"end_at" form:"end_at" binding:"required"`
	Type        uint16  `json:"g_type" form:"g_type" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	MinPrice    uint64 `json:"min_price" form:"min_price" binding:"required"`
	IsFixed     bool   `json:"is_fixed" form:"is_fixed" binding:"required"`
	Description string `json:"description" form:"description"`
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Goods)

	var claims = ginhelper.GetCustomFields(c)
	obj.Seller = uint(claims.UID)
	obj.EndAt = req.EndAt
	obj.Type = req.Type
	obj.Name = req.Name
	obj.MinPrice = req.MinPrice
	obj.IsFixed = req.IsFixed
	obj.Description = req.Description
	obj.Status = types.GoodsStatusUnFinished

	// fill here
	return obj
}
