package routers

import (
	"log"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/mangenotwork/ManGe-Notes/controllers"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

//Check Install
var Check_Install = func(ctx *context.Context) {
	log.Println("中间件")
	log.Println(ctx.Request.RequestURI)

	flag := false

	if util.FileExist(object.InstallJsonPath) {
		infodata := object.OpenInstallFile()
		if infodata.Step == -1 {
			flag = true
		}
	}

	if !flag && strings.Index(ctx.Request.RequestURI, "install") == -1 {
		ctx.Redirect(302, "/install")
	} else if flag && strings.Index(ctx.Request.RequestURI, "install") == 1 {
		ctx.Redirect(302, "/")
	}

}

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter, Check_Install)

	//pg 页面请求   PGController
	beego.Router("/", &controllers.PGController{}, "get:IndexPG") //登录页面
	beego.Router("/login", &controllers.PGController{}, "get:Login")
	beego.Router("/islogin", &controllers.PGController{}, "get:GetIsLogin")
	beego.Router("/install", &controllers.PGController{}, "get:Install")         //安装页面
	beego.Router("/index", &controllers.PGController{}, "get:IndexPG")           //主页
	beego.Router("/mdeditor", &controllers.PGController{}, "get:MdEditorPG")     //MD编辑器
	beego.Router("/home", &controllers.PGController{}, "get:HomePG")             //首页
	beego.Router("/tool", &controllers.PGController{}, "get:ToolPG")             //工具页面
	beego.Router("/mangenotes", &controllers.PGController{}, "get:MangeNotes")   //笔记本管理
	beego.Router("/mangelinks", &controllers.PGController{}, "get:MangeLinks")   //管理收藏的链接的管理页面
	beego.Router("/chartnotes", &controllers.PGController{}, "get:ChartNotes")   //图表模块 笔记本数量分布图
	beego.Router("/mychart", &controllers.PGController{}, "get:MyChart")         //图表模块 我的综合统计
	beego.Router("/myusedspace", &controllers.PGController{}, "get:MyUsedSpace") //图表模块 我的使用空间

	//api
	api := beego.NewNamespace("/api",
		//install
		beego.NSRouter("/install/start", &controllers.InstallController{}, "get:InstallStart"),               //开始安装
		beego.NSRouter("/install/mysqltset", &controllers.InstallController{}, "get:InstallMysqlTset"),       //测试mysql连接
		beego.NSRouter("/install/mysql", &controllers.InstallController{}, "get:InstallMysql"),               // mysql
		beego.NSRouter("/install/pgsqltset", &controllers.InstallController{}, "get:InstallPgsqlTset"),       //测试pgsql连接
		beego.NSRouter("/install/pgsql", &controllers.InstallController{}, "get:InstallPgsql"),               // pgsql
		beego.NSRouter("/install/localsql", &controllers.InstallController{}, "get:InstallSqlite"),           // sqlite
		beego.NSRouter("/install/localmedia", &controllers.InstallController{}, "get:InstallLocalMedia"),     //本地磁盘存储多媒体资源
		beego.NSRouter("/install/alimedia", &controllers.InstallController{}, "get:InstallAliyunMedia"),      //阿里云存储多媒体资源
		beego.NSRouter("/install/tencentmedia", &controllers.InstallController{}, "get:InstallTencentMedia"), //腾讯云存储多媒体资源
		beego.NSRouter("/install/admin", &controllers.InstallController{}, "get:InstallAdmin"),               //创建账号
	)

	beego.AddNamespace(api)

	//登录注册    LoginController
	beego.Router("/userreg", &controllers.LoginController{}, "post:UserRegistered") //用户注册
	beego.Router("/userlogin", &controllers.LoginController{}, "post:UserLogin")    //用户登录
	beego.Router("/outlogin", &controllers.LoginController{}, "get:OutLogin")       //退出登录

	//笔记本    NotesController
	beego.Router("/cnotes", &controllers.NotesController{}, "post:CreateNotes")          //创建笔记本
	beego.Router("/noteslist", &controllers.NotesController{}, "get:GetNotesList")       //获取当前笔记本列表
	beego.Router("/getallnotes", &controllers.NotesController{}, "get:GetAllNotes")      //mange 管理模块  获取所有笔记本数据
	beego.Router("/updatenotes", &controllers.NotesController{}, "post:UpdateNotesInfo") //mange 管理模块  修改笔记本信息
	beego.Router("/delnotes/:notesid:*", &controllers.NotesController{}, "get:DelNotes") //mange 管理模块  删除笔记本

	//MD笔记    MDController
	beego.Router("/cmd", &controllers.MDController{}, "post:CreateMD")                    //创建MD笔记
	beego.Router("/allnote", &controllers.MDController{}, "get:GetAllNote")               //获取所有笔记
	beego.Router("/notes/:id:int", &controllers.MDController{}, "get:NotesMDList")        //获取笔记本对应的笔记
	beego.Router("/noteshow/:mdid:*", &controllers.MDController{}, "get:MDShow")          //显示MD笔记内容
	beego.Router("/editmdnote/:mdid:*", &controllers.MDController{}, "get:MDEditPG")      //修改MD笔记内容编辑页面
	beego.Router("/modifynote/:mdid:*", &controllers.MDController{}, "post:MDNoteModify") //提交修改笔记内容
	beego.Router("/searchnote", &controllers.MDController{}, "get:SearchNote")            //搜索笔记
	beego.Router("/delnote/:mdid:*", &controllers.MDController{}, "get:DelNote")          //删除笔记到回收站
	beego.Router("/recycler", &controllers.MDController{}, "get:NoteRecycler")            //回收站
	beego.Router("/rnoteshow/:mdid:*", &controllers.MDController{}, "get:RMDShow")        //显示回收站MD笔记内容
	beego.Router("/schen/:mdid:*", &controllers.MDController{}, "get:SchenNote")          //恢复到笔记本
	beego.Router("/restore/:mdid:*", &controllers.MDController{}, "get:RestoreNote")      //笔记永久删除
	beego.Router("/todraft", &controllers.MDController{}, "post:ToDraft")                 //笔记保存到草稿
	beego.Router("/draftlist", &controllers.MDController{}, "get:DraftList")              //草稿笔记列表
	beego.Router("/dranoteshow/:mdid:*", &controllers.MDController{}, "get:DraNoteShow")  //显示草稿笔记

	//收藏（链与工具）   TandLController
	beego.Router("/collectlink", &controllers.TandLController{}, "post:AddCollectLink") //添加网络资源
	beego.Router("/tandllist", &controllers.TandLController{}, "get:GetTandL")          //获取网络工具列表
	beego.Router("/linkshow", &controllers.TandLController{}, "get:LinkShow")           //收藏链接的显示页
	beego.Router("/edlink", &controllers.TandLController{}, "post:EDLink")              //编辑收藏的链接
	beego.Router("/dellink", &controllers.TandLController{}, "get:DELLink")             //删除收藏的链接
	beego.Router("/getlinks", &controllers.TandLController{}, "get:GetLinks")           //mange 管理模块 管理收藏链接，获取收藏链接信息
	beego.Router("/mageedlink", &controllers.TandLController{}, "post:MageEDLink")      //mange 管理模块 编辑收藏的链接

	//上传功能
	beego.Router("/imgupload", &controllers.UploadController{}, "post:UploadImg") //写笔记上传图片

	//漫鸽笔记社区
	beego.Router("/shequ", &controllers.PGController{}, "get:Shequ") //主页

	//素材模块
	beego.Router("/sucai", &controllers.PGController{}, "get:SuCai")               //主页
	beego.Router("/addsc", &controllers.PGController{}, "get:AddSuCai")            //添加素材
	beego.Router("/addlinkimg", &controllers.SucaiController{}, "post:AddLinkImg") //添加网络图片素材
	beego.Router("/myimg", &controllers.SucaiController{}, "get:MyImg")            //获取我的图片
	beego.Router("/tomangeimg", &controllers.SucaiController{}, "get:ToMangeImg")  //分享到图库
	beego.Router("/mangeimg", &controllers.SucaiController{}, "get:MangeImg")      //漫鸽图库
}
