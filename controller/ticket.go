package controller

import (
	"github.com/aveyuan/syt/models"
	"github.com/gin-gonic/gin"
)

//工单控制器

//所有工单
func ListTickets(c *gin.Context)  {
	ticktes := &models.Ticket{}
	tickeslist,err := ticktes.List()
	if err !=nil{
		ResJson(402,"获取信息有误",c)
	}
	c.JSON(200,tickeslist)
}

//用户的工单
func UserTickets(c *gin.Context)  {
	user:=JwtUser(c)
	tickets,err := user.UserTickets()
	if err!=nil{
		ResJson(402,"用户工单获取有误",c)
	}
	user.Tickets=tickets
	c.JSON(200,user)
}

//创建工单
func CreateTicket(c *gin.Context)  {

}