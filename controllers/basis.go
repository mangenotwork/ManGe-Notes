package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/conn"

	"github.com/astaxie/beego"
	"github.com/mangenotwork/ManGe-Notes/util"
)

type Controller struct {
	beego.Controller
}

/*
约定
code:
		1   成功
		0   错误(后台)
		2	失败

*/
type RetuenJsonData struct {
	Code  int         `json:"code"`
	Count int         `json:"count"`
	Datas interface{} `json:"data"`
}

//接口返回的结构 RetuenJsonData
func (this *Controller) RetuenJson(code int, count int, data interface{}) {
	returndata := &RetuenJsonData{code, count, data}
	this.Data["json"] = returndata
	this.ServeJSON()
}

type MangeJsonData struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

//接口返回的结构 MangeJsonData
func (this *Controller) MangeJson(code int, msg string, count int, data interface{}) {
	returndata := &MangeJsonData{code, msg, count, data}
	this.Data["json"] = returndata
	this.ServeJSON()
}

//判断是否会话  一般会话不用刷新token 便于请求的流畅性
func (this *Controller) IsSession() bool {
	token := this.Ctx.GetCookie("token")
	uid, err := util.ParseJwtToken(token)
	if err != nil {
		beego.Error("[Token]解析Token错误:", err.Error(), err)
		this.Redirect("/login", 302)
		return false
	}
	//判断缓存里的token
	uidKey := fmt.Sprintf("user:%s", uid.Data)
	isToken := conn.CachesGetStr(uidKey)
	fmt.Println(isToken)
	if token == isToken {
		return true
	}
	this.Redirect("/login", 302)
	return false
}

func (this *Controller) IsLogin() bool {
	token := this.Ctx.GetCookie("token")
	uid, err := util.ParseJwtToken(token)
	if err != nil {
		beego.Error("[Token]解析Token错误:", err.Error(), err)
		return false
	}
	//判断缓存里的token
	uidKey := fmt.Sprintf("user:%s", uid.Data)
	isToken := conn.CachesGetStr(uidKey)
	fmt.Println(isToken)
	if token == isToken {
		return true
	}
	return false
}

//解析post接收的参数
func (this *Controller) ResolvePostData(obj interface{}) error {
	fmt.Println(this.Ctx.Input.RequestBody)
	jsonerr := json.Unmarshal(this.Ctx.Input.RequestBody, obj)
	if jsonerr != nil {
		fmt.Println(" 解析Json错误 ： ", jsonerr)
	}
	return jsonerr
}

//获取客户ip
/*
如果是Nginx 需要设置如下
location / {
        proxy_set_header            X-real-ip $remote_addr;
        proxy_pass http://upstream/;
    }
再使用this.Ctx.Request.Header.Get("X-Real-ip")获取ip
*/
func (this *Controller) GetIP() string {
	ip := this.Ctx.Input.IP()
	return ip
}

//设置 Token
func (this *Controller) SetToken(token string) {
	this.Ctx.SetCookie("token", token, 3600*24*7, "/")
}

//解析Token 获取到Uid
func (this *Controller) GetUid() (uid string) {
	if this.IsLogin() {
		token := this.Ctx.GetCookie("token")
		jwtobj, _ := util.ParseJwtToken(token)
		uid = jwtobj.Data
	}
	return
}

//ClearToken 清空Token
func (this *Controller) ClearToken() {
	fmt.Println("退出登录")
	token := this.Ctx.GetCookie("token")
	jwtobj, _ := util.ParseJwtToken(token)
	uidKey := fmt.Sprintf("user:%s", jwtobj.Data)
	conn.CachesSet(uidKey, "")
	this.Ctx.SetCookie("token", "")
	return
}
