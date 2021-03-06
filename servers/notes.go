package servers

/*
	主要为笔记本，笔记分类等服务
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mangenotwork/ManGe-Notes/dao"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

type NotesServers struct{}

//创建新的笔记本
func (this *NotesServers) CreateNotes(datas *object.CNotes, uid string) (code int, count int, data string) {
	fmt.Println(datas)
	notesdatas := &models.Notes{
		NotesName:       datas.NotesName,
		UID:             uid,
		NotesDes:        datas.Description,
		NotesCreatetime: time.Now().Unix(),
	}
	//1. 检查笔记本命名是否存在
	if new(dao.DaoNotes).IsNotesName(datas.NotesName, uid) {
		new(dao.DaoNotes).AddNotes(notesdatas)
		return 1, 1, "创建成功"
	}
	return 2, 2, "笔记本命名已存在"
}

//获取所有笔记本
func (this *NotesServers) AllNotes(uid string) (code int, count int, data interface{}) {

	defnumber, _ := new(dao.DaoNotes).NotesNumber(uid, 0)
	fmt.Println("默认笔记本笔记数量=", defnumber)
	darnumber, _ := new(dao.DaoNotes).NotesNumber(uid, -2)
	recnumber, _ := new(dao.DaoNotes).NotesNumber(uid, -1)
	allnumber, _ := new(dao.DaoNotes).NotesAllNumber(uid)
	allnotes, err := new(dao.DaoNotes).GetNotesPgs(uid, 0, 20)
	if err != nil {
		fmt.Println(err)
		return 1, 1, "获取所有笔记本错误，后端错误"
	}

	rdata := &object.ReturnNotes{
		All:   allnumber,
		Def:   defnumber,
		Dar:   darnumber,
		Rec:   recnumber,
		Notes: allnotes,
	}

	return 1, 1, rdata
}

//管理模块 获取所有笔记本信息
func (this *NotesServers) GetAllNotesInfo(uid string) (count int, data interface{}) {
	//allnotes,err := new(models.Notes).GetNotesPgs(uid,0,20)
	//allnotes, err := new(dao.DaoNotes).GetNotesPgsInfo(uid, 0, 20)
	//获取所有笔记本
	allnotes, err := new(dao.DaoNotes).GetAllNotes(uid)
	if err != nil {
		fmt.Println(err)
		return 0, "获取所有笔记本错误，后端错误"
	}

	allNotesId := make([]int, 0)
	notesNameMap := make(map[int]string, 0)
	for _, v := range allnotes {
		allNotesId = append(allNotesId, v.NotesId)
		notesNameMap[v.NotesId] = v.NotesName
	}

	//获取所有笔记
	allnote, err := new(dao.DaoMDInof).GetData2Census(allNotesId)
	log.Println("allnotes = ", allnote, err)
	if err != nil {
		fmt.Println(err)
		return 0, "获取所有笔记本错误，后端错误"
	}
	log.Println(allnote)
	//记录笔记本笔记对应个数
	notesmap := make(map[string]int, 0)
	for _, v := range allnote {
		notes := notesNameMap[v.MDNotesid]
		notesmap[notes]++
	}
	log.Println("notesmap = ", notesmap)

	mangenotelist := make([]*object.MangeNotesList, 0)

	for _, v := range allnotes {
		mangenotelist = append(mangenotelist, &object.MangeNotesList{
			NotesID:    v.NotesId,
			NotesName:  v.NotesName,
			NotesDes:   v.NotesDes,
			CreateTime: time.Unix(v.NotesCreatetime, 0).Format("2006-01-02 15:04:05"),
			NoteNumber: notesmap[v.NotesName],
		})
	}

	fmt.Println(mangenotelist)

	return len(mangenotelist), mangenotelist
}

//修改笔记本信息
func (this *NotesServers) UpdateNotesInfo(datas *object.UpdateNotes, uid string) (code int, count int, data string) {
	err := new(dao.DaoNotes).UpdateInfo(datas, uid)
	if err != nil {
		fmt.Println("修改笔记信息失败", err)
		return 0, 1, "修改笔记信息失败"
	}
	return 1, 1, "修改笔记信息成功"
}

//删除笔记本
func (this *NotesServers) DeleteNotes(notesid string, uid string) (code int, count int, data string) {
	nid, _ := new(util.Str).NumberToInt(notesid)
	err := new(dao.DaoNotes).DeleteInfo(nid, uid)
	if err != nil {
		fmt.Println("修改笔记信息失败", err)
		return 0, 1, "修改笔记信息失败"
	}

	return 1, 1, "修改笔记信息成功"
}

//NotesChartData  图表模块 获取笔记本笔记数量分部
func (this *NotesServers) NotesChartData(uid string) (code int, namelist []string, data interface{}) {
	//获取所有笔记本
	allnotes, err := new(dao.DaoNotes).GetAllNotes(uid)
	log.Println("allnotes = ", allnotes, err)
	if err != nil {
		fmt.Println(err)
		return 0, make([]string, 0), "获取所有笔记本错误，后端错误"
	}

	allNotesId := make([]int, 0)       //记录笔记本id
	notesname := make(map[int]string)  //记录笔记本id 对应笔记本名
	notesNameList := make([]string, 0) //记录笔记本名列表
	for _, notesdata := range allnotes {
		allNotesId = append(allNotesId, notesdata.NotesId)
		notesname[notesdata.NotesId] = notesdata.NotesName
		notesNameList = append(notesNameList, notesdata.NotesName)
	}
	log.Println("allNotesId = ", allNotesId)

	//获取所有笔记
	allnote, err := new(dao.DaoMDInof).GetData2Census(allNotesId)
	log.Println("allnotes = ", allnote, err)
	if err != nil {
		fmt.Println(err)
		return 0, make([]string, 0), "获取所有笔记本错误，后端错误"
	}
	log.Println(allnote)

	//记录笔记本笔记对应个数
	notesmap := make(map[string]int, 0)
	for _, v := range allnote {
		notes := notesname[v.MDNotesid]
		notesmap[notes]++
	}
	log.Println("notesmap = ", notesmap)

	//整理输出
	notesdata := make([]*object.NotesCount, 0)
	for k, v := range notesmap {
		notesdata = append(notesdata, &object.NotesCount{
			NotesName: k,
			Number:    v,
		})
	}

	//输出结果序列化成json
	jsondata, err := json.Marshal(notesdata)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	return 1, notesNameList, string(jsondata)
}

//MyChartData  图表模块 我的综合统计
func (this *NotesServers) MyChartData(uid string) {
	// 1.需要记录每天创建笔记的增量
	// 2.需要记录每天修改笔记的增量
	// 3.需要记录每天查看笔记的增量
	// 4.需要记录每天收藏链接的增量
	// 5.需要记录每天上传素材的增量
	fmt.Println("还未开发")

}

//UsedSpace  图表模块 我的使用空间
func (this *NotesServers) UsedSpace(uid string) (data interface{}) {
	//笔记总数
	notecount, noteErr := new(dao.DaoMDInof).GetNoteCount(uid)
	fmt.Println(notecount, noteErr)
	//素材存储空间
	sizeSpace, sizeErr := new(dao.DaoIMGInfo).GetSize(uid)
	fmt.Println(sizeSpace, sizeErr)
	var sizeValue float64
	sizeValue = float64(sizeSpace) / 1024 / 1024
	sizeStr := fmt.Sprintf("%fM", sizeValue)
	//笔记本数量
	notesNumber, notesErr := new(dao.DaoNotes).GetNotesCount(uid)
	fmt.Println(notesNumber, notesErr)

	fmt.Println(int((float32(notecount) / 100) * 100))
	fmt.Println(int((float32(sizeSpace) / (20 * 1024 * 1024)) * 100))
	fmt.Println(int((float32(notesNumber) / 20) * 100))

	data = &object.UsedInfo{
		NoteMax:      100,
		NoteNow:      notecount,
		NotePercent:  int64((float32(notecount) / 100) * 100),
		SpaceMax:     "20M",
		SpaceNow:     sizeStr,
		SpacePercent: int64((float32(sizeSpace) / (20 * 1024 * 1024)) * 100),
		NotesMax:     20,
		NotesNow:     notesNumber,
		NotesPercent: int64((float32(notesNumber) / 20) * 100),
	}
	return
}
