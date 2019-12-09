package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

/**
TODO
可以是买物品身份（买家）：填写想购买的二手物品信息；查询或修改自己发布的购买物品需求信息；删除已发布的购买物品需求信息；查询所有购买物品需求信息；查询所有在售物品信息。点击某一具体物品可看到明细信息，点击购买可模拟竞价购买的请求。不能看到全部的卖家信息
买家支付中介费（成交价格*0.02）

中介收益汇总表：物品类型

*/

type BuyRequest struct {
}

func (srv *Service) Buy(c controller.MContext) {
	var req BuyRequest
	id, ok := ginhelper.ParseUintAndBind(c, "goid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := srv.goodsDB.Buy(id, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
	//if goods.Status
}





