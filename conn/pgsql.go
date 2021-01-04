package conn

import (
	"fmt"
	_ "os"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Pgsql 测试连接
func PgsqlConnTest(host, user, password, dbname string) (bool, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, password, dbname)
	db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	defer db.Close()
	if err != nil {
		beego.Error("[pgsql]连接异常:", err.Error(), connStr)
		return false, err
	}
	return true, nil
}
