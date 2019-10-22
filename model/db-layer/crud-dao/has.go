package crud_dao

import "github.com/jinzhu/gorm"

func Has(db *gorm.DB, replacement interface{}) func(id uint) (has bool, err error) {
	return func(id uint) (has bool, err error) {
		rdb := db.First(replacement, id)
		err = rdb.Error
		if err == nil {
			has = true
		} else if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
}
