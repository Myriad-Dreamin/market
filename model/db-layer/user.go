package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	crud_dao "github.com/Myriad-Dreamin/ginx/model/db-layer/crud-dao"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	userModel         *dorm.Model
	userIDFunc           = crud_dao.ID(db)
	userCreateFunc       = crud_dao.Create(db)
	userDeleteFunc       = crud_dao.Delete(db)
	userUpdateFunc       = crud_dao.Update(db)
	userUpdateFieldsFunc = crud_dao.UpdateFields(userModel)
)

type User struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (User) TableName() string {
	return "user"
}

func (User) migrate() error {
	err := db.AutoMigrate(&User{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	userModel, err = dormDB.Model(&User{})
	return err
}

func (d User) GetID() uint {
	return d.ID
}

func (d *User) Create() (int64, error) {
	return userCreateFunc(d)
}

func (d *User) Update() (int64, error) {
	return userUpdateFunc(d)
}

func (d *User) UpdateFields(fields []string) (int64, error) {
	return userUpdateFieldsFunc(d, fields)
}

func (d *User) Delete() (int64, error) {
	return userDeleteFunc(d)
}

type UserDB struct{}

func NewUserDB(logger types.Logger) (*UserDB, error) {
	return new(UserDB), nil
}

func GetUserDB(logger types.Logger) (*UserDB, error) {
	return new(UserDB), nil
}

func (userDB *UserDB) ID(id uint) (user *User, err error) {
	user = new(User)
	err = userIDFunc(user, id)
	return
}

type UserQuery struct {
	db *gorm.DB
}

func (userDB *UserDB) QueryChain() *UserQuery {
	return &UserQuery{db: db}
}

func (userDB *UserQuery) Order(order string, reorder ...bool) *UserQuery {
	userDB.db = userDB.db.Order(order, reorder...)
	return userDB
}

func (userDB *UserQuery) Page(page, pageSize int) *UserQuery {
	userDB.db = userDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return userDB
}
func (userDB *UserQuery) BeforeID(id uint) *UserQuery {
	userDB.db = userDB.db.Where("id <= ?", id)
	return userDB
}

func (userDB *UserQuery) Preload() *UserQuery {
	userDB.db = userDB.db.Set("gorm:auto_preload", true)
	return userDB
}

func (userDB *UserQuery) Query() (users []User, err error) {
	err = userDB.db.Find(&users).Error
	return
}
