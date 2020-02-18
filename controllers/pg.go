package controllers

import (
	"fmt"
	"encoding/json"

	rdb "man/ManNotes/models/redis"
	object "man/ManNotes/object"
)

type PGController struct {
	Controller
}

//登录页面
func (this *PGController) LoginPG() {
	this.IsLogin()
	this.TplName = "pg/login.html"
}

//主页
func (this *PGController) IndexPG() {
	//获取用户基本信息
	uid := this.GetUid()
	fmt.Println(fmt.Sprintf("uinfo:%s",uid))
	userbinfoStr := new(rdb.RDB).HashGet(fmt.Sprintf("uinfo:%s",uid),"basis")
	fmt.Println(userbinfoStr)

	var userbinfo object.UserBasisInfo
 	// 将字符串反解析为结构体
 	json.Unmarshal([]byte(userbinfoStr), &userbinfo)
 	fmt.Println(userbinfo)

 	this.Data["userbinfo"] = &userbinfo
 	this.Data["IsShowNav"] = "index"
	this.TplName = "index.html"
}

//MD编辑器
func (this *PGController) MdEditorPG() {
	this.IsSession()
	this.TplName = "pg/createmd.html"
}

//首页
func (this *PGController) HomePG(){
	this.IsSession()
	this.TplName = "pg/home.html"
}

//PC端页面渲染
func (this *PGController) ToolPG(){
	//获取用户基本信息
	uid := this.GetUid()
	fmt.Println(fmt.Sprintf("uinfo:%s",uid))
	userbinfoStr := new(rdb.RDB).HashGet(fmt.Sprintf("uinfo:%s",uid),"basis")
	fmt.Println(userbinfoStr)

	var userbinfo object.UserBasisInfo
 	// 将字符串反解析为结构体
 	json.Unmarshal([]byte(userbinfoStr), &userbinfo)
 	fmt.Println(userbinfo)

 	this.Data["userbinfo"] = &userbinfo
 	this.Data["IsShowNav"] = "tool"
	this.TplName = "index.html"
}
