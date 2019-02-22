package router

import (
	"github.com/gin-gonic/gin"
	"github.com/aveyuan/syt/controller"
	"github.com/aveyuan/syt/middleware"
)

func Init()*gin.Engine  {
	r := gin.Default()
	//注册中间件
	middleware.Init(r)
	//单独注册jwt中间件
	authMiddleware:=middleware.Jwtmiddleware(r)
	//把r给controller
	controller.R = r

	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})

	r.NoRoute(func(c *gin.Context) {
		controller.ResJson(404,"Page Not Found",c)
	})

	r.POST("/login", authMiddleware.LoginHandler)

	//用户路由组
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/reg",controller.UserRgeist)
		user.GET("/home",controller.UserHome)
		user.GET("/tickets",controller.UserTickets)

	}
	//工单路由组
	ticket:=r.Group("/ticket")
	ticket.Use(authMiddleware.MiddlewareFunc())
	{
		ticket.GET("/listtk",controller.ListTickets)
		ticket.POST("/createtk",controller.CreateTicket)
		ticket.POST("/updatetk/:id",controller.UpdateTicket)
	}

	return r
}
