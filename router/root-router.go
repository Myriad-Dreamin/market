package router

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/service"
)


type RootRouter struct {
	Root       *Router
	Router     *Router
	AuthRouter *Router
	Auth       *Middleware

	//ObjectRouter *ObjectRouter
	UserRouter *UserRouter
	GoodsRouter *GoodsRouter
}

func NewRootRouter(serviceProvider *service.Provider, jwtMW *jwt.Middleware, routerAuthMW *Middleware) (r *RootRouter) {
	rr := NewRouterGroup()
	apiRouterV1 := rr.Group("/v1")
	authRouterV1 := apiRouterV1.Group("", jwtMW.Build())

	r = &RootRouter{
		Root:       rr,
		Router:     apiRouterV1,
		AuthRouter: authRouterV1,
		Auth:       routerAuthMW,
	}
	//r.ObjectRouter = BuildObjectRouter(r, serviceProvider)
	r.UserRouter = BuildUserRouter(r, serviceProvider)
	r.GoodsRouter = BuildGoodsRouter(r, serviceProvider)
	return
}
