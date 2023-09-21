package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	ArticleTags := []*ArticleTag{}
	query := db.Model(&ArticleTag{}).Where("tag_id = ? AND is_del = ?", a.TagID, 0)

	err := query.Find(&ArticleTags).Error
	if err != nil {
		return nil, err
	}

	return ArticleTags, nil
}

func (a ArticleTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Updates(values).Error
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&a).Error
}
