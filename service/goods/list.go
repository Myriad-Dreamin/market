package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code    int            `json:"code"`
	Goodss []model.Goods `json:"goodss"`
}

/*
物品信息查询功能模块：默认显示所有物品信息，可以输入物品类型、物品名称、价格区间条件，显示满足条件的物品信息，注意货主姓名、电话号码信息不可见。
求购物品信息查询功能模块：默认显示所有求购物品信息，可以输入物品类型、物品名称、价格区间条件，显示满足条件的求购物品信息，注意求购用户姓名、电话号码信息不可见。
*/

/*
查询一定条件的物品及其状态信息，点击某一物品标识可显示货主信息
查询购物需求信息，点击某一标识可显示求购用户基本信息
*/

func GoodssToListReply(obj []model.Goods) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Goodss = obj
	return
}

func (srv *Service) FilterOn(c *gin.Context, page, pageSize int) (interface{}, error) {
	// parse c


	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if err != nil {
		return nil, err
	}
	return GoodssToListReply(objs), nil
}

