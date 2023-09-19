package errcode

//使用者相關錯誤
var (
	ErrorCreateUserFail     = NewErrorCode(2001002, "建立使用者失敗")
	ErrorUserNotFound       = NewErrorCode(2001003, "查無使用者")
	ErrorEmailNotFound      = NewErrorCode(2001004, "查無email")
	ErrorUserNotActivated   = NewErrorCode(2001005, "使用者未啟用")
	ErrorPasswordNotCorrect = NewErrorCode(2001006, "密碼錯誤")
	ErrorUserLoggedOut      = NewErrorCode(2001007, "使用者已登出")
	ErrorCreateUserInfoFail = NewErrorCode(2001008, "建立使用者資訊失敗")
	ErrorGetUserInfoFail    = NewErrorCode(2001009, "取得使用者資訊失敗")
	ErrorUnknown            = NewErrorCode(2001999, "未知錯誤")
)

//Tag相關錯誤
var (
	ErrorCreateTagFail = NewErrorCode(2002001, "建立標籤失敗")
	ErrorCountTagFail  = NewErrorCode(2002002, "查詢標籤數量失敗")
	ErrorListTagFail   = NewErrorCode(2002003, "查詢標籤清單失敗")
	ErrorUpdateTagFail = NewErrorCode(2002005, "更新標籤失敗")
	ErrorDeleteTagFail = NewErrorCode(2002004, "刪除標籤失敗")
)

//文章相關錯誤
var (
	ErrorCreateArticleFail = NewErrorCode(2003001, "建立文章失敗")
	ErrorCountArticleFail  = NewErrorCode(2003002, "查詢文章數量失敗")
	ErrorListArticleFail   = NewErrorCode(2003003, "查詢文章清單失敗")
)
