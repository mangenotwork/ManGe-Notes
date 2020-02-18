package controllers

/*
		工具与链接模块
*/

import (
	"fmt"

	object "man/ManNotes/object"
	servers "man/ManNotes/servers"
	//util "man/ManNotes/util"
)

type TandLController struct {
	Controller
}

//添加网络资源
func (this *TandLController) AddCollectLink(){
	uid := this.GetUid()

	var obj object.AddLink

	err := this.ResolvePostData(&obj)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(obj,uid)

	code,count,data := new(servers.TandLServers).AddLink(&obj,uid)

	this.RetuenJson(code,count,data)
}

//获取网络工具列表
func (this *TandLController) GetTandL(){
	uid := this.GetUid()
	code,count,data := new(servers.TandLServers).GetToolList(uid)
	this.RetuenJson(code,count,data)
}

