package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func (a Article) List(ctx *gin.Context) {
	ctx.String(http.StatusOK, "文章列表")
}

func (a Article) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}
