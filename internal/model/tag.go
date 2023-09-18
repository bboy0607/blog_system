package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) List(db *gorm.DB) ([]*Tag, error) {
	tags := []*Tag{}

	query := db.Model(&Tag{})
	if t.Name != "" {
		query.Where("name = ?", t.Name)
	}

	query = query.Where("state = ?", t.State)

	if err := query.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
