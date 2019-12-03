package router

import (
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/Myriad-Dreamin/market/service"
)

type StatisticRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth     *mgin.Middleware
}

func BuildStatisticRouter(parent *RootRouter, serviceProvider *service.Provider) (router *StatisticRouter) {
	router = &StatisticRouter{
		Router: parent.Router.Extend("statistic"),
		AuthRouter: parent.AuthRouter.Extend("statistic"),
		Auth:   parent.Auth.Copy(),
	}
	return
}

func (s *Provider) StatisticRouter() *StatisticRouter {
	return s.statisticRouter
}
