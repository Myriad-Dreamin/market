package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GoodsUserReply struct {
	ID           uint   `json:"id" form:"id"`
	NickName     string `json:"nick_name" form:"nick_name"`
	RegisterCity string `json:"register_city" form:"register_city"`
}

type GoodsReply struct {
	ID          uint              `json:"id" form:"id"`
	CreatedAt   time.Time         `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" form:"updated_at"`
	EndAt       time.Time         `json:"end_at" form:"end_at"`
	Seller      *GoodsUserReply   `json:"seller" form:"seller"`
	Buyer       *GoodsUserReply   `json:"buyer" form:"buyer"`
	Type        uint16            `json:"g_type" form:"g_type"`
	Name        string            `json:"name" form:"name"`
	CurPrice    uint64            `json:"cur_price" form:"cur_price"`
	MinPrice    uint64            `json:"min_price" form:"min_price"`
	IsFixed     bool              `json:"is_fixed" form:"is_fixed"`
	Description string            `json:"description" form:"description"`
	Status      types.GoodsStatus `json:"status" form:"status"`
}

type ListReply struct {
	Code   int          `json:"code"`
	Goodss []GoodsReply `json:"goodss"`
}

func (srv *Service) GoodssToListReply(c controller.MContext, obj []model.Goods) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Goodss = srv.FromGoodss(c, obj)
	return
}

func (srv *Service) fetchUser(c controller.MContext, u uint) *GoodsUserReply {
	if u == 0 {
		return nil
	}
	usr, err := srv.userDB.ID(u)
	if ginhelper.MaybeSelectError(c, usr, err) {
		return nil
	}
	return &GoodsUserReply{
		ID:           usr.ID,
		NickName:     usr.NickName,
		RegisterCity: usr.RegisterCity,
	}
}

func (srv *Service) FromGoodss(c controller.MContext, goodss []model.Goods) (gr []GoodsReply) {
	for i := range goodss {
		gr = append(gr, GoodsReply{
			ID:          goodss[i].ID,
			CreatedAt:   goodss[i].CreatedAt,
			UpdatedAt:   goodss[i].UpdatedAt,
			EndAt:       goodss[i].EndAt,
			Seller:      srv.fetchUser(c, goodss[i].Seller),
			Buyer:       srv.fetchUser(c, goodss[i].Buyer),
			Type:        goodss[i].Type,
			Name:        goodss[i].Name,
			MinPrice:    goodss[i].MinPrice,
			IsFixed:     goodss[i].IsFixed,
			Description: goodss[i].Description,
			Status:      goodss[i].Status,
		})
		if c.IsAborted() {
			return
		}
	}
	return
}

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return srv.GoodssToListReply(c, result.([]model.Goods))
}
