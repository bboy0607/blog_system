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
	ErrorDeleteTagFail = NewErrorCode(2002006, "刪除標籤失敗")
)

//文章相關錯誤
var (
	ErrorCreateArticleFail = NewErrorCode(2003001, "建立文章失敗")
	ErrorCountArticleFail  = NewErrorCode(2003002, "查詢文章數量失敗")
	ErrorListArticleFail   = NewErrorCode(2003003, "查詢文章清單失敗")
	ErrorGetArticleFail    = NewErrorCode(2003004, "取得文章失敗")
	ErrorUpdateArticleFail = NewErrorCode(2003005, "更新文章失敗")
	ErrorDeleteArticleFail = NewErrorCode(2003006, "刪除文章失敗")
)

//文章評論錯誤
var (
	ErrorCreateArticleCommentFail = NewErrorCode(2004001, "建立文章評論失敗")
	ErrorCountArticleCommentFail  = NewErrorCode(2004002, "查詢文章評論數量失敗")
	ErrorListArticleCommentFail   = NewErrorCode(2004004, "查詢文章評論清單失敗")
	ErrorUpdateArticleCommentFail = NewErrorCode(2004005, "更新文章評論失敗")
)

//上傳檔案相關錯誤
var (
	ErrorUploadFileFail = NewErrorCode(2005001, "上傳檔案失敗")
	ErrorOpenFileFail   = NewErrorCode(2006002, "開啟檔案失敗")
)
