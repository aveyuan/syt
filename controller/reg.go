package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/aveyuan/syt/models"
)


//注册信息
func ClientRegPost(c *gin.Context)  {
	var user models.VliUser
	if err := c.ShouldBindJSON(&user);err!=nil{
		c.JSON(200,gin.H{"message":"用户信息有误"})
	}else{
		if user.Password !=user.RePassword{
			c.JSON(200,gin.H{"message":"注册失败","content":"两次密码不一致"})
		}else {
			if err:= user.Reg();err !=nil{
				c.JSON(200,gin.H{"message":"用户注册失败"})
			}else {
				//写入session
				c.JSON(200,gin.H{"message":"用户注册成功"})
			}
		}

	}
}
