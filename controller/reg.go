package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/utrack/gin-csrf"
	"github.com/aveyuan/syt/models"
)

//用户自主注册
func ClientRegGet(c *gin.Context)  {
	c.HTML(http.StatusOK,"client/reg.html",gin.H{"token":csrf.GetToken(c),"title":"用户注册页面"})
}

//注册信息
func ClientRegPost(c *gin.Context)  {
	var client models.Client
	repassword := c.PostForm("repassword")
	if err := c.ShouldBind(&client);err!=nil{
		c.JSON(200,gin.H{"message":"用户信息有误"})
	}else{
		if client.Password !=repassword{
			c.JSON(200,gin.H{"message":"两次密码验证不一致"})
		}else {
			if err:= client.Add();err !=nil{
				c.JSON(200,gin.H{"message":"用户注册失败"})
			}else {
				//写入session
				c.JSON(200,gin.H{"message":"用户注册成功"})
			}
		}

	}
}
