package models

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/aveyuan/syt/libs"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"` //账户名称
	Password string `json:"password" form:"password" binding:"required"`//密码
	Salt	string //密码加盐
	Nickname string //昵称
	Email string //邮箱
	Avatar string //头像
	Token string //token
	Phone string //电话
	LastTime time.Time //最后一次登录时间
	Tkcontents []Tkcontent //关联工单内容
	Tickets []Ticket `gorm:"many2many:user_tiket"` //关联工单
}

//新增用户
func (this *User)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//用户信息修改（包括密码）
func (this *User)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}

//账号密码验证
//先查询有没有这个用户，然后再将这个用户的salt拿出来和传过来的密码进行加密，最后再比对密码是否匹配
func (this *User)Valid()error  {
	var user User
	if err := db.Where("username = ?",this.Username).Find(&user).Error;err!=nil{
		//没有找到用户
		return err
	}
	//找到了用户，把密码拿出来加密比对
	pass := libs.Md5([]byte(this.Password+user.Salt))
	if pass != user.Password{
		return errors.New("密码错误")
	}
	//记录最后一次登录时间
	user.LastTime=time.Now().Local()
	db.Save(&user)
	return nil

}