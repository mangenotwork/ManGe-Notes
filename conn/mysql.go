package conn

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	user     = beego.AppConfig.DefaultString("mysql::user", "")
	password = beego.AppConfig.DefaultString("mysql::password", "")
	host     = beego.AppConfig.DefaultString("mysql::host", "")
	port     = beego.AppConfig.DefaultString("mysql::port", "")
	dbs      map[string]*gorm.DB
)

func init() {

	dbs = make(map[string]*gorm.DB)
	SetDBConn("notes")
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

func MysqlDB(dbName string) *gorm.DB {
	fmt.Println(user, password, host, port)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName) + "?parseTime=true"
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		beego.Error("[mysql]连接异常:", err.Error(), connStr)
		//添加连接错误通知或触发解决事件
	}
	return db
}

func SetDBConn(dbName string) {

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

func GetDBConns(dbName string) *gorm.DB {
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
