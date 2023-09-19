package dao

import (
	"membership_system/internal/model"
	"membership_system/pkg/app"
)

func (d Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := &model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d Dao) CountTag(name string, state uint8) (int, error) {
	tag := &model.Tag{
		Name:  name,
		State: state,
	}
	count, err := tag.Count(d.engine)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d Dao) ListTag(name string, state uint8, page int, pageSize int) ([]*model.Tag, error) {

	offset := app.GetPageOffset(page, pageSize)

	tag := &model.Tag{
		Name:  name,
		State: state,
	}

	tags, err := tag.List(d.engine, offset, pageSize)
	if err != nil {
		return nil, err
	}

	return tags, nil

}
