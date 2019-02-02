package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

//工单

type Ticket struct {
	gorm.Model
	Tkcontents []Tkcontent //关联到跟进内容
	UserID uint //关联创建工单的用户
	Solveusers []User `gorm:"many2many:user_solvetikets"` //关联支持用户
	TksourceId uint //工单来源
	SatisfactionId uint //工单满意度
	Status int //状态 0.结案,1.新的,2跟进中,3.已解决,4.挂起

}

//用于保持得到了工单数据的表,这个没有和数据库关联,用于显示的
type Tkbase struct {
	Ticket Ticket
	Tksource Tksource
	Satisfaction Satisfaction
	Tkcontent Tkcontent
}

//新增工单
func (this *Ticket)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//更新工单
func (this *Ticket)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}

//显示工单数据
func (this *Ticket)Detail()(*Tkbase)  {
	id := this.ID
	var ticket Ticket
	var tksource Tksource
	var satis Satisfaction
	var tkcontent Tkcontent
	db.Where("id=?",id).Find(&ticket)
	//工单来源
	db.Model(&ticket).Related(&tksource).Find(&tksource)
	db.Model(&ticket).Related(&satis).Find(&satis)
	db.Model(&ticket).Related(&ticket.Tkcontents).Find(&tkcontent)
	//组合数据
	tb := &Tkbase{Ticket:ticket,Tksource:tksource,Satisfaction:satis,Tkcontent:tkcontent}
	return tb
}

//显示所有工单
func (this *Ticket)List()([]Ticket,error)  {
	var tickets []Ticket
	if err := db.Find(&tickets).Order("ID desc").Error;err !=nil{
		return tickets,err
	}
	for _,v := range tickets{
		fmt.Println(v.TksourceId)
	}
	return tickets,nil
}