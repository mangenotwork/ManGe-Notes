package controllers

import (
	"fmt"

	object "github.com/mangenotwork/ManGe-Notes/object"
	servers "github.com/mangenotwork/ManGe-Notes/servers"
	util "github.com/mangenotwork/ManGe-Notes/util"
)

type MDController struct {
	Controller
}

//创建MD笔记
func (this *MDController) CreateMD() {
	uid := this.GetUid()

	var obj object.CMDData

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	code, count, data := new(servers.MDServers).CreateMDNote(&obj, uid)

	this.RetuenJson(code, count, data)
}

//笔记保存到草稿
func (this *MDController) ToDraft() {
	uid := this.GetUid()

	var obj object.CMDData

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	code, count, data := new(servers.MDServers).CreateMDToDraft(&obj, uid)

	this.RetuenJson(code, count, data)
}

//获取所有笔记
func (this *MDController) GetAllNote() {
	uid := this.GetUid()
	if uid == "" {
		this.RetuenJson(-1, 0, "")
		return
	}
	pg, err := this.GetInt("pg")
	fmt.Println(pg, uid)
	if err != nil || pg == 0 {
		fmt.Println("获取页数失败")
		pg = 1
	}
	code, count, data := new(servers.MDServers).GetAllNote(pg, uid)
	this.RetuenJson(code, count, data)
}

//MDShow  显示MD笔记内容
func (this *MDController) MDShow() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	fmt.Println(uid, mdid)
	code, count, data, title := new(servers.MDServers).GetMDContent(uid, mdid, 1) //1 查看请求的内容
	fmt.Println(code, count)

	//this.RetuenJson(1,1,1)
	this.Data["Error"] = code
	this.Data["MDText"] = data
	this.Data["MDID"] = mdid
	this.Data["MDTitle"] = title
	this.TplName = "pg/mdshow.html"
}

//显示回收站里的笔记内容
func (this *MDController) RMDShow() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	code, _, data, title := new(servers.MDServers).GetMDContent(uid, mdid, 1) //1 查看请求的内容

	this.Data["Error"] = code
	this.Data["MDText"] = data
	this.Data["MDID"] = mdid
	this.Data["MDTitle"] = title
	this.TplName = "pg/delmdshow.html"
}

//MDEdit  修改MD笔记内容
func (this *MDController) MDEditPG() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	fmt.Println(mdid)
	code, count, data, title := new(servers.MDServers).GetMDContent(uid, mdid, 2) //2 修改页面请求的内容
	fmt.Println(code, count)

	this.Data["MDTitle"] = title
	this.Data["MDText"] = data
	this.Data["MDID"] = mdid
	this.TplName = "pg/editmd.html"
}

//NotesMDList  笔记本对应的笔记列表
func (this *MDController) NotesMDList() {
	uid := this.GetUid()
	//notesid := this.Ctx.Input.Param(":id")
	notesid, interr := new(util.Str).NumberToInt(this.Ctx.Input.Param(":id"))
	if interr != nil {
		this.RetuenJson(0, 1, "后台错误数字字符串转Int错误")
		return
	}

	pg, err := this.GetInt("pg")
	fmt.Println(pg, uid)
	if err != nil || pg == 0 {
		fmt.Println("获取页数失败")
		pg = 1
	}

	code, count, data := new(servers.MDServers).GetNoteList(notesid, pg, uid)
	this.RetuenJson(code, count, data)

}

//笔记内容提交修改
func (this *MDController) MDNoteModify() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	var obj object.CMDData

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
	fmt.Println(mdid, uid)
	code, count, data := new(servers.MDServers).ModifyMDNote(&obj, uid, mdid)

	this.RetuenJson(code, count, data)
}

//搜索笔记
func (this *MDController) SearchNote() {
	uid := this.GetUid()
	word := this.GetString("word")

	code, count, data := new(servers.MDServers).SearchNoteinfo(word, uid)
	this.RetuenJson(code, count, data)
}

//删除笔记
func (this *MDController) DelNote() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	code, count, data := new(servers.MDServers).DeleteNote(mdid, uid)
	this.RetuenJson(code, count, data)
}

//SchenNote 永久删除笔记
func (this *MDController) SchenNote() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	code, count, data := new(servers.MDServers).SchenNote(mdid, uid)
	this.RetuenJson(code, count, data)
}

//NoteRecycler 回收站
func (this *MDController) NoteRecycler() {
	uid := this.GetUid()
	pg, err := this.GetInt("pg")
	fmt.Println(pg, uid)
	if err != nil || pg == 0 {
		fmt.Println("获取页数失败")
		pg = 1
	}
	code, count, data := new(servers.MDServers).GetRecyclerNote(pg, uid)
	this.RetuenJson(code, count, data)
}

//DraftList  草稿
func (this *MDController) DraftList() {
	uid := this.GetUid()
	pg, err := this.GetInt("pg")
	fmt.Println(pg, uid)
	if err != nil || pg == 0 {
		fmt.Println("获取页数失败")
		pg = 1
	}
	code, count, data := new(servers.MDServers).DraftNote(pg, uid)
	this.RetuenJson(code, count, data)
}

//恢复到笔记本
func (this *MDController) RestoreNote() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	notes, err := this.GetInt("notes")
	if err != nil {
		notes = 0
	}

	code, count, data := new(servers.MDServers).RestoreToNotes(mdid, uid, notes)
	this.RetuenJson(code, count, data)
}

//DraNoteShow  显示草稿笔记
func (this *MDController) DraNoteShow() {
	uid := this.GetUid()
	mdid := this.Ctx.Input.Param(":mdid")
	code, _, data, title := new(servers.MDServers).GetMDContent(uid, mdid, 1) //1 查看请求的内容

	this.Data["Error"] = code
	this.Data["MDText"] = data
	this.Data["MDID"] = mdid
	this.Data["MDTitle"] = title
	this.TplName = "pg/dramdshow.html"
}
