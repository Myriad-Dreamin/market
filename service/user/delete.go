package userservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
)

func (srv *Service) deleteHook(c controller.MContext, object *model.User) (canDelete bool) {
	return true
}
