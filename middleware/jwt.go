package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/appleboy/gin-jwt"
	"github.com/aveyuan/syt/models"
)

func Jwtmiddleware(r *gin.Engine)*jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "username", //用来加密的key,没什么其他用，不要被demo误导了
		//这个函数是将认证后的信息用来保存到生成的claims中,这里只保存了一个
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Jwtuser); ok {
				return jwt.MapClaims{
					//定义存储的内容
					"user" : v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		//标识处理，用于获取
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Jwtuser{
				Username: claims["user"].(string),
			}
		},
		//用于认证,返回的data用于生成key
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user models.VliUser
			if err := c.ShouldBindJSON(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			user.Lastip=c.ClientIP()
			if err := user.Valid();err ==nil {
				return &models.Jwtuser{Username:user.Username}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//授权信息，定义用户可以访问到哪些资源
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//if v, ok := data.(*models.Jwtuser); ok && v.Username == "ww" {
			if _, ok := data.(*models.Jwtuser); ok  {
				return true
			}

			return false
		},
		//未授权
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//验证信息
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
