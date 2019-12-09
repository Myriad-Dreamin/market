package server

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/lib/jwt"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/gin-contrib/cors"
	//"github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	"strconv"
)

func (srv *Server) PrepareMiddleware() bool {
	srv.jwtMW = jwt.NewMiddleWare(func() *jwt.CustomClaims {
		var cc = new(jwt.CustomClaims)
		cc.CustomField = &types.CustomFields{}
		return cc
	}, func(c controller.MContext, cc *jwt.CustomClaims) error {
		c.Set("uid", strconv.FormatInt(cc.CustomField.(*types.CustomFields).UID, 10))
		return nil
	})
	srv.jwtMW.ExpireSecond = 3600
	srv.jwtMW.RefreshSecond = 3600 * 24 * 7

	srv.routerAuthMW = controller.NewMiddleware(srv.DatabaseProvider.Enforcer(),
		"user:", "uid", ginhelper.MissID, ginhelper.AuthFailed)

	srv.corsMW = cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"http://127.0.0.1:80", "https://127.0.0.1:80"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	})

	return true
}
