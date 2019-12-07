package service

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control"
	"github.com/Myriad-Dreamin/market/model"
	objectservice "github.com/Myriad-Dreamin/market/service/object"
	"github.com/Myriad-Dreamin/market/types"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ObjectService = control.ObjectService

func NewObjectService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (ObjectService, error) {
	return objectservice.NewService(logger, provider, config)
}

func (s *Provider) ObjectService() ObjectService {
	return s.objectService
}
