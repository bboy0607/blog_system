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

func (d *Dao) CreateUser(username string, password string, email string, state uint8, createdBy string) error {
	user := model.User{
		Username: username,
		Password: password,
		Email:    email,
		State:    state,
		Model:    &model.Model{CreatedBy: createdBy},
	}

	return user.Create(d.engine)
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

func (d *Dao) ActivateUser(username string, modifiedBy string) error {
	user := model.User{
		Username: username,
		State:    1,
		Model:    &model.Model{ModifiedBy: modifiedBy},
	}

	return user.Activate(d.engine)
}

func (d *Dao) CheckEmail(email string) error {
	user := model.User{
		Email: email,
	}
	return user.CheckEmail(d.engine)
}

func (d *Dao) ResetUserPassword(email string, newPassword string, modifiedBy string) error {
	user := model.User{
		Email:    email,
		Password: newPassword,
		Model:    &model.Model{ModifiedBy: modifiedBy},
	}

	return user.ResetUserPassword(d.engine)
}

func (d *Dao) ValidateUserCredentials(username string, password string) error {
	user := model.User{
		Username: username,
		Password: password,
	}

	return user.ValidateUserCredentials(d.engine)
}
