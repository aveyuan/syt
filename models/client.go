package models

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	Username string //账户名称
	Password string //密码
	Nickname string //昵称
	Sex int //性别
	Phone1 string //电话1
	Phone2 string //电话2
	Email string //邮箱
	Address string //地址
	Avatar string //头像
	Remarks string //备注
	Tickets []Ticket //关联的工单
}

//新增客户
func (this *Client)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//客户信息修改（包括密码）
func (this *Client)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}

//客户删除
func (this *Client)Delete()error  {
	if err :=db.Delete(this).Error;err!=nil{
		return err
	}
	return nil
}

//客户列表
func (this *Client)List()(*[]Client,error)  {
	var clients []Client
	if err :=db.Find(&clients).Error;err!=nil{
		return &clients,err
	}
	return &clients,nil
}