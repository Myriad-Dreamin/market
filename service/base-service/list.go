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
	filterDB FilterDB
	k    string
}

func NewListService(tool ListableObjectToolLite, filterDB FilterDB, k string) ListService {
	return ListService{
		tool: tool,
		filterDB: filterDB,
		k:    k,
	}
}

type FilterDB interface {
	FilterI(f interface{}) (interface{}, error)
}

type ListReply struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}


func (srv *ListService) List(c *gin.Context) {
	var f = srv.tool.CreateFilter()
	if !ginhelper.BindRequest(c, f) {
		return
	}
	result, err := srv.filterDB.FilterI(f)
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
