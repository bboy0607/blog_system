package dao

import "github.com/jinzhu/gorm"

func (d Dao) BeginTransaction() *gorm.DB {
	return d.engine.Begin()
}
