package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/servers"
)

type PGController struct {
	Controller
}

//登录页面
func (this *PGController) Login() {
	this.TplName = "pg/login.html"
}

func (this *PGController) GetIsLogin() {
	if this.IsLogin() {
		this.RetuenJson(0, 0, "")
	} else {
		this.RetuenJson(1, 0, "")
	}

}

//Install  安装页面
func (this *PGController) Install() {
	step := 1
	//检查安装文件
	isInstall := FileExist(object.InstallJsonPath)
	if isInstall {
		//读取文件
		installInfo := object.OpenInstallFile()
		log.Println("masterconf = ", installInfo)
		step = installInfo.Step
	}
	this.Data["step"] = step
	this.TplName = "pg/install.html"
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//主页
func (this *PGController) IndexPG() {

	this.IsSession()

	//获取用户基本信息
	uid := this.GetUid()

	if uid == "" {
		this.Redirect("/", 301)
		return
	}

	fmt.Println(fmt.Sprintf("uinfo:%s", uid))
	//TODO 在缓冲获取用户信息
	userbinfoStr := "" //new(rdb.RDB).HashGet(fmt.Sprintf("uinfo:%s", uid), "basis")
	fmt.Println(userbinfoStr)

	var userbinfo object.UserBasisInfo
	// 将字符串反解析为结构体
	json.Unmarshal([]byte(userbinfoStr), &userbinfo)
	fmt.Println(userbinfo)

	this.Data["userbinfo"] = &userbinfo
	this.Data["IsShowNav"] = "index"

	this.TplName = "index.html"
	return
}

//MD编辑器
func (this *PGController) MdEditorPG() {
	this.Data["IsLogin"] = this.IsLogin()
	this.TplName = "pg/createmd.html"
}

//首页
func (this *PGController) HomePG() {
	this.Data["IsLogin"] = this.IsLogin()
	this.TplName = "pg/home.html"
}

//PC端页面渲染
func (this *PGController) ToolPG() {

	//获取用户基本信息
	uid := this.GetUid()
	fmt.Println(fmt.Sprintf("uinfo:%s", uid))
	//TODO 在缓存获取用户信息
	userbinfoStr := "" //new(rdb.RDB).HashGet(fmt.Sprintf("uinfo:%s", uid), "basis")
	fmt.Println(userbinfoStr)

	var userbinfo object.UserBasisInfo
	// 将字符串反解析为结构体
	json.Unmarshal([]byte(userbinfoStr), &userbinfo)
	fmt.Println(userbinfo)

	this.Data["userbinfo"] = &userbinfo
	this.Data["IsShowNav"] = "tool"
	this.TplName = "index.html"
}

//MangeNotes  笔记本管理
func (this *PGController) MangeNotes() {
	this.TplName = "pg/mange_notes.html"
}

//MangeLinks 收藏链接管理
func (this *PGController) MangeLinks() {
	this.TplName = "pg/mange_links.html"
}

// ChartNotes  图表模块  笔记本数量分布图
func (this *PGController) ChartNotes() {
	uid := this.GetUid()

	code, namelist, data := new(servers.NotesServers).NotesChartData(uid)
	log.Println(code, namelist, data)

	this.Data["Code"] = code
	this.Data["Name"] = namelist
	this.Data["Data"] = data
	this.TplName = "chart/notes.html"
}

//MyChart 图表模块  我的综合统计
func (this *PGController) MyChart() {
	uid := this.GetUid()

	//code,namelist,data := new(servers.NotesServers).NotesChartData(uid)
	new(servers.NotesServers).MyChartData(uid)

	this.Data["Code"] = uid
	this.TplName = "chart/zhonghe.html"
}

//MyUsedSpace 图表模块 我的使用空间
func (this *PGController) MyUsedSpace() {
	uid := this.GetUid()
	log.Println("uid  = ", uid)

	//code,namelist,data := new(servers.NotesServers).NotesChartData(uid)
	data := new(servers.NotesServers).UsedSpace(uid)
	fmt.Println("data = ", data)

	this.Data["Data"] = data
	this.TplName = "chart/main.html"
}

//Shequ  漫鸽笔记社区主页
func (this *PGController) Shequ() {

	this.TplName = "community/index.html"
}

//素材 模块 主页
func (this *PGController) SuCai() {
	uid := this.GetUid()
	if uid == "" {
		this.TplName = "pg/login.html"
		return
	}

	code, count, data := new(servers.SUCai).GetMyImg(uid)

	this.Data["My"] = true
	this.Data["Code"] = code
	this.Data["Count"] = count
	this.Data["Data"] = data
	this.TplName = "sucai/sc.html"
}

//AddSuCai  素材模块 添加素材
func (this *PGController) AddSuCai() {
	this.TplName = "sucai/addsc.html"
}
