package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type StatisticRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	FeeXY   *LeafRouter
	CountXY *LeafRouter
}

func BuildStatisticRouter(parent H, serviceProvider *service.Provider) (router *StatisticRouter) {

	statisticService := serviceProvider.StatisticService()

	router = &StatisticRouter{
		Router:     parent.GetRouter().Group("statistic"),
		AuthRouter: parent.GetAuthRouter().Group("statistic"),
		Auth:       parent.GetAuth().Copy(),
	}

	router.FeeXY = router.AuthRouter.GET("/fee", statisticService.StatGoodsFeeXY)
	router.CountXY = router.AuthRouter.GET("/count", statisticService.StatGoodsCountXY)
	return
}

func (s *Provider) StatisticRouter() *StatisticRouter {
	return s.statisticRouter
}
