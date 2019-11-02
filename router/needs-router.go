package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type NeedsRouter struct {
	*Router
	Auth     *Middleware
	IDRouter *NeedsIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type NeedsIDRouter struct {
	*Router
	Auth *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildNeedsRouter(parent *RootRouter, serviceProvider *service.Provider) (router *NeedsRouter) {
	needsService := serviceProvider.NeedsService()
	router = &NeedsRouter{
		Router: parent.Router.Extend("needs"),
		Auth:   parent.Auth.Copy(),
	}
	router.GetList = router.GET("needs-list", needsService.List)
	router.Post = router.POST("/needs", needsService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*NeedsIDRouter) subBuild(parent *NeedsRouter, serviceProvider *service.Provider) (
	router *NeedsIDRouter) {

	needsService := serviceProvider.NeedsService()

	router = &NeedsIDRouter{
		Router: parent.Group("/needs/:nid"),
		Auth:   parent.Auth.MustGroup("needs", "nid"),
	}

	router.Get = router.GET("", needsService.Get)
	router.Put = router.PUT("", needsService.Put)
	router.Delete = router.DELETE("", needsService.Delete)
	return
}

func (s *Provider) NeedsRouter() *NeedsRouter {
	return s.needsRouter
}
