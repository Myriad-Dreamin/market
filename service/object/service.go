//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	base_service "github.com/Myriad-Dreamin/market/service/base-service"
	"github.com/Myriad-Dreamin/market/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.ObjectDB
	cfg *config.ServerConfig
	logger types.Logger
}

func (srv *Service) ObjectSignatureXXX() interface{} { return srv }

func NewService(logger types.Logger, provider *model.Provider, cfg *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.db = provider.ObjectDB()
	a.logger = logger
	a.cfg = cfg
	a.CRUDService = base_service.NewCRUDService(a, "oid")
	a.ListService = base_service.NewListService(a, "oid")
	return
}
