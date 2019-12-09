package router

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/service"
)

type StatisticRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth     *controller.Middleware

	FeeXY *controller.LeafRouter
	CountXY *controller.LeafRouter
}

func BuildStatisticRouter(parent *RootRouter, serviceProvider *service.Provider) (router *StatisticRouter) {

	statisticService := serviceProvider.StatisticService()

	router = &StatisticRouter{
		Router:     parent.Router.Group("statistic"),
		AuthRouter: parent.AuthRouter.Group("statistic"),
		Auth:       parent.Auth.Copy(),
	}

	router.FeeXY = router.AuthRouter.GET("/fee", statisticService.StatGoodsFeeXY)
	router.CountXY = router.AuthRouter.GET("/count", statisticService.StatGoodsCountXY)
	return
}

func (s *Provider) StatisticRouter() *StatisticRouter {
	return s.statisticRouter
}
