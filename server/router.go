package server

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control/router"
	"github.com/gin-gonic/gin"
)

func (srv *Server) BuildRouter() bool {
	gin.DefaultErrorWriter = srv.LoggerWriter
	gin.DefaultWriter = srv.LoggerWriter
	srv.HttpEngine = gin.New()
	srv.HttpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: srv.LoggerWriter,
	}), gin.Recovery())
	srv.HttpEngine.Use(srv.corsMW)

	srv.Router = router.NewRootRouter(srv.ServiceProvider, srv.jwtMW, srv.routerAuthMW)
	srv.Module.Provide(config.ModulePath.Global.HttpEngine, srv.HttpEngine)
	return true
}
