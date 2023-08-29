package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

//註冊帳號
func (u User) Create(c *gin.Context) {}

//查會員
func (u User) GET(c *gin.Context) {}
