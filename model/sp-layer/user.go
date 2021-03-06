package splayer

import (
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type User = dblayer.User

//struct {
//	dblayer.User
//}

//func (c *User) Update() (int64, error) {
//	return c.User.Update()
//}

//type _userConverter struct {}
//var UserConverter _userConverter
//func (_userConverter) convert(obj dblayer.User) User { return User{User: obj} }
//func (_userConverter) reconvert(obj User) dblayer.User { return obj.User }
//func (_userConverter) converts(objs []dblayer.User) []User { return *(*[]User)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []User) []dblayer.User { return *(*[]dblayer.User)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.User) *User { return (*User)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *User) *dblayer.User { return (*dblayer.User)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.User) []*User { return *(*[]*User)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*User) []*dblayer.User { return *(*[]*dblayer.User)(unsafe.Pointer(&objs)) }

type UserDB struct {
	dblayer.UserDB
}

func NewUserDB(m module.Module) (*UserDB, error) {
	cdb, err := dblayer.NewUserDB(m)
	if err != nil {
		return nil, err
	}
	db := new(UserDB)
	db.UserDB = *cdb
	return db, nil
}

func GetUserDB(m module.Module) (*UserDB, error) {
	cdb, err := dblayer.GetUserDB(m)
	if err != nil {
		return nil, err
	}
	db := new(UserDB)
	db.UserDB = *cdb
	return db, nil
}

func (s *Provider) UserDB() *UserDB {
	return s.userDB
}
