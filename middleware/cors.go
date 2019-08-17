package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Corsmiddleware(r *gin.Engine) {
	r.Use(cors.Default())
}
