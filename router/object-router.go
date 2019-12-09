package router

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/service"
)

type ObjectRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth     *controller.Middleware
	IDRouter *ObjectIDRouter

	Post    *controller.LeafRouter
	GetList *controller.LeafRouter
}

type ObjectIDRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth *controller.Middleware

	Get    *controller.LeafRouter
	Put    *controller.LeafRouter
	Delete *controller.LeafRouter
}

func BuildObjectRouter(parent *RootRouter, serviceProvider *service.Provider) (router *ObjectRouter) {
	objectService := serviceProvider.ObjectService()
	router = &ObjectRouter{
		Router:     parent.Router.Extend("object"),
		AuthRouter: parent.AuthRouter.Extend("object"),
		Auth:       parent.Auth.Copy(),
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
		Router:     parent.Group("/object/:oid"),
		AuthRouter: parent.AuthRouter.Group("/object/:oid"),
		Auth:       parent.Auth.MustGroup("object", "oid"),
	}

	router.Get = router.GET("", objectService.Get)
	router.Put = router.AuthRouter.PUT("", objectService.Put)
	router.Delete = router.AuthRouter.DELETE("", objectService.Delete)
	return
}

func (s *Provider) ObjectRouter() *ObjectRouter {
	return s.objectRouter
}
