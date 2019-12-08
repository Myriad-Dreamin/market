package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/gin-gonic/gin"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Object{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return srv.AfterPost(ObjectToPostReply(obj.(*model.Object)))
}

func (srv *Service) DeleteHook(c *gin.Context, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.Object))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return ObjectToGetReply(obj.(*model.Object))
}

func (srv *Service) ResponseInspect(obj base_service.CRUDEntity) interface{} {
	return ObjectToGetReply(obj.(*model.Object))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c *gin.Context, object base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, object.(*model.Object), req.(*PutRequest))
}
