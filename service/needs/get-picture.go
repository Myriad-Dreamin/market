package needsservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"strconv"
)

func (srv *Service) GetPicture(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, srv.key)
	if !ok {
		return
	}

	c.File(srv.cfg.BaseParametersConfig.NeedsPicturePath + strconv.Itoa(int(id)))
}

