package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleService struct{}

func (a ArticleService) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}

func (a ArticleService) Get(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}

func (a ArticleService) Detele(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}

func (a ArticleService) List(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加用户")
}
