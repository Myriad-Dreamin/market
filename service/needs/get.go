package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/service/reply"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code        int                  `json:"code"`
	ID          uint                 `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	EndAt       time.Time            `json:"end_at"`
	Buyer       *reply.ShortUserInfo `json:"buyer" form:"buyer"`
	Type        types.GoodsType      `json:"g_type"`
	Name        string               `json:"name"`
	MaxPrice    uint64               `json:"max_price"`
	CurPrice    uint64               `json:"cur_price"`
	Description string               `json:"description"`
	Status      types.GoodsStatus    `json:"status"`
}

func (srv *Service) NeedsToGetReply(c controller.MContext, obj *model.Needs) *GetReply {
	return &GetReply{
		Code:        types.CodeOK,
		ID:          obj.ID,
		Buyer:       reply.FetchUser(c, srv.userDB, obj.Buyer),
		CreatedAt:   obj.CreatedAt,
		EndAt:       obj.EndAt,
		Type:        obj.Type,
		Name:        obj.Name,
		MaxPrice:    obj.MaxPrice,
		CurPrice:    obj.CurPrice,
		Description: obj.Description,
		Status:      obj.Status,
	}
}
