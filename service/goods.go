package service

import (
	"github.com/Myriad-Dreamin/market/control"
	goodsservice "github.com/Myriad-Dreamin/market/service/goods"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type GoodsService = control.GoodsService

func NewGoodsService(m module.Module) (GoodsService, error) {
	return goodsservice.NewService(m)
}

func (s *Provider) GoodsService() GoodsService {
	return s.goodsService
}
