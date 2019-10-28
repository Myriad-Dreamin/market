package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/rbac"
	"github.com/Myriad-Dreamin/market/types"
)

type dbResult struct {
	dbName string
	types.DecayResult
}

func (srv *Server) registerDatabaseService() bool {

	for _, dbResult := range []dbResult{
		{"needsDB", types.Decay(model.NewNeedsDB(srv.Logger, srv.cfg))},
		{"goodsDB", types.Decay(model.NewGoodsDB(srv.Logger, srv.cfg))},
		{"userDB", types.Decay(model.NewUserDB(srv.Logger, srv.cfg))},
		{"objectDB", types.Decay(model.NewObjectDB(srv.Logger, srv.cfg))},
	} {
		if dbResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("init %T DB error", dbResult.First), "error", dbResult.Err)
			return false
		}
		srv.DatabaseProvider.Register(dbResult.dbName, dbResult.First)
	}
	return true
}

func (srv *Server) PrepareDatabase() bool {
	var err error
	cfg := srv.cfg
	srv.DB, err = model.OpenORM(cfg)
	if err != nil {
		srv.Logger.Error("open database error", "error", err)
		return false
	}

	srv.Logger.Info("connected to database",
		"connection-type", cfg.DatabaseConfig.ConnectionType,
		"user", cfg.DatabaseConfig.User,
		"database", cfg.DatabaseConfig.DatabaseName,
		"charset", cfg.DatabaseConfig.Charset,
		"location", cfg.DatabaseConfig.Location,
	)

	err = model.Register(srv.DB, srv.Logger)
	if err != nil {
		srv.Logger.Error("register and migrate error", "error", err)
		return false
	}

	//srv.RedisPool, err = model.OpenRedis(cfg)
	//if err != nil {
	//	srv.Logger.Debug("create redis pool error", "error", err)
	//	return false
	//}
	//
	//srv.Logger.Info("connected to redis",
	//	"connection-type", cfg.RedisConfig.ConnectionType,
	//	"host", cfg.RedisConfig.Host,
	//	"connection-timeout", cfg.RedisConfig.ConnectionTimeout,
	//	"database", cfg.RedisConfig.Database,
	//	"read-timeout", cfg.RedisConfig.ReadTimeout,
	//	"write-timeout", cfg.RedisConfig.WriteTimeout,
	//	"idle-timeout", cfg.RedisConfig.IdleTimeout,
	//	"wait", cfg.RedisConfig.Wait,
	//	"max-active", cfg.RedisConfig.MaxActive,
	//	"max-idle", cfg.RedisConfig.MaxIdle,
	//)
	//err = model.RegisterRedis(srv.RedisPool, srv.Logger)
	//if err != nil {
	//	srv.Logger.Debug("register redis error", "error", err)
	//	return false
	//}
	err = rbac.InitGorm(srv.DB, "rbac.conf")
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ServiceProvider.Register(rbac.GetEnforcer())


	return srv.registerDatabaseService()
}


func (srv *Server) MockDatabase() bool {
	var err error
	srv.DB, err = model.MockORM(srv.cfg)
	if err != nil {
		srv.Logger.Error("open database error", "error", err)
		return false
	}

	err = model.Register(srv.DB, srv.Logger)
	if err != nil {
		srv.Logger.Error("register and migrate error", "error", err)
		return false
	}

	err = rbac.InitGorm(srv.DB, "rbac.conf")
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.DatabaseProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.registerDatabaseService()
}

