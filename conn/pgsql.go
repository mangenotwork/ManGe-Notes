package conn

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
)

//Pgsql 测试连接
func PgsqlConnTest(host, user, password, dbname string) (bool, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, password, dbname)
	db, err := gorm.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		beego.Error("[pgsql]连接异常:", err.Error(), connStr)
		return false, err
	}
	return true, nil
}

//Pgsql 连接并创建
func CreatePgsqlDB(host, user, password, dbname string) (bool, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, password, dbname)
	db, err := gorm.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		beego.Error("[pgsql]连接异常:", err.Error(), connStr)
		return false, err
	}

	//检查table
	//如果没有就创建
	if !db.HasTable(&models.ACC{}) {
		log.Println("ACC 不存在")
		db.CreateTable(&models.ACC{})
	}
	if !db.HasTable(&models.IMGInfo{}) {
		log.Println("IMGInfo 不存在")
		db.CreateTable(&models.IMGInfo{})
	}
	if !db.HasTable(&models.MDInof{}) {
		log.Println("MDInof 不存在")
		db.CreateTable(&models.MDInof{})
	}
	if !db.HasTable(&models.MDText{}) {
		log.Println("MDText 不存在")
		db.CreateTable(&models.MDText{})
	}
	if !db.HasTable(&models.Notes{}) {
		log.Println("Notes 不存在")
		db.CreateTable(&models.Notes{})
	}
	if !db.HasTable(&models.SCIMGInfo{}) {
		log.Println("SCIMGInfo 不存在")
		db.CreateTable(&models.SCIMGInfo{})
	}
	if !db.HasTable(&models.ToolandLink{}) {
		log.Println("ToolandLink 不存在")
		db.CreateTable(&models.ToolandLink{})
	}

	return true, nil
}

//获取Pgsql连接
func GetPgsqlConn() (*gorm.DB, error) {

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", object.GlobalPgsqlHost,
		object.GlobalPgsqlUser, object.GlobalPgsqlPassword, object.GlobalPgsqlDBName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		beego.Error("[pgsql]连接异常:", err.Error(), connStr)
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
