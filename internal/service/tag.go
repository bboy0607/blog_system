package service

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	State     uint8  `form:"state, default=0" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
}

type CountTagRequest struct {
	Name  string `form:"name"`
	State uint8  `form:"state, default=0" binding:"oneof=0 1"`
}

type ListTagRequest struct {
	Name  string `form:"name"`
	State uint8  `form:"state, default=0"`
}

func (svc Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc Service) ListTag(param *ListTagRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.ListTag(param.Name, param.State, pager.Page, pager.PageSize)
}
