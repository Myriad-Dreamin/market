package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code        int               `json:"code"`
	ID          uint              `json:"id"`
	CreatedAt   time.Time         `json:"created_at"`
	EndAt       time.Time         `json:"end_at"`
	Type        uint16            `json:"g_type"`
	Name        string            `json:"name"`
	MinPrice    uint64            `json:"min_price"`
	IsFixed     bool              `json:"is_fixed"`
	Description string            `json:"description"`
	Status      types.GoodsStatus `json:"status"`
}

func GoodsToGetReply(obj *model.Goods) *GetReply {
	return &GetReply{
		Code:        types.CodeOK,
		ID:          obj.ID,
		CreatedAt:   obj.CreatedAt,
		EndAt:       obj.EndAt,
		Type:        obj.Type,
		Name:        obj.Name,
		MinPrice:    obj.MinPrice,
		IsFixed:     obj.IsFixed,
		Description: obj.Description,
		Status:      obj.Status,
	}
}
