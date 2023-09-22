package service

import "github.com/jinzhu/gorm"

func (svc *Service) BeginTransaction() *gorm.DB {
	return svc.dao.BeginTransaction()
}
