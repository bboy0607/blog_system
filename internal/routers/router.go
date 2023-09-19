package routers

import (
	"membership_system/internal/middleware"
	v1 "membership_system/internal/routers/api/v1"
	"net/http"

	_ "membership_system/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLGlob("static/*.html")

	user := v1.NewUser()
	userApi := r.Group("/api/v1/users")
	{
		//使用者登入
		userApi.POST("login", user.Login)

		//使用者登出
		userApi.GET("logout", user.Logout)

		//註冊會員
		userApi.POST("register", user.CreateEmailConfirmUser)
		userApi.GET("verify_email/:token", user.ActivateEmailConfirmUser)

		//忘記密碼
		userApi.POST("reset_password", user.SendResetPasswordEmail)
		userApi.GET("reset_password/:token", middleware.ValidatePasswordResetToken(), func(c *gin.Context) { c.HTML(http.StatusOK, "reset_password.html", nil) })
		userApi.POST("reset_password/:token", middleware.ValidatePasswordResetToken(), user.ResetPassword)

		//建立使用者資訊
		userApi.POST("info", user.CreateUserInfo)

		//使用登入Token接收會員資料
		userApi.GET("info", middleware.ValidateLoginToken(), user.GetUserInfo)
	}

	tag := v1.NewTag()
	tagApi := r.Group("/api/v1/tags")
	{
		//建標籤路由
		tagApi.POST("", tag.Create)

		//查詢標籤清單
		tagApi.GET("", tag.List)

		//更新標籤路由
		tagApi.PATCH("/:id", tag.Update)

		//刪除標籤路由
		tagApi.DELETE("/:id", tag.Delete)

	}

	article := v1.NewArticle()
	articleApi := r.Group("/api/v1/articles")
	{
		//新增文章
		articleApi.POST("", article.Create)

		//查找文章
		articleApi.GET("/:id", article.Get)

		//列出文章清單
		articleApi.GET("", article.List)

		//修改文章
		articleApi.PATCH("/:id", article.Update)

		//刪除文章
		articleApi.DELETE("/:id", article.Delete)

	}

	articleComment := v1.NewArticleComment()
	articleCommentApi := r.Group("/api/v1/article_comment")
	{
		articleCommentApi.POST("/", articleComment.Create)

		articleCommentApi.GET("/", articleComment.List)

		articleCommentApi.PATCH("/:id", articleComment.Update)

		articleCommentApi.DELETE("/:id", articleComment.Delete)
	}

	return r
}
