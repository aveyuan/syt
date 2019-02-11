package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Corsmiddleware(r *gin.Engine)   {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8848"}
	r.Use(cors.New(config))
}
