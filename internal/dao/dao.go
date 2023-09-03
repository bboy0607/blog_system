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

func (d *Dao) ActivateUser(username string, modifiedBy string) error {
	user := model.User{
		Username: username,
		State:    1,
		Model:    &model.Model{ModifiedBy: modifiedBy},
	}

	return user.Activate(d.engine)
}

func (d *Dao) ResetUserPassword(username string, newPassword string, modifiedBy string) error {
	user := model.User{
		Username: username,
		Password: newPassword,
		Model:    &model.Model{ModifiedBy: modifiedBy},
	}

	return user.ResetUserPassword(d.engine)
}
