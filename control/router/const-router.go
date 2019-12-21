package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type ConstRouter struct {
	*Router
	AuthRouter *Router
	Auth     *Middleware

	GetServiceCode *LeafRouter
	GetCities *LeafRouter
	GetGoodsTypes *LeafRouter
}

func BuildConstRouter(parent H, serviceProvider *service.Provider) (router *ConstRouter) {
	constService := serviceProvider.ConstService()
	router = &ConstRouter{
		Router: parent.GetRouter().Group("const"),
		AuthRouter: parent.GetAuthRouter().Group("const"),
		Auth:   parent.GetAuth().Copy(),
	}

	router.GetServiceCode = router.Router.GET("service-codes", constService.GetServiceCode)
	router.GetCities = router.Router.GET("user-cities", constService.GetCities)
	router.GetGoodsTypes = router.Router.GET("goods-types", constService.GetGoodsTypes)
	return
}

func (s *Provider) ConstRouter() *ConstRouter {
	return s.constRouter
}
