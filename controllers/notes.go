package controllers

import (
	"fmt"
	"log"

	object "github.com/mangenotwork/ManGe-Notes/object"
	servers "github.com/mangenotwork/ManGe-Notes/servers"
)

type NotesController struct {
	Controller
}

//创建笔记本
func (this *NotesController) CreateNotes() {
	uid := this.GetUid()

	var obj object.CNotes

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	code, count, data := new(servers.NotesServers).CreateNotes(&obj, uid)
	this.RetuenJson(code, count, data)
}

//获取当前笔记本列表
func (this *NotesController) GetNotesList() {
	log.Println("GetNotesList")
	uid := this.GetUid()
	log.Println("uid = ", uid)
	if uid == "" {
		this.RetuenJson(-1, 0, "")
		return
	}
	log.Println("AllNotes")
	code, count, data := new(servers.NotesServers).AllNotes(uid)
	this.RetuenJson(code, count, data)
}

//mange管理模块 获取所有笔记本信息
func (this *NotesController) GetAllNotes() {
	uid := this.GetUid()
	count, data := new(servers.NotesServers).GetAllNotesInfo(uid)
	this.MangeJson(0, "", count, data)
}

//mange管理模块 修改笔记本信息 UpdateNotesInfo
func (this *NotesController) UpdateNotesInfo() {
	uid := this.GetUid()
	var obj object.UpdateNotes

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	code, count, data := new(servers.NotesServers).UpdateNotesInfo(&obj, uid)
	this.RetuenJson(code, count, data)
}

//删除笔记本  DelNotes
func (this *NotesController) DelNotes() {
	uid := this.GetUid()
	notesid := this.Ctx.Input.Param(":notesid")
	code, count, data := new(servers.NotesServers).DeleteNotes(notesid, uid)
	this.RetuenJson(code, count, data)
}
