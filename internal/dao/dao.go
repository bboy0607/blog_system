package dao

import (
	"membership_system/internal/model"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CreateUserInfo(userID string, nickname string, gender string, createdBy string) error {
	userInfo := model.UserInfo{
		UserID:   userID,
		Nickname: nickname,
		Gender:   gender,
		Model:    &model.Model{CreatedBy: createdBy},
	}
	return userInfo.Create(d.engine)
}

func (d *Dao) GetUserInfo(userID string) (*model.UserInfo, error) {
	userInfo := model.UserInfo{
		UserID: userID,
	}
	return userInfo.Get(d.engine)
}
