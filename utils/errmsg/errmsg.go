package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	// code = 2000... 分类模块错误
	ERROR_CATEGORY_USED      = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002
	// code = 3000... 文章模块错误
	ERROR_ART_NOT_EXIST = 3001
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_TOKEN_RUNTIME:  "TOKEN已过期",
	ERROR_TOKEN_EXIST:    "TOKEN不存在",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_USER_NOT_EXIST: "用户不存在",

	ERROR_CATEGORY_USED:      "该分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "该分类不存在",

	ERROR_ART_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}