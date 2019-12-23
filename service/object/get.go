package objectservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type GetReply struct {
	Code   types.CodeType `json:"code"`
	Object *model.Object  `json:"object"`
}

func ObjectToGetReply(obj *model.Object) *GetReply {
	return &GetReply{
		Code:   types.CodeOK,
		Object: obj,
	}
}
