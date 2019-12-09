package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code        int               `json:"code"`
	ID          uint              `json:"id"`
	Seller      *GoodsUserReply   `json:"seller" form:"seller"`
	CreatedAt   time.Time         `json:"created_at"`
	EndAt       time.Time         `json:"end_at"`
	Type        uint16            `json:"g_type"`
	Name        string            `json:"name"`
	MinPrice    uint64            `json:"min_price"`
	CurPrice    uint64            `json:"cur_price"`
	IsFixed     bool              `json:"is_fixed"`
	Description string            `json:"description"`
	Status      types.GoodsStatus `json:"status"`
}

func (srv *Service) GoodsToGetReply(c controller.MContext, obj *model.Goods) *GetReply {
	return &GetReply{
		Code:        types.CodeOK,
		ID:          obj.ID,
		Seller:      srv.fetchUser(c, obj.Seller),
		CreatedAt:   obj.CreatedAt,
		EndAt:       obj.EndAt,
		Type:        obj.Type,
		Name:        obj.Name,
		MinPrice:    obj.MinPrice,
		CurPrice:    obj.CurPrice,
		IsFixed:     obj.IsFixed,
		Description: obj.Description,
		Status:      obj.Status,
	}
}
