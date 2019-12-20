//go:generate package-attach-to-path -generate_register_map
package needsservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
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

type GoodsTypesReply struct {
	Types map[string]types.GoodsType `json:"types"`
}

var goodsTypesReply = GoodsTypesReply{types.GoodsTypesMap}

func (srv *Service) GetTypes(c controller.MContext) {
	c.JSON(http.StatusOK, goodsTypesReply)
}

func (srv *Service) NeedsSignatureXXX() interface{} { return srv }

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.needsDB = provider.NeedsDB()
	a.userDB = provider.UserDB()
	a.enforcer = provider.Enforcer()
	a.key = "nid"
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.forceDelete = base_service.NewDService(forceDeleteService{a}, a.key)
	a.ListService = base_service.NewListService(a, a.needsDB.FilterI)
	return
}
