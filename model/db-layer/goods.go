package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/config"
	crud_dao "github.com/Myriad-Dreamin/market/model/db-layer/crud-dao"
	general_dao "github.com/Myriad-Dreamin/market/model/db-layer/general-dao"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToGoods(goods interface{}, err error) (*Goods, error) {
	return goods.(*Goods), err
}

func GoodsFactory() interface{} {
	return new(Goods)
}

func GoodssFactory() interface{} {
	return new([]Goods)
}

var (
	goodsModel         *dorm.Model
	goodsIDFunc           = crud_dao.ID(GoodsFactory, db)
	goodsCreateFunc       = crud_dao.Create(db)
	goodsDeleteFunc       = crud_dao.Delete(db)
	goodsUpdateFunc       = crud_dao.Update(db)
	goodsUpdateFieldsFunc = crud_dao.UpdateFields(goodsModel)
	goodsFilter = general_dao.GoodsFilterFunc(GoodssFactory, db)
)

type Goods struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;"`
	EndAt time.Time `dorm:"end_at" gorm:"column:end_at;default:CURRENT_TIMESTAMP;not null;"`

	Seller uint `dorm:"seller" gorm:"column:seller;not_null"`
	Buyer uint `dorm:"buyer" gorm:"column:buyer;not_null"`
	Type uint8 `dorm:"g_type" gorm:"column:g_type;not_null"`
	Name string `dorm:"name" gorm:"column:name;not_null"`
	MinPrice uint64 `dorm:"min_price" gorm:"column:min_price;not_null"`
	IsFixed bool `dorm:"is_fixed" gorm:"column:is_fixed;not_null"`
	Description string `dorm:"description" gorm:"column:description;not_null"`
	Status uint8 `dorm:"status" gorm:"column:status;not_null"`

	BuyerFee uint64 `dorm:"buyer_fee" gorm:"column:buyer_fee;not_null"`
	SellerFee uint64 `dorm:"seller_fee" gorm:"column:seller_fee;not_null"`
}

// TableName specification
func (Goods) TableName() string {
	return "goods"
}

func (Goods) migrate() error {
	err := db.AutoMigrate(&Goods{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	goodsModel, err = dormDB.Model(&Goods{})
	return err
}

func (d Goods) GetID() uint {
	return d.ID
}

func (d *Goods) Create() (int64, error) {
	return goodsCreateFunc(d)
}

func (d *Goods) Update() (int64, error) {
	return goodsUpdateFunc(d)
}

func (d *Goods) UpdateFields(fields []string) (int64, error) {
	return goodsUpdateFieldsFunc(d, fields)
}

func (d *Goods) Delete() (int64, error) {
	return goodsDeleteFunc(d)
}

type GoodsDB struct{}

func NewGoodsDB(logger types.Logger, _ *config.ServerConfig) (*GoodsDB, error) {
	return new(GoodsDB), nil
}

func GetGoodsDB(logger types.Logger, _ *config.ServerConfig) (*GoodsDB, error) {
	return new(GoodsDB), nil
}

func (goodsDB *GoodsDB) ID(id uint) (goods *Goods, err error) {
	return wrapToGoods(goodsIDFunc(id))
}

type GoodsQuery struct {
	db *gorm.DB
}

func (goodsDB *GoodsDB) Filter(f *GoodsFilter) (goodss []Goods, err error) {
	i, err := goodsFilter(f)
	if err != nil {
		return nil, err
	}
	goodss = *i.(*[]Goods)
	return
}

func (goodsDB *GoodsDB) FilterI(f *GoodsFilter) (interface{}, error) {
	return goodsDB.Filter(f)
}

func (goodsDB *GoodsDB) QueryChain() *GoodsQuery {
	return &GoodsQuery{db: db}
}

func (goodsDB *GoodsQuery) Order(order string, reorder ...bool) *GoodsQuery {
	goodsDB.db = goodsDB.db.Order(order, reorder...)
	return goodsDB
}

func (goodsDB *GoodsQuery) Page(page, pageSize int) *GoodsQuery {
	goodsDB.db = goodsDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return goodsDB
}
func (goodsDB *GoodsQuery) BeforeID(id uint) *GoodsQuery {
	goodsDB.db = goodsDB.db.Where("id <= ?", id)
	return goodsDB
}

func (goodsDB *GoodsQuery) Preload() *GoodsQuery {
	goodsDB.db = goodsDB.db.Set("gorm:auto_preload", true)
	return goodsDB
}

func (goodsDB *GoodsQuery) Query() (goodss []Goods, err error) {
	err = goodsDB.db.Find(&goodss).Error
	return
}
