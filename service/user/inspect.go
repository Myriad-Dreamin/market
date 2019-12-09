package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type InspectReply struct {
	Code int         `json:"code"`
	User *model.User `json:"user"`
}

func UserToInspectReply(obj *model.User) *InspectReply {
	return &InspectReply{
		Code: types.CodeOK,
		User: obj,
	}
}
