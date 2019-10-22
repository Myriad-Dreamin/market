package router

import (
	"github.com/Myriad-Dreamin/ginx/service"
)

type GoodsRouter struct {
	*Router
	Auth       *Middleware
	IDRouter   *GoodsIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type GoodsIDRouter struct {
	*Router
	Auth *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildGoodsRouter(parent *RootRouter, serviceProvider *service.Provider) (router *GoodsRouter) {
	goodsService := serviceProvider.GoodsService()
	router = &GoodsRouter{
		Router:     parent.Router.Extend("goods"),
		Auth:       parent.Auth.Copy(),
	}
	router.GetList = router.GET("goods-list", goodsService.List)
	router.Post = router.POST("/goods", goodsService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*GoodsIDRouter) subBuild(parent *GoodsRouter, serviceProvider *service.Provider) (
	router *GoodsIDRouter) {

	goodsService := serviceProvider.GoodsService()

	router = &GoodsIDRouter{
		Router: parent.Group("/goods/:goid"),
		Auth:   parent.Auth.MustGroup("goods", "goid"),
	}

	router.Get = router.GET("", goodsService.Get)
	router.Put = router.PUT("", goodsService.Put)
	router.Delete = router.DELETE("", goodsService.Delete)
	return
}

func (s *Provider) GoodsRouter() *GoodsRouter {
	return s.goodsRouter
}
