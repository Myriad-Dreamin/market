//go:generate package-attach-to-path -generate_register_map
package statisticservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
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

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.statFeeDB = provider.StatFeeDB()
	a.statGoodsFeeService = base_service.NewListService(statGoodsFeeService{}, a.statFeeDB.FilterFeeI)
	a.statGoodsFeeXYService = base_service.NewListService(statGoodsFeeXYService{}, a.statFeeDB.FilterFeeI)
	a.statGoodsFeeCountXYService = base_service.NewListService(statGoodsFeeCountXYService{}, a.statFeeDB.FilterFeeCountI)
	//a.filterFunc = base_service.ListFilter(a.statFeeDB)
	return
}
