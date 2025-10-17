package service

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/keepchen/go-sail/v3/sail"
	"nav-server/app/admin/config"
	"nav-server/pkg/constants"
)

var tokenTTL = time.Hour * 24 * 3 //令牌有效期

// 颁发令牌
func issueToken(uid, username string) (string, error) {
	fields := map[string]interface{}{
		"username": username,
		"iat":      time.Now().Unix(),
	}
	exp := time.Now().Add(tokenTTL).Unix() //令牌3天后过期
	return sail.JWT().MakeToken(uid, exp, fields)
}

// ForceExpireToken 强制失效某时间点之前颁发的token
func ForceExpireToken(ctx context.Context, uid string, pointAtSeconds int64) error {
	redisKey := sail.Utils().String().WrapRedisKey(config.Get().AppName, fmt.Sprintf(constants.RedisKeyUserLogoutAt, uid))
	_, err := sail.GetRedis().Set(ctx, redisKey, pointAtSeconds, tokenTTL).Result()

	return err
}

// FetchRemoteAsset 拉取远端资源
func FetchRemoteAsset(ctx context.Context, url string, ginContext *gin.Context, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	req.Header.Set("User-Agent", ginContext.Request.UserAgent()) //透传客户端信息
	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET failed: %w", err)
	}
	if resp.StatusCode >= 400 {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("server returned %d", resp.StatusCode)
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.LastIndex(strings.ToLower(contentType), "image/") != 0 {
		return nil, fmt.Errorf("content type is not image/%s", contentType)
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return byt, nil
}

// 断言资源扩展名
func assertAssetExtension(b []byte) string {
	mType := mimetype.Detect(b)
	if mType != nil {
		return mType.Extension()
	}
	return ""
}

// 删除本地图标
func removeLocalIcon(url string) {
	if len(url) == 0 {
		return
	}
	//是引用的远端图标，不操作
	if strings.LastIndex(strings.ToLower(url), config.Get().Nav.IconEndpoint) != 0 {
		return
	}
	//移除域名及icons引用前缀
	iconFile := strings.Replace(url, fmt.Sprintf("%s/icons/", config.Get().Nav.IconEndpoint), "", 1)
	//拼接为真实的目录路径
	targetPath := fmt.Sprintf("%s/%s", config.Get().Nav.IconPath, iconFile)
	//执行删除操作
	fmt.Println("删除本地icon文件：", targetPath, os.Remove(targetPath))
}
