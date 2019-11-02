package goodsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type GetReply struct {
	Code  int          `json:"code"`
	Goods *model.Goods `json:"goods"`
}

func GoodsToGetReply(obj *model.Goods) *GetReply {
	return &GetReply{
		Code:  types.CodeOK,
		Goods: obj,
	}
}
