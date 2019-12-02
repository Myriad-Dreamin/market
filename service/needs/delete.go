package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 删除已发布未成交的自己物品信息

func (srv *Service) deleteHook(c *gin.Context, object *model.Needs) (canDelete bool) {
	if object.Seller != 0 {
		c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: "this needs has been consumed",
		})
		return false
	}
	return true
}
