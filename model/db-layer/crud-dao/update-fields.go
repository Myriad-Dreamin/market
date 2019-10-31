package crud_dao

import (
	"github.com/Myriad-Dreamin/dorm"
)

func (model CRUDModel) UpdateFields(d dorm.ORMObject, fields []string) (int64, error) {
	return model.i.GetDormModel().Anchor(d).Select(fields...).UpdateFields()
}
