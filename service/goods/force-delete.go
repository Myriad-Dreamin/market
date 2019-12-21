package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (svc *Service) ForceDelete(c controller.MContext) {
	svc.forceDelete.Delete(c)
	return
}

type forceDeleteService struct {
	base_service.FObject
}

func (srv forceDeleteService) DeleteHook(c controller.MContext, goods base_service.CRUDEntity) bool {
	return srv.deleteHook(c, goods.(*model.Goods))
}

func (srv forceDeleteService) deleteHook(c controller.MContext, object *model.Goods) (canDelete bool) {
	return true
}
