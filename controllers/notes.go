package controllers

import (
	"fmt"

	object "man/ManNotes/object"
	servers "man/ManNotes/servers"
)

type NotesController struct {
	Controller
}

//创建笔记本
func (this *NotesController) CreateNotes(){
	this.IsSession()

	var obj object.CNotes
	
	err := this.ResolvePostData(&obj)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(obj)

	code,count,data := new(servers.NotesServers).CreateNotes(&obj)
	this.RetuenJson(code,count,data)
}

//获取当前笔记本列表
func (this *NotesController) GetNotesList(){
	this.IsSession()
	code,count,data := new(servers.NotesServers).AllNotes()
	this.RetuenJson(code,count,data)
}
