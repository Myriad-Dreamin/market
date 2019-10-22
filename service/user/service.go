package userservice

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.UserDB
	logger types.Logger
}


func NewService(logger types.Logger, provider *model.Provider, _ *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.db = provider.UserDB()
	a.logger = logger
	a.CRUDService = base_service.NewCRUDService(a, "id")
	a.ListService = base_service.NewListService(a, "id")
	return
}

/*
type User struct {
}
*/
