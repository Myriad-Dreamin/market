package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

// 删除已发布未成交的自己物品信息

func (srv *Service) deleteHook(c controller.MContext, object *model.Needs) (canDelete bool) {
	if object.Seller != 0 {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: "this needs has been consumed",
		})
		return false
	}
	return true
}
