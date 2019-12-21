//go:generate package-attach-to-path -generate_register_map
package constservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	cfg    *config.ServerConfig
	logger types.Logger
	cities map[string]types.CityObject
	citiesReply CitiesReply
}

func (svc *Service) ConstSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (svc *Service, err error) {
	svc = new(Service)
	svc.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	svc.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	svc.cities = m.Require(config.ModulePath.Global.Cities).(map[string]types.CityObject)
	svc.citiesReply = CitiesReply{svc.cities}
	return
}
