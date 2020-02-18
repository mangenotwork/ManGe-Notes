package controllers

import (
	"fmt"
	"encoding/json"

	util "man/ManNotes/util"
	rdb "man/ManNotes/models/redis"

	"github.com/astaxie/beego"
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

//判断是否登录
func (this *Controller) IsLogin() {
	token := this.Ctx.GetCookie("token")
	if token != "" &&  util.VerifyJwtToken(token){
		fmt.Println(token)
		uid,err := util.ParseJwtToken(token)
		if err!=nil{
			beego.Error("[Token]解析Token错误:", err.Error(), err)
		}
		fmt.Println(uid.Data)
		
		//判断redis里的token
		uidKey := fmt.Sprintf("login:%s",uid.Data)
		isToken,_ := new(rdb.RDB).StringJudge(uidKey, token)
		fmt.Println(isToken)
		if isToken{
			fmt.Println("Token 匹配成功")
			//新的token
			newToken,newTokenErr := util.CreateJwtToken(uid.Data)
			if newTokenErr!=nil{
				beego.Error("[Token]生成Token错误:", newTokenErr.Error(), newTokenErr)
			}

			//Redis保存新的token
			new(rdb.RDB).StringSet(uidKey, newToken)
			//客户端设置新的token
			this.SetToken(newToken)

			this.Redirect("/index",302)
		}
		
	}
}

//判断是否会话
func (this *Controller) IsSession() bool {
	token := this.Ctx.GetCookie("token")
	uid,err := util.ParseJwtToken(token)
	if err!=nil{
		beego.Error("[Token]解析Token错误:", err.Error(), err)
		this.Redirect("/",302)
		return false
	}
	//判断redis里的token
	uidKey := fmt.Sprintf("login:%s",uid.Data)
	isToken,_ := new(rdb.RDB).StringJudge(uidKey, token)
	fmt.Println(isToken)
	if isToken{
		//新的token
		newToken,newTokenErr := util.CreateJwtToken(uid.Data)
		if newTokenErr!=nil{
			beego.Error("[Token]生成Token错误:", newTokenErr.Error(), newTokenErr)
			this.Redirect("/",302)
			return false
		}

		//Redis保存新的token
		new(rdb.RDB).StringSet(uidKey, newToken)
		//客户端设置新的token
		this.SetToken(newToken)
		return true
	}
	this.Redirect("/",302)
	return false
}

//接口返回的结构 RetuenJsonData
func (this *Controller) RetuenJson(code int, count int, data interface{}) {
	returndata := &RetuenJsonData{code, count, data}
	this.Data["json"] = returndata
	this.ServeJSON()
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
func (this *Controller) GetIP() string{
	ip := this.Ctx.Input.IP()
	return ip
}

//设置 Token
func (this *Controller) SetToken(token string){
	
	this.Ctx.SetCookie("token", token, 3600*24, "/")
}

//解析Token 获取到Uid
func (this *Controller) GetUid() (uid string){
	if this.IsSession(){
		token := this.Ctx.GetCookie("token")
		jwtobj,_ := util.ParseJwtToken(token)
		uid = jwtobj.Data
	}
	return
}

//ClearToken 清空Token
func (this *Controller) ClearToken(){
	fmt.Println("退出登录")
	token := this.Ctx.GetCookie("token")
	jwtobj,_ := util.ParseJwtToken(token)
	uid := jwtobj.Data
	new(rdb.RDB).DELKey(fmt.Sprintf("login:%s",uid))
	this.Ctx.SetCookie("token", "")
	return
}
