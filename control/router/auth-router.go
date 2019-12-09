package router

import (
	"github.com/Myriad-Dreamin/market/service"
)

type AuthRouter struct {
	*Router
	AuthRouter *Router
	Auth     *Middleware
}

func BuildAuthRouter(parent *RootRouter, serviceProvider *service.Provider) (router *AuthRouter) {
	router = &AuthRouter{
		Router: parent.Router.Extend("auth"),
		AuthRouter: parent.AuthRouter.Extend("auth"),
		Auth:   parent.Auth.Copy(),
	}
	return
}

func (s *Provider) AuthRouter() *AuthRouter {
	return s.authRouter
}
