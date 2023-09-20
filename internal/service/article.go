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

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ListArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
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
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
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

func (svc Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc Service) ListAricle(param *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.ListArticle(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.CoverImageURL, param.Content, param.ModifiedBy, param.State)
}

func (svc Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
