package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/lib/errorc"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToGoods(goods interface{}, err error) (*Goods, error) {
	if goods == nil {
		return nil, err
	}
	return goods.(*Goods), err
}

var (
	goodsTraits = NewGoodsTraits(Goods{})
)

type Goods struct {
	ID          uint              `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt   time.Time         `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   time.Time         `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;"`
	EndAt       time.Time         `dorm:"end_at" gorm:"column:end_at;default:CURRENT_TIMESTAMP;not null;"`
	Seller      uint              `dorm:"seller" gorm:"column:seller;not_null"`
	Buyer       uint              `dorm:"buyer" gorm:"column:buyer;not_null"`
	Type        types.GoodsType   `dorm:"g_type" gorm:"column:g_type;not_null"`
	Name        string            `dorm:"name" gorm:"column:name;not_null"`
	CurPrice    uint64            `dorm:"cur_price" gorm:"column:cur_price;not_null"`
	MinPrice    uint64            `dorm:"min_price" gorm:"column:min_price;not_null"`
	IsFixed     bool              `dorm:"is_fixed" gorm:"column:is_fixed;not_null"`
	Description string            `dorm:"description" gorm:"column:description;not_null"`
	PicName     string            `dorm:"pic_name" gorm:"column:pic_name;not_null"`
	Status      types.GoodsStatus `dorm:"status" gorm:"column:status;not_null"`
	BuyerFee    uint64            `dorm:"buyer_fee" gorm:"column:buyer_fee;not_null"`
	SellerFee   uint64            `dorm:"seller_fee" gorm:"column:seller_fee;not_null"`
}

// TableName specification
func (Goods) TableName() string {
	return "goods"
}

func (Goods) migrate() error {
	return goodsTraits.Migrate()
}

func (d Goods) GetID() uint {
	return d.ID
}

func (d *Goods) Create() (int64, error) {
	return goodsTraits.Create(d)
}

func (d *Goods) Update() (int64, error) {
	return goodsTraits.Update(d)
}

func (d *Goods) UpdateFields(fields []string) (int64, error) {
	return goodsTraits.UpdateFields(d, fields)
}

func (d *Goods) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return goodsTraits.UpdateFields_(db, d, fields)
}

func (d *Goods) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return goodsTraits.UpdateFields__(db, d, fields)
}

func (d *Goods) Delete() (int64, error) {
	return goodsTraits.Delete(d)
}

type GoodsDB struct{}

func NewGoodsDB(_ module.Module) (*GoodsDB, error) {
	return new(GoodsDB), nil
}

func GetGoodsDB(_ module.Module) (*GoodsDB, error) {
	return new(GoodsDB), nil
}

func (goodsDB *GoodsDB) ID(id uint) (goods *Goods, err error) {
	return wrapToGoods(goodsTraits.ID(id))
}

func (goodsDB *GoodsDB) ID_(db *gorm.DB, id uint) (goods *Goods, err error) {
	return wrapToGoods(goodsTraits.ID_(db, id))
}

type GoodsQuery struct {
	db *gorm.DB
}

func (goodsDB *GoodsDB) Filter(f *GoodsFilter) (goodss []Goods, err error) {
	err = goodsTraits.GoodsFilter(f, &goodss)
	return
}

