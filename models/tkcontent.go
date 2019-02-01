package models

import "github.com/jinzhu/gorm"

//工单跟进内容
type Tkcontent struct {
	gorm.Model
	TicketID uint //关联到的工单
	Content string //跟进内容
	File	string //文件
	UserID  uint //关联到跟进人
	ClientID uint //关联到用户
	Remarks string //备注
}

//新增工单跟进内容
func (this *Tkcontent)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}