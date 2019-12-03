package router

import (
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/Myriad-Dreamin/market/service"
)

type UserRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth     *mgin.Middleware
	IDRouter *UserIDRouter

	Login    *mgin.LeafRouter
	Register *mgin.LeafRouter
	GetList  *mgin.LeafRouter
}

type UserIDRouter struct {
	*mgin.Router
	AuthRouter *mgin.Router
	Auth *mgin.Middleware

	ChangePassword *mgin.LeafRouter
	Get            *mgin.LeafRouter
	Put            *mgin.LeafRouter
	Delete         *mgin.LeafRouter
}

func BuildUserRouter(parent *RootRouter, serviceProvider *service.Provider) (router *UserRouter) {
	userService := serviceProvider.UserService()
	router = &UserRouter{
		Router: parent.Router.Extend("user"),
		AuthRouter: parent.AuthRouter.Extend("user"),
		Auth:   parent.Auth.Copy(),
	}
	router.GetList = router.GET("user-list", userService.List)
	router.Register = router.POST("/user", userService.Register)
	router.Login = router.POST("/login", userService.Login)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*UserIDRouter) subBuild(parent *UserRouter, serviceProvider *service.Provider) (
	router *UserIDRouter) {

	userService := serviceProvider.UserService()

	router = &UserIDRouter{
		Router: parent.Group("/user/:id"),
		AuthRouter: parent.AuthRouter.Group("/user/:id"),
		Auth:   parent.Auth.MustGroup("user", "id"),
	}

	router.Get = router.GET("", userService.Get)
	router.ChangePassword = router.AuthRouter.PUT("/password", userService.ChangePassword)
	router.Put = router.AuthRouter.PUT("", userService.Put)
	router.Delete = router.AuthRouter.DELETE("", userService.Delete)
	return
}

func (s *Provider) UserRouter() *UserRouter {
	return s.userRouter
}
