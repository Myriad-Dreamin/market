package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code             int       `json:"code"`
	NickName         string    `json:"nick_name" form:"nick_name"`
	RegisterProvince string    `json:"register_province" form:"register_province"`
	RegisterCity     string    `json:"register_city" form:"register_city"`
	LastLogin        time.Time `json:"last_login" form:"last_login"`
}

func UserToGetReply(obj *model.User) GetReply {
	return GetReply{
		Code:             types.CodeOK,
		NickName:         obj.NickName,
		RegisterProvince: obj.RegisterProvince,
		RegisterCity:     obj.RegisterCity,
		LastLogin:        obj.LastLogin,
	}
}
