package models

//工单来源表
type Tksource struct {
	Id uint //id
	Content string //来源内容
	Tickets Ticket //关联到的工单
}

//新增工单来源
func (this *Tksource)Add()error  {
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//更新工单来源
func (this *Tksource)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}