//go:generate package-attach-to-path -generate_register_map
package needsservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	forceDelete base_service.DServiceInterface
	needsDB     *model.NeedsDB
	userDB      *model.UserDB
	logger      types.Logger
	enforcer    *model.Enforcer
	cfg         *config.ServerConfig
	filterFunc  func(c controller.MContext) interface{}
	key         string
}

func (srv *Service) NeedsSignatureXXX() interface{} { return srv }


func NewService(logger types.Logger, provider *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.needsDB = provider.NeedsDB()
	a.userDB = provider.UserDB()
	a.enforcer = provider.Enforcer()
	a.logger = logger
	a.cfg = cfg
	a.key = "nid"
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.forceDelete = base_service.NewDService(forceDeleteService{a}, a.key)
	a.ListService = base_service.NewListService(a, a.needsDB.FilterI)
	return
}

/*
type Needs struct {
}
*/
