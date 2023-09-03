package middleware

import (
	"context"
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

		username, err := global.Redis.Get(ctx, token).Result()
		if err != nil {
			errRsp := errcode.InvalidParms.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			c.Abort()
		}
		c.Set("token", token)
		c.Set("username", username)
		c.Next()
	}

}
