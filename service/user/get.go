package userservice

import (
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"time"
)

type GetReply struct {
	Code             types.CodeType       `json:"code"`
	NickName         string    `json:"nick_name" form:"nick_name"`
	CityCode string    `json:"city_code" form:"city_code"`
	LastLogin        time.Time `json:"last_login" form:"last_login"`
}

func UserToGetReply(obj *model.User) GetReply {
	return GetReply{
		Code:             types.CodeOK,
		NickName:         obj.NickName,
		CityCode: obj.CityCode,
		LastLogin:        obj.LastLogin,
	}
}
