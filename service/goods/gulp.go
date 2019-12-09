package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Goods{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.goodsDB.ID(id)
}

func (srv *Service) CreateFilter() interface{} {
	return new(model.GoodsFilter)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return srv.AfterPost(GoodsToPostReply(obj.(*model.Goods)))
}

func (srv *Service) DeleteHook(c controller.MContext, goods base_service.CRUDEntity) bool {
	return srv.deleteHook(c, goods.(*model.Goods))
}

func (srv *Service) ResponseGet(c controller.MContext, obj base_service.CRUDEntity) interface{} {
	return srv.GoodsToGetReply(c, obj.(*model.Goods))
}

func (srv *Service) ResponseInspect(obj base_service.CRUDEntity) interface{} {
	return GoodsToInspectReply(obj.(*model.Goods))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c controller.MContext, goods base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, goods.(*model.Goods), req.(*PutRequest))
}
