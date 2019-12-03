package service

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model"
	statisticservice "github.com/Myriad-Dreamin/market/service/statistic"
	"github.com/Myriad-Dreamin/market/types"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type StatisticService = statisticservice.Service

func NewStatisticService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*StatisticService, error) {
	return statisticservice.NewService(logger, provider, config)
}

func (s *Provider) StatisticService() *StatisticService {
	return s.statisticService
}
