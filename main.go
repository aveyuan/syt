package main

import (
	"github.com/aveyuan/syt/models"
	"github.com/aveyuan/syt/router"
)

func main() {
	//初始化数据库
	models.Init()
	//初始化路由
	r := router.Init()
	//启动服务器
	r.Run()

}
