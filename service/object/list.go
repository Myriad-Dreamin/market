package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code    int            `json:"code"`
	Objects []model.Object `json:"objects"`
}

func ObjectsToListReply(obj []model.Object) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Objects = obj
	return
}

func (srv *Service) ProcessListResults(_ *gin.Context, result interface{}) interface{} {
	return ObjectsToListReply(result.([]model.Object))
}
