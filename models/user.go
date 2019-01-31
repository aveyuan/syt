package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string //用户名
	Password string //密码
	Nickname string //昵称
	Email string //邮箱
	Avatar string //头像
	Token string //token
	Phone string //电话

	Tickets []Ticket `gorm:"many2many:user_tiket"` //关联工单
}

//新增用户
func (this *User)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}