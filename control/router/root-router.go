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

type BaseH struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

}

func (r *BaseH) GetRouter() *Router {
	return r.Router
}

func (r *BaseH) GetAuthRouter() *Router {
	return r.AuthRouter
}

func (r *BaseH) GetAuth() *Middleware {
	return r.Auth
}

type RootRouter struct {
	H
	Root       *Router

	//ObjectRouter *ObjectRouter
	AuthApiRouter *AuthRouter
	UserRouter      *UserRouter
	GoodsRouter     *GoodsRouter
	StatisticRouter *StatisticRouter
	NeedsRouter     *NeedsRouter

	Ping *LeafRouter
	GoodsPictureFiles *LeafRouter
	NeedsPictureFiles *LeafRouter

	V2 *RootRouterV2
}


type RootRouterV2 struct {
	H

	ConstRouter *ConstRouter
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
	apiRouterV2 := rr.Group("/v2")
	b := m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware).Build()
	authRouterV1 := apiRouterV1.Group("", b)
	authRouterV2 := apiRouterV2.Group("", b)


	r = &RootRouter{
		Root:       rr,
		H: &BaseH{
			Router:     apiRouterV1,
			AuthRouter: authRouterV1,
			Auth:       m.Require(config.ModulePath.Middleware.RouteAuth).(*Middleware),
		},
		V2: &RootRouterV2{
			H: &BaseH{
				Router: apiRouterV2,
				AuthRouter: authRouterV2,
				Auth: m.Require(config.ModulePath.Middleware.RouteAuth).(*Middleware),
			},
		},
	}

	r.Ping = r.Root.GET("/ping", PingFunc)

	//r.ObjectRouter = BuildObjectRouter(r, serviceProvider)

	serviceProvider := m.Require(config.ModulePath.Provider.Service).(*service.Provider)

	r.UserRouter = BuildUserRouter(r, serviceProvider)
	r.NeedsRouter = BuildNeedsRouter(r, serviceProvider)
	r.GoodsRouter = BuildGoodsRouter(r, serviceProvider)
	r.StatisticRouter = BuildStatisticRouter(r, serviceProvider)
	r.AuthApiRouter = BuildAuthRouter(r, serviceProvider)
	r.V2.ConstRouter = BuildConstRouter(r.V2, serviceProvider)

	cfg := m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)

	r.GoodsPictureFiles = r.GetAuthRouter().StaticFS("goods-picture", http.Dir(cfg.BaseParametersConfig.GoodsPicturePath))
	r.NeedsPictureFiles = r.GetAuthRouter().StaticFS("needs-picture", http.Dir(cfg.BaseParametersConfig.NeedsPicturePath))

	ApplyAuth(r)
	return
}
