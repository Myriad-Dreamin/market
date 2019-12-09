package service

import (
	"github.com/Myriad-Dreamin/market/control"
	needsservice "github.com/Myriad-Dreamin/market/service/needs"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type NeedsService = control.NeedsService

func NewNeedsService(m module.Module) (NeedsService, error) {
	return needsservice.NewService(m)
}

func (s *Provider) NeedsService() NeedsService {
	return s.needsService
}
