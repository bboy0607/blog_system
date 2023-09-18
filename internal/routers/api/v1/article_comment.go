package v1

import "github.com/gin-gonic/gin"

type ArticleComment struct{}

func NewArticleComment() ArticleComment {
	return ArticleComment{}
}

func (a ArticleComment) Create(c *gin.Context) {}

func (a ArticleComment) List(c *gin.Context) {}

func (a ArticleComment) Update(c *gin.Context) {}

func (a ArticleComment) Delete(c *gin.Context) {}
