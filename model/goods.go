package model

import (
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Goods = splayer.Goods
type GoodsDB = splayer.GoodsDB

func NewGoodsDB(m module.Module) (*GoodsDB, error) {
	return splayer.NewGoodsDB(m)
}

func GetGoodsDB(m module.Module) (*GoodsDB, error) {
	return splayer.GetGoodsDB(m)
}
