package service

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/control"
	"github.com/Myriad-Dreamin/market/model"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/market/types"
)

type GoodsService = control.GoodsService

func NewGoodsService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (GoodsService, error) {
	return goodsservice.NewService(logger, provider, config)
}

func (s *Provider) GoodsService() GoodsService {
	return s.goodsService
}
