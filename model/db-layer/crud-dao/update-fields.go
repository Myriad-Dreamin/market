package crud_dao

import (
	"fmt"
	"github.com/Myriad-Dreamin/dorm"
)

func (model CRUDModel) UpdateFields(d dorm.ORMObject, fields []string) (int64, error) {
	fmt.Println(model.i.GetDormModel())
	return model.i.GetDormModel().Anchor(d).Select(fields...).UpdateFields()
}
