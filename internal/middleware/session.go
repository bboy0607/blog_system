package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ValidateSessionID() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("session_id")

		if sessionID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未經授權的訪問"})
			c.Abort()
			return
		}

		c.Next()
	}
}
