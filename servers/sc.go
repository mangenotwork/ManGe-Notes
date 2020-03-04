package servers
/*
	主要为素材模块的业务
*/
import (
	"fmt"
	"time"
	_ "encoding/json"
	"strings"

	object "man/ManNotes/object"
	//util "man/ManNotes/util"
	models "man/ManNotes/models"
	//rdb "man/ManNotes/models/redis"

	"github.com/astaxie/beego"
)

type SUCai struct {}

//将上传的图片信息存入数据库
func (this *SUCai) UploadImg(uid string,imgpath string, imgsize int64, imgdatas *object.UpLoadImgData) (string,string) {
	//1. 图片上传成功后将图片信息存入图片表
	mainurl := beego.AppConfig.DefaultString("img::mainurl", "")	
	returnurl := fmt.Sprintf("%s%s",mainurl,imgpath)
	nowtime := time.Now().Unix()

	var returnname string
	if imgdatas.ImgName == "" {
		returnname = imgpath
	}else{
		returnname = imgdatas.ImgName
	}

	tags := strings.Replace(imgdatas.ImgTag, "\n", "", -1)
	tags = strings.Trim(tags, " ")
	fmt.Println(tags)
	imginfo := &models.IMGInfo{
		ImgName : returnurl,
		Imgdec : returnname,
		Uid : uid,
		Time : nowtime,
		Date : time.Unix(nowtime, 0).Format("2006-01-02 15:04:05"),
		Size : imgsize,
		Imgtag : tags,
	}
	imginfo.CreateImg()

	//2.返回图片访问链接

	
	fmt.Println(returnurl)
	return returnurl,returnname

}