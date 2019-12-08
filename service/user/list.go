package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
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
可以查询当前所有用户基本信息；查询一定条件的物品及其状态信息，点击某一物品标识可显示货主信息；查询购物需求信息，点击某一标识可显示求购用户基本信息；查询一定条件下当前已经成交物品的累计中介费收益信息。
*/

func (srv *Service) ProcessListResults(c *gin.Context, result interface{}) interface{} {
	return UsersToListReply(result.([]model.User))
}
