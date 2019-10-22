package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	crud_dao "github.com/Myriad-Dreamin/ginx/model/db-layer/crud-dao"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToObject(object interface{}, err error) (*Object, error) {
	return object.(*Object), err
}

func ObjectFactory() interface{} {
	return new(Object)
}

var (
	objectModel         *dorm.Model
	objectIDFunc           = crud_dao.ID(ObjectFactory, db)
	objectCreateFunc       = crud_dao.Create(db)
	objectDeleteFunc       = crud_dao.Delete(db)
	objectUpdateFunc       = crud_dao.Update(db)
	objectUpdateFieldsFunc = crud_dao.UpdateFields(objectModel)
)

type Object struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (Object) TableName() string {
	return "object"
}

func (Object) migrate() error {
	err := db.AutoMigrate(&Object{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	objectModel, err = dormDB.Model(&Object{})
	return err
}

func (d Object) GetID() uint {
	return d.ID
}

func (d *Object) Create() (int64, error) {
	return objectCreateFunc(d)
}

func (d *Object) Update() (int64, error) {
	return objectUpdateFunc(d)
}

func (d *Object) UpdateFields(fields []string) (int64, error) {
	return objectUpdateFieldsFunc(d, fields)
}

func (d *Object) Delete() (int64, error) {
	return objectDeleteFunc(d)
}

type ObjectDB struct{}

func NewObjectDB(logger types.Logger) (*ObjectDB, error) {
	return new(ObjectDB), nil
}

func GetObjectDB(logger types.Logger) (*ObjectDB, error) {
	return new(ObjectDB), nil
}

func (objectDB *ObjectDB) ID(id uint) (object *Object, err error) {
	return wrapToObject(objectIDFunc(id))
}

type ObjectQuery struct {
	db *gorm.DB
}

func (objectDB *ObjectDB) QueryChain() *ObjectQuery {
	return &ObjectQuery{db: db}
}

func (objectDB *ObjectQuery) Order(order string, reorder ...bool) *ObjectQuery {
	objectDB.db = objectDB.db.Order(order, reorder...)
	return objectDB
}

func (objectDB *ObjectQuery) Page(page, pageSize int) *ObjectQuery {
	objectDB.db = objectDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return objectDB
}
func (objectDB *ObjectQuery) BeforeID(id uint) *ObjectQuery {
	objectDB.db = objectDB.db.Where("id <= ?", id)
	return objectDB
}

func (objectDB *ObjectQuery) Preload() *ObjectQuery {
	objectDB.db = objectDB.db.Set("gorm:auto_preload", true)
	return objectDB
}

func (objectDB *ObjectQuery) Query() (objects []Object, err error) {
	err = objectDB.db.Find(&objects).Error
	return
}
