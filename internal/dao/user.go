package dao

import (
	"membership_system/internal/model"
)

type User struct {
	ID         uint32
	Username   string
	Password   string
	Email      string
	State      uint8
	CreatedBy  string
	ModifiedBy string
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

func (d *Dao) GetUserByUsername(username string) (model.User, error) {
	user := model.User{
		Username: username,
	}
	return user.GetByUsername(d.engine)
}

func (d *Dao) UpdateUser(param *User) error {
	user := model.User{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Password != "" {
		values["password"] = param.Password
	}
	if param.Email != "" {
		values["email"] = param.Email
	}
	return user.Update(d.engine, values)
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
