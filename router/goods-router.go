package router

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/service"
)

type GoodsRouter struct {
	Router *controller.Router
	AuthRouter *controller.Router
	Auth     *controller.Middleware
	IDRouter *GoodsIDRouter

	Post    *controller.LeafRouter
	GetList *controller.LeafRouter
}

type GoodsIDRouter struct {
	Router *controller.Router
	AuthRouter *controller.Router
	Auth *controller.Middleware

	Get         *controller.LeafRouter
	Inspect     *controller.LeafRouter
	Put         *controller.LeafRouter
	Delete      *controller.LeafRouter
	ForceDelete *controller.LeafRouter
	PutPicture  *controller.LeafRouter
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
	return
}

func (s *Provider) GoodsRouter() *GoodsRouter {
	return s.goodsRouter
}
