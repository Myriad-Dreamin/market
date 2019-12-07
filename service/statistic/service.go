//go:generate package-attach-to-path -generate_register_map
package statisticservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-gonic/gin"
)

type Service struct {
	cfg *config.ServerConfig
	logger types.Logger
}

func (svc *Service) StatisticSignatureXXX() interface{} { return svc }

func NewService(logger types.Logger, _ *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.logger = logger
	a.cfg = cfg
	return
}
