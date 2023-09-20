package dao

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

func (d Dao) CreateArticle(title, desc, coverImageURL, content, createdBy string, state uint8) (*model.Article, error) {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: coverImageURL,
		Content:       content,
		State:         state,
		Model:         &model.Model{CreatedBy: createdBy},
	}

	return article.Create(d.engine)
}

func (d Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{
		Title: title,
		State: state,
	}

	return article.Count(d.engine)
}

func (d Dao) ListArticle(title string, state uint8, page int, pageSize int) ([]*model.Article, error) {

	pageOffset := app.GetPageOffset(page, pageSize)

	article := model.Article{
		Title: title,
		State: state,
	}

	return article.List(d.engine, pageOffset, pageSize)
}

func (d Dao) GetArticle(id uint32) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
	}

	return article.Get(d.engine)
}

func (d Dao) UpdateArticle(id uint32, title, desc, coverImageURL, content, modifiedBy string, state uint8) error {
	article := model.Article{Model: &model.Model{ID: id}}
	values := map[string]interface{}{
		"modifiedBy": modifiedBy,
		"state":      state,
	}
	if title != "" {
		values["title"] = title
	}
	if coverImageURL != "" {
		values["cover_image_url"] = coverImageURL
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}

	return article.Update(d.engine, values)
}

func (d Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}

	return article.Delete(d.engine)
}
