package traits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

type BaseTraitsInterface interface {
	ObjectFactory() interface{}
	SliceFactory() interface{}
}

type GORMTraitsInterface interface {
	GetDB() *gorm.DB
}

type DORMTraitsInterface interface {
	GetDormModel() *dorm.Model
	GetDormDB() *dorm.DB
}

type ModelInterface interface {
	BaseTraitsInterface
	GORMTraitsInterface
	DORMTraitsInterface
}

type DatabaseTraitsInterface interface {
	Migrate() error
}

type Interface interface {
	ModelInterface
	DatabaseTraitsInterface
}

type Traits = ModelTraits

func NewTraits(t dorm.ORMObject, db *gorm.DB, dormDB *dorm.DB) Interface {
	model := NewModelTraits(t, db, dormDB)
	return &model
}
