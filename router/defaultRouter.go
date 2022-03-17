package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouter := r.Group("/")

	r.LoadHTMLGlob("template/*")

	defaultRouter.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "Main website",
		})
	})
}
