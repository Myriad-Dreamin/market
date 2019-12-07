package dblayer

import (
	"fmt"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/lib/errorc"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToNeeds(needs interface{}, err error) (*Needs, error) {
	if needs == nil {
		return nil, err
	}
	return needs.(*Needs), err
}

func NeedsFactory() interface{} {
	return new(Needs)
}

func NeedssFactory() interface{} {
	return new([]Needs)
}

var (
	needsModel  = new(dorm.Model)
	needsTraits = NewNeedsTraits(Needs{})
)

type Needs struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
	EndAt     time.Time `dorm:"end_at" gorm:"column:end_at;default:CURRENT_TIMESTAMP;not null;"`

	Buyer       uint          `dorm:"buyer" gorm:"column:buyer;not_null"`
	Seller      uint          `dorm:"seller" gorm:"column:seller;not_null"`
	Type        uint16        `dorm:"g_type" gorm:"column:g_type;not_null"`
	Name        string        `dorm:"name" gorm:"column:name;not_null"`
	MinPrice    uint64        `dorm:"min_price" gorm:"column:min_price;not_null"`
	MaxPrice    uint64        `dorm:"min_price" gorm:"column:max_price;not_null"`
	EndDuration time.Duration `dorm:"ddd" gorm:"column:ddd;not_null"`
	Description string        `dorm:"description" gorm:"column:description;not_null"`
	Status      types.GoodsStatus         `dorm:"status" gorm:"column:status;not_null"`

	BuyerFee  uint64 `dorm:"buyer_fee" gorm:"column:buyer_fee;not_null"`
	SellerFee uint64 `dorm:"seller_fee" gorm:"column:seller_fee;not_null"`
}

// TableName specification
func (Needs) TableName() string {
	return "needs"
}

func (Needs) migrate() error {
	return needsTraits.Migrate()
}

func (d Needs) GetID() uint {
	return d.ID
}

func (d *Needs) Create() (int64, error) {
	return needsTraits.Create(d)
}

func (d *Needs) Update() (int64, error) {
	return needsTraits.Update(d)
}

func (d *Needs) UpdateFields(fields []string) (int64, error) {
	return needsTraits.UpdateFields(d, fields)
}

func (d *Needs) Delete() (int64, error) {
	return needsTraits.Delete(d)
}

type NeedsDB struct{}

func NewNeedsDB(logger types.Logger, _ *config.ServerConfig) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func GetNeedsDB(logger types.Logger, _ *config.ServerConfig) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func (needsDB *NeedsDB) ID(id uint) (needs *Needs, err error) {
	return wrapToNeeds(needsTraits.ID(id))
}

type NeedsQuery struct {
	db *gorm.DB
}

func (needsDB *NeedsDB) Filter(f *GoodsFilter) (needss []Needs, err error) {
	err = needsTraits.GoodsFilter(f, &needss)
	return
}

func (needsDB *NeedsDB) FilterI(f *GoodsFilter) (interface{}, error) {
	return needsDB.Filter(f)
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


var needsStatusField = []string{"status", "seller"}

func (needsDB *NeedsDB) Sell(id, uid uint) (int, string) {

	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	needs := new(Needs)
	rdb := tx.First(&needs, id)
	err := rdb.Error
	if rdb.RecordNotFound() {
		needs = nil
		err = nil
	}
	if code, errs := errorc.MaybeSelectError(needs, err); code != types.CodeOK {
		tx.Rollback()
		if tx.Error != nil {
			fmt.Println("rollback error", tx.Error)
		}
		return code, errs
	}
	if needs.Status != types.GoodsStatusUnFinished {
		tx.Rollback()
		if tx.Error != nil {
			fmt.Println("rollback error", tx.Error)
		}
		return types.CodeGoodsStatusNotBeUnfinished, needs.Status.String()
	}
	if needs.EndAt.Before(time.Now()) {
		tx.Rollback()
		if tx.Error != nil {
			fmt.Println("rollback error", tx.Error)
		}
		return types.CodeGoodsLifeTimeout, "expired time"
	}
	needs.Status = types.GoodsStatusPending
	needs.Seller = uid
	err = tx.Model(&needs).Select(needsStatusField).Updates(&needs).Error

	if err != nil {
		errr := tx.Rollback().Error
		if errr != nil {
			fmt.Println("rollback error", errr)
		}
		return types.CodeUpdateError, err.Error()
	}


	err = tx.Commit().Error
	if err != nil {
		errr := tx.Rollback().Error
		if errr != nil {
			fmt.Println("rollback error", errr)
		}
		return types.CodeCommitTransactionError, err.Error()
	} else {
		return types.CodeOK, ""
	}
}
