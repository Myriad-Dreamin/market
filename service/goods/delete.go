package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

func (srv *Service) deleteHook(c controller.MContext, object *model.Goods) (canDelete bool) {
	if object.Status == types.GoodsStatusFinished {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: "this goods has been sold",
		})
		return false
	}
	return true
}
