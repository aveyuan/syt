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
	//注册静态文件
	r.Static("/static", ("static/"))
	//把r给controller
	controller.R = r

	r.LoadHTMLGlob("views/**/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})

	r.GET("/login",controller.ClientLoginGet)
	r.POST("/login",controller.ClientLoginPost)
	r.GET("/admin/login",controller.UserLoginGet)
	r.POST("/admin/login",controller.UserLoginPost)

	r.POST("/home",controller.Home)


	return r
}
