package traits

import (
	"fmt"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

type ModelTraits struct {
	BaseTraits
	DormModel *dorm.Model
	DB *gorm.DB
	DormDB *dorm.DB
}

func NewModelTraits(t dorm.ORMObject, db *gorm.DB, dormDB *dorm.DB) ModelTraits {
	traits := ModelTraits{}
	traits.BaseTraits = NewBaseTraits(t)
	traits.DormModel = new(dorm.Model)
	traits.DB = db
	traits.DormDB = dormDB
	return traits
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
	fmt.Println(traits.DormModel)
	*traits.DormModel = *model
	fmt.Println(traits.DormModel)
	return nil
}

