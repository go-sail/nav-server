package req

import (
	"fmt"
	sailConstants "github.com/keepchen/go-sail/v3/constants"
	"nav-server/pkg/constants"
	"unicode/utf8"
)

// UserLoginReq 用户登录-请求
//
//swagger:model
type UserLoginReq struct {
	Username string `form:"username" validate:"required" format:"string"` //用户名
	Password string `form:"password" validate:"required" format:"string"` //密码
}

func (v UserLoginReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.Username) == 0 || len(v.Password) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("username or password is empty")
	}
	return constants.ErrNone, nil
}

// UserLogoutReq 用户退出登录-请求
//
//swagger:model
type UserLogoutReq struct {
}

func (v UserLogoutReq) Validator() (sailConstants.ICodeType, error) {
	return constants.ErrNone, nil
}

// CommonUploadReq 通用上传-请求
//
//swagger:model
type CommonUploadReq struct {
	Action string `form:"action" validate:"required" format:"string" enums:"category,site"` //操作类型
}

func (v CommonUploadReq) Validator() (sailConstants.ICodeType, error) {
	if v.Action != "category" && v.Action != "site" {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("action %s is invalid", v.Action)
	}

	return constants.ErrNone, nil
}

// CommonSyncRemoteAssetReq 同步远端资源-请求
//
//swagger:model
type CommonSyncRemoteAssetReq struct {
	URL    string `form:"url" validate:"required" format:"string"`                          //远端资源地址
	Action string `form:"action" validate:"required" format:"string" enums:"category,site"` //操作类型
}

func (v CommonSyncRemoteAssetReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.URL) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("url is empty")
	}
	if v.Action != "category" && v.Action != "site" {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("action %s is invalid", v.Action)
	}

	return constants.ErrNone, nil
}

// NavListReq 获取全部导航列表-请求
//
//swagger:model
type NavListReq struct{}

func (v NavListReq) Validator() (sailConstants.ICodeType, error) {
	return constants.ErrNone, nil
}

// NavCategoryListReq 获取全部分类列表-请求
//
//swagger:model
type NavCategoryListReq struct{}

func (v NavCategoryListReq) Validator() (sailConstants.ICodeType, error) {
	return constants.ErrNone, nil
}

// NavSiteListReq 获取站点列表-请求
//
//swagger:model
type NavSiteListReq struct {
	CategoryIdentity string `json:"categoryIdentity" form:"categoryIdentity" validate:"required" format:"string"` //分类唯一标识
}

func (v NavSiteListReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.CategoryIdentity) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("categoryIdentity is empty")
	}

	return constants.ErrNone, nil
}

// NavCategorySortedReq 对分类进行排序-请求
//
//swagger:model
type NavCategorySortedReq struct {
	Identities []string `json:"identities" form:"identities" validate:"required" format:"array<string>"` //分类唯一标识列表
}

func (v NavCategorySortedReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.Identities) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("identities is empty")
	}

	return constants.ErrNone, nil
}

// NavSiteSortedReq 对站点进行排序-请求
//
//swagger:model
type NavSiteSortedReq struct {
	CategoryIdentity string   `json:"categoryIdentity" form:"categoryIdentity" validate:"required" format:"string"` //分类唯一标识
	Identities       []string `json:"identities" form:"identities" validate:"required" format:"array<string>"`      //站点唯一标识列表
}

func (v NavSiteSortedReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.CategoryIdentity) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("categoryIdentity is empty")
	}
	if len(v.Identities) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("identities is empty")
	}

	return constants.ErrNone, nil
}

// NavCategorySaveReq 保存分类-请求
//
//swagger:model
type NavCategorySaveReq struct {
	Identity string `json:"identity" format:"string"`                                  //唯一标识符（为空则表示创建）
	Name     string `json:"name" validate:"required" format:"string" maxLength:"255"`  //名称
	Icon     string `json:"icon" validate:"required" format:"string" maxLength:"1024"` //图标地址
}

func (v NavCategorySaveReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.Name) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("name is empty")
	}
	if utf8.RuneCountInString(v.Name) > 255 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("name is too long")
	}
	if len(v.Icon) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("icon is empty")
	}
	if len(v.Icon) > 1024 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("icon is too long")
	}

	return constants.ErrNone, nil
}

// NavCategoryDeleteReq 删除分类-请求
//
//swagger:model
type NavCategoryDeleteReq struct {
	Identity string `json:"identity" form:"identity" validate:"required" format:"string"` //唯一标识符
}

func (v NavCategoryDeleteReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.Identity) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("identity is empty")
	}

	return constants.ErrNone, nil
}

// NavSiteDeleteReq 删除站点-请求
//
//swagger:model
type NavSiteDeleteReq struct {
	Identity string `json:"identity" form:"identity" validate:"required" format:"string"` //唯一标识符
}

func (v NavSiteDeleteReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.Identity) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("identity is empty")
	}

	return constants.ErrNone, nil
}

// NavSiteSaveReq 保存分类-请求
//
//swagger:model
type NavSiteSaveReq struct {
	CategoryIdentity string `json:"categoryIdentity" form:"categoryIdentity" validate:"required" format:"string"`       //分类唯一标识
	Identity         string `json:"identity" form:"identity" format:"string"`                                           //唯一标识符（为空则表示创建）
	Name             string `json:"name" form:"name" validate:"required" format:"string" maxLength:"255"`               //名称
	Description      string `json:"description" form:"description" validate:"required" format:"string" maxLength:"255"` //描述
	Icon             string `json:"icon" form:"icon" validate:"required" format:"string" maxLength:"1024"`              //图标地址
	URL              string `json:"url" form:"url" validate:"required" format:"string" maxLength:"255"`                 //地址
}

func (v NavSiteSaveReq) Validator() (sailConstants.ICodeType, error) {
	if len(v.CategoryIdentity) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("categoryIdentity is empty")
	}
	if len(v.Name) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("name is empty")
	}
	if utf8.RuneCountInString(v.Name) > 255 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("name is too long")
	}
	if len(v.Icon) == 0 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("icon is empty")
	}
	if len(v.Icon) > 1024 {
		return constants.ErrRequestParamsInvalid, fmt.Errorf("icon is too long")
	}

	return constants.ErrNone, nil
}
