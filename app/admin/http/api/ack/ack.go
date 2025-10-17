package ack

// UserLoginAck 用户登录-响应
//
//swagger:model
type UserLoginAck struct {
	Token string `json:"token" validate:"required" format:"string"` //访问令牌
}

// UserLogoutAck 用户退出登录-响应
//
//swagger:model
type UserLogoutAck struct {
}

// CommonUploadAck 通用上传-响应
//
//swagger:model
type CommonUploadAck struct {
	URL string `json:"url" validate:"required" format:"string"` //访问地址
}

// CommonSyncRemoteAssetAck 同步远端资源-响应
//
//swagger:model
type CommonSyncRemoteAssetAck struct {
	URL string `json:"url" validate:"required" format:"string"` //访问地址
}

// NavListAck 分类结构-响应
//
//swagger:model
type NavListAck struct {
	List []NavCategory `json:"list" validate:"required" format:"array<object>"` //列表
}

// NavCategoryListAck 分类列表-响应
//
//swagger:model
type NavCategoryListAck struct {
	List []NavCategory `json:"list" validate:"required" format:"array<object>"` //列表
}

// NavCategory 分类结构
//
//swagger:model
type NavCategory struct {
	Identity  string    `json:"identity" validate:"required" format:"string"`     //唯一标识符
	Name      string    `json:"name" validate:"required" format:"string"`         //名称
	Icon      string    `json:"icon" validate:"required" format:"string"`         //图标地址
	SortIndex int       `json:"sortIndex" validate:"required" format:"number"`    //排序，越小越靠前
	Sites     []NavSite `json:"sites" validate:"required" format:"array<object>"` //下属站点
}

// NavSiteListAck 分类列表-响应
//
//swagger:model
type NavSiteListAck struct {
	List []NavSite `json:"list" validate:"required" format:"array<object>"` //列表
}

// NavSite 站点结构
//
//swagger:model
type NavSite struct {
	Identity         string `json:"identity" validate:"required" format:"string"`         //唯一标识符
	CategoryIdentity string `json:"categoryIdentity" validate:"required" format:"string"` //所属分类唯一标识符
	Name             string `json:"name" validate:"required" format:"string"`             //名称
	Description      string `json:"description" validate:"required" format:"string"`      //描述
	Icon             string `json:"icon" validate:"required" format:"string"`             //图标地址
	URL              string `json:"url" validate:"required" format:"string"`              //站点链接地址
	SortIndex        int    `json:"sortIndex" validate:"required" format:"number"`        //排序，越小越靠前
}

// NavCategorySortedAck 对分类进行排序-响应
//
//swagger:model
type NavCategorySortedAck struct {
}

// NavSiteSortedAck 对站点进行排序-响应
//
//swagger:model
type NavSiteSortedAck struct {
}

// NavCategorySaveAck 保存分类-响应
//
//swagger:model
type NavCategorySaveAck struct {
}

// NavCategoryDeleteAck 删除分类-响应
//
//swagger:model
type NavCategoryDeleteAck struct {
}

// NavSiteSaveAck 保存站点-响应
//
//swagger:model
type NavSiteSaveAck struct {
}

// NavSiteDeleteAck 删除站点-响应
//
//swagger:model
type NavSiteDeleteAck struct {
}
