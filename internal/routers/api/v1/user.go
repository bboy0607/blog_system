package v1

import (
	"membership_system/global"
	"membership_system/internal/service"
	"membership_system/pkg/app"
	"membership_system/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// 註冊帳號
func (u User) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "使用者創建成功"})
	return
}

// 註冊需email認證帳號
func (u User) CreateEmailConfirmUser(c *gin.Context) {
	param := service.CreateEmailConfirmUserRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateEmailConfirmUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "使用者創建成功"})
	return
}

// 註冊會員，但會先發送email驗證
func (u User) ActivateEmailConfirmUser(c *gin.Context) {
	param := service.ActivateUserRequest{
		Token: c.Param("token"),
	}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.ActivateEmailConfirmUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "使用者創建成功"})
	return
}
