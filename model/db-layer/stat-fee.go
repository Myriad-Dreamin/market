package dblayer

import (
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToStatFee(statFee interface{}, err error) (*StatFee, error) {
	if statFee == nil {
		return nil, err
	}
	return statFee.(*StatFee), err
}

func StatFeeFactory() interface{} {
	return new(StatFee)
}

var (
	statFeeTraits = NewStatFeeTraits(StatFee{})
)

type StatFee struct {
	ID uint `dorm:"id" gorm:"column:id;primary_key;not_null"`

	Month    time.Time `dorm:"month" gorm:"column:month;not null" json:"month"`
	Province string    `dorm:"province" gorm:"column:province;not_null"`
	City     string    `dorm:"city" gorm:"column:city;not_null"`

	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	BuyFeeSum   int64 `dorm:"buy_fee_sum" gorm:"column:buy_fee_sum;not_null"`
	SellFeeSum  int64 `dorm:"sell_fee_sum" gorm:"column:sell_fee_sum;not_null"`
	FinishCount int64 `dorm:"finish_count" gorm:"column:finish_count;not_null"`
}

const stateFeeName = "stat_fee"

// TableName specification
func (StatFee) TableName() string {
	return stateFeeName
}

func (st StatFee) migrate() error {
	if err := statFeeTraits.Migrate(); err != nil {
		return err
	}

	return db.Model(&st).AddUniqueIndex("mpc", "month", "province", "city").Error
}

func (st StatFee) GetID() uint {
	return st.ID
}

func (st *StatFee) Create() (int64, error) {
	return statFeeTraits.Create(st)
}

func (st *StatFee) Update() (int64, error) {
	return statFeeTraits.Update(st)
}

func (st *StatFee) UpdateFields(fields []string) (int64, error) {
	return statFeeTraits.UpdateFields(st, fields)
}

func (st *StatFee) Delete() (int64, error) {
	return statFeeTraits.Delete(st)
}

type StatFeeDB struct{}

func NewStatFeeDB(logger types.Logger, _ *config.ServerConfig) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func GetStatFeeDB(logger types.Logger, _ *config.ServerConfig) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func StatFeeFilterOption(db *gorm.DB, f *StatFeeRequest) *gorm.DB {
	db = gorm_crud_dao.FilterOption(db, &f.Filter)

	if f.LEThan != nil {
		db = db.Where("month <= ?", f.LEThan)
	}
	if f.GEThan != nil {
		db = db.Where("month >= ?", f.GEThan)
	}
	return db
}

func (statFeeDB *StatFeeDB) FilterFee(f *StatFeeRequest) (results []StatFeeXYResult, err error) {
	err = StatFeeFilterOption(db.Table(stateFeeName), f).Select("month, sum(buy_fee_sum)").Scan(&results).Error
	return
}

func (statFeeDB *StatFeeDB) FilterFeeI(f interface{}) (interface{}, error) {
	return statFeeDB.FilterFee(f.(*StatFeeRequest))
}

func (statFeeDB *StatFeeDB) FilterFeeCount(f *StatFeeRequest) (results []StatFeeCountXYResult, err error) {
	err = StatFeeFilterOption(db.Table(stateFeeName), f).Select("month, sum(finish_count)").Scan(&results).Error
	return
}

func (statFeeDB *StatFeeDB) FilterFeeCountI(f interface{}) (interface{}, error) {
	return statFeeDB.FilterFeeCount(f.(*StatFeeRequest))
}

func (statFeeDB *StatFeeDB) ID(id uint) (statFee *StatFee, err error) {
	return wrapToStatFee(statFeeTraits.ID(id))
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

func (statFeeDB *StatFeeQuery) LETime(t time.Time) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("month <= ?", t)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) GETime(t time.Time) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("month >= ?", t)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) InProvince(pv string) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("province = ?", pv)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) InCity(ct string) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("city = ?", ct)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) GroupBy(s string) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Group(s)
	return statFeeDB
}

func (statFeeDB *StatFeeQuery) Query() (statFees []StatFee, err error) {
	err = statFeeDB.db.Find(&statFees).Error
	return
}

func (statFeeDB *StatFeeQuery) Scan(desc interface{}) (err error) {
	err = statFeeDB.db.Scan(desc).Error
	return
}

type StatFeeCountXYResult struct {
	Month time.Time `json:"month"`
	Count int64     `json:"count"`
}

type StatFeeXYResult struct {
	Month time.Time `json:"month"`
	Sum   int64     `json:"count"`
}

type StatFeeRequest struct {
	Filter
	LEThan *time.Time `json:"le" form:"le" binding:"required"`
	GEThan *time.Time `json:"ge" form:"ge" binding:"required"`
}
