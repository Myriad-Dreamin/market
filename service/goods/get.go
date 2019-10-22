package goodsservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
)

type GetReply struct {
	Code   int           `json:"code"`
	Goods *model.Goods `json:"goods"`
}

func GoodsToGetReply(obj *model.Goods) *GetReply {
	return &GetReply{
		Code:   types.CodeOK,
		Goods: obj,
	}
}

