package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/utrack/gin-csrf"
	"github.com/aveyuan/syt/models"
	"github.com/gin-contrib/sessions"
)

//前台登录
func ClientLoginGet(c *gin.Context)  {
	c.HTML(http.StatusOK,"client/login.html",gin.H{"token":csrf.GetToken(c)})
}

//验证账号密码
func ClientLoginPost(c *gin.Context)  {
	var client models.Client
	if err := c.ShouldBind(&client);err!=nil{
		c.JSON(200,gin.H{"message":"用户名密码必填"})
	}else{
		if err:= client.Valid();err !=nil{
			c.JSON(200,gin.H{"message":"客户登录失败"})
			}else {
				//写入session
				session := sessions.Default(c)
				session.Set("username",client.Username)
				//这里再跳转到不同的页面
				c.Request.URL.Path = "/home"
				R.HandleContext(c)
			}
		}
}

//后台登录
func UserLoginGet(c *gin.Context)  {
	c.HTML(http.StatusOK,"admin/login.html",gin.H{"token":csrf.GetToken(c)})
}

//验证账号密码
func UserLoginPost(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBind(&user);err!=nil{
		c.JSON(200,gin.H{"message":"用户名密码必填"})
	}else{
		if err:= user.Valid();err !=nil{
			c.JSON(200,gin.H{"message":"系统用户登录失败"})
		}else {
			//写入session
			session := sessions.Default(c)
			session.Set("username",user.Username)
			//这里再跳转到不同的页面
			c.Request.URL.Path = "/home"
			R.HandleContext(c)
		}
	}
}