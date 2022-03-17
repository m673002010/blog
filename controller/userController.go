package controller

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u User) Login(ctx *gin.Context) {
	// 用户名、密码
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := model.User{}
	user.Username = username
	user.Password = password

	result := model.DB.Take(&user)

	// 用户查询成功
	if result.RowsAffected > 0 {
		fmt.Println(user)
		tokenString, err := middleware.JWTClaims{}.GetToken(user.Id, user.Username, user.Account)
		fmt.Println(err)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "登录失败", "data": gin.H{"token": tokenString}})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功", "data": gin.H{"token": tokenString}})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "帐密不符"})
	}
}

func (u User) List(ctx *gin.Context) {
	// 查询数据库，将查询到的数据存入切片
	userList := []model.User{}

	model.DB.Find(&userList)

	ctx.JSON(http.StatusOK, gin.H{
		"userList": userList,
	})
}

func (u User) Add(ctx *gin.Context) {
	user := model.User{
		Username: "zhangsan",
		Account:  "13128235@qq.com",
		Password: "123456",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	model.DB.Create(&user)

	ctx.String(http.StatusOK, "添加用户成功")
}

func (u User) Edit(ctx *gin.Context) {
	// user := model.User{Id: 1}
	// model.DB.Find(&user)
	// user.Username = "admin"
	// user.UpdateAt = time.Now()
	// model.DB.Save(user)

	// user := model.User{}
	// model.DB.Model(&user).Where("id = ?", 1).Update("username", "哈哈")

	user := model.User{}
	model.DB.Where("id = ?", 1).Find(&user)
	user.Username = "testEdit"
	model.DB.Save(user)

	ctx.String(http.StatusOK, "编辑用户信息成功")
}

func (u User) Delete(ctx *gin.Context) {
	user := model.User{Id: 4}
	// model.DB.Delete(&user)
	model.DB.Delete(user)
	ctx.String(http.StatusOK, "删除用户成功")
}

func (u User) Info(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户信息")
}
