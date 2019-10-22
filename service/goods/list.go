package goodsservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code    int            `json:"code"`
	Goodss []model.Goods `json:"goodss"`
}

func GoodssToListReply(obj []model.Goods) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Goodss = obj
	return
}

func (srv *Service) FilterOn(c *gin.Context, page, pageSize int) (interface{}, error) {
	// parse c


	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if err != nil {
		return nil, err
	}
	return GoodssToListReply(objs), nil
}

