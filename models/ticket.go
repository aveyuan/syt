package models

import (
	"fmt"
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
	Status int  `gorm:"default:'2'"`//状态 1.结案,2.新的,3跟进中,4.已解决,5.挂起

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
	ID 	int
	TksourceId uint `json:"tksourceid" binding:"required"`  //工单来源
	Status int  `json:"Status" binding:"required"`//工单状态
	Solveuser	[]User `json:"solveuser"` //工单处理人
}

//更新工单信息
func (this *TkSave)Update()error  {
	var tk Ticket

	if err :=db.Where("id = ?",this.ID).Model(&tk).Updates(Ticket{Status:this.Status, TksourceId: this.TksourceId}).Error;err !=nil{
		return err
	}

	//再次查找工单，来更新关系,这里有个坑，必须再次查询才行,不能利用上面的updates
	db.Where("id = ?",this.ID).First(&tk)
	if err := db.Model(&tk).Association("Solveusers").Replace(&this.Solveuser).Error;err!=nil{
		fmt.Print(err)
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
	db.Model(&ticket).Association("solveusers").Find(&solveuser)

	//组合数据
	tb := &Tkbase{Ticket:ticket,Tksource:tksource,Satisfaction:satis,Tkcontent:tkcontent,User:user,Solveuser:solveuser}
	return tb
}

//显示所有工单
//可以根据工单的状态来进行筛选
func (this *Ticket)List(status int,search string)([]Ticket,error)  {
	//status=0为显示所有，其他的根据id来选
	var tickets []Ticket
	if status==0&&search==""{
		if err := db.Find(&tickets).Order("ID desc").Error;err !=nil{
			return tickets,err
		}
	}else if status==0&&search!="" {
		if err := db.Where("title LIKE ?","%"+search+"%").Find(&tickets).Order("ID desc").Error;err !=nil{
			return tickets,err
		}
	}else {
		if err := db.Where("status = ?",status).Where("title LIKE ?","%"+search+"%").Find(&tickets).Order("ID desc").Error;err !=nil{
			return tickets,err
		}
	}

	return tickets,nil
}