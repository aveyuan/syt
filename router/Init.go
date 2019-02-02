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

	r.POST("/reg",controller.ClientRegPost)


	r.NoRoute(func(c *gin.Context) {
		c.JSON(404,gin.H{"message":"404 Not founds"})
	})

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/home",controller.Home)
		auth.GET("/detail",controller.UserTickets)
		auth.GET("/listtk",controller.ListTickets)
	}


	return r
}
