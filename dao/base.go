package dao

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mangenotwork/ManGe-Notes/conn"
	"github.com/mangenotwork/ManGe-Notes/object"
)

//获取连接
func GetConn() *gorm.DB {
	log.Println("object.GlobalDBType = ", object.GlobalDBType)
	switch object.GlobalDBType {
	case "sqlite":
		c, err := conn.GetSqliteConn()
		if err != nil {
			log.Println("连接错误 = ", err)
		}
		return c
	case "mysql":
		return conn.NotesDB()
	case "pgsql":
		c, err := conn.GetPgsqlConn()
		if err != nil {
			log.Println("连接错误 = ", err)
		}
		return c
	}
	return nil
}
