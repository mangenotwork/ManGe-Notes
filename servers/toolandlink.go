package servers

/*
		工具与链接模块
*/

import (
	"fmt"
	_"time"
	_ "encoding/json"

	object "man/ManNotes/object"
	util "man/ManNotes/util"
	models "man/ManNotes/models"
	//rdb "man/ManNotes/models/redis"
)

type TandLServers struct {}

func (this *TandLServers) AddLink(datas *object.AddLink, uid string) (code int, count int, data string) {
	//判断上传链接类型
	//如果是  0,1 保存为  工具&链接  0 链接  1 工具
	//如果是  2,3,4 保存为  网络素材  2 网络图片链接   3网络视频链接   4网络音乐链接
	if datas.LinkType == "0" || datas.LinkType == "1" {
		return this.AddTandL(datas,uid)
	}else{
		fmt.Println("占时不支持添加网络素材链接")
		return 2,1,"占时不支持添加网络素材链接"
	}
}

//保存工具&链接
func (this *TandLServers) AddTandL(datas *object.AddLink, uid string) (code int, count int, data string) {
	//判断是否命名，如果命名为空则取网络链接的title
	var linkname string
	if datas.LinkName == ""{
		linkname,_ = new(util.WebInfo).GetWebTitle(datas.Link)
	}else{
		linkname = datas.LinkName
	}
	//取网络//主域名+ /favicon.ico
	linkico := new(util.WebInfo).GetIcoLink(datas.Link)
	linktype,_ := new(util.Str).NumberToInt(datas.LinkType)
	//添加链接信息
	linkinfos := &models.ToolandLink{
		UID : uid,
		Name : linkname,
		Des : datas.LinkDes,
		Link : datas.Link,
		Ico : linkico,
		LinkType : linktype,
	}
	err := linkinfos.Create()
	if err != nil{
		fmt.Println("添加收藏错误",err)
		return 0,1,"添加收藏链接错误"
	}
	return 1,1,"添加成功"

}	

//获取网络工具列表
func (this *TandLServers) GetToolList(uid string) (code int, count int, data interface{}) {
	toollist,err := new(models.ToolandLink).GeTools(0,20,uid)
	if err != nil {
		fmt.Println(err)
		return 1,1,"获取工具列表错误，后端错误"
	}
	return 1,1,toollist
}

//获取链接列表
func (this *TandLServers) GetLinks(uid string) (code int, count int, data interface{}) {
	toollist,err := new(models.ToolandLink).GetLinks(0,20,uid)
	if err != nil {
		fmt.Println(err)
		return 1,1,"获取工具列表错误，后端错误"
	}
	return 1,1,toollist
}

//修改收藏的链接EDLink
func (this *TandLServers) EditLink(datas *object.EDLinks, uid string) (code int, count int, data string) {
	//判断是否命名，如果命名为空则取网络链接的title
	var linkname string
	if datas.LinkName == ""{
		linkname,_ = new(util.WebInfo).GetWebTitle(datas.Link)
	}else{
		linkname = datas.LinkName
	}
	//取网络//主域名+ /favicon.ico
	linkico := new(util.WebInfo).GetIcoLink(datas.Link)
	linkinfos := &models.ToolandLink{
		Id : datas.LinkID,
		UID : uid,
		Name : linkname,
		Des : datas.LinkDes,
		Link : datas.Link,
		Ico : linkico,
	}
	err := linkinfos.UpdateLink()
	if err != nil{
		fmt.Println("修改收藏的链接错误",err)
		return 0,1,"修改收藏的链接错误"
	}
	return 1,1,"修改成功"

}

//删除收藏的链接
func (this *TandLServers) DELLink(uid string,linkid int) (code int, count int, data string) {
	err := new(models.ToolandLink).DELLink(uid,linkid)
	if err != nil{
		fmt.Println("删除收藏的链接错误",err)
		return 0,1,"删除收藏的链接错误"
	}
	return 1,1,"删除成功"
}

//GetAllLinksInfo mange管理模块  获取收藏链接列表
func (this *TandLServers) GetAllLinksInfo(uid string) (count int,data interface{}) {
	alldata,err := new(models.ToolandLink).GetAll(0,100,uid)
	if err != nil{
		fmt.Println("删除收藏的链接错误",err)
		return 0,"删除收藏的链接错误"
	}

	linkslist := make([]*object.LinksInfo,0)

	for _,k := range alldata{
		var typeStr string 
		if k.LinkType == 0{
			typeStr = "网络链接"
		}else if k.LinkType == 1{
			typeStr = "网络工具地址"
		}

		linkslist = append(linkslist,&object.LinksInfo{
			Id : k.Id,
			Name : k.Name,
			Des : k.Des,
			Link : k.Link,
			Ico : k.Ico,
			Tag : k.Tag,
			LinkTypeStr : typeStr,
			})
	}

	return len(alldata),linkslist
}

//MangeEditLink 管理模块修改收藏链接
func (this *TandLServers) MangeEditLink(datas *object.EDLinksInfo, uid string) (code int, count int, data string) {

	idnumber,_ := new(util.Str).NumberToInt(datas.Id)
	typenumber,_ := new(util.Str).NumberToInt(datas.LinkType)
	linkinfos := &models.ToolandLink{
		Id : idnumber,
		UID : uid,
		Name : datas.Name,
		Des : datas.Des,
		Link : datas.Link,
		Ico : datas.Ico,
		Tag : datas.Tag,
		LinkType : typenumber,
	}
	err := linkinfos.EDLink()
	if err != nil{
		fmt.Println("修改收藏的链接错误",err)
		return 0,1,"修改收藏的链接错误"
	}
	return 1,1,"修改成功"

}