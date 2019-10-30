package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/gin-gonic/gin"
)


// 删除已发布未成交的自己物品信息

func (srv *Service) deleteHook(c *gin.Context, object *model.Goods) (canDelete bool) {
	return false
}
