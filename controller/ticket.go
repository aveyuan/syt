package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"github.com/aveyuan/syt/models"
)

//工单控制器

//所有工单
func ListTickets(c *gin.Context)  {
	ticktes := &models.Ticket{}
	tickeslist,err := ticktes.List()
	if err !=nil{
		c.JSON(402,gin.H{"message":"工单信息获取有误"})
	}
	c.JSON(200,tickeslist)
}

//用户的工单
func UserTickets(c *gin.Context)  {
	claims := jwt.ExtractClaims(c)
	username := claims["user"]
	user := &models.User{Username:username.(string)}
	//取得了用户的详细信息
	userdetail,err := user.Detail()
	if err !=nil{
		c.JSON(402,gin.H{"message":"用户信息获取失败"})
	}
	c.JSON(200,userdetail)
}