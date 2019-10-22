package crud_dao

import "github.com/jinzhu/gorm"

func Where1(creator func() interface{}, template string, db *gorm.DB) func(id interface{}) (interface{}, error) {
	return func(id interface{}) (object interface{}, err error) {
		object = creator()
		rdb := db.Where(template, id).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
	}
}
