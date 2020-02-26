package controllers

/*
		上传模块
*/
import (
	"fmt"
	"strings"
	"path"
	"path/filepath"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	util "man/ManNotes/util"
	models "man/ManNotes/models"

	"github.com/nfnt/resize"
	"github.com/astaxie/beego"
	"github.com/rs/xid"
)


type UploadController struct {
	Controller
}


//写笔记上传图片
// v1 版本只支持存本机服务器磁盘，网络访问需要服务器nginx配置图片请求
func (this *UploadController) UploadImg(){
	uid := this.GetUid()

	f, h, err := this.GetFile("file")
	if err != nil{
		fmt.Println("获取到上传的文件错误",err)
		this.RetuenJson(0,1,"获取到上传的文件错误")
		return 
	}

	fmt.Println(f)
	fmt.Println(h.Filename)
	fmt.Println(h.Size)
	fileExt := filepath.Ext(h.Filename)
	fmt.Println(fileExt)
	//1.验证图片类型
	// v1 版本只支持 ".jpg", ".png", ".jpeg"   v2 版本开始支持 ".gif"
	var ExtList = []string{".jpg", ".png", ".jpeg"}
	fileExt = strings.ToLower(fileExt)
	if !new(util.Str).IsItemStr(ExtList,fileExt){
		fmt.Println("上传的图片类型不符合标准，请上传jpg,png类型的图片")
		this.RetuenJson(0,1,"上传的图片类型不符合标准，请上传jpg,png类型的图片")
		return
	}

	//图片命名
	id := xid.New()
	newimgName := fmt.Sprintf("%s%s",id.String(),fileExt)
	savedir := beego.AppConfig.DefaultString("img::savepath", "")	
	savepath := fmt.Sprintf("%s%s",savedir,newimgName)
	fmt.Println(savepath)
	fmt.Println(path.Join(savepath))

	//2.验证图片大小 如果图片大于指定大小对图片进行压缩
	if h.Size > 1204*50{
		fmt.Println("上传的图片大于50kb 将进行压缩处理")
		var img image.Image
		var imgerr error
		// decode jpeg into image.Image
		if fileExt == ".png"{
			img, imgerr = png.Decode(f)
		}else{
			img, imgerr = jpeg.Decode(f)
		}
		
		if imgerr != nil {
			fmt.Println(imgerr)
			this.RetuenJson(0,1,"压缩图片错误")
			return
		}
		f.Close()
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		//resize.Lanczos3
		m := resize.Resize(480, 0, img, resize.NearestNeighbor)
		out, err := os.Create(savepath)
		if err != nil {
			fmt.Println(err)
		}
		
		// write new image to file
		jpeg.Encode(out, m, nil)
		out.Close()


	}else{
		f.Close()   
		//保存文件到指定的位置
		//static/uploadfile,这个是文件的地址，第一个static前面不要有/
		this.SaveToFile("file", path.Join(savepath)) 
	}

	imgsize := fileSize(savepath)

	//3.保存图片，并将图片名存入数据库
	nowtime := time.Now().Unix()
	imginfo := &models.IMGInfo{
		ImgName : newimgName,
		Imgdec : "",
		Uid : uid,
		Time : nowtime,
		Date : time.Unix(nowtime, 0).Format("2006-01-02 15:04:05"),
		Size : imgsize,
		Imgtag : "",
	}
	imginfo.CreateImg()
	//5.返回图片访问链接

	mainurl := beego.AppConfig.DefaultString("img::mainurl", "")	
	returnurl := fmt.Sprintf("%s%s",mainurl,newimgName)
	fmt.Println(returnurl)

	this.RetuenJson(1,1,&ReturnImgInfo{newimgName,returnurl})
}

type ReturnImgInfo struct{
	Name string `json:"name"`
	Url string `json:"url"`
}


func fileSize(path string) int64 {
	f, e := os.Stat(path)
	if e != nil {
		fmt.Println(e.Error())
		return 0
	}
	fsize := f.Size() 
	return fsize
}