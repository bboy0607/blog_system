package v1

import (
	"membership_system/global"
	"membership_system/internal/service"
	"membership_system/pkg/app"
	"membership_system/pkg/convert"
	"membership_system/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	err = svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) List(c *gin.Context) {
	param := service.ListArticleRequest{}
	response := app.NewResponse(c)

	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	totalRows, err := svc.CountArticle(&service.CountArticleRequest{Title: param.Title, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}

	pager := app.Pager{
		Page:      app.GetPage(c),
		PageSize:  app.GetPageSize(c),
		TotalRows: totalRows,
	}

	articles, err := svc.ListAricle(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.ListAricle: %v", err)
		response.ToErrorResponse(errcode.ErrorListArticleFail)
		return
	}

	response.ToResponseList(articles, totalRows)
}

func (a Article) Get(c *gin.Context) {
	param := service.GetArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
}

func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Errorf("gin.Context ShouldBind err: %v", err)
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	err = svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) Delete(c *gin.Context) {

}
