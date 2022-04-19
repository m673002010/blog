package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouter := r.Group("/user")

	userRouter.POST("/login", controller.UserController{}.Login)

	userRouter.POST("/register", controller.UserController{}.Register)

	userRouter.GET("/info", controller.UserController{}.Info)

	userRouter.POST("/edit", controller.UserController{}.Edit)

	userRouter.GET("/delete", controller.UserController{}.Delete)
}
