package models

import (
	"github.com/jinzhu/gorm"
	"github.com/aveyuan/syt/libs"
	"errors"
	"time"
)

type Client struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"` //账户名称
	Password string `json:"password" form:"password" binding:"required"`//密码
	Salt	string //密码加盐
	Nickname string //昵称
	Sex int //性别
	Phone1 string //电话1
	Phone2 string //电话2
	Email string //邮箱
	Address string //地址
	Avatar string //头像
	Remarks string //备注
	LastTime time.Time //最后一次登录时间
	Tickets []Ticket //关联的工单
	Tkcontents []Tkcontent //关联到内容
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

//账号密码验证
//先查询有没有这个用户，然后再将这个用户的salt拿出来和传过来的密码进行加密，最后再比对密码是否匹配
func (this *Client)Valid()error  {
	var client Client
	if err := db.Where("username = ?",this.Username).Find(&client).Error;err!=nil{
		//没有找到用户
		return err
	}
	//找到了用户，把密码拿出来加密比对
	pass := libs.Md5([]byte(this.Password+client.Salt))
	if pass != client.Password{
		return errors.New("密码错误")
	}
	//记录最后一次登录时间
	client.LastTime=time.Now().Local()
	db.Save(&client)
	return nil

}

func (this *Client)Detail()(*Client,error)  {
	username := this.Username
	var client Client
	db.Where("username=?",username).Find(&client)
	//查询工单
	if err :=db.Model(&client).Related(&client.Tickets).Find(&client.Tickets).Error;err !=nil{
		return &client,err
	}
	return &client,nil
}