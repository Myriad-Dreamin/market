package plugin

import (
	"context"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	"github.com/Myriad-Dreamin/market/service"
	"github.com/Myriad-Dreamin/market/types"
)

type Logger = types.Logger
type ConfigLoader = types.ConfigLoader
type ServiceProvider = service.Provider
type DatabaseProvider = model.Provider
type ServerConfig = config.ServerConfig
type Module = types.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin)
	Inject(services *ServiceProvider, dbs *DatabaseProvider, module Module) (plg Plugin)
	Work(ctx context.Context)
}
