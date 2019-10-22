package crud_dao

import "github.com/jinzhu/gorm"

func ID(creator func() interface{}, db *gorm.DB) func(id uint) (obj interface{}, err error) {
	return func(id uint) (obj interface{}, err error) {
		obj = creator()
		rdb := db.First(obj, id)
		err = rdb.Error
		if rdb.RecordNotFound() {
			obj = nil
			err = nil
		}
		return
	}
}

