package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
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

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)

	err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Article, error) {
	var articles []*Article
	query := db.Model(&Article{})

	//如果標題不為空，則在搜索條件內新增查詢標題
	if a.Title != "" {
		query = db.Where("title = ?", a.Title)
	}

	//加入查詢狀態條件與分頁查詢
	query = query.Where("state = ?", a.State).Offset(pageOffset).Limit(pageSize)

	err := query.Where("is_del = ?", 0).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}
