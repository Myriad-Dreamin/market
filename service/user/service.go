//go:generate package-attach-to-path -generate_register_map
package userservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/model/sp-layer"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/casbin/casbin/v2"
	"net/http"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	userDB     *model.UserDB
	goodsDB    *model.GoodsDB
	needsDB    *splayer.NeedsDB
	enforcer   *casbin.SyncedEnforcer
	logger     types.Logger
	cfg        *config.ServerConfig
	middleware *jwt.Middleware
	cities     map[string]types.CityObject
}

type CitiesReply struct {
	Cities map[string]types.CityObject `json:"cities"`
}

var citiesReply CitiesReply

func (srv *Service) GetCities(c controller.MContext) {
	c.JSON(http.StatusOK, citiesReply)
}

func (srv *Service) UserSignatureXXX() interface{} { return srv }

func NewService(m module.Module) (a control.UserService, err error) {
	srv := new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	srv.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	srv.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	srv.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	srv.cities = m.Require(config.ModulePath.Global.Cities).(map[string]types.CityObject)
	citiesReply = CitiesReply{srv.cities}
	srv.userDB = provider.UserDB()
	srv.goodsDB = provider.GoodsDB()
	srv.needsDB = provider.NeedsDB()
	srv.enforcer = provider.Enforcer()
	srv.CRUDService = base_service.NewCRUDService(srv, srv.cfg.BaseParametersConfig.PathPlaceholder.User)
	srv.ListService = base_service.NewListService(srv, srv.userDB.FilterI)
	a = srv
	return
}

/*
type User struct {
}
*/
