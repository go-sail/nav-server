package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/keepchen/go-sail/v3/sail"
	"go.uber.org/zap"
	"nav-server/app/admin/config"
	"nav-server/pkg/constants"
	"time"
)

// UserCredentials 用户凭证
type UserCredentials struct {
	UID      string //用户唯一标识
	Username string //用户名
	IssueAt  int64  //令牌颁发时间
}

// AuthCheck 授权验证
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			sail.Response(c).Wrap(constants.ErrAuthorizationTokenInvalid, nil).Send()
			return
		}
		ok, claims, err := sail.JWT().ValidToken(token)

		if err != nil {
			sail.LogTrace(c).Error("验证访问令牌出错", zap.String("err", err.Error()))
			sail.Response(c).Wrap(constants.ErrAuthorizationTokenInvalid, nil).Send()
			return
		}
		if !ok {
			sail.Response(c).Wrap(constants.ErrAuthorizationTokenInvalid, nil).Send()
			return
		}

		var userCredentials UserCredentials
		if val, ok := claims["uid"]; ok {
			userCredentials.UID = val.(string)
		}
		if val, ok := claims["username"]; ok {
			userCredentials.Username = val.(string)
		}
		if val, ok := claims["iat"]; ok {
			userCredentials.IssueAt = int64(val.(float64))
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		ok, err = assertTokenExpiration(ctx, userCredentials.UID, userCredentials.IssueAt)
		if err != nil {
			sail.LogTrace(c).Error("断言令牌有效性出错", zap.String("err", err.Error()))
			sail.Response(c).Wrap(constants.ErrAuthorizationTokenInvalid, nil).Send()
			return
		}
		if !ok {
			sail.Response(c).Wrap(constants.ErrAuthorizationTokenInvalid, nil).Send()
			return
		}

		c.Set("userCredentials", userCredentials)
		c.Next()
	}
}

// 断言令牌是否有效(是否被ban)
func assertTokenExpiration(ctx context.Context, uid string, issueAt int64) (bool, error) {
	redisKey := sail.Utils().String().WrapRedisKey(config.Get().AppName, fmt.Sprintf(constants.RedisKeyUserLogoutAt, uid))
	logoutAt, err := sail.GetRedis().Get(ctx, redisKey).Int64()
	if errors.Is(err, redis.Nil) {
		return true, nil
	}
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}
	//在这一刻之前颁发的令牌统统失效
	if logoutAt != 0 && issueAt <= logoutAt {
		return false, nil
	}

	return true, nil
}
