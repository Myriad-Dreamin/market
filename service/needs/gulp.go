package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Needs{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.needsDB.ID(id)
}

func (srv *Service) CreateFilter() interface{} {
	return new(model.GoodsFilter)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return srv.AfterPost(NeedsToPostReply(obj.(*model.Needs)))
}

func (srv *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.Needs))
}

func (srv *Service) ResponseGet(c controller.MContext, obj base_service.CRUDEntity) interface{} {
	return srv.NeedsToGetReply(c, obj.(*model.Needs))
}

func (srv *Service) ResponseInspect(obj base_service.CRUDEntity) interface{} {
	return NeedsToInspectReply(obj.(*model.Needs))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c controller.MContext, needs base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, needs.(*model.Needs), req.(*PutRequest))
}
