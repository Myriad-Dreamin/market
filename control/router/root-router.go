package router

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	"github.com/Myriad-Dreamin/market/service"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RootRouter struct {
	Root       *Router
	Router     *Router
	AuthRouter *Router
	Auth       *Middleware

	Ping *LeafRouter
	//ObjectRouter *ObjectRouter
	AuthApiRouter *AuthRouter
	UserRouter      *UserRouter
	GoodsRouter     *GoodsRouter
	StatisticRouter *StatisticRouter
	NeedsRouter     *NeedsRouter


	GoodsPictureFiles *LeafRouter
	NeedsPictureFiles *LeafRouter
}

// @title Ping
// @description result
func PingFunc(c controller.MContext) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewRootRouter(m module.Module) (r *RootRouter) {
	rr := controller.NewRouterGroup()
	apiRouterV1 := rr.Group("/v1")
	authRouterV1 := apiRouterV1.Group("", m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware).Build())

	r = &RootRouter{
		Root:       rr,
		Router:     apiRouterV1,
		AuthRouter: authRouterV1,
		Auth:       m.Require(config.ModulePath.Middleware.RouteAuth).(*Middleware),
	}

	r.Ping = r.Root.GET("/ping", PingFunc)

	//r.ObjectRouter = BuildObjectRouter(r, serviceProvider)

	serviceProvider := m.Require(config.ModulePath.Provider.Service).(*service.Provider)

	r.UserRouter = BuildUserRouter(r, serviceProvider)
	r.NeedsRouter = BuildNeedsRouter(r, serviceProvider)
	r.GoodsRouter = BuildGoodsRouter(r, serviceProvider)
	r.StatisticRouter = BuildStatisticRouter(r, serviceProvider)
	r.AuthApiRouter = BuildAuthRouter(r, serviceProvider)

	cfg := m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)

	r.GoodsPictureFiles = r.AuthRouter.StaticFS("goods-picture", http.Dir(cfg.BaseParametersConfig.GoodsPicturePath))
	r.NeedsPictureFiles = r.AuthRouter.StaticFS("needs-picture", http.Dir(cfg.BaseParametersConfig.NeedsPicturePath))

	ApplyAuth(r)
	return
}
