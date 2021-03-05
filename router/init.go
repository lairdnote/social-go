package router

import (
	"github.com/lairdnote/social/api"
	"github.com/lairdnote/social/middleware/auth"
	v1 "github.com/lairdnote/social/router/v1"

	"github.com/gin-gonic/gin"
)

// 初始化路由配置
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(auth.Cors())
	// 主页
	r.GET("/", api.Index)

	// api服务
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", api.Index)

		// 添加v1路由
		v1.AddV1Group(apiGroup)
	}

	return r
}
