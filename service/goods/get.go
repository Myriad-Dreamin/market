package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/service/reply"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code        types.CodeType                  `json:"code"`
	ID          uint                 `json:"id"`
	Seller      *reply.ShortUserInfo `json:"seller" form:"seller"`
	CreatedAt   time.Time            `json:"created_at"`
	EndAt       time.Time            `json:"end_at"`
	Type        types.GoodsType      `json:"g_type"`
	Name        string               `json:"name"`
	MinPrice    uint64               `json:"min_price"`
	CurPrice    uint64               `json:"cur_price"`
	IsFixed     bool                 `json:"is_fixed"`
	Description string               `json:"description"`
	PicName string               `json:"pic_name"`
	Status      types.GoodsStatus    `json:"status"`
}

func (svc *Service) GoodsToGetReply(c controller.MContext, obj *model.Goods) *GetReply {
	return &GetReply{
		Code:        types.CodeOK,
		ID:          obj.ID,
		Seller:      reply.FetchUser(c, svc.userDB, obj.Seller),
		CreatedAt:   obj.CreatedAt,
		EndAt:       obj.EndAt,
		Type:        obj.Type,
		Name:        obj.Name,
		MinPrice:    obj.MinPrice,
		CurPrice:    obj.CurPrice,
		IsFixed:     obj.IsFixed,
		Description: obj.Description,
		PicName: obj.PicName,
		Status:      obj.Status,
	}
}
