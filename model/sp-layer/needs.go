package splayer

import (
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Needs = dblayer.Needs

//struct {
//	dblayer.Needs
//}

//func (c *Needs) Update() (int64, error) {
//	return c.Needs.Update()
//}

//type _userConverter struct {}
//var NeedsConverter _userConverter
//func (_userConverter) convert(obj dblayer.Needs) Needs { return Needs{Needs: obj} }
//func (_userConverter) reconvert(obj Needs) dblayer.Needs { return obj.Needs }
//func (_userConverter) converts(objs []dblayer.Needs) []Needs { return *(*[]Needs)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []Needs) []dblayer.Needs { return *(*[]dblayer.Needs)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.Needs) *Needs { return (*Needs)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *Needs) *dblayer.Needs { return (*dblayer.Needs)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.Needs) []*Needs { return *(*[]*Needs)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*Needs) []*dblayer.Needs { return *(*[]*dblayer.Needs)(unsafe.Pointer(&objs)) }

type NeedsDB struct {
	dblayer.NeedsDB
}

func NewNeedsDB(m module.Module) (*NeedsDB, error) {
	cdb, err := dblayer.NewNeedsDB(m)
	if err != nil {
		return nil, err
	}
	db := new(NeedsDB)
	db.NeedsDB = *cdb
	return db, nil
}

func GetNeedsDB(m module.Module) (*NeedsDB, error) {
	cdb, err := dblayer.GetNeedsDB(m)
	if err != nil {
		return nil, err
	}
	db := new(NeedsDB)
	db.NeedsDB = *cdb
	return db, nil
}

func (s *Provider) NeedsDB() *NeedsDB {
	return s.needsDB
}
