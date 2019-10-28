package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/config"
	crud_dao "github.com/Myriad-Dreamin/market/model/db-layer/crud-dao"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToStatFee(statFee interface{}, err error) (*StatFee, error) {
	return statFee.(*StatFee), err
}

func StatFeeFactory() interface{} {
	return new(StatFee)
}

var (
	statFeeModel         *dorm.Model
	statFeeIDFunc           = crud_dao.ID(StatFeeFactory, db)
	statFeeCreateFunc       = crud_dao.Create(db)
	statFeeDeleteFunc       = crud_dao.Delete(db)
	statFeeUpdateFunc       = crud_dao.Update(db)
	statFeeUpdateFieldsFunc = crud_dao.UpdateFields(statFeeModel)
)

type StatFee struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (StatFee) TableName() string {
	return "stat_fee"
}

func (StatFee) migrate() error {
	err := db.AutoMigrate(&StatFee{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	statFeeModel, err = dormDB.Model(&StatFee{})
	return err
}

func (d StatFee) GetID() uint {
	return d.ID
}

func (d *StatFee) Create() (int64, error) {
	return statFeeCreateFunc(d)
}

func (d *StatFee) Update() (int64, error) {
	return statFeeUpdateFunc(d)
}

func (d *StatFee) UpdateFields(fields []string) (int64, error) {
	return statFeeUpdateFieldsFunc(d, fields)
}

func (d *StatFee) Delete() (int64, error) {
	return statFeeDeleteFunc(d)
}

type StatFeeDB struct{}

func NewStatFeeDB(logger types.Logger, _ *config.ServerConfig) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func GetStatFeeDB(logger types.Logger, _ *config.ServerConfig) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func (statFeeDB *StatFeeDB) ID(id uint) (statFee *StatFee, err error) {
	return wrapToStatFee(statFeeIDFunc(id))
}

type StatFeeQuery struct {
	db *gorm.DB
}

func (statFeeDB *StatFeeDB) QueryChain() *StatFeeQuery {
	return &StatFeeQuery{db: db}
}

func (statFeeDB *StatFeeQuery) Order(order string, reorder ...bool) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Order(order, reorder...)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) Page(page, pageSize int) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return statFeeDB
}
func (statFeeDB *StatFeeQuery) BeforeID(id uint) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("id <= ?", id)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) Preload() *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Set("gorm:auto_preload", true)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) Query() (statFees []StatFee, err error) {
	err = statFeeDB.db.Find(&statFees).Error
	return
}
