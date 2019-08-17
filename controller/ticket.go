package controller

import (
	"github.com/aveyuan/syt/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

//工单控制器

//所有工单
//可以根据用户和工单的状态来查询，1个添加，2个条件,系统级权限
func ListTickets(c *gin.Context) {
	ticktes := &models.Ticket{}
	statusq := c.Query("status")
	useridq := c.Query("userid")
	status, _ := strconv.Atoi(statusq)
	userid, _ := strconv.Atoi(useridq)
	search := c.Query("search")
	var ticketslist []models.Ticket
	var err error
	if userid == 0 {
		ticketslist, err = ticktes.List(status, search)
	} else {
		user, err := models.IdUser(userid)
		if err != nil {
			ResJson(202, "用户信息错误", c)
			c.Abort()
		}
		ticketslist, err = user.UserTickets(status, search)
	}

	if err != nil {
		ResJson(402, "获取信息有误", c)
	}
	//定义一个map用来获取里面的数据
	tickmap := make([]map[string]interface{}, len(ticketslist))
	for k, v := range ticketslist {
		//组合工单内容
		ticket := make(map[string]interface{})
		detail := v.Detail()
		ticket["ID"] = v.ID
		ticket["User"] = detail.User.Nickname
		ticket["Username"] = detail.User.Username
		ticket["Title"] = v.Title
		ticket["Tksource"] = detail.Tksource.Content
		ticket["Satisfaction"] = detail.Satisfaction.Content
		ticket["Status"] = v.Status
		ticket["CreateAt"] = v.CreatedAt
		ticket["UpdateAt"] = v.UpdatedAt
		soveuser := make([]map[string]interface{}, len(detail.Solveuser))
		for k, v := range detail.Solveuser {
			suser := make(map[string]interface{})
			suser["id"] = v.ID
			suser["username"] = v.Username
			suser["nickname"] = v.Nickname
			soveuser[k] = suser
		}
		ticket["solveuser"] = soveuser
		tickmap[k] = ticket
	}
	c.JSON(200, tickmap)
}

//用户的工单
//只能看到用户自己的工单
func UserTickets(c *gin.Context) {
	user := JwtUser(c)
	search := c.Query("search")
	tickeslist, err := user.UserTickets(0, search)
	if err != nil {
		ResJson(402, "获取用户工单失败", c)
	}
	//定义一个map用来获取里面的数据
	tickmap := make([]map[string]interface{}, len(tickeslist))
	for k, v := range tickeslist {
		//组合工单内容
		ticket := make(map[string]interface{})
		detail := v.Detail()
		ticket["ID"] = v.ID
		ticket["User"] = detail.User.Nickname
		ticket["Username"] = detail.User.Username
		ticket["Title"] = v.Title
		ticket["Tksource"] = detail.Tksource.Content
		ticket["Satisfaction"] = detail.Satisfaction.Content
		ticket["Status"] = v.Status
		ticket["CreateAt"] = v.CreatedAt
		ticket["UpdateAt"] = v.UpdatedAt
		soveuser := make([]map[string]interface{}, len(detail.Solveuser))
		for k, v := range detail.Solveuser {
			suser := make(map[string]interface{})
			suser["id"] = v.ID
			suser["username"] = v.Username
			suser["nickname"] = v.Nickname
			soveuser[k] = suser
		}
		ticket["solveuser"] = soveuser
		tickmap[k] = ticket
	}
	c.JSON(200, tickmap)
}

//正在进行的工单
func UserTicketsing(c *gin.Context) {
	user := JwtUser(c)
	search := c.Query("search")
	tickeslist, err := user.UserTickets(8, search)
	if err != nil {
		ResJson(402, "获取用户工单失败", c)
	}
	//定义一个map用来获取里面的数据
	tickmap := make([]map[string]interface{}, len(tickeslist))
	for k, v := range tickeslist {
		//组合工单内容
		ticket := make(map[string]interface{})
		detail := v.Detail()
		ticket["ID"] = v.ID
		ticket["User"] = detail.User.Nickname
		ticket["Username"] = detail.User.Username
		ticket["Title"] = v.Title
		ticket["Tksource"] = detail.Tksource.Content
		ticket["Satisfaction"] = detail.Satisfaction.Content
		ticket["Status"] = v.Status
		ticket["CreateAt"] = v.CreatedAt
		ticket["UpdateAt"] = v.UpdatedAt
		soveuser := make([]map[string]interface{}, len(detail.Solveuser))
		for k, v := range detail.Solveuser {
			suser := make(map[string]interface{})
			suser["id"] = v.ID
			suser["username"] = v.Username
			suser["nickname"] = v.Nickname
			soveuser[k] = suser
		}
		ticket["solveuser"] = soveuser
		tickmap[k] = ticket
	}
	c.JSON(200, tickmap)
}

//创建工单
//这个是普通用户创建工单，请求者都不可以选的这种
func CreateTicket(c *gin.Context) {
	user := JwtUser(c)
	var tkcr models.TkCreate
	if err := c.ShouldBindJSON(&tkcr); err != nil {
		ResJson(402, "创建工单参数有误", c)
	} else {
		tkcr.User = user
		if err := tkcr.Add(); err != nil {
			ResJson(402, "创建工单失败", c)
		} else {
			ResJson(402, "创建工单成功", c)
		}
	}
}

//系统创建工单
//由管理员创建的工单，可以选择用户，工单来源，状态，等
//还没想好怎么写

//更新/分配工单
func UpdateTicket(c *gin.Context) {
	tkidp := c.Param("id")
	tkid, _ := strconv.Atoi(tkidp) //需要更新的工单地址
	var tksave models.TkSave       //绑定得到更新内容
	if err := c.ShouldBindJSON(&tksave); err != nil {
		ResJson(200, "更新工单参数有误", c)
	} else {

		var user []models.User //取得用户参数
		tksave.ID = tkid
		thisusers := tksave.Solveuser //将用户信息分解出来
		for _, v := range thisusers {
			u, _ := v.Detail()
			user = append(user, *u)
		}
		//更新工单
		tksave.Solveuser = user
		if err := tksave.Update(); err != nil {
			ResJson(402, "更新工单失败", c)
		} else {
			ResJson(402, "更新工单成功", c)
		}
	}
}
