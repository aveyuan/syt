package models

//工单满意度
type Satisfaction struct {
	Id int //id
	Content string //内容
	TicketId int //关联工单
}
