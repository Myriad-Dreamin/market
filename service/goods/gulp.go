package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Goods{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return GoodsToPostReply(obj.(*model.Goods))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return GoodsToGetReply(obj.(*model.Goods))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(goods base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(goods.(*model.Goods), req.(*PutRequest))
}
