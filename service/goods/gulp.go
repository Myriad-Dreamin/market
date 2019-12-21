package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (svc *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Goods{ID: id}
}

func (svc *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return svc.goodsDB.ID(id)
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.GoodsFilter)
}

func (svc *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return svc.AfterPost(GoodsToPostReply(obj.(*model.Goods)))
}

func (svc *Service) DeleteHook(c controller.MContext, goods base_service.CRUDEntity) bool {
	return svc.deleteHook(c, goods.(*model.Goods))
}

func (svc *Service) ResponseGet(c controller.MContext, obj base_service.CRUDEntity) interface{} {
	return svc.GoodsToGetReply(c, obj.(*model.Goods))
}

func (svc *Service) ResponseInspect(obj base_service.CRUDEntity) interface{} {
	return GoodsToInspectReply(obj.(*model.Goods))
}

func (svc *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, goods base_service.CRUDEntity, req interface{}) []string {
	return svc.fillPutFields(c, goods.(*model.Goods), req.(*PutRequest))
}
