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

// @Summary 註冊需email認證帳號
// @Produce json
// @Param username formData string true "使用者帳號" minlength(3) maxlength(100)
// @Param password formData string true "使用者密碼" minlength(6) maxlength(100)
// @Param email formData string true "使用者Email" format(email)
// @Param created_by formData string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/register [post]
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

// ActivateEmailConfirmUser 驗證使用者帳戶Email
// @Summary 驗證使用者帳戶Email
// @Description 驗證使用者帳戶Email
// @Accept json
// @Produce json
// @Param token query string true "email驗證token"
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/verify_email/{token} [get]
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

// SendResetPasswordEmail 發送重置密碼的Email
// @Summary 發送重置密碼的Email
// @Description 發送重置密碼的Email
// @Accept json
// @Produce json
// @Param email formData string true "用戶Email"
// @Success 200 {object} gin.H "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/reset_password_email [post]
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

// ResetPassword 重置使用者密碼
// @Summary 重置使用者密碼
// @Description 重置使用者密碼
// @Accept json
// @Produce json
// @Param password formData string true "新密碼"
// @Success 200 {object} gin.H "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/reset_password [post]
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

// Login 使用者登入
// @Summary 使用者登入
// @Description 使用者登入
// @Accept json
// @Produce json
// @Param username formData string true "使用者帳號"
// @Param password formData string true "使用者密碼"
// @Success 200 {object} gin.H "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 401 {object} errcode.Error "使用者未找到、密碼不正確、使用者未啟用等錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/login [post]
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
	_, err = svc.UserLogin(&param, c)
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

	response.ToResponse(gin.H{"message": "使用者登入成功", "login_token": ""})
	return
}

// Logout 使用者登出
// @Summary 使用者登出
// @Description 使用者登出
// @Produce json
// @Success 200 {object} gin.H "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 401 {object} errcode.Error "使用者已登出等錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/logout [get]
func (u User) Logout(c *gin.Context) {
	param := service.UserLogoutRequest{}
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	err := svc.UserLogout(&param, c)
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

// @Summary 創建使用者資訊
// @Description 創建使用者資訊
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id formData string true "使用者ID"
// @Param nickname formData string true "暱稱"
// @Param gender formData string true "性別"
// @Success 200 {object} gin.H{"message": "使用者資訊建立成功"}
// @Failure 400 {object}  errcode.Error "參數錯誤"
// @Failure 500 {object}  errcode.Error "內部錯誤"
// @Router /api/v1/users/info [post]
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

// @Summary 取得使用者資訊
// @Description 取得特定使用者的詳細資訊
// @Tags 使用者
// @Accept json
// @Produce json
// @Param id path int true "使用者ID" Format(int64)
// @Success 200 {object} model.UserInfo "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/users/info/{id} [get]
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
