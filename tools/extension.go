package tools

import "gorm.io/gorm"

func Paging(pageIndex, pageSize int) func(db *gorm.DB) *gorm.DB {
	pageIndex = pageIndex - 1
	if pageIndex < 0 {
		pageIndex = 0
	}
	return func(db *gorm.DB) *gorm.DB {
		if pageSize != 0 {
			return db.Limit(pageSize).Offset(pageIndex * pageSize)
		}
		return db
	}
}
