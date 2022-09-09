package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ces", func(data *gin.Context) {
		name := data.PostForm("username")
		password := data.PostForm("password")

		// check
		if name == "a" && password == "123" {
		}
	})
	r.Run()
}
