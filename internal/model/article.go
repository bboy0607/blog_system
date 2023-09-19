package model

import (
	"membership_system/internal/model"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*model.Model
	Title         string
	Desc          string
	CoverImageUrl string
	Content       string
	State         uint8
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}
