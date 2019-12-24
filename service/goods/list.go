package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/service/reply"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GoodsReply struct {
	ID          uint                 `json:"id" form:"id"`
	CreatedAt   time.Time            `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" form:"updated_at"`
	EndAt       time.Time            `json:"end_at" form:"end_at"`
	Seller      *reply.ShortUserInfo `json:"seller" form:"seller"`
	Type        types.GoodsType      `json:"g_type" form:"g_type"`
	Name        string               `json:"name" form:"name"`
	CurPrice    uint64               `json:"cur_price" form:"cur_price"`
	MinPrice    uint64               `json:"min_price" form:"min_price"`
	IsFixed     bool                 `json:"is_fixed" form:"is_fixed"`
	Description string               `json:"description" form:"description"`
	PicName     string               `json:"pic_name"`
	Status      types.GoodsStatus    `json:"status" form:"status"`
}

type ListReply struct {
	Code   types.CodeType `json:"code"`
	Goodss []GoodsReply   `json:"goodss"`
}

func (svc *Service) GoodssToListReply(c controller.MContext, obj []model.Goods) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Goodss = svc.FromGoodss(c, obj)
	return
}

func (svc *Service) FromGoodss(c controller.MContext, goodss []model.Goods) (gr []GoodsReply) {
	for i := range goodss {
		gr = append(gr, GoodsReply{
			ID:          goodss[i].ID,
			CreatedAt:   goodss[i].CreatedAt,
			UpdatedAt:   goodss[i].UpdatedAt,
			EndAt:       goodss[i].EndAt,
			Seller:      reply.FetchUser(c, svc.userDB, goodss[i].Seller),
			Type:        goodss[i].Type,
			Name:        goodss[i].Name,
			CurPrice:    goodss[i].CurPrice,
			MinPrice:    goodss[i].MinPrice,
			IsFixed:     goodss[i].IsFixed,
			Description: goodss[i].Description,
			PicName:     goodss[i].PicName,
			Status:      goodss[i].Status,
		})
		if c.IsAborted() {
			return
		}
	}
	return
}

func (svc *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return svc.GoodssToListReply(c, result.([]model.Goods))
}
