package modeltraits

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/lib/traits"
	"github.com/jinzhu/gorm"
)

type ModelTraits struct {
	traits.BaseTraits
	DormModel *dorm.Model
	DB *gorm.DB
	DormDB *dorm.DB
}

func NewModelTraits(t dorm.ORMObject, db *gorm.DB, dormDB *dorm.DB) ModelTraits {
	tr := ModelTraits{}
	tr.BaseTraits = traits.NewBaseTraits(t)
	tr.DormModel = new(dorm.Model)
	tr.DB = db
	tr.DormDB = dormDB
	return tr
}

func (traits ModelTraits) GetDormModel() *dorm.Model {
	return traits.DormModel
}

func (traits ModelTraits) GetDB() *gorm.DB {
	return traits.DB
}

func (traits ModelTraits) GetDormDB() *dorm.DB {
	return traits.DormDB
}

func (traits *ModelTraits) Migrate() error {
	err := traits.DB.AutoMigrate(traits.ObjectFactory()).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	model, err := traits.DormDB.Model(traits.ObjectFactory().(dorm.ORMObject))
	if err != nil {
		return err
	}
	*traits.DormModel = *model
	return nil
}

