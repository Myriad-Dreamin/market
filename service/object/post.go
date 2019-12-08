package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code   int           `json:"code"`
	Object *model.Object `json:"object"`
}

func ObjectToPostReply(obj *model.Object) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		Object: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Object)
	// fill here
	return obj
}

func (srv *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
