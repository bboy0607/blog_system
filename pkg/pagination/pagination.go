package pagination

import (
	"membership_system/global"
	"membership_system/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	//從url中取得page參數
	page := convert.StrTo(c.Query("page")).MustInt()
	//如果 page=0 或不存在，則返回 1
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()

	//如果pageSize不存在，則返回設定檔的DefaultPageSize
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}

	//如果PageSize大於設定檔的MaxPageSize，則使用設定檔的設定
	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pagesize int) int {
	result := 0

	//如果page大於0，才會計算PageOffset，否則返回 0
	if page > 0 {
		result = (page - 1) * pagesize
	}

	return result
}
