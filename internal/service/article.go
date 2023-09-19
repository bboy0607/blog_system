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

func (svc Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.CoverImageUrl, param.Content, param.CreatedBy, param.State)
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
