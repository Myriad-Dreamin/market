package splayer

import (
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Goods = dblayer.Goods

type GoodsDB struct {
	dblayer.GoodsDB
}

func NewGoodsDB(m module.Module) (*GoodsDB, error) {
	cdb, err := dblayer.NewGoodsDB(m)
	if err != nil {
		return nil, err
	}
	db := new(GoodsDB)
	db.GoodsDB = *cdb
	return db, nil
}

func GetGoodsDB(m module.Module) (*GoodsDB, error) {
	cdb, err := dblayer.GetGoodsDB(m)
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
