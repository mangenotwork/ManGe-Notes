package conn

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
)

var (
	user     = object.GlobalMysqlUser
	password = object.GlobalMysqlPassword
	host     = object.GlobalMysqlHost
	port     = object.GlobalMysqlPort
	dbName   = object.GlobalMysqlDBName
	dbs      map[string]*gorm.DB
)

func MysqlInit() {

	dbs = make(map[string]*gorm.DB)
	SetDBConn()
}

/*
//在内部db取值
var (
	user = "root"
	password = "123"
	host = "127.0.0.1"
	port = "3306"
)
*/

func MysqlDB() *gorm.DB {
	fmt.Println(user, password, host, port)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName) + "?parseTime=true"
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		beego.Error("[mysql]连接异常:", err.Error(), connStr)
		//添加连接错误通知或触发解决事件
	}
	return db
}

func SetDBConn() {

	var (
		db  *gorm.DB
		err error
	)

	fmt.Println(user, password, host, port)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName) + "?parseTime=true"
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		beego.Error("[mysql]连接异常:", err.Error(), connStr)
		//添加连接错误通知或触发解决事件
	}
	db.CommonDB()
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	dbs[dbName] = db
}

func GetDBConns() *gorm.DB {
	return dbs[dbName]
}

func NotesDB() *gorm.DB {
	db := dbs["notes"]
	db.LogMode(true)
	return db
}

func MysqlConnTest(host, port, user, password, dbname string) (bool, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname) + "?parseTime=true"
	_, err := gorm.Open("mysql", connStr)
	if err != nil {
		beego.Error("[mysql]连接异常:", err.Error(), connStr)
		return false, err
	}
	return true, nil
}

func CreateMysqlDB(host, port, user, password, dbname string) (bool, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname) + "?parseTime=true"
	db, err := gorm.Open("mysql", connStr)
	defer db.Close()
	if err != nil {
		beego.Error("[mysql]连接异常:", err.Error(), connStr)
		return false, err
	}

	//检查table
	//如果没有就创建
	if !db.HasTable(&models.ACC{}) {
		log.Println("ACC 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT  CHARSET=utf8").CreateTable(&models.ACC{})
	}
	if !db.HasTable(&models.IMGInfo{}) {
		log.Println("IMGInfo 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.IMGInfo{})
	}
	if !db.HasTable(&models.MDInof{}) {
		log.Println("MDInof 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.MDInof{})
	}
	if !db.HasTable(&models.MDText{}) {
		log.Println("MDText 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.MDText{})
	}
	if !db.HasTable(&models.Notes{}) {
		log.Println("Notes 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.Notes{})
	}
	if !db.HasTable(&models.SCIMGInfo{}) {
		log.Println("SCIMGInfo 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.SCIMGInfo{})
	}
	if !db.HasTable(&models.ToolandLink{}) {
		log.Println("ToolandLink 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.ToolandLink{})
	}
	return true, err
}
