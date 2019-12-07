package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CRUDObjectToolLite interface {
	FObject
	DObject
	SerializePost(c *gin.Context) CRUDEntity
	ResponsePost(obj CRUDEntity) interface{}
	ResponseGet(obj CRUDEntity) interface{}
	GetPutRequest() interface{}
	FillPutFields(c *gin.Context, object CRUDEntity, req interface{}) []string
}

type CRUDService struct {
	Tool CRUDObjectToolLite
	k    string
	DServiceInterface
}

func NewCRUDService(tool CRUDObjectToolLite, k string) CRUDService {
	return CRUDService{
		Tool: tool,
		k:    k,
		DServiceInterface: NewDService(tool, k),
	}
}



func (srv *CRUDService) Get(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, srv.Tool.ResponseGet(obj))
}

func (srv *CRUDService) Put(c *gin.Context) {
	var req = srv.Tool.GetPutRequest()
	id, ok := ginhelper.ParseUintAndBind(c, srv.k, req)
	if !ok {
		return
	}

	object, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, object, err) {
		return
	}

	fields := srv.Tool.FillPutFields(c, object, req)
	if c.IsAborted() {
		return
	}

	if ginhelper.UpdateFields(c, object, fields) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (srv *CRUDService) Post(c *gin.Context) {
	var obj = srv.Tool.SerializePost(c)
	if c.IsAborted() {
		return
	}
	if ginhelper.CreateObj(c, obj) {
		c.JSON(http.StatusOK, srv.Tool.ResponsePost(obj))
	}
}
