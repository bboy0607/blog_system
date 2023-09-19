package service

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,max=100"`
	State     uint8  `form:"state, default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,max=100"`
}

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type ListTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
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

func (svc Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
