package controller

import (
	"blog/model"
	"blog/param"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 登录
func (UserController) Login(ctx *gin.Context) {
	uParam := param.LoginParam{}
	err := ctx.Bind(&uParam)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "参数错误"})
		return
	}

	res := service.UserService{}.Login(uParam)

	ctx.JSON(http.StatusOK, res)
}

// 注册用户
func (UserController) Register(ctx *gin.Context) {
	uParam := param.AddUserParam{}
	err := ctx.Bind(&uParam)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "参数错误"})
		return
	}

	res := service.UserService{}.AddUser(uParam)
	if res {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "注册失败"})
}

// 编辑用户
func (UserController) Edit(ctx *gin.Context) {
	uParam := param.EditUserParam{}

	err := ctx.Bind(&uParam)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "参数错误"})
		return
	}

	res := service.UserService{}.EditUser(uParam)
	if res {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "更新失败"})
}

// 删除用户
func (UserController) Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "参数错误"})
		return
	}

	res := service.UserService{}.DeteleUser(userId)

	if res {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "删除失败"})
}

// 获取用户信息
func (UserController) Info(ctx *gin.Context) {
	id := ctx.Query("id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "参数错误"})
		return
	}

	userInfo := service.UserService{}.Info(userId)

	if userInfo == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户信息获取成功"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "用户信息获取失败"})
}

// 获取用户列表
func (UserController) List(ctx *gin.Context) {
	userList := []model.User{}

	model.DB.Find(&userList)

	ctx.JSON(http.StatusOK, gin.H{
		"userList": userList,
	})
}
