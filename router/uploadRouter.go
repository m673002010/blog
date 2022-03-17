package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func UploadRouterInit(r *gin.Engine) {
	fileRouter := r.Group("/file")

	fileRouter.POST("/uploadPic", controller.File{}.UploadPic)

	fileRouter.POST("/upload", controller.File{}.Upload)

	fileRouter.POST("/MultiUpload", controller.File{}.MultiUpload)
}
