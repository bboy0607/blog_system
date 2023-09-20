package dao

import "membership_system/internal/model"

//創建文章與標籤的關聯
func (d Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
		TagID:     tagID,
		Model:     &model.Model{CreatedBy: createdBy},
	}

	return articleTag.Create(d.engine)
}
