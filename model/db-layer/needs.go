package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/lib/errorc"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
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

	Buyer       uint              `dorm:"buyer" gorm:"column:buyer;not_null"`
	Seller      uint              `dorm:"seller" gorm:"column:seller;not_null"`
	Type        types.GoodsType   `dorm:"g_type" gorm:"column:g_type;not_null"`
	Name        string            `dorm:"name" gorm:"column:name;not_null"`
	CurPrice    uint64            `dorm:"cur_price" gorm:"column:cur_price;not_null"`
	MaxPrice    uint64            `dorm:"min_price" gorm:"column:max_price;not_null"`
	EndDuration time.Duration     `dorm:"ddd" gorm:"column:ddd;not_null"`
	Description string            `dorm:"description" gorm:"column:description;not_null"`
	PicName     string            `dorm:"pic_name" gorm:"column:pic_name;not_null"`
	Status      types.GoodsStatus `dorm:"status" gorm:"column:status;not_null"`

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

func (d *Needs) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return needsTraits.UpdateFields_(db, d, fields)
}

func (d *Needs) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return needsTraits.UpdateFields__(db, d, fields)
}

func (d *Needs) UpdateFields(fields []string) (int64, error) {
	return needsTraits.UpdateFields(d, fields)
}

func (d *Needs) Delete() (int64, error) {
	return needsTraits.Delete(d)
}

type NeedsDB struct{}

func NewNeedsDB(_ module.Module) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func GetNeedsDB(_ module.Module) (*NeedsDB, error) {
	return new(NeedsDB), nil
}

func (needsDB *NeedsDB) ID(id uint) (needs *Needs, err error) {
	return wrapToNeeds(needsTraits.ID(id))
}

func (needsDB *NeedsDB) ID_(db *gorm.DB, id uint) (needs *Needs, err error) {
	return wrapToNeeds(needsTraits.ID_(db, id))
}

type NeedsQuery struct {
	db *gorm.DB
}

func (needsDB *NeedsDB) Filter(f *GoodsFilter) (needss []Needs, err error) {
	err = needsTraits.GoodsFilter(f, &needss)
	return
}

func (needsDB *NeedsDB) FilterI(f interface{}) (interface{}, error) {
	return needsDB.Filter(f.(*GoodsFilter))
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

var needsStatusField = []string{"status", "seller", "buyer_fee", "seller_fee"}

func (needsDB *NeedsDB) Sell(price uint64, id, uid uint) (types.CodeType, string) {

	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	needs, err := needsDB.ID_(tx, id)
	if code, errs := errorc.MaybeSelectError(needs, err); code != types.CodeOK {
		rollback(tx)
		return code, errs
	}
	if needs.Status != types.GoodsStatusUnFinished {
		return types.CodeGoodsStatusNotBeUnfinished, needs.Status.String()
	}
	if needs.EndAt.Before(time.Now()) {
		rollback(tx)
		return types.CodeGoodsLifeTimeout, "expired time"
	}

	if needs.CurPrice <= price {
		return types.CodeGoodsOverflowValue, "need lower price"
	}

	needs.Seller = uid
	needs.Status = types.GoodsStatusPending
	needs.CurPrice = price
	_, err = needs.UpdateFields__(tx.CommonDB(), needsStatusField)

	if err != nil {
		rollback(tx)
		return types.CodeUpdateError, err.Error()
	}

	return commitOrRollback(tx)
}

func (needsDB *NeedsDB) ConfirmSell(id uint, confirm bool, uid uint) (types.CodeType, string) {
	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	needs, err := needsDB.ID_(tx, id)
	if code, errs := errorc.MaybeSelectError(needs, err); code != types.CodeOK {
		rollback(tx)
		return code, errs
	}
	if needs.Status != types.GoodsStatusPending {
		rollback(tx)
		return types.CodeGoodsStatusNotBePending, needs.Status.String()
	}

	if confirm {
		needs.Status = types.GoodsStatusFinished
		needs.BuyerFee = needs.CurPrice / 50
		needs.SellerFee = needs.BuyerFee >> 1

		buyer, err := wrapToUser(userTraits.ID_(tx, needs.Buyer))
		if code, errs := errorc.MaybeSelectError(buyer, err); code != types.CodeOK {
			rollback(tx)
			return code, errs
		}
		seller, err := wrapToUser(userTraits.ID_(tx, needs.Seller))
		if code, errs := errorc.MaybeSelectError(seller, err); code != types.CodeOK {
			rollback(tx)
			return code, errs
		}
		now := time.Now()
		y, m, _ := now.Date()

		x := new(StatFee)
		x.Month = time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
		x.CityCode = buyer.CityCode
		if hs, err := Exists_(tx, x); err != nil {
			rollback(tx)
			return types.CodeSelectError, err.Error()
		} else {
			x.BuyFeeSum += needs.BuyerFee
			x.BuyFinishCount++
			if hs {
				_, err = x.UpdateFields__(tx.CommonDB(), []string{"buy_fee_sum", "buy_finish_count"})
				if err != nil {
					rollback(tx)
					return types.CodeUpdateError, err.Error()
				}
				//fmt.Println("old record")
			} else {
				_, err = x.Create_(tx)
				if err != nil {
					rollback(tx)
					return types.CodeUpdateError, err.Error()
				}
				//fmt.Println("new record")
			}
		}

		x2 := new(StatFee)
		x2.Month = x.Month
		x2.CityCode = seller.CityCode
		if hs, err := Exists_(tx, x2); err != nil {
			rollback(tx)
			return types.CodeSelectError, err.Error()
		} else {
			x2.SellFeeSum += needs.SellerFee
			x2.SellFinishCount++
			if hs {
				_, err = x2.UpdateFields__(tx.CommonDB(), []string{"sell_fee_sum", "sell_finish_count"})
				if err != nil {
					rollback(tx)
					return types.CodeUpdateError, err.Error()
				}
				//fmt.Println("old record")
			} else {
				_, err = x2.Create_(tx)
				if err != nil {
					rollback(tx)
					return types.CodeUpdateError, err.Error()
				}
				//fmt.Println("new record")
			}
		}
	} else {
		needs.Status = types.GoodsStatusCancelled
	}

	_, err = needs.UpdateFields__(tx.CommonDB(), needsStatusField)
	if err != nil {
		rollback(tx)
		return types.CodeUpdateError, err.Error()
	}
	return commitOrRollback(tx)
}
