package middleware

import "github.com/gin-gonic/gin"

//用于注册管理中间件
//没用用到了，直接使用jwt认证了
func Init(r *gin.Engine)  {
	//Csrfmiddleware(r)
	//Corsmiddleware(r)
	//Jwtmiddleware(r)

}
