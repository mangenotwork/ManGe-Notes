package conn

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateSqliteDB() (bool, error) {
	dbFile := "./db/base.db"
	//判断是否存在db目录

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
	return true, nil
}
