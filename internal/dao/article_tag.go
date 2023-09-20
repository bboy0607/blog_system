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

func (d Dao) UpdateArticleTag(articleID uint32, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}

	values := map[string]interface{}{
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}

	return articleTag.Update(d.engine, values)
}

func (d Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}

	return articleTag.Delete(d.engine)
}
