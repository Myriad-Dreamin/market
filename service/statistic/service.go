//go:generate package-attach-to-path -generate_register_map
package statisticservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type Service struct {
	cfg        *config.ServerConfig
	logger     types.Logger
	statFeeDB  *model.StatFeeDB
	filterFunc func(c *gin.Context) interface{}
}

func (svc *Service) StatisticSignatureXXX() interface{} { return svc }

func NewService(logger types.Logger, provider *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.logger = logger
	a.cfg = cfg
	a.statFeeDB = provider.StatFeeDB()
	//a.filterFunc = base_service.ListFilter(a.statFeeDB)
	return
}
