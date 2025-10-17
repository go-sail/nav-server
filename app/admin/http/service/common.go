package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keepchen/go-sail/v3/sail"
	"go.uber.org/zap"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/api/ack"
	"nav-server/app/admin/http/api/req"
	"nav-server/pkg/constants"
	"nav-server/pkg/utils"
	"net/url"
	"time"
)

type commonSvc struct{}

var Common = &commonSvc{}

func (*commonSvc) Upload(c *gin.Context) {
	var (
		form      req.CommonUploadReq
		resp      ack.CommonUploadAck
		_, cancel = context.WithTimeout(context.Background(), time.Minute*10)
		loggerSvc = sail.LogTrace(c).GetLogger()
	)
	defer cancel()
	if err := c.ShouldBind(&form); err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, nil, err.Error()).Send()
		return
	}
	if code, err := form.Validator(); err != nil {
		sail.Response(c).Wrap(code, nil, err.Error()).Send()
		return
	}

	file, _ := c.FormFile("file")
	dst := fmt.Sprintf("%s/%s/%s", config.Get().Nav.IconPath, form.Action, file.Filename)
	err := sail.Utils().File().Save2Dst(file, dst)
	if err != nil {
		loggerSvc.Error("上传文件时保存文件出错", zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, nil, err.Error()).Send()
		return
	}

	//地址需要与注册的图标路由一致
	resp.URL = fmt.Sprintf("%s/icons/%s/%s", config.Get().Nav.IconEndpoint, form.Action, file.Filename)
	sail.Response(c).Data(resp)
}

func (*commonSvc) SyncRemoteAsset(c *gin.Context) {
	var (
		form        req.CommonSyncRemoteAssetReq
		resp        ack.CommonSyncRemoteAssetAck
		ctx, cancel = context.WithTimeout(context.Background(), time.Minute*10)
		loggerSvc   = sail.LogTrace(c).GetLogger()
	)
	defer cancel()
	if err := c.ShouldBind(&form); err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, nil, err.Error()).Send()
		return
	}
	if code, err := form.Validator(); err != nil {
		sail.Response(c).Wrap(code, nil, err.Error()).Send()
		return
	}

	urlObj, err := url.Parse(form.URL)
	if err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, resp, err.Error()).Send()
		return
	}

	host := urlObj.Host
	result, err := FetchRemoteAsset(ctx, form.URL, c, time.Second*8)
	if err != nil {
		loggerSvc.Error("FetchRemoteAsset error", zap.String("url", form.URL), zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrFetchRemoteAssetFailed, resp, err.Error()).Send()
		return
	}
	//断言资源扩展名
	ext := assertAssetExtension(result)
	if len(ext) == 0 {
		ext = ".ico"
	}
	filename := fmt.Sprintf("%s%s", host, ext) //默认文件名设置为域名+扩展名，例如：github.com.ico
	if form.Action == "category" {             //如果是分类，则文件名使用uuid
		filename = fmt.Sprintf("%s%s", utils.MakeIdentity(), ext)
	}

	dst := fmt.Sprintf("%s/%s/%s", config.Get().Nav.IconPath, form.Action, filename)
	err = sail.Utils().File().PutContents(result, dst)
	if err != nil {
		loggerSvc.Error("同步远端资源时保存文件出错", zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
		return
	}

	//地址需要与注册的图标路由一致
	resp.URL = fmt.Sprintf("%s/icons/%s/%s", config.Get().Nav.IconEndpoint, form.Action, filename)
	sail.Response(c).Data(resp)
}
