package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

/**
TODO
可以是卖二手物品身份（卖家）：填写发布物品信息，上传物品照片（选项：先到先得or竞价（有效时间））；查询自己发布的物品信息，对于有竞价该物品的用户可点击同意或拒绝，如果同意则平台按照规则计算中介费（具体的计算规则？）并提交数据库存储；修改自己已发布未成交的自己物品信息；删除已发布未成交的自己物品信息；查询所有在售物品信息；查询所有物品需求信息。
卖家支付中介费（成交价格*0.01）。
*/

type SellRequest struct {
	Price uint64 `json:"price"`
}

func (srv *Service) Sell(c controller.MContext) {

	var req SellRequest
	id, ok := ginhelper.ParseUintAndBind(c, "nid", &req)
	if !ok {
		return
	}
	var claims = ginhelper.GetCustomFields(c)

	code, err := srv.needsDB.Sell(req.Price, id, uint(claims.UID))
	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: err,
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
}

