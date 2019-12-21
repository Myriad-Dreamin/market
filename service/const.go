package service

import (
	"github.com/Myriad-Dreamin/market/control"
	constservice "github.com/Myriad-Dreamin/market/service/const"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ConstService = control.ConstService

func NewConstService(m module.Module) (ConstService, error) {
	return constservice.NewService(m)
}

func (s *Provider) ConstService() ConstService {
	return s.constService
}
