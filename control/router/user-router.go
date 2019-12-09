package router

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/service"
)

type UserRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth       *controller.Middleware
	IDRouter   *UserIDRouter

	Login    *controller.LeafRouter
	Register *controller.LeafRouter
	GetList  *controller.LeafRouter
}

type UserIDRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth       *controller.Middleware
	GoodsIDRouter *UserGoodsIDRouter
	NeedsIDRouter *UserNeedsIDRouter

	ChangePassword *controller.LeafRouter
	Get            *controller.LeafRouter
	Put            *controller.LeafRouter
	Delete         *controller.LeafRouter
}

type UserGoodsIDRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth       *controller.Middleware

	Buy            *controller.LeafRouter
	ConfirmBuy     *controller.LeafRouter
}

type UserNeedsIDRouter struct {
	*controller.Router
	AuthRouter *controller.Router
	Auth       *controller.Middleware

	Sell           *controller.LeafRouter
	ConfirmSell    *controller.LeafRouter
}

func BuildUserRouter(parent *RootRouter, serviceProvider *service.Provider) (router *UserRouter) {
	userService := serviceProvider.UserService()
	router = &UserRouter{
		Router:     parent.Router.Extend("user"),
		AuthRouter: parent.AuthRouter.Extend("user"),
		Auth:       parent.Auth.Copy(),
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
