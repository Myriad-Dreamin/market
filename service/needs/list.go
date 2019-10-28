package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code    int            `json:"code"`
	Needss []model.Needs `json:"needss"`
}

func NeedssToListReply(obj []model.Needs) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Needss = obj
	return
}

func (srv *Service) FilterOn(c *gin.Context, page, pageSize int) (interface{}, error) {
	// parse c


	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if err != nil {
		return nil, err
	}
	return NeedssToListReply(objs), nil
}

