package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/gin-gonic/gin"
)

func (srv *Service) ForceDelete(c *gin.Context) {
	srv.forceDelete.Delete(c)
	return
}

type forceDeleteService struct {
	base_service.FObject
}

func (srv forceDeleteService) DeleteHook(c *gin.Context, goods base_service.CRUDEntity) bool {
	return srv.deleteHook(c, goods.(*model.Goods))
}

func (srv forceDeleteService) deleteHook(c *gin.Context, object *model.Goods) (canDelete bool) {
	return true
}
