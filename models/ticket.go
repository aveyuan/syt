package models

import (
	"github.com/jinzhu/gorm"
)

//工单

type Ticket struct {
	gorm.Model
	Tkcontents []Tkcontent //关联到跟进内容
	UserID uint //关联创建工单的用户
	Title string //工单主题
	Solveusers []User `gorm:"many2many:user_solvetikets"` //关联支持用户
	TksourceId uint //工单来源
	SatisfactionId uint //工单满意度
	Status int  `gorm:"default:'1'"`//状态 0.结案,1.新的,2跟进中,3.已解决,4.挂起

}

//用于保存得到了工单数据的表,这个没有和数据库关联,用于显示的
type Tkbase struct {
	Ticket Ticket
	Tksource Tksource
	Satisfaction Satisfaction
	Tkcontent Tkcontent
	User	User
	Solveuser []User
}

//创建工单时所需要的数据
type TkCreate struct {
	User User `json:"-" binding:"required"`  //哪个用户提起的工单
	Title string `json:"title" binding:"required"` //标题是什么
	Tksource uint `json:"tksource" binding:"required"` //工单从哪里提交的
} 

//更新工单主体内容
//包含工单的来源，处理人，状态
type TkSave struct {
	ID uint  `json:"id" binding:"required"`//工单ID
	TksourceId uint `json:"tksourceid" binding:"required"`  //工单来源
	Status int  `json:"Status" binding:"required"`//工单状态
	Solveuser	[]User `json:"solveuser"`
}

//更新用户信息
func (this *TkSave)Update()error  {
	var tk Ticket
	db.Find(this.ID).First(&tk)
	//tk.Solveusers=this.Solveuser
	if err :=db.Model(&tk).Updates(Ticket{Status:this.Status, TksourceId: this.TksourceId,Solveusers:this.Solveuser}).Error;err !=nil{
		return err
	}
	return nil
}

//新增工单
func (this *TkCreate)Add()error  {
	//解析到数据表
	tk := &Ticket{}
	tk.Title=this.Title
	tk.UserID=this.User.ID
	tk.TksourceId=this.Tksource
	if err :=db.Create(tk).Error;err!=nil{
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
	var solveuser []User
	var user User
	db.Where("id=?",id).Find(&ticket)
	//工单来源
	db.Model(&ticket).Related(&tksource).Find(&tksource)
	db.Model(&ticket).Related(&satis).Find(&satis)
	db.Model(&ticket).Related(&ticket.Tkcontents).Find(&tkcontent)
	db.Model(&ticket).Related(&user).Find(&user)
	db.Model(&ticket).Related(&ticket.Solveusers).Find(&solveuser)
	db.Model(&ticket).Association("solveusers").Find(&solveuser)

	//组合数据
	tb := &Tkbase{Ticket:ticket,Tksource:tksource,Satisfaction:satis,Tkcontent:tkcontent,User:user,Solveuser:solveuser}
	return tb
}

//显示所有工单
func (this *Ticket)List()([]Ticket,error)  {
	var tickets []Ticket
	if err := db.Find(&tickets).Order("ID desc").Error;err !=nil{
		return tickets,err
	}
	return tickets,nil
}