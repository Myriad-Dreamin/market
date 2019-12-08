package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListableObjectToolLite interface {
	CreateFilter() interface{}
	ProcessListResults(c *gin.Context, r interface{}) interface{}
}

type ListService struct {
	tool ListableObjectToolLite
	filterFunc FilterFunc
}

func NewListService(tool ListableObjectToolLite, filterFunc FilterFunc) ListService {
	return ListService{
		tool: tool,
		filterFunc: filterFunc,
	}
}

type FilterFunc = func(f interface{}) (interface{}, error)

type ListReply struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}


func (srv *ListService) List(c *gin.Context) {
	var f = srv.tool.CreateFilter()
	if !ginhelper.BindRequest(c, f) {
		return
	}
	result, err := srv.filterFunc(f)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
	}

	result = srv.tool.ProcessListResults(c, result)
	if !c.IsAborted() {
		c.JSON(http.StatusOK, result)
	}
	return
}
