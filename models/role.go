package models

//用户角色表

type Role struct {
	ID      uint
	content string
	User    []User `gorm:"many2many:user_role"`
}
