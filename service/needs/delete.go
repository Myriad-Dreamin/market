package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/gin-gonic/gin"
)

func (srv *Service) deleteHook(c *gin.Context, object *model.Needs) (canDelete bool) {
	return false
}
