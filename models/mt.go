package models

import (
	"log"
	"github.com/aveyuan/syt/libs"
	"fmt"
)


//测试数据
func TClient()  {
	//测试客户添加
	pass,salt := libs.Password("123456")
	client := &Client{Username:"张三",Password:pass,Salt:salt,Phone1:"12345678901",Email:"zs@qq.com",Address:"重庆"}
	if err := client.Add();err !=nil{
		log.Print("添加客户失败")
	}else {
		log.Print("添加客户成功")
	}

	client1 := &Client{Username:"李四",Password:pass,Salt:salt,Phone1:"12345678901",Email:"zs@qq.com",Address:"重庆"}
	if err := client1.Add();err !=nil{
		log.Print("添加客户失败")
	}else {
		log.Print("添加客户成功")
	}
}

//添加工单来源
func Tsource()  {
	tksource1 := &Tksource{Content:"微信"}
	tksource2 := &Tksource{Content:"工单提交"}
	if err :=tksource1.Add();err !=nil{
		log.Print("添加来源失败",err)
	}else {
		log.Print("添加来源成功")
	}
	if err :=tksource2.Add();err !=nil{
		log.Print("添加客户失败")
	}else {
		log.Print("添加客户成功")
	}
}

//添加满意度
func TSatisfactions()  {
	s1 := &Satisfaction{Content:"满意"}
	s2 := &Satisfaction{Content:"一般"}
	s3 := &Satisfaction{Content:"不满意"}
	if err :=s1.Add();err !=nil{
		log.Print("添加满意度失败",err)
	}else {
		log.Print("添加满意度成功")
	}
	if err :=s2.Add();err !=nil{
		log.Print("添加满意度失败")
	}else {
		log.Print("添加满意度成功")
	}
	if err :=s3.Add();err !=nil{
		log.Print("添加满意度失败")
	}else {
		log.Print("添加满意度成功")
	}
}

//创建工单
func Tkcreates()  {
	var client Client
	db.First(&client)
	//工单来源
	tksource := &Tksource{Id:1}

	//添加工单


	tkc := &Tkcontent{Content:"需要帮助,我的电脑坏了",ClientID:client.ID}
	tks := &Ticket{ClientID:client.ID,Status:1,TksourceId:tksource.Id,Contents:[]Tkcontent{*tkc}}
	if err :=tks.Add();err !=nil{
		log.Print("添加工单失败")
	}else {
		log.Print("添加工单成功")
	}
}

//分配工单给处理人员
func Tuser()  {
	var user1 User
	db.First(&user1)

	var user2 User
	db.Last(&user2)

	var tk Ticket
	db.First(&tk)
	tk.Users=[]User{user1,user2}
	tk.Status=3
	if err := tk.Update();err !=nil{
		fmt.Println("工单更新失败")
	}else {
		fmt.Println("工单更新成功")
	}
}

func TCreateuser()  {
	//创建前检查管理员账户是否存在
	user := &User{}
	db.Where("username=?","zhangsan").Find(&user)
	//没有找到管理员账户则创建一个
	if user.Nickname==""{
		pass,salt := libs.Password("123456")
		user = &User{Username:"zhangsan",Password:pass,Salt:salt,Nickname:"张三"}
		if err:=user.Add();err!=nil{
			log.Println("用户添加失败")
		}
		log.Printf("用户添加成功")
	}

}

//客户列表
func Tclist()  {
	clint := &Client{}
	cl,err := clint.List()
	if err !=nil{
		fmt.Println("客户列表获取失败")
	}
	for _,v := range *cl{
		fmt.Println(v.Username)
	}
}