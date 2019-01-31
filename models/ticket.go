package models

import (
	"github.com/jinzhu/gorm"
)

//工单

type Ticket struct {
	gorm.Model
	Content string //工单内容
	ClientID int //管理的客户ID
	Users []User `gorm:"many2many:user_tiket"` //关联支持用户
	Tksource Tksource //工单来源
	Satisfaction Satisfaction //工单满意度
	Status int //状态 0.结案,1.新的,2跟进中,3.已解决,4.挂起

}

//新增工单
func (this *Ticket)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}