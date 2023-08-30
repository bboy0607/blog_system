package service

import (
	"context"
	"fmt"
	"membership_system/global"
	"membership_system/pkg/email"
	"membership_system/pkg/util"
	"time"
)

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=6,max=100"`
	Password  string `form:"password" binding:"required,min=6,max=100"`
	Email     string `form:"email" binding:"required,max=100,email"`
	State     uint8  `form:"state,default=1" binding:"oneof=01"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type CreateEmailConfirmUserRequest struct {
	Username  string `form:"username" binding:"required,min=6,max=100"`
	Password  string `form:"password" binding:"required,min=6,max=100"`
	Email     string `form:"email" binding:"required,max=100,email"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type ActivateUserRequest struct {
	Token string `form:"token" binding:"required"`
}

func (svc Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.Email, param.State, param.CreatedBy)
}

func (svc Service) CreateEmailConfirmUser(param *CreateEmailConfirmUserRequest) error {
	var ctx = context.Background()
	token := util.GenerateSecureToken(10)
	err := global.Redis.Set(ctx, token, param.Username, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	email := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	to := []string{param.Email}
	err = email.SendConfirmationEmail(to, "email驗證郵件", token)
	if err != nil {
		return err
	}

	return svc.dao.CreateUser(param.Username, param.Password, param.Email, 0, param.CreatedBy)
}

func (svc Service) ActivateEmailConfirmUser(param *ActivateUserRequest) error {
	var ctx = context.Background()
	username, err := global.Redis.Get(ctx, param.Token).Result()
	fmt.Print(username)
	if err != nil {
		return err
	}
	return svc.dao.ActivateUser(username, "backend_system")
}
