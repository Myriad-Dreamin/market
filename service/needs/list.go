package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
	"time"
)

type NeedsUserReply struct {
	ID           uint   `json:"id" form:"id"`
	NickName     string `json:"nick_name" form:"nick_name"`
	RegisterCity string `json:"register_city" form:"register_city"`
}

type NeedsReply struct {
	ID          uint              `json:"id" form:"id"`
	CreatedAt   time.Time         `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" form:"updated_at"`
	EndAt       time.Time         `json:"end_at" form:"end_at"`
	Seller      *NeedsUserReply   `json:"seller" form:"seller"`
	Buyer       *NeedsUserReply   `json:"buyer" form:"buyer"`
	EndDuration time.Duration     `json:"ddd" form:"ddd"`
	Type        uint16            `json:"g_type" form:"g_type"`
	Name        string            `json:"name" form:"name"`
	MinPrice    uint64            `json:"min_price" form:"min_price"`
	MaxPrice    uint64            `json:"min_price" form:"min_price"`
	IsFixed     bool              `json:"is_fixed" form:"is_fixed"`
	Description string            `json:"description" form:"description"`
	Status      types.GoodsStatus `json:"status" form:"status"`
}

type ListReply struct {
	Code   int          `json:"code"`
	Needss []NeedsReply `json:"needss"`
}

func (srv *Service) NeedssToListReply(c *gin.Context, obj []model.Needs) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Needss = srv.FromNeedss(c, obj)
	return
}

func (srv *Service) fetchUser(c *gin.Context, u uint) *NeedsUserReply {
	if u == 0 {
		return nil
	}
	usr, err := srv.userDB.ID(u)
	if ginhelper.MaybeSelectError(c, usr, err) {
		return nil
	}
	return &NeedsUserReply{
		ID:           usr.ID,
		NickName:     usr.NickName,
		RegisterCity: usr.RegisterCity,
	}
}

func (srv *Service) FromNeedss(c *gin.Context, needss []model.Needs) (gr []NeedsReply) {
	for i := range needss {
		gr = append(gr, NeedsReply{
			ID:          needss[i].ID,
			CreatedAt:   needss[i].CreatedAt,
			UpdatedAt:   needss[i].UpdatedAt,
			EndAt:       needss[i].EndAt,
			Seller:      srv.fetchUser(c, needss[i].Seller),
			Buyer:       srv.fetchUser(c, needss[i].Buyer),
			Type:        needss[i].Type,
			Name:        needss[i].Name,
			EndDuration: needss[i].EndDuration,
			MinPrice:    needss[i].MinPrice,
			MaxPrice:    needss[i].MaxPrice,
			Description: needss[i].Description,
			Status:      needss[i].Status,
		})
		if c.IsAborted() {
			return
		}
	}
	return
}

func (srv *Service) FilterOn(c *gin.Context) (interface{}, error) {
	result := srv.filterFunc(c)
	if c.IsAborted() {
		return nil, nil
	}
	return srv.NeedssToListReply(c, result.([]model.Needs)), nil
}
