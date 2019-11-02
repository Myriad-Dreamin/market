package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/lib/traits"
	"github.com/Myriad-Dreamin/market/service"
)
type serviceResult struct {
	serviceName string
	traits.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"needsService", traits.Decay(service.NewNeedsService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
		{"goodsService", traits.Decay(service.NewGoodsService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
		{"userService", traits.Decay(service.NewUserService(srv.Logger, srv.DatabaseProvider, srv.jwtMW, srv.cfg))},
		{"objectService", traits.Decay(service.NewObjectService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.serviceName, serviceResult.First)
	}
	return true
}


