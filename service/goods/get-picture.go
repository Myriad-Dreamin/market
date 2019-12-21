package goodsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"strconv"
)

func (svc *Service) GetPicture(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, svc.key)
	if !ok {
		return
	}

	c.File(svc.cfg.BaseParametersConfig.GoodsPicturePath + strconv.Itoa(int(id)))
}
