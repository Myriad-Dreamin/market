package service

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	needsservice "github.com/Myriad-Dreamin/ginx/service/needs"
	"github.com/Myriad-Dreamin/ginx/types"
)

type NeedsService = needsservice.Service

func NewNeedsService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*NeedsService, error) {
	return needsservice.NewService(logger, provider, config)
}

func (s *Provider) NeedsService() *NeedsService {
	return s.needsService
}
