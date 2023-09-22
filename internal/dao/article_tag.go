package dao

import (
	"membership_system/internal/model"

	"github.com/jinzhu/gorm"
)

// 創建文章與標籤的關聯
func (d Dao) CreateArticleTag(articleID uint32, tagIDs []uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
		Model:     &model.Model{CreatedBy: createdBy},
	}

	return articleTag.Create(d.engine, tagIDs)
}

// 創建文章與標籤的關聯(使用交易)
func (d Dao) CreateArticleTagInTransaction(tx *gorm.DB, articleID uint32, tagIDs []uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
		Model:     &model.Model{CreatedBy: createdBy},
	}

	return articleTag.CreateInTransaction(tx, tagIDs)
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
