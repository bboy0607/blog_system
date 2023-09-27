package service

import (
	"context"
	"errors"
	"fmt"
	"membership_system/global"
	"membership_system/internal/model"
	"membership_system/pkg/email"
	"membership_system/pkg/errcode"
	"membership_system/pkg/pwd"
	"membership_system/pkg/util"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=6,max=100"`
	Password  string `form:"password" binding:"required,min=6,max=100"`
	Email     string `form:"email" binding:"required,max=100,email"`
	State     uint8  `form:"state,default=1" binding:"oneof=01"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type CreateEmailConfirmUserRequest struct {
	Username string `form:"username" binding:"required,min=6,max=100"`
	Password string `form:"password" binding:"required,min=6,max=100"`
	Email    string `form:"email" binding:"required,max=100,email"`
}

type ActivateUserRequest struct {
	Token string `form:"token" binding:"required"`
}

type SendResetPasswordEmail struct {
	Email string `form:"email" binding:"email,required"`
}

type ResetUserPasswordRequest struct {
	Email              string
	NewPassword        string `form:"new_password" binding:"min=3,max=100,required"`
	NewConfirmPassword string `form:"confirm_new_password" binding:"min=3,max=100,required"`
}

type UserLoginRequest struct {
	Username string `form:"username" binding:"min=3,max=100,required"`
	Password string `form:"password" binding:"min=3,max=100,required"`
}

type UserLogoutRequest struct {
	Username string `form:"username" binding:"min=3,max=100,required"`
}

type CreateUserInfoRequest struct {
	UserID   string `form:"user_id" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	Gender   string `form:"gender" binding:"required,oneof=男 女"`
}

type GetUserInfoRequest struct {
	UserID string `form:"user_id" binding:"required"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.Email, param.State, param.CreatedBy)
}

func (svc *Service) CreateEmailConfirmUser(param *CreateEmailConfirmUserRequest) error {
	//密碼使用bcrypt加密
	hashedPassword, err := pwd.HashPassword(param.Password)
	if err != nil {
		return err
	}

	//使用dao創建使用方法
	err = svc.dao.CreateUser(param.Username, hashedPassword, param.Email, 0, "backend_system")
	if err != nil {
		return err
	}

	//建立驗證Token存入Redis中
	var ctx = context.Background()
	token := util.GenerateSecureToken(10)
	err = global.Redis.Set(ctx, token, param.Username, 10*time.Minute).Err()
	if err != nil {
		return err
	}

	//發送驗證郵件
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

	return nil
}

func (svc *Service) CreateUserInfo(param *CreateUserInfoRequest) error {
	return svc.dao.CreateUserInfo(param.UserID, param.Nickname, param.Gender, "backend_system")
}

func (svc *Service) GetUserInfo(param *GetUserInfoRequest) (*model.UserInfo, error) {
	return svc.dao.GetUserInfo(param.UserID)
}

func (svc *Service) ActivateEmailConfirmUser(param *ActivateUserRequest) error {
	var ctx = context.Background()
	username, err := global.Redis.Get(ctx, param.Token).Result()
	if err != nil {
		return err
	}
	return svc.dao.ActivateUser(username, "backend_system")
}

func (svc *Service) SendResetPasswordEmail(param *SendResetPasswordEmail) error {
	if err := svc.dao.CheckEmail(param.Email); err != nil {
		return err
	}

	ctx := context.Background()
	resetPasswordToken := util.GenerateSecureToken(10)
	key := fmt.Sprintf("resetPasswordToken:%v", resetPasswordToken)
	global.Redis.Set(ctx, key, param.Email, 10*time.Minute)

	email := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	to := []string{param.Email}
	err := email.SendResetPasswordEmail(to, "email驗證郵件", resetPasswordToken)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) ResetUserPassword(param *ResetUserPasswordRequest) error {
	if param.NewPassword != param.NewConfirmPassword {
		return errors.New("新密碼與確認密碼不符合")
	}
	return svc.dao.ResetUserPassword(param.Email, param.NewPassword, "backend_system")
}

func (svc *Service) UserLogin(param *UserLoginRequest, c *gin.Context) (loginToken string, err error) {
	user, err := svc.dao.GetUserByUsername(param.Username)
	if err != nil {
		return "", err
	}

	//驗證密碼
	if err := pwd.VerifyPassword(user.Password, param.Password); err != nil {
		return "", errcode.ErrorPasswordNotCorrect
	}

	session := sessions.Default(c)
	sessionID := session.Get("session_id")

	if sessionID == nil {
		//使用寫好的GenerateRandomSessionID函數產生SessionID
		sessionID := util.GenerateRandomSessionID()
		session.Set("session_id", sessionID)
		err = session.Save()
		if err != nil {
			return "", err
		}
	}

	// ctx := context.Background()
	// loginToken = util.GenerateSecureToken(10)
	// key := fmt.Sprintf("loginToken:%v", param.Username)
	// global.Redis.Set(ctx, key, loginToken, 0)
	// return loginToken, nil

	return "", nil
}

func (svc *Service) UserLogout(param *UserLogoutRequest, c *gin.Context) error {
	// ctx := context.Background()
	// key := fmt.Sprintf("loginToken:%v", param.Username)
	// count, err := global.Redis.Del(ctx, key).Result()
	// if err != nil {
	// 	return err
	// }
	// if count == 0 {
	// 	return errcode.ErrorUserLoggedOut
	// }
	session := sessions.Default(c)
	session.Delete("session_id")
	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}
