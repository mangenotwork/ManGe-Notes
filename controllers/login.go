package controllers

import (
	"fmt"

	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/servers"
)

type LoginController struct {
	Controller
}

//用户注册  账号密码注册
func (this *LoginController) UserRegistered() {
	ip := this.GetIP()

	var obj object.UserRegInfo

	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}

	code, count, data, token := new(servers.LoginServers).UserReg(&obj, ip)
	this.SetToken(token)
	this.RetuenJson(code, count, data)
}

//用户登录  账号密码登录
func (this *LoginController) UserLogin() {
	ip := this.GetIP()
	var obj object.Logininfo
	err := this.ResolvePostData(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip, obj)

	code, count, data, token := new(servers.LoginServers).UserAccLogin(&obj, ip)
	this.SetToken(token)
	this.RetuenJson(code, count, data)
}

//退出登录
func (this *LoginController) OutLogin() {
	this.ClearToken()
	this.Redirect("/", 302)
}
