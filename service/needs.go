package service

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control"
	"github.com/Myriad-Dreamin/market/model"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/market/types"
)

type NeedsService = control.NeedsService

func NewNeedsService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (NeedsService, error) {
	return needsservice.NewService(logger, provider, config)
}

func (s *Provider) NeedsService() NeedsService {
	return s.needsService
}
