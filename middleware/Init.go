package middleware

import "github.com/gin-gonic/gin"

//用于注册管理中间件
func Init(r *gin.Engine)  {
	Csrfmiddleware(r)
	Corsmiddleware(r)
}
