package model

import (
	"github.com/Myriad-Dreamin/market/config"
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type Goods = splayer.Goods
type GoodsDB = splayer.GoodsDB

func NewGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	return splayer.NewGoodsDB(logger, cfg)
}

func GetGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	return splayer.GetGoodsDB(logger, cfg)
}
