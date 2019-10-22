package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code   int           `json:"code"`
	User *model.User `json:"user"`
}

func UserToPostReply(obj *model.User) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		User: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req = new(PostRequest)
	if !ginhelper.BindRequest(c, req) {
		return nil
	}

	var obj = new(model.User)
	// fill here
	return obj
}

