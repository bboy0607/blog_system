package service

import (
	"errors"
	"membership_system/global"
	"membership_system/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName, err := upload.GetFileName(fileHeader.Filename)
	if err != nil {
		return nil, errors.New("failed to get fileName.")
	}

	//使用寫好的GetSavePath()取得設定檔中的儲存位置
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	//如果副檔名或類型不在設定中，則返回錯誤
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix or file type is not supporterd.")
	}

	//如果儲存位置目錄不存在，則創建目錄
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}

	//確認檔案大小是否超過所屬類型的最大值
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	//確認檔案權限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	//上傳檔案
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil

}
