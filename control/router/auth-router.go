package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type AuthRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	RefreshToken *LeafRouter
	Group        *AuthGroupRouter
}

type AuthGroupRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	GrantGroup  *LeafRouter
	RevokeGroup *LeafRouter
	CheckGroup  *LeafRouter
}

func BuildAuthRouter(parent *RootRouter, serviceProvider *service.Provider) (router *AuthRouter) {
	authService := serviceProvider.AuthService()

	router = &AuthRouter{
		Router:     parent.Router.Group("auth"),
		AuthRouter: parent.AuthRouter.Group("auth"),
		Auth:       parent.Auth.Copy(),
	}
	router.RefreshToken = router.AuthRouter.GET("/refresh-token", authService.RefreshToken)

	router.Group = router.Group.subBuild(router, serviceProvider)
	return
}

func (*AuthGroupRouter) subBuild(parent *AuthRouter, serviceProvider *service.Provider) (
	router *AuthGroupRouter) {

	authService := serviceProvider.AuthService()

	router = &AuthGroupRouter{
		Router:     parent.Router.Group("/group/user/:id"),
		AuthRouter: parent.AuthRouter.Group("/group/user/:id"),
		Auth:       parent.Auth.MustGroup("user", "id"),
	}
	router.GrantGroup = router.AuthRouter.PUT("", authService.GrantGroup)
	router.RevokeGroup = router.AuthRouter.DELETE("", authService.RevokeGroup)
	router.CheckGroup = router.Router.GET("", authService.CheckGroup)
	return router
}

func (s *Provider) AuthRouter() *AuthRouter {
	return s.authRouter
}
