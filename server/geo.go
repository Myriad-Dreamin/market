package server

import "github.com/Myriad-Dreamin/market/config"

func (srv *Server) PrepareGeography() bool {
	srv.Module.Provide(config.ModulePath.Global.Cities, config.Cities)
	return true
}
