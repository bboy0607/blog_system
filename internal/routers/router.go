package routers

import (
	v1 "membership_system/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := v1.NewUser()
	apiv1 := r.Group("/api/v1")
	{
		//註冊會員
		apiv1.POST("/user", user.Create)
		//會員登入
		apiv1.GET("/user/:username", user.GET)

	}

	return r
}
