package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type InspectReply struct {
	Code  int          `json:"code"`
	Goods *model.Goods `json:"goods"`
}

func GoodsToInspectReply(obj *model.Goods) *InspectReply {
	return &InspectReply{
		Code:  types.CodeOK,
		Goods: obj,
	}
}