func (goodsDB *GoodsDB) FilterI(f interface{}) (interface{}, error) {
	return goodsDB.Filter(f.(*GoodsFilter))
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

var goodsStatusBuyFixedField = []string{"status", "buyer"}
var goodsStatusBuyField = []string{"status", "buyer", "cur_price"}
var goodsStatusConfirmBuyField = []string{"status", "buyer", "buyer_fee", "seller_fee"}

func (goodsDB *GoodsDB) BuyFixed(id, uid uint) (types.CodeType, string) {
	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	goods, err := goodsDB.ID_(tx, id)
	if code, errs := errorc.MaybeSelectError(goods, err); code != types.CodeOK {
		rollback(tx)
		return code, errs
	}
	if goods.Status != types.GoodsStatusUnFinished {
		rollback(tx)
		return types.CodeGoodsStatusNotBeUnfinished, goods.Status.String()
	}
	if goods.EndAt.Before(time.Now()) {
		rollback(tx)
		return types.CodeGoodsLifeTimeout, "expired time"
	}
	if goods.IsFixed == false {
		return types.CodeGoodsBuyTypeInvalid, "unfixed goods form"
	}

	goods.Status = types.GoodsStatusPending
	goods.Buyer = uid

	_, err = goods.UpdateFields__(tx.CommonDB(), goodsStatusBuyFixedField)
	if err != nil {
		rollback(tx)
		return types.CodeUpdateError, err.Error()
	}
	return commitOrRollback(tx)
}

func (goodsDB *GoodsDB) Buy(id, uid uint, price uint64) (types.CodeType, string) {
	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	goods, err := goodsDB.ID_(tx, id)
	if code, errs := errorc.MaybeSelectError(goods, err); code != types.CodeOK {
		rollback(tx)
		return code, errs
	}
	if goods.Status != types.GoodsStatusUnFinished && goods.Status != types.GoodsStatusPending {
		rollback(tx)
		return types.CodeGoodsStatusNotBeUnfinishedOrPending, goods.Status.String()
	}
	if goods.EndAt.Before(time.Now()) {
		rollback(tx)
		return types.CodeGoodsLifeTimeout, "expired time"
	}
	if goods.IsFixed {
		return types.CodeGoodsBuyTypeInvalid, "fixed goods form"
	}
	if goods.CurPrice >= price {
		return types.CodeGoodsInsufficientValue, "need higher price"
	}
	goods.Buyer = uid
	goods.Status = types.GoodsStatusPending
	goods.CurPrice = price

	_, err = goods.UpdateFields__(tx.CommonDB(), goodsStatusBuyField)
	if err != nil {
		rollback(tx)
		return types.CodeUpdateError, err.Error()
	}
	return commitOrRollback(tx)
}


var statFeeBuyFields = []string{"buy_fee_sum", "buy_finish_count"}
func (goodsDB *GoodsDB) ConfirmBuy(id uint, confirm bool, uid uint) (types.CodeType, string) {
	tx := db.Begin()
	if tx.Error != nil {
		return types.CodeBeginTransactionError, tx.Error.Error()
	}
	goods, err := goodsDB.ID_(tx, id)
	if code, errs := errorc.MaybeSelectError(goods, err); code != types.CodeOK {
		rollback(tx)
		return code, errs
	}
	if goods.Status != types.GoodsStatusPending {
		rollback(tx)
		return types.CodeGoodsStatusNotBePending, goods.Status.String()
	}
	if confirm {
		goods.Status = types.GoodsStatusFinished
		goods.BuyerFee = uint64(float64(goods.CurPrice) * config.GoodsBuyerRaito + 0.5)
		goods.SellerFee = uint64(float64(goods.CurPrice) * config.GoodsSellerRaito + 0.5)

		buyer, err := wrapToUser(userTraits.ID_(tx, goods.Buyer))
		if code, errs := errorc.MaybeSelectError(buyer, err); code != types.CodeOK {
			rollback(tx)
			return code, errs
		}
		seller, err := wrapToUser(userTraits.ID_(tx, goods.Seller))
		if code, errs := errorc.MaybeSelectError(seller, err); code != types.CodeOK {
			rollback(tx)
			return code, errs
		}
		now := time.Now()
		y, m, _ := now.Date()

		x := new(StatFee)
		x.Month = time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
		x.CityCode = buyer.CityCode
		x.GoodsType = goods.Type
		if hs, err := Exists_(tx, x); err != nil {
			rollback(tx)
			return types.CodeSelectError, err.Error()
		} else {
			x.BuyFeeSum += goods.BuyerFee
			x.BuyFinishCount++
			if hs {
				_, err = x.UpdateFields__(tx.CommonDB(), statFeeBuyFields)
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
		x2.GoodsType = goods.Type
		if hs, err := Exists_(tx, x2); err != nil {
			rollback(tx)
			return types.CodeSelectError, err.Error()
		} else {
			x2.SellFeeSum += goods.SellerFee
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
		goods.Status = types.GoodsStatusCancelled
	}

	_, err = goods.UpdateFields__(tx.CommonDB(), goodsStatusConfirmBuyField)
	if err != nil {
		rollback(tx)
		return types.CodeUpdateError, err.Error()
	}
	return commitOrRollback(tx)
}

func Exists_(tx *gorm.DB, v interface{}) (has bool, err error) {
	rdb := tx.Where(v).First(v)
	err = rdb.Error
	if err == nil {
		has = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
