package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/service/reply"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type NeedsUserReply struct {
	ID           uint   `json:"id" form:"id"`
	NickName     string `json:"nick_name" form:"nick_name"`
	RegisterCity string `json:"register_city" form:"register_city"`
}

type NeedsReply struct {
	ID          uint                 `json:"id" form:"id"`
	CreatedAt   time.Time            `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" form:"updated_at"`
	EndAt       time.Time            `json:"end_at" form:"end_at"`
	Buyer       *reply.ShortUserInfo `json:"buyer" form:"buyer"`
	EndDuration time.Duration        `json:"ddd" form:"ddd"`
	Type        types.GoodsType      `json:"g_type" form:"g_type"`
	Name        string               `json:"name" form:"name"`
	CurPrice    uint64               `json:"cur_price" form:"cur_price"`
	MaxPrice    uint64               `json:"max_price" form:"max_price"`
	IsFixed     bool                 `json:"is_fixed" form:"is_fixed"`
	Description string               `json:"description" form:"description"`
	Status      types.GoodsStatus    `json:"status" form:"status"`
}

type ListReply struct {
	Code   types.CodeType          `json:"code"`
	Needss []NeedsReply `json:"needss"`
}

func (srv *Service) NeedssToListReply(c controller.MContext, obj []model.Needs) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Needss = srv.FromNeedss(c, obj)
	return
}

func (srv *Service) FromNeedss(c controller.MContext, needss []model.Needs) (gr []NeedsReply) {
	for i := range needss {
		gr = append(gr, NeedsReply{
			ID:          needss[i].ID,
			CreatedAt:   needss[i].CreatedAt,
			UpdatedAt:   needss[i].UpdatedAt,
			EndAt:       needss[i].EndAt,
			Buyer:       reply.FetchUser(c, srv.userDB, needss[i].Buyer),
			Type:        needss[i].Type,
			Name:        needss[i].Name,
			EndDuration: needss[i].EndDuration,
			CurPrice:    needss[i].CurPrice,
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

func (srv *Service) ProcessListResults(c controller.MContext, results interface{}) interface{} {
	return srv.NeedssToListReply(c, results.([]model.Needs))
}
