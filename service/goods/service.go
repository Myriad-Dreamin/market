//go:generate package-attach-to-path -generate_register_map
package goodsservice

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
	goodsDB     *model.GoodsDB
	userDB      *model.UserDB
	enforcer    *model.Enforcer
	logger      types.Logger
	cfg         *config.ServerConfig
	filterFunc  func(c controller.MContext) interface{}
	key         string
}

type GoodsTypesReply struct {
	Types map[string]types.GoodsType `json:"types"`
}

var goodsTypesReply = GoodsTypesReply{types.GoodsTypesMap}

func (svc *Service) GetTypes(c controller.MContext) {
	c.JSON(http.StatusOK, goodsTypesReply)
}

func (svc *Service) GoodsSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.goodsDB = provider.GoodsDB()
	a.userDB = provider.UserDB()
	a.enforcer = provider.Enforcer()
	a.key = "goid"
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.forceDelete = base_service.NewDService(forceDeleteService{a}, a.key)
	a.ListService = base_service.NewListService(a, a.goodsDB.FilterI)
	return
}
