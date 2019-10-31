package crud_dao

import (
	"github.com/Myriad-Dreamin/market/model/traits"
)

type CRUDModel struct {
	i traits.ModelInterface
	replacement interface{}
}


func NewCRUDModel(t traits.ModelInterface) CRUDModel {
	return CRUDModel{
		i: t,
		replacement: t.ObjectFactory(),
	}
}


func (model CRUDModel) ID(id uint) (obj interface{}, err error) {
	obj = model.i.ObjectFactory()
	rdb := model.i.GetDB().First(obj, id)
	err = rdb.Error
	if rdb.RecordNotFound() {
		obj = nil
		err = nil
	}
	return
}

