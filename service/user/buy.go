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
可以查询当前所有用户基本信息；查询一定条件的物品及其状态信息，点击某一物品标识可显示货主信息；查询购物需求信息，点击某一标识可显示求购用户基本信息；查询一定条件下当前已经成交物品的累计中介费收益信息。
*/

func (srv *Service) Buy(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, "goid")
	if !ok {
		return
	}
	res, ok := ginhelper.RawJson(c)
	if !ok {
		return
	}
	var (
		claims = ginhelper.GetCustomFields(c)
		code   types.CodeType
		errs   string
	)
	if fixed := res.Get("fixed"); fixed.Exists() && fixed.Bool() {
		code, errs = srv.goodsDB.BuyFixed(id, uint(claims.UID))
	} else if price := res.Get("price"); price.Exists() {
		code, errs = srv.goodsDB.Buy(id, uint(claims.UID), price.Uint())
	} else {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "price must be given",
		})
	}

	if code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  code,
			Error: errs,
		})
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
	//if goods.Status
}
