package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type GetReply struct {
	Code  int          `json:"code"`
	Needs *model.Needs `json:"needs"`
}

func NeedsToGetReply(obj *model.Needs) *GetReply {
	return &GetReply{
		Code:  types.CodeOK,
		Needs: obj,
	}
}
