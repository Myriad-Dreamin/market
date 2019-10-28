package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/service"
	"github.com/Myriad-Dreamin/market/types"
)
type serviceResult struct {
	serviceName string
	types.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"needsDB", types.Decay(service.NewNeedsService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
		{"goodsDB", types.Decay(service.NewGoodsService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
		{"userDB", types.Decay(service.NewUserService(srv.Logger, srv.DatabaseProvider, srv.jwtMW, srv.cfg))},
		{"objectService", types.Decay(service.NewObjectService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.First)
	}
	return true
}


