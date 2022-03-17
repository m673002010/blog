package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct{}

func (c Category) List(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户列表")
}

func (c Category) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}

func (c Category) Info(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户信息")
}
