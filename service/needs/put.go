package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
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
	MaxPrice    *uint64   `json:"max_price" form:"max_price"`
	Description string    `json:"description" form:"description"`
}

func (srv *Service) fillPutFields(c *gin.Context, needs *model.Needs, req *PutRequest) (fields []string) {
	if needs.EndAt.Sub(time.Now()) <= time.Minute {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsLifeTimeout,
			Error: "needs life time is consumed",
		})
		return
	}

	if needs.Status == types.GoodsStatusUnknown {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsStatusUnknown,
			Error: "unknown status of needs",
		})
		return
	}
	if needs.Status == types.GoodsStatusFinished {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsStatusFinished,
			Error: "needs is sold",
		})
		return
	}

	if req.EndAt.Sub(time.Now()) >= time.Minute {
		fields = append(fields, "end_at")
		needs.EndAt = req.EndAt
	} else if !req.EndAt.IsZero() {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "req.EndAt.Sub(time.Now()) must greater than or equal to time.Minute",
		})
		return
	}

	if req.MinPrice != nil {
		if req.MaxPrice != nil && *req.MaxPrice < *req.MinPrice {
			c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeInvalidParameters,
				Error: "max price should not be less than min price",
			})
			return
		}
		fields = append(fields, "min_price")
		needs.MinPrice = *req.MinPrice
	}
	if req.MaxPrice != nil {
		if *req.MaxPrice < needs.MinPrice {
			c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeInvalidParameters,
				Error: "max price should not be less than min price",
			})
			return
		}
		fields = append(fields, "max_price")
		needs.MaxPrice = *req.MaxPrice
	}

	if req.Type != types.GoodsTypeUnknown {
		fields = append(fields, "g_type")
		needs.Type = req.Type
	}

	if len(req.Description) != 0 {
		fields = append(fields, "description")
		needs.Description = req.Description
	}

	if len(req.Name) != 0 {
		fields = append(fields, "name")
		needs.Name = req.Name
	}
	return
}
