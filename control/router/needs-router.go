package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type NeedsRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *NeedsIDRouter

	Post     *LeafRouter
	GetList  *LeafRouter
	GetTypes *LeafRouter
}

type NeedsIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Get         *LeafRouter
	Inspect     *LeafRouter
	Put         *LeafRouter
	Delete      *LeafRouter
	ForceDelete *LeafRouter
	PutPicture  *LeafRouter
	GetPicture  *LeafRouter
}

func BuildNeedsRouter(parent *RootRouter, serviceProvider *service.Provider) (router *NeedsRouter) {
	needsService := serviceProvider.NeedsService()
	router = &NeedsRouter{
		Router:     parent.Router.Extend("needs"),
		AuthRouter: parent.AuthRouter.Extend("needs"),
		Auth:       parent.Auth.Copy(),
	}
	router.GetList = router.GET("needs-list", needsService.List)
	router.GetTypes = router.Router.GET("needs-types", needsService.GetTypes)
	router.Post = router.AuthRouter.POST("/needs", needsService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*NeedsIDRouter) subBuild(parent *NeedsRouter, serviceProvider *service.Provider) (
	router *NeedsIDRouter) {

	needsService := serviceProvider.NeedsService()

	router = &NeedsIDRouter{
		Router:     parent.Group("/needs/:nid"),
		AuthRouter: parent.AuthRouter.Group("/needs/:nid"),
		Auth:       parent.Auth.MustGroup("needs", "nid"),
	}

	router.Get = router.GET("", needsService.Get)
	router.Inspect = router.AuthRouter.GET("/inspect", needsService.Inspect)
	router.Put = router.AuthRouter.PUT("", needsService.Put)
	router.Delete = router.AuthRouter.DELETE("", needsService.Delete)
	router.ForceDelete = router.AuthRouter.DELETE("/force", needsService.ForceDelete)
	router.PutPicture = router.AuthRouter.PUT("/picture", needsService.PutPicture)
	router.GetPicture = router.AuthRouter.GET("/picture", needsService.GetPicture)
	return
}

func (s *Provider) NeedsRouter() *NeedsRouter {
	return s.needsRouter
}
