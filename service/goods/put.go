package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PutRequest struct {
	EndAt       time.Time `json:"end_at" form:"end_at"`
	Type        uint16    `json:"g_type" form:"g_type"`
	Name        string    `json:"name" form:"name"`
	MinPrice    *uint64   `json:"min_price" form:"min_price"`
	IsFixed     *bool     `json:"is_fixed" form:"is_fixed"`
	Description string    `json:"description" form:"description"`
}

func (srv *Service) fillPutFields(c *gin.Context, goods *model.Goods, req *PutRequest) (fields []string) {
	if goods.EndAt.Sub(time.Now()) <= time.Minute {
		c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
			Code:  types.CodeGoodsLifeTimeout,
			Error: "goods life time is consumed",
		})
		return
	}

	if goods.Status == types.GoodsStatusUnknown {
		c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
			Code:  types.CodeGoodsStatusUnknown,
			Error: "unknown status of goods",
		})
		return
	}
	if goods.Status == types.GoodsStatusFinished {
		c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
			Code:  types.CodeGoodsStatusFinished,
			Error: "goods is sold",
		})
		return
	}

	if req.EndAt.Sub(time.Now()) >= time.Minute {
		fields = append(fields, "end_at")
		goods.EndAt = req.EndAt
	} else if !req.EndAt.IsZero() {
		c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "req.EndAt.Sub(time.Now()) must greater than or equal to time.Minute",
		})
		return
	}

	if req.IsFixed != nil {
		if goods.IsFixed && !*req.IsFixed {
			c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
				Code:  types.CodeInvalidParameters,
				Error: "cant set fixed goods to be not fixed",
			})
			return
		}
		if goods.IsFixed != *req.IsFixed {
			fields = append(fields, "is_fixed")
			goods.IsFixed = *req.IsFixed
		}
	}

	if req.MinPrice != nil {
		if !goods.IsFixed {
			c.AbortWithStatusJSON(http.StatusOK, ginhelper.ErrorSerializer{
				Code:  types.CodeInvalidParameters,
				Error: "not fixed price should not be update",
			})
			return
		}
		fields = append(fields, "min_price")
		goods.MinPrice = *req.MinPrice
	}

	if req.Type != types.GoodsTypeUnknown {
		fields = append(fields, "g_type")
		goods.Type = req.Type
	}

	if len(req.Description) != 0 {
		fields = append(fields, "description")
		goods.Description = req.Description
	}

	if len(req.Name) != 0 {
		fields = append(fields, "name")
		goods.Name = req.Name
	}

	return
}
