package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *Service) deleteHook(c *gin.Context, object *model.Object) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
		Code:  types.CodeDeleteError,
		Error: "generated delete api has not been implemented yet",
	})
	return false
}
