package models

import (
	"log"
	"testing"
)

//测试client
func Testclient(t *testing.T)  {

	testdata()
}

//测试数据
func testdata()  {
	//测试客户添加
	client := &Client{Username:"张三",Password:"123",Phone1:"12345678901",Email:"zs@qq.com",Address:"重庆"}
	if err := client.Add();err !=nil{
		log.Print("添加客户失败")
	}else {
		log.Print("添加客户成功")
	}

	client1 := &Client{Username:"李四",Password:"123",Phone1:"12345678901",Email:"zs@qq.com",Address:"重庆"}
	if err := client1.Add();err !=nil{
		log.Print("添加客户失败")
	}else {
		log.Print("添加客户成功")
	}

}