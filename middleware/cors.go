package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Corsmiddleware(r *gin.Engine)   {
	r.Use(cors.Default())
}
