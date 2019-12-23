package reply

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
)

type ShortUserInfo struct {
	ID       uint   `json:"id" form:"id"`
	NickName string `json:"nick_name" form:"nick_name"`
	CityCode string `json:"city_code" form:"city_code"`
}

func FetchUser(c controller.MContext, userDB *model.UserDB, u uint) *ShortUserInfo {
	if u == 0 {
		return nil
	}
	usr, err := userDB.ID(u)
	if ginhelper.MaybeSelectError(c, usr, err) {
		return nil
	}
	return &ShortUserInfo{
		ID:       usr.ID,
		NickName: usr.NickName,
		CityCode: usr.CityCode,
	}
}
