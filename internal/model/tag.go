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

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	query := db.Model(&t)

	//如果名稱不為空值，則查詢條件增加名稱查詢
	if t.Name != "" {
		query = query.Where("name = ?", t.Name)
	}

	query = query.Where("state = ?", t.State)

	err := query.Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Tag, error) {
	tags := []*Tag{}

	query := db.Model(&Tag{})
	if t.Name != "" {
		query.Where("name = ?", t.Name)
	}

	//新增查詢State條件與分頁功能
	query = query.Where("state = ?", t.State).Offset(pageOffset).Limit(pageSize)

	if err := query.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

//更新Tag
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND is_del = ? ", t.ID, 0).Updates(values).Error
}

//刪除Tag
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
