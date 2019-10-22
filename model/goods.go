package model

import (
	"github.com/Myriad-Dreamin/ginx/config"
	splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Goods = splayer.Goods
type GoodsDB = splayer.GoodsDB

func NewGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	return splayer.NewGoodsDB(logger, cfg)
}

func GetGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	return splayer.GetGoodsDB(logger, cfg)
}