package servers
/*
	主要为Markdown类型笔记内容服务
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

type MDServers struct {}

//创建MD笔记
func (this *MDServers) CreateMDNote(datas *object.CMDData, uid string) (code int, count int, data string) {
	fmt.Println(datas)
	//创建笔记id
	id,uidErr := util.Int64ID()
	if uidErr != nil{
		fmt.Println("生成 int64 ID错误")
		return 0,1,"生成 int64 ID错误"
	}
	mdid := fmt.Sprintf("md_%d",id)
	nowtime := time.Now().Unix()
	tagsinfo := new(util.Str).Strip(datas.Tags," \n")

	//判断笔记内容是否存在图片链接
	mdimglink := new(util.Str).GetMDImgLink(datas.Detail)
	mdisimg := 0
	if mdimglink != ""{
		mdisimg = 1
	}

	var destxt string
	fmt.Println(len(datas.Detail))
	if len(datas.Detail) > 30{
		destxt = new(util.Str).RepMDDesc(datas.Detail[0 : 30],30)
	}else{
		destxt = new(util.Str).RepMDDesc(datas.Detail[0 : len(datas.Detail)-1],len(datas.Detail)-1)
	}
	
	
	notesid,_ := new(util.Str).NumberToInt(datas.NotesId)

	mdinfos := &models.MDInof{
		MDTitle : datas.Title,
		MDDes : destxt,
		MDIMG : mdimglink,
		IsIMG : mdisimg,
		MDId : mdid,
		Uid : uid,
		MDNotesid : notesid,
		MDCreatetime : nowtime,
		MDSavetime : nowtime,
		MDTag :tagsinfo,
	}
	//存储笔记和笔记信息
	mdText := &models.MDText{
		MDId: mdid,
		MDContent:datas.Detail,
	}

	err := mdinfos.InsertMDNote(mdText)
	if err != nil {
		fmt.Println("创建笔记错误",err)
		return 0,1,"创建笔记错误"
	}
	return 1,1,"创建笔记成功"
	//如果创建成功将dm笔记缓存到redis
}


//修改笔记内容
func (this *MDServers) ModifyMDNote(datas *object.CMDData,uid string,mdid string) (code int, count int, data string) {
	//1. 权限判断
	ismd,_,err := new(models.MDInof).IsMD(uid,mdid)
	if err != nil && err.Error() != "record not found"{
		fmt.Println("判断查看MD权限错误，错误信息:",err)
		return 0,1,"判断查看MD权限错误，错误信息"
	}
	if ismd {
		//2. 修改笔记的描述 取笔记内容前30个字符作为笔记的描述内容
		var destxt string
		if len(datas.Detail) > 30{
			destxt = new(util.Str).RepMDDesc(datas.Detail[0 : 30],30)
		}else{
			destxt = new(util.Str).RepMDDesc(datas.Detail[0 : len(datas.Detail)-1],len(datas.Detail)-1)
		}

		//判断笔记内容是否存在图片链接
		mdimglink := new(util.Str).GetMDImgLink(datas.Detail)
		mdisimg := 0
		if mdimglink != ""{
			mdisimg = 1
		}

		mdinfos := &models.MDInof{
			MDTitle : datas.Title,
			MDDes : destxt,
			MDIMG : mdimglink,
			IsIMG : mdisimg,
		}

		//3. 修改笔记内容
		updateerr := mdinfos.UpdateMDNote(datas.Detail,mdid)
		if updateerr != nil{
			fmt.Println("修改笔记错误:",updateerr)
			return 0,1,"修改笔记错误"
		}
		return 1,1,"修改笔记成功"
	}
	return 2,1,"你没有权限修改笔记"
	
}


//将获取到的笔记信息列表处理后的结构形式再返回给接口使用
func (this *MDServers) NoteListFormat(datas []*models.MDInof) (returndatas []*object.ReturnNoteInfo) {
	returndatas = make([]*object.ReturnNoteInfo, 0) 
	for _,v := range datas{

		fmt.Println(v,v.MDTitle)
		destxt := new(util.Str).ToNbsp(v.MDDes)
		noteinfo := &object.ReturnNoteInfo{
			Title : v.MDTitle,
			Des : destxt,
			IsImg : v.IsIMG,
			ImgLink : v.MDIMG,
			Id :  v.MDId,
			Savetime :  time.Unix(v.MDSavetime, 0).Format("2006-01-02 15:04:05"),
			Tags :  v.MDTag,
			ViewTimes :  v.MDViewTimes,
			Modifytimes :  v.MDModifytimes,
			}
		fmt.Println(noteinfo)
		returndatas = append(returndatas, noteinfo)
		
	}
	return
}

//获取所有笔记
func (this *MDServers) GetAllNote(pg int,uid string) (code int, count int, data interface{}) {
	pgsize := (pg-1)*20
	datas,err := new(models.MDInof).GetToPG(pgsize,20,uid)
	if err != nil{
		fmt.Println("获取所有笔记错误，后台错误",err)
		return 0,1,"获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)
	
	return 0,1,rdata
}

//获取回收站的笔记
func (this *MDServers) GetRecyclerNote(pg int,uid string) (code int, count int, data interface{}) {
	pgsize := (pg-1)*20
	datas,err := new(models.MDInof).GetRecycler(pgsize,20,uid)
	if err != nil{
		fmt.Println("获取所有笔记错误，后台错误",err)
		return 0,1,"获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)
	
	return 0,1,rdata
}

//GetNoteList 获取笔记本下的笔记
func (this *MDServers) GetNoteList(notesid int,pg int,uid string) (code int, count int, data interface{}) {
	pgsize := (pg-1)*20
	datas,err := new(models.MDInof).GetNotesToPG(pgsize,20,uid,notesid)
	if err != nil{
		fmt.Println("获取笔记本笔记列表错误，后台错误",err)
		return 0,1,"获取笔记本笔记列表错误"
	}
	rdata := this.NoteListFormat(datas)
	return 0,1,rdata
}

//通过用户id mdid查找MD笔记内容
func (this *MDServers) GetMDContent(uid string,mid string,types int) (code int, count int, data interface{},title string) {
	//1. 判断uid用户是否拥有这个mid,如果拥有则返回md内容，不拥有则提示你没有权限访问此MD内容
	ismd,title,err := new(models.MDInof).IsMD(uid,mid)
	if err != nil && err.Error() != "record not found"{
		fmt.Println("判断查看MD权限错误，错误信息:",err)
		return 0,1,"判断查看MD权限错误，错误信息",""
	}
	if ismd {
		mdtxt,err := new(models.MDText).GetMDTxt(mid)
		if err != nil {
			fmt.Println("查询MD内容错误，错误信息:",err)
			return 0,1,"查询MD内容错误，错误信息",""
		}
		
		if types == 1 {
			//2. 增加查看次数
			go new(models.MDInof).AddMDViewTimes(mid)
		}else if types == 2 {
			//2. 增加修改次数
			go new(models.MDInof).AddMDModifytimes(mid)
		}
		
		return 1,1,mdtxt,title
	}
	fmt.Println("你没有权限查看此笔记")
	return 2,1,"你没有权限查看此笔记",""
}


//通过笔记名模糊搜索查询笔记
func (this *MDServers) SearchNoteinfo(word,uid string) (code int, count int, data interface{}) {
	datas,err := new(models.MDInof).SearchTitle(word,uid)
	if err != nil{
		fmt.Println("获取所有笔记错误，后台错误",err)
		return 0,1,"获取所有笔记错误"
	}
	rdata := this.NoteListFormat(datas)
	
	return 0,1,rdata
}

//删除笔记到回收站
func (this *MDServers) DeleteNote(mdid,uid string) (code int, count int, data interface{}) {
	ismd,_,err := new(models.MDInof).IsMD(uid,mdid)
	if err != nil && err.Error() != "record not found"{
		fmt.Println("判断查看MD权限错误，错误信息:",err)
		return 0,1,"判断查看MD权限错误，错误信息"
	}
	if ismd {
		err := new(models.MDInof).ToDEL(mdid,uid)
		if err != nil {
			fmt.Println("删除笔记错误",err)
			return 0,1,"删除笔记错误"
		}
		return 1,1,"删除笔记成功"
	}
	return 2,1,"你没有权限删除此笔记"
}

//SchenNote 永久删除笔记
func (this *MDServers) SchenNote(mdid,uid string) (code int, count int, data interface{}) {
	ismd,_,err := new(models.MDInof).IsMD(uid,mdid)
	if err != nil && err.Error() != "record not found"{
		fmt.Println("判断查看MD权限错误，错误信息:",err)
		return 0,1,"判断查看MD权限错误，错误信息"
	}
	if ismd {
		err := new(models.MDInof).Schen(mdid,uid)
		if err != nil {
			fmt.Println("删除笔记错误",err)
			return 0,1,"删除笔记错误"
		}
		return 1,1,"删除笔记成功"
	}
	return 2,1,"你没有权限删除此笔记"
}

//恢复笔记到笔记本
func (this *MDServers) RestoreToNotes(mdid,uid string,notes int) (code int, count int, data interface{}) {
	err := new(models.MDInof).NoteToNotes(mdid,uid,notes)
	if err != nil {
		fmt.Println("恢复笔记错误",err)
		return 0,1,"恢复笔记错误"
	}
	return 1,1,"恢复笔记成功"
}