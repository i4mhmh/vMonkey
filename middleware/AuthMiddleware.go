package middleware

import (
	"github.com/gin-gonic/gin"
	"m0nk3y/pocScanner/common"
	User "m0nk3y/pocScanner/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取header
		tokenString := ctx.GetHeader("Authorization")

		// token vaild
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 丢弃Bearer
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "权限不足",
			})
			return
		}

		// 通过验证 返回user信息
		userId := claims.UserId
		DB := common.GetDB()
		var user User.User
		DB.First(&user, userId)

		// 用户token生效期间已被删除
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "权限不足",
			})
		}

		// 用户存在
		ctx.Set("user", user)
		ctx.Next()
	}
}
