package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Object{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return ObjectToPostReply(obj.(*model.Object))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return ObjectToGetReply(obj.(*model.Object))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(object base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(object.(*model.Object), req.(*PutRequest))
}
