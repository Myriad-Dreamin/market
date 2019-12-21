package dblayer

import (
	"fmt"
	"github.com/Myriad-Dreamin/dorm"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/Myriad-Dreamin/minimum-lib/module"
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
	CityCode string    `dorm:"city_code" gorm:"column:city_code;not_null"`

	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	BuyFeeSum       uint64 `dorm:"buy_fee_sum" gorm:"column:buy_fee_sum;not_null"`
	SellFeeSum      uint64 `dorm:"sell_fee_sum" gorm:"column:sell_fee_sum;not_null"`
	BuyFinishCount  uint64 `dorm:"buy_finish_count" gorm:"column:buy_finish_count;not_null"`
	SellFinishCount uint64 `dorm:"sell_finish_count" gorm:"column:sell_finish_count;not_null"`
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

	return db.Model(&st).AddUniqueIndex("mpc", "month", "city_code").Error
}

func (st StatFee) GetID() uint {
	return st.ID
}

func (st *StatFee) Create() (int64, error) {
	return statFeeTraits.Create(st)
}

func (st *StatFee) Create_(tx *gorm.DB) (interface{}, error) {
	return statFeeTraits.Create_(tx, st)
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

func (st *StatFee) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return statFeeTraits.UpdateFields_(db, st, fields)
}

func (st *StatFee) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return statFeeTraits.UpdateFields__(db, st, fields)
}

type StatFeeDB struct{}

func NewStatFeeDB(_ module.Module) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func GetStatFeeDB(_ module.Module) (*StatFeeDB, error) {
	return new(StatFeeDB), nil
}

func StatFeeFilterOption(db *gorm.DB, f *StatFeeRequest) *gorm.DB {
	fmt.Println("page, pagesize", f.Page, f.PageSize)
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
	err = StatFeeFilterOption(db.Table(stateFeeName), f).Select("month, sum(buy_fee_sum)+sum(sell_fee_sum) as sum").Scan(&results).Error
	return
}

func (statFeeDB *StatFeeDB) FilterFeeI(f interface{}) (interface{}, error) {
	return statFeeDB.FilterFee(f.(*StatFeeRequest))
}

func (statFeeDB *StatFeeDB) FilterFeeCount(f *StatFeeRequest) (results []StatFeeCountXYResult, err error) {
	err = StatFeeFilterOption(db.Table(stateFeeName), f).Select("month, sum(buy_finish_count)+sum(sell_finish_count) as count").Scan(&results).Error
	return
}

func (statFeeDB *StatFeeDB) FilterFeeCountI(f interface{}) (interface{}, error) {
	return statFeeDB.FilterFeeCount(f.(*StatFeeRequest))
}

func (statFeeDB *StatFeeDB) ID(id uint) (statFee *StatFee, err error) {
	return wrapToStatFee(statFeeTraits.ID(id))
}

func (statFeeDB *StatFeeDB) First() (statFee *StatFee, err error) {
	statFee = new(StatFee)
	err = db.First(statFee).Error
	return
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

func (statFeeDB *StatFeeQuery) InCityCode(ct string) *StatFeeQuery {
	statFeeDB.db = statFeeDB.db.Where("city_code = ?", ct)
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
	Month time.Time `json:"month" gorm:"column:month"`
	Count int64     `json:"count" gorm:"column:count"`
}

type StatFeeXYResult struct {
	Month time.Time `json:"month" gorm:"column:month"`
	Sum   int64     `json:"sum" gorm:"column:sum"`
}

type StatFeeRequest struct {
	Filter `binding:"dive"`
	LEThan *time.Time `json:"le" form:"le" binding:"required"`
	GEThan *time.Time `json:"ge" form:"ge" binding:"required"`
}
