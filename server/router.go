package server

import (
	"github.com/Myriad-Dreamin/market/router"
	"github.com/gin-gonic/gin"
)



func (srv *Server) BuildRouter() bool {

	srv.RouterEngine = gin.Default()
	srv.RouterEngine.Use(srv.corsMW)

	srv.Router = router.NewRootRouter(srv.ServiceProvider, srv.jwtMW, srv.routerAuthMW)
	return true
}
