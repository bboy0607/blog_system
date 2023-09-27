package model

import "github.com/jinzhu/gorm"

type ArticleComment struct {
	*Model
	ArticleID uint32
	Nickname  string
	Comment   string
}

func (a ArticleComment) TableName() string {
	return "blog_article_comment"
}

func (a ArticleComment) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a ArticleComment) List(db *gorm.DB) ([]*ArticleComment, error) {
	articleComments := []*ArticleComment{}
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Find(&articleComments).Error
	if err != nil {
		return nil, err
	}

	return articleComments, nil
}

func (a ArticleComment) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&a).Where("ID = ? AND is_del = ?", a.Model.ID, 0).Updates(values).Error
}

func (a ArticleComment) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&a).Error
}
