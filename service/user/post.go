package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code int         `json:"code"`
	User *model.User `json:"user"`
}

func UserToPostReply(obj *model.User) *PostReply {
	return &PostReply{
		Code: types.CodeOK,
		User: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}


	var obj = new(model.User)
	// fill here
	return obj
}
