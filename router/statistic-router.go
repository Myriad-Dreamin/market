package router

import (
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/Myriad-Dreamin/market/service"
)

type StatisticRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth     *mgin.Middleware

	FeeXY *mgin.LeafRouter
	CountXY *mgin.LeafRouter
}

func BuildStatisticRouter(parent *RootRouter, serviceProvider *service.Provider) (router *StatisticRouter) {

	statisticService := serviceProvider.StatisticService()

	router = &StatisticRouter{
		Router: parent.Router.Group("statistic"),
		AuthRouter: parent.AuthRouter.Group("statistic"),
		Auth:   parent.Auth.Copy(),
	}

	router.FeeXY = router.AuthRouter.GET("/fee", statisticService.StatGoodsFeeXY)
	router.CountXY = router.AuthRouter.GET("/count", statisticService.StatGoodsCountXY)
	return
}

func (s *Provider) StatisticRouter() *StatisticRouter {
	return s.statisticRouter
}
