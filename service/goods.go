package service

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	goodsservice "github.com/Myriad-Dreamin/ginx/service/goods"
	"github.com/Myriad-Dreamin/ginx/types"
)

type GoodsService = goodsservice.Service

func NewGoodsService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*GoodsService, error) {
	return goodsservice.NewService(logger, provider, config)
}

func (s *Provider) GoodsService() *GoodsService {
	return s.goodsService
}
