package service

import "membership_system/internal/model"

type CreateArticleCommentRequest struct {
	ArticleID uint32 `form:"article_id" binding:"required,gte=1"`
	Nickname  string `form:"nickname" binding:"required,min=3,max=100"`
	Comment   string `form:"comment" binding:"required,min=3,max=300"`
	CreatedBy string `form:"created_by" binding:"required"`
}

type ListArticleCommentRequest struct {
	ArticleID uint32 `form:"article_id" binding:"required,gte=1"`
}

type UpdateArticleCommentRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	NickName   string `form:"nickname" binding:"required,min=3,max=100"`
	Comment    string `form:"comment" binding:"required,min=3,max=300"`
	ModifiedBy string `form:"modified_by" binding:"required"`
}

type DeleteArticleCommentRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc Service) CreateArticleComment(param *CreateArticleCommentRequest) error {
	return svc.dao.CreateArticleComment(param.ArticleID, param.Nickname, param.Comment, param.CreatedBy)
}

func (svc Service) ListArticleComment(param *ListArticleCommentRequest) ([]*model.ArticleComment, error) {
	return svc.dao.ListArticleComment(param.ArticleID)
}

func (svc Service) UpdateArticleComment(param *UpdateArticleCommentRequest) error {
	return svc.dao.UpdateArticleComment(param.ID, param.NickName, param.Comment, param.ModifiedBy)
}

func (svc Service) DeleteArticleComment(param *DeleteArticleCommentRequest) error {
	return svc.dao.DeleteArticleComment(param.ID)
}
