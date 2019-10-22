package needsservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Needs{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return NeedsToPostReply(obj.(*model.Needs))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return NeedsToGetReply(obj.(*model.Needs))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(needs base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(needs.(*model.Needs), req.(*PutRequest))
}
