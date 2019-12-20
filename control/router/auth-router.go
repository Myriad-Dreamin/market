package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type AuthRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	RefreshToken *LeafRouter
}

func BuildAuthRouter(parent *RootRouter, serviceProvider *service.Provider) (router *AuthRouter) {
	authService := serviceProvider.AuthService()

	router = &AuthRouter{
		Router:     parent.Router.Extend("auth"),
		AuthRouter: parent.AuthRouter.Extend("auth"),
		Auth:       parent.Auth.Copy(),
	}
	router.RefreshToken = router.AuthRouter.GET("/refresh-token", authService.RefreshToken)
	return
}

func (s *Provider) AuthRouter() *AuthRouter {
	return s.authRouter
}
