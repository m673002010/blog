package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func CateRouterInit(r *gin.Engine) {
	cateRouter := r.Group("/cate")

	cateRouter.GET("/", controller.Category{}.Info)

	cateRouter.GET("/list", controller.Category{}.List)

	cateRouter.POST("/add", controller.Category{}.Add)
}
