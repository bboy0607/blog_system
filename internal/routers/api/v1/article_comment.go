package v1

import (
	"membership_system/global"
	"membership_system/internal/service"
	"membership_system/pkg/app"
	"membership_system/pkg/convert"
	"membership_system/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type ArticleComment struct{}

func NewArticleComment() ArticleComment {
	return ArticleComment{}
}

func (a ArticleComment) Create(c *gin.Context) {
	param := service.CreateArticleCommentRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	err = svc.CreateArticleComment(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticleComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleCommentFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a ArticleComment) List(c *gin.Context) {}

func (a ArticleComment) Update(c *gin.Context) {
	param := service.UpdateArticleCommentRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	err = svc.UpdateArticleComment(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticleComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleCommentFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a ArticleComment) Delete(c *gin.Context) {
	param := service.DeleteArticleCommentRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	err = svc.DeleteArticleComment(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticleComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleCommentFail)
		return
	}

	response.ToResponse(gin.H{})
}
