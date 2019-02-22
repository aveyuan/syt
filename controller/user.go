package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/aveyuan/syt/models"
	"github.com/appleboy/gin-jwt"
)


//用户注册信息
func UserRgeist(c *gin.Context)  {
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

//用户目录
//可以修改自己的资料
//可以看到自己提起的工单
//正在处理的工单
//跟进反馈工单
//创建工单
//满意度评判
func UserHome(c *gin.Context)  {
	claims := jwt.ExtractClaims(c)
	username := claims["user"]
	if username==""{
		c.JSON(200,gin.H{"message":"没有用户信息"})
		c.Abort()
	}else {
		var user models.User
		user.Username=username.(string)
		cl,err := user.Detail()
		if err !=nil{
			ResJson(202,"用户信息获取失败",c)
			c.Abort()
		}

		tikets,err := cl.UserTickets(0,"")
		if err !=nil{
			ResJson(202,"用户工单获取失败",c)
			c.Abort()
		}

		ticktesing,err := cl.UserTickets(8,"")
		if err !=nil{
			ResJson(202,"没有正在进行的工单",c)
			c.Abort()
		}

		message := struct {
			UserID uint
			NickName string
			Tikets int
			Tiketsing int
		}{cl.ID,cl.Nickname,len(tikets),len(ticktesing)}
		c.JSON(200,message)

	}
}