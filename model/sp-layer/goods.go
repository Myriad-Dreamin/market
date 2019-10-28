package splayer

import (
	"github.com/Myriad-Dreamin/market/config"
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type Goods = dblayer.Goods

//struct {
//	dblayer.Goods
//}

//func (c *Goods) Update() (int64, error) {
//	return c.Goods.Update()
//}

//type _userConverter struct {}
//var GoodsConverter _userConverter
//func (_userConverter) convert(obj dblayer.Goods) Goods { return Goods{Goods: obj} }
//func (_userConverter) reconvert(obj Goods) dblayer.Goods { return obj.Goods }
//func (_userConverter) converts(objs []dblayer.Goods) []Goods { return *(*[]Goods)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []Goods) []dblayer.Goods { return *(*[]dblayer.Goods)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.Goods) *Goods { return (*Goods)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *Goods) *dblayer.Goods { return (*dblayer.Goods)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.Goods) []*Goods { return *(*[]*Goods)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*Goods) []*dblayer.Goods { return *(*[]*dblayer.Goods)(unsafe.Pointer(&objs)) }

type GoodsDB struct {
	dblayer.GoodsDB
}

func NewGoodsDB(logger types.Logger, _ *config.ServerConfig) (*GoodsDB, error) {
	cdb, err := dblayer.NewGoodsDB(logger)
	if err != nil {
		return nil, err
	}
	db := new(GoodsDB)
	db.GoodsDB = *cdb
	return db, nil
}

func GetGoodsDB(logger types.Logger, _ *config.ServerConfig) (*GoodsDB, error) {
	cdb, err := dblayer.GetGoodsDB(logger)
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
