package main

import (
	"github.com/astaxie/beego"
	"github.com/mangenotwork/ManGe-Notes/conn"
	"github.com/mangenotwork/ManGe-Notes/object"
	_ "github.com/mangenotwork/ManGe-Notes/routers"
	"github.com/mangenotwork/ManGe-Notes/util"
)

func main() {

	conn.CachesInit()

	//检查安装
	if util.FileExist(object.InstallJsonPath) {
		installInfo := object.OpenInstallFile()
		switch installInfo.DBType {
		case "sqlite":
			object.GlobalDBType = "sqlite"
		case "mysql":
			object.GlobalDBType = "mysql"
			object.GlobalMysqlHost = installInfo.MysqlHost
			object.GlobalMysqlPort = installInfo.MysqlPort
			object.GlobalMysqlUser = installInfo.MysqlUser
			object.GlobalMysqlPassword = installInfo.MysqlPassword
			object.GlobalMysqlDBName = installInfo.MysqlDBName
			conn.MysqlInit()
		case "pgsql":
			object.GlobalDBType = "pgsql"
			object.GlobalPgsqlHost = installInfo.PgsqlHost
			object.GlobalPgsqlUser = installInfo.PgsqlUser
			object.GlobalPgsqlPassword = installInfo.PgsqlPassword
			object.GlobalPgsqlDBName = installInfo.PgsqlDBName
		}
	}

	beego.Run()
}
