package conn

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/util"
)

func CreateSqliteDB() (bool, error) {
	dbFile := "./db/base.db"
	//判断是否存在db目录
	util.CreateMutiDir("./db")

	_, err := os.Lstat(dbFile)
	//没有则创建
	if os.IsNotExist(err) {
		f, err := os.Create(dbFile)
		defer f.Close()
		if err != nil {
			log.Println("创建文件 err = ", err.Error())
		}
	}

	db, err := gorm.Open("sqlite3", dbFile)
	defer db.Close()
	if err != nil {
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

//获取Sqlite 连接
func GetSqliteConn() (*gorm.DB, error) {
	dbFile := "./db/base.db"
	//判断是否存在db目录
	util.CreateMutiDir("./db")

	_, err := os.Lstat(dbFile)
	//没有则创建
	if os.IsNotExist(err) {
		f, err := os.Create(dbFile)
		defer f.Close()
		if err != nil {
			log.Println("创建文件 err = ", err.Error())
		}
	}
	db, err := gorm.Open("sqlite3", dbFile)
	db.LogMode(true)
	return db, err
}
