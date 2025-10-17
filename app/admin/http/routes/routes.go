package routes

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	sailMiddleware "github.com/keepchen/go-sail/v3/http/middleware"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/handler"
	"nav-server/app/admin/http/middleware"
	"net/http"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) {
	//全局开启跨域设置
	r.Use(sailMiddleware.WithCors(nil))

	apis := r.Group("/apis")
	{
		user := apis.Group("/user")
		{
			user.POST("/login", handler.User.Login)
			user.Use(middleware.AuthCheck()).
				POST("/logout", handler.User.Logout)
		}
		nav := apis.Group("/nav", middleware.AuthCheck())
		{
			nav.GET("/list", handler.Nav.List).
				GET("/categories", handler.Nav.Categories).
				POST("/category", handler.Nav.SaveCategory).
				DELETE("/category", handler.Nav.DeleteCategory).
				POST("/category/sorted", handler.Nav.SortCategories).
				GET("/sites", handler.Nav.Sites).
				POST("/site", handler.Nav.SaveSite).
				DELETE("/site", handler.Nav.DeleteSite).
				POST("/site/sorted", handler.Nav.SortSites)
		}
		common := apis.Group("/common")
		{
			common.PUT("/upload", handler.Common.Upload).
				POST("/sync-remote-asset", handler.Common.SyncRemoteAsset)
		}
	}

	index := apis.Group("/index").Use(gzip.Gzip(gzip.DefaultCompression))
	{
		index.GET("/list", handler.Index.List)
	}

	//注册静态图标地址
	r.Static("/icons", config.Get().Nav.IconPath)
	//注册404路由
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}
