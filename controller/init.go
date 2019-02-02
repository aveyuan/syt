package controller

import (
	"github.com/appleboy/gin-jwt"
	"github.com/aveyuan/syt/models"
	"github.com/gin-gonic/gin"
)

//定义一个全局的r用于跳转
var R *gin.Engine

//定义用于jwt识别的函数，返回用户信息
func JwtUser(c *gin.Context)models.User {
	claims := jwt.ExtractClaims(c)
	username := claims["user"]
	user := &models.User{Username:username.(string)}
	u,err := user.Detail()
	if err !=nil{
		ResJson(402,"用户识别失败",c)
		c.Abort()
	}
	return *u
}

//定义一个响应的json函数
func ResJson(code int,message string,c *gin.Context)  {
	c.JSON(code,gin.H{"message":message})
}