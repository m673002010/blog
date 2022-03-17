package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 登录接口不验证
		path := ctx.Request.URL.Path
		if path == "/user/login" {
			ctx.Next()
			return
		}

		// 从请求头获取token字符串
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"code":    -1,
				"message": "无令牌",
			})
			ctx.Abort()
		}

		// 解析token字符串
		token, jwtClaims, err := JWTClaims{}.ParseToken(tokenString)

		// token验证失败
		if err != nil || !token.Valid {
			fmt.Println(err)
			ctx.JSON(401, gin.H{
				"code":    -1,
				"message": "令牌验证失败",
			})
			ctx.Abort()
		}

		// 将用户信息写入请求上下文
		userInfo := gin.H{
			"id":       jwtClaims.Id,
			"username": jwtClaims.Username,
			"account":  jwtClaims.Account,
		}

		ctx.Set("userInfo", userInfo)
	}
}
