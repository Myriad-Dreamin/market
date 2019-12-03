package router

import (
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/Myriad-Dreamin/market/service"
)

type ObjectRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth     *mgin.Middleware
	IDRouter *ObjectIDRouter

	Post    *mgin.LeafRouter
	GetList *mgin.LeafRouter
}

type ObjectIDRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth *mgin.Middleware

	Get    *mgin.LeafRouter
	Put    *mgin.LeafRouter
	Delete *mgin.LeafRouter
}

func BuildObjectRouter(parent *RootRouter, serviceProvider *service.Provider) (router *ObjectRouter) {
	objectService := serviceProvider.ObjectService()
	router = &ObjectRouter{
		Router: parent.Router.Extend("object"),
		AuthRouter: parent.AuthRouter.Extend("object"),
		Auth:   parent.Auth.Copy(),
	}
	router.GetList = router.GET("object-list", objectService.List)
	router.Post = router.AuthRouter.POST("/object", objectService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*ObjectIDRouter) subBuild(parent *ObjectRouter, serviceProvider *service.Provider) (
	router *ObjectIDRouter) {

	objectService := serviceProvider.ObjectService()

	router = &ObjectIDRouter{
		Router: parent.Group("/object/:oid"),
		AuthRouter: parent.AuthRouter.Group("/object/:oid"),
		Auth:   parent.Auth.MustGroup("object", "oid"),
	}

	router.Get = router.GET("", objectService.Get)
	router.Put = router.AuthRouter.PUT("", objectService.Put)
	router.Delete = router.AuthRouter.DELETE("", objectService.Delete)
	return
}

func (s *Provider) ObjectRouter() *ObjectRouter {
	return s.objectRouter
}
