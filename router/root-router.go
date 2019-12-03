package router

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/Myriad-Dreamin/market/service"
	"github.com/gin-gonic/gin"
)

type RootRouter struct {
	Root       *mgin.Router
	Router     *mgin.Router
	AuthRouter *mgin.Router
	Auth       *mgin.Middleware

	Ping *mgin.LeafRouter
	//ObjectRouter *ObjectRouter
	UserRouter  *UserRouter
	GoodsRouter *GoodsRouter
	StatisticRouter *StatisticRouter
	NeedsRouter *NeedsRouter
}

// @title Ping
// @description result
func PingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewRootRouter(serviceProvider *service.Provider, jwtMW *jwt.Middleware, routerAuthMW *mgin.Middleware) (r *RootRouter) {
	rr := mgin.NewRouterGroup()
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