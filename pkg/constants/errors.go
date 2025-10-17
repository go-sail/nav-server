package constants

import (
	"github.com/keepchen/go-sail/v3/constants"
	"github.com/keepchen/go-sail/v3/sail"
)

const (
	ErrInternalServerError       = constants.CodeType(9999) //服务器内部错误
	ErrNone                      = constants.CodeType(200)  //无错误（请求成功）
	ErrRequestParamsInvalid      = constants.CodeType(1000) //请求参数有误
	ErrAuthorizationTokenInvalid = constants.CodeType(1001) //令牌已失效
	ErrUsernamePasswordNotMatch  = constants.CodeType(1002) //帐号或密码错误
	ErrFetchRemoteAssetFailed    = constants.CodeType(1003) //拉取远端资源失败
	ErrCategoryNotFound          = constants.CodeType(1004) //分类不存在
	ErrSiteNotFound              = constants.CodeType(1005) //站点不存在
)

var errorsMap = map[constants.ICodeType]string{
	ErrNone:                      "SUCCESS",
	ErrRequestParamsInvalid:      "Bad request parameters",
	ErrAuthorizationTokenInvalid: "Authorization token invalid",
	ErrInternalServerError:       "Internal server error",
	ErrUsernamePasswordNotMatch:  "Username password not match",
	ErrFetchRemoteAssetFailed:    "Failed to fetch remote asset",
	ErrCategoryNotFound:          "Category not found",
	ErrSiteNotFound:              "Site not found",
}

// RegisterErrorCode 注册错误码
func RegisterErrorCode() {
	for code, msg := range errorsMap {
		sail.Code().Register("en", code.Int(), msg)
	}
}
