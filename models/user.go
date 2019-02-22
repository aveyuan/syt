package models

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/aveyuan/syt/libs"
	"time"
)

type User struct {
	gorm.Model
	Username string  //账户名称
	Password string `json:"-"`//密码
	Salt	string  `json:"-"`//密码加盐
	Nickname string //昵称
	Email string //邮箱
	Avatar string //头像
	Token string `json:"-"`//token
	Phone1 string //电话1
	Phone2 string //电话2
	LastTime time.Time //最后一次登录时间
	Lastip string `json:"-"`//最后一次登录IP
	Tkcontents []Tkcontent //关联工单内容
	Solvetickets []Ticket `gorm:"many2many:user_solvetikets"` //解决的工单
	Tickets []Ticket  //创建的工单
	Role []Role `gorm:"many2many:user_role"`
}

//为了避免密码暴露，在登录验证的时候使用VliUser
type VliUser struct {
	Username string `json:"username"  binding:"required"` //账户名称
	Password string `json:"password"  binding:"required"`//密码
	RePassword string `json:"repassword" gorm:"-"`//重复密码
	Nickname string  `json:"nickname"`//昵称
	Email string `json:"email"`//邮箱
	Phone1 string `json:"phone1"`//电话1
	Lastip string //最后一次登录IP
}


//用户注册
func (this *VliUser)Reg()error  {
	var user User
	user.Username=this.Username
	user.Password=this.Password
	user.Nickname=this.Nickname
	user.Email=this.Email
	user.Phone1=this.Phone1
	//添加用户
	if err := user.Add();err !=nil{
		return err
	}
	return nil
}

//用于生成jwt信息
type Jwtuser struct {
	Username string
}

func IdUser(user int)(*User,error)  {
	var u User
	if err :=db.First(&u,user).Error;err !=nil{
		return &u,err
	}
	return &u,nil
}

//新增用户
func (this *User)Add()error  {
	password,salt := libs.Password(this.Password)
	this.Password=password
	this.Salt=salt
	if err :=db.Create(this).Error;err!=nil{
		return err
	}
	return nil
}

//用户信息修改
func (this *User)Update()error  {
	if err :=db.Save(this).Error;err!=nil{
		return err
	}
	return nil
}

//账号密码验证
//先查询有没有这个用户，然后再将这个用户的salt拿出来和传过来的密码进行加密，最后再比对密码是否匹配
func (this *VliUser)Valid()error  {
	var user User
	if err := db.Where("username = ?",this.Username).Find(&user).Error;err!=nil{
		//没有找到用户
		return err
	}
	//找到了用户，把密码拿出来加密比对
	pass := libs.Md5([]byte(this.Password+user.Salt))
	if pass != user.Password{
		return errors.New("密码错误")
	}
	//记录最后一次登录时间
	user.LastTime=time.Now().Local()
	user.Lastip=this.Lastip
	db.Save(&user)
	return nil

}

//用户详情,只有基本的用户信息
func (this *User)Detail()(*User,error)  {
	username := this.Username
	var user User
	if err :=db.Where("username=?",username).Find(&user).Error;err !=nil{
		return &user,err
	}
	return &user,nil
}

//用户工单
//将这个函数进行复用，根据用户传入工单状态，返回工单信息
func (this *User)UserTickets(status int,search string)([]Ticket,error)  {
	var tickets []Ticket
	user,_ := this.Detail()
	//传入0表示查询所有的数据，不筛选
	//传入8表示查询非结案工单，表示进行中的工单
	if status==0 && search==""{
		if err :=db.Model(user).Association("Tickets").Find(&tickets).Error;err!=nil{
			return tickets,err
		}
	}else if(status==8){
		if err :=db.Not("status = ?",0).Model(user).Association("Tickets").Find(&tickets).Error;err!=nil{
			return tickets,err
		}
	}else if status==0 && search!="" {
		if err :=db.Where("title LIKE ?","%"+search+"%").Model(user).Association("Tickets").Find(&tickets).Error;err!=nil{
			return tickets,err
		}
	}else {
		if err :=db.Where("status = ?",status).Where("title LIKE ?","%"+search+"%").Model(user).Association("Tickets").Find(&tickets).Error;err!=nil{
			return tickets,err
		}
	}

	return tickets,nil
}
