package router

import "github.com/gin-gonic/gin"

func Init()*gin.Engine  {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})

	return r
}
