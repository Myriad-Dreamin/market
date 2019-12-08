//go:generate package-attach-to-path -generate_register_map
package goodsservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	forceDelete base_service.DServiceInterface
	goodsDB    *model.GoodsDB
	userDB     *model.UserDB
	logger     types.Logger
	cfg        *config.ServerConfig
	filterFunc func(c *gin.Context) interface{}
	key        string
}

func (srv *Service) GoodsSignatureXXX() interface{} { return srv }


func NewService(logger types.Logger, provider *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.goodsDB = provider.GoodsDB()
	a.userDB = provider.UserDB()
	a.logger = logger
	a.cfg = cfg
	a.key = "goid"
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.forceDelete = base_service.NewDService(forceDeleteService{a}, a.key)
	a.ListService = base_service.NewListService(a, a.goodsDB.FilterI)
	return
}

/*
type Goods struct {
}
*/
