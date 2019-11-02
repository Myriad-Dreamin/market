package dblayer

import (
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	general_dao "github.com/Myriad-Dreamin/market/model/db-layer/general-dao"
)

type Traits = traits.ModelTraits
type Interface = traits.Interface
type TraitsAcceptObject = traits.ORMObject

func NewTraits(t TraitsAcceptObject) Traits {
	tt := traits.NewModelTraits(t, db, dormDB)
	return tt
}

type GoodsTraits struct {
	Traits
	general_dao.GoodsModel
}

var (
	_                = "TraitsDefinitionHook"
	NewObjectTraits  = NewTraits
	NewUserTraits    = NewTraits
	NewStatFeeTraits = NewTraits
)

func NewGoodsTraits(t TraitsAcceptObject) (g GoodsTraits) {
	g = GoodsTraits{}
	g.Traits = NewTraits(t)
	g.GoodsModel = general_dao.NewGoodsModel(g.Traits)
	return
}

func NewNeedsTraits(t TraitsAcceptObject) (g GoodsTraits) {
	g = GoodsTraits{}
	g.Traits = NewTraits(t)
	g.GoodsModel = general_dao.NewGoodsModel(g.Traits)
	return
}
