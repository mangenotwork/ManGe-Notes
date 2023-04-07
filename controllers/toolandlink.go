package controllers

/*
	工具与链接模块
*/

import (
	"fmt"

	object "github.com/mangenotwork/ManGe-Notes/object"
	servers "github.com/mangenotwork/ManGe-Notes/servers"
	//util "man/ManNotes/util"
)

type TandLController struct {
	Controller
}

//添加网络资源
func (this *TandLController) AddCollectLink() {
	uid := this.GetUid()

	var obj object.AddLink

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj, uid)

	code, count, data := new(servers.TandLServers).AddLink(&obj, uid)

	this.RetuenJson(code, count, data)
}

//获取网络工具列表
func (this *TandLController) GetTandL() {
	uid := this.GetUid()
	code, count, data := new(servers.TandLServers).GetToolList(uid)
	this.RetuenJson(code, count, data)
}

//收藏链接的显示页
func (this *TandLController) LinkShow() {
	uid := this.GetUid()

	fmt.Println(uid)

	code, count, data := new(servers.TandLServers).GetLinks(uid)
	fmt.Println(code, count)

	this.Data["Links"] = data
	this.TplName = "pg/linkshow.html"
}

//修改收藏的链接
func (this *TandLController) EDLink() {
	uid := this.GetUid()

	var obj object.EDLinks

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj, uid)

	code, count, data := new(servers.TandLServers).EditLink(&obj, uid)

	this.RetuenJson(code, count, data)
}

//删除收藏的链接
func (this *TandLController) DELLink() {
	uid := this.GetUid()
	linkid, _ := this.GetInt("linkid")

	code, count, data := new(servers.TandLServers).DELLink(uid, linkid)

	this.RetuenJson(code, count, data)
}

//GetLinks  mange 管理模块 获取收藏链接信息
func (this *TandLController) GetLinks() {
	uid := this.GetUid()

	count, data := new(servers.TandLServers).GetAllLinksInfo(uid)
	this.MangeJson(0, "", count, data)
}

//MageEDLink  mange 管理模块  收藏链接修改
func (this *TandLController) MageEDLink() {
	uid := this.GetUid()

	var obj object.EDLinksInfo
	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj, uid)

	code, count, data := new(servers.TandLServers).MangeEditLink(&obj, uid)

	this.RetuenJson(code, count, data)
}
