// ----------- api doc definition -----------------

// @title          admin - <nav-server>
// @version        1.0
// @description    This is an API documentation of nav-server module.
// @termsOfService https://nav.go-sail.dev

// @contact.name  Go-Sail
// @contact.url   https://nav.go-sail.dev
// @contact.email support@go-sail.dev

// @Scheme
// @Host
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                Access Token protects our entity endpoints

// ----------- api doc definition -----------------

package admin

import (
	"fmt"
	"github.com/keepchen/go-sail/v3/http/api"
	"github.com/keepchen/go-sail/v3/sail"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/routes"
	"nav-server/pkg/constants"
	"nav-server/pkg/models"
)

// StartServer 启动服务
func StartServer() {
	afterFunc := func() {
		//注册错误码
		fmt.Println("注册错误码...")
		constants.RegisterErrorCode()
		fmt.Println("注册错误码[√]")
		//自动同步表结构
		fmt.Println("自动同步表结构和数据...")
		if err := models.AutoMigrate(sail.GetDBW()); err != nil {
			fmt.Println("自动同步表结构和数据出错：", err.Error())
		} else {
			fmt.Println("自动同步表结构和数据[√]")
		}
	}
	//设置响应器行为
	opts := &api.Option{
		ErrNoneCode:      constants.ErrNone,
		ErrNoneCodeMsg:   constants.ErrNone.String(),
		ForceHttpCode200: true,
	}

	sail.WakeupHttp(config.Get().AppName, &config.Get().SailConf).
		SetupApiOption(opts).
		Hook(routes.RegisterRoutes, nil, afterFunc).
		Launch()
}
