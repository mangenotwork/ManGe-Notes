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

//收藏链接的显示页
func (this *TandLController) LinkShow(){
	uid := this.GetUid()

	fmt.Println(uid)

	code,count,data := new(servers.TandLServers).GetLinks(uid)
	fmt.Println(code,count)

 	this.Data["Links"] = data
	this.TplName = "pg/linkshow.html"
}


//修改收藏的链接
func (this *TandLController) EDLink() {
	uid := this.GetUid()

	var obj object.EDLinks

	err := this.ResolvePostData(&obj)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(obj,uid)

	code,count,data := new(servers.TandLServers).EditLink(&obj,uid)

	this.RetuenJson(code,count,data)
}

//删除收藏的链接
func (this *TandLController) DELLink(){
	uid := this.GetUid()
	linkid,_ := this.GetInt("linkid")

	code,count,data := new(servers.TandLServers).DELLink(uid,linkid)

	this.RetuenJson(code,count,data)
}