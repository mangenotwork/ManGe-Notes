package servers

/*
	主要为素材模块的业务
*/
import (
	_ "encoding/json"
	"fmt"
	"strings"
	"time"

	object "github.com/mangenotwork/ManGe-Notes/object"
	//util "man/ManNotes/util"
	models "github.com/mangenotwork/ManGe-Notes/models"
	//rdb "man/ManNotes/models/redis"

	"github.com/astaxie/beego"
	"github.com/rs/xid"
)

type SUCai struct{}

//将上传的图片信息存入数据库
func (this *SUCai) UploadImg(uid string, imgpath string, imgsize int64, imgdatas *object.UpLoadImgData) (string, string) {
	//1. 图片上传成功后将图片信息存入图片表
	mainurl := beego.AppConfig.DefaultString("img::mainurl", "")
	returnurl := fmt.Sprintf("%s%s", mainurl, imgpath)
	nowtime := time.Now().Unix()

	var returnname string
	if imgdatas.ImgName == "" {
		returnname = imgpath
	} else {
		returnname = imgdatas.ImgName
	}

	tags := strings.Replace(imgdatas.ImgTag, "\n", "", -1)
	tags = strings.Trim(tags, " ")
	fmt.Println(tags)
	imginfo := &models.IMGInfo{
		ImgName: returnurl,
		Imgdec:  returnname,
		Uid:     uid,
		Time:    nowtime,
		Date:    time.Unix(nowtime, 0).Format("2006-01-02 15:04:05"),
		Size:    imgsize,
		Imgtag:  tags,
	}
	imginfo.CreateImg()

	//2.返回图片访问链接

	fmt.Println(returnurl)
	return returnurl, returnname

}

//添加网络图片链接
func (this *SUCai) ADDLinkImg(uid string, datas *object.LinkImgData) (code int, count int, data string) {
	nowtime := time.Now().Unix()
	var returnname string
	if datas.ImgName == "" {
		id := xid.New()
		returnname = fmt.Sprintf("%s", id.String())
	} else {
		returnname = datas.ImgName
	}

	tags := strings.Replace(datas.ImgTag, "\n", "", -1)
	tags = strings.Trim(tags, " ")

	imginfo := &models.IMGInfo{
		ImgName: datas.ImgLink,
		Imgdec:  returnname,
		Uid:     uid,
		Time:    nowtime,
		Date:    time.Unix(nowtime, 0).Format("2006-01-02 15:04:05"),
		Size:    0,
		Imgtag:  tags,
	}
	err := imginfo.CreateImg()
	if err != nil {
		fmt.Println("添加网络图片链接错误")
		return 0, 1, "添加网络图片链接错误"
	}
	return 1, 1, "添加网络图片链接成功"

}

//获取我的图片 GetMyImg
func (this *SUCai) GetMyImg(uid string) (code int, count int, data interface{}) {
	rdata, err := new(models.IMGInfo).GetMyImg(uid)
	if err != nil {
		fmt.Println("获取我的图片错误 : ", err)
		return 0, 1, "获取我的图片错误"
	}

	imgList := make([]*object.ImgInfoShow, 0)

	for _, v := range rdata {
		imgList = append(imgList, &object.ImgInfoShow{
			ImgId:     v.Id,
			ImgUrl:    v.ImgName,
			ImgName:   v.Imgdec,
			ImgTag:    v.Imgtag,
			ImgCreate: v.Date,
		})
	}

	return 1, 1, imgList
}

//分享到 图片库 ToMangeImg
func (this *SUCai) ToMangeImg(uid string, imgid int) (code int, count int, data interface{}) {
	//图片权限验证
	ismyImg, imginfo := new(models.IMGInfo).IsMyImg(imgid, uid)
	fmt.Println(ismyImg, imginfo)
	if !ismyImg {

		nowtime := time.Now().Unix()
		addmangeimg := &models.SCIMGInfo{
			ImgName: imginfo.ImgName,
			Imgdec:  imginfo.Imgdec,
			Uid:     imginfo.Uid,
			Time:    nowtime,
			Date:    time.Unix(nowtime, 0).Format("2006-01-02 15:04:05"),
			Size:    imginfo.Size,
			Imgtag:  imginfo.Imgtag,
		}
		addmangeimg.AddSCImg()
		return 1, 1, "感谢你的分享"
	}
	return 0, 1, "你没有权限操作此图片"

}

// GetMangeImg 获取漫鸽图库
func (this *SUCai) GetMangeImg(uid string) (code int, count int, data interface{}) {
	rdata, err := new(models.SCIMGInfo).MangeImgList()
	if err != nil {
		fmt.Println("获取我的图片错误 : ", err)
		return 0, 1, "获取我的图片错误"
	}

	imgList := make([]*object.ImgInfoShow, 0)

	for _, v := range rdata {
		imgList = append(imgList, &object.ImgInfoShow{
			ImgId:     v.Id,
			ImgUrl:    v.ImgName,
			ImgName:   v.Imgdec,
			ImgTag:    v.Imgtag,
			ImgCreate: v.Date,
		})
	}

	return 1, 1, imgList
}
