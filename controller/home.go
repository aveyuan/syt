package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"github.com/aveyuan/syt/models"
	"fmt"
)

//客户主目录
//可以修改自己的资料
//可以看到自己提起的工单
//正在处理的工单
//跟进反馈工单
//创建工单
//满意度评判
func Home(c *gin.Context)  {
	claims := jwt.ExtractClaims(c)
	username := claims["user"]
	if username==""{
		c.JSON(200,gin.H{"message":"没有用户信息"})
		c.Abort()
	}else {
		var user models.User
		user.Username=username.(string)
		cl,err := user.Detail()
		if err !=nil{
			fmt.Println("用户信息获取有误")
			c.Abort()
		}else {
			fmt.Println("用户电话:",cl.Phone1)
			for _,v := range cl.Tickets{
				fmt.Println("工单ID",v.ID,"工单创建时间",v.CreatedAt)
				tb := v.Detail()
				fmt.Println("工单来源",tb.Tksource.Content)
				fmt.Println("工单满意度",tb.Satisfaction.Content)
				fmt.Println("工单内容:",tb.Tkcontent.Content)

			}
		}

	}
}
