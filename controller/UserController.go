package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"m0nk3y/pocScanner/common"
	"m0nk3y/pocScanner/model"
	"m0nk3y/pocScanner/util"
	"net/http"
)

var cipherKey = "asdfgsImFSNXUATM"

func Register(ctx *gin.Context) {
	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 过滤判断
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "密码长度必须不小于六位"})
		return
	}
	if isTelephoneExist(db, telephone) == true {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "不能重复注册"})
		return
	}
	if name == "" {
		name = util.RandomName(6)
	}
	// 全通过要对密码进行处理再保存
	password = util.EnAES(password, cipherKey)
	log.Println(name, password, telephone)
	newUser := User.User{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}
	db.Create(&newUser)

	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "密码长度不小于6位"})
	}

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "手机号必须为11位"})
		return
	}

	// 判断用户是否存在
	var user User.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		// 用户不存在 直接返回
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "用户名或密码错误"})
		return
	}
	// 判断密码是否正确
	password = util.EnAES(password, cipherKey)
	if user.Password != password {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "用户名或密码错误"})
		return
	}
	// 生成token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": "系统异常0x01"})
		log.Printf("token gengerate err %v", err)
		return
	}

	// 返回状态码
	ctx.JSON(200, gin.H{
		"token": token,
		"msg":   "登录成功",
	})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{"user": user},
	})
}
