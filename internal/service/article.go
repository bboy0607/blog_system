package service

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	CoverImageUrl string `form:"cover_image_url"`
	Content       string `form:"content"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,max=100"`
	TagID         uint32 `form:"tag_id" binding:"max=100"`
}

type CountArticleByTitleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CountArticleRequest struct {
	TagID uint32 `form:"tag_id" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// type ListArticleRequest struct {
// 	Title string `form:"title" binding:"max=100"`
// 	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
// }

type ListArticleRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	CoverImageURL string `form:"cover_image_url"`
	Content       string `form:"content"`
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	CoverImageURL string     `json:"cover_image_url"`
	Content       string     `json:"content"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(param.Title, param.Desc, param.CoverImageUrl, param.Content, param.CreatedBy, param.State)
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc Service) CountArticle(param *CountArticleByTitleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc Service) ListAricle(param *ListArticleRequest, pager *app.Pager) ([]*Article, int, error) {
	totalRows, err := svc.dao.CountArticleByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.ListArticleByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.Title,
			Desc:          article.Desc,
			CoverImageURL: article.CoverImageUrl,
			Content:       article.Content,
			State:         article.State,
			Tag:           &model.Tag{Model: &model.Model{ID: article.TagID}, Name: article.TagName},
		})
	}

	return articleList, totalRows, nil
}

func (svc Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.CoverImageURL, param.Content, param.ModifiedBy, param.State)
	if err != nil {
		return err
	}

	err = svc.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc Service) DeleteArticle(param *DeleteArticleRequest) error {
	//刪除文章
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}

	//刪除與文章標籤關聯資料
	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}
