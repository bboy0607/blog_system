package dao

import "membership_system/internal/model"

func (d Dao) CreateArticleComment(articleID uint32, nickname string, comment string, createdBy string) error {
	articleComment := model.ArticleComment{
		ArticleID: articleID,
		Nickname:  nickname,
		Comment:   comment,
		Model:     &model.Model{CreatedBy: createdBy},
	}

	return articleComment.Create(d.engine)
}

func (d Dao) UpdateArticleComment(id uint32, nickname string, comment string, modifiedBy string) error {
	articleComment := model.ArticleComment{
		Model: &model.Model{ID: id},
	}

	values := map[string]interface{}{
		"nickname":    nickname,
		"comment":     comment,
		"modified_by": modifiedBy,
	}

	return articleComment.Update(d.engine, values)
}

func (d Dao) DeleteArticleComment(id uint32) error {
	articleComment := model.ArticleComment{Model: &model.Model{ID: id}}
	return articleComment.Delete(d.engine)
}
