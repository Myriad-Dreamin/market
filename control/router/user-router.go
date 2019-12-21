package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type UserRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *UserIDRouter

	Login     *LeafRouter
	Register  *LeafRouter
	GetList   *LeafRouter
	GetCities *LeafRouter
}

type UserIDRouter struct {
	*Router
	AuthRouter    *Router
	Auth          *Middleware
	GoodsIDRouter *UserGoodsIDRouter
	NeedsIDRouter *UserNeedsIDRouter

	ChangePassword *LeafRouter
	Get            *LeafRouter
	Put            *LeafRouter
	Delete         *LeafRouter
}

type UserGoodsIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Buy        *LeafRouter
	ConfirmBuy *LeafRouter
}

type UserNeedsIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Sell        *LeafRouter
	ConfirmSell *LeafRouter
}

func BuildUserRouter(parent *RootRouter, serviceProvider *service.Provider) (router *UserRouter) {
	userService := serviceProvider.UserService()
	router = &UserRouter{
		Router:     parent.Router.Extend("user"),
		AuthRouter: parent.AuthRouter.Extend("user"),
		Auth:       parent.Auth.Copy(),
	}
	router.GetList = router.GET("user-list", userService.List)
	router.GetCities = router.GET("user-cities", userService.GetCities)
	router.Register = router.POST("/user", userService.Register)
	router.Login = router.POST("/login", userService.Login)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*UserIDRouter) subBuild(parent *UserRouter, serviceProvider *service.Provider) (
	router *UserIDRouter) {

	userService := serviceProvider.UserService()

	router = &UserIDRouter{
		Router:     parent.Group("/user/:id"),
		AuthRouter: parent.AuthRouter.Group("/user/:id"),
		Auth:       parent.Auth.MustGroup("user", "id"),
	}

	router.Get = router.GET("", userService.Get)
	router.ChangePassword = router.AuthRouter.PUT("/password", userService.ChangePassword)
	router.Put = router.AuthRouter.PUT("", userService.Put)
	router.Delete = router.AuthRouter.DELETE("", userService.Delete)

	router.GoodsIDRouter = router.GoodsIDRouter.subBuild(router, serviceProvider)
	router.NeedsIDRouter = router.NeedsIDRouter.subBuild(router, serviceProvider)

	return
}

func (t *UserGoodsIDRouter) subBuild(parent *UserIDRouter, serviceProvider *service.Provider) *UserGoodsIDRouter {

	userService := serviceProvider.UserService()

	router := &UserGoodsIDRouter{
		Router:     parent.Group("/goods/:goid"),
		AuthRouter: parent.AuthRouter.Group("/goods/:goid"),
		Auth:       parent.Auth.MustGroup("goods", "goid"),
	}

	router.Buy = router.AuthRouter.POST("/buy", userService.Buy)
	router.ConfirmBuy = router.AuthRouter.POST("/confirm-buy", userService.ConfirmBuy)
	return router
}

func (t *UserNeedsIDRouter) subBuild(parent *UserIDRouter, serviceProvider *service.Provider) *UserNeedsIDRouter {

	userService := serviceProvider.UserService()

	router := &UserNeedsIDRouter{
		Router:     parent.Group("/needs/:nid"),
		AuthRouter: parent.AuthRouter.Group("/needs/:nid"),
		Auth:       parent.Auth.MustGroup("needs", "nid"),
	}

	router.Sell = router.AuthRouter.POST("/sell", userService.Sell)
	router.ConfirmSell = router.AuthRouter.POST("/confirm-sell", userService.ConfirmSell)
	return router
}

func (s *Provider) UserRouter() *UserRouter {
	return s.userRouter
}
