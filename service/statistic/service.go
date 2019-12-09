//go:generate package-attach-to-path -generate_register_map
package statisticservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
)

type Service struct {
	cfg                        *config.ServerConfig
	logger                     types.Logger
	statFeeDB                  *model.StatFeeDB
	filterFunc                 func(c controller.MContext) interface{}
	statGoodsFeeService        base_service.ListService
	statGoodsFeeXYService      base_service.ListService
	statGoodsFeeCountXYService base_service.ListService
}

func (svc *Service) StatisticSignatureXXX() interface{} { return svc }

func NewService(logger types.Logger, provider *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.logger = logger
	a.cfg = cfg
	a.statFeeDB = provider.StatFeeDB()
	a.statGoodsFeeService = base_service.NewListService(statGoodsFeeService{}, a.statFeeDB.FilterFeeI)
	a.statGoodsFeeXYService = base_service.NewListService(statGoodsFeeXYService{}, a.statFeeDB.FilterFeeI)
	a.statGoodsFeeCountXYService = base_service.NewListService(statGoodsFeeCountXYService{}, a.statFeeDB.FilterFeeCountI)
	//a.filterFunc = base_service.ListFilter(a.statFeeDB)
	return
}
