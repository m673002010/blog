package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func ArticleRouterInit(r *gin.Engine) {
	articleRouter := r.Group("/article")

	articleRouter.GET("/", controller.Article{}.List)

	articleRouter.GET("/list", controller.Article{}.List)

	articleRouter.POST("/add", controller.Article{}.Add)
}
