package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)


type FilterDB interface {
	FilterI(f interface{}) (interface{}, error)
}

type ListReply struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

func ListFilter(newFilter func()interface{}, db FilterDB) func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		var f = newFilter()
		if !ginhelper.BindRequest(c, f) {
			return nil
		}
		result, err := db.FilterI(f)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeSelectError,
				Error: err.Error(),
			})
		}
		return result
	}
}
