package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mangenotwork/ManGe-Notes/conn"
	object "github.com/mangenotwork/ManGe-Notes/object"
	servers "github.com/mangenotwork/ManGe-Notes/servers"
)

type LoginController struct {
	Controller
}

//用户注册  账号密码注册
func (this *LoginController) UserRegistered() {
	this.IsLogin()
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

type InstallInfo struct {
	Versions      string `json:"versions"` //安装版本
	Step          int    `json:"step"`     //安装进度
	DBType        string `json:"db_type"`  //数据存放类型， mysql,pgsql,sqlite
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDBName   string `json:"mysql_dbname"`
	PgsqlHost     string `json:"pgsql_host"`
	PgsqlUser     string `json:"pgsql_user"`
	PgsqlPassword string `json:"pgsql_password"`
	PgsqlDBName   string `json:"pgsql_dbname"`
	SqlitePath    string `json:"sqlite_path"`
	MediaType     string `json:"media_type"` //多媒体资源存放位置
	MediaPath     string `json:"media_path"`
	Aliyun        string `json:"aliyun"`     //阿里云对象存储
	Tencentyun    string `json:"tencentyun"` //腾讯云对象存储
}

//写入数据到install
func WriteInstallInfo(installInfo string) error {
	f, err := os.Create("./install.json")
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(installInfo)
	if err != nil {
		return err
	}
	return nil
}

//开始安装
func (this *LoginController) InstallStart() {
	//创建安装文件，并写入安装信息
	installInfo, err := json.Marshal(&InstallInfo{
		Versions: "v0.1",
		Step:     2,
	})

	log.Println("installInfo = ", string(installInfo), err)
	f, err := os.Create("./install.json")
	if err != nil {
		this.RetuenJson(0, 0, err)
		return
	}
	l, err := f.WriteString(string(installInfo))
	if err != nil {
		log.Println(err)
		f.Close()
		this.RetuenJson(0, 0, err)
		return
	}
	log.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Println(err)
		this.RetuenJson(0, 0, err)
		return
	}

	this.RetuenJson(0, 0, string(installInfo))
}

//测试mysql连接
func (this *LoginController) InstallMysqlTset() {
	host := this.GetString("host")
	port := this.GetString("port")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, port, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "mysql连接信息不全")
	}

	if port == "" {
		port = "3306"
	}

	//mysql 测试
	isconn, err := conn.MysqlConnTest(host, port, user, password, dbname)
	if isconn {
		this.RetuenJson(0, 0, "连接成功")
		return
	}
	this.RetuenJson(1, 0, err)
}

//mysql
func (this *LoginController) InstallMysql() {
	host := this.GetString("host")
	port := this.GetString("port")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, port, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "mysql连接信息不全")
	}

	if port == "" {
		port = "3306"
	}

	//mysql 连接测试
	isconn, _ := conn.MysqlConnTest(host, port, user, password, dbname)
	if !isconn {
		this.RetuenJson(1, 0, "连接不上mysql.")
		return
	}

	//将信息写入 install.json
	infodata := OpenInstallFile()
	infodata.Step = 3
	infodata.DBType = "mysql"
	infodata.MysqlHost = host
	infodata.MysqlPort = port
	infodata.MysqlUser = user
	infodata.MysqlPassword = password
	infodata.MysqlDBName = dbname

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}

	this.RetuenJson(0, 0, "ok")
}

//测试pgsql连接
func (this *LoginController) InstallPgsqlTset() {
	host := this.GetString("host")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "pgsql连接信息不全")
	}

	//pgsql 测试
	isconn, err := conn.PgsqlConnTest(host, user, password, dbname)
	if isconn {
		this.RetuenJson(0, 0, "连接成功")
		return
	}
	this.RetuenJson(1, 0, err)
}

//pgsql
func (this *LoginController) InstallPgsql() {
	host := this.GetString("host")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "pgsql连接信息不全")
	}

	//pgsql 测试
	isconn, _ := conn.PgsqlConnTest(host, user, password, dbname)
	if !isconn {
		this.RetuenJson(1, 0, "连接不上pgsql.")
		return
	}

	//将信息写入 install.json
	infodata := OpenInstallFile()
	infodata.Step = 3
	infodata.DBType = "pgsql"
	infodata.PgsqlHost = host
	infodata.PgsqlUser = user
	infodata.PgsqlPassword = password
	infodata.PgsqlDBName = dbname

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}

	this.RetuenJson(0, 0, "ok")
}

//sqlite
func (this *LoginController) InstallSqlite() {
	isconn, err := conn.CreateSqliteDB()
	log.Println("sqlite 连接 = ", isconn, err)
	if isconn {

		//将信息写入 install.json
		infodata := OpenInstallFile()
		infodata.Step = 3
		infodata.DBType = "sqlite"
		infodata.SqlitePath = "./db/base.db"
		installInfo, err := json.Marshal(infodata)
		if err != nil {
			this.RetuenJson(1, 0, err)
		}
		err = WriteInstallInfo(string(installInfo))
		if err != nil {
			this.RetuenJson(1, 0, err)
		}
		this.RetuenJson(0, 0, "创建成功")
		return
	}

	this.RetuenJson(1, 0, err)
}
