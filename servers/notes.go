package servers
/*
	主要为笔记本，笔记分类等服务
*/
import (
	"fmt"
	"time"
	_ "encoding/json"

	object "man/ManNotes/object"
	util "man/ManNotes/util"
	models "man/ManNotes/models"
	//rdb "man/ManNotes/models/redis"
)

type NotesServers struct {}

//创建新的笔记本
func (this *NotesServers) CreateNotes(datas *object.CNotes,uid string) (code int, count int, data string) {
	fmt.Println(datas)
	notesdatas := &models.Notes{
		NotesName : datas.NotesName,
		UID : uid,
		NotesDes : datas.Description,
		NotesCreatetime : time.Now().Unix(),
	}
	//1. 检查笔记本命名是否存在
	if notesdatas.IsNotesName(datas.NotesName,uid){
		notesdatas.AddNotes()
		return 1,1,"创建成功"
	}
	return 2,2,"笔记本命名已存在"
}



//获取所有笔记本
func (this *NotesServers) AllNotes(uid string) (code int, count int, data interface{}) {

	defnumber,_ := new(models.Notes).NotesNumber(uid,0)
	fmt.Println("默认笔记本笔记数量=",defnumber)
	darnumber,_ := new(models.Notes).NotesNumber(uid,-2)
	recnumber,_ := new(models.Notes).NotesNumber(uid,-1)
	allnumber,_ := new(models.Notes).NotesAllNumber(uid)

	allnotes,err := new(models.Notes).GetNotesPgs(uid,0,20)
	if err != nil {
		fmt.Println(err)
		return 1,1,"获取所有笔记本错误，后端错误"
	}

	rdata := &object.ReturnNotes{
		All : allnumber,
		Def : defnumber,
		Dar : darnumber,
		Rec : recnumber,
		Notes : allnotes,
	}

	return 1,1,rdata
}

//管理模块 获取所有笔记本信息
func (this *NotesServers) GetAllNotesInfo(uid string) (count int,data interface{}) {
	//allnotes,err := new(models.Notes).GetNotesPgs(uid,0,20)
	allnotes,err := new(models.Notes).GetNotesPgsInfo(uid,0,20)
	if err != nil {
		fmt.Println(err)
		return 0,"获取所有笔记本错误，后端错误"
	}

	mangenotelist := make([]*object.MangeNotesList,0)

	for _,k := range allnotes{
		mangenotelist = append(mangenotelist,&object.MangeNotesList{
			NotesID : k.NotesId,
			NotesName : k.NotesName,
			NotesDes : k.NotesDes,
			CreateTime : time.Unix(k.NotesCreatetime, 0).Format("2006-01-02 15:04:05"),
			NoteNumber : k.N,
			})
	}

	fmt.Println(mangenotelist)
	
	return len(mangenotelist),mangenotelist
}

//修改笔记本信息
func (this *NotesServers) UpdateNotesInfo(datas *object.UpdateNotes,uid string) (code int, count int, data string) {
	err := new(models.Notes).UpdateInfo(datas,uid)
	if err != nil {
		fmt.Println("修改笔记信息失败",err)
		return 0,1,"修改笔记信息失败"
	}
	return 1,1,"修改笔记信息成功"
}

//删除笔记本
func (this *NotesServers) DeleteNotes(notesid string,uid string) (code int, count int, data string) {
	nid,_ := new(util.Str).NumberToInt(notesid)
	err := new(models.Notes).DeleteInfo(nid,uid)
	if err != nil {
		fmt.Println("修改笔记信息失败",err)
		return 0,1,"修改笔记信息失败"
	}

	return 1,1,"修改笔记信息成功"
}