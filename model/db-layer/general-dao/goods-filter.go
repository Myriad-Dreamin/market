package general_dao

import (
	modeltraits "github.com/Myriad-Dreamin/go-model-traits"
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/jinzhu/gorm"
	"time"
)


type GoodsFilter struct {
	traits.Filter
	BySeller   uint       `json:"seller" form:"seller"`
	ByBuyer    uint       `json:"buyer" form:"buyer"`
	HasType    uint8      `json:"type" form:"type"`
	HasStatus  uint8      `json:"status" form:"status"`
	WithName   string     `json:"name" form:"name"`
	MinPriceL  *uint      `json:"min_price_l" form:"min_price_l"`
	MinPriceR  *uint      `json:"min_price_r" form:"min_price_r"`
	MaxPriceL  *uint      `json:"max_price_l" form:"max_price_l"`
	MaxPriceR  *uint      `json:"max_price_r" form:"max_price_r"`
	FixedTag   *bool      `json:"fixed" form:"fixed"`
	EndBefore  *time.Time `json:"end_before" form:"end_before"`
	BeginAfter *time.Time `json:"begin_after" form:"begin_after"`
}

func GoodsFilterOption(db *gorm.DB, f *GoodsFilter) *gorm.DB {
	db = gorm_crud_dao.FilterOption(db, &f.Filter)
	if f.BySeller != 0 {
		db = db.Where("seller = ?", f.BySeller)
	}
	if f.ByBuyer != 0 {
		db = db.Where("buyer = ?", f.ByBuyer)
	}
	if f.HasType != 0 {
		db = db.Where("type = ?", f.HasType)
	}
	if f.HasStatus != 0 {
		db = db.Where("status = ?", f.HasStatus)
	}
	if len(f.WithName) != 0 {
		db = db.Where("name = ?", f.WithName)
	}
	if f.MinPriceL != nil && f.MinPriceR != nil {
		db = db.Where("min_price between ? and ?", *f.MinPriceL, *f.MinPriceR)
	} else if f.MinPriceL != nil {
		db = db.Where("min_price > ?", *f.MinPriceL)
	} else if f.MinPriceR != nil {
		db = db.Where("min_price < ?", *f.MinPriceR)
	}

	if f.MaxPriceL != nil && f.MaxPriceR != nil {
		db = db.Where("max_price between ? and ?", *f.MaxPriceL, *f.MaxPriceR)
	} else if f.MaxPriceL != nil {
		db = db.Where("max_price > ?", *f.MaxPriceL)
	} else if f.MaxPriceR != nil {
		db = db.Where("max_price < ?", *f.MaxPriceR)
	}
	if f.FixedTag != nil {
		db = db.Where("is_fixed = ?", *f.FixedTag)
	}
	if f.EndBefore != nil {
		db = db.Where("end_at < ?", *f.EndBefore)
	}
	if f.BeginAfter != nil {
		db = db.Where("created_at > ?", *f.BeginAfter)
	}
	return db
}

type GoodsModelOperatingModel interface {
	modeltraits.BaseTraitsInterface
	modeltraits.GORMTraitsInterface
}

type GoodsModel struct {
	i GoodsModelOperatingModel
}

func NewGoodsModel(modelInterface GoodsModelOperatingModel) GoodsModel {
	return GoodsModel{i: modelInterface}
}

func (model GoodsModel) GoodsFilter(f *GoodsFilter, goodss interface{}) (err error) {
	err = GoodsFilterOption(model.i.GetGormDB(), f).Find(goodss).Error
	return
}
