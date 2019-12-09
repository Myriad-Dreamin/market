package service

import (
	"github.com/Myriad-Dreamin/market/control"
	statisticservice "github.com/Myriad-Dreamin/market/service/statistic"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type StatisticService = control.StatisticService

func NewStatisticService(m module.Module) (StatisticService, error) {
	return statisticservice.NewService(m)
}

func (s *Provider) StatisticService() StatisticService {
	return s.statisticService
}
