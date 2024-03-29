package dao

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

func (d Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}

func (d Dao) ListTag(name string, state uint8, page int, pageSize int) ([]*model.Tag, error) {

	pageOffset := app.GetPageOffset(page, pageSize)

	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.List(d.engine, pageOffset, pageSize)
}

func (d Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	values := map[string]interface{}{
		"state":      state,
		"modifiedBy": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
