package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/keepchen/go-sail/v3/sail"
	"go.uber.org/zap"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/api/ack"
	"nav-server/app/admin/http/api/req"
	"nav-server/app/admin/http/middleware"
	"nav-server/pkg/constants"
	"nav-server/pkg/models"
	"time"
)

type userSvc struct{}

var User = &userSvc{}

func (*userSvc) Login(c *gin.Context) {
	var (
		form        req.UserLoginReq
		resp        ack.UserLoginAck
		loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	)
	defer cancel()
	if err := c.ShouldBind(&form); err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, resp, err.Error()).Send()
		return
	}
	if code, err := form.Validator(); err != nil {
		sail.Response(c).Wrap(code, resp, err.Error()).Send()
		return
	}

	var user models.User
	sail.GetDBR().WithContext(ctx).Where("username = ?", form.Username).First(&user)
	if user.ID == 0 {
		sail.Response(c).Wrap(constants.ErrUsernamePasswordNotMatch, resp).Send()
		return
	}
	passwordDecrypted, err := sail.JWT().Decrypt(user.Password)
	if err != nil {
		loggerSvc.Error("从数据表解析用户密码出错", zap.String("uid", user.UID), zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
		return
	}
	if passwordDecrypted != form.Password {
		sail.Response(c).Wrap(constants.ErrUsernamePasswordNotMatch, resp).Send()
		return
	}
	var logoutAt = time.Now().Add(time.Second * -1).Unix()
	token, err := issueToken(user.UID, user.Username)
	if err != nil {
		loggerSvc.Error("颁发令牌出错", zap.String("uid", user.UID), zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
		return
	}
	//如果不允许重复登录，需要将此前颁发的所有令牌标记为失效
	if !config.Get().Nav.RepeatLogin {
		_ = ForceExpireToken(ctx, user.UID, logoutAt)
	}
	//记录最近一次登录时间
	go func() {
		now := time.Now()
		sail.GetDBW().Where("id = ?", user.ID).Updates(models.User{
			LatestLogin: &now,
		})
	}()

	resp.Token = token
	sail.Response(c).Data(resp)
}

func (*userSvc) Logout(c *gin.Context) {
	var (
		form            req.UserLogoutReq
		resp            ack.UserLogoutAck
		loggerSvc       = sail.LogTrace(c).GetLogger()
		userCredentials = c.MustGet("userCredentials").(middleware.UserCredentials)
		ctx, cancel     = context.WithTimeout(context.Background(), time.Second*5)
	)
	defer cancel()
	if err := c.ShouldBind(&form); err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, resp, err.Error()).Send()
		return
	}
	if code, err := form.Validator(); err != nil {
		sail.Response(c).Wrap(code, resp, err.Error()).Send()
		return
	}

	err := ForceExpireToken(ctx, userCredentials.UID, time.Now().Unix())
	if err != nil {
		loggerSvc.Error("用户退出登录时，标记退出状态出错", zap.String("uid", userCredentials.UID), zap.String("err", err.Error()))
	}

	sail.Response(c).Success()
}
