package controllers

import (
	_ "encoding/json"
	"fmt"

	//rdb "man/ManNotes/models/redis"
	object "github.com/mangenotwork/ManGe-Notes/object"
	servers "github.com/mangenotwork/ManGe-Notes/servers"
)

type SucaiController struct {
	Controller
}

//AddLinkImg  添加网络链接图片
func (this *SucaiController) AddLinkImg() {
	uid := this.GetUid()
	if uid == "" {
		this.RetuenJson(-1, 1, "请登录")
		return
	}

	var obj object.LinkImgData

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj, uid)

	code, count, data := new(servers.SUCai).ADDLinkImg(uid, &obj)

	this.RetuenJson(code, count, data)
}

//MyImg  获取我的图片
func (this *SucaiController) MyImg() {
	uid := this.GetUid()
	if uid == "" {
		this.RetuenJson(-1, 1, "请登录")
		return
	}

	code, count, data := new(servers.SUCai).GetMyImg(uid)

	this.RetuenJson(code, count, data)

}

//ToMangeImg 分享到mange 图库
func (this *SucaiController) ToMangeImg() {
	uid := this.GetUid()
	if uid == "" {
		this.RetuenJson(-1, 1, "请登录")
		return
	}

	imgid, err := this.GetInt("imgid")
	if err != nil {
		this.RetuenJson(-1, 1, "错误请求")
		return
	}

	code, count, data := new(servers.SUCai).ToMangeImg(uid, imgid)

	this.RetuenJson(code, count, data)
}

//MangeImg 漫鸽图库图片素材
func (this *SucaiController) MangeImg() {
	uid := this.GetUid()
	if uid == "" {
		this.TplName = "pg/login.html"
		return
	}
	code, count, data := new(servers.SUCai).GetMangeImg(uid)

	this.Data["Mange"] = true
	this.Data["Code"] = code
	this.Data["Count"] = count
	this.Data["Data"] = data
	this.TplName = "sucai/sc.html"
}
