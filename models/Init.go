package models

import (
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

//设置一个全局的db方便其他函数调用
var db *gorm.DB
var err error
//初始化db，数据库连接，创建等
func Init()  {
	//获取配置信息
	ini,err := ini.Load("config/config.ini")
	if err !=nil{
		log.Fatal("配置文件读取出错,请检查")
		return
	}
	db_user := ini.Section("").Key("db_user").String()
	db_pass := ini.Section("").Key("db_pass").String()
	db_host := ini.Section("").Key("db_host").String()
	db_db := ini.Section("").Key("db_db").String()
	db_port := ini.Section("").Key("db_port").String()
	if db_port==""{
		db_port="3306"
	}

	log.Println("正在初始化数据库")
	con := db_user+":"+db_pass+"@("+db_host+":"+db_port+")/"+db_db+"?charset=utf8&parseTime=True&loc=Local"
	db,err  = gorm.Open("mysql",con)
	if err !=nil{
		log.Fatal("connect err",err)
	}

	//设置数据库前缀
	gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
		return "syt_"+defaultTableName
	}
	log.Println("初始化数据库完成")

	if len(os.Args)>=2 && os.Args[1]=="install"{
		log.Println("安装正在检查迁移数据库信息")
		db.AutoMigrate(&User{},&Ticket{},&Satisfaction{},&Tkcontent{},&Tksource{})
		createAdmin()

		//测试数据信息
		Tsource()
		TSatisfactions()
		Tkcreates()
		TCreateuser()
		Tuser()
		}
	}

//创建管理账户
func createAdmin()  {
	//创建前检查管理员账户是否存在
	user := &User{}
	db.Where("username=?","admin").Find(&user)
	//没有找到管理员账户则创建一个
	if user.Nickname==""{
		user = &User{Username:"admin",Password:"123456",Nickname:"管理员"}
		if err:=user.Add();err!=nil{
			log.Println("用户添加失败")
		}
		log.Printf("系统用户初始化成功，根据以下信息建议您及时修改密码:\n用户名:admin\n密码:123456")
	}

}

