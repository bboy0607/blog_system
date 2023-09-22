package routers

import (
	"membership_system/global"
	"membership_system/internal/service"
	"membership_system/pkg/app"
	"membership_system/pkg/convert"
	"membership_system/pkg/errcode"
	"membership_system/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUplaod() Upload {
	return Upload{}
}

// 上傳單檔案
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParms)
		return
	}

	svc := service.New(c)
	//將fileType轉換成FileType類型
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

// 上傳多檔案
func (u Upload) UploadMultipleFiles(c *gin.Context) {
	response := app.NewResponse(c)
	//擷取所有MultipartForm至map中: 10 * 2^20 bytes = 10485760 Bytes = 10 MB
	err := c.Request.ParseMultipartForm(global.AppSetting.UploadMultiImageTotalMaxSize)
	if err != nil {
		errRsp := errcode.InvalidParms.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	//擷取type欄位
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	for _, fileHeaders := range c.Request.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				global.Logger.Errorf("Error open file: %v", err)
				response.ToErrorResponse(errcode.ErrorOpenFileFail)
				return
			}
			defer file.Close()

			svc := service.New(c)
			_, err = svc.UploadFile(upload.FileType(fileType), file, fileHeader)

			if err != nil {
				global.Logger.Errorf("svc.UploadFile err: %v", err)
				response.ToErrorResponse(errcode.ErrorUploadFileFail)
				return
			}

		}
	}

	response.ToResponse(gin.H{"message": "upload successed"})
}
