package dao

import (
	"frame/internal/request"
	"github.com/jinzhu/gorm"
)

func PaginateInfo(r *request.PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page == 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 100:
			r.PageSize = 100
		case r.PageSize <= 0:
			r.PageSize = 10
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}
