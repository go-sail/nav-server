package handler

import (
	"github.com/gin-gonic/gin"
	"nav-server/app/admin/http/service"
)

type indexCtrl struct{}

var Index = &indexCtrl{}

// List 获取完整的导航列表
// @Tags        index - 首页模块
// @Summary     apis/index/list - 获取完整的导航列表
// @Description 获取完整的导航列表
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string false "授权token"
// @Param parameter query req.NavListReq true "查询参数"
// @Success     200   {object} ack.NavListAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/index/list [get]
func (*indexCtrl) List(c *gin.Context) {
	service.Index.List(c)
}
