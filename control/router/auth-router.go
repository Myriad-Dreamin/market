package router

import (
	"github.com/Myriad-Dreamin/market/service"
	"github.com/Myriad-Dreamin/market/types"
)

type AuthRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	RefreshToken *LeafRouter
	Group        *AuthGroupRouter
	Sugar 		*AuthSugarRouter
}

type AuthGroupRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	GrantGroup  *LeafRouter
	RevokeGroup *LeafRouter
	CheckGroup  *LeafRouter
}

type AuthSugarRouterApiGroup struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Grant *LeafRouter
	Revoke *LeafRouter
	Check *LeafRouter
}

type AuthSugarRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	DynamicList map[string]*AuthSugarRouterApiGroup
}

func BuildAuthRouter(parent *RootRouter, serviceProvider *service.Provider) (router *AuthRouter) {
	authService := serviceProvider.AuthService()

	router = &AuthRouter{
		Router:     parent.GetRouter().Group("auth"),
		AuthRouter: parent.GetAuthRouter().Group("auth"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.RefreshToken = router.AuthRouter.GET("/refresh-token", authService.RefreshToken)

	router.Group = router.Group.subBuild(router, serviceProvider)
	router.Sugar = router.Sugar.subBuild(router, serviceProvider)
	return
}

func (*AuthSugarRouter) subBuild(parent *AuthRouter, serviceProvider *service.Provider) *AuthSugarRouter {
	authService := serviceProvider.AuthService()

	router := &AuthSugarRouter{
		Router:     parent.Router.Group("sugar"),
		AuthRouter: parent.AuthRouter.Group("sugar"),
		Auth:       parent.Auth.Copy(),
	}
	router.DynamicList = make(map[string]*AuthSugarRouterApiGroup)
	for _, groupName := range types.Groups {
		dyn := &AuthSugarRouterApiGroup{
			Router:     router.Router.Group("/group/" + groupName + "/user/:id"),
			AuthRouter: router.AuthRouter.Group("/group/" + groupName + "/user/:id"),
			Auth:       router.Auth.MustGroup("user", "id"),
		}
		router.DynamicList[groupName] = dyn
		dyn.Grant = dyn.AuthRouter.PUT("", authService.GroupGranter(groupName))
		dyn.Check = dyn.Router.GET("", authService.GroupChecker(groupName))
		dyn.Revoke = dyn.AuthRouter.DELETE("", authService.GroupRevoker(groupName))
	}

	return router
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
