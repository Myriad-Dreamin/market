package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code  int          `json:"code"`
	Users []model.User `json:"users"`
}

func UsersToListReply(obj []model.User) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Users = obj
	return
}

/*
可以查询当前所有用户基本信息
*/

func (srv *Service) FilterOn(c *gin.Context, page, pageSize int) (interface{}, error) {
	// parse c

	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if err != nil {
		return nil, err
	}
	return UsersToListReply(objs), nil
}
