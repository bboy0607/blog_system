package routers

import (
	"membership_system/internal/middleware"
	v1 "membership_system/internal/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("static/*.html")

	user := v1.NewUser()
	apiv1 := r.Group("/api/v1")
	{
		//使用者登入
		apiv1.POST("/user/login", user.Login)

		//註冊會員
		apiv1.POST("/user/register", user.CreateEmailConfirmUser)
		apiv1.GET("/user/verify_email/:token", user.ActivateEmailConfirmUser)

		//忘記密碼
		apiv1.POST("user/reset_password", user.SendResetPasswordEmail)
		apiv1.GET("user/reset_password/:token", middleware.ValidatePasswordResetToken(), func(c *gin.Context) { c.HTML(http.StatusOK, "reset_password.html", nil) })
		apiv1.POST("/user/reset_password/:token", middleware.ValidatePasswordResetToken(), user.ResetPassword)
	}

	return r
}
