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

func (a ArticleTag) Create(db *gorm.DB, tagIDs []uint32) error {
	//建立新交易
	tx := db.Begin()

	for _, tagID := range tagIDs {
		articleTag := &ArticleTag{
			ArticleID: a.ArticleID,
			TagID:     tagID,
		}

		err := tx.Create(&articleTag).Error
		// 如果在交易中有任何錯誤，回滾交易
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//如果所有操作都成功，提交交易
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
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
