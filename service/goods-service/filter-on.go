package goods_service

import (
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 3.9 3.10 3.15 3.16
// 4.9

type FilterDB interface {
	FilterI(f *model.GoodsFilter) (interface{}, error)
}

type ListReply struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

func ListFilter(db FilterDB) func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		var f model.GoodsFilter
		if !ginhelper.BindRequest(c, &f) {
			return nil
		}
		result, err := db.FilterI(&f)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
				Code:  types.CodeSelectError,
				Error: err.Error(),
			})
		}
		return result
	}
}
