package middleware

import (
	"context"
	"fmt"
	"membership_system/global"
	"membership_system/pkg/app"
	"membership_system/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidatePasswordResetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := app.NewResponse(c)
		ctx := context.Background()
		token := c.Param("token")
		if token == "" {
			c.AbortWithStatus(http.StatusNotFound)
		}

		key := fmt.Sprintf("resetPasswordToken:%v", token)
		email, err := global.Redis.Get(ctx, key).Result()
		if err != nil {
			errRsp := errcode.InvalidParms.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			c.Abort()
		}
		c.Set("token", token)
		c.Set("email", email)
		c.Next()
	}

}
