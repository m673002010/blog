package main

import (
	"blog/middleware"
	"blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Verify())

	// 路由
	router.UserRouterInit(r)
	router.ArticleRouterInit(r)
	router.CateRouterInit(r)
	router.DefaultRouterInit(r)
	router.UploadRouterInit(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
