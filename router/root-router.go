package router

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	"github.com/Myriad-Dreamin/market/service"
	"github.com/gin-gonic/gin"
)

type RootRouter struct {
	Root       *controller.Router
	Router     *controller.Router
	AuthRouter *controller.Router
	Auth       *controller.Middleware

	Ping *controller.LeafRouter
	//ObjectRouter *ObjectRouter
	UserRouter  *UserRouter
	GoodsRouter *GoodsRouter
	StatisticRouter *StatisticRouter
	NeedsRouter *NeedsRouter
}

// @title Ping
// @description result
func PingFunc(c controller.MContext) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewRootRouter(serviceProvider *service.Provider, jwtMW *jwt.Middleware, routerAuthMW *controller.Middleware) (r *RootRouter) {
	rr := controller.NewRouterGroup()
	apiRouterV1 := rr.Group("/v1")
	authRouterV1 := apiRouterV1.Group("", jwtMW.Build())

	r = &RootRouter{
		Root:       rr,
		Router:     apiRouterV1,
		AuthRouter: authRouterV1,
		Auth:       routerAuthMW,
	}

	r.Ping = r.Root.GET("/ping", PingFunc)

	//r.ObjectRouter = BuildObjectRouter(r, serviceProvider)
	r.UserRouter = BuildUserRouter(r, serviceProvider)
	r.NeedsRouter = BuildNeedsRouter(r, serviceProvider)
	r.GoodsRouter = BuildGoodsRouter(r, serviceProvider)
	r.StatisticRouter = BuildStatisticRouter(r, serviceProvider)


	ApplyAuth(r)
	return
}