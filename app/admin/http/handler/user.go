package handler

import (
	"github.com/gin-gonic/gin"
	"nav-server/app/admin/http/service"
)

type userCtrl struct{}

var User = &userCtrl{}

// Login 用户登录
// @Tags        user - 用户模块
// @Summary     apis/user/login - 用户登录
// @Description 用户登录
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string false "授权token"
// @Param payload body req.UserLoginReq true "请求载荷"
// @Success     200   {object} ack.UserLoginAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/user/login [post]
func (*userCtrl) Login(c *gin.Context) {
	service.User.Login(c)
}

// Logout 用户退出登录
// @Tags        user - 用户模块
// @Summary     apis/user/logout - 用户退出登录
// @Description 用户退出登录
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload body req.UserLogoutReq true "请求载荷"
// @Success     200   {object} ack.UserLogoutAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/user/logout [post]
func (*userCtrl) Logout(c *gin.Context) {
	service.User.Logout(c)
}
