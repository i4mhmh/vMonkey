package main

import (
	"github.com/gin-gonic/gin"
	"m0nk3y/pocScanner/controller"
	"m0nk3y/pocScanner/middleware"
)

func CollectionRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/userinfo", middleware.AuthMiddleware(), controller.UserInfo)

	return r
}
