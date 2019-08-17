package router

import (
	"github.com/aveyuan/syt/controller"
	"github.com/aveyuan/syt/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	//注册中间件
	middleware.Init(r)
	//单独注册jwt中间件
	authMiddleware := middleware.Jwtmiddleware(r)
	//把r给controller
	controller.R = r
	r.Static("/upload", "./upload")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.NoRoute(func(c *gin.Context) {
		controller.ResJson(404, "Page Not Found", c)
	})

	r.POST("/login", authMiddleware.LoginHandler)

	//用户路由组
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/reg", controller.UserRgeist)
		user.GET("/home", controller.UserHome)
		user.GET("/tickets", controller.UserTickets)        //用户所有的工单列表
		user.GET("/ticketsing", controller.UserTicketsing)  //用户进行中的工单列表
		user.GET("/ticket/:id", controller.UserTicket)      //用户工单详情
		user.PUT("/ticket/:id", controller.UserTicketClose) //关闭用户工单
		user.PUT("/ticket/:id/re", controller.UserTicketRe) //回复工单
		user.GET("/info", controller.UserInfo)              //用户详细信息
		user.PUT("/info", controller.UserUpdate)            //更新用户详细信息
		user.PUT("/chpass", controller.UserChpass)          //更新用户详细信息

	}
	//工单路由组
	ticket := r.Group("/ticket")
	ticket.Use(authMiddleware.MiddlewareFunc())
	{
		ticket.GET("/listtk", controller.ListTickets)
		ticket.POST("/createtk", controller.CreateTicket)
		ticket.POST("/updatetk/:id", controller.UpdateTicket)
		ticket.POST("/upload", controller.FileUpload)
	}

	return r
}
