package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
)

type GetReply struct {
	Code int         `json:"code"`
	User *model.User `json:"user"`
}

func UserToGetReply(obj *model.User) *GetReply {
	return &GetReply{
		Code: types.CodeOK,
		User: obj,
	}
}
