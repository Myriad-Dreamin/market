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
		{"authService", functional.Decay(service.NewAuthService(srv.Module))},
		{"statisticService", functional.Decay(service.NewStatisticService(srv.Module))},
		{"needsService", functional.Decay(service.NewNeedsService(srv.Module))},
		{"goodsService", functional.Decay(service.NewGoodsService(srv.Module))},
		{"userService", functional.Decay(service.NewUserService(srv.Module))},
		{"objectService", functional.Decay(service.NewObjectService(srv.Module))},
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
