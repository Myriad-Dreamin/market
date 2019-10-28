package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/config"
	crud_dao "github.com/Myriad-Dreamin/market/model/db-layer/crud-dao"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToNeeds(needs interface{}, err error) (*Needs, error) {
	return needs.(*Needs), err
}

func NeedsFactory() interface{} {
	return new(Needs)
}

var (
	needsModel         *dorm.Model
	needsIDFunc           = crud_dao.ID(NeedsFactory, db)
	needsCreateFunc       = crud_dao.Create(db)
	needsDeleteFunc       = crud_dao.Delete(db)
	needsUpdateFunc       = crud_dao.Update(db)
	needsUpdateFieldsFunc = crud_dao.UpdateFields(needsModel)
)

type Needs struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null" json:"updated_at"`

	Buyer uint `dorm:"buyer" gorm:"column:buyer;not_null"`
	Seller uint `dorm:"seller" gorm:"column:seller;not_null"`
	Type uint8 `dorm:"g_type" gorm:"column:g_type;not_null"`
	Name string `dorm:"name" gorm:"column:name;not_null"`
	MinPrice uint `dorm:"min_price" gorm:"column:min_price;not_null"`
	MaxPrice uint `dorm:"min_price" gorm:"column:min_price;not_null"`
	IsFixed bool `dorm:"is_fixed" gorm:"column:is_fixed;not_null"`
	EndDuration time.Duration `dorm:"ddd" gorm:"column:ddd;not_null"`
	Description string `dorm:"description" gorm:"column:description;not_null"`
	Status uint8 `dorm:"status" gorm:"column:status;not_null"`

	BuyerFee uint64 `dorm:"buyer_fee" gorm:"column:buyer_fee;not_null"`
	SellerFee uint64 `dorm:"seller_fee" gorm:"column:seller_fee;not_null"`
}

// TableName specification
func (Needs) TableName() string {
	return "needs"
}

func (Needs) migrate() error {
	err := db.AutoMigrate(&Needs{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	needsModel, err = dormDB.Model(&Needs{})
	return err
}

func (d Needs) GetID() uint {
	return d.ID
}

func (d *Needs) Create() (int64, error) {
	return needsCreateFunc(d)
}

func (d *Needs) Update() (int64, error) {
	return needsUpdateFunc(d)
}

func (d *Needs) UpdateFields(fields []string) (int64, error) {
	return needsUpdateFieldsFunc(d, fields)
}

func (d *Needs) Delete() (int64, error) {
	return needsDeleteFunc(d)
}

type NeedsDB struct{}

func NewNeedsDB(logger types.Logger, _ *config.ServerConfig) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func GetNeedsDB(logger types.Logger, _ *config.ServerConfig) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func (needsDB *NeedsDB) ID(id uint) (needs *Needs, err error) {
	return wrapToNeeds(needsIDFunc(id))
}

type NeedsQuery struct {
	db *gorm.DB
}

func (needsDB *NeedsDB) QueryChain() *NeedsQuery {
	return &NeedsQuery{db: db}
}

func (needsDB *NeedsQuery) Order(order string, reorder ...bool) *NeedsQuery {
	needsDB.db = needsDB.db.Order(order, reorder...)
	return needsDB
}

func (needsDB *NeedsQuery) Page(page, pageSize int) *NeedsQuery {
	needsDB.db = needsDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return needsDB
}
func (needsDB *NeedsQuery) BeforeID(id uint) *NeedsQuery {
	needsDB.db = needsDB.db.Where("id <= ?", id)
	return needsDB
}

func (needsDB *NeedsQuery) Preload() *NeedsQuery {
	needsDB.db = needsDB.db.Set("gorm:auto_preload", true)
	return needsDB
}

func (needsDB *NeedsQuery) Query() (needss []Needs, err error) {
	err = needsDB.db.Find(&needss).Error
	return
}
