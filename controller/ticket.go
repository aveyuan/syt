package controller

import (
	"github.com/aveyuan/syt/models"
	"github.com/gin-gonic/gin"
	"fmt"
)

//工单控制器

//所有工单
func ListTickets(c *gin.Context)  {
	ticktes := &models.Ticket{}
	tickeslist,err := ticktes.List()
	if err !=nil{
		ResJson(402,"获取信息有误",c)
	}
	//定义一个map用来获取里面的数据
	tickmap := make([]map[string]interface{},len(tickeslist))
	for k,v := range tickeslist{
		//组合工单内容
		ticket := make(map[string]interface{})
		detail := v.Detail()
		ticket["ID"]=v.ID
		ticket["User"]=detail.User.Nickname
		ticket["Username"]=detail.User.Username
		ticket["Title"]=v.Title
		ticket["Tksource"]=detail.Tksource.Content
		ticket["Satisfaction"]=detail.Satisfaction.Content
		ticket["Status"]=v.Status
		ticket["CreateAt"]=v.CreatedAt
		ticket["UpdateAt"]=v.UpdatedAt
		soveuser := make([]map[string]interface{},len(detail.Solveuser))
		for k,v := range detail.Solveuser{
			suser := make(map[string]interface{})
			suser["username"]=v.Username
			suser["nickname"]=v.Nickname
			soveuser[k]=suser
		}
		ticket["solveuser"]=soveuser
		tickmap[k]=ticket
	}
	c.JSON(200,tickmap)
}

//用户的工单
func UserTickets(c *gin.Context)  {
	user:=JwtUser(c)
	tickeslist,err := user.UserTickets()
	if err !=nil{
		ResJson(402,"获取用户工单失败",c)
	}
	//定义一个map用来获取里面的数据
	tickmap := make([]map[string]interface{},len(tickeslist))
	for k,v := range tickeslist{
		//组合工单内容
		ticket := make(map[string]interface{})
		detail := v.Detail()
		ticket["ID"]=v.ID
		ticket["User"]=detail.User.Nickname
		ticket["Username"]=detail.User.Username
		ticket["Title"]=v.Title
		ticket["Tksource"]=detail.Tksource.Content
		ticket["Satisfaction"]=detail.Satisfaction.Content
		ticket["Status"]=v.Status
		ticket["CreateAt"]=v.CreatedAt
		ticket["UpdateAt"]=v.UpdatedAt
		soveuser := make([]map[string]interface{},len(detail.Solveuser))
		for k,v := range detail.Solveuser{
			suser := make(map[string]interface{})
			suser["username"]=v.Username
			suser["nickname"]=v.Nickname
			soveuser[k]=suser
		}
		ticket["solveuser"]=soveuser
		tickmap[k]=ticket
	}
	c.JSON(200,tickmap)
}

//创建工单
func CreateTicket(c *gin.Context)  {
	user:=JwtUser(c)
	var tkcr models.TkCreate
	if err := c.ShouldBindJSON(&tkcr);err!=nil{
		ResJson(402,"创建工单参数有误",c)
	}else{
		tkcr.User=user
		if err :=tkcr.Add();err !=nil{
			ResJson(402,"创建工单失败",c)
		}else {
			ResJson(402,"创建工单成功",c)
		}
	}
}

//更新/分配工单
func SaveTicket(c *gin.Context)  {
	var tksave models.TkSave
	if err := c.ShouldBindJSON(&tksave);err!=nil{
		fmt.Println(err)
		ResJson(402,"创建工单参数有误",c)
	}else {
		var user []models.User
		thisusers := tksave.Solveuser
		for _,v := range thisusers{
			u,_ := v.Detail()
			user = append(user,*u)
		}
		fmt.Println(user)
		if err := tksave.Update();err !=nil{
			ResJson(402,"更新工单失败",c)
		}else {
			ResJson(402,"更新工单成功",c)
		}
	}
}