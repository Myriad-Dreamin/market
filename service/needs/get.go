package needsservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
)

type GetReply struct {
	Code   int           `json:"code"`
	Needs *model.Needs `json:"needs"`
}

func NeedsToGetReply(obj *model.Needs) *GetReply {
	return &GetReply{
		Code:   types.CodeOK,
		Needs: obj,
	}
}

