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

// 驗證登入token是否正確
func ValidateLoginToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			requestToken string
			username     string
			ecode        = errcode.Success
		)

		requestToken = c.GetHeader("token")
		username = c.GetHeader("username")

		if requestToken == "" || username == "" {
			ecode = errcode.InvalidParms
		} else {
			ctx := context.Background()
			key := fmt.Sprintf("loginToken:%v", username)
			userLoginToken, err := global.Redis.Get(ctx, key).Result()
			if err != nil || requestToken != userLoginToken {
				ecode = errcode.UnauthorizedTokenError
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()

	}
}

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
