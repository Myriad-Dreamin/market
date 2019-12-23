package needsservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type InspectReply struct {
	Code  types.CodeType `json:"code"`
	Needs *model.Needs   `json:"needs"`
}

func NeedsToInspectReply(obj *model.Needs) *InspectReply {
	return &InspectReply{
		Code:  types.CodeOK,
		Needs: obj,
	}
}
