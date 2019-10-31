package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	crud_dao "github.com/Myriad-Dreamin/market/model/db-layer/crud-dao"
	general_dao "github.com/Myriad-Dreamin/market/model/db-layer/general-dao"
	"github.com/Myriad-Dreamin/market/model/modeltraits"
)

type Traits struct {
	modeltraits.Interface
	crud_dao.CRUDModel
}
type TraitsAcceptObject = dorm.ORMObject

func NewTraits(t TraitsAcceptObject) Traits {
	tt := Traits{}
	tt.Interface = modeltraits.NewTraits(t, db, dormDB)
	tt.CRUDModel = crud_dao.NewCRUDModel(tt)
	return tt
}

type GoodsTraits struct {
	Traits
	general_dao.GoodsModel
}

var (
	_ = "TraitsDefinitionHook"
	NewObjectTraits = NewTraits
	NewUserTraits = NewTraits
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
