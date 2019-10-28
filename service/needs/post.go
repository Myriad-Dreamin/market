package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code   int           `json:"code"`
	Needs *model.Needs `json:"needs"`
}

func NeedsToPostReply(obj *model.Needs) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		Needs: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req = new(PostRequest)
	if !ginhelper.BindRequest(c, req) {
		return nil
	}

	var obj = new(model.Needs)
	// fill here
	return obj
}

