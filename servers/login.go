package servers
/*
	主要为登录注册服务
*/

import (
	"fmt"
	"time"
	"encoding/json"

	object "man/ManNotes/object"
	util "man/ManNotes/util"
	models "man/ManNotes/models"
	rdb "man/ManNotes/models/redis"
)

//默认头像
const DefaultAvatar = "https://avatars3.githubusercontent.com/u/53510864?s=88&v=4"

type LoginServers struct {}

//用户注册 账号密码注册
func (this *LoginServers) UserReg(datas *object.UserRegInfo,ip string) (code int, count int, data string, jwtStr string){
	//验证邀请码
	if datas.InviteCode != "mangenotes2020"{
		return 0,1,"邀请码错误！需要邀请码请加qq群:1060290526",""
	}

	//检查账号是否存在
	if !new(models.ACC).ACCIsAccount(datas.Acc){
		return 0,1,"账号已存在",""
	}

	//创建账号id
	id,uidErr := util.Int64ID()
	if uidErr != nil{
		fmt.Println("生成 int64 ID错误")
		return 0,1,"生成 int64 ID错误",""
	}
	uid := fmt.Sprintf("uid_%d",id)
	//密码转MD5
	nowtime := time.Now().Unix()
	password := util.Md5Crypt(datas.Password,nowtime)
	fmt.Println(uid,datas.Acc,password)
	
	newuser := &models.ACC{
		UserId : uid,
		Account : datas.Acc,
		Password : password,
		Phone : datas.Phone,
		Mail : "",
		Avatar : DefaultAvatar,
		Createtime : nowtime,
		Logintime : 0,
		LoginIP : ip,
	}
	err := newuser.CreateUser()
	if err != nil {
		fmt.Println("注册失败",err)
		return 0,1,fmt.Sprintf("注册失败:%s",err),""
	}
	jwtStr,jwtStrErr := util.CreateJwtToken(uid)
	if jwtStrErr != nil {
		fmt.Println("生成Token失败",err)
		return 0,1,fmt.Sprintf("生成Token失败:%s",jwtStrErr),""
	}

	//将token 保存到Redis
	go new(rdb.RDB).StringSet(fmt.Sprintf("login:%s",uid), jwtStr)

	return 1,1,"注册成功",jwtStr
	

	//如果邀请码验证成功，给邀请码赋予账号id
}


//用户登录 账号密码登录
func (this *LoginServers) UserAccLogin(datas *object.Logininfo,ip string) (code int, count int, data string, jwtStr string){
	//用过账号查询用户信息 
	userinfo := &models.ACC{}
	if err := userinfo.GetACCinfo(datas.LoginAcc); err != nil {
		fmt.Println("后台错误，获取账号信息错误：",err)
		return 0,1,fmt.Sprintf("后台错误，获取账号信息错误:%s",err),""
	}else if err != nil && err.Error() == "record not found" {
		fmt.Println("账号不存在")
		return 2,1,"账号不存在",""
	}

	//匹配密码
	inputPwd := util.Md5Crypt(datas.LoginPassword,userinfo.Createtime)
	fmt.Println(datas.LoginPassword,userinfo.Createtime,inputPwd,userinfo.Password)
	if inputPwd == userinfo.Password{
		//将登陆的用户信息写入redis  uinfo:uid hash
		var lasttime string = "首次登陆"
		if userinfo.Logintime != 0{ 
			lasttime = time.Unix(userinfo.Logintime, 0).Format("2006-01-02 15:04:05")
		}
		ukeys := fmt.Sprintf("uinfo:%s",userinfo.UserId) 
		ubasisInfo := &object.UserBasisInfo{
			UName : userinfo.Account,
			UAvatar : userinfo.Avatar,
			LastLogin : lasttime,
			LastLoginIp : userinfo.LoginIP,
		}
		uvalue,_ := json.Marshal(ubasisInfo)
		ufield := "basis"
		new(rdb.RDB).HashSet(ukeys,ufield,uvalue)

		//更新用户本次登陆的时间和ip
		nowtime := time.Now().Unix()
		go userinfo.UpdateLastLogin(userinfo.UserId,nowtime,ip)

		//设置token
		jwtStr,jwtStrErr := util.CreateJwtToken(userinfo.UserId)
		if jwtStrErr != nil {
			fmt.Println("生成Token失败",jwtStrErr)
			return 0,1,fmt.Sprintf("生成Token失败:%s",jwtStrErr),""
		}

		//将token 保存到Redis
		go new(rdb.RDB).StringSet(fmt.Sprintf("login:%s",userinfo.UserId), jwtStr)
		return 1,1,"登陆成功",jwtStr
	}

	fmt.Println("密码错误")
	return 2,1,"密码错误",""
}
