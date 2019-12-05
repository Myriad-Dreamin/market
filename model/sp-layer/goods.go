package splayer

import (
	"github.com/Myriad-Dreamin/market/config"
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type Goods = dblayer.Goods

type GoodsDB struct {
	dblayer.GoodsDB
}

func NewGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	cdb, err := dblayer.NewGoodsDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(GoodsDB)
	db.GoodsDB = *cdb
	return db, nil
}

func GetGoodsDB(logger types.Logger, cfg *config.ServerConfig) (*GoodsDB, error) {
	cdb, err := dblayer.GetGoodsDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(GoodsDB)
	db.GoodsDB = *cdb
	return db, nil
}

func (s *Provider) GoodsDB() *GoodsDB {
	return s.goodsDB
}
