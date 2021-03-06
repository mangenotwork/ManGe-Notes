package servers

/*
	主要为登录注册服务
*/

import (
	//"encoding/json"
	"fmt"
	"time"

	"github.com/mangenotwork/ManGe-Notes/conn"
	"github.com/mangenotwork/ManGe-Notes/dao"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

//默认头像
const DefaultAvatar = "https://avatars3.githubusercontent.com/u/53510864?s=88&v=4"

type LoginServers struct{}

//用户注册 账号密码注册
func (this *LoginServers) UserReg(datas *object.UserRegInfo, ip string) (code int, count int, data string, jwtStr string) {
	// TODO: 验证邀请码功能重构
	// if datas.InviteCode != "mangenotes2020" {
	// 	return 0, 1, "邀请码错误！需要邀请码请加qq群:1060290526", ""
	// }

	//检查账号是否存在
	if !new(dao.DaoACC).ACCIsAccount(datas.Acc) {
		return 0, 1, "账号已存在", ""
	}

	//创建账号id
	id, uidErr := util.Int64ID()
	if uidErr != nil {
		fmt.Println("生成 int64 ID错误")
		return 0, 1, "生成 int64 ID错误", ""
	}
	uid := fmt.Sprintf("uid_%d", id)
	//密码转MD5
	nowtime := time.Now().Unix()
	password := util.Md5Crypt(datas.Password, nowtime)
	fmt.Println(uid, datas.Acc, password)

	newuser := &models.ACC{
		UserId:     uid,
		Account:    datas.Acc,
		Password:   password,
		Phone:      datas.Phone,
		Mail:       "",
		Avatar:     DefaultAvatar,
		Createtime: nowtime,
		Logintime:  0,
		LoginIP:    ip,
	}
	err := new(dao.DaoACC).CreateUser(newuser)
	if err != nil {
		fmt.Println("注册失败", err)
		return 0, 1, fmt.Sprintf("注册失败:%s", err), ""
	}
	jwtStr, jwtStrErr := util.CreateJwtToken(uid)
	if jwtStrErr != nil {
		fmt.Println("生成Token失败", err)
		return 0, 1, fmt.Sprintf("生成Token失败:%s", jwtStrErr), ""
	}

	//将用户token存入缓冲
	tokenKey := "user:" + uid
	conn.CachesSet(tokenKey, jwtStr)

	return 1, 1, "注册成功", jwtStr

	//如果邀请码验证成功，给邀请码赋予账号id
}

//用户登录 账号密码登录
func (this *LoginServers) UserAccLogin(datas *object.Logininfo, ip string) (code int, count int, data string, jwtStr string) {
	//用过账号查询用户信息
	userinfo, err := new(dao.DaoACC).GetACCinfo(datas.LoginAcc)
	if err != nil {
		fmt.Println("后台错误，获取账号信息错误：", err)
		return 0, 1, fmt.Sprintf("后台错误，获取账号信息错误:%s", err), ""
	} else if err != nil && err.Error() == "record not found" {
		fmt.Println("账号不存在")
		return 2, 1, "账号不存在", ""
	}

	//匹配密码
	inputPwd := util.Md5Crypt(datas.LoginPassword, userinfo.Createtime)
	fmt.Println(datas.LoginPassword, userinfo.Createtime, inputPwd, userinfo.Password)
	if inputPwd == userinfo.Password {

		//更新用户本次登陆的时间和ip
		nowtime := time.Now().Unix()
		go new(dao.DaoACC).UpdateLastLogin(userinfo.UserId, nowtime, ip)

		//设置token
		jwtStr, jwtStrErr := util.CreateJwtToken(userinfo.UserId)
		if jwtStrErr != nil {
			fmt.Println("生成Token失败", jwtStrErr)
			return 0, 1, fmt.Sprintf("生成Token失败:%s", jwtStrErr), ""
		}

		//将用户token存入缓冲
		tokenKey := "user:" + userinfo.UserId
		conn.CachesSet(tokenKey, jwtStr)

		return 1, 1, "登陆成功", jwtStr
	}

	fmt.Println("密码错误")
	return 2, 1, "密码错误", ""
}
