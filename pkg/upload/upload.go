package upload

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"membership_system/global"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

func generateRandomFileName(string) (string, error) {
	randomBytes := make([]byte, 16)
	// 生成隨機16位bytes切片
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 將[]byte轉換成16進位
	fileName := hex.EncodeToString(randomBytes)

	return fileName, nil
}

func GetFileName(name string) (string, error) {
	var err error
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName, err = generateRandomFileName(fileName)
	if err != nil {
		return "", err
	}

	return fileName + ext, nil
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 檢查是副檔名是否有在設定檔的允許清單，有則True，沒有則False
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	default:
		return false
	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		//如果是Image類型，檔案大小限制為設定檔的UploadImageMaxSize
		//1 MB = 1024 KB, 1KB = 1024 Bytes
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckPermission(dst string) bool {
	//讀取目標位置的State
	_, err := os.Stat(dst)
	//如果錯誤中含有Permission權限錯誤，則返回true
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	//創建多層級目錄
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
