package servers

/*
	主要为Markdown类型笔记内容服务
*/

import (
	_ "encoding/json"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/mangenotwork/ManGe-Notes/dao"

	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

type MDServers struct{}

func (this *MDServers) CreateMDInfo(datas *object.CMDData) (mdid string, mdimglink string, mdisimg int, destxt string) {

	mdTxt := new(util.Str).Escape(datas.Detail)
	//创建笔记id
	id, uidErr := util.Int64ID()
	if uidErr != nil {
		fmt.Println("生成 int64 ID错误")
	}
	mdid = fmt.Sprintf("md_%d", id)

	//判断笔记内容是否存在图片链接
	mdimglink = new(util.Str).GetMDImgLink(mdTxt)

	mdisimg = 0
	if mdimglink != "" {
		mdisimg = 1
	}

	fmt.Println("ASCII 字符串长度", len(mdTxt))
	fmt.Println("Unicode 字符串长度", utf8.RuneCountInString(mdTxt))
	if len(mdTxt) > 30 {
		//destxt = new(util.Str).RepMDDesc(datas.Detail[0 : 30],30)
		destxt = new(util.Str).RepMDDesc(util.ShowSubstr(mdTxt, 30), 30)
	} else {
		destxt = new(util.Str).RepMDDesc(util.ShowSubstr(mdTxt, len(mdTxt)-1), len(mdTxt)-1)
	}

	return
}

//创建MD笔记
func (this *MDServers) CreateMDNote(datas *object.CMDData, uid string) (code int, count int, data string) {
	fmt.Println(datas)
	mdid, mdimglink, mdisimg, destxt := this.CreateMDInfo(datas)
	nowtime := time.Now().Unix()
	tagsinfo := new(util.Str).Strip(datas.Tags, " \n")
	notesid, _ := new(util.Str).NumberToInt(datas.NotesId)

	mdTxt := new(util.Str).Escape(datas.Detail)

	mdinfos := &models.MDInof{
		MDTitle:      datas.Title,
		MDDes:        destxt,
		MDIMG:        mdimglink,
		IsIMG:        mdisimg,
		MDId:         mdid,
		Uid:          uid,
		MDNotesid:    notesid,
		MDCreatetime: nowtime,
		MDSavetime:   nowtime,
		MDTag:        tagsinfo,
	}
	//存储笔记和笔记信息
	mdText := &models.MDText{
		MDId:      mdid,
		MDContent: mdTxt,
	}

	err := new(dao.DaoMDInof).InsertMDNote(mdText, mdinfos)
	if err != nil {
		fmt.Println("创建笔记错误", err)
		return 0, 1, "创建笔记错误"
	}
	return 1, 1, "创建笔记成功"
	//如果创建成功将dm笔记缓存到redis
}

func (this *MDServers) CreateMDToDraft(datas *object.CMDData, uid string) (code int, count int, data string) {
	mdid, mdimglink, mdisimg, destxt := this.CreateMDInfo(datas)
	nowtime := time.Now().Unix()
	tagsinfo := new(util.Str).Strip(datas.Tags, " \n")
	notesid := -2
	mdTxt := new(util.Str).Escape(datas.Detail)

	//如果存的草稿没有设置 title 那么就给添加一个
	timeStr := time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")
	title := fmt.Sprintf("%s保存的草稿", timeStr)

	mdinfos := &models.MDInof{
		MDTitle:      title,
		MDDes:        destxt,
		MDIMG:        mdimglink,
		IsIMG:        mdisimg,
		MDId:         mdid,
		Uid:          uid,
		MDNotesid:    notesid,
		MDCreatetime: nowtime,
		MDSavetime:   nowtime,
		MDTag:        tagsinfo,
	}
	fmt.Println(mdinfos)
	//存储笔记和笔记信息
	mdText := &models.MDText{
		MDId:      mdid,
		MDContent: mdTxt,
	}

	err := new(dao.DaoMDInof).InsertMDNote(mdText, mdinfos)
	if err != nil {
		fmt.Println("创建笔记错误", err)
		return 0, 1, "创建笔记错误"
	}
	return 1, 1, "创建笔记成功"
	//如果创建成功将dm笔记缓存到redis
}

//修改笔记内容
func (this *MDServers) ModifyMDNote(datas *object.CMDData, uid string, mdid string) (code int, count int, data string) {
	//1. 权限判断
	ismd, _, err := new(dao.DaoMDInof).IsMD(uid, mdid)
	if err != nil && err.Error() != "record not found" {
		fmt.Println("判断查看MD权限错误，错误信息:", err)
		return 0, 1, "判断查看MD权限错误，错误信息"
	}
	if ismd {

		mdTxt := new(util.Str).Escape(datas.Detail)
		_, mdimglink, mdisimg, destxt := this.CreateMDInfo(datas)

		mdinfos := &models.MDInof{
			MDTitle: datas.Title,
			MDDes:   destxt,
			MDIMG:   mdimglink,
			IsIMG:   mdisimg,
		}

		//3. 修改笔记内容
		fmt.Println("\n\n\n****************修改笔记内容", mdTxt, "\n\n\n**************************")
		updateerr := new(dao.DaoMDInof).UpdateMDNote(mdinfos, mdTxt, mdid)
		if updateerr != nil {
			fmt.Println("修改笔记错误:", updateerr)
			return 0, 1, "修改笔记错误"
		}
		return 1, 1, "修改笔记成功"
	}
	return 2, 1, "你没有权限修改笔记"

}

//将获取到的笔记信息列表处理后的结构形式再返回给接口使用
func (this *MDServers) NoteListFormat(datas []*models.MDInof) (returndatas []*object.ReturnNoteInfo) {
	returndatas = make([]*object.ReturnNoteInfo, 0)
	for _, v := range datas {

		fmt.Println(v, v.MDTitle)
		destxt := new(util.Str).ToNbsp(v.MDDes)
		noteinfo := &object.ReturnNoteInfo{
			Title:       v.MDTitle,
			Des:         destxt,
			IsImg:       v.IsIMG,
			ImgLink:     v.MDIMG,
			Id:          v.MDId,
			Savetime:    time.Unix(v.MDSavetime, 0).Format("2006-01-02 15:04:05"),
			Tags:        v.MDTag,
			ViewTimes:   v.MDViewTimes,
			Modifytimes: v.MDModifytimes,
		}
		fmt.Println(noteinfo)
		returndatas = append(returndatas, noteinfo)

	}
	return
}

//获取所有笔记
func (this *MDServers) GetAllNote(pg int, uid string) (code int, count int, data interface{}) {
	pgsize := (pg - 1) * 20
	datas, err := new(dao.DaoMDInof).GetToPG(pgsize, 20, uid)
	if err != nil {
		fmt.Println("获取所有笔记错误，后台错误", err)
		return 0, 1, "获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)

	return 0, 1, rdata
}

//获取回收站的笔记
func (this *MDServers) GetRecyclerNote(pg int, uid string) (code int, count int, data interface{}) {
	pgsize := (pg - 1) * 20
	datas, err := new(dao.DaoMDInof).GetRecycler(pgsize, 20, uid)
	if err != nil {
		fmt.Println("获取所有笔记错误，后台错误", err)
		return 0, 1, "获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)

	return 0, 1, rdata
}

//DraftNote 获取草稿笔记
func (this *MDServers) DraftNote(pg int, uid string) (code int, count int, data interface{}) {
	pgsize := (pg - 1) * 20
	datas, err := new(dao.DaoMDInof).DraftNote(pgsize, 20, uid)
	if err != nil {
		fmt.Println("获取所有笔记错误，后台错误", err)
		return 0, 1, "获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)

	return 0, 1, rdata
}

//GetNoteList 获取笔记本下的笔记
func (this *MDServers) GetNoteList(notesid int, pg int, uid string) (code int, count int, data interface{}) {
	pgsize := (pg - 1) * 20
	datas, err := new(dao.DaoMDInof).GetNotesToPG(pgsize, 20, uid, notesid)
	if err != nil {
		fmt.Println("获取笔记本笔记列表错误，后台错误", err)
		return 0, 1, "获取笔记本笔记列表错误"
	}
	rdata := this.NoteListFormat(datas)
	return 0, 1, rdata
}

//通过用户id mdid查找MD笔记内容
func (this *MDServers) GetMDContent(uid string, mid string, types int) (code int, count int, data interface{}, title string) {
	//1. 判断uid用户是否拥有这个mid,如果拥有则返回md内容，不拥有则提示你没有权限访问此MD内容
	ismd, title, err := new(dao.DaoMDInof).IsMD(uid, mid)
	if err != nil && err.Error() != "record not found" {
		fmt.Println("判断查看MD权限错误，错误信息:", err)
		return 0, 1, "判断查看MD权限错误，错误信息", ""
	}
	if ismd {
		mdtxt, err := new(dao.DaoMDText).GetMDTxt(mid)
		if err != nil {
			fmt.Println("查询MD内容错误，错误信息:", err)
			return 0, 1, "查询MD内容错误，错误信息", ""
		}

		if types == 1 {
			//2. 增加查看次数
			go new(dao.DaoMDInof).AddMDViewTimes(mid)
		} else if types == 2 {
			//2. 增加修改次数
			go new(dao.DaoMDInof).AddMDModifytimes(mid)
		}

		return 1, 1, mdtxt, title
	}
	fmt.Println("你没有权限查看此笔记")
	return 2, 1, "你没有权限查看此笔记", ""
}

//通过笔记名模糊搜索查询笔记
func (this *MDServers) SearchNoteinfo(word, uid string) (code int, count int, data interface{}) {
	datas, err := new(dao.DaoMDInof).SearchTitle(word, uid)
	if err != nil {
		fmt.Println("获取所有笔记错误，后台错误", err)
		return 0, 1, "获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)

	return 0, 1, rdata
}

//删除笔记到回收站
func (this *MDServers) DeleteNote(mdid, uid string) (code int, count int, data interface{}) {
	ismd, _, err := new(dao.DaoMDInof).IsMD(uid, mdid)
	if err != nil && err.Error() != "record not found" {
		fmt.Println("判断查看MD权限错误，错误信息:", err)
		return 0, 1, "判断查看MD权限错误，错误信息"
	}
	if ismd {
		err := new(dao.DaoMDInof).ToDEL(mdid, uid)
		if err != nil {
			fmt.Println("删除笔记错误", err)
			return 0, 1, "删除笔记错误"
		}
		return 1, 1, "删除笔记成功"
	}
	return 2, 1, "你没有权限删除此笔记"
}

//SchenNote 永久删除笔记
func (this *MDServers) SchenNote(mdid, uid string) (code int, count int, data interface{}) {
	ismd, _, err := new(dao.DaoMDInof).IsMD(uid, mdid)
	if err != nil && err.Error() != "record not found" {
		fmt.Println("判断查看MD权限错误，错误信息:", err)
		return 0, 1, "判断查看MD权限错误，错误信息"
	}
	if ismd {
		err := new(dao.DaoMDInof).Schen(mdid, uid)
		if err != nil {
			fmt.Println("删除笔记错误", err)
			return 0, 1, "删除笔记错误"
		}
		return 1, 1, "删除笔记成功"
	}
	return 2, 1, "你没有权限删除此笔记"
}

//恢复笔记到笔记本
func (this *MDServers) RestoreToNotes(mdid, uid string, notes int) (code int, count int, data interface{}) {
	err := new(dao.DaoMDInof).NoteToNotes(mdid, uid, notes)
	if err != nil {
		fmt.Println("恢复笔记错误", err)
		return 0, 1, "恢复笔记错误"
	}
	return 1, 1, "恢复笔记成功"
}
