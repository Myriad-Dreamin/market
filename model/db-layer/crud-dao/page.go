package crud_dao

import "github.com/jinzhu/gorm"

func PageOption(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Limit(pageSize).Offset((page - 1) * pageSize)
}
