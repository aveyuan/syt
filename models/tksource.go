package models

//工单来源表
type Tksource struct {
	Id int //id
	Content string //来源内容
	TicketId int //关联工单ID
}
