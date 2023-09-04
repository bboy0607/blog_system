package errcode

var (
	ErrorCreateUserFail     = NewErrorCode(2001002, "建立使用者失敗")
	ErrorUserNotFound       = NewErrorCode(2001003, "查無使用者")
	ErrorEmailNotFound      = NewErrorCode(2001004, "查無email")
	ErrorUserNotActivated   = NewErrorCode(2001005, "使用者未啟用")
	ErrorPasswordNotCorrect = NewErrorCode(2001006, "密碼錯誤")
	ErrorUnknown            = NewErrorCode(2001999, "未知錯誤")
)
