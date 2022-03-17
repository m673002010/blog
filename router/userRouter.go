package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouter := r.Group("/user")

	userRouter.POST("/login", controller.User{}.Login)

	userRouter.GET("/info", controller.User{}.Info)

	userRouter.GET("/list", controller.User{}.List)

	userRouter.GET("/add", controller.User{}.Add)

	userRouter.GET("/edit", controller.User{}.Edit)

	userRouter.GET("/delete", controller.User{}.Delete)
}
