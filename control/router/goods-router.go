package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type GoodsRouter struct {
	Router *Router
	AuthRouter *Router
	Auth     *Middleware
	IDRouter *GoodsIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type GoodsIDRouter struct {
	Router *Router
	AuthRouter *Router
	Auth *Middleware

	Get         *LeafRouter
	Inspect     *LeafRouter
	Put         *LeafRouter
	Delete      *LeafRouter
	ForceDelete *LeafRouter
	PutPicture  *LeafRouter
	GetPicture  *LeafRouter
}

func BuildGoodsRouter(parent *RootRouter, serviceProvider *service.Provider) (router *GoodsRouter) {
	goodsService := serviceProvider.GoodsService()
	router = &GoodsRouter{
		Router: parent.Router.Extend("goods"),
		AuthRouter: parent.AuthRouter.Extend("goods"),
		Auth:   parent.Auth.Copy(),
	}
	router.GetList = router.Router.GET("goods-list", goodsService.List)
	router.Post = router.AuthRouter.POST("/goods", goodsService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*GoodsIDRouter) subBuild(parent *GoodsRouter, serviceProvider *service.Provider) (
	router *GoodsIDRouter) {

	goodsService := serviceProvider.GoodsService()

	router = &GoodsIDRouter{
		Router: parent.Router.Group("/goods/:goid"),
		AuthRouter: parent.AuthRouter.Group("/goods/:goid"),
		Auth:   parent.Auth.MustGroup("goods", "goid"),
	}

	router.Get = router.Router.GET("", goodsService.Get)
	router.Inspect = router.AuthRouter.GET("/inspect", goodsService.Inspect)
	router.Put = router.AuthRouter.PUT("", goodsService.Put)
	// todo
	router.Delete = router.AuthRouter.DELETE("", goodsService.Delete)
	router.ForceDelete = router.AuthRouter.DELETE("/force", goodsService.ForceDelete)
	router.PutPicture = router.AuthRouter.PUT("/picture", goodsService.PutPicture)
	router.GetPicture = router.AuthRouter.GET("/picture", goodsService.GetPicture)
	return
}

func (s *Provider) GoodsRouter() *GoodsRouter {
	return s.goodsRouter
}
