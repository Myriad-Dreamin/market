package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"net/http"
	"time"
)

type PutRequest struct {
	EndAt       time.Time       `json:"end_at" form:"end_at"`
	Type        types.GoodsType `json:"g_type" form:"g_type"`
	Name        string          `json:"name" form:"name"`
	MinPrice    *uint64         `json:"min_price" form:"min_price"`
	Description string          `json:"description" form:"description"`
}

func (srv *Service) fillPutFields(c controller.MContext, goods *model.Goods, req *PutRequest) (fields []string) {
	if goods.EndAt.Sub(time.Now()) <= time.Minute {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsLifeTimeout,
			Error: "goods life time is consumed",
		})
		return
	}

	if goods.Status == types.GoodsStatusUnknown {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsStatusUnknown,
			Error: "unknown status of goods",
		})
		return
	}
	if goods.Status == types.GoodsStatusFinished {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsStatusFinished,
			Error: "goods is sold",
		})
		return
	}
	if goods.Status == types.GoodsStatusCancelled {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeGoodsStatusCancelled,
			Error: "goods is sold",
		})
		return
	}

	if req.EndAt.Sub(time.Now()) >= time.Minute {
		fields = append(fields, "end_at")
		goods.EndAt = req.EndAt
	} else if !req.EndAt.IsZero() {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "req.EndAt.Sub(time.Now()) must greater than or equal to time.Minute",
		})
		return
	}

	//if req.IsFixed != nil {
	//	if goods.IsFixed && !*req.IsFixed {
	//		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
	//			Code:  types.CodeInvalidParameters,
	//			Error: "cant set fixed goods to be not fixed",
	//		})
	//		return
	//	}
	//	if goods.IsFixed != *req.IsFixed {
	//		fields = append(fields, "is_fixed")
	//		goods.IsFixed = *req.IsFixed
	//	}
	//}

	if req.MinPrice != nil {
		if !goods.IsFixed {
			c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
				Code:  types.CodeInvalidParameters,
				Error: "not fixed price should not be update",
			})
			return
		}
		fields = append(fields, "min_price")
		goods.MinPrice = *req.MinPrice
		if goods.MinPrice < goods.CurPrice {
			fields = append(fields, "cur_price", "buyer")
			goods.CurPrice = goods.MinPrice
			goods.Buyer = 0
		}
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
