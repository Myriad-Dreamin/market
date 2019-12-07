package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DObjectToolLite interface {
	FObject
	DObject
}

type DServiceInterface interface {
	Delete(c *gin.Context)
}

type DService struct {
	Tool DObjectToolLite
	k    string
}

func NewDService(tool DObjectToolLite, k string) DService {
	return DService{
		Tool: tool,
		k:    k,
	}
}

func (srv DService) Delete(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}
	if !srv.Tool.DeleteHook(c, obj) {
		return
	}

	if ginhelper.DeleteObj(c, obj) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	} else {
	}
}
