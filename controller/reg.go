package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/aveyuan/syt/models"
)


//注册信息
func ClientRegPost(c *gin.Context)  {
	var user models.VliUser
	if err := c.ShouldBindJSON(&user);err!=nil{
		ResJson(402,"用户信息有误",c)
	}else{
		if user.Password !=user.RePassword{
			ResJson(402,"两次密码不一致",c)
		}else {
			if err:= user.Reg();err !=nil{
				ResJson(200,"用户注册失败",c)
			}else {
				//写入session
				ResJson(402,"用户注册成功",c)
			}
		}

	}
}
