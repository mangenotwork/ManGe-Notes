package servers
/*
	主要为笔记本，笔记分类等服务
*/
import (
	"fmt"
	"time"
	_ "encoding/json"

	object "man/ManNotes/object"
	//util "man/ManNotes/util"
	models "man/ManNotes/models"
	//rdb "man/ManNotes/models/redis"
)

type NotesServers struct {}

//创建新的笔记本
func (this *NotesServers) CreateNotes(datas *object.CNotes) (code int, count int, data string) {
	fmt.Println(datas)
	notesdatas := &models.Notes{
		NotesName : datas.NotesName,
		NotesDes : datas.Description,
		NotesCreatetime : time.Now().Unix(),
	}
	//1. 检查笔记本命名是否存在
	if notesdatas.IsNotesName(datas.NotesName){
		notesdatas.AddNotes()
		return 1,1,"创建成功"
	}
	return 2,2,"笔记本命名已存在"
}

//获取所有笔记本
func (this *NotesServers) AllNotes() (code int, count int, data interface{}) {
	allnotes,err := new(models.Notes).GetNotesPgs(0,20)
	if err != nil {
		fmt.Println(err)
		return 1,1,"获取所有笔记本错误，后端错误"
	}
	return 1,1,allnotes
}