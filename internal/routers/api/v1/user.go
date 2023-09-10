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

// @Summary 創建帳號
// @Produce json
// @Param username query string true "使用者帳號" minlength(3) maxlength(100)
// @Param password query string true "使用者密碼" minlength(6) maxlength(100)
// @Param email query string true "使用者Email" format(email)
// @Param state query uint8 false "使用者狀態" Enums(0, 1) default(1)
// @Param created_by query string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users [post]
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
		response.ToErrorResponse(errcode.ErrorCreateUserFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"message": "使用者創建成功,等待Email驗證"})
	return
}

// 驗證Email，如果成功，啟動帳號 (state = 1)
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

	response.ToResponse(gin.H{"message": "Email驗證成功"})
	return
}

func (u User) SendResetPasswordEmail(c *gin.Context) {
	param := service.SendResetPasswordEmail{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err = svc.SendResetPasswordEmail(&param)
	if err != nil {
		global.Logger.Errorf("svc.SendResetPasswordEmail err: %v", err)
		response.ToErrorResponse(errcode.ErrorEmailNotFound.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"message": "密碼重置Email已發送"})
	return
}

func (u User) ResetPassword(c *gin.Context) {
	param := service.ResetUserPasswordRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	email, ok := c.Get("email")
	if !ok {
		global.Logger.Errorf("gin.Context Get err: %v", err)
		errRsp := errcode.ServerError.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	param.Email = email.(string)

	svc := service.New(c.Request.Context())

	err = svc.ResetUserPassword(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "密碼重置成功"})
	return
}

// 登入
func (u User) Login(c *gin.Context) {
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	loginToken, err := svc.UserLogin(&param)
	if err != nil {
		global.Logger.Errorf("svc.UserLogin err: %v", err)
		switch err {
		case errcode.ErrorUserNotFound:
			response.ToErrorResponse(errcode.ErrorUserNotFound)
		case errcode.ErrorPasswordNotCorrect:
			response.ToErrorResponse(errcode.ErrorPasswordNotCorrect)
		case errcode.ErrorUserNotActivated:
			response.ToErrorResponse(errcode.ErrorUserNotActivated)
		default:
			response.ToErrorResponse(errcode.ErrorUnknown.WithDetails(err.Error()))
		}
		return
	}

	response.ToResponse(gin.H{"message": "使用者登入成功", "login_token": loginToken})
	return
}

//登出

func (u User) Logout(c *gin.Context) {
	param := service.UserLogoutRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UserLogout(&param)
	if err != nil {
		switch err {
		case errcode.ErrorUserLoggedOut:
			global.Logger.Errorf("svc.UserLogout err: %v", err)
			response.ToErrorResponse(errcode.ErrorUserLoggedOut)
			return
		default:
			global.Logger.Errorf("svc.UserLogout err: %v", err)
			response.ToErrorResponse(errcode.ErrorUnknown.WithDetails(err.Error()))
			return
		}

	}

	response.ToResponse(gin.H{"message": "使用者成功登出"})
	return
}

func (u User) CreateUserInfo(c *gin.Context) {
	param := service.CreateUserInfoRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CreateUserInfo(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUserInfo err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserInfoFail)
		return
	}

	response.ToResponse(gin.H{"message": "使用者資訊建立成功"})
	return
}

func (u User) GetUserInfo(c *gin.Context) {
	param := service.GetUserInfoRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	result, err := svc.GetUserInfo(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetUserInfo err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserInfoFail)
		return
	}

	response.ToResponse(result)
	return
}
