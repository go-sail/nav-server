package handler

import (
	"github.com/gin-gonic/gin"
	"nav-server/app/admin/http/service"
)

type commonCtrl struct{}

var Common = &commonCtrl{}

// Upload 上传文件
// @Tags        common - 通用模块
// @Summary     apis/common/upload - 上传文件
// @Description 上传文件
// @Security    ApiKeyAuth
// @Accept      multipart/form-data
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param file formData file true "文件"
// @Param payload formData req.CommonUploadReq true "请求载荷"
// @Success     200   {object} ack.CommonUploadAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/common/upload [put]
func (*commonCtrl) Upload(c *gin.Context) {
	service.Common.Upload(c)
}

// SyncRemoteAsset 同步远端资源
// @Tags        common - 通用模块
// @Summary     apis/common/sync-remote-asset - 同步远端资源
// @Description 同步远端资源
// @Security    ApiKeyAuth
// @Accept      application/json
// @Produce     json
// @Param Authorization header string true "授权token"
// @Param payload formData req.CommonSyncRemoteAssetReq true "请求载荷"
// @Success     200   {object} ack.CommonSyncRemoteAssetAck
// @Failure     400   {object} dto.Error400
// @Failure     500   {object} dto.Error500
// @Router      /apis/common/sync-remote-asset [post]
func (*commonCtrl) SyncRemoteAsset(c *gin.Context) {
	service.Common.SyncRemoteAsset(c)
}
