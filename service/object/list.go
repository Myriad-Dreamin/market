package objectservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type ListReply struct {
	Code    types.CodeType            `json:"code"`
	Objects []model.Object `json:"objects"`
}

func ObjectsToListReply(obj []model.Object) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Objects = obj
	return
}

func (srv *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return ObjectsToListReply(result.([]model.Object))
}
