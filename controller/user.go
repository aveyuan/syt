package controller

import (
	"fmt"
	"github.com/aveyuan/syt/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

//用户注册信息
func UserRgeist(c *gin.Context) {
	var user models.VliUser
	if err := c.ShouldBindJSON(&user); err != nil {
		ResJson(402, "用户信息有误", c)
	} else {
		if user.Password != user.RePassword {
			ResJson(402, "两次密码不一致", c)
		} else {
			if err := user.Reg(); err != nil {
				ResJson(200, "用户注册失败", c)
			} else {
				//写入session
				ResJson(402, "用户注册成功", c)
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
func UserHome(c *gin.Context) {
	cl := JwtUser(c)

	tikets, err := cl.UserTickets(0, "")
	if err != nil {
		ResJson(202, "用户工单获取失败", c)
		c.Abort()
	}

	ticktesing, err := cl.UserTickets(8, "")
	if err != nil {
		ResJson(202, "没有正在进行的工单", c)
		c.Abort()
	}

	message := struct {
		UserID    uint
		NickName  string
		Tikets    int
		Tiketsing int
	}{cl.ID, cl.Nickname, len(tikets), len(ticktesing)}
	c.JSON(200, message)

}

//某个工单的详细信息
func UserTicket(c *gin.Context) {
	tkidp := c.Param("id")
	tkid, _ := strconv.Atoi(tkidp) //需要更新的工单地址

	v := &models.Ticket{}
	v.ID = uint(tkid)

	//定义一个map用来获取里面的数据
	ticket := make(map[string]interface{})
	detail := v.Detail()
	ticket["ID"] = detail.Ticket.ID
	ticket["User"] = detail.User.Nickname
	ticket["Username"] = detail.User.Username
	ticket["Title"] = detail.Ticket.Title
	ticket["Tksource"] = detail.Tksource.Content
	ticket["Satisfaction"] = detail.Satisfaction.Content
	ticket["Status"] = detail.Ticket.Status
	ticket["CreateAt"] = detail.Ticket.CreatedAt
	ticket["UpdateAt"] = detail.Ticket.UpdatedAt
	soveuser := make([]map[string]interface{}, len(detail.Solveuser))
	for k, v := range detail.Solveuser {
		suser := make(map[string]interface{})
		suser["id"] = v.ID
		suser["username"] = v.Username
		suser["nickname"] = v.Nickname
		soveuser[k] = suser
	}
	ticket["solveuser"] = soveuser
	c.JSON(200, ticket)
}

//关闭某个工单
func UserTicketClose(c *gin.Context) {
	tkidp := c.Param("id")
	tkid, _ := strconv.Atoi(tkidp) //需要更新的工单地址
	v := &models.Ticket{}
	v.ID = uint(tkid)
	user := JwtUser(c)
	if err := v.UserTicketClose(user.Username); err != nil {
		fmt.Println("返回内容", err)
		ResJson(202, "关闭工单失败,请检查权限", c)
	} else {
		ResJson(202, "关闭工单成功", c)
	}

}

//回复工单
func UserTicketRe(c *gin.Context) {
	tkidp := c.Param("id")
	tkid, _ := strconv.Atoi(tkidp) //需要更新的工单地址

	print(tkid)
}

//用户详细信息
func UserInfo(c *gin.Context) {
	cl := JwtUser(c)
	c.JSON(200, cl)
}

//修改用户信息
func UserUpdate(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ResJson(202, "参数错误", c)
		c.Abort()
	}
	if err := user.Update(); err != nil {
		ResJson(202, "用户信息修改失败", c)
		c.Abort()
	} else {
		ResJson(202, "用户信息修改成功", c)
	}

}

//修改用户信息
func UserChpass(c *gin.Context) {
	var pass models.Password
	if err := c.ShouldBindJSON(&pass); err != nil {
		ResJson(202, "参数错误", c)
		c.Abort()
	}
	user := JwtUser(c)
	pass.Username = user.Username
	if err := pass.Update(); err != nil {
		ResJson(202, "密码修改失败", c)
		c.Abort()
	} else {
		ResJson(202, "密码修改成功", c)
	}

}
