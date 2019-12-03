package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/market/service"
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"statisticService", functional.Decay(service.NewStatisticService(srv.Logger, srv.DatabaseProvider, srv.Cfg))},
		{"needsService", functional.Decay(service.NewNeedsService(srv.Logger, srv.DatabaseProvider, srv.Cfg))},
		{"goodsService", functional.Decay(service.NewGoodsService(srv.Logger, srv.DatabaseProvider, srv.Cfg))},
		{"userService", functional.Decay(service.NewUserService(srv.Logger, srv.DatabaseProvider, srv.jwtMW, srv.Cfg))},
		{"objectService", functional.Decay(service.NewObjectService(srv.Logger, srv.DatabaseProvider, srv.Cfg))},
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
