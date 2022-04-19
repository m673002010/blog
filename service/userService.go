package service

import (
	"blog/middleware"
	"blog/model"
	"blog/param"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

func (UserService) Login(uParam param.LoginParam) gin.H {

	user := model.User{}
	model.DB.Where(model.User{Account: uParam.Account, Password: uParam.Password}).Find(&user)

	if user.Username != "" {
		tokenString, err := middleware.JWTClaims{}.GetToken(user.Id, user.Username, user.Account)

		if err == nil {
			return gin.H{
				"code":    0,
				"message": "登录成功",
				"data": gin.H{
					"token":    tokenString,
					"username": user.Username,
					"account":  user.Account,
				},
			}
		}
	}

	return gin.H{
		"code":    -1,
		"message": "登录失败",
	}
}

func (UserService) FindUserById(id int) bool {

	return false
}

func (UserService) AddUser(uParam param.AddUserParam) bool {
	user := model.User{
		Username: uParam.Username,
		Account:  uParam.Account,
		Password: uParam.Password,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	err := model.DB.Create(&user)

	return err != nil
}

func (UserService) EditUser(uParam param.EditUserParam) bool {
	user := model.User{Id: uParam.Id}
	model.DB.Find(&user)
	user.Username = uParam.Username
	user.Account = uParam.Account
	err := model.DB.Save(user)

	return err != nil
}

func (UserService) DeteleUser(id int) bool {
	user := model.User{Id: id}
	err := model.DB.Delete(user)
	return err != nil
}

func (UserService) Info(id int) gin.H {
	return gin.H{}
}
