package service

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	userservice "github.com/Myriad-Dreamin/ginx/service/user"
	"github.com/Myriad-Dreamin/ginx/types"
)

type UserService = userservice.Service

func NewUserService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*UserService, error) {
	return userservice.NewService(logger, provider, config)
}

func (s *Provider) UserService() *UserService {
	return s.userService
}
