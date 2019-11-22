package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/gin-gonic/gin"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Needs{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.needsDB.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return NeedsToPostReply(obj.(*model.Needs))
}

func (srv *Service) DeleteHook(c *gin.Context, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.Needs))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return NeedsToGetReply(obj.(*model.Needs))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c *gin.Context, needs base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, needs.(*model.Needs), req.(*PutRequest))
}
