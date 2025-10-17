package handler

import (
	"github.com/gin-gonic/gin"
	"nav-server/app/admin/http/service"
)

type navCtrl struct{}

var Nav = &navCtrl{}

// List 获取完整的导航列表
// @Tags        nav - 导航模块
// @Summary     apis/nav/list - 获取完整的导航列表
// @Description 获取完整的导航列表
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param parameter query req.NavListReq true "查询参数"
// @Success     200   {object} ack.NavListAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/list [get]
func (*navCtrl) List(c *gin.Context) {
	service.Nav.List(c)
}

// Categories 获取分类列表
// @Tags        nav - 导航模块
// @Summary     apis/nav/categories - 获取分类列表
// @Description 获取分类列表
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param parameter query req.NavCategoryListReq true "查询参数"
// @Success     200   {object} ack.NavCategoryListAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/categories [get]
func (*navCtrl) Categories(c *gin.Context) {
	service.Nav.Categories(c)
}

// Sites 按分类获取站点列表
// @Tags        nav - 导航模块
// @Summary     apis/nav/sites - 按分类获取站点列表
// @Description 按分类获取站点列表
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param parameter query req.NavSiteListReq true "查询参数"
// @Success     200   {object} ack.NavSiteListAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/sites [get]
func (*navCtrl) Sites(c *gin.Context) {
	service.Nav.Sites(c)
}

// SaveCategory 保存分类
// @Tags        nav - 导航模块
// @Summary     apis/nav/category - 保存分类
// @Description 保存分类
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload body req.NavCategorySaveReq true "请求载荷"
// @Success     200   {object} ack.NavCategorySaveAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/category [post]
func (*navCtrl) SaveCategory(c *gin.Context) {
	service.Nav.SaveCategory(c)
}

// DeleteCategory 删除分类
// @Tags        nav - 导航模块
// @Summary     apis/nav/category - 删除分类
// @Description 删除分类
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param parameter query req.NavCategoryDeleteReq true "查询参数"
// @Success     200   {object} ack.NavCategoryDeleteAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/category [delete]
func (*navCtrl) DeleteCategory(c *gin.Context) {
	service.Nav.DeleteCategory(c)
}

// SaveSite 保存站点
// @Tags        nav - 导航模块
// @Summary     apis/nav/site - 保存站点
// @Description 保存站点
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload body req.NavSiteSaveReq true "请求载荷"
// @Success     200   {object} ack.NavSiteSaveAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/site [post]
func (*navCtrl) SaveSite(c *gin.Context) {
	service.Nav.SaveSite(c)
}

// DeleteSite 删除站点
// @Tags        nav - 导航模块
// @Summary     apis/nav/site - 删除站点
// @Description 删除站点
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param parameter query req.NavSiteDeleteReq true "查询参数"
// @Success     200   {object} ack.NavSiteDeleteAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/site [delete]
func (*navCtrl) DeleteSite(c *gin.Context) {
	service.Nav.DeleteSite(c)
}

// SortCategories 保存分类排序
// @Tags        nav - 导航模块
// @Summary     apis/nav/category/sorted - 保存分类排序
// @Description 保存分类排序
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload body req.NavCategorySortedReq true "请求载荷"
// @Success     200   {object} ack.NavCategorySortedAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/category/sorted [post]
func (*navCtrl) SortCategories(c *gin.Context) {
	service.Nav.SortCategories(c)
}

// SortSites 保存站点排序
// @Tags        nav - 导航模块
// @Summary     apis/nav/site/sorted - 保存站点排序
// @Description 保存站点排序
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload body req.NavSiteSortedReq true "请求载荷"
// @Success     200   {object} ack.NavSiteSortedAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/nav/site/sorted [post]
func (*navCtrl) SortSites(c *gin.Context) {
	service.Nav.SortSites(c)
}
