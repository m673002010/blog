package controller

import (
	"blog/lib"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type File struct{}

func (f File) UploadPic(ctx *gin.Context) {
	file, err := ctx.FormFile("pic")

	if err == nil {
		// 获取后缀名
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		if _, ok := allowExtMap[extName]; !ok {
			ctx.String(http.StatusOK, "文件类型不合法")
			return
		}

		// 创建保存目录
		day := lib.GetDay()
		dir := "./static/pic" + day

		if err := os.MkdirAll(dir, 0666); err != nil {
			ctx.String(http.StatusOK, "目录创建失败")
			return
		}

		// 生成文件名
		fileName := strconv.FormatInt(lib.GetUnix(), 10) + extName
		dst := path.Join(dir, fileName)

		ctx.SaveUploadedFile(file, dst)
	}

	ctx.JSON(http.StatusOK, "图片上传成功")
}

func (f File) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err == nil {
		dst := path.Join("./static/file", file.Filename)
		ctx.SaveUploadedFile(file, dst)
	}

	ctx.JSON(http.StatusOK, "上传成功")
}

func (f File) MultiUpload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()

	if err == nil {
		files := form.File["files[]"]

		for _, file := range files {
			dst := path.Join("./static/file", file.Filename)
			ctx.SaveUploadedFile(file, dst)
		}
	}

	ctx.JSON(http.StatusOK, "所有文件上传成功")
}
