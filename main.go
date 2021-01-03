package main

import (
	"github.com/astaxie/beego"
	_ "github.com/mangenotwork/ManGe-Notes/routers"
)

func main() {

	//检查安装
	//这里采用检查本地key文件方案

	beego.Run()
}
