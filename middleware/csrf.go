package middleware

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"github.com/gin-gonic/gin"
)

func Csrfmiddleware(r *gin.Engine)  {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context){
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}


