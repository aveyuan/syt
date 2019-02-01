package models

//工单满意度
type Satisfaction struct {
	Id uint //id
	Content string //内容
	Tickets Ticket //关联工单
}

//工单满意度
func (this *Satisfaction)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//更新满意度
func (this *Satisfaction)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}