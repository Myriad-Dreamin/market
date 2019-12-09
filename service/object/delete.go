package objectservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
)

func (srv *Service) deleteHook(c controller.MContext, object *model.Object) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
		Code:  types.CodeDeleteError,
		Error: "generated delete api has not been implemented yet",
	})
	return false
}
